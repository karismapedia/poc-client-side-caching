package constant

const (
	RedisAddress = "127.0.0.1:6379"

	CommandGet  = "get"
	CommandQuit = "quit"
	CommandSet  = "set"

	TCP = "tcp"
)

var (
	Commands = map[string]bool{
		CommandGet:  true,
		CommandQuit: true,
		CommandSet:  true,
	}
)
