package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type User struct {
	ID           int64
	Name         string
	Email        string
	Password     string
	RegisteredAt time.Time
}

func main() {
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=postgres sslmode=disable password=goLANGn1nja")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected")

	// #1
	// rows, err := db.Query("SELECT id, name, email, password, registered_at FROM users")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer rows.Close()

	// users := make([]User, 0)
	// for rows.Next() {
	// 	u := User{}
	// 	err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.RegisteredAt)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	users = append(users, u)
	// }

	// err = rows.Err()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("%+v\n", users)

	// #2
	// row := db.QueryRow("SELECT id, name, email, password, registered_at FROM users WHERE id = $1", 1)

	// u := User{}
	// err = row.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.RegisteredAt)
	// if err != nil {
	// 	if errors.Is(err, sql.ErrNoRows) {
	// 		fmt.Println("no rows")
	// 		return
	// 	}
	// 	log.Fatal(err)
	// }

	// fmt.Printf("%+v\n", u)

	err = insertUser(db, User{
		Name:     "Petya",
		Email:    "petya@gmail.com",
		Password: "awlwlwlallspspwls",
	})

	// err = updateUser(db, 1, User{
	// 	Name:  "John",
	// 	Email: "john@doe.com",
	// })

	if err != nil {
		log.Fatal(err)
	}

	users, err := getUsers(db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", users)

	// err = deleteUser(db, 2)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// u, err := getUserById(db, 1)
	// fmt.Printf("%+v", u)

}

func getUsers(db *sql.DB) ([]User, error) {

	rows, err := db.Query("SELECT id, name, email, password, registered_at FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]User, 0)
	for rows.Next() {
		u := User{}
		err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.RegisteredAt)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func getUserById(db *sql.DB, id int) (User, error) {
	u := User{}
	row := db.QueryRow("SELECT id, name, email, password, registered_at FROM users WHERE id = $1", id)
	err := row.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.RegisteredAt)

	return u, err
}

func insertUser(db *sql.DB, u User) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	_, err = tx.Exec("INSERT INTO users (name, email, password) VALUES ($1, $2, $3)",
		u.Name, u.Email, u.Password)

	if err != nil {
		return err
	}

	_, err = tx.Exec("INSERT INTO logs (entity, action) VALUES ($1, $2)",
		"user", "created")

	if err != nil {
		return err
	}

	return tx.Commit()
}

func insertUser2(db *sql.DB, u User) error {
	_, err := db.Exec("INSERT INTO users (name, email, password) VALUES ($1, $2, $3)",
		u.Name, u.Email, u.Password)

	return err
}

func deleteUser(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM users WHERE id = $1", id)

	return err
}

func updateUser(db *sql.DB, id int, newUser User) error {
	_, err := db.Exec("UPDATE users SET name=$1, email=$2 WHERE id = $3",
		newUser.Name, newUser.Email, id)

	return err
}
