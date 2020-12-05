package main

import (
	"bytes"
	"log"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
	"github.com/koron-go/aquestalk"
)

// playWave play wave in-synch
func playWave(b []byte) error {
	s, f, err := wav.Decode(bytes.NewReader(b))
	if err != nil {
		return err
	}
	defer s.Close()
	err = speaker.Init(f.SampleRate, f.SampleRate.N(time.Second/10))
	if err != nil {
		return err
	}
	done := make(chan struct{})
	speaker.Play(beep.Seq(s, beep.Callback(func() {
		close(done)
	})))
	defer speaker.Close()
	<-done
	return nil
}

func talk(koe string) error {
	b, err := aquestalk.Synthe(koe, 100)
	if err != nil {
		return err
	}
	return playWave(b)
}

func main() {
	err := talk("こんにちわ/ごーふぁー")
	if err != nil {
		log.Fatal(err)
	}
}
