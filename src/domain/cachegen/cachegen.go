package cachegen

type CacheGen interface {
}

type CacheGenImpl struct {
}

func NewCacheGen() *CacheGenImpl {
	return &CacheGenImpl{}
}
