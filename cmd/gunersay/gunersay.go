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
	minTemplateLines int    = 5  // Template has 4 lines of formatted text, the rest need to be above
	template         string = "" +
		"                   \x1b[31mXKKKX\x1b[0m                 %v\n" +
		"              \x1b[31m0kxooololoolodOK\x1b[0m           %v\n" +
		"           \x1b[31mKxollllcccccllccllllxK\x1b[0m        %v\n" +
		"         \x1b[31mXdllccccc:cccccccccccclodX\x1b[0m      %v \n" +
		"        \x1b[31m0lllcccc:::c::::cccccccccll0\x1b[0m     %v\n" +
		"       \x1b[31mXllllc:::::::::::::ccccc:ccclK\x1b[0m   /\n" +
		"       \x1b[31mkclllc:::;:::;:'',;::::ccc..cx\x1b[0m\n" +
		"       \x1b[31mxllccc::;;;;;;:  \x1b[37m..\x1b[31m;::::cc\x1b[37m .\x1b[31m:o\x1b[0m\n" +
		"       \x1b[31m0olccc::;;;;;;;.  \x1b[31m.::::;:cc:ld\x1b[0m\n" +
		"        \x1b[31molccc::,,,,;;;;::::;:;c;ccclK\x1b[0m\n" +
		"    \x1b[36mX\x1b[31m0dc:cc:::;',,,;,;;;;,;::cccc:ccodxk0X\x1b[0m\n" +
		" \x1b[31m0o:;,,,,;:;;,,..'.''',,;;;:::;;;;:clllllllk\x1b[0m\n" +
		"\x1b[31m0:;,'.;clccc:::;,',;::::ccllllc;;;:::clllol\x1b[36md\x1b[0m\n" +
		" \x1b[36md;..\x1b[31m;c:::::;;;'..,'::::::cclll. .',;:ccclx\x1b[0m\n" +
		"     \x1b[36mK\x1b[31m:;;;;;,.;oK   \x1b[36mx:;:;;::ccoKXkoc:::cok\x1b[0m\n" +
		"       \x1b[36m0kxxx0        \x1b[36mXxc:;:cxK\x1b[0m\n"
)

var (
	defaultMessages = []string{
		"\x1b[31m...\x1b[0m",
		"I will \x1b[1;36mNOT\x1b[m go in the \x1b[1;31mCRACK!\x1b[m",
		"Give me your \x1b[1;31mTHOUGHTS\x1b[m to suck!!",
		"GRR",
		"\x1b[31mAMGERY\x1b[0m",
		"\x1b[1;37m%!*$%!*$&\x1b[m!!!",
	}
)

func gunersay(c *cli.Context) error {
	var (
		message        = strings.Join(c.Args().Slice(), " ")
		defaultMessage = common.GetRandomFrom(defaultMessages)
		noWrap         = c.Bool("no-wrap")
		termWidth      = c.Int("output-width")
		escape         = c.Bool("extended-formatting")
	)

	if noWrap {
		termWidth = len(template) + len(message) + colPadding
	}

	return common.Say(os.Stdout, message, defaultMessage, template, colPadding, minTemplateLines, termWidth, escape)
}

var (
	version string
)

func main() {
	app := &cli.App{
		Name: "gunersay",
		Authors: []*cli.Author{
			{Name: "Ian McLinden", Email: ""},
		},
		Version:              version,
		Compiled:             time.Now(),
		Usage:                "a speaking Güner",
		UsageText:            "gunersay [-en] [-w width] [message]",
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
				Aliases: []string{"width", "w"},
				Usage:   "Set the output width (Useful for chaining commands)",
				Value:   common.GetTermSize(),
			},
		},
		Action: gunersay,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}