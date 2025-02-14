package main

import (
	"log"

	"github.com/MohamedMosalm/To-Do-List/cmd"
	"github.com/MohamedMosalm/To-Do-List/config"
)

func main() {
	cfg, err := config.SetupEnv()
	if err != nil {
		log.Fatalf("config file setup failed, err: %v", err)
	}
	cmd.StartServer(cfg)
}
