package main

import (
	"encoding/json"
	"flag"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

// Configuration singleton for the project
var conf Config

// Config file format
type Config struct {
	ApiKey      string `json:"apiKey"`
	ReleaseMode bool   `json:"releaseMode"`
}

// LoadConf Loads a config file into the conf var
func LoadConf(file string) {
	configFile, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	fileBytes, _ := io.ReadAll(configFile)
	err = json.Unmarshal(fileBytes, &conf)
	if err != nil {
		panic(err)
	}
}

// Main entry point for the program
func main() {
	config := flag.String("c", "config.json", "Config File")
	address := flag.String("a", "localhost", "Address to serve from")
	port := flag.Int("p", 8080, "Port to serve from")
	LoadConf(*config)
	if conf.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	PhoneHomerAPIv1(*address, *port)
}
