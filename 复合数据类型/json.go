package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie = struct{
	Title string
	// 因为go里成员都是大写的，但是有些json是小写的，所以用tag来控制
	Year int `json: "released"`		//Tag 来控制编码和解码，将 Year 对应到 json 中的 released.
	Color bool `json:"color, omitempty"`
	Actors []string
}

func main() {
	var movies = []Movie{
		{Title: "Casablance", Year: 1942, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true,
			Actors: []string{"Paul Newman"}},
		// ...
	}
	data, err := json.Marshal(movies)
	if err != nil{
		log.Fatal("JSON marshling failed: %s", err)
	}
	fmt.Printf("%s \n", data)

	// 便于阅读，prefix 是每一行输出的前缀，indent是每一层输出的前缀
	data2, err2 := json.MarshalIndent(movies, "", "    ")
	if err2 != nil {
		log.Fatal("JSON marshling failed: %s", err)
	}
	fmt.Printf("%s \n", data2)

	var titles []struct{ Title string}
	if err3 := json.Unmarshal(data, &titles); err3 != nil {
		log.Fatal("JSON unmarshaling failed: %s", err3)
	}
	fmt.Println(titles)
}


