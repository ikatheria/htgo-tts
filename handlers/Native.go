package handlers

import (
	"bytes"
	"os"
	"time"

	"github.com/ikatheria/go-mp3"
	"github.com/ikatheria/oto/v2"
)

type Native struct {
}

func (n *Native) Play(fileName string) error {
	// Read the mp3 file into memory
	fileBytes, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	fileBytesReader := bytes.NewReader(fileBytes)

	// Decode file
	decodedMp3, err := mp3.NewDecoder(fileBytesReader)
	if err != nil {
		return err
	}

	// Set audio parameters
	options := oto.NewContextOptions()
	options.SampleRate = decodedMp3.SampleRate()
	options.NumChannels = 2
	options.BitDepth = 2

	otoCtx, readyChan, err := oto.NewContext(options)
	if err != nil {
		return err
	}
	<-readyChan

	player := otoCtx.NewPlayer(decodedMp3)

	player.Play()

	for player.IsPlaying() {
		time.Sleep(time.Millisecond)
	}

	return player.Close()
}
