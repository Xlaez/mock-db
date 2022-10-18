package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	v "student-reg/input"
	"student-reg/student"
	"text/tabwriter"
)

func main() {
	// Get number of students user wants to add
	fmt.Println("Enter the number of students to be registerd: ")
	num, err := v.NumChecker(v.InputValidatorInt())

	if err != nil {
		log.Fatal("An error has occured: ", err)
	}

	// create a hashmap database to hold user data
	addedStudents := map[string]student.Student{}

	value := student.Student{} //string value type to hold the keys for hashmap

	var name string

	// loop to accept data as long as the number inputed for students entry
	for {
		fmt.Print("Enter FirstName: ")
		_, err := fmt.Scanln(&name)
		if err != nil {
			log.Fatal(err)
		}
		name = strings.TrimSpace(name)

		// firstname and lastname
		fmt.Print("Enter LastName: ")
		err = value.SetName(name, v.InputValidatorStrings())
		if err != nil {
			log.Fatal(err)
		}

		//date of birth
		fmt.Println("Enter your date of birth")
		fmt.Print("Day: ")
		err = value.SetDay(v.InputValidatorInt())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print("Month: ")
		err = value.SetMonth(v.InputValidatorInt())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print("Year: ")
		err = value.SetYear(v.InputValidatorInt())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v-%v-%v\n", value.Day(), value.Month(), value.Year())

		//email
		fmt.Print("Enter a Valid Email Address: ")
		err = value.SetEmail(v.InputValidatorStrings())
		if err != nil {
			log.Fatal(err)
		}

		//department
		fmt.Print("Enter Student Department: ")
		err = value.SetDept(v.InputValidatorStrings())
		if err != nil {
			log.Fatal(err)
		}

		// matric Number
		fmt.Print("Enter Student Matric Number: ")
		err = value.SetRegNo(v.InputValidatorStrings())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("")

		fmt.Println("Enter the details for another student.")

		addedStudents[name] = value // add the values to addedStudents map

		if len(addedStudents) == num {
			break
		}
	}

	// print data from map in table format
	// var values student.Student

	w := tabwriter.NewWriter(os.Stdout, 10, 0, 3, ' ', tabwriter.Debug)

	fmt.Fprintln(w, "FirstName\tSurname\tGender\tDateOfBirth\tEmail\tDepartment\tMatric-No\t")

	for _, values := range addedStudents {
		fmt.Fprint(w, values.FirstName(), "\t", values.LastName(), "\t", values.GetGender(), "\t", values.Date, "\t", values.GetEmail(), "\t", values.GetDept(), "\t", values.GetRegNo(), "\t", "\n")

	}

	w.Flush()

	fmt.Println("")

	//Enabling deletion of student
	fmt.Println("Would you like to delete a student name?")
	fmt.Print("Please enter 'yes' or 'no': ")

	answer := v.InputValidatorStrings() //user enters yes or no
	fmt.Println("")

	var deleteName string // name of key to delete

	if answer == "yes" {
		fmt.Print("Enter a Student name: ")

		_, ok := addedStudents[deleteName] // check if student name is in the database

		if !ok {
			fmt.Println("")
			fmt.Printf("%q is not in the database\n", deleteName)
			fmt.Println("Please Enter a valid name")
		} else {
			delete(addedStudents, deleteName) // deletes student from database
			fmt.Fprintln(w, "FirstName\tSurname\tGender\tDateOfBirth\tEmail\tDepartment\tMatric-No\t")

			for _, values := range addedStudents {
				fmt.Fprint(w, values.FirstName(), "\t", values.LastName(), "\t", values.GetGender(), "\t", values.Date, "\t", values.GetEmail(), "\t", values.GetDept(), "\t", values.GetRegNo(), "\t", "\n")

			}
		}
	} else if answer == "no" {
		fmt.Println("No data was deleted")
	} else {
		fmt.Println("Invalid input! You're to enter yes or no!")
	}

	fmt.Println("")

}
