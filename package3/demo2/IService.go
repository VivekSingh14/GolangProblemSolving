package demo2

import "fmt"

type IPerson interface {
	studentDetails(string, int) (string, error)
}

type Person struct {
	Id         int
	Name, City string
}

func (d1 Person) studentDetails(street string, age int) (string, error) {
	fmt.Println("Address: ", d1.City, " ", street)
	fmt.Println("Name and Age: ", d1.Id, " ", d1.Name, " ", age)
	return "Vivek", nil
}

func PersonService() {
	var person2 IPerson

	person2 = Person{
		Id:   1,
		Name: "Vivek",
		City: "Jaipur",
	}

	_, err := person2.studentDetails("Shastri Nagar", 23)

	if err == nil {
		fmt.Println("Good")
	}
}
