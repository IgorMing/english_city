package main

type Room struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Rooms []Room
