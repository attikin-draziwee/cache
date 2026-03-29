package cache

import "errors"

type Cache struct {
	mem map[string]any
}

func New() Cache {
	return Cache{
		mem: make(map[string]any),
	}
}

func (c *Cache) Set(key string, value any) {
	c.mem[key] = value
}

func (c Cache) Get(key string) (value any, err error) {
	if result, err := c.mem[key]; !err {
		return nil, errors.New(key + " not found.")
	} else {
		return result, nil
	}
}

func (c *Cache) Delete(key string) (value any, err error) {
	if result, err := c.mem[key]; !err {
		return nil, errors.New(key + " not found.")
	} else {
		delete(c.mem, key)
		return result, nil
	}
}
