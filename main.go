package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello, Kids!")

	var person = createPerson()
	fmt.Println("Person created!")
	fmt.Println("Name: " + person.name)
	fmt.Println("Age: " + strconv.Itoa(person.age))
	fmt.Println("Weight: " + strconv.FormatFloat(person.weight, 'f', 2, 64))
	fmt.Println("Height: " + strconv.FormatFloat(person.height, 'f', 2, 64))
	fmt.Printf("Street: %s \n", person.address.street)
	fmt.Printf("Number: %d \n", person.address.number)
	fmt.Printf("Postcode: %d%s \n", person.address.postcode.numbers, person.address.postcode.letters)

}

func ReadConsole(deleteLines ...int) string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")

	if deleteLines != nil {
		DeleteLinesFromConsole(deleteLines[0])
	}

	return text
}

func DeleteLinesFromConsole(lines int) {
	for i := 0; i < lines; i++ {
		// Move up the cursor and clear the line
		fmt.Print("\033[A\033[K")
	}
}

type postcode struct {
	numbers int
	letters string
}
type address struct {
	street   string
	number   int
	postcode postcode
}
type person struct {
	name    string
	age     int
	height  float64
	weight  float64
	address address
}

func createPerson() person {
	fmt.Println("Welcome to the person creator!")
	fmt.Println("Please fill in the following information:")

	fmt.Println("What is your name? Ex: John Doe")
	var name = ReadConsole(2)
	fmt.Println("What is your age? Ex: 20")
	var age = ReadConsole(2)
	fmt.Println("What is your weight? Ex: 75.5 (In kilograms)")
	var weight = ReadConsole(2)
	fmt.Println("What is your height? Ex: 1.75")
	var height = ReadConsole(2)

	// ask for address
	fmt.Println("What is your street name? Ex: Main Street")
	var street = ReadConsole(2)
	fmt.Println("What is your house number? Ex: 16")
	var number = ReadConsole(2)
	fmt.Println("What is your postcode? Ex: 1234AB")
	var postcodeQ = ReadConsole(2)

	var numbers = postcodeQ[0:4]
	var letters = postcodeQ[4:6]

	var ageConverted, _ = strconv.Atoi(age)
	var heightConverted, _ = strconv.ParseFloat(height, 64)
	var weightConverted, _ = strconv.ParseFloat(weight, 64)
	var numberConverted, _ = strconv.Atoi(number)
	var numbersConverted, _ = strconv.Atoi(numbers)

	var newPerson = person{
		name:   name,
		age:    ageConverted,
		height: heightConverted,
		weight: weightConverted,
		address: address{
			street: street,
			number: numberConverted,
			postcode: postcode{
				numbers: numbersConverted,
				letters: letters,
			},
		},
	}

	return newPerson
}
