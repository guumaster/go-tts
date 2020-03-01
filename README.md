[![Tests](https://img.shields.io/github/workflow/status/guumaster/go-tts/Test)](https://github.com/guumaster/go-tts/actions?query=workflow%3ATest)
[![GitHub Release](https://img.shields.io/github/release/guumaster/go-tts.svg?logo=github&labelColor=262b30)](https://github.com/guumaster/go-tts/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/guumaster/go-tts)](https://goreportcard.com/report/github.com/guumaster/go-tts)
[![License](https://img.shields.io/github/license/guumaster/go-tts)](https://github.com/guumaster/go-tts/LICENSE)

# Go-TTS

This CLI is just a wrapper on Google Translate that can play TTS audio files in all languages that are supported on the web. 

Some languages are available for translation but not for TTS, so those would fail. Don't now which ones, I didn't go through the whole list.

Although it didn't fail while it was under development, it may fail randomly if you abuse it.

* [Module docs](https://pkg.go.dev/github.com/guumaster/go-tts@v1.0.0/pkg/tts?tab=doc)

## Prerequisites

This program should be cross-platform. Uses [hajimehoshi/oto](https://github.com/hajimehoshi/oto) to play sounds, so it has the same requirements: 

 * _macOS_: `AudioToolbox.framework` required (automatically linked).
 * _Linux_: libasound2-dev required. (`apt install libasound2-dev`)
 * _FreeBSD_: OpenAL required. (`pkg install openal-soft`)
 * _OpenBSD_: OpenAL required. (`pkg_add -r openal`)

I've only tested on Linux. Open [an issue](https://github.com/guumaster/go-tts/issues/new) if you found any.


## Installation

### Install binary directly

Feel free to change the path from `/usr/local/bin`, just make sure `go-tts` is available on your `$PATH` (check with `go-tts -h`).

#### Linux/MacOS

```
$ curl -sfL https://raw.githubusercontent.com/guumaster/go-tts/master/install.sh | bash -s -- -b /usr/local/bin
```

Depending on the path you choose, it may need `sudo`
```
$ curl -sfL https://raw.githubusercontent.com/guumaster/go-tts/master/install.sh | sudo bash -s -- -b /usr/local/bin
```


### Release page download

Go to the [Release page](https://github.com/guumaster/go-tts/releases) and pick one.


### With Go tools
```
go get -u github.com/guumaster/go-tts

```

## Module Usage

```
	package main

	import (
	  "fmt"
	  "github.com/guumaster/go-tts/pkg/tts"
	)

	func main() {
        t := tts.NewGoogleTTS()
        defer t.Close()

        err := t.Say("it" "questo programma è fantastico", &SayOptions{})

	  fmt.Println(f(":boom: Hello World :beer:"))
	}
    // Output:
    // will play an audio for "this program is amazing" in italian

```

## CLI Usage

You can pass arguments: 

```
$> go-tts --slow --lang "it" "questo programma è fantastico"
// Output:
// will play an audio for "this program is amazing" in italian
```

Or echo from other commands: 
```
$> echo "all done." | go-tts
// Output:
// will play an audio for "all done" in english
```


## CLI Options

```
	USAGE:
	   go-tts <TEXT_TO_SAY>
	   echo "TEXT_TO_SAY" | go-tts

	VERSION:
	   dev

	AUTHOR:
	   guumaster <guuweb@gmail.com>

	COMMANDS:
	   help, h  Shows a list of commands or help for one command

	GLOBAL OPTIONS:
	   --slow, -s     play audio slower (default: false)
	   --no-cache     don't use file cache (default: false)
	   --help, -h     show help (default: false)
	   --version, -v  print the version (default: false)
```


## Dependencies 

 * Uses [gosimple/slug](https://github.com/gosimple/slug) to calc cache key
 * Uses [hajimehoshi/oto](https://github.com/hajimehoshi/oto) to play audio files
 * Uses [tosone/minimp3](https://github.com/tosone/minimp3) to decode MP3 files
 * Uses [urfave/cli/v2](https://github.com/urfave/cli/v2) to run as CLI 


## Acknowledgements

This is loosely based on [hegedustibor/htgo-tts](https://github.com/hegedustibor/htgo-tts), but without depending on mplayer, and with a better cache implementation.


## License

[MIT license](LICENSE)
