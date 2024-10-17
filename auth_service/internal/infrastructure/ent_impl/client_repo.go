package ent_impl

import (
	"context"
	"fmt"
	"hype-casino-platform/auth_service/internal/domain/aggregate"
	"hype-casino-platform/auth_service/internal/domain/entity"
	"hype-casino-platform/auth_service/internal/domain/repository"
	"hype-casino-platform/auth_service/internal/infrastructure/ent_impl/ent"
	"hype-casino-platform/auth_service/internal/infrastructure/ent_impl/ent/authclient"
	"hype-casino-platform/auth_service/internal/infrastructure/ent_impl/ent/role"
	"hype-casino-platform/pkg/db"
	"hype-casino-platform/pkg/enum"
	"hype-casino-platform/pkg/kgserr"
	"hype-casino-platform/pkg/kgsotel"
)

const (
	ClientInfoPrefix = "auth_client_info"
	RolePrefix       = "auth_role"
)

type ClientRepoImpl struct {
	db    db.Database
	cache db.Cache
}

var _ repository.ClientRepo = (*ClientRepoImpl)(nil)

// NewClientRepoImpl creates a new instance of ClientRepoImpl
func NewClientRepoImpl(db db.Database, cache db.Cache) *ClientRepoImpl {
	return &ClientRepoImpl{
		db:    db,
		cache: cache,
	}
}

func (c *ClientRepoImpl) Find(ctx context.Context, id int64) (*aggregate.Client, *kgserr.KgsError) {
	// Start trace
	ctx, span := kgsotel.StartTrace(ctx)
	defer span.End()

	// Get client with transaction if exists
	var client *ent.Client
	tx, ok := c.db.GetTx(ctx).(*ent.Tx)
	if ok {
		client = tx.Client()
	} else {
		client = c.db.GetConn(ctx).(*ent.Client)
	}

	// Fetch client from cache
	aggregateClient := &aggregate.Client{}
	key := fmt.Sprintf("%s:%d", ClientInfoPrefix, id)
	KgsErr := c.cache.GetObject(ctx, key, aggregateClient)
	if KgsErr == nil {
		setClientLoader(c.db, aggregateClient)
		return aggregateClient, nil
	}

	// Find client
	entEntity, err := client.AuthClient.Get(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			kgsErr := kgserr.New(kgserr.ResourceNotFound, "client not found", err)
			kgsotel.Error(ctx, kgsErr.Error())
			return nil, kgsErr
		}
		kgsErr := kgserr.New(kgserr.InternalServerError, "failed to find client", err)
		kgsotel.Error(ctx, err.Error())
		return nil, kgsErr
	}

	// Map client type to enum
	clientType, kgsErr := enum.ClientTypeFromId(int(entEntity.ClientType))
	if kgsErr != nil {
		kgsotel.Error(ctx, kgsErr.Error())
		return nil, kgsErr
	}

	// Create aggregate client
	authClient := &aggregate.Client{
		Id:               entEntity.ID,
		ClientType:       clientType,
		MerchantId:       entEntity.MerchantID,
		Secret:           entEntity.Secret,
		Active:           entEntity.Active,
		TokenExpireSecs:  entEntity.TokenExpireSecs,
		LoginFailedTimes: entEntity.LoginFailedTimes,
	}
	setClientLoader(c.db, authClient)

	// Save client in cache
	kgsErr = c.cache.SetObject(ctx, key, authClient, 0)
	if kgsErr != nil {
		kgsotel.Error(ctx, kgsErr.Error())
		return nil, kgsErr
	}

	return authClient, nil
}

