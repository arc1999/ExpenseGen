package expenses

import(
"net/http"
"context"
"fmt"
"errors"
"log"
"strconv"
"github.com/go-chi/chi"
"github.com/go-chi/chi/middleware"
"github.com/go-chi/render"
)
var obj Expense
var expenses Expenses
func Init(){
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
	log.Fatal(http.ListenAndServe(":8084", r))
}

func CrudContext(next http.Handler) http.Handler {
return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        var err error
        flag:=1
		ID := chi.URLParam(r, "id")
		b,_:=strconv.Atoi(ID)
		for _, a := range expenses{
			if a.Id == b {
				flag=0
				ctx := context.WithValue(r.Context(), "key", a )
				next.ServeHTTP(w, r.WithContext(ctx))
			}
		}
		if flag ==1{
			err=errors.New("ID not Found")
			//render.Render(w, r, errrs.ErrRender(err))
			fmt.Println(err)
		}
		})
}
func Create(writer http.ResponseWriter , request *http.Request){
var req Createreq
	var err error
	err = render.Bind(request, &req)
	if err != nil {
		log.Println(err)
		return
	}
     expenses= append(expenses, *req.Expense)
	render.Render(writer, request, List1(req.Expense))
//TODO implementation

}
func Update(writer http.ResponseWriter , request *http.Request){
obj = request.Context().Value("key").(Expense)
var req Updatereq
err:= render.Bind(request,&req)
    if err != nil {
        log.Println(err)
        return
    }
// for loop -> get index of expense
        for index, e := range expenses {
            if e.Id == obj.Id {
                  expenses[index] = *req.Expense
            }
    }
        errs:=render.Render(writer, request,List1(&obj))
    if errs != nil {
       fmt.Println(err)
       return
    }
}
func Delete(writer http.ResponseWriter , request *http.Request){
obj = request.Context().Value("key").(Expense)
    expenses=append(expenses[:obj.Id], expenses[obj.Id+1:]...)
     err:=render.Render(writer,request,ListAll(&expenses))
    if err != nil {
        fmt.Println(err)
        return
            }
}
func GetAll(writer http.ResponseWriter , request *http.Request){
    err := render.Render(writer, request, ListAll(&expenses))
        if err != nil {
              panic(err)
             //render.Render(writer,request,errrs.ErrRender(err))
            return
    }

}
func GetId(writer http.ResponseWriter , request *http.Request){
            obj = request.Context().Value("key").(Expense)
            err:=render.Render(writer, request, List1(&obj))
                if err != nil {
         //render.Render(writer,request,errrs.ErrRender(err))
                   panic(err)
                   return
                   }
}

