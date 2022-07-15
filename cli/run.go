package cli

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/karismapedia/poc-client-side-caching/constant"
)

func (c *cli) Run() {
	quit := make(chan bool, 1)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		var input, command, response string
		var h func(...string) string
		var ok bool
		scanner := bufio.NewScanner(os.Stdin)

		fmt.Print("> ")

		for scanner.Scan() {
			input = scanner.Text()

			sliceOfInput := strings.Fields(strings.ToLower(input))
			if len(sliceOfInput) < 1 {
				goto FINISH
			}

			command = sliceOfInput[0]

			if command == constant.CommandQuit {
				goto QUIT
			}

			h, ok = c.handler[sliceOfInput[0]]
			if !ok {
				fmt.Println("unrecognized command")
				goto FINISH
			}

			response = h(sliceOfInput[1:]...)
			fmt.Println(response)

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
