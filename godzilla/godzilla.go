package main

import (
	"image/color"
	"os"
	"time"

	"github.com/eliukblau/pixterm/pkg/ansimage"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/sirupsen/logrus"
)

const (
	fireImageFile = "godzilla/godzilla.png"
	roarSoundFile = "godzilla/godzilla.mp3"
)

func main() {
	command := os.Args[1]
	switch command {
	case "roar":
		roar()
	case "fire":
		fire()
	}
}

func fire() {
	logrus.Info("I burn things!!!")
	img, err := ansimage.NewFromFile(fireImageFile, color.Black, ansimage.NoDithering)
	if err != nil {
		panic(err)
	}
	img.Draw()
}

func roar() {
	logrus.Info("ROAR!!!!")
	f, err := os.Open(roarSoundFile)
	if err != nil {
		panic(err)
	}
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		panic(err)
	}
	defer streamer.Close()
	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	if err != nil {
		panic(err)
	}
	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))
	<-done
}
