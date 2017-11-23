package bank_account

import "test"
import "testing"

func TestAccountAcceptance(t *testing.T) {

  var account = NewBankAccount()

  account.Deposit(1000 , "10-01-2012")
  account.Deposit(2000 , "13-01-2012")
  account.Withdrawal(500 ,"14-01-2012")

  var moveList = account.Print()

  test.AssertEquals(
    "date || credit || debit || balance\n"+
    "14-01-2012 || || 500 || 2500\n" +
    "13-01-2012 || 2000 || || 3000\n" +
    "10-01-2012 || 1000 || || 1000", moveList, t)

}
