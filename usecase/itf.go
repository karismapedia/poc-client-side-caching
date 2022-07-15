package usecase

type Usecase interface {
	Get(key string) (val interface{}, err error)
	Set(key string, val interface{}) (err error)
	Refresh(key string) (err error)
}