func (c *ClientRepoImpl) Create(ctx context.Context, authClient *aggregate.Client) (*aggregate.Client, *kgserr.KgsError) {
	// Start trace
	ctx, span := kgsotel.StartTrace(ctx)
	defer span.End()

	// Get Client with transaction
	tx, ok := c.db.GetTx(ctx).(*ent.Tx)
	if !ok {
		kgsErr := kgserr.New(kgserr.InternalServerError, "transaction not found in context", nil)
		kgsotel.Error(ctx, kgsErr.Error())
		return nil, kgsErr
	}

	// Validate parameters
	if kgsErr := c.validateParameters(authClient); kgsErr != nil {
		kgsotel.Error(ctx, kgsErr.Error())
		return nil, kgsErr
	}

	// Check if  Merchant already has a client with the same type
	isExists, err := tx.AuthClient.Query().
		Where(authclient.MerchantID(authClient.MerchantId)).
		Where(authclient.ClientType(authClient.ClientType.Id)).
		Exist(ctx)
	if err != nil {
		kgsErr := kgserr.New(kgserr.InternalServerError, "failed to check if client exists", err)
		kgsotel.Error(ctx, kgsErr.Error())
		return nil, kgsErr
	}
	if isExists {
		kgsErr := kgserr.New(kgserr.ResourceIsExist, "client already exists for merchant", nil)
		kgsotel.Error(ctx, kgsErr.Error())
		return nil, kgsErr
	}

	// Create client
	entity, err := tx.AuthClient.Create().
		SetID(authClient.Id).
		SetMerchantID(authClient.MerchantId).
		SetClientType(authClient.ClientType.Id).
		SetSecret(authClient.Secret).
		SetActive(authClient.Active).
		SetTokenExpireSecs(authClient.TokenExpireSecs).
		SetLoginFailedTimes(authClient.LoginFailedTimes).
		Save(ctx)
	if err != nil {
		kgsErr := kgserr.New(kgserr.InternalServerError, "failed to create client", err)
		kgsotel.Error(ctx, kgsErr.Error())
		return nil, kgsErr
	}

	// Create aggregate client
	createdClient := &aggregate.Client{
		Id:               entity.ID,
		ClientType:       authClient.ClientType,
		MerchantId:       entity.MerchantID,
		Secret:           entity.Secret,
		Active:           entity.Active,
		TokenExpireSecs:  entity.TokenExpireSecs,
		LoginFailedTimes: entity.LoginFailedTimes,
	}
	setClientLoader(c.db, createdClient)

	// Save client in cache
	key := fmt.Sprintf("%s:%d", ClientInfoPrefix, createdClient.Id)
	kgsErr := c.cache.SetObject(ctx, key, createdClient, 0)
	if kgsErr != nil {
		kgsotel.Error(ctx, kgsErr.Error())
		return nil, kgsErr
	}

	return createdClient, nil
}

func (c *ClientRepoImpl) Update(ctx context.Context, authClient *aggregate.Client) (*aggregate.Client, *kgserr.KgsError) {
	// Start trace
	ctx, span := kgsotel.StartTrace(ctx)
	defer span.End()

	// Get Client with transaction
	tx, ok := c.db.GetTx(ctx).(*ent.Tx)
	if !ok {
		kgsErr := kgserr.New(kgserr.InternalServerError, "transaction not found in context", nil)
		kgsotel.Error(ctx, kgsErr.Error())
		return nil, kgsErr
	}

	// Update client
	entity, err := tx.AuthClient.UpdateOneID(authClient.Id).
		SetActive(authClient.Active).
		SetTokenExpireSecs(authClient.TokenExpireSecs).
		SetLoginFailedTimes(authClient.LoginFailedTimes).
		Save(ctx)
	if err != nil {
		kgsErr := kgserr.New(kgserr.InternalServerError, "failed to update client", err)
		kgsotel.Error(ctx, kgsErr.Error())
		return nil, kgsErr
	}

	// Create aggregate client
	updatedClient := &aggregate.Client{
		Id:               entity.ID,
		ClientType:       authClient.ClientType,
		MerchantId:       entity.MerchantID,
		Secret:           entity.Secret,
		Active:           entity.Active,
		TokenExpireSecs:  entity.TokenExpireSecs,
		LoginFailedTimes: entity.LoginFailedTimes,
	}
	setClientLoader(c.db, updatedClient)

	// Save client in cache
	key := fmt.Sprintf("%s:%d", ClientInfoPrefix, updatedClient.Id)
	kgsErr := c.cache.SetObject(ctx, key, updatedClient, 0)
	if kgsErr != nil {
		kgsotel.Error(ctx, kgsErr.Error())
		return nil, kgsErr
	}

	return updatedClient, nil
}

