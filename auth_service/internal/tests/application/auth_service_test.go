package application_tests

import (
	"context"
	"hype-casino-platform/auth_service/internal/application"
	domainService "hype-casino-platform/auth_service/internal/domain/service"
	"hype-casino-platform/auth_service/internal/domain/vo"
	"hype-casino-platform/auth_service/internal/infrastructure/ent_impl"
	"hype-casino-platform/auth_service/internal/infrastructure/ent_impl/ent"
	"hype-casino-platform/auth_service/internal/infrastructure/token_helper"
	"hype-casino-platform/auth_service/internal/tests"
	"hype-casino-platform/pkg/db"
	redis_cache "hype-casino-platform/pkg/db/redis"
	"hype-casino-platform/pkg/enum"
	"hype-casino-platform/pkg/kgscrypto"
	"hype-casino-platform/pkg/kgserr"
	"hype-casino-platform/pkg/pb/gen/auth"
	"hype-casino-platform/pkg/req_analyzer"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupAuthApplication() (authApp *application.AuthService, db db.Database, cache db.Cache, closeFunc func()) {
	db = tests.NewMemoryDB()
	redis, closeFunc := tests.NewMemoryRedis()
	cache = redis_cache.NewRedisCache(redis)
	clientRepo := ent_impl.NewClientRepoImpl(db, cache)
	userRepo := ent_impl.NewUserRepoImpl(db)
	tokenHelper := token_helper.NewJwtToken()

	authService := domainService.NewAuthService(clientRepo, userRepo, cache, tokenHelper)
	clientService := domainService.NewClientService(clientRepo)
	reqAnalyzer := req_analyzer.NewReqAnalyzer()
	authApp = application.NewAuthService(authService, clientService, db, reqAnalyzer)

	return authApp, db, cache, closeFunc
}

func TestClientAuth(t *testing.T) {
	authApp, db, _, closeFunc := setupAuthApplication()
	defer closeFunc()

	ctx := context.Background()

	clientInfo := vo.ClientInfo{
		Id:               12345,
		MerchantId:       11111,
		ClientType:       enum.ClientType.Frontend,
		LoginFailedTimes: 3,
		TokenExpireSecs:  3600,
	}

	// Begin a transaction
	ctx, err := db.Begin(ctx)
	require.Nil(t, err)

	// Get the transaction
	tx, ok := db.GetTx(ctx).(*ent.Tx)
	require.True(t, ok)

	// Create a client
	client, e := tx.AuthClient.Create().
		SetID(clientInfo.Id).
		SetMerchantID(clientInfo.MerchantId).
		SetClientType(clientInfo.ClientType.Id).
		SetLoginFailedTimes(clientInfo.LoginFailedTimes).
		SetTokenExpireSecs(clientInfo.TokenExpireSecs).
		SetActive(true).
		SetSecret("secret").
		Save(ctx)
	require.Nil(t, e)
	require.NotNil(t, client)

	// Commit the transaction
	ctx, err = db.Commit(ctx)
	require.Nil(t, err)

	t.Run("Client Auth", func(t *testing.T) {
		req := &auth.ClientAuthRequest{
			ClientId: clientInfo.Id,
		}

		res, err := authApp.ClientAuth(ctx, req)
		assert.Nil(t, err)
		assert.NotNil(t, res)
	})

	t.Run("Client Auth Fail", func(t *testing.T) {
		req := &auth.ClientAuthRequest{
			ClientId: 0,
		}

		res, err := authApp.ClientAuth(ctx, req)
		assert.NotNil(t, err)
		assert.Nil(t, res)
	})
}

func TestLogin(t *testing.T) {
	authApp, db, _, closeFunc := setupAuthApplication()
	defer closeFunc()

	ctx := context.Background()

	clientInfo := vo.ClientInfo{
		Id:               12345,
		MerchantId:       11111,
		ClientType:       enum.ClientType.Frontend,
		LoginFailedTimes: 3,
		TokenExpireSecs:  3600,
	}

	userInfo := vo.UserInfo{
		Id:       12345,
		Account:  "account",
		Password: "password",
		Status:   enum.UserStatusType.Active,
	}

	// Begin a transaction
	ctx, err := db.Begin(ctx)
	require.Nil(t, err)

	// Get the transaction
	tx, ok := db.GetTx(ctx).(*ent.Tx)
	require.True(t, ok)

	// Create a client
	client, e := tx.AuthClient.Create().
		SetID(clientInfo.Id).
		SetMerchantID(clientInfo.MerchantId).
		SetClientType(clientInfo.ClientType.Id).
		SetLoginFailedTimes(clientInfo.LoginFailedTimes).
		SetTokenExpireSecs(clientInfo.TokenExpireSecs).
		SetActive(true).
		SetSecret("secret").
		Save(ctx)
	require.Nil(t, e)
	require.NotNil(t, client)

	// Create a user
	crypto := kgscrypto.New()
	hashPwd, err := crypto.HashPassword(ctx, userInfo.Password)
	require.Nil(t, err)

	user, e := tx.User.Create().
		SetID(userInfo.Id).
		SetAccount(userInfo.Account).
		SetPassword(hashPwd).
		SetPasswordFailTimes(0).
		SetStatus(userInfo.Status.Int()).
		SetRolesID(1).
		Save(ctx)
	require.Nil(t, e)
	require.NotNil(t, user)

	// Commit the transaction
	ctx, err = db.Commit(ctx)
	require.Nil(t, err)

	t.Run("Login", func(t *testing.T) {
		// Create a c token
		cToken, err := authApp.ClientAuth(ctx, &auth.ClientAuthRequest{
			ClientId: clientInfo.Id,
		})
		assert.Nil(t, err)
		assert.NotNil(t, cToken)

		req := &auth.LoginRequest{
			UserId:      userInfo.Id,
			AccessToken: cToken.AccessToken,
			Password:    "password",
		}

		res, err := authApp.Login(ctx, req)
		assert.Nil(t, err)
		assert.NotNil(t, res)
	})

	t.Run("Login without cToken", func(t *testing.T) {
		req := &auth.LoginRequest{
			UserId:      userInfo.Id,
			AccessToken: "",
			Password:    "password",
		}

		res, err := authApp.Login(ctx, req)
		assert.NotNil(t, err)
		assert.Nil(t, res)
	})

	t.Run("Login with wrong password", func(t *testing.T) {
		// Create a c token
		cToken, err := authApp.ClientAuth(ctx, &auth.ClientAuthRequest{
			ClientId: clientInfo.Id,
		})
		assert.Nil(t, err)
		assert.NotNil(t, cToken)

		req := &auth.LoginRequest{
			UserId:      userInfo.Id,
			AccessToken: cToken.AccessToken,
			Password:    "111",
		}

		res, err := authApp.Login(ctx, req)
		kgsErr := err.(*kgserr.KgsError)
		assert.NotNil(t, err)
		assert.Nil(t, res)
		assert.Equal(t, kgserr.AccountPasswordError, kgsErr.Code().Int())
	})
}

func TestValidToken(t *testing.T) {
	authApp, db, _, closeFunc := setupAuthApplication()
	defer closeFunc()

	ctx := context.Background()

	clientInfo := vo.ClientInfo{
		Id:               12345,
		MerchantId:       11111,
		ClientType:       enum.ClientType.Frontend,
		LoginFailedTimes: 3,
		TokenExpireSecs:  3600,
	}

	userInfo := vo.UserInfo{
		Id:       12345,
		Account:  "account",
		Password: "password",
		Status:   enum.UserStatusType.Active,
	}

	// Begin a transaction
	ctx, err := db.Begin(ctx)
	require.Nil(t, err)

	// Get the transaction
	tx, ok := db.GetTx(ctx).(*ent.Tx)
	require.True(t, ok)

	// Create a client
	client, e := tx.AuthClient.Create().
		SetID(clientInfo.Id).
		SetMerchantID(clientInfo.MerchantId).
		SetClientType(clientInfo.ClientType.Id).
		SetLoginFailedTimes(clientInfo.LoginFailedTimes).
		SetTokenExpireSecs(clientInfo.TokenExpireSecs).
		SetActive(true).
		SetSecret("secret").
		Save(ctx)
	require.Nil(t, e)
	require.NotNil(t, client)

	// Create a user
	crypto := kgscrypto.New()
	hashPwd, err := crypto.HashPassword(ctx, userInfo.Password)
	require.Nil(t, err)

	user, e := tx.User.Create().
		SetID(userInfo.Id).
		SetAccount(userInfo.Account).
		SetPassword(hashPwd).
		SetPasswordFailTimes(0).
		SetStatus(userInfo.Status.Int()).
		SetRolesID(1).
		Save(ctx)
	require.Nil(t, e)
	require.NotNil(t, user)

	// Commit the transaction
	ctx, err = db.Commit(ctx)
	require.Nil(t, err)

	t.Run("Valid cToken", func(t *testing.T) {
		// Create a c token
		cToken, err := authApp.ClientAuth(ctx, &auth.ClientAuthRequest{
			ClientId: clientInfo.Id,
		})
		assert.Nil(t, err)
		assert.NotNil(t, cToken)

		req := &auth.ValidTokenRequest{
			AccessToken: cToken.AccessToken,
		}

		res, err := authApp.ValidToken(ctx, req)
		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, clientInfo.Id, res.ClientId)
		assert.Equal(t, clientInfo.MerchantId, res.MerchantId)
		assert.Nil(t, res.UserAccount)
		assert.Nil(t, res.UserId)
		assert.Nil(t, res.Role)
	})

	t.Run("Valid uToken missing role", func(t *testing.T) {
		// Create a c token
		cToken, err := authApp.ClientAuth(ctx, &auth.ClientAuthRequest{
			ClientId: clientInfo.Id,
		})
		assert.Nil(t, err)
		assert.NotNil(t, cToken)

		// Create a u token
		uToken, err := authApp.Login(ctx, &auth.LoginRequest{
			UserId:      userInfo.Id,
			AccessToken: cToken.AccessToken,
			Password:    "password",
		})
		assert.Nil(t, err)

		req := &auth.ValidTokenRequest{
			AccessToken: uToken.AccessToken,
		}

		res, err := authApp.ValidToken(ctx, req)
		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, clientInfo.Id, res.ClientId)
		assert.Equal(t, clientInfo.MerchantId, res.MerchantId)
		assert.Equal(t, userInfo.Account, *res.UserAccount)
		assert.Equal(t, userInfo.Id, *res.UserId)
		assert.Nil(t, res.Role)
	})
}
