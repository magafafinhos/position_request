package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main(){
	err:= godotenv.Load()
		if err != nil{
			log.Fatal(".env can't be loaded...")
		}

	position_get := os.Getenv("position_url")

	response, err := http.Get(position_get)
	if err != nil{
		panic(err)
	}
	
	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)

	content, _ :=ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}