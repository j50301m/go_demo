package v1_handler

import (
	"hype-casino-platform/frontend_api/internal/config"
	"hype-casino-platform/frontend_api/internal/infrastructure/grpc_client"
	"hype-casino-platform/frontend_api/internal/model/request"
	"hype-casino-platform/pkg/kgserr"
	"hype-casino-platform/pkg/kgsotel"
	"hype-casino-platform/pkg/pb/gen/auth"
	"hype-casino-platform/pkg/responder"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authGrpc *grpc_client.AuthClient
}

func NewAuthHandler(authGrpc *grpc_client.AuthClient) *AuthHandler {
	return &AuthHandler{
		authGrpc: authGrpc,
	}
}

func (a *AuthHandler) ClientAuth(c *gin.Context) {
	ctx := c.Request.Context()

	// Get access token from cookie
	if accessToken, err := c.Cookie(config.TokenKey); err == nil {
		// Validate the access token first
		// if the token is valid, we don't need to create a new one
		_, validErr := a.authGrpc.ValidToken(ctx, accessToken)
		if validErr == nil {
			responder.Ok(nil).WithContext(c)
			return
		}
	}

	// Get the client id from Header
	val := c.GetHeader("client_id")
	if val == "" {
		kgsErr := kgserr.New(kgserr.InvalidArgument, "Client id not found in header", nil)
		kgsotel.Warn(ctx, kgsErr.Error())
		responder.Error(kgsErr).WithContext(c)
	}
	// Convert the client id to int
	clientId, err := strconv.Atoi(val)
	if err != nil {
		kgsErr := kgserr.New(kgserr.InvalidArgument, "Invalid client id", err)
		kgsotel.Warn(ctx, kgsErr.Error())
		responder.Error(kgsErr).WithContext(c)
		return
	}

	// Call the auth grpc
	res, kgsErr := a.authGrpc.ClientAuth(ctx, int64(clientId))
	if kgsErr != nil {
		responder.Error(kgsErr).WithContext(c)
		return
	}

	// Set Cookie and return the empty response
	c.SetCookie(config.TokenKey, res.AccessToken, int(res.TokenExpireSecs), "/", "", false, true)
	responder.Ok(nil).WithContext(c)
}

func (a *AuthHandler) Login(c *gin.Context) {
	ctx := c.Request.Context()

	// Get access token from cookie
	accessToken, err := c.Cookie(config.TokenKey)
	if err != nil {
		kgsErr := kgserr.New(kgserr.MissingAccessToken, "Access token is missing", err)
		kgsotel.Warn(ctx, kgsErr.Error())
		responder.Error(kgsErr).WithContext(c)
		return
	}

	// Get the login request from the body
	var loginRequest request.LoginRequest
	if err := c.ShouldBindBodyWithJSON(&loginRequest); err != nil {
		// New a Kgserr with InvalidArgument error code and log it berfore returning
		kgsErr := kgserr.New(kgserr.InvalidArgument, "Invalid request", err)
		kgsotel.Warn(ctx, kgsErr.Error())
		responder.Error(kgsErr).WithContext(c)
		return
	}

	// Get user agent and IP address from th header
	userAgent := c.GetHeader("User-Agent")
	ip := c.ClientIP()

	// TODO: Get user Id from the user grpc service

	// Authenticate the user
	loginInfo := &auth.LoginRequest{
		UserAgent:   userAgent,
		Ip:          ip,
		UserId:      2, // TODO: Get user Id from the user grpc service
		AccessToken: accessToken,
		Password:    loginRequest.Password,
	}
	res, kgsErr := a.authGrpc.Login(ctx, loginInfo)
	if kgsErr != nil {
		responder.Error(kgsErr).WithContext(c)
		return
	}

	// Set Cookie and return the empty response
	c.SetCookie(config.TokenKey, res.AccessToken, int(res.TokenExpireSecs), "/", "", false, true)
	responder.Ok(nil).WithContext(c)
}
