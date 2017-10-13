package person

type Person struct {
	ID        string   `json:"id,ommitempty"`
	FirstName string   `json:"first_name,ommitempty`
	LastName  string   `json:"last_name,ommitempty`
	Address   *Address `json: "address,ommitempty"`
}

func NewPerson(id string, fname string, lname string, address *Address) Person {
	return Person{
		ID:        id,
		FirstName: fname,
		LastName:  lname,
		Address:   address,
	}
}
