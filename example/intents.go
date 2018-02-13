package main

import (
	"fmt"
	"log"

	df "github.com/kompiuter/dialogflow-go-client"
	"github.com/kompiuter/dialogflow-go-client/model"
)

func main() {
	client, err := df.NewDialogFlowClient(model.Options{
		AccessToken: "<API.AI TOKEN GOES HERE>",
	})
	if err != nil {
		log.Fatal(err)
	}

	intents, err := client.IntentsFindAllRequest()
	if err != nil {
		log.Fatal(err)
	}
	for _, intent := range intents {
		fmt.Println(intent.Name)
	}
}
