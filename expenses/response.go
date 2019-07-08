package expenses

import (
	"net/http"
)

type ListResponse struct {
	*Expense
}

func (ListResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func List1(exp *Expense) *ListResponse {
	resp := &ListResponse{Expense: exp}
	return resp
}

type ListAllResponse struct {
	Multiple *Expenses
}

func ListAll(m *Expenses) *ListAllResponse {
	return &ListAllResponse{Multiple: m}
}

func (e *ListAllResponse) Render(w http.ResponseWriter, r *http.Request) error {

	return nil
}
