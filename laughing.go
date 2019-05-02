package main

import (
	"os"
	"io/ioutil"
	"log"
	"fmt"
	"net/http"
	"flag"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"

)

const zKillboardURL = "https://redisq.zkillboard.com/listen.php"


var (
	tee_file *os.File
	err error

	tee_filename = flag.String("s","","Tee output to givenfilename")
	append_tee   = flag.Bool("a", false, "Append to existing output (if not given, -s will overwrite file")
	stream_name  = flag.String("ks","killboard","Name of kinesis stream to push to")


)

func main() {

	aws_session := session.Must(session.NewSession())
	
	kc := kinesis.New(aws_session)

	streams, err := kc.DescribeStream(&kinesis.DescribeStreamInput{StreamName: stream_name})
	if err != nil {
		panic(err)
	}
  fmt.Printf("%v\n", streams)
	
	

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
		fmt.Println("file_mode = ", file_mode)
		tee_file, err = os.OpenFile(*tee_filename, file_mode, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer tee_file.Close()
	}

	i := 0
	for true {  // Main loop
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

		putOutput, err := kc.PutRecord(&kinesis.PutRecordInput{
			Data: body,
			StreamName: stream_name,
			PartitionKey: aws.String("kill"),
		})

		if err != nil {
			panic(err)
		}
		fmt.Printf("%v\n", putOutput)

		// fmt.Printf("%d %d Processed JSON: %#v", i,  killmail.Package.KillId, killmail)
		if write_output {
			tee_file.Write(body)
			tee_file.Write([]byte("\n"))
		}
		i++
	}
}