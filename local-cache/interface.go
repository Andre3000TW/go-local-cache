package localcache

type Cache interface {
	Get(key string) (any, bool)
	Set(key string, val any)
}
