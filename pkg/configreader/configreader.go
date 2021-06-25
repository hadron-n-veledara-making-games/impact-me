package configreader

import (
	"log"

	"github.com/BurntSushi/toml"
)

func Read(path string, c interface{}) {
	if _, err := toml.DecodeFile(path, c); err != nil {
		log.Fatal(err.Error())
	}
}
