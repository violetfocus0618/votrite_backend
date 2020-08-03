package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"./managers/database"
	"./managers/network"
)

var (
	env = []string{"VR_API_HOST",
		"VR_API_PORT",
		"VR_DB_HOST",
		"VR_DB_PORT",
		"VR_DB_USR",
		"VR_DB_PWD",
		"VR_DB_NAME"}
	svrChan = make(chan bool)
)

func run(host string, port int) {
	svrChan <- true
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), nil); err != nil {
		log.Fatal(err)
	}
	svrChan <- false
}

func checkEnv() {
	for _, key := range env {
		if _, ok := os.LookupEnv(key); !ok {
			log.Fatalf("Environment variable %s is not set", key)
		}
	}
}

func main() {
	//checkEnv()
	os.Setenv("VR_API_HOST", "0.0.0.0")
	os.Setenv("VR_API_PORT", "9191")
	os.Setenv("VR_DB_HOST", "localhost")
	os.Setenv("VR_DB_PORT", "5432")
	os.Setenv("VR_DB_USR", "postgres")
	os.Setenv("VR_DB_PWD", "Koukla2020=")
	os.Setenv("VR_DB_NAME", "votrite")
	os.Setenv("GCP_API_KEY", "AIzaSyDSswF9CPPmvoLgx-00fjF4_U3_mBw9iiM")

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.HandleFunc("/api/", network.RequestHandler)

	dbPort, err := strconv.Atoi(os.Getenv("VR_DB_PORT"))
	if err != nil {
		log.Fatal(err)
	}

	dbConnStr := fmt.Sprintf("user=%s password='%s' host=%s port=%d dbname=%s sslmode=disable",
		os.Getenv("VR_DB_USR"),
		os.Getenv("VR_DB_PWD"),
		os.Getenv("VR_DB_HOST"),
		dbPort,
		os.Getenv("VR_DB_NAME"))
	database.SetDbConnString(dbConnStr)

	port, err := strconv.Atoi(os.Getenv("VR_API_PORT"))
	if err != nil {
		log.Fatal(err)
	}

	go run(os.Getenv("VR_API_HOST"), port)

	for {
		if !<-svrChan {
			if _, err := fmt.Println("Server exiting"); err != nil {
				log.Fatal(err)
			}
			break
		}
	}
}
