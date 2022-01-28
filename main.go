package main

import (
	"fmt"
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Ponto de partida")

	aplicacao := app.Gerar()

	err := aplicacao.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
