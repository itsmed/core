package hooks

import (
	"context"
	"encoding/json"
	"fmt"

	"entgo.io/ent"

	"github.com/rs/zerolog/log"
	"github.com/stripe/stripe-go/v81"

	entgen "github.com/theopenlane/core/internal/ent/generated"
	"github.com/theopenlane/core/internal/ent/generated/organizationsetting"
	"github.com/theopenlane/core/pkg/events/soiree"
)

// EventID is a struct to hold the ID of an event
type EventID struct {
	ID string `json:"id"`
}

// EmitEventHook emits an event to the event pool when a mutation is performed
func EmitEventHook(pool *soiree.EventPool) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, mutation ent.Mutation) (ent.Value, error) {
			retVal, err := next.Mutate(ctx, mutation)
			if err != nil {
				return nil, err
			}

			op := mutation.Op()
			typ := mutation.Type()

			fields := mutation.Fields()

			out, err := json.Marshal(retVal)
			if err != nil {
				log.Err(err).Msg("Failed to marshal return value")
			}

			other := EventID{}
			if err := json.Unmarshal(out, &other); err != nil {
				log.Err(err).Msg("Failed to unmarshal return value")
			}

			event := soiree.NewBaseEvent(fmt.Sprintf("%s.%s", typ, op), mutation)

			event.Properties().Set("ID", other.ID)

			for _, field := range fields {
				value, exists := mutation.Field(field)
				if exists {
					event.Properties().Set(field, value)
				}
			}

			event.SetContext(context.WithoutCancel(ctx))
			event.SetClient(pool.GetClient())
			pool.Emit(event.Topic(), event)

			return retVal, err
		})
	}
}

// RegisterGlobalHooks registers global hooks for the ent client with the initialized event pool
func RegisterGlobalHooks(client *entgen.Client, pool *soiree.EventPool) {
	client.Use(EmitEventHook(pool))
}

func InitEventPool(client interface{}) *soiree.EventPool {
	return soiree.NewEventPool(soiree.WithPool(soiree.NewPondPool(soiree.WithMaxWorkers(100), soiree.WithName("ent_event_pool"))), soiree.WithErrorHandler(func(event soiree.Event, err error) error { // nolint: mnd
		log.Printf("Error encountered during event '%s': %v, with payload: %v", event.Topic(), err, event.Payload())
		return nil
	}), soiree.WithClient(client))
}

// RegisterListeners registers listeners for events
func RegisterListeners(pool *soiree.EventPool) {
	for _, event := range []string{"OrganizationSetting.OpCreate", "OrganizationSetting.OpUpdate", "OrganizationSetting.OpUpdateOne"} {
		_, err := pool.On(event, handleCustomerCreate)
		if err != nil {
			log.Error().Err(ErrFailedToRegisterListener)
		}
	}
}

