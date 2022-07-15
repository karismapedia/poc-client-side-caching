package main

import (
	"github.com/karismapedia/poc-client-side-caching/cli"
	"github.com/karismapedia/poc-client-side-caching/constant"
	"github.com/karismapedia/poc-client-side-caching/repository/memory"
	"github.com/karismapedia/poc-client-side-caching/repository/redis"
	"github.com/karismapedia/poc-client-side-caching/usecase"
)

func main() {
	mem := memory.Init()
	red, _ := redis.Init(constant.RedisAddress)

	use := usecase.Init(&mem, &red)

	cli := cli.Init(&use)

	red.AssignTrackHandler(cli.TrackHandler)
	red.Track()

	cli.Run()
}
