package main

import "fmt"

type Person struct {
	Name   string
	Age    int32
	Gender string
	Vote   string
}

func (P Person) Introduction() {
	fmt.Println("Hi, my name is ",P.Name, "My age is",P.Age,)	
}
func checkforvote(P *Person){
    //stat:="yes"
    if P.Age >=21 {
        fmt.Println("Yes this person can vote")
    }else{
        fmt.Println("No this person can not vote")
    }
}
func (P Person) UpdateAge(){ //This will pass a copy of it
	P.Age = 21
	fmt.Println("Updated age of this user is", P.Age)
}
func main() {
	var person Person
	fmt.Println("Enter your name:") //user input for name
	fmt.Scanln(&person.Name)

	fmt.Println("Enter your age:") //user input for age
	fmt.Scanln(&person.Age)

	fmt.Println("Enter Gender:") //user input for gender
	fmt.Scanln(&person.Gender)

	fmt.Println("Person's Details:", person)
	person.Introduction()
	checkforvote(&person)
	person.UpdateAge()
}
