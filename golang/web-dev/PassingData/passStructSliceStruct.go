package main

import (
	"log"
	"os"
)

type Student struct {
	Name  string
	Class string
}

type Class struct {
	ClassNumber      string
	NumberOfStudents int
}

type School struct {
	Students []Student
	Classes  []Class
}

func passStructSliceStruct() {

	student1 := Student{Name: "Benjmin", Class: "xi"}
	student2 := Student{Name: "James", Class: "xii"}
	class1 := Class{ClassNumber: "m2", NumberOfStudents: 7}
	class2 := Class{ClassNumber: "m3", NumberOfStudents: 13}
	var mayflower School
	mayflower.Classes = append(mayflower.Classes, class1)
	mayflower.Classes = append(mayflower.Classes, class2)
	mayflower.Students = append(mayflower.Students, student1)
	mayflower.Students = append(mayflower.Students, student2)
	err := tpl.ExecuteTemplate(os.Stdout, "structslicestruct.gohtml", mayflower)
	if err != nil {
		log.Fatalln(err)
	}
}
