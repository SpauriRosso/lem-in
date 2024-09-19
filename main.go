package main

import (
	"fmt"
	"log"
	"os"
)

type Ant struct {
	Name     string
	Location Room
}

var Ants []Ant

type Room struct {
	Name    string
	IsEmpty bool
	Visited bool
	Links   []string
}

var Rooms []Room

var Path []Room

var Paths [][]Room

var startRoom Room

var endRoom Room

func main() {

	fileContent := ReadFile()

	// fmt.Println(fileContent)

	DecodeFile(fileContent)

	if !IsFileValid() {
		log.Fatal("Invalid file format.")
		os.Exit(0)
	}

	// fmt.Println(startRoom.Name)
	// fmt.Println(endRoom.Name)
	// fmt.Println(Ants)
	// fmt.Println(Rooms[0])

	for _, room := range Rooms {
		PrintRoom(room)
	}
}

func PrintRoom(room Room) {

	fmt.Printf("Room: %s\n", room.Name)
	fmt.Printf("Links: %v\n", room.Links)
	fmt.Printf("Visited: %v\n", room.Visited)
	fmt.Printf("IsEmpty: %v\n", room.IsEmpty)
	fmt.Println()
}
