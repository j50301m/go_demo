package grpc_client

import (
	"context"
	"hype-casino-platform/frontend_api/internal/config"
	authMiddleware "hype-casino-platform/frontend_api/internal/middleware/auth"
	"hype-casino-platform/pkg/enum"
	"hype-casino-platform/pkg/kgserr"
	"hype-casino-platform/pkg/kgsotel"
	otelgrpc "hype-casino-platform/pkg/kgsotel/grpc"
	"hype-casino-platform/pkg/pb/gen/auth"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthClient struct {
	conn             *grpc.ClientConn
	authGrpcClient   auth.AuthServiceClient
	userGrpcClient   auth.UserServiceClient
	clientGrpcClient auth.ClientServiceClient
}

var _ authMiddleware.AuthClient = (*AuthClient)(nil)

// NewAuthClient creates and returns a new AuthClient instance.
// It establishes a gRPC connection to the authentication service using the address
// specified in the application configuration.
//
// The function sets up the connection with insecure credentials and a tracing middleware.
// If the connection fails, it returns an error.
//
// Returns:
//   - *AuthClient: A pointer to the newly created AuthClient instance.
//   - error: An error if the connection could not be established, nil otherwise.
func NewAuthClient() (*AuthClient, error) {
	// Get address from config
	gprcAddr := config.GetConfig().OauthUrl

	// New grpc client with own tracing middleware
	conn, err := grpc.NewClient(
		gprcAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithStatsHandler(otelgrpc.TracingMiddleware(otelgrpc.RoleClient)),
	)
	if err != nil {
		return &AuthClient{}, err
	}
	authGrpc := auth.NewAuthServiceClient(conn)
	clientGrpc := auth.NewClientServiceClient(conn)
	userGrpc := auth.NewUserServiceClient(conn)

	return &AuthClient{
		conn:             conn,
		authGrpcClient:   authGrpc,
		clientGrpcClient: clientGrpc,
		userGrpcClient:   userGrpc,
	}, nil
}

// Close terminates the gRPC connection associated with the AuthClient.
// It should be called when the client is no longer needed to free up resources.
//
// Returns:
//   - error: An error if closing the connection fails, nil otherwise.
func (c *AuthClient) Close() error {
	return c.conn.Close()
}

// ValidToken validates the provided access token and returns user information if the token is valid.
// It makes a gRPC call to the authentication service to perform the validation.
//
// Parameters:
//   - c: A context.Context for the gRPC call.
//   - token: The access token to validate as a string.
//
// Returns:
//   - *authMiddleware.UserInfo: A pointer to a UserInfo struct containing the user's information if the token is valid.
//   - *kgserr.KgsError: A pointer to a KgsError if an error occurs during validation or if the token is invalid.
//     This could be a MissingAccessToken error if the token is empty, an Unauthorized error if validation fails,
//     or an InvalidPermission error if the user has invalid permissions.
//
// The function performs the following steps:
//  1. Checks if the token is empty.
//  2. Makes a gRPC call to validate the token.
//  3. Handles any errors returned by the gRPC call.
//  4. If successful, converts the response into a UserInfo struct.
//  5. Validates the permissions returned by the authentication service.
func (a *AuthClient) ValidToken(ctx context.Context, token string) (*authMiddleware.UserInfo, *kgserr.KgsError) {
	if token == "" {
		err := kgserr.New(kgserr.MissingAccessToken, "missing access token")
		kgsotel.Error(ctx, err.Error())
		return nil, err
	}

	req := auth.ValidTokenRequest{AccessToken: token}

	// Call the oauth service to validate the token
	res, grpcErr := a.authGrpcClient.ValidToken(ctx, &req)

	// Handle the error if the err is kgserr.KgsError return the error
	// else return the unauthorized error
	if grpcErr != nil {
		if err, ok := kgserr.FromGrpcErr(grpcErr); ok {
			kgsotel.Error(ctx, err.Error())
			return nil, err
		}
		err := kgserr.New(kgserr.Unauthorized, "failed to validate token", grpcErr)
		kgsotel.Error(ctx, err.Error())
		return nil, err
	}

	// Map the permissions
	permissions := make([]enum.Permission, 0)
	if res.Role != nil {
		for i, pid := range res.Role.PermIds {
			perm, err := enum.PermissionById(pid)
			if err != nil {
				kgsotel.Error(ctx, err.Error())
				return nil, err
			}
			permissions[i] = perm
		}
	}

	userInfo := authMiddleware.NewUserInfo(
		permissions,
		res.ClientId,
		res.MerchantId,
		res.UserAccount,
		res.UserId,
	)
	return userInfo, nil
}

// ClientAuth authenticates a client with the provided client ID and returns an access token.
func (a *AuthClient) ClientAuth(ctx context.Context, clientId int64) (*auth.AuthResponse, *kgserr.KgsError) {
	if clientId == 0 {
		err := kgserr.New(kgserr.InvalidArgument, "invalid client id")
		kgsotel.Error(ctx, err.Error())
		return nil, err
	}

	res, grpcErr := a.authGrpcClient.ClientAuth(ctx, &auth.ClientAuthRequest{
		ClientId: clientId,
	})
	if grpcErr != nil {
		if err, ok := kgserr.FromGrpcErr(grpcErr); ok {
			kgsotel.Error(ctx, err.Error())
			return nil, err
		}
		err := kgserr.New(kgserr.InternalServerError, "can't found the kgsErr from grpcErr", grpcErr)
		kgsotel.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

// Login authenticates a user with the provided login information and returns an access token.
func (a *AuthClient) Login(ctx context.Context, req *auth.LoginRequest) (*auth.AuthResponse, *kgserr.KgsError) {
	// Check if the access token is empty
	if req.AccessToken == "" {
		err := kgserr.New(kgserr.MissingAccessToken, "missing access token")
		kgsotel.Error(ctx, err.Error())
		return nil, err
	}

	// Call the oauth service to login
	res, grpcErr := a.authGrpcClient.Login(ctx, req)
	if grpcErr != nil {
		if err, ok := kgserr.FromGrpcErr(grpcErr); ok {
			kgsotel.Error(ctx, err.Error())
			return nil, err
		}
		err := kgserr.New(kgserr.InternalServerError, "can't found the kgsErr from grpcErr", grpcErr)
		kgsotel.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}
