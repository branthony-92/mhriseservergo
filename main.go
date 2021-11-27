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

func LoadURL(fname string) {
	err := godotenv.Load(fname)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	server.URL = os.Getenv("MONGODB_URI")
}

func main() {
	numArgs := len(os.Args)
	if numArgs < 4 {
		return
	}

	hostname := os.Args[1]
	portstring := os.Args[2]
	envFileName := os.Args[3]

	url := fmt.Sprintf("%v:%v", hostname, portstring)

	http.HandleFunc("/api/v1/skills", handleSkills)
	http.HandleFunc("/api/v1/armoursets", handleArmour)
	http.HandleFunc("/api/v1/applyfilter", handleApplyFilter)

	fmt.Printf("Server Listening On %v\n", url)
	LoadURL(envFileName)
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

func handleApplyFilter(w http.ResponseWriter, req *http.Request) {
	q := req.URL.Query()
	filter, ok := q["filter"]
	if !ok {

	}

	armour := server.QueryFilter(filter[0])
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

	fmt.Printf("%v", filter[0])
}
