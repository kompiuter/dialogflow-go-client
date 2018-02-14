package main

import (
	"fmt"
	"log"

	df "github.com/kompiuter/go-dialogflow"
	"github.com/kompiuter/go-dialogflow/model"
)

func main() {
	client, err := df.NewDialogFlowClient(model.Options{
		AccessToken: "<API.AI TOKEN GOES HERE>",
	})
	if err != nil {
		log.Fatal(err)
	}

	contexts, err := client.ContextsFindAllRequest()
	if err != nil {
		log.Fatal(err)
	}
	for _, context := range contexts {
		fmt.Printf("%s:%d\n", context.Name, context.Lifespan)
	}
}
