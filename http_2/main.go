package main

import (
	"fmt"
	"golang-ninja/coincap"
	"log"
	"time"
)

func main() {
	// # 1
	// resp, err := http.DefaultClient.Get("https://api.coincap.io/v2/assets")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer resp.Body.Close()
	// fmt.Println("Response status: ", resp.StatusCode)

	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// var r assetsResponse

	// if err = json.Unmarshal(body, &r); err != nil {
	// 	log.Fatal(err)
	// }

	// for _, asset := range r.Data {
	// 	fmt.Println(asset.Info())
	// }

	// #2

	coincapClient, err := coincap.NewClient(time.Second * 10)
	if err != nil {
		log.Fatal(err)
	}

	assets, err := coincapClient.GetAssests()
	if err != nil {
		log.Fatal(err)
	}
	for _, asset := range assets {
		fmt.Println(asset.Info())
	}

	bitcoin, err := coincapClient.GetAssest("bitcoin")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()
	fmt.Println(bitcoin.Info())
}
