package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type jsonPerson struct {
	ID        int
	FirstName string
	LastName  string
	UserName  string
}

type Person struct {
	XMLName   xml.Name `xml:"person"`
	ID        int      `xml:"id"`
	FirstName string   `xml:"firstname"`
	LastName  string   `xml:"lastname"`
	UserName  string   `xml:"username"`
}

type Company struct {
	XMLName xml.Name `xml:"company"`
	Persons []Person `xml:"person"`
}

// Function of object Person
func (p Person) String() string {
	return fmt.Sprintf("\t ID : %d - FirstName : %s - LastName : %s - UserName : %s \n", p.ID, p.FirstName, p.LastName, p.UserName)
}

func main() {
	xmlFile, err := os.Open("Employees.xml")
	if err != nil {
		fmt.Println("Opening file error : ", err)
		return
	}
	defer xmlFile.Close()

	xmlData, _ := ioutil.ReadAll(xmlFile)

	var c Company
	xml.Unmarshal(xmlData, &c)

	// Write XML on screen
	fmt.Println(c.Persons)

	// Convert to JSON
	var person jsonPerson
	var persons []jsonPerson

	for _, value := range c.Persons {
		person.ID = value.ID
		person.FirstName = value.FirstName
		person.LastName = value.LastName
		person.UserName = value.UserName

		persons = append(persons, person)
	}

	jsonData, err := json.Marshal(persons)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Write JSON on screen
	fmt.Println(string(jsonData))

	// Write to JSON file
	jsonFile, err := os.Create("./Employees.json")

	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	jsonFile.Write(jsonData)
	jsonFile.Close()
}
