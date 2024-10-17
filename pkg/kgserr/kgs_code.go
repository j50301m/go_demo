package kgserr

import (
	"net/http"

	"google.golang.org/grpc/codes"
)

// KgsCode is a custom error code type.
// It consists of 7 digits: the first 3 represent the HTTP standard status code,
// and the last 4 are our custom error codes.
// The last 4 digits start from 0000 and increase sequentially (skipping numbers is prohibited).
type KgsCode int

const (
	// 200 status code from here
	OK = 200_0000 // 成功

	// 400 status code from here
	BadRequest      = 400_0000 // 請求錯誤
	InvalidArgument = 400_0001 // 請求參數錯誤
	ThirdPartyError = 400_0002 // oauth provider error

	// 401 status code from here
	Unauthorized         = 401_0000 // 未授權
	MissingAccessToken   = 401_0001 // 沒有帶Token
	AccountLocked        = 401_0002 // 帳號被鎖定
	AccountPasswordError = 401_0003 // 密碼錯誤
	TokenExpired         = 401_0004 // Token過期
	ClientInactive       = 401_0005 // 客戶端未啟用

	// 403 status code from here
	Forbidden    = 403_0000 // 禁止訪問
	NoPermission = 403_0001 // 沒有權限
	NoRole       = 403_0002 // 沒有角色

	// 404 status code from here
	ResponseNotFound = 404_0000 // 沒有Response
	ResourceNotFound = 404_0001 // 找不到資源

	// 409 status code from here
	Conflict        = 409_0000 // 衝突
	ResourceIsExist = 409_0001 // 資源已存在

	// 429 status code from here
	TooManyRequests = 429_0000 // 請求過多

	// 500 status code from here
	InternalServerError = 500_0000 // 内部錯誤
	InvalidPermission   = 500_0001 // 無效的權限

	// 501 status code from here
	NotImplemented = 501_0000 // 功能未實現
)

// HttpCode returns the standard HTTP status code.
func (c KgsCode) HttpCode() int {

	// Get the http code from the KgsCode
	httpCode := int(c) / 10000

	// Check if the http code is valid
	if http.StatusText(httpCode) == "" {
		return http.StatusInternalServerError
	}

	return httpCode
}

// GrpcCode returns the corresponding gRPC status code.
func (c KgsCode) GrpcCode() codes.Code {
	switch c.HttpCode() {
	case http.StatusOK:
		return codes.OK
	case http.StatusBadRequest:
		return codes.InvalidArgument
	case http.StatusUnauthorized:
		return codes.Unauthenticated
	case http.StatusForbidden:
		return codes.PermissionDenied
	case http.StatusNotFound:
		return codes.NotFound
	case http.StatusConflict:
		return codes.AlreadyExists
	case http.StatusTooManyRequests:
		return codes.ResourceExhausted
	case http.StatusInternalServerError:
		return codes.Internal
	case http.StatusNotImplemented:
		return codes.Unimplemented
	default:
		return codes.Unknown
	}
}

// Int returns the integer value of the KgsCode.
func (c KgsCode) Int() int {
	return int(c)
}
