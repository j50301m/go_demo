package enum

import (
	"hype-casino-platform/pkg/kgserr"
)

type UserStatus int

var UserStatusType = struct {
	Active UserStatus
	Locked UserStatus
}{
	Active: 1,
	Locked: 2,
}

func (s UserStatus) Int() int {
	return int(s)
}

func UserStatusFromInt(val int) (UserStatus, *kgserr.KgsError) {
	switch val {
	case int(UserStatusType.Active):
		return UserStatusType.Active, nil
	case int(UserStatusType.Locked):
		return UserStatusType.Locked, nil
	default:
		return 0, kgserr.New(kgserr.InvalidArgument, "invalid user status")
	}
}
