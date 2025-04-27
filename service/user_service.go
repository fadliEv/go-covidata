package service

import (
	"be-covidata/config"
	"be-covidata/entity"
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

var db *sql.DB

func SetupDB() {
	// Ambil path database dari .env
	dbPath := config.GetEnv("DB_PATH")

	// Membuka koneksi database dengan path yang sudah disesuaikan
	var err error
	db, err = sql.Open("sqlite", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database connected successfully at", dbPath)
}

func CreateUser(user entity.User) error {
    stmt, err := db.Prepare("INSERT INTO users(name, age, status, location) VALUES(?, ?, ?, ?)")
    if err != nil {
        return err
    }
    _, err = stmt.Exec(user.Name, user.Age, user.Status, user.Location)
    return err
}

func GetUsers() ([]entity.User, error) {
    rows, err := db.Query("SELECT id, name, age, status, location FROM users order by id DESC")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []entity.User
    for rows.Next() {
        var user entity.User
        if err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Status, &user.Location); err != nil {
            return nil, err
        }
        users = append(users, user)
    }

    return users, nil
}
