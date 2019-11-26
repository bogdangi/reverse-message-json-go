package main

import (
    "fmt"
    "os"

    "encoding/json"
    "net/http"
)

type Message struct {
    Id      string `json:"id"`
    Message string `json:"message"`
}

type API struct {
    Client  *http.Client
    baseURL string
}

func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func (api *API) ReverseMessage(res http.ResponseWriter, req *http.Request) {
    resp, err := api.Client.Get(api.baseURL + "/")
    if err != nil {
        http.Error(res, err.Error(), http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()

    var message Message

    err = json.NewDecoder(resp.Body).Decode(&message)
    if err != nil {
        http.Error(res, err.Error(), http.StatusInternalServerError)
        return
    }

    message.Message = Reverse(message.Message) 

    js, err := json.Marshal(message)

    if err != nil {
        http.Error(res, err.Error(), http.StatusInternalServerError)
        return
    }

    res.Header().Set("Content-Type", "application/json")
    res.Write(js)
}

func Healthz(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(res, "OK")
}

func main() {
    api := API{&http.Client{}, os.Getenv("MESSAGE_SERVICE_URL")}

    http.HandleFunc("/", api.ReverseMessage)
    http.HandleFunc("/healthz", Healthz)
    http.ListenAndServe(":3000", nil)
}
