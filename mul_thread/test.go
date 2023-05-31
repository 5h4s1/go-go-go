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
    url := "https://aws-uat-int-api.mcredit.com.vn/externalchannel/mc-admin-api/api/links"
    token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6InBlbnRlc3QwMDFAbWNyZWRpdC5jb20udm4iLCJGdWxsTmFtZSI6InBlbnRlc3QwMDEiLCJLZXlMb2dpbiI6ImY5MWQ2MzhhLTdiZjgtNGE1Ni1hMzNmLWIzNTYwNjQ4MWRlMyIsIlBob25lIjoiIiwiU3RhdHVzIjoiYWN0aXZlIiwiZXhwIjoxNjg1NTIzOTIxfQ.xRjcc9KDw784SGfCwxl1EGg72wSgU4IUMCa3SeEND_Q"
    requestPost(url, token)
}
