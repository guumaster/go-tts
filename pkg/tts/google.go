package tts

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/hajimehoshi/oto"
	"github.com/tosone/minimp3"
	"golang.org/x/text/language"

	"github.com/guumaster/go-tts/internal/cache"
)

// GoogleTTS struct with default options to use
type GoogleTTS struct {
	cache  *cache.FileCache
	player *oto.Player
}

// SayOptions struct with all options you can pass to Say()
type SayOptions struct {
	Slow    bool
	NoCache bool
}

// NewGoogleTTS returns a new instance with default options
func NewGoogleTTS() *GoogleTTS {
	return &GoogleTTS{
		cache: &cache.FileCache{
			BaseDir: "/tmp",
		},
	}
}

// Say will play an audio of the given text
func (t *GoogleTTS) Say(l, text string, opts *SayOptions) error {
	ltag, err := language.Parse(l)
	if err != nil {
		return err
	}
	lang := ltag.String()
	slow := false
	noCache := false
	if opts != nil {
		if opts.NoCache {
			noCache = true
		}
		if opts.Slow {
			slow = true
		}
	}

	key := t.cache.Key(fmt.Sprintf("%s_%s_%t", text, lang, slow), lang)
	path := t.cache.FullPath(fmt.Sprintf("%s.mp3", key))

	var audio io.ReadCloser
	if !noCache && t.cache.Exist(path) {
		audio, err := t.cache.Open(path)
		if err != nil {
			return err
		}
		defer audio.Close()
		return t.play(audio)
	}

	audio, err = t.googleTTSReader(text, lang, slow)
	if err != nil {
		return err
	}
	defer audio.Close()

	var audioBuf bytes.Buffer
	audioReader := io.TeeReader(audio, &audioBuf)

	err = t.play(audioReader)
	if err != nil {
		return err
	}

	if noCache {
		return nil
	}

	return t.cache.Save(path, audioBuf)
}

// MustSay is a helper that wraps a call to Say() and panics if the error is non-nil.
func (t *GoogleTTS) MustSay(lang, text string, opts *SayOptions) {
	l := language.MustParse(lang)

	err := t.Say(l.String(), text, opts)
	if err != nil {
		panic(err)
	}
}

// Save an audio MP3 translation file to a specific destination
func Save(l, text string, destination string, opts *SayOptions) error {

	if destination == "" {
		return fmt.Errorf("destination must not be empty")
	}

	if !opts.NoCache {
		if _, err := os.Stat(destination); err == nil {
			return nil
		}
	}

	ltag, err := language.Parse(l)
	if err != nil {
		return err
	}
	lang := ltag.String()
	t := NewGoogleTTS()
	t.SetCacheDir(destination)

	f, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer f.Close()

	audio, err := t.googleTTSReader(text, lang, opts.Slow)
	if err != nil {
		return err
	}
	defer audio.Close()

	b, err := io.ReadAll(audio)
	if err != nil {
		return err
	}

	_, err = f.Write(b)

	return err
}

// Close discard the underlying mp3 player
func (t *GoogleTTS) Close() error {
	if t.player != nil {
		return t.player.Close()
	}
	return nil
}

// SetCacheDir set a new path to save mp3 files. it fails if path doesn't exist
func (t *GoogleTTS) SetCacheDir(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		return err
	}
	t.cache.BaseDir = path
	return nil
}

func (t *GoogleTTS) play(audio io.Reader) error {
	b, err := ioutil.ReadAll(audio)
	if err != nil {
		return err
	}

	dec, data, _ := minimp3.DecodeFull(b)
	if t.player == nil {
		player, _ := oto.NewContext(dec.SampleRate, dec.Channels, 2, 1024)
		t.player = player.NewPlayer()
	}
	_, err = t.player.Write(data)
	return err
}

func (t *GoogleTTS) googleTTSReader(text, lang string, slow bool) (io.ReadCloser, error) {
	speed := "1"
	if slow {
		speed = "0.24"
	}

	q := url.Values{}
	q.Set("ie", "UTF-8")
	q.Set("total", "1")
	q.Set("idx", "0")
	q.Set("client", "tw-ob")
	q.Set("tl", lang)
	q.Set("ttsspeed", speed)
	q.Set("q", text)
	q.Set("textlen", strconv.Itoa(len(text)))

	u := &url.URL{
		Scheme:   "https",
		Host:     "translate.google.com",
		Path:     "translate_tts",
		RawQuery: q.Encode(),
	}

	response, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}

	return response.Body, nil
}
