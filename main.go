package main

import "fmt"

type employee interface {
	getName() string
	getSalary() int
}

type contractor struct {
	name         string
	hourlyPay    int
	hoursPerYear int
}

type fullTime struct {
	name   string
	salary int
}

func (c contractor) getName() string {
	return c.name
}
func (c contractor) getSalary() int {
	return c.hourlyPay * c.hoursPerYear
}

func (ft fullTime) getName() string {
	return ft.name
}
func (ft fullTime) getSalary() string {
	ft.getSalary()
}

func test(e employee) {
	fmt.Println("Employee Name:", e.getName())
	fmt.Println("Employee Salary:", e.getSalary())
	fmt.Println("------")
}

func main() {
	contractor1 := contractor{
		name:         "Yaqoob",
		hourlyPay:    40,
		hoursPerYear: 2000,
	}
	fullTime1 := fullTime{name: "Ali", salary: 80000}

	test(contractor1)
	test(fullTime1)
}
