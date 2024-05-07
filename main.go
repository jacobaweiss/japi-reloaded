package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Clues struct {
	Clues []Clue `json:"clues"`
}

type Clue struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
	Value    string `json:"value"`
	Category string `json:"category"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handleGet).Methods("GET")
	http.ListenAndServe(":8080", r)
}

func getClue() (clue Clue, err error) {
	jsonFile, err := os.Open("questions.json")
	if err != nil {
		fmt.Println("Unable to read JSON")
		return clue, errors.New("cant read file")
	}
	defer jsonFile.Close()

	bytevalue, _ := ioutil.ReadAll(jsonFile)

	var clues Clues
	json.Unmarshal(bytevalue, &clues)

	clueId := rand.Intn(216929)

	return clues.Clues[clueId], err
}

func handleGet(w http.ResponseWriter, req *http.Request) {
	clue, err := getClue()
	if err != nil {
		fmt.Println("Unable to get clue")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	resp, err := json.Marshal(clue)
	if err != nil {
		fmt.Println("Unable to encode JSON")
		return
	}

	w.Write(resp)
}
