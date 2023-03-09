package singleton

type Singleton interface {
	GetInstance() Singleton
}
