package main

import (
	"Expense/ExpenseGen/expenses"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetall(t *testing.T) {

	req, err := http.NewRequest("GET", "/expenses", nil)
	if err != nil {
		fmt.Println(err)
	}
	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(expenses.GetAll)
	handler.ServeHTTP(rec, req)
	if status := rec.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `"Expenses":[{"ID":6,"Description":"ayush","Type":"ash","Amount":4415.32,"CreatedOn":"2019-06-21T02:25:38Z","UpdatedOn":"2019-06-24T06:49:46Z"},{"ID":7,"Description":"ayush","Type":"ash","Amount":4415.32,"CreatedOn":"2019-06-24T05:21:19Z","UpdatedOn":"2019-06-24T05:21:19Z"},{"ID":8,"Description":"ayush","Type":"ash","Amount":4415.32,"CreatedOn":"2019-06-24T05:34:00Z","UpdatedOn":"2019-06-24T05:34:00Z"},{"ID":9,"Description":"ayush","Type":"ash","Amount":4415.32,"CreatedOn":"2019-06-24T06:19:05Z","UpdatedOn":"2019-06-24T06:19:05Z"},{"ID":10,"Description":"Ayushand","Type":"ashooooo","Amount":12.32,"CreatedOn":"2019-06-24T07:40:06Z","UpdatedOn":"2019-07-10T11:52:04Z"},{"ID":11,"Description":"test","Type":"Tut","Amount":1000,"CreatedOn":"2019-07-09T12:59:58Z","UpdatedOn":"2019-07-09T12:59:58Z"},{"ID":12,"Description":"Ayushdove","Type":"ash","Amount":4455.32,"CreatedOn":"2019-07-10T07:18:52Z","UpdatedOn":"2019-07-10T07:18:52Z"}]`
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rec.Body.String(), expected)
	}

}
