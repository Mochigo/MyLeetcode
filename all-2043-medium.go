package main

type Bank struct {
	accounts []int64
	UpBound  int
}

func Constructor(balance []int64) Bank {
	return Bank{balance, len(balance) + 1}
}

func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
	// fmt.Println(account1, " to ", account2, " transfer ", money)
	if account1 >= this.UpBound || account2 >= this.UpBound {
		return false
	}
	if this.accounts[account1-1] < money {
		return false
	}
	this.accounts[account1-1] -= money
	this.accounts[account2-1] += money
	return true
}

func (this *Bank) Deposit(account int, money int64) bool {
	if account >= this.UpBound {
		return false
	}
	this.accounts[account-1] += money
	return true
}

func (this *Bank) Withdraw(account int, money int64) bool {
	if account >= this.UpBound {
		return false
	}
	if this.accounts[account-1] < money {
		return false
	}
	this.accounts[account-1] -= money
	return true
}

/**
 * Your Bank object will be instantiated and called as such:
 * obj := Constructor(balance);
 * param_1 := obj.Transfer(account1,account2,money);
 * param_2 := obj.Deposit(account,money);
 * param_3 := obj.Withdraw(account,money);
 */
