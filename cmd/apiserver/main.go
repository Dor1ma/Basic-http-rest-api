package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/Dor1ma/Basic-http-rest-api/internal/app/apiserver"
	"log"
	"os"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.json", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()

	content, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(content, &config)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(config)

	s := apiserver.New(config)

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}

}
