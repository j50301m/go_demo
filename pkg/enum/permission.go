package enum

import "hype-casino-platform/pkg/kgserr"

type Permission struct {
	Id   int64
	Name string
}

var PermissionType = struct {
	Withdraw Permission
	Deposit  Permission
	PlayGame Permission
}{
	Withdraw: Permission{
		Id:   1,
		Name: "PlayerWithdraw",
	},
	Deposit: Permission{
		Id:   2,
		Name: "PlayerDeposit",
	},
	PlayGame: Permission{
		Id:   3,
		Name: "PlayerPlayGame",
	},
}

func PermissionById(id int64) (Permission, *kgserr.KgsError) {
	switch id {
	case PermissionType.Withdraw.Id:
		return PermissionType.Withdraw, nil
	case PermissionType.Deposit.Id:
		return PermissionType.Deposit, nil
	case PermissionType.PlayGame.Id:
		return PermissionType.PlayGame, nil
	default:
		return Permission{},
			kgserr.New(kgserr.InvalidArgument, "invalid permission id")
	}
}
