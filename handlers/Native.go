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

	// Decode MP3 file
	decodedMp3, err := mp3.NewDecoder(fileBytesReader)
	if err != nil {
		return err
	}

	// Create new audio context
	otoCtx, err := oto.NewContext(oto.WithSR(decodedMp3.SampleRate()), oto.WithChannels(2), oto.WithBitDepth(16))
	if err != nil {
		return err
	}

	// Create a new player
	player := otoCtx.NewPlayer()

	// Start playing
	player.Write(decodedMp3)

	// Wait for the player to finish
	time.Sleep(decodedMp3.Duration())

	// Close the player and context
	player.Close()
	otoCtx.Close()

	return nil
}
