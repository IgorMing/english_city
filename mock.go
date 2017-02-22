package main

import "fmt"

var rooms Rooms

func init() {
	addRoom("New York")
	addRoom("Chicago")
	addRoom("Los Angeles")
	addRoom("Miami")
}

func getRoomByID(id int) Room {
	for _, r := range rooms {
		if r.ID == id {
			return r
		}
	}
	return Room{}
}

func getRooms() Rooms {
	return rooms
}

func addRoom(name string) (Room, error) {
	if alreadyExists(name) {
		return Room{}, fmt.Errorf("The resource with the name \"%s\" already exists", name)
	}
	nextID := len(rooms) + 1
	var room = Room{ID: nextID, Name: name}
	rooms = append(rooms, room)
	return room, nil
}

func alreadyExists(name string) bool {
	for _, r := range rooms {
		if r.Name == name {
			return true
		}
	}
	return false
}

func deleteRoom(id int) error {
	for i, r := range rooms {
		if r.ID == id {
			rooms = append(rooms[:i], rooms[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Cannot find the element with id: %d", id)
}
