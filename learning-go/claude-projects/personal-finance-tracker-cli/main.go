package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type finance interface {
	GetBalance() float64
	Deposit(amount float64)
	Withdraw(amount float64) bool
	seeTransHistory()
	SeeAccountDetailsPretty()
	GetAccountNumber()
}

type transaction interface {
	formateDate() string
}

type Account struct {
	name         string
	accType      string
	balance      float64
	accNumber    int
	transNumber  int
	transHistory map[int]*Transaction
}

type Transaction struct {
	transactionType string
	transactionAmt  float64
	transactionDate time.Time
	accNumber       int
}

func generateAccNum() int {
	return rand.Intn(900000) + 100000
}

func NewAccount(name, accType string, initialBalance float64) *Account {
	return &Account{
		name:         name,
		accType:      accType,
		balance:      initialBalance,
		transNumber:  0,
		transHistory: make(map[int]*Transaction),
		accNumber:    generateAccNum(),
	}
}

func NewTransaction(transactionType string, transactionAmt float64, accNumber int) *Transaction {
	return &Transaction{
		transactionType: transactionType,
		transactionAmt:  transactionAmt,
		transactionDate: time.Now(),
		accNumber:       accNumber,
	}
}

func (t *Transaction) formateDate() string {
	return t.transactionDate.Format("2006-01-02 15:04:05")
}

func (a *Account) GetAccountNumber() int {
	return a.accNumber
}

func (a *Account) GetBalance() float64 {
	return a.balance
}

func (a *Account) Deposit(depositAmt float64) {
	DepositTrans := NewTransaction("Deposit", depositAmt, a.GetAccountNumber())
	a.balance += depositAmt
	a.transNumber++
	a.transHistory[a.transNumber] = DepositTrans
}

func (a *Account) Withdraw(withdrawAmt float64) bool {
	WithdrawTrans := NewTransaction("Withdraw", withdrawAmt*-1, a.GetAccountNumber())
	if a.balance < withdrawAmt {
		fmt.Println("Insufficient funds")
		return false
	}
	a.balance -= withdrawAmt
	a.transNumber++
	a.transHistory[a.transNumber] = WithdrawTrans
	return true
}

func (a *Account) SeeAccountDetailsPretty() string {
	s := "Listing Acount Details:\n"
	s += "------------------------------\n"
	s += "Account Name: " + a.name + "\n"
	s += "Type of account: " + a.accType + "\n"
	s += "Account Number: " + strconv.FormatInt(int64(a.accNumber), 16) + "\n"
	s += "Balance: " + strconv.FormatFloat(a.balance, 'f', 2, 64) + "\n\n"
	s += SeeTransactionHistoryDetailsPretty(a)
	return s
}

func (a *Account) seeTransHistory() {
	fmt.Println(SeeTransactionHistoryDetailsPretty(a))
}

func SeeTransactionHistoryDetailsPretty(a *Account) string {
	s := ""
	for k, v := range a.transHistory {

		s += ("Transaction #" + strconv.Itoa(k) + " " + v.transactionType + " at " + v.formateDate() + " for " + strconv.FormatFloat(v.transactionAmt, 'f', 2, 64) + "\n")
	}
	return s
}

func main() {
	Bob := NewAccount("Bob", "checking", 100.0)

	Bob.Deposit(50.0)
	Bob.Withdraw(30.0)
	fmt.Println(Bob.SeeAccountDetailsPretty())

}
