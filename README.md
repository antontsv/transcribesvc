# Transcription service
Transcribe audio file using Google Speech API

`go run main.go`
```
reading results:
#1: hello Google (0.993815)
```

To make sample file on Mac you can use these steps:
* QuickTime to record new `audio.m4a`
* `brew install ffmpeg`
* `ffmpeg -i audio.m4a -c:a flac -ac 1 audio.flac`
