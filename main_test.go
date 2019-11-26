package main

import (
    "fmt"
    "net/http"
    "net/http/httptest"
    "testing"
)

func Test_ReverseMessage(t *testing.T) {
    req, err := http.NewRequest("GET", "http://localhost:3000/", nil)
    if err != nil {
        t.Fatal(err)
    }

    res := httptest.NewRecorder()

    server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
        rw.Write([]byte(`{"id":"1","message":"Hello world"}`))
    }))

    api := API{server.Client(), server.URL}

    api.ReverseMessage(res, req)

    exp := fmt.Sprintf(`{"id":"1","message":"dlrow olleH"}`)
    act := res.Body.String()
    if exp != act {
        t.Fatalf("Expected %s got %s", exp, act)
    }
}

func Test_Healthz(t *testing.T) {
    req, err := http.NewRequest("GET", "http://localhost:3000/healthz", nil)
    if err != nil {
        t.Fatal(err)
    }

    res := httptest.NewRecorder()
    Healthz(res, req)

    exp := fmt.Sprintf("OK")
    act := res.Body.String()
    if exp != act {
        t.Fatalf("Expected %s got %s", exp, act)
    }
}
