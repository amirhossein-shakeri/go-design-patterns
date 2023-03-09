package singleton

// Singleton: This pattern is used to ensure a class has only one
// instance, and provides a global point of access to it. In Go,
// this can be achieved by creating a struct with only private fields,
// and providing a global function that returns a pointer to the same
// instance of the struct.
type Singleton interface {
	GetInstance() Singleton
}
