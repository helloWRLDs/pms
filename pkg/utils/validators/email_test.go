package validators

import (
	"testing"

	"pms.pkg/errs"
)

func TestValidateEmail(t *testing.T) {
	tests := []struct {
		name    string
		email   string
		wantErr bool
	}{
		{
			name:    "valid email",
			email:   "test@example.com",
			wantErr: false,
		},
		{
			name:    "valid email with subdomain",
			email:   "test@sub.example.com",
			wantErr: false,
		},
		{
			name:    "valid email with numbers",
			email:   "test123@example.com",
			wantErr: false,
		},
		{
			name:    "valid email with dots",
			email:   "test.name@example.com",
			wantErr: false,
		},
		{
			name:    "invalid email - no @",
			email:   "testexample.com",
			wantErr: true,
		},
		{
			name:    "invalid email - no domain",
			email:   "test@",
			wantErr: true,
		},
		{
			name:    "invalid email - no local part",
			email:   "@example.com",
			wantErr: true,
		},
		{
			name:    "invalid email - special characters",
			email:   "test@#$@example.com",
			wantErr: true,
		},
		{
			name:    "invalid email - spaces",
			email:   "test name@example.com",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateEmail(tt.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if errObj, ok := err.(errs.ErrInvalidInput); ok {
					if errObj.Object != "email" {
						t.Errorf("ValidateEmail() error object = %v, want email", errObj.Object)
					}
					if errObj.Reason != "invalid" {
						t.Errorf("ValidateEmail() error reason = %v, want invalid", errObj.Reason)
					}
				} else {
					t.Errorf("ValidateEmail() error type = %T, want ErrInvalidInput", err)
				}
			}
		})
	}
}
