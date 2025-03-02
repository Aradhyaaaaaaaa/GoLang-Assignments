package main

import ("fmt"
"github.com/go-playground/validator/v10")

type Person struct {
	Name   string 
	Age    uint //for only positive value
	Gender string `validate:"required,oneof=male female m f"`
	Vote   string
}

func validatePerson(p *Person) error {
	validate := validator.New()
	return validate.Struct(p)
}

func (P Person) introduction() {
	fmt.Println("Hi, my name is ",P.Name, "My age is",P.Age,)	
}

func checkforvote(P *Person){
    //stat:="yes"
    if P.Age >=21 {
        fmt.Println("Yes this person can vote")
    }else{
        fmt.Println("No this person can not vote")
		P.UpdateAge()
    }
}

func (P Person) UpdateAge(){ //for updating the age
	P.Age = 21
	fmt.Println("For being elgible to vote person have to be ", P.Age, "or more then that")
}

func main() {
	var person Person
	fmt.Println("Enter your name:") //user input for name
	fmt.Scanln(&person.Name)

	fmt.Println("Enter your age:") //user input for age
	fmt.Scanln(&person.Age)

	fmt.Println("Enter Gender(Male, Female or m, f  ):") //user input for gender
	fmt.Scanln(&person.Gender)

	if err := validatePerson(&person); err != nil {
		fmt.Println("Person could either Female or Male", err)
		return
	}

	fmt.Println("Person's Details:", person)
	person.introduction()
	checkforvote(&person)
}
