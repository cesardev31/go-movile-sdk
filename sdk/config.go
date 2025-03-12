package sdk

import (
	_ "embed"
	"log"
	"sync"

	"github.com/BurntSushi/toml"
)

var embeddedConfig []byte

type Config struct {
	App struct {
		Name    string
		Version string
		Debug   bool
	}
	Server struct {
		Host string
		Port int
	}
	Auth struct {
		TokenExpiry int
		SecretKey   string
	}
	UI struct {
		Theme string
	}
}

var (
	config     Config
	configOnce sync.Once
)

func Init() error {
	var initErr error
	configOnce.Do(func() {
		if _, err := toml.Decode(string(embeddedConfig), &config); err != nil {
			initErr = err
			log.Printf("Error decodificando config embed: %v", err)
			return
		}
		log.Printf("Configuración cargada (embed): %+v", config)
	})
	return initErr
}

func GetConfig() Config {
	return config
}
