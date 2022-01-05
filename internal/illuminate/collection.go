package illuminate

type Collection struct {
	items map[interface{}]interface{}
}

func New() *Collection {
	return &Collection{
		items: make(map[interface{}]interface{}),
	}
}

func (c *Collection) Put(key, value interface{}) {
	c.items[key] = value
}

func (c *Collection) Contains(key interface{}) bool {
	_, ok := c.items[key]
	return ok
}

func (c *Collection) Get(key interface{}) interface{} {
	return c.items[key]
}

func (c *Collection) All() map[interface{}]interface{} {
	return c.items
}
