package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
    "strings"
    "strconv"
	"math/rand"
    "time"
)


func requestPost(url string, token string){
	rand.Seed(time.Now().UnixNano())

    // Tạo một số ngẫu nhiên có độ dài 16
    num := rand.Intn(9) + 1
    for i := 0; i < 15; i++ {
        num *= 10
        num += rand.Intn(10)
    }

    filename := "./data.txt"
    jsonData, err := ioutil.ReadFile(filename)

    data := strings.Replace(string(jsonData), "{}",strconv.Itoa(num), -1)
    request, err := http.NewRequest("POST", url, bytes.NewReader([]byte(data)))
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    request.Header.Set("Authorization", "Bearer " + token)

    client := &http.Client{}

    response, err := client.Do(request)
    defer response.Body.Close()

    responseData, err := ioutil.ReadAll(response.Body)

    fmt.Println(string(responseData))
}

func main() {
    url := ""
    token := ""
    requestPost(url, token)
}
