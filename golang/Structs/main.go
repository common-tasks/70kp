package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	cinfo     contact
}

type contact struct {
	email   string
	pincode string
}

func main() {
	fmt.Println("hello from structs")
	var c contact
	c.email = "draupadi.murmu@gmail.com"
	c.pincode = "101"
	p1 := person{"anurag", "kumar", c}
	// fmt.Println(p1)
	p1 = person{firstName: "michael", lastName: "topno"}
	fmt.Println(p1)

	// p4:=person{firstName: "ram",lastName: "naam"}
	// fmt.Println(p4);
	var p2 person
	p2.firstName = "james"
	p2.lastName = "murmu"
	// fmt.Printf("%+v", p2)
	// p2.print()
	var p3 person
	p3.cinfo.email="daggu.gmail.com"
	p3.cinfo.pincode="100"
	p3.firstName="Michale"
	p3.lastName="oraon"

	p3.updateName("gareeb")
	p3.print()

	pointp3:= &p3
	fmt.Println("memory address ")
	fmt.Println(*pointp3)
	pointp3.updateNamebyPointer("vinod")
	pointp3.print()

	var p4 person
	p4.firstName="jklf"
	p4.lastName= "hghg"
	c1 :=contact{email: "g",pincode: "11"}
	p4.cinfo = c1

	p4.updateNamebyPointer("kunar")
	p4.print()

	myslice := []string {"hello","hi","sukkhi"}
	updateSlice(myslice)
	fmt.Println(myslice)
    name:="kachchap"
	fmt.Println(*(&name))
	fmt.Scanf("h") // wait for the input

}

func updateSlice(slice []string){
slice[0]="bye"
}

func(p *person)updateNamebyPointer(newFirstname string){
	(*p).firstName=newFirstname
}
func (p person) updateName(newfirstName string) {
	p.firstName = newfirstName
	fmt.Print(p.firstName)
}

func (p person) print() {
	fmt.Println(p)
}
