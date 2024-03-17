package apiCall

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
)

func NetHttp() {
	baseUrl := "https://jsonplaceholder.typicode.com"
	var wg sync.WaitGroup

	getReq, err := http.NewRequest("GET", baseUrl+"/posts", nil)
	if err != nil {
		fmt.Print(err.Error())
	}

	getRes, err := http.DefaultClient.Do(getReq)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer getRes.Body.Close()
	body, readErr := io.ReadAll(getRes.Body)
	if readErr != nil {
		fmt.Print(err.Error())
	}
	//fmt.Println(body)
	posts := []Post{}
	json.Unmarshal(body, &posts)

	wg.Add(len(posts))

	for i, val := range posts {
		go func(p Post, n int) {
			deleteReq, err := http.NewRequest("DELETE", baseUrl+"/posts/"+strconv.Itoa(p.Id), nil)
			if err != nil {
				fmt.Println(err)
			}
			_, delErr := http.DefaultClient.Do(deleteReq)
			if delErr != nil {
				fmt.Print(delErr.Error())
				fmt.Println("error in deleting:", n)
			}
			fmt.Println("Deleted:", n)
			wg.Done()

		}(val, i)
	}

	wg.Wait()
}
