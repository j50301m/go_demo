package vo

import "hype-casino-platform/pkg/enum"

type ClientInfo struct {
	Id               int64
	MerchantId       int64
	ClientType       enum.Client
	LoginFailedTimes int
	TokenExpireSecs  int
	Active           bool
}
