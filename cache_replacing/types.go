package cachereplacing

// cache node with key and value
type data struct {
	key   int
	value int
}

// LRUKCache history count struct
type historyCounter struct {
	key     int
	visited uint
}
