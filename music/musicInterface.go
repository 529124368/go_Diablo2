package music

type MusicInterface interface {
	PlayMusic(name, ty string)
	CloseMusic()
}
