package bank_account


import "test"
import "testing"


func TestPrintEmptyAccount(t *testing.T) {

  var account = NewBankAccount()

  test.AssertEquals("date || credit || debit || balance\n" , account.Print() ,t)

}

func TestPrintOneDeposit(t *testing.T) {

  var account = NewBankAccount()

  account.Deposit(150 ,"10-01-2012")

  test.AssertEquals(
	"date || credit || debit || balance\n"+
	"10-01-2012 || 150 || || 150",
	account.Print() ,t)

}

func TestPrintTwoDeposit(t *testing.T) {

  var account = NewBankAccount()

  account.Deposit(150 ,"10-01-2012")
  account.Deposit(100 ,"13-01-2012")


  test.AssertEquals(
	    "date || credit || debit || balance\n"+
		"13-01-2012 || 100 || || 250\n" +
		"10-01-2012 || 150 || || 150",
	account.Print() ,t)

}

func TestPrint3Deposit(t *testing.T) {

  var account = NewBankAccount()

  account.Deposit(150 ,"10-01-2012")
  account.Deposit(100 ,"13-01-2012")
  account.Deposit(105 ,"17-01-2012")



  test.AssertEquals(
	"date || credit || debit || balance\n"+
		"17-01-2012 || 105 || || 355\n" +
		"13-01-2012 || 100 || || 250\n" +
		"10-01-2012 || 150 || || 150",
	account.Print() ,t)

}

func TestPrintWithdrawal(t *testing.T) {

  var account = NewBankAccount()

  account.Deposit(105 ,"10-01-2012")
  account.Withdrawal(100 ,"11-01-2012")

  test.AssertEquals(
	"date || credit || debit || balance\n"+
		"11-01-2012 || || 100 || 5\n" +
		"10-01-2012 || 105 || || 105",
	account.Print() ,t)

}