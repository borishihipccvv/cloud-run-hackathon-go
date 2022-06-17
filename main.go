package main

import (
	"encoding/json"
	"fmt"
	"log"
	rand2 "math/rand"
	"net/http"
	"os"
	"strconv"
)

func main() {
	port := "8080"
	if v := os.Getenv("PORT"); v != "" {
		port = v
	}
	http.HandleFunc("/", handler)

	log.Printf("starting server on port :%s", port)
	err := http.ListenAndServe(":"+port, nil)
	log.Fatalf("http listen error: %v", err)
}

func handler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		fmt.Fprint(w, "Let the battle begin!")
		return
	}

	var v ArenaUpdate
	defer req.Body.Close()
	d := json.NewDecoder(req.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&v); err != nil {
		log.Printf("WARN: failed to decode ArenaUpdate in response body: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp := play(v)
	fmt.Fprint(w, resp)
}

func play(input ArenaUpdate) (response string) {
	//log.Printf("IN: %#v", input)
	//test
	doubleMay := input.Links.Self.Href
	doubleMayObj := input.Arena.State[doubleMay]
	areaXmax := input.Arena.Dimensions[0]
	areaYmax := input.Arena.Dimensions[1]
	doubleMayDirection := doubleMayObj.Direction
	doubleMayX := doubleMayObj.X
	doubleMayY := doubleMayObj.Y
	log.Printf("DoubleMay URL:" + doubleMay + " areaXmax:" + strconv.Itoa(areaXmax) + " doubleMayX:" + strconv.Itoa(doubleMayX) + " areaYmax:" + strconv.Itoa(areaYmax) + " doubleMayY:" + strconv.Itoa(doubleMayY) + " doubleMayDirection:" + doubleMayDirection)
	//Check chikan location
	for key, chikan := range input.Arena.State {
		log.Println("Key:", key, "=>", "Element:", chikan)
		if key == doubleMay {
			continue
		}
		if doubleMayDirection == "N" {
			log.Printf("Hit N")
		}
		if doubleMayDirection == "E" {
			log.Printf("Hit E")
		}
		if doubleMayDirection == "S" {
			log.Printf("Hit S")
		}
		if doubleMayDirection == "W" {
			log.Printf("Hit W")
		}
	}
	commands := []string{"F", "R", "L", "T"}
	rand := rand2.Intn(4)
	return commands[rand]
}
