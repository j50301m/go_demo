package application_tests

import (
	"context"
	"hype-casino-platform/auth_service/internal/infrastructure/ent_impl"
	"hype-casino-platform/auth_service/internal/infrastructure/ent_impl/ent"
	"hype-casino-platform/auth_service/internal/tests"
	"hype-casino-platform/pkg/db"
	redis_cache "hype-casino-platform/pkg/db/redis"
	"hype-casino-platform/pkg/enum"
	"hype-casino-platform/pkg/kgserr"
	"hype-casino-platform/pkg/pb/gen/auth"
	"testing"

	"hype-casino-platform/auth_service/internal/application"
	domainService "hype-casino-platform/auth_service/internal/domain/service"
	"hype-casino-platform/auth_service/internal/domain/vo"

	"github.com/stretchr/testify/require"
)

func setupUserApplication() (userApp *application.UserService, db db.Database, cache db.Cache, closeFunc func()) {
	db = tests.NewMemoryDB()
	redis, closeFunc := tests.NewMemoryRedis()
	cache = redis_cache.NewRedisCache(redis)
	clientRepo := ent_impl.NewClientRepoImpl(db, cache)
	userRepo := ent_impl.NewUserRepoImpl(db)

	userService := domainService.NewUserService(clientRepo, userRepo)
	userApp = application.NewUserService(userService, db)
	return userApp, db, cache, closeFunc
}
func TestCreateUser(t *testing.T) {
	userApp, db, _, closeFunc := setupUserApplication()
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

	t.Run("Create user", func(t *testing.T) {
		req := &auth.CreateUserRequest{
			ClientId: clientInfo.Id,
			Id:       123456,
			Account:  "test",
			Password: "test",
			Status:   int32(enum.UserStatusType.Active),
		}

		res, err := userApp.CreateUser(ctx, req)
		require.Nil(t, err)
		require.NotNil(t, res)

		// Find the user
		tx, ok := db.GetConn(ctx).(*ent.Client)
		require.True(t, ok)
		entUser, e := tx.User.Get(ctx, req.Id)
		require.Nil(t, e)
		require.NotNil(t, entUser)
		require.Equal(t, req.Account, entUser.Account)
	})

	t.Run("Create user with invalid status", func(t *testing.T) {
		req := &auth.CreateUserRequest{
			ClientId: clientInfo.Id,
			Id:       123,
			Account:  "test",
			Password: "test",
			Status:   100,
		}

		res, err := userApp.CreateUser(ctx, req)
		kgsErr, ok := err.(*kgserr.KgsError)
		require.NotNil(t, err)
		require.Nil(t, res)
		require.True(t, ok)
		require.Equal(t, kgserr.InvalidArgument, kgsErr.Code().Int())
	})

	t.Run("Create user with invalid account", func(t *testing.T) {
		req := &auth.CreateUserRequest{
			ClientId: clientInfo.Id,
			Id:       123,
			Account:  "",
			Password: "test",
			Status:   int32(enum.UserStatusType.Active),
		}

		res, err := userApp.CreateUser(ctx, req)
		kgeErr := err.(*kgserr.KgsError)
		require.NotNil(t, err)
		require.Nil(t, res)
		require.Equal(t, kgserr.InvalidArgument, kgeErr.Code().Int())
	})
}

func TestUpdateUser(t *testing.T) {
	userApp, db, _, closeFunc := setupUserApplication()
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
		Id:       123456,
		Account:  "test",
		Password: "test",
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
	user, e := tx.User.Create().
		SetID(userInfo.Id).
		SetAccount(userInfo.Account).
		SetPassword(userInfo.Password).
		SetStatus(userInfo.Status.Int()).
		SetPasswordFailTimes(0).
		SetRolesID(1).
		Save(ctx)
	require.Nil(t, e)
	require.NotNil(t, user)

	// Commit the transaction
	ctx, err = db.Commit(ctx)
	require.Nil(t, err)

	t.Run("Update user", func(t *testing.T) {
		req := &auth.UpdateUserRequest{
			Id:       123456,
			Account:  "newAccount",
			Password: "newPassword",
			Status:   int32(enum.UserStatusType.Active),
		}

		res, err := userApp.UpdateUser(ctx, req)
		require.Nil(t, err)
		require.NotNil(t, res)

		// Find the user
		tx, ok := db.GetConn(ctx).(*ent.Client)
		expectedPwd := "FKiLnS9SxVtfvPnF2cEYdQ=="
		require.True(t, ok)
		entUser, e := tx.User.Get(ctx, req.Id)
		require.Nil(t, e)
		require.NotNil(t, entUser)
		require.Equal(t, req.Account, entUser.Account)
		require.Equal(t, expectedPwd, entUser.Password)
	})

	t.Run("Update user with invalid status", func(t *testing.T) {
		req := &auth.UpdateUserRequest{
			Id:       123456,
			Account:  "newAccount",
			Password: "newPassword",
			Status:   100,
		}

		res, err := userApp.UpdateUser(ctx, req)
		kgsErr, ok := err.(*kgserr.KgsError)
		require.NotNil(t, err)
		require.Nil(t, res)
		require.True(t, ok)
		require.Equal(t, kgserr.InvalidArgument, kgsErr.Code().Int())
	})
}
