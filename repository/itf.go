package repository

type Memory interface {
	Get(key string) (val interface{}, err error)
	Set(key string, val interface{}) (err error)
}

type Redis interface {
	AssignTrackHandler(handler func(interface{}))
	Track() (err error)
	Get(key string) (val interface{}, err error)
	Set(key string, val interface{}) (err error)
}
