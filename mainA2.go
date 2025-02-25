package main

import (
	"fmt"
	"strconv"
)

type Employee struct {
	Name   string
	Age    int
	Id     int
	Salary int
}
type Department struct{
	Name   string
	Employees []Employee//this allow us to store multiple employees in single department
}
func Employeesdetails() []string{ 
	var Name string
	fmt.Println("Enter employee's name:") //user input for name
	fmt.Scanln(&Name)
	var Age int
	fmt.Println("Enter employee's age:") //user input for age
	fmt.Scanln(&Age)
	var Salary int
	fmt.Println("Enter salary:") //user input for salary
	fmt.Scanln(&Salary)
	var Id int
	fmt.Println("Enter ID:")
	fmt.Scanln(&Id)
	return []string{Name, 
		fmt.Sprintf("%d",Age), 
		fmt.Sprintf("%d",Id), 
		fmt.Sprintf("%d",Salary),
	}
}

func (dep *Department) AvrageSalary(){
	sum:= 0
	count:= 0
	average:= 0
	for _, eachEmp := range dep.Employees{
		sum += eachEmp.Salary
		count++
	}
	average = sum / count
	fmt.Printf("The average salary of employees is: %d", average)
}

func (dep * Department) Raise(){
	var empID int
	fmt.Printf("Name the employee whom you want to give raise")
	_, err := fmt.Scanln(&empID)
	if err != nil{
		fmt.Println("Error",err)
	}
}
func (Dep *Department) AddEmp(){
    // newEmp := []Employee{}
	newEmp := Employeesdetails()
	age, _ := strconv.Atoi(newEmp[1])//this will convertstring to int
	salary, _ := strconv.Atoi(newEmp[2])
	Id, _ := strconv.Atoi(newEmp[3])
	
		Dep.Employees = append(Dep.Employees, 
			Employee{
				newEmp[0],
				age,
				salary,
				Id,
			},
		)
	
	for _, values := range Dep.Employees {
		fmt.Println(values)
	}
	// fmt.Println(newEmp)
}

func main1(){
	//creating a department
	Department1 := Department{Name: "Golnang", Employees:[]Employee{
			{"ABC", 24, 565, 50000},
			{"XYZ", 25, 787, 60000},
		},
	}

	// Employeesdetails()
	Department1.AddEmp()
	Department1.Raise()
	Department1.AvrageSalary()
}
