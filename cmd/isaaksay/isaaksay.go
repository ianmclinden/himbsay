package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ianmclinden/himbsay/common"
	"github.com/urfave/cli/v2"
)

const (
	colPadding       int    = 41 // Left hand offset
	minTemplateLines int    = 4  // Template has 4 lines of formatted text, the rest need to be above
	template         string = "" +
		"                \x1b[36mX0\x1b[34mOkxooodxkO\x1b[36mKX\x1b[0m           %v\n" +
		"             \x1b[36m0\x1b[34mxxxxxdolllllcloodk\x1b[36mK\x1b[0m        %v\n" +
		"           \x1b[34m0dddddooolcccllllllllox\x1b[36mK\x1b[0m      %v\n" +
		"          \x1b[34mkxddddooodllllllllllloodd\x1b[36m0\x1b[0m     %v\n" +
		"         \x1b[36mK\x1b[34mOkxxxddooollollooddoddddxk\x1b[36mK\x1b[0m   /\n" +
		"         \x1b[36m00\x1b[34mOkkxdddoollll:,:odddddd\x1b[36m:\x1b[34m:x\x1b[0m\n" +
		"         \x1b[36m00O\x1b[34mOkxxddoolll:\x1b[37m.. \x1b[34m;ooddxx\x1b[37m. \x1b[34mx\x1b[0m\n" +
		"         \x1b[36mK0O\x1b[34mOkkxxdooc:::\x1b[37m. .\x1b[34mcoooddxlc\x1b[36mk\x1b[0m\n" +
		"       \x1b[36m X00\x1b[34mOkkkxdolc;::::ccclc:clolo\x1b[36mOKKKKX\x1b[0m\n" +
		"   \x1b[36mKO\x1b[34mkolllloooollc::::;;,,,,;;:;,',:lcllloo\x1b[31mk\x1b[0m\n" +
		" \x1b[34mkdolccc:lclc:;,,'',,,,,,;;:;;;,'';cc::ccll\x1b[31mk\x1b[0m\n" +
		"\x1b[31mX\x1b[34mxdo\x1b[31mll\x1b[34mlddllcc::;:;,,:lclclcc:ccc,\x1b[31m'\x1b[34m;:cclcccc\x1b[31mX\x1b[0m\n" +
		"  \x1b[31mOdl:\x1b[34mxxdolc\x1b[31mc:::;;:llc\x1b[34moodollclllo\x1b[31ml:,,,;;cdK\x1b[0m\n" +
		"       \x1b[31mOdc::ccoxK     kllc:ccclk\x1b[0m\n" +
		"                        \x1b[31mkolclxX\x1b[0m\n"
)

var (
	defaultMessages = []string{
		"\x1b[34m...\x1b[0m",
		"I do not wish to go in the crack...",
		"Got any \x1b[1;36mTHOUGHTS\x1b[m to suck?",
		"I got no thoughts in my little head...",
		"... just a little boy",
	}
)

func isaaksay(c *cli.Context) error {
	var (
		message        = strings.Join(c.Args().Slice(), " ")
		defaultMessage = common.GetRandomFrom(defaultMessages)
		noWrap         = c.Bool("no-wrap")
		width          = c.Int("output-width")
		escape         = c.Bool("extended-formatting")
		list           = c.Bool("list")
	)

	if list {
		_, err := fmt.Println("isaaksay only supports one type of isasak:\ndefault")
		return err
	}

	if noWrap {
		width = 2 * len(message)
	}

	return common.Say(os.Stdout, message, defaultMessage, template, colPadding, minTemplateLines, width, escape)
}

var (
	version string
)

func main() {
	app := &cli.App{
		Name: "isaaksay",
		Authors: []*cli.Author{
			{Name: "Ian McLinden", Email: ""},
		},
		Version:              version,
		Compiled:             time.Now(),
		Usage:                "a speaking isaak",
		UsageText:            "isaaksay [-en] [-wW width] [message]",
		EnableBashCompletion: true,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "extended-formatting",
				Aliases: []string{"e"},
				Usage:   "Allow ANSI character escape sequences",
				Value:   false,
			},
			&cli.BoolFlag{
				Name:    "no-wrap",
				Aliases: []string{"n"},
				Usage:   "Don't wrap output text",
				Value:   false,
			},
			&cli.IntFlag{
				Name:    "output-width",
				Aliases: []string{"width", "W", "w"},
				Usage:   "Set the output width of the speech baloon in columns (Useful for chaining commands)",
				Value:   common.GetTermSize() - (colPadding + 4), // speech bubble padding
			},
			&cli.BoolFlag{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "(Compatabiliity) List supported isaaks",
				Hidden:  true,
				Value:   false,
			},
			&cli.StringFlag{
				Name:    "isaakfile",
				Aliases: []string{"f"},
				Usage:   "(Compatabiliity) Specify isaakfile",
				Hidden:  true,
				Value:   "",
			},
		},
		Action: isaaksay,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