// handleCustomerCreate handles the creation of a customer in Stripe
func handleCustomerCreate(event soiree.Event) error {
	props := event.Properties()
	// TODO: add funcs for default unmarshalling of props and send all props to stripe
	orgsettingID, exists := props["ID"]
	if !exists {
		log.Info().Msg("organizationSetting ID not found in event properties")
		return nil
	}

	orgSetting, err := fetchOrganizationSettingbyID(event.Context(), orgsettingID.(string), event.Client())
	if err != nil {
		log.Err(err).Msg("failed to fetch OrganizationSetting")
		return err
	}

	if orgSetting.StripeID != "" {
		stripecustomer, err := event.Client().(*entgen.Client).EntitlementManager.Client.Customers.Get(orgSetting.StripeID, nil)
		if err != nil {
			log.Err(err).Msg("Failed to retrieve Stripe customer by ID")
			return err
		}

		if stripecustomer.Email != orgSetting.BillingEmail {
			_, err := event.Client().(*entgen.Client).EntitlementManager.Client.Customers.Update(orgSetting.StripeID, &stripe.CustomerParams{
				Email: &orgSetting.BillingEmail,
			})
			if err != nil {
				log.Err(err).Msg("failed to update Stripe customer email")
				return err
			}

			log.Info().Msg("updated Stripe customer email")
		}

		if orgSetting.BillingPhone != "" && stripecustomer.Phone != orgSetting.BillingPhone {
			_, err := event.Client().(*entgen.Client).EntitlementManager.Client.Customers.Update(orgSetting.StripeID, &stripe.CustomerParams{
				Phone: &orgSetting.BillingPhone,
			})
			if err != nil {
				log.Err(err).Msg("failed to update Stripe customer phone")
				return err
			}

			log.Info().Msg("updated Stripe customer phone")
		}

		// TODO: split out address fields in ent schema
		if orgSetting.BillingAddress != "" && stripecustomer.Address.Line1 != orgSetting.BillingAddress {
			_, err := event.Client().(*entgen.Client).EntitlementManager.Client.Customers.Update(orgSetting.StripeID, &stripe.CustomerParams{
				Address: &stripe.AddressParams{
					Line1: &orgSetting.BillingAddress,
				},
			})
			if err != nil {
				log.Err(err).Msg("failed to update Stripe customer address")
				return err
			}

			log.Info().Msg("updated Stripe customer address")
		}

		subs, err := event.Client().(*entgen.Client).EntitlementManager.ListOrCreateStripeSubscriptions(orgSetting.StripeID)
		if err != nil {
			log.Err(err).Msg("failed to list or create Stripe subscriptions")
			return err
		}

		log.Info().Msgf("Created stripe subscription with ID %s", subs.ID)
	}

	billingEmail, exists := props["billing_email"]
	if exists && billingEmail != "" {
		email := billingEmail.(string)

		i := event.Client().(*entgen.Client).EntitlementManager.Client.Customers.List(&stripe.CustomerListParams{Email: &email})

		if !i.Next() {
			customer, err := event.Client().(*entgen.Client).EntitlementManager.Client.Customers.New(&stripe.CustomerParams{Email: &email})
			if err != nil {
				log.Err(err).Msg("Failed to create Stripe customer")
				return err
			}

			log.Info().Msgf("Created Stripe customer with ID: %s", customer.ID)

			if err := updateOrganizationSettingWithCustomerID(event.Context(), orgsettingID.(string), customer.ID, event.Client()); err != nil {
				log.Err(err).Msg("Failed to update OrganizationSetting with Stripe customer ID")
				return err
			}

			log.Info().Msgf("Updated OrganizationSetting with Stripe customer ID: %s", customer.ID)
		}

		if err := updateOrganizationSettingWithCustomerID(event.Context(), orgsettingID.(string), i.Customer().ID, event.Client()); err != nil {
			log.Err(err).Msg("Failed to update OrganizationSetting with Stripe customer ID")
			return err
		}

		log.Info().Msgf("Updated OrganizationSetting with Stripe customer ID: %s", i.Customer().ID)
	}

	return nil
}

// fetchOrganizationSettingbyID fetches an OrganizationSetting by ID
func fetchOrganizationSettingbyID(ctx context.Context, orgID string, client interface{}) (*entgen.OrganizationSetting, error) {
	orgsetting, err := client.(*entgen.Client).OrganizationSetting.Query().Where(organizationsetting.ID(orgID)).Only(ctx)
	if err != nil {
		log.Err(err).Msgf("Failed to fetch OrganizationSetting ID %s", orgID)

		return nil, err
	}

	return orgsetting, nil
}

// updateOrganizationSettingWithCustomerID updates an OrganizationSetting with a Stripe customer ID
func updateOrganizationSettingWithCustomerID(ctx context.Context, orgID, customerID string, client interface{}) error {
	if _, err := client.(*entgen.Client).OrganizationSetting.UpdateOneID(orgID).SetStripeID(customerID).Save(ctx); err != nil {
		log.Err(err).Msgf("Failed to update OrganizationSetting ID %s with Stripe customer ID %s", orgID, customerID)

		return err
	}

	log.Info().Msgf("Updated OrganizationSetting ID %s with Stripe customer ID %s", orgID, customerID)

	return nil
}