package configs

import (
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

type Config struct {
	Router *httprouter.Router
}

func New() Config {
	return Config{}
}

func (c *Config) Start() error {
	ch := make(chan bool)

	go func() {
		err := c.Start()
		if err != nil {
			return
		}
	}()

	<-ch
	return nil
}

func (c *Config) InitEnv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	return nil
}

func Catch(err error) {
	if err != nil {
		panic(err)
	}
}
