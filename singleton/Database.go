package singleton

var instance *Database

type Database struct {
	// anything could be here like DB_URI or whatever else ...
}

func GetInstance() *Database {
	if instance != nil {
		return instance
	}
	// create new instance
	instance = &Database{}
	return instance
}
