package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Stringer interface {
	String() string
}

type NameNo struct {
	name string
	no int
	sex bool
}

func(this *NameNo) String() string{
	var sex string
	if this.sex {
		sex = "男"
	} else {
		sex = "女"
	}
	return strings.Join([]string{this.name, strconv.Itoa(this.no),sex},",")
}

func main(){
	var v Stringer
	nameNo := new(NameNo)
	nameNo.name = "mi-autumn"
	nameNo.no = 5
	nameNo.sex = false

	v = nameNo

	if sv, ok := v.(Stringer); ok {
		fmt.Printf("v implements String(): %s\n", sv.String()) // note: sv, not v
	}
}