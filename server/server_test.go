package server_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/Lucasbyte/DB-KV-GO/routes"
	"github.com/Lucasbyte/DB-KV-GO/server"
	"github.com/gin-gonic/gin"

	"github.com/stretchr/testify/assert"
)

func BenchmarkSETk(b *testing.B) {
	all := func(r *gin.Engine) {
		reqBodyGetAll := server.Request_Get{
			Method: "ALL",
		}
		jsonValueGetAll, _ := json.Marshal(reqBodyGetAll)
		req, _ := http.NewRequest("GET", "/all", bytes.NewBuffer(jsonValueGetAll))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
	}

	s := server.NewServer()
	router := routes.RunRoutes(s)

	var wg sync.WaitGroup
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			reqBody := server.Request_Post{
				Method: "SET",
				Key:    fmt.Sprintf("KEY %d", i),
				Value:  fmt.Sprintf("VALUE %d", i),
			}
			jsonValue, _ := json.Marshal(reqBody)
			req, _ := http.NewRequest("POST", "/set", bytes.NewBuffer(jsonValue))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
		}(i)
	}
	wg.Wait()

	all(router)

	var wg2 sync.WaitGroup

	for i := 0; i < b.N; i++ {
		wg2.Add(1)
		go func(i int) {
			defer wg2.Done()
			reqBody := server.Request_Post{
				Method: "DEL",
				Key:    fmt.Sprintf("KEY %d", i),
			}
			jsonValue, _ := json.Marshal(reqBody)

			req, _ := http.NewRequest("DELETE", "/delete", bytes.NewBuffer(jsonValue))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
		}(i)
	}
	wg2.Wait()

	all(router)
}

func TestSETk(t *testing.T) {
	s := server.NewServer()
	router := routes.RunRoutes(s)

	reqBody := server.Request_Post{
		Method: "SET",
		Key:    "foo",
		Value:  "bar",
	}
	jsonValue, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "/set", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, "Value set successfully!", response["message"])
	assert.Equal(t, "foo", response["key"])
	assert.Equal(t, "bar", response["value"])
}

func TestGETk(t *testing.T) {
	s := server.NewServer()
	router := routes.RunRoutes(s)

	// Primeiro, definimos um valor
	reqBody := server.Request_Post{
		Method: "SET",
		Key:    "foo",
		Value:  "bar",
	}
	jsonValue, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "/set", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Agora, recuperamos o valor
	reqBodyGet := server.Request_Get{
		Method: "GET",
		Key:    "foo",
	}
	jsonValueGet, _ := json.Marshal(reqBodyGet)
	req, _ = http.NewRequest("GET", "/get", bytes.NewBuffer(jsonValueGet))
	req.Header.Set("Content-Type", "application/json")

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, "Value get successfully!", response["message"])
	assert.Equal(t, "bar", response["value"])
}

func TestALLkv(t *testing.T) {
	s := server.NewServer()
	router := routes.RunRoutes(s)
	// Primeiro, definimos alguns valores
	reqBody1 := server.Request_Post{
		Method: "SET",
		Key:    "foo",
		Value:  "bar",
	}
	jsonValue1, _ := json.Marshal(reqBody1)
	req1, _ := http.NewRequest("POST", "/set", bytes.NewBuffer(jsonValue1))
	req1.Header.Set("Content-Type", "application/json")

	w1 := httptest.NewRecorder()
	router.ServeHTTP(w1, req1)

	reqBody2 := server.Request_Post{
		Method: "SET",
		Key:    "baz",
		Value:  "qux",
	}
	jsonValue2, _ := json.Marshal(reqBody2)
	req2, _ := http.NewRequest("POST", "/set", bytes.NewBuffer(jsonValue2))
	req2.Header.Set("Content-Type", "application/json")

	w2 := httptest.NewRecorder()
	router.ServeHTTP(w2, req2)

	// Agora, recuperamos todos os valores
	reqBodyGetAll := server.Request_Get{
		Method: "ALL",
	}
	jsonValueGetAll, _ := json.Marshal(reqBodyGetAll)
	req, _ := http.NewRequest("GET", "/all", bytes.NewBuffer(jsonValueGetAll))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, "Values get successfully!", response["message"])
	assert.NotNil(t, response["values"])
}
