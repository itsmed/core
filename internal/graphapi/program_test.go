package graphapi_test

import (
	"context"
	"testing"
	"time"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/theopenlane/core/pkg/enums"
	"github.com/theopenlane/core/pkg/openlaneclient"
	"github.com/theopenlane/utils/ulids"
)

func (suite *GraphTestSuite) TestQueryProgram() {
	t := suite.T()

	program := (&ProgramBuilder{client: suite.client}).MustNew(testUser1.UserCtx, t)

	testCases := []struct {
		name     string
		queryID  string
		client   *openlaneclient.OpenlaneClient
		ctx      context.Context
		errorMsg string
	}{
		{
			name:    "happy path",
			queryID: program.ID,
			client:  suite.client.api,
			ctx:     testUser1.UserCtx,
		},
		{
			name:    "happy path using personal access token",
			queryID: program.ID,
			client:  suite.client.apiWithPAT,
			ctx:     context.Background(),
		},
		{
			name:     "no access, user of same org",
			queryID:  program.ID,
			client:   suite.client.api,
			ctx:      viewOnlyUser.UserCtx,
			errorMsg: notFoundErrorMsg,
		},
		{
			name:     "no access, user of different org",
			queryID:  program.ID,
			client:   suite.client.api,
			ctx:      testUser2.UserCtx,
			errorMsg: notFoundErrorMsg,
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			resp, err := tc.client.GetProgramByID(tc.ctx, tc.queryID)

			if tc.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)

			assert.Equal(t, program.ID, resp.Program.ID)
			assert.Equal(t, program.Name, resp.Program.Name)
		})
	}
}

func (suite *GraphTestSuite) TestQueryPrograms() {
	t := suite.T()

	// programs for the first organization
	(&ProgramBuilder{client: suite.client}).MustNew(testUser1.UserCtx, t)
	(&ProgramBuilder{client: suite.client}).MustNew(testUser1.UserCtx, t)

	// program created by an admin user of the first organization
	(&ProgramBuilder{client: suite.client}).MustNew(adminUser.UserCtx, t)

	// program for the other organization
	(&ProgramBuilder{client: suite.client}).MustNew(testUser2.UserCtx, t)

	testCases := []struct {
		name            string
		client          *openlaneclient.OpenlaneClient
		ctx             context.Context
		expectedResults int
		errorMsg        string
	}{
		{
			name:            "happy path, org owner should see all programs (3)",
			client:          suite.client.api,
			ctx:             testUser1.UserCtx,
			expectedResults: 3,
		},
		{
			name:            "happy path using personal access token",
			client:          suite.client.apiWithPAT,
			ctx:             context.Background(),
			expectedResults: 3,
		},
		{
			name:            "view only user has not been added to any programs",
			client:          suite.client.api,
			ctx:             viewOnlyUser.UserCtx,
			expectedResults: 0,
		},
		{
			name:            "admin user should see the program they created",
			client:          suite.client.api,
			ctx:             adminUser.UserCtx,
			expectedResults: 1,
		},
		{
			name:            "owner of the other organization should see the program they created",
			client:          suite.client.api,
			ctx:             testUser2.UserCtx,
			expectedResults: 1,
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			resp, err := tc.client.GetAllPrograms(tc.ctx)

			if tc.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			assert.Len(t, resp.Programs.Edges, tc.expectedResults)
		})
	}
}

