package app

import (
	"github.com/aystream/bank-account-exercises/src/app/db"
	"github.com/aystream/bank-account-exercises/src/app/handler"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// Приложение
type App struct {
	Router *mux.Router //экземпляры маршрутизатора
	DB     *db.DB
	// TODO таже можно добавать базу данных и какие то другие настройки
}

// Инициализация приложение TODO также если была бы какая нибудь база, то можно было бы её тут иницилизировать
func (a *App) Initialize() {
	// Явно ининцилизируем нашу БД
	a.DB = &db.DB{Account: nil}

	a.Router = mux.NewRouter()
	a.setRouters()
}

// Устанавливает все необходимые маршрутизаторы
func (a *App) setRouters() {
	a.Post("/account", a.CreateAccount)
	a.Put("/account", a.DepositAccount)
	a.Get("/account", a.GetBalanceAccount)
	a.Delete("/account", a.CloseAccount)
}

// маршрутизатор для метода GET
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// маршрутизатор для метода POST
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// маршрутизатор для метода PUT
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// маршрутизатор для метода DELETE
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

/*
 * Handler-ы банковского счета
 */
func (a *App) CreateAccount(w http.ResponseWriter, r *http.Request) {
	handler.CreateAccount(a.DB, w, r)
}
func (a *App) GetBalanceAccount(w http.ResponseWriter, r *http.Request) {
	handler.GetBalanceAccount(a.DB, w, r)
}

func (a *App) DepositAccount(w http.ResponseWriter, r *http.Request) {
	handler.DepositAccount(a.DB, w, r)
}

func (a *App) CloseAccount(w http.ResponseWriter, r *http.Request) {
	handler.CloseAccount(a.DB, w, r)
}

// Запустите приложение по определенном порту
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
