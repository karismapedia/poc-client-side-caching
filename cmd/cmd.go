package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/gomodule/redigo/redis"
	"github.com/karismapedia/poc-client-side-caching/constant"
)

func main() {
	quit := make(chan bool, 1)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		var input, command string
		scanner := bufio.NewScanner(os.Stdin)

		fmt.Print("> ")

		for scanner.Scan() {
			input = scanner.Text()

			sliceOfInput := strings.Fields(strings.ToLower(input))
			if len(sliceOfInput) < 1 {
				goto FINISH
			}

			command = sliceOfInput[0]

			if !constant.Commands[command] {
				goto FINISH
			}
			if command == constant.CommandQuit {
				goto QUIT
			}

			fmt.Printf("got command: %s\n", command)

		FINISH:
			fmt.Print("> ")
		}

	QUIT:
		quit <- true
	}()

	var sigRcv os.Signal

	select {
	case sigRcv = <-sigCh:
		fmt.Printf("\ngot signal %v, quitting...\n", sigRcv)
	case <-quit:
		fmt.Println("quitting...")
	}

	fmt.Println("bye")
}

func init() {
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
