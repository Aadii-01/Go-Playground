package features

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	Name  string
	PRN   string
	Marks int
}
type Class struct {
	Students []Student
}

func (c *Class) NewClass() {

	myStudents := []Student{}

	file, err := os.Open("db.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, " ")
		marks, err := strconv.Atoi(fields[2])

		if err != nil {
			panic(err)
		}

		obj := Student{
			Name:  fields[0],
			PRN:   fields[1],
			Marks: marks,
		}

		myStudents = append(myStudents, obj)
	}
	fmt.Println(myStudents)
}

func (C *Class) WriteToFile() {
	for _, y := range C.Students {
		line := fmt.Sprintf("%v %v %v", v.Name, v.PRN, v.Marks)
		fmt.Println(line)
	}
}

func (c *Class) ShowStudents() {
	for _, val := range c.Students {
		name := val.Name
		prn := val.PRN
		marks := val.Marks
		fmt.Println(name, " ", prn, " ", marks, " ")
	}
}

func (c *Class) AddStudent() {
	fmt.Println("Enter name")
	var name string
	fmt.Scanln(&name)

	fmt.Println("Enter PRN")
	var prn string
	fmt.Scanln(&prn)

	fmt.Println("Enter marks")
	var marks int
	fmt.Scanln(&marks)

	obj := Student{
		Name:  name,
		PRN:   prn,
		Marks: marks,
	}
	c.Students = append(c.Students, obj)

	c.WriteToFile()

	//Writing data to the file

	fmt.Println("Data Added")
}

func (c *Class) UpdateStudent() {
	fmt.Println("Enter PRN")
	var prn string
	fmt.Scanln(&prn)

	for idx, val := range c.Students {

		Prn := val.PRN

		if prn == Prn {
			fmt.Println("Enter name")
			var name string
			fmt.Scanln(&name)

			fmt.Println("Enter PRN")
			var prn string
			fmt.Scanln(&prn)

			fmt.Println("Enter marks")
			var marks int
			fmt.Scanln(&marks)
			c.Students[idx] = Student{Name: name, PRN: Prn, Marks: marks}
		}
	}

	fmt.Println("Data Updated")
}

func (c *Class) DeleteStudent() {
	fmt.Println("Enter PRN")
	var prn string
	fmt.Scanln(&prn)

	var idx int
	for i, val := range c.Students {
		if val.PRN == prn {
			idx = i
			break
		}
	}
	tempArr := c.Students
	tempArr = append(tempArr[:idx], tempArr[idx+1:]...)
	c.Students = tempArr

	fmt.Println("Data Delted")
}
