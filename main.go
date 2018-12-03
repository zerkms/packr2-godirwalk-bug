package main

import (
	"log"

	packr "github.com/gobuffalo/packr/v2"
)

func main() {
	box := packr.New("resources", "./resources")

	wsdlRaw, err := box.FindString("calculator.wsdl")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(wsdlRaw)
}
