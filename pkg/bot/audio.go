package bot

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	dgo "github.com/bwmarrin/discordgo"
)

// Reads the file and send the frames to the ch(channel), when done reading
// it sends a nil error to the doneCh.
func loadSong(filepath string, ch chan []byte, doneCh chan error) error {
	fmt.Println("Loading song to buffer")
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}

	defer func() {
		doneCh <- nil
	}()

	var opuslen int16

	for {
		err = binary.Read(file, binary.LittleEndian, &opuslen)
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			// As the other goroutine hasn't finished sending the song
			// we give it a little time.
			time.Sleep(time.Second * 30)
			file.Close()
			return nil
		}

		if err != nil {
			fmt.Println(err)
			return err
		}

		// The fist few []byte are huge and not part of the song, so we don't
		// send it. Without this the beggining of the song won't load.
		if opuslen > 500 {
			continue
		}

		InBuf := make([]byte, opuslen)
		err := binary.Read(file, binary.LittleEndian, &InBuf)

		if err != nil {
			fmt.Println(err)
			return err
		}

		ch <- InBuf
	}
}

// Functions cuncurrently with loadSong but with a delay to make the sound more reliable.
func playSong(s *dgo.Session, m *dgo.MessageCreate, ch chan []byte, doneCh chan error) error {
	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		return err
	}
	vc, ok := s.VoiceConnections[c.GuildID]
	if !ok {
		return errors.New("Voice connection not found.")
	}

	vc.Speaking(true)

	fmt.Println("Sending audio")

	for {
		select {
		case frame := <-ch:
			vc.OpusSend <- frame
		case <-doneCh:
			vc.Speaking(false)
			fmt.Println("Done sending")
			return nil
		}
	}
}
