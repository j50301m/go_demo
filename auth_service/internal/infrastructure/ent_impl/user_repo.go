package ent_impl

import (
	"context"
	"hype-casino-platform/auth_service/internal/domain/aggregate"
	"hype-casino-platform/auth_service/internal/domain/entity"
	"hype-casino-platform/auth_service/internal/domain/repository"
	"hype-casino-platform/auth_service/internal/infrastructure/ent_impl/ent"
	"hype-casino-platform/auth_service/internal/infrastructure/ent_impl/ent/loginrecord"
	"hype-casino-platform/auth_service/internal/infrastructure/ent_impl/ent/user"
	"hype-casino-platform/pkg/db"
	"hype-casino-platform/pkg/enum"
	"hype-casino-platform/pkg/kgserr"
	"hype-casino-platform/pkg/kgsotel"
)

type UserRepoImpl struct {
	db db.Database
}

var _ repository.UserRepo = (*UserRepoImpl)(nil)

func NewUserRepoImpl(db db.Database) *UserRepoImpl {
	return &UserRepoImpl{
		db: db,
	}
}

func (u *UserRepoImpl) Find(ctx context.Context, id int64) (*aggregate.User, *kgserr.KgsError) {
	// Start trace
	ctx, span := kgsotel.StartTrace(ctx)
	defer span.End()

	// Get client with transaction if exists
	var client *ent.Client
	tx, ok := u.db.GetTx(ctx).(*ent.Tx)
	if ok {
		client = tx.Client()
	} else {
		client = u.db.GetConn(ctx).(*ent.Client)
	}

	// Find user
	entUser, err := client.User.Query().
		Where(user.ID(id)).
		WithAuthClients().
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			kgsErr := kgserr.New(kgserr.ResourceNotFound, "user not found", err)
			kgsotel.Error(ctx, kgsErr.Error())
			return nil, kgsErr
		}
		kgsErr := kgserr.New(kgserr.InternalServerError, "find user failed", err)
		kgsotel.Error(ctx, kgsErr.Error())
		return nil, kgsErr
	}

	// Map to enum.UserStatus
	status, kgsErr := enum.UserStatusFromInt(entUser.Status)
	if kgsErr != nil {
		kgsotel.Error(ctx, kgsErr.Error())
		return nil, kgsErr
	}

	// Map to aggregate.User
	user := &aggregate.User{
		Id:                entUser.ID,
		Account:           entUser.Account,
		Password:          entUser.Password,
		PasswordFailTimes: entUser.PasswordFailTimes,
		Status:            status,
	}
	setUserLoader(u.db, user)

	return user, nil
}

func (u *UserRepoImpl) Create(ctx context.Context, clientId int64, user *aggregate.User) (*aggregate.User, *kgserr.KgsError) {
	// Start trace
	ctx, span := kgsotel.StartTrace(ctx)
	defer span.End()

	// Get Tx from context
	tx, ok := u.db.GetTx(ctx).(*ent.Tx)
	if !ok {
		err := kgserr.New(kgserr.InternalServerError, "get tx from context failed")
		kgsotel.Error(ctx, err.Error())
		return nil, err
	}

	// Validate parameters
	kgsErr := u.validateParameters(user)
	if kgsErr != nil {
		kgsotel.Error(ctx, kgsErr.Error())
		return nil, kgsErr
	}

	// Check the client is exist
	_, err := tx.AuthClient.Get(ctx, clientId)
	if err != nil {
		kgsErr := kgserr.New(kgserr.ResourceNotFound, "client not found", err)
		kgsotel.Error(ctx, kgsErr.Error())
		return nil, kgsErr
	}

	// Create user
	entUser, err := tx.User.Create().
		SetID(user.Id).
		SetAccount(user.Account).
		SetPassword(user.Password).
		SetPasswordFailTimes(user.PasswordFailTimes).
		SetStatus(user.Status.Int()).
		SetAuthClientsID(clientId).
		Save(ctx)
	if err != nil {
		kgsErr := kgserr.New(kgserr.InternalServerError, "create user failed", err)
		kgsotel.Error(ctx, kgsErr.Error())
		return nil, kgsErr
	}

	// Map to enum.UserStatus
	status, kgsErr := enum.UserStatusFromInt(entUser.Status)
	if kgsErr != nil {
		kgsotel.Error(ctx, kgsErr.Error())
		return nil, kgsErr
	}

	// Map to aggregate.User
	newUser := &aggregate.User{
		Id:                entUser.ID,
		Account:           entUser.Account,
		Password:          entUser.Password,
		PasswordFailTimes: entUser.PasswordFailTimes,
		Status:            status,
	}
	setUserLoader(u.db, newUser)

	return newUser, nil
}

