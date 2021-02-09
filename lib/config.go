package lib

// Config specificaton for all in application
type Config struct {
	Monitor struct {
		Port string
		Addr string
	} `toml:"monitor"`
	Hub map[string]struct {
		Port     string
		Baudrate string
		Parity   string
		Stop     string
		Size     string
		Timeout  string
		Type     string
	} `toml:"hub"`
}
