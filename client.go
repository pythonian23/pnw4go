package pnw4go

func NewClient(key string) Client {
	return Client{
		client{
			key: key,
		},
	}
}

type Client struct {
	client
}

type client struct {
	key string
}
