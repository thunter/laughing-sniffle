package main

import (
	"os"
	"io/ioutil"
	"log"
	"fmt"
	"net/http"
	"flag"
)

const zKillboardURL = "https://redisq.zkillboard.com/listen.php"


var (
	tee_file *os.File
	err error
)

func main() {

	tee_filename := flag.String("s","","Tee output to givenfilename")
	append_tee := flag.Bool("a", false, "Append to existing output (if not given, -s will overwrite file")
	
	flag.Parse()

	write_output := false
	if *tee_filename != "" {
		write_output = true
	}

	if write_output {
		file_mode := os.O_CREATE|os.O_WRONLY|os.O_TRUNC
		if *append_tee {
			file_mode = os.O_APPEND|os.O_CREATE|os.O_WRONLY
		}
		fmt.Println("file_mode = %d", file_mode)
		tee_file, err = os.OpenFile(*tee_filename, file_mode, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}

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

		killmail, err := parseKillMail(body)

		if killmail.Package == nil {
			fmt.Printf("Empty result, restarting")
			continue
		}

		// fmt.Printf("%d %d Processed JSON: %#v", i,  killmail.Package.KillId, killmail)
		if write_output {
			tee_file.Write(body)
			tee_file.Write([]byte("\n"))
		}
		i++
	}
}