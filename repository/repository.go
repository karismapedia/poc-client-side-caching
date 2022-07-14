package repository

import (
	"github.com/gomodule/redigo/redis"
	"github.com/karismapedia/poc-client-side-caching/constant"
)

type repository struct {
	client1 redis.Conn
	client2 redis.Conn

	handler func()
}

func Init(address string) (r repository, err error) {
	c1, err := redis.Dial(constant.TCP, address)
	if err != nil {
		return
	}

	c2, err := redis.Dial(constant.TCP, address)
	if err != nil {
		return
	}

	r = repository{
		client1: c1,
		client2: c2,
	}
	return
}

func (r *repository) AssignHandler(handler func()) {
	r.handler = handler
}
