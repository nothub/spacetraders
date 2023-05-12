package main

import (
	"context"
	"log"
	"os"
	"runtime/debug"

	"github.com/nothub/spacetraders/api"
)

func main() {
	config := api.NewConfiguration()

	buildInfo, ok := debug.ReadBuildInfo()
	if ok {
		config.UserAgent = buildInfo.Main.Path + "/" + buildInfo.Main.Version
	}

	config.AddDefaultHeader("Accept", "application/json")
	config.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("SPACETRADERS_TOKEN"))

	client := api.NewAPIClient(config)

	res, _, err := client.AgentsApi.GetMyAgent(context.Background()).Execute()
	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Printf("%++v", res.GetData())
}
