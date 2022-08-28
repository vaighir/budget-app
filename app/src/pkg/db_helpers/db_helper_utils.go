package db_helpers

import (
	"log"
)

func cleanup(){

	if r := recover(); r != nil {
		log.Println("Recovered. Error:\n", r)
	}
}