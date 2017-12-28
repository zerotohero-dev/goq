package main

import "fmt"

// Comments that explains the algorithm are to be added.
//
// Basically, we are using a HashMap in conjunction with a
// doubly-linked list for efficient lookup.

type node struct {
	key   string
	value string
	prev  *node
	next  *node
}

func newNode() node {
	return node{"", "", nil, nil}
}

// Cache is a LRU cache.
type Cache struct {
	head     *node
	tail     *node
	cache    *(map[string]*node)
	capacity int
}

func (c *Cache) moveToHead(n *node) {
	c.removeNode(n)
	c.addNode(n)
}

func (c *Cache) popTail() *node {
	n := c.tail.prev

	if n == c.head {
		return nil
	}

	c.removeNode(n)
	return n
}

func (c *Cache) removeNode(n *node) {
	prev := n.prev
	next := n.next
	prev.next = next
	next.prev = prev
}

func (c *Cache) addNode(n *node) {
	n.prev = c.head
	n.next = c.head.next
	c.head.next.prev = n
	c.head.next = n
}

// Get gets value from cache.
func (c *Cache) Get(key string) string {
	n := (*(c.cache))[key]

	if n == nil {
		return ""
	}

	c.moveToHead(n)
	return n.value
}

// Set sets value to cache.
func (c *Cache) Set(key, value string) {
	cc := (*c.cache)
	n := cc[key]

	if n == nil {
		nn := newNode()
		nn.key = key
		nn.value = value

		cc[key] = &nn

		c.addNode(&nn)

		if len(cc) > c.capacity {
			t := c.popTail()
			delete(cc, t.key)
		}

		return
	}

	n.value = value
	c.moveToHead(n)
}

func newCache(capacity int) Cache {
	nodeMap := make(map[string]*node)

	head := node{"", "", nil, nil}
	tail := node{"", "", nil, nil}

	head.next = &tail
	tail.prev = &head

	return Cache{
		&head,
		&tail,
		&nodeMap,
		capacity}
}

func main() {
	cache := newCache(3)

	cache.Set("foo", "bar")
	cache.Set("baz", "bat")
	cache.Set("bingo", "bazinga")
	cache.Set("bobo", "dodo")
	cache.Set("foo", "bat")
	cache.Set("bat", "bar")
	cache.Set("baz", "foo")
	cache.Set("ein", "zwei")
	cache.Set("vier", "funf")

	fmt.Println(cache.Get("bingo"))
	fmt.Println(cache.Get("bobo"))
	fmt.Println(cache.Get("vier"))
}
