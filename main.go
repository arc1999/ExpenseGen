package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"
	"github.com/gobuffalo/packr/v2"
)
type Json struct{
	Struct string


}
// go:generate main -package="ExpenseGen" -Opname="Create" -Opname1="Update" -Opname2="Delete" -Opname3="GetbyId" -Opname4="GetAll"
func main() {
	var struc string

	flag.StringVar(&struc,"Structurename","","name of the structure")
	flag.Parse()
	box:= packr.New("temp","./templates")
	t,err:=box.FindString("request-create.gotpl")
	if err != nil {
		log.Fatal(err)
	}
	tpl,err:= template.New("request").Parse(t)
	if err != nil {
		fmt.Println(err)
	}
	data :=Json{
		Struct: struc,
	}
	file,err := os.Create("./request/request.go")
	err= tpl.Execute(file,data)
	if err != nil {
		fmt.Println(err)
	}
}