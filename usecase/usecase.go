package usecase

import (
	"log"

	"github.com/karismapedia/poc-client-side-caching/repository"
)

type usecase struct {
	memory repository.Memory
	redis  repository.Redis
}

func Init(m repository.Memory, r repository.Redis) (u usecase) {
	u = usecase{m, r}
	return
}

func (u *usecase) Get(key string) (val interface{}, err error) {
	val, err = u.memory.Get(key)
	if err != nil {
		return
	}

	// means we got value from memory cache
	if val != nil {
		log.Print("got value from memory cache")
		return
	}

	val, err = u.redis.Get(key)
	if err != nil {
		return
	}

	// means we got value from redis
	if val != nil {
		log.Print("got value from redis")

		// since val exist on redis and not on memory cache,
		// store retrieved value to memory cache
		u.memory.Set(key, val)

		return
	}

	log.Print("got no value from anywhere")

	return
}

func (u *usecase) Set(key string, val interface{}) (err error) {
	err = u.redis.Set(key, val)
	return
}

func (u *usecase) Refresh(key string) (err error) {
	val, err := u.redis.Get(key)
	if err != nil {
		return
	}

	err = u.memory.Set(key, val)
	return
}
