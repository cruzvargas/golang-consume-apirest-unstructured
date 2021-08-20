package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var results map[string]interface{}
	json.Unmarshal(responseData, &results)
	if _, ok := results["descriptions"]; ok {
		for _, item := range results["descriptions"].([]interface{}) {
			fmt.Println(item.(map[string]interface{})["description"])
		}
	}
}
