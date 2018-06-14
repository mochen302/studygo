package main

import (
	"encoding/json"
	"fmt"
	"os"
	"log"
)

type Address struct {
	Type    string
	City    string
	Country string
}
type VCard struct {
	FirstName string
	LastName  string
	Address   []*Address
	Remark    string
}

func main() {
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}

	js, _ := json.Marshal(vc)
	fmt.Printf("JSON format: %s", js)

	file, _ := os.OpenFile("resources/vcard.json", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	enc := json.NewEncoder(file)
	err := enc.Encode(vc)
	if err == nil {
		str, _ := json.Marshal(vc)
		log.Println("new JsonFile: \n", str)
	}
}
