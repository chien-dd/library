package rabbitmq

import "fmt"

type Config struct {
	Address  string `json:"address" yaml:"address"`
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
}

func (conf *Config) String() string {
	// Success
	return fmt.Sprintf("amqp://%s:%s@%s", conf.Username, conf.Password, conf.Address)
}
