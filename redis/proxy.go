package redis

type proxy struct {
	client any
}

func NewProxy(client any) Factory {
	return &proxy{client: client}
}

func (p *proxy) MakeRedisClient() any {
	return p.client
}
