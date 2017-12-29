package main

import "fmt"

// LRU (Last Recently Used) Cache
//
// An LRU cache is a data structure that has an upper limit on
// the number of items it can hold. Once you reach that limit,
// and want to add a new item, the least-recently-used item will
// be discarded to make space for the new item.
//
// See [Cache Replacement Policies](https://en.wikipedia.org/wiki/Cache_replacement_policies)
// for a discussion of different cache replacement strategies including LRU.
//
// In our implementation we’ll use a `map[string]node` where `node` is a struct that we
// will use to store cache data.
//
// We will join `node`s in a [doubly-linked list](https://en.wikipedia.org/wiki/Doubly_linked_list)
// so that it will be efficient to keep track of the least-recently-used item: We move
// the recently-used items to the “head” of the linked list, and that (by extension) results
// in the least-recently-used item to remain at the “tail” of the list.
//
// When adding a new item, if the cache is full (i.e. the number of items in it is equal to
// the `capacity` of the `Cache`), then we pop the item at the “tail” of the linked list (which
// is the least-recently-used item).
//
// It is also possible to create a simpler solution by keeping a “sort order” on each of
// the `node` and updating all “sort oder”s whenever you `get`, or `set` an item
// to/from the cache. This, however, will be computationally expensive — especially if the
// size of the list is large. That’s because changing the position of an item in a
//  linked-list is a constant-time operation,whereas updating all the sort orders in
// a collection is linear time at best.

// This is the struct that we store the data of each cache entry
// `value` is the data we store. — We are using `string` for the sake of simplicity;
// however we could have used any other type (even an `interface{}`) as well.
type node struct {
	key   string
	value string
	prev  *node
	next  *node
}

// Helper function to instantiate a new node with no value.
func newNode() node {
	return node{"", "", nil, nil}
}

// Cache is a LRU cache.
// It has to pseudo nodes to help keep track of head
// and tail of it easily.
//
// For example…
//
// (head) <-> (tail) : empty cache.
//
// (head) <-> (1) <-> (2) <-> (3) <-> (tail) : A cache with three nodes
//   where (1) is the actual head node and (3) is the actual tail node.
//   We will use (head) and (tail) to reach them with ease.
//     (i.e. (head).next will be the actual head;
//      (tail)->prev will be the actual tail)
type Cache struct {
	head *node
	tail *node

	// cache is a reference to the actual `map[string]*node` cache
	// structure that this `Cache` type manges.
	//
	// Cache only does book-keeping of capacity, pseudo head, and pseudo tail
	// nodes, whereas Cache.cache does the actual legwork.
	cache *(map[string]*node)

	capacity int
}

// Removes and returns the tail node.
// Returns `nil` if there is no tail node.
func (c *Cache) popTail() *node {
	n := c.tail.prev

	if n == c.head {
		return nil
	}

	c.removeNode(n)
	return n
}

// Removing a node is essentially linking the node
// before and the node after together.
func (c *Cache) removeNode(n *node) {
	prev := n.prev
	next := n.next
	prev.next = next
	next.prev = prev
}

// Adding a node means, the pseudo head (head) now
// points the the new node, and the new node points to
// the old head node of the linked list.
func (c *Cache) addNode(n *node) {
	n.prev = c.head
	n.next = c.head.next
	c.head.next.prev = n
	c.head.next = n
}

// Moving a node to head means we first break its links,
// and then add it to the cache as if it was a new node.
func (c *Cache) moveToHead(n *node) {
	c.removeNode(n)
	c.addNode(n)
}

// Get gets value (of the `node`) from the cache.
func (c *Cache) Get(key string) string {

	// Get the node from the inner cache.
	n := (*(c.cache))[key]

	// If node not found, return the zero value of
	// `string` which is an empty string.
	if n == nil {
		return ""
	}

	// Value found.
	// Since we have accessed the value, move it to the
	// head of the underlying linked list.
	c.moveToHead(n)

	// Return the value of the cache item.
	return n.value
}

// Set sets value to cache.
func (c *Cache) Set(key, value string) {

	// cc is the internal cache.
	cc := (*c.cache)
	n := cc[key]

	// If no cache entry, create a new one.
	if n == nil {
		nn := newNode()
		nn.key = key
		nn.value = value

		cc[key] = &nn

		c.addNode(&nn)

		// If we are out of capacity, pop the tail
		// and reduce the size of the inner cache by
		// deleting the entry.
		if len(cc) > c.capacity {
			t := c.popTail()
			delete(cc, t.key)
		}

		return
	}

	// We already have a cache entry, so update its value.
	n.value = value

	// Since we’ve accessed to the cache entry, we should
	// move it to the head of the linked list.
	c.moveToHead(n)
}

// Helper function to create a new cache with a set capacity.
func newCache(capacity int) Cache {
	nodeMap := make(map[string]*node)

	// Two dummy nodes for the pseudo head and pseudo tail.
	// These nodes never change “after” they are created.
	// We use them as “helper”s to simplify operations.
	head := node{"", "", nil, nil}
	tail := node{"", "", nil, nil}

	// At its zero state, the pseudo head and the pseudo tail
	// point at each other.
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
