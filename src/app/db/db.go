package db

import "github.com/aystream/bank-account-exercises/src/app/account"

// Наша база данных для одного аккаунта (сиситем авторизации в тз не говорится, поэтому воспользуемся данным способом)
type DB struct {
	Account *account.Account
}

// Сохранение аккаунта
func (db *DB) SaveAccount(account *account.Account) {
	db.Account = account
}
