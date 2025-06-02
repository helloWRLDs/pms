package logic

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	userdata "pms.auth/internal/data/user"
	"pms.pkg/errs"
	"pms.pkg/transport/grpc/dto"
	"pms.pkg/utils"
)

type mockUserRepo struct {
	users map[string]userdata.User
	log   *zap.SugaredLogger
}

func (m *mockUserRepo) GetByEmail(ctx context.Context, email string) (userdata.User, error) {
	for _, u := range m.users {
		if u.Email == email {
			return u, nil
		}
	}
	return userdata.User{}, errs.ErrNotFound{
		Object: "user",
		Field:  "email",
		Value:  email,
	}
}

func (m *mockUserRepo) GetByID(ctx context.Context, id string) (userdata.User, error) {
	if user, ok := m.users[id]; ok {
		return user, nil
	}
	return userdata.User{}, errs.ErrNotFound{
		Object: "user",
		Field:  "id",
		Value:  id,
	}
}

func (m *mockUserRepo) Exists(ctx context.Context, field string, value interface{}) bool {
	if field == "email" {
		email := value.(string)
		for _, u := range m.users {
			if u.Email == email {
				return true
			}
		}
	}
	return false
}

func (m *mockUserRepo) Create(ctx context.Context, user userdata.User) error {
	m.users[user.ID] = user
	return nil
}

// func setupTest(t *testing.T) (*Logic, *mockUserRepo) {
// 	logger, _ := zap.NewDevelopment()
// 	sugar := logger.Sugar()

// 	cfg := &config.Config{
// 		JWT: jwtoken.Config{
// 			Secret: "test_secret",
// 			TTL:    int64(24 * time.Hour),
// 		},
// 	}

// 	mockRepo := &mockUserRepo{
// 		users: make(map[string]userdata.User),
// 		log:   sugar,
// 	}

// 	l := &Logic{
// 		Repo: &data.Repository{
// 			User: mockRepo,
// 		},
// 		conf: cfg,
// 		log:  sugar,
// 	}

// 	return l, mockRepo
// }

func TestLoginUser(t *testing.T) {
	tests := []struct {
		name     string
		provider *string
		creds    *dto.UserCredentials
		wantErr  bool
	}{
		{
			name:     "successful login",
			provider: nil,
			creds: &dto.UserCredentials{
				Email:    "test@example.com",
				Password: "testpass",
			},
			wantErr: false,
		},
		{
			name:     "wrong password",
			provider: nil,
			creds: &dto.UserCredentials{
				Email:    "test@example.com",
				Password: "wrongpass",
			},
			wantErr: true,
		},
		{
			name:     "user not found",
			provider: nil,
			creds: &dto.UserCredentials{
				Email:    "nonexistent@example.com",
				Password: "testpass",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := logic.LoginUser(context.Background(), tt.provider, tt.creds)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}

func TestRegisterUser(t *testing.T) {
	tests := []struct {
		name    string
		newUser *dto.NewUser
		wantErr bool
	}{
		{
			name: "successful registration",
			newUser: &dto.NewUser{
				Email:     "new@example.com",
				Password:  "newpass",
				FirstName: "New",
				LastName:  "User",
			},
			wantErr: false,
		},
		{
			name: "duplicate email",
			newUser: &dto.NewUser{
				Email:     "test@example.com",
				Password:  "testpass",
				FirstName: "Test",
				LastName:  "User",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := logic.RegisterUser(context.Background(), tt.newUser)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}

// func TestGetProfile(t *testing.T) {
// 	l, repo := setupTest(t)

// 	// Create test user
// 	testUser := userdata.User{
// 		ID:        uuid.New().String(),
// 		Email:     "test@example.com",
// 		FirstName: "Test",
// 		LastName:  utils.Ptr("User"),
// 		CreatedAt: time.Now(),
// 		UpdatedAt: time.Now(),
// 	}
// 	repo.Create(context.Background(), testUser)

// 	tests := []struct {
// 		name    string
// 		userID  string
// 		wantErr bool
// 	}{
// 		{
// 			name:    "successful profile retrieval",
// 			userID:  testUser.ID,
// 			wantErr: false,
// 		},
// 		{
// 			name:    "user not found",
// 			userID:  uuid.New().String(),
// 			wantErr: true,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			_, err := l.GetProfile(context.Background(), tt.userID)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("GetProfile() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

func Test_RegisterUser(t *testing.T) {
	newUser := &dto.NewUser{
		Email:     "admin@example.com",
		FirstName: "admin",
		LastName:  "admin",
		Password:  "admin",
	}
	created, err := logic.RegisterUser(context.Background(), newUser)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(utils.JSON(created))
}

func Test_GetUser(t *testing.T) {
	profile, err := logic.GetProfile(context.Background(), "eb306dc5-52bb-4009-88af-347b4d040718")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(utils.JSON(profile))
}

func Test_UpdateUser(t *testing.T) {
	userID := "eb306dc5-52bb-4009-88af-347b4d040718"
	user, err := logic.GetProfile(context.Background(), userID)
	if err != nil {
		t.Fatal(err)
	}
	user.FirstName = "admin2"

	updated, err := logic.UpdateUser(context.Background(), userID, user)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(utils.JSON(updated))
}

func TestCreateCompany(t *testing.T) {
	userID := "be10a73c-0927-4e3d-afe5-b4bae2e84946"

	tests := []struct {
		name       string
		newCompany *dto.NewCompany
		wantErr    bool
	}{
		{
			name: "create new company",
			newCompany: &dto.NewCompany{
				Name:        "Test Company",
				Codename:    "TEST",
				Bin:         "123456789012",
				Address:     "Test Address",
				Description: "Test Description",
			},
			wantErr: false,
		},
		{
			name: "create company with invalid user",
			newCompany: &dto.NewCompany{
				Name:        "Test Company 2",
				Codename:    "TEST2",
				Bin:         "123456789013",
				Address:     "Test Address 2",
				Description: "Test Description 2",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			created, err := logic.CreateCompany(context.Background(), userID, tt.newCompany)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, created)
			assert.Equal(t, tt.newCompany.Name, created.Name)
			assert.Equal(t, tt.newCompany.Codename, created.Codename)
			assert.Equal(t, tt.newCompany.Bin, created.Bin)
			assert.Equal(t, tt.newCompany.Address, created.Address)
			assert.Equal(t, tt.newCompany.Description, created.Description)
		})
	}
}

func TestListCompanies(t *testing.T) {
	tests := []struct {
		name      string
		filter    *dto.CompanyFilter
		wantCount int
		wantErr   bool
	}{
		{
			name: "list all companies",
			filter: &dto.CompanyFilter{
				Page:    1,
				PerPage: 10,
			},
			wantCount: 10,
			wantErr:   false,
		},
		{
			name: "list companies with name filter",
			filter: &dto.CompanyFilter{
				Page:        1,
				PerPage:     10,
				CompanyName: "Test",
			},
			wantCount: 5,
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			companies, err := logic.ListCompanies(context.Background(), tt.filter)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Len(t, companies.Items, tt.wantCount)
			assert.Equal(t, tt.filter.Page, companies.Page)
			assert.Equal(t, tt.filter.PerPage, companies.PerPage)
		})
	}
}
