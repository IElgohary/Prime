package person

import (
	"encoding/json"
	"net/http"

	"../common"
)

type PersonController struct {
	common.Controller
}

var people []Person

func (c *PersonController) Seed() {
	people = append(people, NewPerson("1", "Islam", "Elgohary", nil))
	people = append(people, NewPerson("2", "Abdelrahman", "Elmeniawy", nil))
}

func getPerson(w http.ResponseWriter, req *http.Request) {

}
func (c *PersonController) GetPeople(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func createPerson(w http.ResponseWriter, req *http.Request) {

}

func deletePerson(w http.ResponseWriter, req *http.Request) {

}
