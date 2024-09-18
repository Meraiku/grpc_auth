package cache

type Cache interface {
	Write(key string, val any) error
	Get(key string) (any, error)
	Delete(key string) error
}
