package redis

type Factory interface {
	MakeRedisClient() any
}
