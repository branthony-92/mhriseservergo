package main

import (
	"fmt"
	"io"
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
	http.HandleFunc("/api/v1/optimize", handleOptimize)

	fmt.Printf("Server Listening On %v\n", url)
	LoadURL(envFileName)

	if err := http.ListenAndServe(url, nil); err != nil {
		panic(err)
	}
}

func handleSkills(w http.ResponseWriter, req *http.Request) {
	resp := server.ResponseBody{
		Code:         0,
		ErrorMessage: "",
		Body:         nil,
	}
	sendReply := func(r server.ResponseBody) {
		data, err := json.MarshalIndent(resp, "", "  ")
		if err != nil {
			resp.Body = nil
			resp.ErrorMessage = err.Error()
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
	skills, err := server.QueryAllSkills()

	if err != nil {
		resp.ErrorMessage = err.Error()
	} else {
		resp.Body = skills
	}
	sendReply(resp)
}

func handleArmour(w http.ResponseWriter, req *http.Request) {
	resp := server.ResponseBody{
		Code:         0,
		ErrorMessage: "",
		Body:         nil,
	}
	sendReply := func(r server.ResponseBody) {
		data, err := json.MarshalIndent(resp, "", "  ")
		if err != nil {
			resp.Body = nil
			resp.ErrorMessage = err.Error()
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}

	fmt.Println("Armour Query Request")
	armour, err := server.QueryAllArmour()

	if err != nil {
		resp.ErrorMessage = err.Error()
		return
	} else {
		resp.Body = armour
	}
	sendReply(resp)
}

func handleApplyFilter(w http.ResponseWriter, req *http.Request) {
	resp := server.ResponseBody{
		Code:         0,
		ErrorMessage: "",
		Body:         nil,
	}
	sendReply := func(r server.ResponseBody) {
		data, err := json.MarshalIndent(resp, "", "  ")
		if err != nil {
			resp.Body = nil
			resp.ErrorMessage = err.Error()
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}

	q := req.URL.Query()

	filter, ok := q["filter"]
	if !ok {
		resp.ErrorMessage = "No Filters In Query"
		sendReply(resp)
		return
	}

	armour, err := server.QueryFilter([]byte(filter[0]))

	if err != nil {
		resp.ErrorMessage = err.Error()
		return
	} else {
		resp.Body = armour
	}
	sendReply(resp)
}

func handleOptimize(w http.ResponseWriter, req *http.Request) {
	resp := server.ResponseBody{
		Code:         0,
		ErrorMessage: "",
		Body:         nil,
	}
	sendReply := func(r server.ResponseBody) {
		data, err := json.MarshalIndent(resp, "", "  ")
		if err != nil {
			resp.Body = nil
			resp.ErrorMessage = err.Error()
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}

	data, err := io.ReadAll(req.Body)
	if err != nil {
		resp.ErrorMessage = "No Body"
		sendReply(resp)
		return
	}

	customSet, err := server.Optimize(data)
	summary := *server.Summarize(customSet)
	customSet.Pieces = append(customSet.Pieces, summary)

	if err != nil {
		resp.ErrorMessage = err.Error()
		return
	} else {
		resp.Body = customSet
	}
	sendReply(resp)
}
