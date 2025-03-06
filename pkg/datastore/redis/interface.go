package redis

type Cachable interface {
	GetDB() int
}
