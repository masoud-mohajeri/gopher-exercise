package patterns

import "fmt"

type EvictionAlgo interface {
	evict(c *Cache)
}

type FIFO struct {
}

func (l *FIFO) evict(c *Cache) {
	fmt.Println("Evicting By FIFI strategy")
}

type Lru struct {
}

func (l *Lru) evict(c *Cache) {
	fmt.Println("Evicting By Lru strategy")
}

type Lfu struct {
}

func (l *Lfu) evict(c *Cache) {
	fmt.Println("Evicting By Lfu strategy")
}

type Cache struct {
	storage      map[string]string
	evictionAlgo EvictionAlgo
	capacity     int
	maxCapacity  int
}

func initCache(e EvictionAlgo) *Cache {
	storage := make(map[string]string)
	return &Cache{
		storage:      storage,
		evictionAlgo: e,
		capacity:     0,
		maxCapacity:  2,
	}
}

func (c *Cache) setEvictionAlgo(e EvictionAlgo) {
	c.evictionAlgo = e
}
func (c *Cache) add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}

	c.capacity++
	c.storage[key] = value
}

func (c *Cache) get(key string) {
	delete(c.storage, key)
}

func (c *Cache) evict() {
	c.evictionAlgo.evict(c)
	c.capacity--
}

func StrategyPattern() {
	fifo := &FIFO{}

	cache := initCache(fifo)

	cache.add("1", "2")
	cache.add("3", "4")
	cache.add("5", "6")

}
