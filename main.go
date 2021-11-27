package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"encoding/json"

	server "github.com/branthony-92/mhriseservergo/server"
	"github.com/joho/godotenv"
)

func LoadURL() {
	err := godotenv.Load("serverconfig.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	server.URL = os.Getenv("MONGODB_URI")
}

func main() {
	numArgs := len(os.Args)
	if numArgs < 3 {
		return
	}

	hostname := os.Args[1]
	portstring := os.Args[2]

	url := fmt.Sprintf("%v:%v", hostname, portstring)

	http.HandleFunc("/api/v1/skills", handleSkills)
	http.HandleFunc("/api/vi/armoursets", handleArmour)

	fmt.Printf("Server Listening On %v\n", url)
	LoadURL()
	if err := http.ListenAndServe(url, nil); err != nil {
		panic(err)
	}
}

func handleSkills(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Skill Query Request")
	skills := server.QueryAllSkills()

	resp := server.ResponseBody{
		Code:         server.EOK,
		ErrorMessage: "",
		Body:         skills,
	}

	data, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {

	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func handleArmour(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Armour Query Request")
	armour := server.QueryAllArmour()

	resp := server.ResponseBody{
		Code:         server.EOK,
		ErrorMessage: "",
		Body:         armour,
	}

	data, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {

	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
