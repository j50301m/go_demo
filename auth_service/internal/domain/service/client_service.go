package service

import (
	"context"
	"hype-casino-platform/auth_service/internal/domain/aggregate"
	"hype-casino-platform/auth_service/internal/domain/entity"
	"hype-casino-platform/auth_service/internal/domain/repository"
	"hype-casino-platform/auth_service/internal/domain/vo"
	"hype-casino-platform/pkg/enum"
	"hype-casino-platform/pkg/kgscrypto"
	"hype-casino-platform/pkg/kgserr"
	"hype-casino-platform/pkg/kgsotel"
)

type ClientService struct {
	clientRepo repository.ClientRepo
	crypto     kgscrypto.KgsCrypto
}

func NewClientService(clientRepo repository.ClientRepo) *ClientService {
	return &ClientService{
		clientRepo: clientRepo,
		crypto:     kgscrypto.New(),
	}
}

func (c *ClientService) CreateClient(ctx context.Context, clientInfo vo.ClientInfo) (*aggregate.Client, *kgserr.KgsError) {
	// Start trace
	ctx, span := kgsotel.StartTrace(ctx)
	defer span.End()

	// Generate random secret
	secretByte, err := c.crypto.GenerateRandomSecret(ctx, 32)
	if err != nil {
		return nil, err
	}
	secret := c.crypto.EncodeHex(ctx, secretByte)

	client := &aggregate.Client{
		Id:               clientInfo.Id,
		MerchantId:       clientInfo.MerchantId,
		ClientType:       clientInfo.ClientType,
		LoginFailedTimes: clientInfo.LoginFailedTimes,
		TokenExpireSecs:  clientInfo.TokenExpireSecs,
		Secret:           secret,
		Active:           clientInfo.Active,
	}

	// Create client
	client, err = c.clientRepo.Create(ctx, client)
	if err != nil {
		return nil, err
	}

	// Create default roles for the client
	var roles []entity.Role
	switch clientInfo.ClientType {
	case enum.ClientType.Frontend:
		roles = entity.AllFrontendRoles
	case enum.ClientType.Backend:
		roles = entity.AllBackendRoles
	default:
		err = kgserr.New(kgserr.InvalidArgument, "invalid client type")
		kgsotel.Error(ctx, err.Error())
		return nil, err
	}

	// Bind system roles to the client
	err = c.clientRepo.BindSystemRoles(ctx, client.Id, roles...)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (c *ClientService) UpdateClient(ctx context.Context, clientInfo vo.ClientInfo) (*aggregate.Client, *kgserr.KgsError) {
	// Start trace
	ctx, span := kgsotel.StartTrace(ctx)
	defer span.End()

	// Find client
	client, err := c.clientRepo.Find(ctx, clientInfo.Id)
	if err != nil {
		return nil, err
	}

	// Check the parameters
	if clientInfo.LoginFailedTimes == 0 ||
		clientInfo.TokenExpireSecs == 0 {
		err = kgserr.New(kgserr.InvalidArgument, "login failed times and token expire seconds are required")
		kgsotel.Error(ctx, err.Error())
		return nil, err
	}

	client.LoginFailedTimes = clientInfo.LoginFailedTimes
	client.TokenExpireSecs = clientInfo.TokenExpireSecs
	client.Active = clientInfo.Active

	// Update client
	client, err = c.clientRepo.Update(ctx, client)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (c *ClientService) CreateRoles(ctx context.Context, clientId int64, roles ...entity.Role) ([]entity.Role, *kgserr.KgsError) {
	// Start trace
	ctx, span := kgsotel.StartTrace(ctx)
	defer span.End()

	// Find client
	client, err := c.clientRepo.Find(ctx, clientId)
	if err != nil {
		return nil, err
	}

	// Create a new slice for modified roles
	modifiedRoles := make([]entity.Role, len(roles))
	for i, role := range roles {
		modifiedRole := role // Create a copy of the role
		modifiedRole.ClientType = client.ClientType
		modifiedRoles[i] = modifiedRole
	}

	// Create role
	newRoles, err := c.clientRepo.CreateRoles(ctx, clientId, modifiedRoles...)
	if err != nil {
		return nil, err
	}

	return newRoles, nil
}

func (c *ClientService) DeleteRoles(ctx context.Context, clientId int64, roleIds ...int64) *kgserr.KgsError {
	// Start trace
	ctx, span := kgsotel.StartTrace(ctx)
	defer span.End()

	// Define a function to filter out system roles
	filterRoleIds := func(roleIds []int64) (*[]int64, *kgserr.KgsError) {
		// Get client by id
		client, err := c.clientRepo.Find(ctx, clientId)
		if err != nil {
			return nil, err
		}
		roles, err := client.Roles(ctx)
		if err != nil {
			return nil, err
		}
		if roles == nil {
			err = kgserr.New(kgserr.InvalidArgument, "Client roles not found")
			return nil, err
		}

		// Check if the role is not a system role, system roles cannot be deleted
		var ids []int64
		for _, roleId := range roleIds {
			if val, exists := (*roles)[roleId]; exists && !val.IsSystem() {
				ids = append(ids, roleId)
			}
		}
		return &ids, nil
	}

	ids, err := filterRoleIds(roleIds)
	if err != nil {
		return err
	}

	// Delete roles
	err = c.clientRepo.DeleteRoles(ctx, clientId, *ids...)
	if err != nil {
		return err
	}

	return nil
}

func (c *ClientService) UpdateRoles(ctx context.Context, clientId int64, roles ...entity.Role) ([]entity.Role, *kgserr.KgsError) {
	// Start trace
	ctx, span := kgsotel.StartTrace(ctx)
	defer span.End()

	// Update role
	role, err := c.clientRepo.UpdateRoles(ctx, clientId, roles...)
	if err != nil {
		return nil, err
	}

	return role, nil
}

func (c *ClientService) FindRole(ctx context.Context, clientId int64, roleId int64) (*entity.Role, *kgserr.KgsError) {
	// Start trace
	ctx, span := kgsotel.StartTrace(ctx)
	defer span.End()

	// Find role
	role, err := c.clientRepo.FindRole(ctx, clientId, roleId)
	if err != nil {
		return nil, err
	}

	return role, nil
}
