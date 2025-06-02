package validators

import (
	"testing"

	"pms.pkg/errs"
)

func TestValidatePassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		wantErr  bool
		errType  error
	}{
		{
			name:     "valid password",
			password: "Password123",
			wantErr:  false,
		},
		{
			name:     "too short",
			password: "Pass1",
			wantErr:  true,
			errType: errs.ErrInvalidInput{
				Object: "password",
				Reason: "password should contain minimum 8 characters",
			},
		},
		{
			name:     "no letters",
			password: "12345678",
			wantErr:  true,
			errType: errs.ErrInvalidInput{
				Object: "password",
				Reason: "Password should contain at least 1 letter",
			},
		},
		{
			name:     "no numbers",
			password: "Password",
			wantErr:  true,
			errType: errs.ErrInvalidInput{
				Object: "password",
				Reason: "Password should contain at least 1 number",
			},
		},
		{
			name:     "special characters",
			password: "Pass@123",
			wantErr:  true,
			errType: errs.ErrInvalidInput{
				Object: "password",
				Reason: "password should contain minimum 8 characters",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidatePassword(tt.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidatePassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if errObj, ok := err.(errs.ErrInvalidInput); ok {
					if errObj.Object != tt.errType.(errs.ErrInvalidInput).Object {
						t.Errorf("ValidatePassword() error object = %v, want %v", errObj.Object, tt.errType.(errs.ErrInvalidInput).Object)
					}
					if errObj.Reason != tt.errType.(errs.ErrInvalidInput).Reason {
						t.Errorf("ValidatePassword() error reason = %v, want %v", errObj.Reason, tt.errType.(errs.ErrInvalidInput).Reason)
					}
				} else {
					t.Errorf("ValidatePassword() error type = %T, want %T", err, tt.errType)
				}
			}
		})
	}
}