func (u *UserRepoImpl) Update(ctx context.Context, user *aggregate.User) (*aggregate.User, *kgserr.KgsError) {
	// Start trace
	ctx, span := kgsotel.StartTrace(ctx)
	defer span.End()

	// Get Tx from context
	tx, ok := u.db.GetTx(ctx).(*ent.Tx)
	if !ok {
		err := kgserr.New(kgserr.InternalServerError, "get tx from context failed")
		kgsotel.Error(ctx, err.Error())
		return nil, err
	}

	kgsErr := u.validateParameters(user)
	if kgsErr != nil {
		kgsotel.Error(ctx, kgsErr.Error())
		return nil, kgsErr
	}

	// Update user
	entUser, err := tx.User.UpdateOneID(user.Id).
		SetAccount(user.Account).
		SetPassword(user.Password).
		SetPasswordFailTimes(user.PasswordFailTimes).
		SetStatus(user.Status.Int()).
		Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			kgsErr := kgserr.New(kgserr.ResourceNotFound, "user not found", err)
			kgsotel.Error(ctx, kgsErr.Error())
			return nil, kgsErr
		}
		kgsErr := kgserr.New(kgserr.InternalServerError, "update user failed", err)
		kgsotel.Error(ctx, kgsErr.Error())
		return nil, kgsErr
	}

	// Map to enum.UserStatus
	status, kgsErr := enum.UserStatusFromInt(entUser.Status)
	if kgsErr != nil {
		kgsotel.Error(ctx, kgsErr.Error())
		return nil, kgsErr
	}

	// Map to aggregate.User
	updatedUser := &aggregate.User{
		Id:                entUser.ID,
		Account:           entUser.Account,
		Password:          entUser.Password,
		PasswordFailTimes: entUser.PasswordFailTimes,
		Status:            status,
	}
	setUserLoader(u.db, updatedUser)

	return updatedUser, nil
}

func (u *UserRepoImpl) AddLoginRecord(ctx context.Context, userId int64, loginRecord *entity.LoginRecord) (*entity.LoginRecord, *kgserr.KgsError) {
	// Start trace
	ctx, span := kgsotel.StartTrace(ctx)
	defer span.End()

	// Get Tx from context
	tx, ok := u.db.GetTx(ctx).(*ent.Tx)
	if !ok {
		err := kgserr.New(kgserr.InternalServerError, "get tx from context failed")
		kgsotel.Error(ctx, err.Error())
		return nil, err
	}

	// Create login record
	entLoginRecord, err := tx.LoginRecord.Create().
		SetBrowser(loginRecord.Browser).
		SetBrowserVer(loginRecord.BrowserVer).
		SetIP(loginRecord.Ip).
		SetOs(loginRecord.Os).
		SetPlatform(loginRecord.Platform).
		SetCountry(loginRecord.Country).
		SetCountryCode(loginRecord.CountryCode).
		SetCity(loginRecord.City).
		SetAsp(loginRecord.Asp).
		SetIsMobile(loginRecord.IsMobile).
		SetIsSuccess(loginRecord.IsSuccess).
		SetErrMessage(loginRecord.ErrMessage).
		SetUsersID(userId).
		Save(ctx)
	if err != nil {
		kgsErr := kgserr.New(kgserr.InternalServerError, "create login record failed", err)
		kgsotel.Error(ctx, kgsErr.Error())
		return nil, kgsErr
	}

	// Map to entity.LoginRecord
	return &entity.LoginRecord{
		Id:          entLoginRecord.ID,
		Browser:     entLoginRecord.Browser,
		BrowserVer:  entLoginRecord.BrowserVer,
		Ip:          entLoginRecord.IP,
		Os:          entLoginRecord.Os,
		Platform:    entLoginRecord.Platform,
		Country:     entLoginRecord.Country,
		CountryCode: entLoginRecord.CountryCode,
		City:        entLoginRecord.City,
		Asp:         entLoginRecord.Asp,
		IsMobile:    entLoginRecord.IsMobile,
		IsSuccess:   entLoginRecord.IsSuccess,
		ErrMessage:  entLoginRecord.ErrMessage,
		CreateAt:    entLoginRecord.CreatedAt,
	}, nil
}

