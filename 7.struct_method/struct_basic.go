package main
import (
	"fmt"
	"strings"
)

type Person struct {
	firstName   string
	lastName    string
}

func upPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

func upPersonNoP(p Person){
	(&p).firstName = strings.ToUpper((&p).firstName)
	(&p).lastName = strings.ToUpper((&p).lastName)
}

func main() {
	// 1-struct as a value type:
	var pers1 Person
	pers1.firstName = "Chris"
	pers1.lastName = "Woodward"
	upPerson(&pers1)
	fmt.Printf("The name of the person is %s %s\n", pers1.firstName, pers1.lastName)

	// 2—struct as a pointer:
	pers2 := new(Person)
	pers2.firstName = "Chris"
	pers2.lastName = "Woodward"
	(*pers2).lastName = "Woodward"  // 这是合法的
	upPerson(pers2)
	fmt.Printf("The name of the person is %s %s\n", pers2.firstName, pers2.lastName)

	// 3—struct as a literal:
	pers3 := &Person{"Chris","Woodward"}
	upPerson(pers3)
	fmt.Printf("The name of the person is %s %s\n", pers3.firstName, pers3.lastName)

	pers4 := Person{"Yang","Jinyi"}
	pers4P := &pers4
	pers4.firstName="Xiao"
	pers4.lastName="Xiao"
	fmt.Println(pers4)
	fmt.Println(pers4P)
	pers4P.firstName="Yang"
	pers4P.lastName="Jinyi"
	fmt.Println(pers4)
	fmt.Println(pers4P)
	fmt.Println()

	//这样写还是不能改变pers4
	upPersonNoP(pers4)
	fmt.Println(pers4)
/*
	输出
	The name of the person is CHRIS WOODWARD
	The name of the person is CHRIS WOODWARD
	The name of the person is CHRIS WOODWARD
	{Xiao Xiao}
	&{Xiao Xiao}
	{Yang Jinyi}
   	&{Yang Jinyi}

   	{Yang Jinyi}
*/
}