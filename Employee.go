package main

import "fmt"

type Employee struct {
	Id           int
	Name         string
	DepartmentID int
	Projects     []Project
}

type Department struct {
	Id        int
	Name      string
	Employees []Employee
}

type Project struct {
	Id           int
	Name         string
	DepartmentID int
	Employees    []Employee
}

func (a *Employee) AssignToProject(projectID int) {
	project := fetchProjectById(projectID)
	if project != nil {
		a.Projects = append(a.Projects, *project)
	} else {
		fmt.Println("Project not found")
	}
}

func (b *Department) AddEmployee(employee Employee) {
	b.Employees = append(b.Employees, employee)
}

func (c *Department) GetEmployees() []Employee {
	return c.Employees
}

func (d *Project) AddEmployee(employee Employee) {
	d.Employees = append(d.Employees, employee)
}

func fetchProjectById(projectID int) *Project {
	return &Project{
		Id:           projectID,
		Name:         "Campus Project",
		DepartmentID: 1,
		Employees:    nil,
	}
}

func main() {

	employee := Employee{
		Id:           1,
		Name:         "Alice",
		DepartmentID: 1,
		Projects:     nil,
	}

	department:=Department{
		Id: 1,
		Name: "Enginnering",
		Employees: nil,
	}

	department.AddEmployee(employee)

	employees:=department.GetEmployees()
	fmt.Println(employees)

	employee.AssignToProject(642)

	fmt.Println(employee.Projects)



}