func (u *UserRepoImpl) BindRole(ctx context.Context, userId int64, roleId int64) (*aggregate.User, *kgserr.KgsError) {
	// Start trace
	ctx, span := kgsotel.StartTrace(ctx)
	defer span.End()

	// Get Tx from context
	tx, ok := u.db.GetTx(ctx).(*ent.Tx)
	if !ok {
		err := kgserr.New(kgserr.InternalServerError, "get tx from context failed")
		kgsotel.Error(ctx, err.Error())
		return nil, err
	}

	// Check the role is exist
	role, err := tx.Role.Get(ctx, roleId)
	if err != nil {
		kgsErr := kgserr.New(kgserr.ResourceNotFound, "role is not found", err)
		kgsotel.Error(ctx, kgsErr.Error())
		return nil, kgsErr
	}

	// Bind role to user
	entUser, err := tx.User.UpdateOneID(userId).SetRoles(role).Save(ctx)
	if err != nil {
		kgsErr := kgserr.New(kgserr.InternalServerError, "binding role failed", err)
		kgsotel.Error(ctx, kgsErr.Error())
		return nil, kgsErr
	}

	// Map to enum.UserStatus
	status, kgsErr := enum.UserStatusFromInt(entUser.Status)
	if kgsErr != nil {
		kgsotel.Error(ctx, kgsErr.Error())
		return nil, kgsErr
	}

	// Map to aggregate.User
	user := &aggregate.User{
		Id:                entUser.ID,
		Account:           entUser.Account,
		Password:          entUser.Password,
		PasswordFailTimes: entUser.PasswordFailTimes,
		Status:            status,
	}
	setUserLoader(u.db, user)

	return user, nil
}

func (u *UserRepoImpl) validateParameters(user *aggregate.User) *kgserr.KgsError {
	if user == nil {
		return kgserr.New(kgserr.InvalidArgument, "user is nil")
	}

	if user.Id == 0 {
		return kgserr.New(kgserr.InvalidArgument, "user id is required")
	}

	if user.Account == "" {
		return kgserr.New(kgserr.InvalidArgument, "user account is required")
	}

	if user.Status.Int() == 0 {
		return kgserr.New(kgserr.InvalidArgument, "user status is required")
	}

	return nil
}

