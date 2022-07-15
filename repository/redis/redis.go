package redis

import (
	"log"

	"github.com/gomodule/redigo/redis"
	"github.com/karismapedia/poc-client-side-caching/constant"
)

type repository struct {
	client1 redis.Conn
	client2 redis.Conn

	handler func(interface{})
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

func (r *repository) AssignTrackHandler(handler func(interface{})) {
	r.handler = handler
}

func (r *repository) Track() (err error) {
	client1ID, err := r.client1.Do("client", "id")
	if err != nil {
		return err
	}

	r.client1.Send("subscribe", "__redis__:invalidate")
	r.client1.Flush()
	go func() {
		for {
			payload, err := r.client1.Receive()
			if err != nil {
				log.Println("receive err:", err)
				continue
			}

			r.handler(payload)
		}
	}()

	r.client2.Do("client", "tracking", "on", "redirect", client1ID)
	if err != nil {
		return err
	}

	return nil
}
