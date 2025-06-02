package email

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name   string
		config Config
		want   *Email
	}{
		{
			name: "valid config",
			config: Config{
				Host:     "smtp.gmail.com",
				Port:     "587",
				Username: "test@example.com",
				Password: "password123",
			},
			want: &Email{
				Conf: Config{
					Host:     "smtp.gmail.com",
					Port:     "587",
					Username: "test@example.com",
					Password: "password123",
				},
			},
		},
		{
			name: "empty config",
			config: Config{
				Host:     "",
				Port:     "",
				Username: "",
				Password: "",
			},
			want: &Email{
				Conf: Config{
					Host:     "",
					Port:     "",
					Username: "",
					Password: "",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.config)
			assert.NotNil(t, got)
			assert.Equal(t, tt.want.Conf, got.Conf)
			assert.NotNil(t, got.Auth)
		})
	}
}

func TestSend(t *testing.T) {
	tests := []struct {
		name    string
		email   *Email
		data    []byte
		to      []string
		wantErr bool
	}{
		{
			name: "invalid smtp configuration",
			email: &Email{
				Conf: Config{
					Host:     "invalid-host",
					Port:     "587",
					Username: "test@example.com",
					Password: "password123",
				},
			},
			data:    []byte("test email content"),
			to:      []string{"recipient@example.com"},
			wantErr: true,
		},
		{
			name: "empty recipient list",
			email: &Email{
				Conf: Config{
					Host:     "smtp.gmail.com",
					Port:     "587",
					Username: "test@example.com",
					Password: "password123",
				},
			},
			data:    []byte("test email content"),
			to:      []string{},
			wantErr: true,
		},
		{
			name: "empty email content",
			email: &Email{
				Conf: Config{
					Host:     "smtp.gmail.com",
					Port:     "587",
					Username: "test@example.com",
					Password: "password123",
				},
			},
			data:    []byte{},
			to:      []string{"recipient@example.com"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.email.Send(tt.data, tt.to...)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
