package bankex

var deposits = make(chan int)
var balances = make(chan int)
var withdraws = make(chan int)
var withdrawStatus = make(chan bool)

func Withdraw(amount int) bool {
	withdraws <- amount
	return <-withdrawStatus
}

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balances
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case amount := <-withdraws:
			if balance < amount {
				withdrawStatus <- false
				continue
			}
			balance -= amount
			withdrawStatus <- true
		}
	}
}

func init() {
	go teller()
}
