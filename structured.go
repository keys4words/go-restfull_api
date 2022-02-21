package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Users struct {
	Users []User
}
type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"age"`
	Social Social `json:"social"`
}
type Social struct {
	Twitter  string
	Facebook string
}

func PrintUser(u *User) {
	fmt.Println("==================")
	fmt.Printf("Name: %s\n", u.Name)
	fmt.Printf("Type: %s\n", u.Type)
	fmt.Printf("Age: %v\n", u.Age)
	fmt.Printf("Social:\n\t- Twitter: %s\n\t- Facebook:%s\n", u.Social.Twitter, u.Social.Facebook)
}

type Professor struct {
	Name       string     `json:"name"`
	ScienceID  int        `json:"science_id"`
	IsWorking  bool       `json:"is_working"`
	University University `json:"university"`
}
type University struct {
	Name string `json:"name"`
	City string `json:"city"`
}

func main() {
	jsonFile, err := os.Open("users.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	var users Users
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(byteValue, &users)
	for _, v := range users.Users {
		PrintUser(&v)
	}

	// marshalling
	prof1 := Professor{
		Name:      "Nilse Bohr",
		ScienceID: 424535,
		IsWorking: false,
		University: University{
			Name: "GVD",
			City: "Koln",
		},
	}
	byteArr, err := json.MarshalIndent(prof1, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(byteArr))
	err = os.WriteFile("professors.json", byteArr, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
