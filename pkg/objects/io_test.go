package objects_test

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/theopenlane/core/pkg/objects"
)

func TestDetectContentType(t *testing.T) {
	tests := []struct {
		name        string
		data        io.ReadSeeker
		expected    string
		expectError bool
	}{
		{
			name:        "Detect text/plain",
			data:        bytes.NewReader([]byte("Hello, World!")),
			expected:    "text/plain; charset=utf-8",
			expectError: false,
		},
		{
			name:        "Detect image/png",
			data:        bytes.NewReader([]byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}),
			expected:    "image/png",
			expectError: false,
		},
		{
			name:        "Detect application/octet-stream",
			data:        bytes.NewReader([]byte{0x00, 0x01, 0x02, 0x03}),
			expected:    "application/octet-stream",
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			contentType, err := objects.DetectContentType(tt.data)
			if tt.expectError {
				assert.Error(t, err)

				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.expected, contentType)
		})
	}
}

func TestComputeChecksum(t *testing.T) {
	tests := []struct {
		name        string
		data        io.ReadSeeker
		expected    string
		expectError bool
	}{
		{
			name:        "Compute checksum for text data",
			data:        bytes.NewReader([]byte("Hello, World!")),
			expected:    "ZajifYh5KDgxtmS9i38K1A==", // Precomputed MD5 checksum for "Hello, World!"
			expectError: false,
		},
		{
			name:        "Compute checksum for binary data",
			data:        bytes.NewReader([]byte{0x00, 0x01, 0x02, 0x03}),
			expected:    "N7Wa/VknJfkwXkhKXX9RaA==", // Precomputed MD5 checksum for {0x00, 0x01, 0x02, 0x03}
			expectError: false,
		},
		{
			name:        "Error on empty data",
			data:        bytes.NewReader([]byte{}),
			expected:    "1B2M2Y8AsgTpgAmY7PhCfg==", // Precomputed MD5 checksum for empty data
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checksum, err := objects.ComputeChecksum(tt.data)
			if tt.expectError {
				assert.Error(t, err)

				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.expected, checksum)
		})
	}
}

func TestReaderToSeeker(t *testing.T) {
	tests := []struct {
		name        string
		data        io.Reader
		expected    []byte
		expectError bool
	}{
		{
			name:        "Convert text data to ReadSeeker",
			data:        bytes.NewReader([]byte("Hello, World!")),
			expected:    []byte("Hello, World!"),
			expectError: false,
		},
		{
			name:        "Convert binary data to ReadSeeker",
			data:        bytes.NewReader([]byte{0x00, 0x01, 0x02, 0x03}),
			expected:    []byte{0x00, 0x01, 0x02, 0x03},
			expectError: false,
		},
		{
			name:        "nil data",
			data:        nil,
			expected:    nil,
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			readSeeker, err := objects.ReaderToSeeker(tt.data)
			if tt.expectError {
				assert.Error(t, err)

				return
			}

			assert.NoError(t, err)

			if tt.expected != nil {
				result, err := io.ReadAll(readSeeker)
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestStreamToByte(t *testing.T) {
	tests := []struct {
		name        string
		data        io.ReadSeeker
		expected    []byte
		expectError bool
	}{
		{
			name:        "Convert text data to byte slice",
			data:        bytes.NewReader([]byte("Hello, World!")),
			expected:    []byte("Hello, World!"),
			expectError: false,
		},
		{
			name:        "Convert binary data to byte slice",
			data:        bytes.NewReader([]byte{0x00, 0x01, 0x02, 0x03}),
			expected:    []byte{0x00, 0x01, 0x02, 0x03},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := objects.StreamToByte(tt.data)
			if tt.expectError {
				assert.Error(t, err)

				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestNewUploadFile(t *testing.T) {
	tests := []struct {
		name        string
		fileContent []byte
		expected    objects.FileUpload
		expectError bool
	}{
		{
			name:        "Valid text file",
			fileContent: []byte("Hello, World!"),
			expected: objects.FileUpload{
				Filename:    "testfile",
				Size:        int64(len("Hello, World!")),
				ContentType: "text/plain; charset=utf-8",
			},
			expectError: false,
		},
		{
			name:        "Valid binary file",
			fileContent: []byte{0x00, 0x01, 0x02, 0x03},
			expected: objects.FileUpload{
				Filename:    "testfile",
				Size:        int64(len([]byte{0x00, 0x01, 0x02, 0x03})),
				ContentType: "application/octet-stream",
			},
			expectError: false,
		},
		{
			name:        "Non-existent file",
			fileContent: nil,
			expected:    objects.FileUpload{},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				filePath string
				err      error
			)

			if tt.fileContent != nil {
				tmpfile, err := os.CreateTemp("", "testfile-*")
				assert.NoError(t, err)
				defer os.Remove(tmpfile.Name())

				_, err = tmpfile.Write(tt.fileContent)
				assert.NoError(t, err)

				err = tmpfile.Close()
				assert.NoError(t, err)

				filePath = tmpfile.Name()
			} else {
				filePath = "non_existent_file.txt"
			}

			file, err := objects.NewUploadFile(filePath)
			if tt.expectError {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)

			// Check if the filename contains the expected filename, as the filename is generated by the OS
			assert.Contains(t, file.Filename, tt.expected.Filename)
			assert.Equal(t, tt.expected.Size, file.Size)
			assert.Equal(t, tt.expected.ContentType, file.ContentType)
		})
	}
}