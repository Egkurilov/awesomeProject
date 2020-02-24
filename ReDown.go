package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	getUrl("https://www.reddit.com/r/dataisbeautiful/comments/f0qulq/i_made_a_website_to_where_you_can_visualise.json")
}

func getUrl(url string) {

	fmt.Println("URL:>", url)

	req, err := http.NewRequest("GET", url, bytes.NewBuffer([]byte("")))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("User-agent", "your bot 0.1")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
