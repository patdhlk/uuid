package main

import (
	"log"

	"github.com/patdhlk/uuid"
)

func main() {
	uuid, _ := uuid.UUID() //getting the uuid as a string
	log.Println(uuid)
}
