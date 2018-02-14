# go-dialogflow

Original implementation by mlabouardy [dialogflow-go-client](github.com/mlabouardy/dialogflow-go-client).

[![License](https://img.shields.io/badge/License-Apache%202.0-yellowgreen.svg)](https://opensource.org/licenses/Apache-2.0) 

<div align="center">
	<img src="logo.png" width="50%"/>
</div>

This library allows integrating agents from the [DialogFlow](https://dialogflow.com) natural language processing service with your Golang application.

* [Prerequsites](#prerequsites)
* [Installation](#installation)
* [Features](#features)
* [Usage](#usage)

# Prerequsites

Create an [DialogFlow account](https://dialogflow.com/).

# Installation

```shell
go get github.com/kompiuter/go-dialogflow
```

# Features

* Queries
* Contexts
* Intents
* UserIntents
* Entities

# Usage

* Create `main.go` file with the following code:

```golang
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

	entities, err := client.EntitiesFindAllRequest()
	if err != nil {
		log.Fatal(err)
	}
	for _, entity := range entities {
		fmt.Println(entity.Name)
	}
}
```
* Run following command.
```shell
go run main.go
```
* Your can find more examples in [`examples`](examples) directory.