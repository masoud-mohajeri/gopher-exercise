package apiCall

import (
	"encoding/json"
	"fmt"
	"github.com/gojek/heimdall/v7/httpclient"
	"io"
	"strconv"
	"sync"
)

func Heimdall() {
	baseUrl := "https://jsonplaceholder.typicode.com"
	var wg sync.WaitGroup

	client := httpclient.NewClient()
	getRes, err := client.Get(baseUrl+"/posts", nil)
	if err != nil {
		panic(err)
	}
	defer getRes.Body.Close()

	body, err := io.ReadAll(getRes.Body)

	posts := []Post{}
	unmarshallingError := json.Unmarshal(body, &posts)
	if unmarshallingError != nil {
		fmt.Println("unmarshalling error", unmarshallingError)
	}

	wg.Add(len(posts))

	for i, val := range posts {
		go func(p Post, index int) {
			_, err := client.Delete(baseUrl+"/posts/"+strconv.Itoa(p.Id), nil)
			if err != nil {
				panic(err)
			}
			defer getRes.Body.Close()
			fmt.Println("deleted: ", strconv.Itoa(p.Id))
			wg.Done()
		}(val, i)
	}

	wg.Wait()
}
