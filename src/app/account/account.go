package account

import (
	"sync"
)

// Банковский счет
type Account struct {
	AccountPrivate
	AccountPublic
}

// Скрытые поля счета от пользователя
type AccountPrivate struct {
	mutex  sync.Mutex // симофор синхронизации потоков
	active bool       // флаг открытия закрытия счета
}

// Открытые поля счета пользователя
type AccountPublic struct {
	balance int64 // баланс счета
}

// Создание нового счета
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{AccountPrivate{active: true}, AccountPublic{balance: initialDeposit}}
}

// Получение баланса
func (a *Account) Balance() (balance int64, ok bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	if !a.active {
		return 0, false
	}

	return a.balance, true
}

// Зачисление (снятие) средств на счет
func (a *Account) Deposit(amount int64) (balance int64, ok bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	if (a.balance+amount) < 0 || !a.active {
		return 0, false
	}
	a.balance += amount
	return a.balance, true
}

// Закрытие счета
func (a *Account) Close() (payout int64, ok bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	if !a.active {
		return 0, false
	}

	a.active = false
	return a.balance, true
}
