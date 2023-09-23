package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON video")
	EncodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{
			"ReactJS Bootcamp",
			299,
			"LearncodeOnline.in",
			"abc123",
			[]string{"web-dev", "js"},
		}, {
			"MERN Bootcamp",
			199,
			"LearncodeOnline.in",
			"bcd123",
			[]string{"full-stack", "js"},
		}, {
			"Angular Bootcamp",
			399,
			"LearncodeOnline.in",
			"hit123",
			nil,
		},
	}

	// package this data as JSON data

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}
