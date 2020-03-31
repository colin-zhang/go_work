package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"log"
	"time"
)

type Config struct {
	Age int
	Cats []string
	Pi float64
	Perfection []int
	DOB time.Time // requires `import time`
}

func main()  {
	content, err := ioutil.ReadFile("test.toml")
	if err != nil {
		log.Fatal(err)
	}

	var config Config
	_, err = toml.Decode(string(content), &config)
	if err != nil {
		log.Fatal()
	}

	fmt.Println(config)
}