func (suite *GraphTestSuite) TestMutationCreateProgram() {
	t := suite.T()

	startDate := time.Now().AddDate(0, 0, 1)
	endDate := time.Now().AddDate(0, 0, 360)

	testCases := []struct {
		name        string
		request     openlaneclient.CreateProgramInput
		client      *openlaneclient.OpenlaneClient
		ctx         context.Context
		expectedErr string
	}{
		{
			name: "happy path, minimal input",
			request: openlaneclient.CreateProgramInput{
				Name: "mitb program",
			},
			client: suite.client.api,
			ctx:    testUser1.UserCtx,
		},
		{
			name: "happy path, all basic input",
			request: openlaneclient.CreateProgramInput{
				Name:                 "mitb program",
				Description:          lo.ToPtr("being the best"),
				Status:               &enums.ProgramStatusInProgress,
				StartDate:            &startDate,
				EndDate:              &endDate,
				AuditorReady:         lo.ToPtr(false),
				AuditorWriteComments: lo.ToPtr(true),
				AuditorReadComments:  lo.ToPtr(true),
			},
			client: suite.client.api,
			ctx:    testUser1.UserCtx,
		},
		// {
		// 	name: "happy path, using pat",
		// 	request: openlaneclient.CreateProgramInput{
		// 		Name:        "mitb program",
		// 		Description: lo.ToPtr("being the best"),
		// 	},
		// 	client: suite.client.apiWithPAT,
		// 	ctx:    context.Background(),
		// },
		// {
		// 	name: "happy path, using api token",
		// 	request: openlaneclient.CreateProgramInput{
		// 		Name:        "mitb program",
		// 		Description: lo.ToPtr("being the best"),
		// 	},
		// 	client: suite.client.apiWithPAT,
		// 	ctx:    context.Background(),
		// },
		{
			name: "user not authorized, not enough permissions",
			request: openlaneclient.CreateProgramInput{
				Name: "mitb program",
			},
			client:      suite.client.api,
			ctx:         viewOnlyUser.UserCtx,
			expectedErr: notAuthorizedErrorMsg,
		},
		// {
		// 	name: "user not authorized, no permissions",
		// 	request: openlaneclient.CreateProgramInput{
		// 		Name:    "mitb program",
		// 		OwnerID: &testUser1.OrganizationID, // check edges this should not be allowed
		// 	},
		// 	client:      suite.client.api,
		// 	ctx:         testUser2.UserCtx,
		// 	expectedErr: notFoundErrorMsg,
		// },
		{
			name: "missing required field",
			request: openlaneclient.CreateProgramInput{
				Description: lo.ToPtr("soc2 2024"),
			},
			client:      suite.client.api,
			ctx:         testUser1.UserCtx,
			expectedErr: "value is less than the required length",
		},
	}

	for _, tc := range testCases {
		t.Run("Create "+tc.name, func(t *testing.T) {
			resp, err := tc.client.CreateProgram(tc.ctx, tc.request)
			if tc.expectedErr != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.expectedErr)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)

			// check required fields
			assert.Equal(t, tc.request.Name, resp.CreateProgram.Program.Name)

			// check optional fields
			if tc.request.Description == nil {
				assert.Empty(t, resp.CreateProgram.Program.Description)
			} else {
				assert.Equal(t, tc.request.Description, resp.CreateProgram.Program.Description)
			}

			if tc.request.Status == nil {
				assert.Equal(t, enums.ProgramStatusNotStarted, resp.CreateProgram.Program.Status)
			} else {
				assert.Equal(t, *tc.request.Status, resp.CreateProgram.Program.Status)
			}

			if tc.request.StartDate == nil {
				assert.Empty(t, resp.CreateProgram.Program.StartDate)
			} else {
				assert.WithinDuration(t, startDate, *resp.CreateProgram.Program.StartDate, 1*time.Minute)
			}

			if tc.request.EndDate == nil {
				assert.Empty(t, resp.CreateProgram.Program.EndDate)
			} else {
				assert.WithinDuration(t, endDate, *resp.CreateProgram.Program.EndDate, 1*time.Minute)
			}

			if tc.request.AuditorReady == nil {
				assert.False(t, resp.CreateProgram.Program.AuditorReady)
			} else {
				assert.Equal(t, *tc.request.AuditorReady, resp.CreateProgram.Program.AuditorReady)
			}

			if tc.request.AuditorWriteComments == nil {
				assert.False(t, resp.CreateProgram.Program.AuditorWriteComments)
			} else {
				assert.Equal(t, *tc.request.AuditorWriteComments, resp.CreateProgram.Program.AuditorWriteComments)
			}

			if tc.request.AuditorReadComments == nil {
				assert.False(t, resp.CreateProgram.Program.AuditorReadComments)
			} else {
				assert.Equal(t, *tc.request.AuditorReadComments, resp.CreateProgram.Program.AuditorReadComments)
			}
		})
	}
}

