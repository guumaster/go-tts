package cmd

import (
	"bufio"
	"os"
	"strings"

	"github.com/urfave/cli/v2"

	"github.com/guumaster/go-tts/pkg/tts"
)

var (
	// Version placeholder for the version number filled by goreleaser
	Version = ""
)

// RunCLI runs the CLI command
func RunCLI(version string) error {
	app := &cli.App{
		Name:      "go-tts",
		Usage:     "play a google tts audio from the input text",
		Version:   version,
		UsageText: "go-tts <TEXT_TO_SAY>\n   echo \"TEXT_TO_SAY\" | go-tts",
		Authors: []*cli.Author{
			{
				Name:  "guumaster",
				Email: "guuweb@gmail.com",
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "lang",
				Usage:   "language of the text",
				Value:   "en",
				Aliases: []string{"l"},
			},
			&cli.BoolFlag{
				Name:    "slow",
				Usage:   "play audio slower",
				Aliases: []string{"s"},
			},
			&cli.BoolFlag{
				Name:  "no-cache",
				Usage: "don't use file cache",
			},
		},
		Before: func(c *cli.Context) error {
			if c.NArg() == 0 && !isPiped() {
				return cli.Exit("missing text to play", 1)
			}
			return nil
		},
		Action: func(c *cli.Context) error {
			text := strings.Join(c.Args().Slice(), " ")
			slow := c.Bool("slow")
			noCache := c.Bool("no-cache")
			lang := c.String("lang")

			opts := &tts.SayOptions{
				Slow:    slow,
				NoCache: noCache,
			}

			t := tts.NewGoogleTTS()
			defer t.Close()

			// input piped to stdin
			if isPiped() {
				scanner := bufio.NewScanner(os.Stdin)
				for scanner.Scan() {
					err := t.Say(lang, scanner.Text(), opts)
					if err != nil {
						return cli.Exit(err, 1)
					}
				}

				if err := scanner.Err(); err != nil {
					return cli.Exit(err, 1)
				}
				return nil
			}

			return t.Say(lang, text, opts)
		},
	}

	if err := app.Run(os.Args); err != nil {
		return cli.Exit(err, 1)
	}
	return nil
}

func isPiped() bool {
	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
	notPipe := info.Mode()&os.ModeNamedPipe == 0
	return !notPipe
}
