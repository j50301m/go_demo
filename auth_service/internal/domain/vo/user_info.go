package vo

import "hype-casino-platform/pkg/enum"

type UserInfo struct {
	Id       int64
	Account  string
	Password string
	Status   enum.UserStatus
}
