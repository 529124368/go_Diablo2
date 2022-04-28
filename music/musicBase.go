package music

import (
	"bytes"
	"embed"
	"game/tools"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

var cont *audio.Context
var musicRate int = 44100

type scenceMusic struct {
	cont *audio.Player
	name string
}
type MusicBGM struct {
	image                *embed.FS
	AudioMusicScencePool []scenceMusic //场景音乐池
	AudioContBGM         *audio.Player
	BgmMusicName         string
}

func NewMusicBGM(images *embed.FS) *MusicBGM {
	cont = audio.NewContext(musicRate)
	m := &MusicBGM{
		image:                images,
		AudioMusicScencePool: make([]scenceMusic, 0),
		AudioContBGM:         nil,
		BgmMusicName:         "",
	}
	return m
}

//播放背景音乐
func (m *MusicBGM) PlayBGMusic(name string, ty int) {
	if name != m.BgmMusicName {
		if m.AudioContBGM != nil {
			m.AudioContBGM.Close()
		}
		switch ty {
		case tools.MUSICMP3:
			go func() {
				ss := m.LoadMp3(name)
				m.AudioContBGM, _ = cont.NewPlayer(ss)
				m.AudioContBGM.Play()
			}()
		case tools.MUSICWAV:
			go func() {
				ss := m.LoadWav(name)
				m.AudioContBGM, _ = cont.NewPlayer(ss)
				m.AudioContBGM.Play()
			}()
		}
	} else {
		m.AudioContBGM.Rewind()
		m.AudioContBGM.Play()
	}
}

//播放场景音乐
func (m *MusicBGM) PlayMusic(name string, ty int) {
	if !m.checkIsScenceMusicAtPlay(name) {
		//判断是否加载过得音乐
		switch ty {
		case tools.MUSICMP3:
			go func() {
				j := m.getPool(name)
				if j != nil {
					j.Rewind()
					j.Play()
				} else {
					ss := m.LoadMp3(name)
					n, _ := cont.NewPlayer(ss)
					n.Play()
					m.pushPool(name, n)
				}
			}()
		case tools.MUSICWAV:
			go func() {
				j := m.getPool(name)
				if j != nil {
					j.Rewind()
					j.Play()
				} else {
					ss := m.LoadWav(name)
					n, _ := cont.NewPlayer(ss)
					n.Play()
					m.pushPool(name, n)
				}
			}()
		}
	}
}

func (m *MusicBGM) pushPool(name string, a *audio.Player) {
	var s scenceMusic
	s.name = name
	s.cont = a
	m.AudioMusicScencePool = append(m.AudioMusicScencePool, s)
}

func (m *MusicBGM) getPool(name string) *audio.Player {
	for _, v := range m.AudioMusicScencePool {
		if name == v.name {
			return v.cont
		}
	}
	return nil
}
func (m *MusicBGM) checkIsScenceMusicAtPlay(name string) bool {
	for _, v := range m.AudioMusicScencePool {
		if v.name != name && v.cont.IsPlaying() {
			v.cont.Close()
		}
		if v.name == name && v.cont.IsPlaying() {
			return true
		}
	}
	return false
}

//MP3
func (m *MusicBGM) LoadMp3(name string) *mp3.Stream {
	bgm, _ := m.image.ReadFile("resource/BGM/" + name)
	ss, _ := mp3.DecodeWithSampleRate(musicRate, bytes.NewReader(bgm))
	return ss
}

//Wav
func (m *MusicBGM) LoadWav(name string) *wav.Stream {
	bgm, _ := m.image.ReadFile("resource/BGM/" + name)
	ss, _ := wav.DecodeWithSampleRate(musicRate, bytes.NewReader(bgm))
	return ss
}

//关闭音乐
func (m *MusicBGM) CloseBGMusic() {
	m.AudioContBGM.Close()
}

//停止音乐
func (m *MusicBGM) PauseBGMusic() {
	m.AudioContBGM.Pause()
}

func (m *MusicBGM) IsPlayingBGMusic() bool {
	return m.AudioContBGM.IsPlaying()
}
