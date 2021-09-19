package ffmpeg

import (
	"fmt"
	"log"

	"golift.io/ffmpeg"
)

func SaveToDevice(usr string, pswd string, host string, port string, path string) error {
	url := fmt.Sprintf("rtsp://%v:%v@%v:%v/%v", usr, pswd, host, port, path)
	dest := "/tmp/captured_file.mov"

	c := &ffmpeg.Config{
		FFMPEG: "/opt/homebrew/bin/ffmpeg",
		Copy:   true,
		Audio:  true,
		Time:   120,
	}

	encode := ffmpeg.Get(c)
	cmd, out, err := encode.SaveVideo(url, dest, "title")

	log.Println("Command Used:", cmd)
	log.Println("Command Output:", out)

	if err != nil {
		return err
	}

	log.Println("Saved file from", url, "to", dest)

	return nil
}
