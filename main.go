package main

import (
	"flag"
	"fmt"
	"github.com/gertd/go-pluralize"
	"github.com/gobuffalo/packr/v2"
	"html/template"
	"log"
	"os"
	"strings"
)
type Json struct{
	Struct string
	SliceStruct string
	Package string

}
// go:generate main -package="ExpenseGen" -Opname="Create" -Opname1="Update" -Opname2="Delete" -Opname3="GetbyId" -Opname4="GetAll"
func main() {
	var struc string

	flag.StringVar(&struc,"Structurename","","name of the structure")
	flag.Parse()
	box:= packr.New("temp","./templates")
	t,err:=box.FindString("request-create.gotpl")
	t1,err:=box.FindString("responsegen.gotpl")
	t2,err:=box.FindString("crud-op.gotpl")

	if err != nil {
		log.Fatal(err)
	}
	tpl,err:= template.New("request").Parse(t)
	if err != nil {
		fmt.Println(err)
	}
	tpl1,err1:= template.New("request").Parse(t1)
	if err1 != nil {
		fmt.Println(err)
	}
	tpl2,err2:= template.New("request").Parse(t2)
	if err2 != nil {
		fmt.Println(err)
	}
	pluralize := pluralize.NewClient()
	StructSlice:=pluralize.Plural(struc)
	pkgname:=strings.ToLower(StructSlice)
	data :=Json{
		Struct: struc,
		SliceStruct :StructSlice,
		Package:pkgname,

	}
	file,err := os.Create("./"+pkgname+"/request.go")
	file1,err := os.Create("./"+pkgname+"/response.go")
	file2,err := os.Create("./"+pkgname+"/crud.go")

	err= tpl.Execute(file,data)
	err= tpl1.Execute(file1,data)
	err= tpl2.Execute(file2,data)
	if err != nil {
		fmt.Println(err)

}
