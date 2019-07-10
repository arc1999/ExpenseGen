package expenses

import (
	"time"
)

//go:generate reform
//reform:ExpenseTable
type Expense struct {
	ID          int32     `reform:"id,pk"`
	Description string    `reform:"description"`
	Type        string    `reform:"type"`
	Amount      float64   `reform:"amount"`
	CreatedOn   time.Time `reform:"created_on" `
	UpdatedOn   time.Time `reform:"updated_on"`
}


