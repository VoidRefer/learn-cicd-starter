package auth

import (
	"errors"
	"net/http"
	"testing"
)

type testCase struct {
	name          string
	headers       http.Header
	expectedKey   string
	expectedError error
}

func TestGetAPIKey(t *testing.T) {
	testCases := []testCase{
		{
			name:          "No Authorization Header",
			headers:       http.Header{},
			expectedError: ErrNoAuthHeaderIncluded,
		},
		{
			name: "Empty Authorization Header",
			headers: http.Header{
				"Authorization": []string{""},
			},
			expectedError: ErrNoAuthHeaderIncluded,
		},
		{
			name: "Malformed Authorization Header",
			headers: http.Header{
				"Authorization": []string{"Bearer token"},
			},
			expectedError: errors.New("malformed authorization header"),
		},
		{
			name: "Valid Authorization Header",
			headers: http.Header{
				"Authorization": []string{"ApiKey validApiKey"},
			},
			expectedKey:   "validApiKey1",
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			key, err := GetAPIKey(tc.headers)
			if err != tc.expectedError {
				if err == nil || tc.expectedError == nil || err.Error() != tc.expectedError.Error() {
					t.Errorf("Expected error %v, got %v", tc.expectedError, err)
				}
			}
			if key != tc.expectedKey {
				t.Errorf("Expected API key %s, got %s", tc.expectedKey, key)
			}
		})
	}

}
