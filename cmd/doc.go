// Package cmd for contains the code to execute is as a CLI tool to convert strings to audio.
//
/*

	NAME:
	   go-tts - play a google tts audio from the input text

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


EXAMPLES

	t.MustSay("zh-CN", "我喜欢它", &tts.SayOptions{Slow: true})
	t.MustSay("my", "ကိုယ်ကြိုက်တယ်", opts)
	t.MustSay("bn", "আমি এটা ভালোবাসি", opts)
	t.MustSay("de", "ich liebe es", opts)
	t.MustSay("ar", "احب هذا", opts)
	t.MustSay("en", "I love it", opts)
	t.MustSay("ja", "大好きです", opts)
	t.MustSay("tr", "onu seviyorum", opts)
	t.MustSay("es", "Me encanta", &tts.SayOptions{Slow: true})
	t.MustSay("it", "Lo adoro", opts)

*/
package cmd
