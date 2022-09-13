package account

import "sync"

type Account struct {
    balance int64
	open bool
	mu sync.RWMutex
}

func Open(amount int64) *Account {
	if amount < 0 {
        return nil
    }
	return &Account{balance: amount, open: true}
}

func (a *Account) Balance() (int64, bool) {
	a.mu.RLock()
    defer a.mu.RUnlock()
	if a.open == false {
		return a.balance, false
	}
	return a.balance, true

}

// func (a *Account) Deposit(amount int64) (int64, bool) {
// 	a.mu.Lock()
//     defer a.mu.Unlock()
//     if amount < 0 {
//         if a.balance < amount {
//             return a.balance, false
//         } else {
// 			a.balance -= amount
//     		return a.balance, true    
//     	}
// 	} else if a.open == false {
// 		return a.balance, false
// 	} else {
// 		a.balance += amount 
// 		return a.balance, true
// 	}
// }

func (a *Account) Deposit(amount int64) (int64, bool) {
	a.mu.Lock()
    defer a.mu.Unlock()
 	if a.open == false {
		return a.balance, false
	} else {
		if (a.balance + amount) < 0{
            return a.balance, false
        } else {
			a.balance += amount
    		return a.balance, true    
    	}
	}
}

func (a *Account) Close() (int64, bool) {
	a.mu.Lock()
    defer a.mu.Unlock()
	if a.open == false {
        a.balance = 0
		return a.balance, false
	}
	a.open = false
    balance := a.balance
    a.balance = 0
	return balance, true
}