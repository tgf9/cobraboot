package main

import (
	"log"

	"github.com/tgf9/cobraboot/internal/command"
)

func main() {
	if err := command.Run(); err != nil {
		log.Fatalln(err)
	}
}
