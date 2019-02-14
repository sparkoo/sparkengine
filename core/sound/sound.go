package sound

import (
	"github.com/veandco/go-sdl2/mix"
	"io/ioutil"
	"log"
)

type Sound struct {
	data []byte
}

func LoadSound(soundFile string) *Sound {
	//Load entire WAV data from file
	data, err := ioutil.ReadFile(soundFile)
	if err != nil {
		log.Println(err)
	}

	return &Sound{data: data}
}

func (s *Sound) Play(loop int) {
	chunk, err := mix.QuickLoadWAV(s.data)
	if err != nil {
		panic(err)
	}
	if ch, err := chunk.Play(-1, 0); err != nil {
		log.Printf("unable to play sound [%v]\n", err)
	} else {
		log.Printf("playing on channel [%d]\n", ch)
	}
}
