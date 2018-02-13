package main

import (
	"fmt"
	"log"

	df "github.com/kompiuter/dialogflow-go-client"
	"github.com/kompiuter/dialogflow-go-client/models"
)

func main() {
	err, client := df.NewDialogFlowClient(models.Options{
		AccessToken: "<API.AI TOKEN GOES HERE>",
	})
	if err != nil {
		log.Fatal(err)
	}

	entities, err := client.EntitiesFindAllRequest()
	if err != nil {
		log.Fatal(err)
	}
	for _, entity := range entities {
		fmt.Println(entity.Name)
	}
}
