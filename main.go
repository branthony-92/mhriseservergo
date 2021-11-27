package main

import (
	"fmt"
	"net/http"
	"os"

	server "github.com/branthony-92/mhriseservergo/server"
)

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

	if err := http.ListenAndServe(url, nil); err != nil {
		panic(err)
	}

}

func handleSkills(w http.ResponseWriter, req *http.Request) {
	server.QueryAllSkills()
}

func handleArmour(w http.ResponseWriter, req *http.Request) {
	server.QueryAllArmour()
}
