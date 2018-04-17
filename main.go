package main

import (
	"github.com/zouyx/agollo"
	"fmt"
	"encoding/json"
)

func main() {
	agollo.Start()
	v:=agollo.GetStringValue("hello","abc")
	fmt.Println(v)
	for  {
		event := agollo.ListenChangeEvent()
		changeEvent := <-event
		bytes, _ := json.Marshal(changeEvent)
		fmt.Println("event:", string(bytes))
	}
}