package inmemcache

type Cache struct {
	cache map[string]interface{}
}

func NewCache() *Cache {
	return &Cache{
		cache: make(map[string]interface{}),
	}
}

func (c *Cache) Set(key string, value interface{}) {
	c.cache[key] = value
}

func (c *Cache) Get(key string) interface{}{
	value, ok := c.cache[key]
	if !ok {
		return nil
	}

	return value
}

func (c *Cache) Delete(key string) {
	delete(c.cache, key)
}
