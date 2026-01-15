package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

// type Employee struct {
// 	ID, ManagerID int
// 	Name, Address string
// 	DoB time.Time
// 	Position string
// 	Salary int
// }

func main() {

	var dilbert Employee

	dilbert.Salary -= 5000 // demoted
	position := &dilbert.Position
	*position = "Senior " + *position // promoted

	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Position += " (proactive team player)"

	fmt.Println(EmployeeByID(dilbert.ManagerID).Position) // prints boss position

	id := dilbert.ID
	EmployeeByID(id).Salary = 0 // fired
}

func EmployeeByID(id int) *Employee { /* ... */ }
