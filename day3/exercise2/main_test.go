package main

import (
	"bytes"
	"fmt"
	"github.com/exercise2/Routes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetFunc(t *testing.T) {
	router:=Routes.SetupRouter()
	router.GET("/user-api/user/1")

	req,_ := http.NewRequest("GET", "/user-api/user/1",nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp,req)
	fmt.Println(resp.Code)
	if resp.Code != 200 {
		t.Errorf("handler returned unexpected body: got %v want %v",resp.Code,200)
	}
}



func TestPostfunc(t *testing.T) {
	// Define the content of the POST request
	router := Routes.SetupRouter()
	router.POST("/user-api/user/1")
	req,_ := http.NewRequest("POST", "/user-api/user/1", bytes.NewBuffer([]byte(`{"name":"manish","lastname":"prajapat","dob":"12345678","address":"xyz lane","subject":"maths","marks":"100"}`)))
	resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		fmt.Println(resp.Code)
	if resp.Code != 200 {
		t.Errorf("handler returned unexpected body: got %v want %v",resp.Code,200)
	}

}
