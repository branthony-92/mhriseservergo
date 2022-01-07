package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"encoding/json"

	server "github.com/branthony-92/mhriseservergo/server"
	//"github.com/joho/godotenv"
)

func LoadURL(fname string) {
	//err := godotenv.Load(fname)
	//if err != nil {
	//	log.Fatal("Error loading .env file")
	//}
	server.URL = "mongodb+srv://mhrise-user:dYiVISDF8teKybdA@mhrisedb.ckl1w.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
}

func sendErrorReply(w http.ResponseWriter, err error, statusCode int) {
	msg := "Unknown Error"
	if err != nil {
		msg = err.Error()
	}
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "html/text")
	w.Write([]byte(msg))
}
func sendReply(w http.ResponseWriter, r server.ResponseBody) {
	data, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		sendErrorReply(w, err, 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(data)
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

	switch req.Method {
	case http.MethodPut:
		sendErrorReply(w, fmt.Errorf("Method %v not supported for this endpoint", req.Method), 400)
	case http.MethodPost:
		sendErrorReply(w, fmt.Errorf("Method %v not supported for this endpoint", req.Method), 400)
	case http.MethodGet:
		skills, err := server.QueryAllSkills()
		if err != nil {
			resp.ErrorMessage = err.Error()
		} else {
			resp.Body = skills
		}
		sendReply(w, resp)
	case http.MethodDelete:
		sendErrorReply(w, fmt.Errorf("Method %v not supported for this endpoint", req.Method), 400)
	default:
		sendErrorReply(w, fmt.Errorf("Method %v not supported for this endpoint", req.Method), 400)
	}

}

func handleArmour(w http.ResponseWriter, req *http.Request) {
	resp := server.ResponseBody{
		Code:         0,
		ErrorMessage: "",
		Body:         nil,
	}

	fmt.Println("Armour Query Request")

	switch req.Method {
	case http.MethodPut:
		sendErrorReply(w, fmt.Errorf("Method %v not supported for this endpoint", req.Method), 400)
	case http.MethodPost:
		sendErrorReply(w, fmt.Errorf("Method %v not supported for this endpoint", req.Method), 400)
	case http.MethodGet:
		armour, err := server.QueryAllArmour()
		if err != nil {
			resp.ErrorMessage = err.Error()
		} else {
			resp.Body = armour
		}
		sendReply(w, resp)
	case http.MethodDelete:
		sendErrorReply(w, fmt.Errorf("Method %v not supported for this endpoint", req.Method), 400)
	default:
		sendErrorReply(w, fmt.Errorf("Method %v not supported for this endpoint", req.Method), 400)
	}
}

func handleApplyFilter(w http.ResponseWriter, req *http.Request) {
	resp := server.ResponseBody{
		Code:         0,
		ErrorMessage: "",
		Body:         nil,
	}

	switch req.Method {
	case http.MethodPut:
		data, err := io.ReadAll(req.Body)
		if err != nil {
			resp.ErrorMessage = "No Body"
			sendReply(w, resp)
			return
		}
		armour, err := server.QueryFilter(data)
		if err != nil {
			resp.ErrorMessage = err.Error()
			return
		} else {
			resp.Body = armour
		}
		sendReply(w, resp)
	case http.MethodPost:
		sendErrorReply(w, fmt.Errorf("Method %v not supported for this endpoint", req.Method), 400)
	case http.MethodGet:
		sendErrorReply(w, fmt.Errorf("Method %v not supported for this endpoint", req.Method), 400)
	case http.MethodDelete:
		sendErrorReply(w, fmt.Errorf("Method %v not supported for this endpoint", req.Method), 400)
	default:
		sendErrorReply(w, fmt.Errorf("Method %v not supported for this endpoint", req.Method), 400)
	}
}

func handleOptimize(w http.ResponseWriter, req *http.Request) {
	resp := server.ResponseBody{
		Code:         0,
		ErrorMessage: "",
		Body:         nil,
	}
	switch req.Method {
	case http.MethodPut:
		data, err := io.ReadAll(req.Body)
		if err != nil {
			resp.ErrorMessage = "No Body"
			sendReply(w, resp)
			return
		}

		customSet, err := server.Optimize(data)
		fmt.Println(server.Summarize(*customSet))

		if err != nil {
			resp.ErrorMessage = err.Error()
			return
		} else {
			resp.Body = customSet
		}
		sendReply(w, resp)
	case http.MethodPost:
		sendErrorReply(w, fmt.Errorf("Method %v not supported for this endpoint", req.Method), 400)
	case http.MethodGet:
		sendErrorReply(w, fmt.Errorf("Method %v not supported for this endpoint", req.Method), 400)
	case http.MethodDelete:
		sendErrorReply(w, fmt.Errorf("Method %v not supported for this endpoint", req.Method), 400)
	default:
		sendErrorReply(w, fmt.Errorf("Method %v not supported for this endpoint", req.Method), 400)
	}
}
