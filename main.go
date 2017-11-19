package transcribesvc

import (
	"context"
	"fmt"

	speech "github.com/google/go-genproto/googleapis/cloud/speech/v1"
	"google.golang.org/api/option"
	"google.golang.org/api/transport"
)

// Transcribe runs audio file agains Google Speech service
// and returns a top translation alternative
func Transcribe(ctx context.Context, data []byte,
	encoding speech.RecognitionConfig_AudioEncoding,
	sampleRate int32, langLocale string) (string, error) {
	conn, err := transport.DialGRPC(ctx,
		option.WithEndpoint("speech.googleapis.com:443"),
		option.WithScopes("https://www.googleapis.com/auth/cloud-platform"),
	)
	if err != nil {
		return "", fmt.Errorf("could not connect to gRPC server %v", err)
	}
	defer conn.Close()

	client := speech.NewSpeechClient(conn)
	req := &speech.RecognizeRequest{
		Config: &speech.RecognitionConfig{
			Encoding:        encoding,
			SampleRateHertz: sampleRate,
			LanguageCode:    langLocale,
		},
		Audio: &speech.RecognitionAudio{
			AudioSource: &speech.RecognitionAudio_Content{
				Content: data,
			},
		},
	}
	resp, err := client.Recognize(ctx, req)
	if err != nil {
		return "", fmt.Errorf("failed to recognize: %v", err)
	}
	result := ""
	var conf float32 = -1.0
	for _, res := range resp.GetResults() {
		for _, alt := range res.GetAlternatives() {
			if result == "" || alt.GetConfidence() > conf {
				conf = alt.GetConfidence()
				result = alt.GetTranscript()
			}
		}
	}
	return result, nil
}