func setUserLoader(db db.Database, domainUser *aggregate.User) {
	domainUser.SetLoginRecordLoader(func(ctx context.Context) (*entity.LoginRecord, *kgserr.KgsError) {
		// Start trace
		ctx, span := kgsotel.StartTrace(ctx)
		defer span.End()

		// Get client with transaction if exists.
		var client *ent.Client
		tx, ok := db.GetTx(ctx).(*ent.Tx)
		if ok {
			client = tx.Client()
		} else {
			client = db.GetConn(ctx).(*ent.Client)
		}

		// Find last login record, if not found, return error.
		entLoginRecord, err := client.User.
			Query().
			Where(user.ID(domainUser.Id)).
			QueryLoginRecords().
			Order(ent.Desc(loginrecord.FieldCreatedAt)).
			First(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				kgsErr := kgserr.New(kgserr.ResourceNotFound, "login record not found", err)
				kgsotel.Error(ctx, kgsErr.Error())
				return nil, kgsErr
			}
			kgsErr := kgserr.New(kgserr.InternalServerError, "find last login record failed", err)
			kgsotel.Error(ctx, kgsErr.Error())
			return nil, kgsErr
		}

		// Map to entity.LoginRecord.
		return &entity.LoginRecord{
			Id:        entLoginRecord.ID,
			Browser:   entLoginRecord.Browser,
			Ip:        entLoginRecord.IP,
			Os:        entLoginRecord.Os,
			Country:   entLoginRecord.Country,
			City:      entLoginRecord.City,
			IsSuccess: entLoginRecord.IsSuccess,
			CreateAt:  entLoginRecord.CreatedAt,
		}, nil
	})

	domainUser.SetRoleLoader(func(ctx context.Context) (*entity.Role, *kgserr.KgsError) {
		// Start trace
		ctx, span := kgsotel.StartTrace(ctx)
		defer span.End()

		// Get client with transaction if exists.
		var client *ent.Client
		tx, ok := db.GetTx(ctx).(*ent.Tx)
		if ok {
			client = tx.Client()
		} else {
			client = db.GetConn(ctx).(*ent.Client)
		}

		// Find role, if not found, return error.
		entRole, err := client.User.
			Query().
			Where(user.ID(domainUser.Id)).
			QueryRoles().
			Only(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				kgsErr := kgserr.New(kgserr.ResourceNotFound, "role not found", err)
				kgsotel.Error(ctx, kgsErr.Error())
				return nil, kgsErr
			}
			kgsErr := kgserr.New(kgserr.InternalServerError, "find role failed", err)
			kgsotel.Error(ctx, kgsErr.Error())
			return nil, kgsErr
		}

		// Map to entity.Role
		clientType, kgsErr := enum.ClientTypeFromId(entRole.ClientType)
		if kgsErr != nil {
			kgsotel.Error(ctx, kgsErr.Error())
			return nil, kgsErr
		}
		return &entity.Role{
			Id:          entRole.ID,
			Name:        entRole.Name,
			Permissions: entRole.Permissions,
			ClientType:  clientType,
		}, nil

	})

	domainUser.SetClientLoader(func(ctx context.Context) (*aggregate.Client, *kgserr.KgsError) {
		// Start trace
		ctx, span := kgsotel.StartTrace(ctx)
		defer span.End()

		// Get client with transaction if exists.
		var client *ent.Client
		tx, ok := db.GetTx(ctx).(*ent.Tx)
		if ok {
			client = tx.Client()
		} else {
			client = db.GetConn(ctx).(*ent.Client)
		}

		// Find client, if not found, return error.
		entClient, err := client.User.
			Query().
			Where(user.ID(domainUser.Id)).
			QueryAuthClients().
			Only(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				kgsErr := kgserr.New(kgserr.ResourceNotFound, "client not found", err)
				kgsotel.Error(ctx, kgsErr.Error())
				return nil, kgsErr
			}
			kgsErr := kgserr.New(kgserr.InternalServerError, "find client failed", err)
			kgsotel.Error(ctx, kgsErr.Error())
			return nil, kgsErr
		}

		// Map to aggregate.Client
		clientType, kgsErr := enum.ClientTypeFromId(entClient.ClientType)
		if kgsErr != nil {
			kgsotel.Error(ctx, kgsErr.Error())
			return nil, kgsErr
		}
		domainClient := &aggregate.Client{
			Id:               entClient.ID,
			MerchantId:       entClient.MerchantID,
			ClientType:       clientType,
			Secret:           entClient.Secret,
			Active:           entClient.Active,
			TokenExpireSecs:  entClient.TokenExpireSecs,
			LoginFailedTimes: entClient.LoginFailedTimes,
		}
		setClientLoader(db, domainClient)
		return domainClient, nil
	})

}
