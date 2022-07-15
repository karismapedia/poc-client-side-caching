package cli

import (
	"github.com/karismapedia/poc-client-side-caching/constant"
	"github.com/karismapedia/poc-client-side-caching/usecase"
)

type cli struct {
	use usecase.Usecase

	handler map[string]func(...string) string
}

func Init(use usecase.Usecase) (c cli) {
	c = cli{use: use, handler: make(map[string]func(...string) string)}

	c.handler[constant.CommandGet] = c.get
	c.handler[constant.CommandSet] = c.set
	return
}

func (c *cli) get(ss ...string) (s string) {
	if len(ss) < 1 {
		s = "need key name"
		return
	}

	if len(ss) > 1 {
		s = "too many command, require only key name"
		return
	}

	val, err := c.use.Get(ss[0])
	if err != nil {
		s = "err: " + err.Error()
		return
	}

	if val == nil {
		s = "<nil>"
		return
	}

	s, ok := val.(string)
	if !ok {
		s = "err: got unknown data type"
		return
	}

	return
}

func (c *cli) set(ss ...string) (s string) {
	if len(ss) < 2 {
		s = "need key name and value"
		return
	}

	if len(ss) > 2 {
		s = "too many command, require only key name and value"
		return
	}

	err := c.use.Set(ss[0], ss[1])
	if err != nil {
		s = "err: " + err.Error()
		return
	}

	s = "OK"
	return
}
