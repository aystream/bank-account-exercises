package handler

import (
	"encoding/json"
	"github.com/aystream/bank-account-exercises/src/app/account"
	"github.com/aystream/bank-account-exercises/src/app/db"
	"net/http"
)

// Получение баланса аккаунта
func GetBalanceAccount(db *db.DB, w http.ResponseWriter, r *http.Request) {
	if db.Account == nil {
		respondError(w, http.StatusBadRequest, "Account is not created")
		return
	}

	balance, active := db.Account.Balance()
	if active == false {
		respondError(w, http.StatusBadRequest, "Account is closed")
		return
	}
	respondJSON(w, http.StatusOK, struct {
		Balance int64 `json:"amount"`
	}{Balance: balance})
}

// Создание аккаунта
func CreateAccount(db *db.DB, w http.ResponseWriter, r *http.Request) {
	initBalance := struct {
		InitialAmount int64 `json:"initialAmount"`
	}{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&initBalance); err != nil || initBalance.InitialAmount <= 0 {
		respondError(w, http.StatusBadRequest, "Invalid request parameters")
		return
	}
	defer r.Body.Close()

	newAccount := account.Open(initBalance.InitialAmount)
	db.SaveAccount(newAccount)
	respondJSON(w, http.StatusOK, nil)
}

// Зачисление (снятие) средств на счет
func DepositAccount(db *db.DB, w http.ResponseWriter, r *http.Request) {
	initDeposit := struct {
		Amount int64 `json:"amount"`
	}{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&initDeposit); err != nil || initDeposit.Amount == 0 {
		respondError(w, http.StatusBadRequest, "Invalid request parameters")
		return
	}
	defer r.Body.Close()

	_, ok := db.Account.Deposit(initDeposit.Amount)
	if ok == false {
		respondError(w, http.StatusBadRequest, "Not enough money")
		return
	}

	respondJSON(w, http.StatusOK, nil)
}

// Закрытие аккаунта
func CloseAccount(db *db.DB, w http.ResponseWriter, r *http.Request) {
	_, ok := db.Account.Close()
	if ok == false {
		respondError(w, http.StatusBadRequest, "Account is not created")
		return
	}

	respondJSON(w, http.StatusOK, nil)
}
