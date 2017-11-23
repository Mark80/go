package bank_account

import (
  "fmt"
  "strconv"
  "strings"
)

type BankAccount struct {
  Balance   int
  Movements []Movement
}

type Movement struct {
  amount          int
  date            string
  snapshotBalance int
}

func NewBankAccount() *BankAccount {
  return &BankAccount{}
}

func (account *BankAccount) Deposit(amount int, date string) {
  var balance = calculateNewBalance(account.Movements, amount)
  updateAccount(account, amount, balance, date)

}

func (account*BankAccount) Withdrawal(amount int, date string) {
  var balance = calculateNewBalance(account.Movements, -amount)
  updateAccount(account, - amount, balance, date)
}

func updateAccount(account *BankAccount, amount int, balance int, date string) {
  account.Movements = append(account.Movements, Movement{amount, date, balance})
  account.Balance = balance
}

func calculateNewBalance(movements []Movement, amount int) int {
  var balance = amount
  if (len(movements) > 0) {
	balance += movements[len(movements) - 1].snapshotBalance
  }
  return balance
}

func (account *BankAccount) Print() string {
  return header + formatMovements(account.Movements)
}

var header = "date || credit || debit || balance\n"

func formatMovements(movements []Movement) string {
  var rowFormat = "%s || %s|| %s|| %s"
  var rows = []string{}

  for index, move := range movements {
	if index != 0 {
	  rowFormat = "%s || %s|| %s|| %s" + "\n"
	}
	row := fmt.Sprintf(rowFormat, move.date, formatAmount(move.amount), formatAmount(-move.amount), formatBalance(move.snapshotBalance))
	rows = append([]string{row}, rows...)
  }

  return strings.Join(rows, "")

}

func formatAmount(amount int) string {
  result := ""
  if (amount > 0) {
	result = strconv.Itoa(amount) + " "
  }
  return result
}

func formatBalance(amount int) string {
  return strconv.Itoa(amount)
}
