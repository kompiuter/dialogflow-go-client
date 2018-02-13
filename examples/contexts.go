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

	contexts, err := client.ContextsFindAllRequest()
	if err != nil {
		log.Fatal(err)
	}
	for _, context := range contexts {
		fmt.Printf("%s:%d\n", context.Name, context.Lifespan)
	}
}
