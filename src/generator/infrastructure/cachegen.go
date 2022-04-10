package infrastructure

type CacheGen interface {
}

type CacheGenImpl struct {
}

func NewCacheGen() *CacheGenImpl {
	return &CacheGenImpl{}
}
