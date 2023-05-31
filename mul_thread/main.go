package main

import (
	"bytes"
	"fmt"
	"os"
	"net/http"
	"io/ioutil"
	"strconv"
	"time"
	"math/rand"
    "strings"
)

func sendRequest(url string, token string) {
	rand.Seed(time.Now().UnixNano())

    // Tạo một số ngẫu nhiên có độ dài 16
    num := rand.Intn(9) + 1
    for i := 0; i < 15; i++ {
        num *= 10
        num += rand.Intn(10)
    }
	req, err := http.NewRequest("GET", url + strconv.Itoa(num) + "/2", nil)

	if err != nil {
		fmt.Println("Error")
	}
	req.Header.Set("Authorization", "Bearer " + token)

	client := &http.Client{}
	beforeTime := time.Now().Unix()
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error")
	}
	body, _ := ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()
	fmt.Println(time.Now().Unix() - beforeTime)
	fmt.Println(string(body))
}

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
	if len(os.Args) < 4 {
        fmt.Println("Usage: go run main.go [url] [token] [thread] [count]")
        return
    }
	url := os.Args[1]
	token := os.Args[2]
	thread, _ := strconv.Atoi(os.Args[3])
	count, _ := strconv.Atoi(os.Args[4])
	for i := 0; i < count; i++ {
		for j := 0; j < thread; j++ {
			go sendRequest(url,token)
		}

		time.Sleep(time.Second * 5)
	}
}