func (c *ClientRepoImpl) CreateRoles(ctx context.Context, clientId int64, r ...entity.Role) ([]entity.Role, *kgserr.KgsError) {
	// Start trace
	ctx, span := kgsotel.StartTrace(ctx)
	defer span.End()

	// Get Client with transaction
	tx, ok := c.db.GetTx(ctx).(*ent.Tx)
	if !ok {
		kgsErr := kgserr.New(kgserr.InternalServerError, "transaction not found in context", nil)
		kgsotel.Error(ctx, kgsErr.Error())
		return nil, kgsErr
	}

	// Create roles
	entRoles, err := tx.Role.MapCreateBulk(r, func(c *ent.RoleCreate, i int) {
		c.SetName(r[i].Name)
		c.SetPermissions(r[i].Permissions)
		c.AddAuthClientIDs(clientId)
		c.SetIsSystem(r[i].IsSystem())
		c.SetClientType(r[i].ClientType.Id)
	}).Save(ctx)
	if err != nil {
		kgsErr := kgserr.New(kgserr.InternalServerError, "failed to create roles", err)
		kgsotel.Error(ctx, kgsErr.Error())
		return nil, kgsErr
	}

	// Map roles
	roles := make([]entity.Role, 0, len(entRoles))
	for _, entRole := range entRoles {
		clientType, kgsErr := enum.ClientTypeFromId(entRole.ClientType)
		if kgsErr != nil {
			kgsotel.Error(ctx, kgsErr.Error())
			return nil, kgsErr
		}

		role := entity.Role{
			Id:          entRole.ID,
			Name:        entRole.Name,
			Permissions: entRole.Permissions,
			ClientType:  clientType,
		}
		roles = append(roles, role)
	}

	// Save roles in cache
	for _, role := range roles {
		key := fmt.Sprintf("%s:%d:%d", RolePrefix, clientId, role.Id)
		kgsErr := c.cache.SetObject(ctx, key, role, 0)
		if kgsErr != nil {
			kgsotel.Error(ctx, kgsErr.Error())
			return nil, kgsErr
		}
	}

	return roles, nil
}

func (c *ClientRepoImpl) BindSystemRoles(ctx context.Context, clientId int64, sysRoles ...entity.Role) *kgserr.KgsError {
	// Start trace
	ctx, span := kgsotel.StartTrace(ctx)
	defer span.End()

	// Get Client with transaction
	tx, ok := c.db.GetTx(ctx).(*ent.Tx)
	if !ok {
		kgsErr := kgserr.New(kgserr.InternalServerError, "transaction not found in context", nil)
		kgsotel.Error(ctx, kgsErr.Error())
		return kgsErr
	}

	// Check if client exists
	isExists, err := tx.AuthClient.Query().
		Where(authclient.ID(clientId)).
		Exist(ctx)
	if err != nil {
		kgsErr := kgserr.New(kgserr.InternalServerError, "failed to check if client exists", err)
		kgsotel.Error(ctx, kgsErr.Error())
		return kgsErr
	}
	if !isExists {
		kgsErr := kgserr.New(kgserr.ResourceNotFound, "client not found", nil)
		kgsotel.Error(ctx, kgsErr.Error())
		return kgsErr
	}

	// Bind roles
	sysRoleIds := make([]int64, 0, len(sysRoles))
	for _, sysRole := range sysRoles {
		sysRoleIds = append(sysRoleIds, sysRole.Id)
	}
	err = tx.Role.Update().
		Where(role.IDIn(sysRoleIds...)).
		Where(role.IsSystem(true)). // Only bind system roles
		AddAuthClientIDs(clientId).
		Exec(ctx)
	if err != nil {
		kgsErr := kgserr.New(kgserr.InternalServerError, "failed to bind system roles", err)
		kgsotel.Error(ctx, kgsErr.Error())
		return kgsErr
	}

	// Save roles in cache
	for _, role := range sysRoles {
		key := fmt.Sprintf("%s:%d:%d", RolePrefix, clientId, role.Id)
		kgsErr := c.cache.SetObject(ctx, key, role, 0)
		if kgsErr != nil {
			kgsotel.Error(ctx, kgsErr.Error())
			return kgsErr
		}
	}

	return nil
}

