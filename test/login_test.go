package test

import (
	"net/http"
	"testing"

	"github.com/jordanlanch/bia-test/domain"

	"github.com/jordanlanch/bia-test/test/e2e"
)

const (
	statusOK                  = http.StatusOK
	statusUnprocessableEntity = http.StatusUnprocessableEntity
	errorInvalidPayload       = "error.invalid_payload"
	errorLoginEventBaseUserID = "login.event_base.user_id"
	errorDescription          = "[ERROR] invalid payload [" + errorLoginEventBaseUserID + ": String length must be greater than or equal to 1]"
)

func TestLogin(t *testing.T) {
	expect, teardown := e2e.Setup(t, "fixtures/login_test")
	defer teardown()

	t.Run("Login success", func(t *testing.T) {
		t.Log("Iniciando prueba:", t.Name())

		in := domain.LoginRequest{
			Email:    "jordan.dev93@gmail.com",
			Password: "pass1234*",
		}

		obj := expect.POST("/api/v1/login").
			WithJSON(in).
			Expect().
			Status(http.StatusOK).
			JSON().Object()
		obj.Value("res").String().Equal("Successfully logged in!")
		t.Log("Finalizando prueba:", t.Name())
	})

	t.Run("Login failure", func(t *testing.T) {
		t.Log("Iniciando prueba:", t.Name())
		in := domain.LoginRequest{
			Email:    "jordan.dev93@gmail.com",
			Password: "pass1234*",
		}

		obj := expect.POST("/api/v1/login").
			WithJSON(in).
			Expect().
			Status(http.StatusOK).
			JSON().Object()
		obj.Value("res").String().Equal("Successfully logged in!")
		t.Log("Finalizando prueba:", t.Name())
	})

	t.Run("Login error with required field", func(t *testing.T) {
		t.Log("Iniciando prueba:", t.Name())
		in := domain.LoginRequest{
			Email: "jordan.dev93@gmail.com",
		}
		obj := expect.POST("/api/v1/login").
			WithJSON(in).
			Expect().
			Status(http.StatusUnprocessableEntity).
			JSON().Object()
		errorMsg := obj.Value("error").Object()
		errorMsg.Value("code").String().Equal("error.invalid_payload")
		errorMsg.Value("description").String().Equal("[ERROR] invalid payload [login.login_status: login.login_status must be one of the following: \"$failure\", \"$success\" login.login_status: String length must be greater than or equal to 1]")
		t.Log("Finalizando prueba:", t.Name())
	})

	t.Run("Login error with required field", func(t *testing.T) {
		t.Log("Iniciando prueba:", t.Name())
		in := domain.LoginRequest{
			Email: "jordan.dev93@gmail.com",
		}
		obj := expect.POST("/api/v1/login").
			WithJSON(in).
			Expect().
			Status(http.StatusUnprocessableEntity).
			JSON().Object()
		errorMsg := obj.Value("error").Object()
		errorMsg.Value("code").String().Equal("error.invalid_payload")
		errorMsg.Value("description").String().Equal("[ERROR] invalid payload [login.login_status: login.login_status must be one of the following: \"$failure\", \"$success\" login.login_status: String length must be greater than or equal to 1]")
		t.Log("Finalizando prueba:", t.Name())
	})
}
