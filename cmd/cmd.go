package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gomodule/redigo/redis"
)

func main() {
	c1, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Fatal(err)
		return
	}

	rep, err := c1.Do("client", "id")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("client id:", rep)

	c1.Send("subscribe", "__redis__:invalidate")
	c1.Flush()
	go func() {
		for {
			reply, err := c1.Receive()
			if err != nil {
				fmt.Printf("err: %v\n", err)
				continue
			}
			fmt.Println("receive:", reply)
		}
	}()

	c2, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Fatal(err)
		return
	}

	rep2, err := c2.Do("client", "tracking", "on", "redirect", rep)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("client 2 start tracking:", rep2)

	key1, _ := c2.Do("get", "key1")
	fmt.Println("key1:", key1)
	key2, _ := c2.Do("get", "key2")
	fmt.Println("key2:", key2)
	key3, _ := c2.Do("get", "key3")
	fmt.Println("key3:", key3)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)
	<-sigCh
}
