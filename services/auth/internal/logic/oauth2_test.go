package logic

// func setupTest(t *testing.T) (*data.Repository, *config.Config) {
// 	logger, _ := zap.NewDevelopment()
// 	sugar := logger.Sugar()

// 	cfg := &config.Config{
// 		GoogleConfig: google.Config{
// 			ClientID:     "test_client_id",
// 			ClientSecret: "test_client_secret",
// 			RedirectURL:  "http://localhost:5173/oauth2/:provider/callback",

// 		},
// 		GitHubConfig: github.Config{
// 			ClientID:     "test_client_id",
// 			ClientSecret: "test_client_secret",
// 			RedirectURL:  "http://localhost:5173/oauth2/:provider/callback",
// 		},
// 	}

// 	repo := &data.Repository{
// 		User: userdata.New(nil, sugar),
// 	}

// 	return repo, cfg
// }

// func TestInitiateOAuth2(t *testing.T) {
// 	repo, cfg := setupTest(t)
// 	l := New(repo, cfg, nil)

// 	tests := []struct {
// 		name     string
// 		provider string
// 		wantErr  bool
// 	}{
// 		{
// 			name:     "Google OAuth2 initiation",
// 			provider: string(consts.ProviderGoogle),
// 			wantErr:  false,
// 		},
// 		{
// 			name:     "GitHub OAuth2 initiation",
// 			provider: string(consts.ProviderGitHub),
// 			wantErr:  false,
// 		},
// 		{
// 			name:     "Invalid provider",
// 			provider: "invalid",
// 			wantErr:  true,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			authURL, state, err := l.InitiateOAuth2(tt.provider)
// 			if tt.wantErr {
// 				assert.Error(t, err)
// 				return
// 			}

// 			assert.NoError(t, err)
// 			assert.NotEmpty(t, state)
// 			assert.NotEmpty(t, authURL)
// 			assert.Contains(t, authURL, "state="+state)
// 		})
// 	}
// }

// func TestCompleteOAuth2(t *testing.T) {
// 	repo, cfg := setupTest(t)
// 	l := New(repo, cfg, nil)

// 	tests := []struct {
// 		name     string
// 		provider string
// 		code     string
// 		state    string
// 		wantErr  bool
// 	}{
// 		{
// 			name:     "Invalid code",
// 			provider: string(consts.ProviderGoogle),
// 			code:     "invalid_code",
// 			state:    "test_state",
// 			wantErr:  true,
// 		},
// 		{
// 			name:     "Invalid provider",
// 			provider: "invalid",
// 			code:     "test_code",
// 			state:    "test_state",
// 			wantErr:  true,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			user, payload, err := l.CompleteOAuth2(context.Background(), tt.provider, tt.code, tt.state)
// 			if tt.wantErr {
// 				assert.Error(t, err)
// 				return
// 			}

// 			assert.NoError(t, err)
// 			assert.NotNil(t, user)
// 			assert.NotNil(t, payload)
// 		})
// 	}
// }

// func TestHandleOAuthUser(t *testing.T) {
// 	repo, cfg := setupTest(t)
// 	l := New(repo, cfg, nil)

// 	tests := []struct {
// 		name      string
// 		oauthUser *dto.OAuthUser
// 		wantErr   bool
// 	}{
// 		{
// 			name: "New user creation",
// 			oauthUser: &dto.OAuthUser{
// 				ID:        "test_123",
// 				Email:     "test_new@example.com",
// 				FirstName: "Test",
// 				LastName:  "User",
// 				AvatarURL: "https://example.com/avatar.jpg",
// 				Provider:  consts.ProviderGoogle,
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name: "Existing user login",
// 			oauthUser: &dto.OAuthUser{
// 				ID:        "test_456",
// 				Email:     "test_existing@example.com",
// 				FirstName: "Existing",
// 				LastName:  "User",
// 				AvatarURL: "https://example.com/avatar2.jpg",
// 				Provider:  consts.ProviderGitHub,
// 			},
// 			wantErr: false,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			user, payload, err := l.handleOAuthUser(context.Background(), tt.oauthUser)
// 			if tt.wantErr {
// 				assert.Error(t, err)
// 				return
// 			}

// 			assert.NoError(t, err)
// 			assert.NotNil(t, user)
// 			assert.NotNil(t, payload)
// 			assert.Equal(t, tt.oauthUser.Email, user.Email)
// 			assert.Equal(t, tt.oauthUser.FirstName, user.FirstName)
// 			assert.Equal(t, tt.oauthUser.LastName, user.LastName)
// 			assert.Equal(t, tt.oauthUser.AvatarURL, user.AvatarUrl)
// 		})
// 	}
// }
