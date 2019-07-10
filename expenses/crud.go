package expenses

import(
"database/sql"
"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/mysql"
	"net/http"
"context"
"fmt"
"log"
"strconv"
"github.com/go-chi/chi"
"github.com/go-chi/chi/middleware"
"github.com/go-chi/render"
_ "github.com/go-sql-driver/mysql"
	"os"
	"time"
)
var obj Expense
type Expenses []Expense
var expenses Expenses
var db *reform.DB
var err error
var req Createreq

func Init(){
          d,err:=sql.Open("mysql", "root:root@tcp(localhost:3306)/Expense?charset=utf8&parseTime=True")
        	if err != nil {
        		fmt.Println("ayush")
        		fmt.Println(err)
        	}
               lg := log.New(os.Stderr, "SQL: ", log.Flags())
        	db= reform.NewDB(d, mysql.Dialect, reform.NewPrintfLogger(lg.Printf))

       r := chi.NewRouter()
           r.Use(middleware.RequestID)
           r.Use(middleware.RealIP)
           r.Use(middleware.Logger)
           r.Use(middleware.Recoverer)
           r.Use(render.SetContentType(render.ContentTypeJSON))
           r.Route("/expenses", func(r chi.Router) {
               r.Post("/", Create)
               r.Get("/",GetAll)
               r.Route("/{id}", func(r chi.Router) {
                   r.Use(CrudContext)
                   r.Get("/",GetId)
                   r.Put("/",Update)
               })
           })
           log.Fatal(http.ListenAndServe(":8080", r))
   }

func CrudContext(next http.Handler) http.Handler {
return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

ID := chi.URLParam(r, "id")
id,_:=strconv.Atoi(ID)
		temp,err :=db.FindByPrimaryKeyFrom(ExpenseTable,id)

		if err != nil {
        		fmt.Println(err)
        		return
        	}else{
			ctx := context.WithValue(r.Context(), "key", temp)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})

}
func Create(writer http.ResponseWriter , request *http.Request){
err = render.Bind(request, &req)
	temp:=*req.Expense
	temp.CreatedOn=time.Now()
	temp.UpdatedOn=time.Now()
	err=db.Save(&temp)
	if err != nil {
            		fmt.Println(err)
            		return
            	}
	render.Render(writer, request,List1(&temp))
}
func GetId(writer http.ResponseWriter , request *http.Request){
 	temp:= request.Context().Value("key").(*Expense)
 		_ = render.Render(writer, request,List1(temp))

}
func GetAll(writer http.ResponseWriter , request *http.Request){
flag:=1
    tables, err := db1.SelectRows(ExpenseTable, "WHERE id IS NOT NULL")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer tables.Close()

    for flag!=0{
        err = db1.NextRow(&obj, tables)
        expenses=append(expenses,obj)
        //fmt.Println(obj)
        if err!=nil{
            flag=0
            break
        }
    }
    _=render.Render(writer, request, ListAll(&expenses))
}
func Update(writer http.ResponseWriter , request *http.Request){
s:=request.Context().Value("key").(*Expense)
var upreq Updatereq
err:= render.Bind(request,&upreq)
  if err != nil {
      log.Println(err)
      return
  }
   var temp Expense
      temp=*upreq.Expense

           s.Description=temp.Description
                  s.Type=temp.Type
                  s.Amount=temp.Amount
                  s.UpdatedOn=time.Now()

              err1 := db.Update(s)

      if err1 != nil{
                 fmt.Println(err)
                 return
             }else{
                 err=render.Render(writer, request, List1(&temp))
                 fmt.Println(err)
                  }
}
func Delete(writer http.ResponseWriter , request *http.Request){
s:=request.Context().Value("key").(*Expense)
    err=db.Delete(s)
    if err != nil {
        panic(err)
    }else{
        _=render.Render(writer, request, List1(s))
    }
}
