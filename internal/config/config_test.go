package config

import (
	"errors"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	// Original environment variables to restore later if not using t.Setenv
	// For simplicity and given Go 1.17+ context, we'll rely on t.Setenv's cleanup.
	// If t.Setenv were not available, we'd do:
	// origPort := os.Getenv("PORT")
	// origDbUrl := os.Getenv("DATABASE_URL")
	// origJwtSecret := os.Getenv("SUPABASE_JWT_SECRET")
	// t.Cleanup(func() {
	// 	os.Setenv("PORT", origPort)
	// 	os.Setenv("DATABASE_URL", origDbUrl)
	// 	os.Setenv("SUPABASE_JWT_SECRET", origJwtSecret)
	// })

	testCases := []struct {
		name                string
		envVars             map[string]string
		expectedConfig      *Config
		expectedErr         error
		unsetVars           []string // Variables to explicitly unset for the test case
		expectSpecificError string
	}{
		{
			name: "All environment variables set",
			envVars: map[string]string{
				"PORT":                "8080",
				"DATABASE_URL":        "postgres://user:pass@host:port/db",
				"SUPABASE_JWT_SECRET": "testsecret",
			},
			expectedConfig: &Config{
				Port:              "8080",
				DatabaseURL:       "postgres://user:pass@host:port/db",
				SupabaseJWTSecret: "testsecret",
			},
			expectedErr: nil,
		},
		{
			name: "PORT environment variable missing",
			envVars: map[string]string{
				"DATABASE_URL":        "postgres://user:pass@host:port/db",
				"SUPABASE_JWT_SECRET": "testsecret",
			},
			unsetVars:           []string{"PORT"},
			expectedConfig:      nil,
			expectedErr:         errors.New("PORT is not set"),
			expectSpecificError: "PORT is not set",
		},
		{
			name: "DATABASE_URL environment variable missing",
			envVars: map[string]string{
				"PORT":                "8080",
				"SUPABASE_JWT_SECRET": "testsecret",
			},
			unsetVars:           []string{"DATABASE_URL"},
			expectedConfig:      nil,
			expectedErr:         errors.New("DATABASE_URL is not set"),
			expectSpecificError: "DATABASE_URL is not set",
		},
		{
			name: "SUPABASE_JWT_SECRET environment variable missing",
			envVars: map[string]string{
				"PORT":         "8080",
				"DATABASE_URL": "postgres://user:pass@host:port/db",
			},
			unsetVars:           []string{"SUPABASE_JWT_SECRET"},
			expectedConfig:      nil,
			expectedErr:         errors.New("SUPABASE_JWT_SECRET is not set"),
			expectSpecificError: "SUPABASE_JWT_SECRET is not set",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Unset variables specified for the test case or all relevant ones
			// to ensure a clean environment for each test run.
			// t.Setenv with an empty value effectively unsets it for the test's scope.
			varsToManage := []string{"PORT", "DATABASE_URL", "SUPABASE_JWT_SECRET"}
			for _, key := range varsToManage {
				t.Setenv(key, "") // Unset initially
			}

			for _, keyToUnset := range tc.unsetVars {
				// This is slightly redundant if varsToManage already covers it,
				// but ensures clarity if tc.unsetVars has something not in varsToManage.
				t.Setenv(keyToUnset, "")
			}

			for key, value := range tc.envVars {
				t.Setenv(key, value)
			}

			config, err := Load()

			if tc.expectedErr != nil {
				if err == nil {
					t.Errorf("Expected error '%v', but got nil", tc.expectedErr)
				} else if err.Error() != tc.expectedErr.Error() {
					// Use specific error string if provided, otherwise compare error objects
					if tc.expectSpecificError != "" && err.Error() != tc.expectSpecificError {
						t.Errorf("Expected error message '%s', but got '%s'", tc.expectSpecificError, err.Error())
					} else if tc.expectSpecificError == "" && err.Error() != tc.expectedErr.Error() {
						t.Errorf("Expected error '%v', but got '%v'", tc.expectedErr, err)
					}
				}
			} else if err != nil {
				t.Errorf("Expected no error, but got '%v'", err)
			}

			if tc.expectedConfig == nil && config != nil {
				t.Errorf("Expected nil config, but got %v", config)
			}

			if tc.expectedConfig != nil {
				if config == nil {
					t.Errorf("Expected config %v, but got nil", tc.expectedConfig)
				} else {
					if config.Port != tc.expectedConfig.Port {
						t.Errorf("Expected Port '%s', but got '%s'", tc.expectedConfig.Port, config.Port)
					}
					if config.DatabaseURL != tc.expectedConfig.DatabaseURL {
						t.Errorf("Expected DatabaseURL '%s', but got '%s'", tc.expectedConfig.DatabaseURL, config.DatabaseURL)
					}
					if config.SupabaseJWTSecret != tc.expectedConfig.SupabaseJWTSecret {
						t.Errorf("Expected SupabaseJWTSecret '%s', but got '%s'", tc.expectedConfig.SupabaseJWTSecret, config.SupabaseJWTSecret)
					}
				}
			}
		})
	}
}
