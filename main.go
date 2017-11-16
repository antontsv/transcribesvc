package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/Sirupsen/logrus"
	speech "github.com/google/go-genproto/googleapis/cloud/speech/v1"
	"google.golang.org/api/option"
	"google.golang.org/api/transport"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	conn, err := transport.DialGRPC(ctx,
		option.WithEndpoint("speech.googleapis.com:443"),
		option.WithScopes("https://www.googleapis.com/auth/cloud-platform"),
	)
	if err != nil {
		logrus.Fatalf("could not connect to gRPC server %v", err)
	}
	defer conn.Close()

	file := "samples/audio.flac"
	data, err := ioutil.ReadFile(file)
	if err != nil {
		logrus.Fatalf("could not read audio file %v", err)
	}

	client := speech.NewSpeechClient(conn)
	req := &speech.RecognizeRequest{
		Config: &speech.RecognitionConfig{
			Encoding:        speech.RecognitionConfig_FLAC,
			SampleRateHertz: 44100,
			LanguageCode:    "en-US",
		},
		Audio: &speech.RecognitionAudio{
			AudioSource: &speech.RecognitionAudio_Content{
				Content: data,
			},
		},
	}
	resp, err := client.Recognize(ctx, req)
	if err != nil {
		logrus.Fatalf("failed to recognize: %v", err)
	}
	fmt.Println("reading results:")
	for _, res := range resp.GetResults() {
		for i, alt := range res.GetAlternatives() {
			fmt.Printf("#%d: %s (%f)\n", i+1, alt.GetTranscript(), alt.GetConfidence())
		}
	}
}
