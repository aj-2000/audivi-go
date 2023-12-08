package utils

import (
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func Cut() {
	err := ffmpeg.Input("audio.mp3", ffmpeg.KwArgs{"ss": 1}).Output("output.mp3", ffmpeg.KwArgs{"t": 1}).OverWriteOutput().Run()
	if err != nil {
		println("Error: ", err)
	} 
}