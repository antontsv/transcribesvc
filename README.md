# Transcription service 
[![Go Report Card](https://goreportcard.com/badge/github.com/antontsv/transcribesvc)](https://goreportcard.com/report/github.com/antontsv/transcribesvc)

Transcribe audio file using Google Speech API

See [examples_test.go](examples_test.go) for a demo

To make sample file on Mac you can use these steps:
* QuickTime to record new `audio.m4a`
* `brew install ffmpeg`
* `ffmpeg -i audio.m4a -c:a flac -ac 1 audio.flac`
