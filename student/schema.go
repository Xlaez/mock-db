package student

import "fmt"

type Student struct {
	firstName string
	lastname  string
	gender    string
	Date
	email        string
	department   string
	matricNumber string
}

// setter and getter for first and last name
func (s *Student) SetName(first string, last string) error {
	if first == "" && last == "" {
		return fmt.Errorf("error occured! you haven't entered a name")
	}
	s.firstName = first
	s.lastname = last
	return nil
}

func (s *Student) FirstName() string {
	return s.firstName
}
func (s *Student) LastName() string {
	return s.lastname
}

// Set and get method for gender
func (s *Student) SetGender(gender string) error {
	if gender == "" {
		return fmt.Errorf("error occured! you haven't entered a name")
	}
	s.gender = gender
	return nil
}
func (s *Student) GetGender() string {
	return s.gender
}

// Set and get method for Email
func (s *Student) SetEmail(email string) error {
	if email == "" {
		return fmt.Errorf("error occured! you haven't entered an email")
	}
	s.email = email
	return nil
}
func (s *Student) GetEmail() string {
	return s.email
}

// Set and get method for department
func (s *Student) SetDept(department string) error {

	if department == "" {
		return fmt.Errorf("error occured! you haven't entered an Department")
	}
	s.department = department
	return nil
}
func (s *Student) GetDept() string {
	return s.department
}

// Set and get method for MatricNumber
func (s *Student) SetRegNo(id string) error {
	if id == "" {
		return fmt.Errorf("error occured! you haven't entered an email")
	}
	s.matricNumber = id
	return nil
}
func (s *Student) GetRegNo() string {
	return s.matricNumber
}
