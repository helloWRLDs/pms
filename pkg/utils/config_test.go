package utils

import (
	"os"
	"testing"
)

type TestConfig struct {
	Host     string `env:"TEST_HOST"`
	Port     int    `env:"TEST_PORT"`
	Debug    bool   `env:"TEST_DEBUG"`
	Optional string `env:"TEST_OPTIONAL"`
}

func TestLoadConfig(t *testing.T) {
	envContent := `TEST_HOST=localhost
					TEST_PORT=8080
					TEST_DEBUG=true
					TEST_OPTIONAL=value`

	tmpFile, err := os.CreateTemp("", "test.env")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.WriteString(envContent); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tmpFile.Close()

	tests := []struct {
		name        string
		envPath     string
		wantErr     bool
		checkConfig func(t *testing.T, cfg TestConfig)
	}{
		{
			name:    "successful config load",
			envPath: tmpFile.Name(),
			wantErr: false,
			checkConfig: func(t *testing.T, cfg TestConfig) {
				if cfg.Host != "localhost" {
					t.Errorf("Host = %v, want localhost", cfg.Host)
				}
				if cfg.Port != 8080 {
					t.Errorf("Port = %v, want 8080", cfg.Port)
				}
				if !cfg.Debug {
					t.Errorf("Debug = %v, want true", cfg.Debug)
				}
				if cfg.Optional != "value" {
					t.Errorf("Optional = %v, want value", cfg.Optional)
				}
			},
		},
		{
			name:    "non-existent env file",
			envPath: "non_existent.env",
			wantErr: true,
			checkConfig: func(t *testing.T, cfg TestConfig) {

				if cfg.Host != "" {
					t.Errorf("Host = %v, want empty string", cfg.Host)
				}
				if cfg.Port != 0 {
					t.Errorf("Port = %v, want 0", cfg.Port)
				}
				if cfg.Debug {
					t.Errorf("Debug = %v, want false", cfg.Debug)
				}
				if cfg.Optional != "" {
					t.Errorf("Optional = %v, want empty string", cfg.Optional)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg, err := LoadConfig[TestConfig](tt.envPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			tt.checkConfig(t, cfg)
		})
	}
}
