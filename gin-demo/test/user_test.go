package test

import (
	"bytes"
	// "fmt"
	"gin-demo/initRouter"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"

	// "strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var router *gin.Engine

func init() {
	gin.SetMode(gin.TestMode)
	router = initRouter.SetupRouter()
}

// func TestUserSave(t *testing.T) {
// 	username := "lisi"
// 	router := initRouter.SetupRouter()
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest(http.MethodGet, "/user/"+username, nil)
// 	router.ServeHTTP(w, req)
// 	assert.Equal(t, http.StatusOK, w.Code)
// 	assert.Equal(t, "用户"+username+"已经保存", w.Body.String())
// }

// func TestUserSaveQuery(t *testing.T) {
// 	username := "lisi"
// 	age := 18
// 	router := initRouter.SetupRouter()
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest(http.MethodGet, "/user?name="+username+"&age="+strconv.Itoa(age), nil)
// 	router.ServeHTTP(w, req)
// 	assert.Equal(t, http.StatusOK, w.Code)
// 	assert.Equal(t, "用户:"+username+",年龄:"+strconv.Itoa(age)+"已经保存", w.Body.String())
// }

// func TestUserPostForm(t *testing.T) {
// 	value := url.Values{}
// 	value.Add("email", "youngxhui@gmail.com")
// 	value.Add("password", "1234")
// 	value.Add("password-again", "1234")
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest(http.MethodPost, "/user/register", bytes.NewBufferString(value.Encode()))
// 	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; param=value")
// 	router.ServeHTTP(w, req)
// 	assert.Equal(t, http.StatusOK, w.Code)
// }

// func TestUserPostFormEmailErrorAndPasswordError(t *testing.T) {
// 	value := url.Values{}
// 	value.Add("email", "youngxhui")
// 	value.Add("password", "1234")
// 	value.Add("password-again", "qwer")
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest(http.MethodPost, "/user/register", bytes.NewBufferString(value.Encode()))
// 	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; param=value")
// 	router.ServeHTTP(w, req)
// 	assert.Equal(t, http.StatusOK, w.Code)
// }

func TestUserLogin(t *testing.T) {
	email := "youngxhui1"
	value := url.Values{}
	value.Add("email", email)
	value.Add("password", "1234")
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/user/login", bytes.NewBufferString(value.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; param=value")
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, strings.Contains(w.Body.String(), email), true)
}
