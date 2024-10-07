package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Transaction struct {
	ID       int
	Amount   float64
	Category string
	Date     time.Time
	Type     string
}

type BudgetTracker struct {
	transactions []Transaction
	nextID   int
}

type FinancialRecord interface {
	GetAmount() float64
	GetType()   string
}

func (t Transaction) GetAmount() float64 {
	return t.Amount
}

func (t Transaction) GetType() string {
	return t.Type
}

func (bt *BudgetTracker) Addtransaction(amount float64, category string, tType string)  {
	 newTransaction := Transaction {
		ID: bt.nextID,
		Amount: amount,
		Category: category,
        Date: time.Now(),
		Type: tType,
	}

	bt.transactions = append(bt.transactions, newTransaction)
	bt.nextID++
}

func (bt *BudgetTracker) DisplayTransactions()  {
   
	for _, transaction := range bt.transactions {
       fmt.Printf("%d\t%.2f\t%s\t%s\t%s\n", 
	transaction.ID, transaction.Amount, transaction.Category,
	transaction.Date.Format("2006-01-02"), transaction.Type)
	
	}
}

func (bt *BudgetTracker) CalculateTotal(tType string) float64  {
    var total float64
	for _, transaction := range bt.transactions {
		if transaction.Type == tType {
			total += transaction.Amount
		}
	 }
	 return total
}

func (bt *BudgetTracker) SaveToCSV(filename string) error  {
    
	 file, err := os.Create(filename)

	if err != nil {
		return err
	 }

	 defer file.Close();	 
	 writer := csv.NewWriter(file) // creating a new csv
	 defer writer.Flush()
     
	 writer.Write([]string{"ID", "Amount", "Category", "Date", "Type"})

	 for _, t := range bt.transactions {

		 record := []string {
			strconv.Itoa(t.ID),
			fmt.Sprintf("%.2f", t.Amount),
			t.Category,
			t.Date.Format("2006-01-02"),
            t.Type,
		 }

		 writer.Write(record)
	 }
     
	 fmt.Println("transactions saved to", filename)
     return nil	
}

func main() {

	bt := BudgetTracker{}
   for {
	fmt.Println("\n ---Personal Budget Tracker---")
	fmt.Println("1. Add transaction")
	fmt.Println("2. Display transaction")
	fmt.Println("3. Show total income")
	fmt.Println("4. Show total expense")
	fmt.Println("5. Save to csv")
	fmt.Println("6. Exit")
	fmt.Println("7. Choose an option")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		fmt.Print("Enter Amount")
		var amount float64
		fmt.Scanln(&amount)

		fmt.Print("Enter Category")
		var category string
		fmt.Scanln(&category)

		fmt.Print("Enter type (Income / Expense)")
		var theType string
		fmt.Scanln(&theType)

        bt.Addtransaction(amount, category, theType)
		fmt.Println("Transaction Added")
	case 2:
		bt.DisplayTransactions()
	case 3:
		fmt.Printf("Total income: %.2f\n", bt.CalculateTotal("income"))
	case 4:
		fmt.Printf("Total income: %.2f\n", bt.CalculateTotal("expense"))
	case 5:
		fmt.Printf("Enter filenmae (e.g transactions.csv)")
		var filename string
		fmt.Scanln(filename)

		if err := bt.SaveToCSV(filename); err != nil {
			fmt.Println("Error saving transactions:",err)
		}
	case 6:
	   fmt.Println("Existing..")
	   return
	default:
		fmt.Println("Invalid answer. Try again")
	}

   }
	
}