func (c *ClientRepoImpl) DeleteRoles(ctx context.Context, clientId int64, roleIds ...int64) *kgserr.KgsError {
	// Start trace
	ctx, span := kgsotel.StartTrace(ctx)
	defer span.End()

	// Get Client with transaction
	tx, ok := c.db.GetTx(ctx).(*ent.Tx)
	if !ok {
		kgsErr := kgserr.New(kgserr.InternalServerError, "transaction not found in context", nil)
		kgsotel.Error(ctx, kgsErr.Error())
		return kgsErr
	}

	_, err := tx.Role.Delete().
		Where(role.IDIn(roleIds...)).
		Where(role.IsSystem(false)).                             // Can not delete system roles
		Where(role.HasAuthClientsWith(authclient.ID(clientId))). // Only delete roles that belong to client
		Exec(ctx)
	if err != nil {
		kgsErr := kgserr.New(kgserr.InternalServerError, "failed to delete roles", err)
		kgsotel.Error(ctx, kgsErr.Error())
		return kgsErr
	}

	// Delete roles from cache
	for _, rid := range roleIds {
		key := fmt.Sprintf("%s:%d:%d", RolePrefix, clientId, rid)
		kgsErr := c.cache.Delete(ctx, key)
		if kgsErr != nil && kgsErr.Code().Int() != kgserr.ResourceNotFound { // Ignore if key not found
			return kgsErr
		}
	}

	return nil
}

func (c *ClientRepoImpl) UpdateRoles(ctx context.Context, clientId int64, domainRoles ...entity.Role) ([]entity.Role, *kgserr.KgsError) {
	// Start trace
	ctx, span := kgsotel.StartTrace(ctx)
	defer span.End()

	// Get Client with transaction
	tx, ok := c.db.GetTx(ctx).(*ent.Tx)
	if !ok {
		kgsErr := kgserr.New(kgserr.InternalServerError, "transaction not found in context", nil)
		kgsotel.Error(ctx, kgsErr.Error())
		return nil, kgsErr
	}

	// Check if client exists
	isExists, err := tx.AuthClient.Query().
		Where(authclient.ID(clientId)).
		Exist(ctx)
	if err != nil {
		kgsErr := kgserr.New(kgserr.InternalServerError, "failed to check if client exists", err)
		kgsotel.Error(ctx, kgsErr.Error())
		return nil, kgsErr
	}
	if !isExists {
		kgsErr := kgserr.New(kgserr.ResourceNotFound, "client not found", nil)
		kgsotel.Error(ctx, kgsErr.Error())
		return nil, kgsErr
	}

	// Update roles
	roles := make([]entity.Role, 0, len(domainRoles))
	for _, domainRole := range domainRoles {
		entRole, err := tx.Role.UpdateOneID(domainRole.Id).
			Where(role.HasAuthClientsWith(authclient.ID(clientId))). // Only update roles that belong to client
			Where(role.IsSystem(false)).                             // Can not update system roles
			SetName(domainRole.Name).
			SetPermissions(domainRole.Permissions).
			Save(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				kgsErr := kgserr.New(kgserr.ResourceNotFound, "role not found", err)
				kgsotel.Error(ctx, kgsErr.Error())
				return nil, kgsErr
			}
			kgsErr := kgserr.New(kgserr.InternalServerError, "failed to update roles", err)
			kgsotel.Error(ctx, kgsErr.Error())
			return nil, kgsErr
		}

		// Map role
		role := entity.Role{
			Id:          entRole.ID,
			Name:        entRole.Name,
			Permissions: entRole.Permissions,
		}
		roles = append(roles, role)
	}

	// Save roles in cache
	for _, role := range roles {
		key := fmt.Sprintf("%s:%d:%d", RolePrefix, clientId, role.Id)
		kgsErr := c.cache.SetObject(ctx, key, role, 0)
		if kgsErr != nil {
			kgsotel.Error(ctx, kgsErr.Error())
			return nil, kgsErr
		}
	}

	return roles, nil
}

