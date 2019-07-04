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
	Packagename string
	Opname string
	Opname1 string
	Opname2 string
	Opname3 string
	Opname4 string

}
// go:generate main -package="ExpenseGen" -Opname="Create" -Opname1="Update" -Opname2="Delete" -Opname3="GetbyId" -Opname4="GetAll"
func main() {
	var pkgname,Opertname,Opertname1 ,Opertname2 ,Opertname3,Opertname4  string
	flag.StringVar(&pkgname,"package","","name of the package")
	flag.StringVar(&Opertname,"Opname","","name of the operation")
	flag.StringVar(&Opertname1,"Opname1","","name of the operation")
	flag.StringVar(&Opertname2,"Opname2","","name of the operation")
	flag.StringVar(&Opertname3,"Opname3","","name of the operation")
	flag.StringVar(&Opertname4,"Opname4","","name of the operation")
	flag.Parse()
	box:= packr.New("temp","./templates")
	t,err:=box.FindString("crud-op.gotpl")
	if err != nil {
		log.Fatal(err)
	}
	tpl,err:= template.New("expensegen").Parse(t)
	if err != nil {
		fmt.Println(err)
	}
	data :=Json{
		Packagename: pkgname,
		Opname: 	Opertname,
		Opname1: 	Opertname1,
		Opname2: 	Opertname2,
		Opname3: 	Opertname3,
		Opname4: 	Opertname4,
	}
	file,err := os.Create("./crud/crud.go")
	err= tpl.Execute(file,data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("shrikar")
	fmt.Println("pawas")
}