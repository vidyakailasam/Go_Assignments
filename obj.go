package main

import (
	"encoding/json"
	"fmt"
)

type employee struct {
	Name string
	Role string
	Age  int
}

func main() {
	employeeJson := `{ "Name":"VidyaKailasam", "Role":"Projectengineer", "Age":21 }`
	var employeee employee
	json.Unmarshal([]byte(employeeJson), &employeee)
	fmt.Printf("Name: %s, Role: %s, Age: %d", employeee.Name, employeee.Role, employeee.Age)

}
