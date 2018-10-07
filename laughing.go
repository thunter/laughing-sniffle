package main

import (
	"io/ioutil"
	"log"
	"fmt"
	"net/http"
)

const zKillboardURL = "https://redisq.zkillboard.com/listen.php"

func main() {
	i := 0
	for true {
		resp, err := http.Get(zKillboardURL)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(resp)
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%d Raw JSON: %s\n",i, body)

		killmail := parseKillMail(string(body))

		fmt.Printf("%d Processed JSON: %#v", i, killmail)
		i++
	}
}