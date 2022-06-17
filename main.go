package main

import (
	"GolangTrainee/user"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

var PORT = ":8088"

//belajar rest api dengan postman dari go

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users", postgetHanlder)
	r.HandleFunc("/users/{Id}", postgetHanlder)
	//http.HandleFunc("/users/", pathHanlder)
	//http.HandleFunc("/users/", UserHanlder)
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8088",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func pathHanlder(w http.ResponseWriter, r *http.Request) {
	users := []*user.User{
		{
			Id:       1,
			Username: "Joshuandi",
			Email:    "Joshunandi11@gmail.com",
			Age:      20,
		},
		{
			Id:       2,
			Username: "Entahlah",
			Email:    "Joshunandi11@gmail.com",
			Age:      12,
		},
	}
	fmt.Println(r.URL.Path[1:])
	paths := strings.Split(r.URL.Path[1:], "/")
	fmt.Printf("%+v", paths)
	fmt.Println(len(paths))
	if len(paths) == 2 && paths[1] != "" {
		if index, err := strconv.Atoi(paths[1]); err == nil {
			jsonData, _ := json.Marshal(&users[index-1])
			w.Header().Add("Content-Type", "application/json")
			w.Write(jsonData)
		}
	} else {
		fmt.Println("no param")
		jsonData, _ := json.Marshal(&users)
		w.Header().Add("Content-Type", "application/json")
		w.Write(jsonData)
	}
}

var pathUsers = map[int]*user.User{
	1: {
		Id:       1,
		Username: "Joshuandi",
		Email:    "Joshunandi11@gmail.com",
		Age:      20,
	},
	2: {
		Id:       2,
		Username: "Entahlah",
		Email:    "Joshunandi11@gmail.com",
		Age:      12,
	},
}

func postgetHanlder(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	id := param["Id"]
	var userss *user.User
	switch r.Method {
	case "POST":
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&userss)
		if err != nil {
			panic(err)
		}
		pathUsers[int(userss.Id)] = userss
		jsonData, _ := json.Marshal(pathUsers)
		w.Write(jsonData)
		fmt.Println("POST")
	case "GET":
		if id != "" {
			if index, err := strconv.Atoi(id); err == nil {
				jsonData, _ := json.Marshal(pathUsers[index])
				w.Header().Add("Content-Type", "application/json")
				w.Write(jsonData)
			}
		} else { //get all
			var userSlice []*user.User
			for _, v := range pathUsers {
				userSlice = append(userSlice, v)
			}
			jsonData, _ := json.Marshal(&userSlice)
			w.Header().Add("Content-Type", "application/json")
			w.Write(jsonData)
		}
	case "PUT":
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&userss)
		if err != nil {
			panic(err)
		}
		pathUsers[int(userss.Id)] = userss
		jsonData, _ := json.Marshal(pathUsers)
		w.Write(jsonData)
		fmt.Println("UPDATED")
	case "DELETE":
		if id != "" {
			if index, err := strconv.Atoi(id); err == nil {
				delete(pathUsers, index)
				jsonData, _ := json.Marshal(pathUsers)
				w.Write(jsonData)
				fmt.Println("DELETED")
			}
		}
	default:
		http.Error(w, "", http.StatusBadRequest)
	}
	w.Header().Add("Content-Type", "application/json")
}

func UserHanlder(w http.ResponseWriter, r *http.Request) {
	users := []*user.User{
		{
			Id:       1,
			Username: "Joshuandi",
			Email:    "Joshunandi11@gmail.com",
			Age:      20,
		},
		{
			Id:       2,
			Username: "Entahlah",
			Email:    "Joshunandi11@gmail.com",
			Age:      12,
		},
	}
	if r.Method == "GET" {
		json.NewEncoder(w).Encode(users)
		return
	}
	if r.Method == "POST" {

	}
	jsonData, _ := json.Marshal(users)
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)
}