func (suite *GraphTestSuite) TestMutationUpdateProgram() {
	t := suite.T()

	program := (&ProgramBuilder{client: suite.client}).MustNew(testUser1.UserCtx, t)

	testCases := []struct {
		name        string
		request     openlaneclient.UpdateProgramInput
		client      *openlaneclient.OpenlaneClient
		ctx         context.Context
		expectedErr string
	}{
		{
			name: "happy path, update field",
			request: openlaneclient.UpdateProgramInput{
				Description: lo.ToPtr("new description"),
			},
			client: suite.client.api,
			ctx:    testUser1.UserCtx,
		},
		{
			name: "happy path, update multiple fields",
			request: openlaneclient.UpdateProgramInput{
				Status:               &enums.ProgramStatusReadyForAuditor,
				EndDate:              lo.ToPtr(time.Now().AddDate(0, 0, 30)),
				AuditorReady:         lo.ToPtr(true),
				AuditorWriteComments: lo.ToPtr(true),
				AuditorReadComments:  lo.ToPtr(true),
			},
			client: suite.client.apiWithPAT,
			ctx:    context.Background(),
		},
		{
			name: "update not allowed, not enough permissions",
			request: openlaneclient.UpdateProgramInput{
				Description: lo.ToPtr("newer description"),
			},
			client:      suite.client.api,
			ctx:         viewOnlyUser.UserCtx,
			expectedErr: notFoundErrorMsg, // programs are not visible to view only users of the organization
		},
		{
			name: "update not allowed, no permissions",
			request: openlaneclient.UpdateProgramInput{
				Description: lo.ToPtr("newer description"),
			},
			client:      suite.client.api,
			ctx:         testUser2.UserCtx,
			expectedErr: notFoundErrorMsg,
		},
	}

	for _, tc := range testCases {
		t.Run("Update "+tc.name, func(t *testing.T) {
			resp, err := tc.client.UpdateProgram(tc.ctx, program.ID, tc.request)
			if tc.expectedErr != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.expectedErr)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)

			// add checks for the updated fields if they were set in the request
			if tc.request.Description != nil {
				assert.Equal(t, *tc.request.Description, *resp.UpdateProgram.Program.Description)
			}

			if tc.request.Status != nil {
				assert.Equal(t, *tc.request.Status, resp.UpdateProgram.Program.Status)
			}

			if tc.request.StartDate != nil {
				assert.WithinDuration(t, *tc.request.StartDate, *resp.UpdateProgram.Program.StartDate, time.Minute)
			}

			if tc.request.EndDate != nil {
				assert.WithinDuration(t, *tc.request.EndDate, *resp.UpdateProgram.Program.EndDate, time.Minute)
			}

			if tc.request.AuditorReady != nil {
				assert.Equal(t, *tc.request.AuditorReady, resp.UpdateProgram.Program.AuditorReady)
			}

			if tc.request.AuditorWriteComments != nil {
				assert.Equal(t, *tc.request.AuditorWriteComments, resp.UpdateProgram.Program.AuditorWriteComments)
			}

			if tc.request.AuditorReadComments != nil {
				assert.Equal(t, *tc.request.AuditorReadComments, resp.UpdateProgram.Program.AuditorReadComments)
			}
		})
	}
}

func (suite *GraphTestSuite) TestMutationDeleteProgram() {
	t := suite.T()

	// create Programs to be deleted
	program1 := (&ProgramBuilder{client: suite.client}).MustNew(testUser1.UserCtx, t)
	program2 := (&ProgramBuilder{client: suite.client}).MustNew(testUser1.UserCtx, t)

	testCases := []struct {
		name        string
		idToDelete  string
		client      *openlaneclient.OpenlaneClient
		ctx         context.Context
		expectedErr string
	}{
		{
			name:        "not authorized, delete program",
			idToDelete:  program1.ID,
			client:      suite.client.api,
			ctx:         testUser2.UserCtx,
			expectedErr: notFoundErrorMsg,
		},
		{
			name:       "happy path, delete program",
			idToDelete: program1.ID,
			client:     suite.client.api,
			ctx:        testUser1.UserCtx,
		},
		{
			name:        "Program already deleted, not found",
			idToDelete:  program1.ID,
			client:      suite.client.api,
			ctx:         testUser1.UserCtx,
			expectedErr: "not found",
		},
		{
			name:       "happy path, delete program using personal access token",
			idToDelete: program2.ID,
			client:     suite.client.apiWithPAT,
			ctx:        context.Background(),
		},
		{
			name:        "unknown program, not found",
			idToDelete:  ulids.New().String(),
			client:      suite.client.api,
			ctx:         testUser1.UserCtx,
			expectedErr: notFoundErrorMsg,
		},
	}

	for _, tc := range testCases {
		t.Run("Delete "+tc.name, func(t *testing.T) {
			resp, err := tc.client.DeleteProgram(tc.ctx, tc.idToDelete)
			if tc.expectedErr != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.expectedErr)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			assert.Equal(t, tc.idToDelete, resp.DeleteProgram.DeletedID)
		})
	}
}