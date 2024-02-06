package main

var conf Config

type Config struct {
	ApiKey  string `json:"apiKey,omitempty"`
	Address string `json:"address,omitempty"`
	Port    int    `json:"port,omitempty"`
}

func LoadConf(file string) Config {
	return Config{ApiKey: "1234"}
}

func main() {
	conf = LoadConf("config.json")
	PhoneHomerAPIv1("localhost", 8080)
}
