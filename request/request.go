package request
import (
	"github.com/shrikar007/ExpenseGen/types"
	"net/http"
)
type Createreq struct {
	*types.Expense
}


func (c *Createreq) Bind(r *http.Request) error {
	//TODO implement

	return nil
}
type Updatereq struct {
	*Createreq
}

func (u *Updatereq) Bind(r *http.Request) error {

	return u.Createreq.Bind(r)
}

