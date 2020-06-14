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

// Functions cuncurrently with playSong, but with a headstart, the frames are
// sent to the channel passed as parameters.
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
			// When we reach the end of the file, the other goroutine has not
			// so we wait until it collects everything from the ch and send the done
			// signal.
			time.Sleep(time.Second * 30)
			file.Close()
			return nil
		}

		if err != nil {
			fmt.Println(err)
			return err
		}

		// The first few []byte are too big and probably are metadata, so we don't
		// use them, it fixes the starting bug.
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
