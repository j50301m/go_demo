package enum

import "hype-casino-platform/pkg/kgserr"

type Client struct {
	Id     int
	String string
}

var ClientType = struct {
	Frontend Client
	Backend  Client
}{
	Frontend: Client{
		Id:     1,
		String: "Frontend",
	},
	Backend: Client{
		Id:     2,
		String: "Backend",
	},
}

func ClientTypeFromId(id int) (Client, *kgserr.KgsError) {
	switch id {
	case 1:
		return ClientType.Frontend, nil
	case 2:
		return ClientType.Backend, nil
	}
	return Client{}, kgserr.New(kgserr.InvalidArgument, "invalid client type")
}
