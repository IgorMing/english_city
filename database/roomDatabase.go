package database

import (
	ec "github.com/IgorMing/english_city"
	"github.com/IgorMing/english_city/utils"
)

func GetRooms() (ec.Rooms, error) {
	rows, err := database.Query("SELECT id, name FROM room")
	if err != nil {
		return ec.Rooms{}, utils.HandleError("Error executing the query", err)
	}

	var rooms ec.Rooms
	for rows.Next() {
		var room ec.Room
		if err := rows.Scan(&room.ID, &room.Name); err != nil {
			return ec.Rooms{}, utils.HandleError("Error scanning the query", err)
		}
		rooms = append(rooms, room)
	}
	if err := rows.Err(); err != nil {
		return ec.Rooms{}, utils.HandleError("", err)
	}
	return rooms, nil
}

func InsertRoom(room ec.Room) (int, error) {
	var lastInsertID int
	err := database.QueryRow("INSERT INTO room(name) VALUES($1) returning id;", room.Name).Scan(&lastInsertID)
	if err != nil {
		return 0, utils.HandleError("Error while inserting a room", err)
	}

	return lastInsertID, nil
}

func DeleteRoom(id int) error {
	stmt, err := database.Prepare("DELETE FROM room WHERE id = $1")
	if err != nil {
		return utils.HandleError("Error preparing delete", err)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return utils.HandleError("Error while deleting row", err)
	}
	return nil
}

func UpdateRoom(id int, room ec.Room) (ec.Room, error) {
	stmt, err := database.Prepare("UPDATE room SET name=$1 WHERE id=$2")
	if err != nil {
		return ec.Room{}, utils.HandleError("Error preparing update", err)
	}

	_, err = stmt.Exec(room.Name, id)
	if err != nil {
		return ec.Room{}, utils.HandleError("Error while updating row", err)
	}

	room.ID = id

	return room, nil
}

// //Insert
// stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
// checkErr(err)
//
// res, err := stmt.Exec("reckhou", "IT", "2010-10-02")
// checkErr(err)
//
// id, err := res.LastInsertId()
// checkErr(err)
//
// fmt.Println(id)
//
// // Update
// stmt, err = db.Prepare("update userinfo set username=? where uid=?")
// checkErr(err)
//
// res, err = stmt.Exec("update", id)
// checkErr(err)
//
// affect, err := res.RowsAffected()
// checkErr(err)
//
// fmt.Println(affect)
//
// //Query
// rows, err := db.Query("SELECT * FROM userinfo")
// checkErr(err)
//
// for rows.Next() {
//     var uid int
//     var username string
//     var department string
//     var created string
//     err = rows.Scan(&uid, &username, &department, &created)
//     checkErr(err)
//     fmt.Println(uid)
//     fmt.Println(username)
//     fmt.Println(department)
//     fmt.Println(created)
// }
//
// //Delete
// stmt, err = db.Prepare("delete from userinfo where uid=?")
// checkErr(err)
//
// res, err = stmt.Exec(id)
// checkErr(err)
//
// affect, err = res.RowsAffected()
// checkErr(err)
//
// fmt.Println(affect)
//
// db.Close()
//
// }
