package main

import (
	"log"
	"rtsp-uploader/pkg/ffmpeg"
	"rtsp-uploader/pkg/helper"
)

func main() {
	log.Output(1, "Hello")
	usr := helper.GetEnvVar("RTSP_DEVICE_USERNAME")
	pswd := helper.GetEnvVar("RTSP_DEVICE_PASSWORD")
	host := "192.168.0.161"
	port := "8080"
	path := "h264_pcm.sdp"

	err := ffmpeg.SaveToDevice(usr, pswd, host, port, path)
	if err != nil {
		log.Fatalf("[ERROR] %v", err)
	}
}
