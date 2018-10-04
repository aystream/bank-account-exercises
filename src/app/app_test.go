package app

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type TestCase struct {
	Method      string
	RequestBody string
	Request     string
	Response    string
	StatusCode  int
}

func TestApp(t *testing.T) {
	cases := []TestCase{
		TestCase{
			Method:      `GET`,
			RequestBody: ``,
			Request:     `GET /account`,
			Response:    `{"error":"Account is not created"}`,
			StatusCode:  http.StatusBadRequest,
		},
		TestCase{
			Method:      `POST`,
			RequestBody: `{"initialAmount":120}`,
			Request:     `POST /account`,
			Response:    ``,
			StatusCode:  http.StatusOK,
		},
		TestCase{
			Method:      `PUT`,
			RequestBody: `{"amount":-30}`,
			Request:     `PUT /account`,
			Response:    ``,
			StatusCode:  http.StatusOK,
		},
		TestCase{
			Method:      `PUT`,
			RequestBody: `{"amount":-100}`,
			Request:     `PUT /account`,
			Response:    `{"error":"Not enough money"}`,
			StatusCode:  http.StatusBadRequest,
		},
		TestCase{
			Method:      `GET`,
			RequestBody: ``,
			Request:     `GET /account`,
			Response:    `{"amount":90}`,
			StatusCode:  http.StatusOK,
		},
		TestCase{
			Method:      `DELETE`,
			RequestBody: ``,
			Request:     `DELETE /account`,
			Response:    ``,
			StatusCode:  http.StatusOK,
		},
		TestCase{
			Method:      `GET`,
			RequestBody: ``,
			Request:     `GET /account`,
			Response:    `{"error":"Account is closed"}`,
			StatusCode:  http.StatusBadRequest,
		},
	}

	newApp := &App{}
	newApp.Initialize()

	for caseNumber, item := range cases {

		req := httptest.NewRequest(item.Method, "http://localhost:3000/account", strings.NewReader(item.RequestBody))
		w := httptest.NewRecorder()

		switch item.Method {
		case `GET`:
			newApp.GetBalanceAccount(w, req)
			break
		case `POST`:
			newApp.CreateAccount(w, req)
			break
		case `PUT`:
			newApp.DepositAccount(w, req)
			break
		case `DELETE`:
			newApp.CloseAccount(w, req)
			break
		}

		if w.Code != item.StatusCode {
			t.Errorf("[%d] wrong StatusCode: got %d, expected %d", caseNumber, w.Code, item.StatusCode)
		}

		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)

		bodyStr := string(body)
		if bodyStr != item.Response {
			t.Errorf("[%d] wrong Response:\ngot: %+v\nexpected: %+v", caseNumber, bodyStr, item.Response)
		}
	}
}