func (c *ClientRepoImpl) FindRole(ctx context.Context, clientId int64, roleId int64) (*entity.Role, *kgserr.KgsError) {
	// Start trace
	ctx, span := kgsotel.StartTrace(ctx)
	defer span.End()

	// Fetch role from cache
	key := fmt.Sprintf("%s:%d:%d", RolePrefix, clientId, roleId)
	entityRole := &entity.Role{}
	KgsErr := c.cache.GetObject(ctx, key, entityRole)
	if KgsErr == nil {
		return entityRole, nil
	}

	// Get ent client
	var client *ent.Client
	tx, ok := c.db.GetTx(ctx).(*ent.Tx)
	if ok {
		client = tx.Client()
	} else {
		client = c.db.GetConn(ctx).(*ent.Client)
	}

	// Find role
	entRole, err := client.Role.Query().
		Where(role.ID(roleId)).
		Where(role.HasAuthClientsWith(authclient.ID(clientId))).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			kgsErr := kgserr.New(kgserr.ResourceNotFound, "role not found", err)
			kgsotel.Error(ctx, kgsErr.Error())
			return nil, kgsErr
		}
		kgsErr := kgserr.New(kgserr.InternalServerError, "failed to find role", err)
		kgsotel.Error(ctx, kgsErr.Error())
		return nil, kgsErr
	}

	// Map role
	entityRole.Id = entRole.ID
	entityRole.Name = entRole.Name
	entityRole.Permissions = entRole.Permissions

	// Save role in cache
	kgsErr := c.cache.SetObject(ctx, key, entityRole, 0)
	if kgsErr != nil {
		kgsotel.Error(ctx, kgsErr.Error())
		return nil, kgsErr
	}

	return entityRole, nil
}

func (c *ClientRepoImpl) validateParameters(authClient *aggregate.Client) *kgserr.KgsError {
	// Check if client is nil
	if authClient == nil {
		return kgserr.New(kgserr.InvalidArgument, "client is required", nil)
	}

	// Check if client secret is set
	if authClient.Secret == "" {
		return kgserr.New(kgserr.InvalidArgument, "client secret is required", nil)
	}

	// Check if client token expire is set
	if authClient.TokenExpireSecs == 0 {
		return kgserr.New(kgserr.InvalidArgument, "client token expire is required", nil)
	}

	// Check if client login failed times is set
	if authClient.LoginFailedTimes == 0 {
		return kgserr.New(kgserr.InvalidArgument, "client login failed times is required", nil)
	}

	// Check merchant id is set
	if authClient.MerchantId == 0 {
		return kgserr.New(kgserr.InvalidArgument, "merchant id is required", nil)
	}

	// Check client id is set
	if authClient.Id == 0 {
		return kgserr.New(kgserr.InvalidArgument, "client id is required", nil)
	}

	// Check client type is set
	if authClient.ClientType.Id == 0 || authClient.ClientType.String == "" {
		return kgserr.New(kgserr.InvalidArgument, "client type is required", nil)
	}

	return nil
}

func setClientLoader(db db.Database, authClient *aggregate.Client) {
	authClient.SetRolesLoader(
		func(ctx context.Context) (*map[int64]entity.Role, *kgserr.KgsError) {
			// Start trace
			ctx, span := kgsotel.StartTrace(ctx)
			defer span.End()

			// Get ent client
			var client *ent.Client
			tx, ok := db.GetTx(ctx).(*ent.Tx)
			if ok {
				client = tx.Client()
			} else {
				client = db.GetConn(ctx).(*ent.Client)
			}

			// Find roles
			entRoles, err := client.AuthClient.
				Query().
				Where(authclient.ID(authClient.Id)).
				QueryRoles().All(ctx)
			if err != nil {
				err := kgserr.New(kgserr.InternalServerError, "failed to find roles", err)
				kgsotel.Error(ctx, err.Error())
				return nil, err
			}
			if len(entRoles) == 0 {
				err := kgserr.New(kgserr.ResourceNotFound, "roles not found", nil)
				kgsotel.Error(ctx, err.Error())
				return nil, err
			}

			// Map roles
			roleMap := make(map[int64]entity.Role, len(entRoles))
			for _, entRole := range entRoles {
				roleMap[entRole.ID] = entity.Role{
					Id:          entRole.ID,
					Name:        entRole.Name,
					Permissions: entRole.Permissions,
					ClientType:  authClient.ClientType,
				}
			}

			return &roleMap, nil
		},
	)
}
