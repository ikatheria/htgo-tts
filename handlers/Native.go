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

	// Calculate duration
	duration := float64(decodedMp3.Length()) / float64(decodedMp3.SampleRate()) / float64(decodedMp3.BitRate()/8)
	durationInDuration := time.Duration(duration) * time.Second

	// Create new audio context
	otoCtx, _, err := oto.NewContext(44100, 2, 2, 8192) // Adjust these values as per oto library's requirements
	if err != nil {
		return err
	}

	// Create a new player
	player := otoCtx.NewPlayer()

	// Start playing
	buffer := make([]byte, 8192)
	for {
		_, err := decodedMp3.Read(buffer)
		if err != nil {
			break
		}
		player.Write(buffer)
	}

	// Wait for the player to finish
	time.Sleep(durationInDuration)

	// Close the player and context
	player.Close()
	otoCtx.Close()

	return nil
}
