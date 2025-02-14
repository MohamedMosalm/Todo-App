package main

import (
	"log"

	"github.com/MohamedMosalm/Todo-App/cmd"
	"github.com/MohamedMosalm/Todo-App/config"
)

func main() {
	cfg, err := config.SetupEnv()
	if err != nil {
		log.Fatalf("config file setup failed, err: %v", err)
	}
	cmd.StartServer(cfg)
}
