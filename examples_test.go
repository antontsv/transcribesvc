package transcribesvc_test

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	svc "github.com/antontsv/transcribesvc"
)

func sample() {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	file := "samples/audio.flac"
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("cannot read sample file: %v", err)
	}
	res, err := svc.Transcribe(ctx, data)
	if err != nil {
		log.Fatalf("cannot transcribe: %v", err)
	}
	fmt.Println(res)
	//Output: hello Google
}
