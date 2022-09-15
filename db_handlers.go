package main

import (
	"database/sql"
	"log"
)

type UserDetails struct {
	Username  string
	Firstname string
	Lastname  string
	Email     string
	Phone     string
	Aadhaar   string
	PAN       string
	Addr1     string
	Addr2     string
	PIN       string
}

type UserAuth struct {
	Username string
	Password string
}

func getDetailsByUsername(username string, db *sql.DB) (res UserDetails) {
	var details UserDetails
	query := "SELECT username, firstname, lastname, email, phone, aadhaar, pan, addr1, addr2, pin FROM users WHERE username='" + username + "';"
	rows, dbErr := db.Query(query)
	if dbErr != nil {
		log.Fatal(dbErr.Error())
	}
	defer rows.Close()
	for rows.Next() {
		scanErr := rows.Scan(&details.Username, &details.Firstname, &details.Lastname, &details.Email, &details.Phone, &details.Aadhaar, &details.PAN, &details.Addr1, &details.Addr2, &details.PIN)
		if scanErr != nil {
			log.Fatal(scanErr.Error())
		}
	}
	err := rows.Err()
	if err != nil {
		log.Fatal(err.Error())
	}

	return details
}

func getUserAuth(username string, db *sql.DB) (res UserAuth) {
	var auth UserAuth
	query := "SELECT username, password FROM user where username='" + username + "';"
	rows, dbErr := db.Query(query)
	if dbErr != nil {
		log.Fatal(dbErr.Error())
	}
	defer rows.Close()
	for rows.Next() {
		scanErr := rows.Scan(&auth.Username, &auth.Password)
		if scanErr != nil {
			log.Fatal(scanErr.Error())
		}
	}
	err := rows.Err()
	if err != nil {
		log.Fatal(err.Error())
	}

	return auth
}
