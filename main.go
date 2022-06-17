package main

import (
	"encoding/json"
	"fmt"
	"log"
	rand2 "math/rand"
	"net/http"
	"os"
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
	/*log.Printf("IN: %#v", input)
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
	var hit1x, hit1y, hit2x, hit2y, hit3x, hit3y int
	if doubleMayDirection == "E" {
			log.Printf("Hit E")
		hit1x = doubleMayX
		hit2x = doubleMayX
		hit3x = doubleMayX
		hit1y = doubleMayY - 1
		hit2y = doubleMayY - 2
		hit3y = doubleMayY - 3
	}
	if doubleMayDirection == "N" {
			log.Printf("Hit N")
		hit1x = doubleMayX
		hit2x = doubleMayX
		hit3x = doubleMayX
		hit1y = doubleMayY + 1
		hit2y = doubleMayY + 2
		hit3y = doubleMayY + 3
	}
	if doubleMayDirection == "W" {
			log.Printf("Hit W")
		hit1x = doubleMayX + 1
		hit2x = doubleMayX + 2
		hit3x = doubleMayX + 3
		hit1y = doubleMayY
		hit2y = doubleMayY
		hit3y = doubleMayY
	}
	if doubleMayDirection == "S" {
			log.Printf("Hit S")
		hit1x = doubleMayX - 1
		hit2x = doubleMayX - 2
		hit3x = doubleMayX - 3
		hit1y = doubleMayY
		hit2y = doubleMayY
		hit3y = doubleMayY
	}
	log.Println("Hit 1 : " + strconv.Itoa(hit1x) + "," + strconv.Itoa(hit1y))
	log.Println("Hit 2 : " + strconv.Itoa(hit2x) + "," + strconv.Itoa(hit2y))
	log.Println("Hit 3 : " + strconv.Itoa(hit3x) + "," + strconv.Itoa(hit3y))
	for key, chikan := range input.Arena.State {
		//
		/*
			if key == doubleMay {
				continue
			}*/
	//log.Println("Chikan:" + key + " x:" + strconv.Itoa(chikan.X) + " y:" + strconv.Itoa(chikan.Y))
	/*	if chikan.X == hit1x && chikan.Y == hit1y {
			log.Println("Hit Chikan:" + key + " at hit1")
			return "T"
		}
		if chikan.X == hit2x && chikan.Y == hit2y {
			log.Println("Hit Chikan:" + key + " at hit2")
			return "T"
		}
		if chikan.X == hit3x && chikan.Y == hit3y {
			log.Println("Hit Chikan:" + key + " at hit3")
			return "T"
		}

	}*/
	commands := []string{"F", "R", "L", "T"}
	rand := rand2.Intn(4)
	return commands[rand]
}
