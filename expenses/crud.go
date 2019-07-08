package expenses

import(
"net/http"
"context"
"fmt"
"errors"
"log"
"github.com/go-chi/chi"
"github.com/go-chi/chi/middleware"
"github.com/go-chi/render"
"github.com/jinzhu/gorm"
_ "github.com/go-sql-driver/mysql"
"time"
)
var obj Expense
var expenses Expenses
var db *gorm.DB
var err error
var req Createreq

func Init(){
      db, err = gorm.Open("mysql", "root:root@tcp(localhost:3306)/Expense1?charset=utf8&parseTime=True")
      	defer db.Close()
      	if err != nil {
      		fmt.Println(err)
      	}else{
      		fmt.Println("Connection established")
      	}
      	if(!db.HasTable(&Expense{}) ) {
      		db.AutoMigrate(&Expense{})
      	}
       r := chi.NewRouter()
           r.Use(middleware.RequestID)
           r.Use(middleware.RealIP)
           r.Use(middleware.Logger)
           r.Use(middleware.Recoverer)
           r.Use(render.SetContentType(render.ContentTypeJSON))
           r.Route("/expenses", func(r chi.Router) {
               r.Post("/", Create)
               r.Get("/", GetAll)
               r.Route("/{id}", func(r chi.Router) {
                   r.Use(CrudContext)
                   r.Get("/",GetId)
                   r.Put("/", Update)
                   r.Delete("/", Delete)
               })
           })
           log.Fatal(http.ListenAndServe(":8080", r))
   }

func CrudContext(next http.Handler) http.Handler {
return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

ID := chi.URLParam(r, "id")
		var temp Expense
		Db:= db.Table("expenses").Where("id = ?", ID).Find(&temp)

		if Db.RowsAffected == 0{
			err=errors.New("ID not Found")
			return
		} else{
			ctx := context.WithValue(r.Context(), "key", Db)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})

}
func Create(writer http.ResponseWriter , request *http.Request){
err = render.Bind(request, &req)
	temp:=*req.Expense
	temp.CreatedOn=time.Now()
	temp.UpdatedOn=time.Now()
	db.Create(&temp)
	render.Render(writer, request,List1(req.Expense))
}
func Update(writer http.ResponseWriter , request *http.Request){
db=request.Context().Value("key").(*gorm.DB)
    var Req Updatereq
	err:= render.Bind(request,&Req)
	if err != nil {
		fmt.Println(err)
		return
	}
	temp:=*Req.Expense
	temp.UpdatedOn=time.Now()
	dB:= db.Update(&temp)
    			if(dB.RowsAffected == 0){
    				err=errors.New("Expense not found")
    				fmt.Println(err)
    				return
    			}else{
    				_=render.Render(writer, request, List1(&temp))
    			}
 }
func Delete(writer http.ResponseWriter , request *http.Request){
 db=request.Context().Value("key").(*gorm.DB)
                  Db:= db.Delete(&obj)
             if(Db.RowsAffected == 0){
                 err=errors.New("Expense not found")
                 fmt.Println(err)
                 return
             }else{
                 fmt.Fprintf(writer,"sucessful delete")
                 return
             }
}
func GetAll(writer http.ResponseWriter , request *http.Request){
  db.Find(&expenses)
  fmt.Println(expenses)
  	err = render.Render(writer, request, ListAll(&expenses))
  	if err != nil {
  		fmt.Println(err)
  		return
  	}
}
func GetId(writer http.ResponseWriter , request *http.Request){
   db = request.Context().Value("key").(*gorm.DB)
   	dB:=db.Find(&obj)
   	if(dB.RowsAffected == 0){
   		err:=errors.New("Expense"+"not found")
   		fmt.Println(err)
   		return
   	}else{
   		_=render.Render(writer, request, List1(&obj))
   	}

}