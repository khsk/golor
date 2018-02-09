package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"golang.org/x/crypto/ssh/terminal"
	"io/ioutil"
	"math"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	app := cli.NewApp()
	app.Name = "golor"
	app.Usage = "Colorize inputs"
	app.Version = "1.0.0"

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "input, i, in",
			Usage: "Input from stdin",
		},
		cli.StringFlag{
			Name:  "file, f",
			Usage: "Input from file",
		},
		cli.BoolFlag{
			Name:  "bold, b",
			Usage: "Set bold",
		},
		cli.BoolFlag{
			Name:  "italic, it",
			Usage: "Set italic",
		},
		cli.BoolFlag{
			Name:  "underline, u",
			Usage: "Set underline",
		},
		cli.BoolFlag{
			Name:  "reverse, r, invert, inv",
			Usage: "Set reverse",
		},
		cli.BoolFlag{
			Name:  "cancel, c",
			Usage: "Set cancealed",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "rainbow",
			Aliases: []string{"r"},
			Usage:   "Same columns, same colors",
			Action:  rainbowAction,
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "deviation, dev, d",
					Usage: "Each line deviates from the color cycle",
				},
			},
		},
		{
			Name:    "aurora",
			Aliases: []string{"a"},
			Usage:   "Gradationed text",
			Action:  auroraAction,
			Flags: []cli.Flag{
				cli.Float64Flag{
					Name:  "freq, f",
					Usage: "Aurora frequency ",
					Value: 0.01,
				},
				cli.Float64Flag{
					Name:  "spread, s",
					Usage: "Aurora spread",
					Value: 3,
				},
				cli.Int64Flag{
					Name:  "seed",
					Usage: "Set seed",
					Value: 0,
				},
			},
		},
	}

	app.Action = Action
	app.Run(os.Args)
}

func Action(c *cli.Context) {
	text := getText(c)
	if terminal.IsTerminal(0) && c.NArg() == 0 && text == "" {
		cli.ShowAppHelpAndExit(c, 1)
	}
	fmt.Print(rainbow(text, c))
}

func rainbowAction(c *cli.Context) {
	text := getText(c)
	fmt.Print(rainbow(text, c))
}

func auroraAction(c *cli.Context) {
	text := getText(c)
	fmt.Print(aurora(text, c))
}

func rainbow(text string, c *cli.Context) string {
	const (
		red = iota + 31
		green
		yellow
		blue
		magenta
		cyan
	)
	rainbow := []int{magenta, red, yellow, green, cyan, blue}
	buf := make([]rune, 0, len(text))

	var line int
	attribute := getVisualDistinction(c)

	i := 0
	for _, char := range text {
		if char == '\n' {
			if c.Int("deviation") != 0 {
				line++
				i = line * c.Int("deviation")
			} else {
				i = 0
			}
			buf = append(buf, char)
			continue
		}
		buf = append(buf, []rune(fmt.Sprintf("\x1b[%d%sm%c\x1b[0m", rainbow[i%6], attribute, char))...)
		i++
	}

	return string(buf)
}

func aurora(text string, c *cli.Context) string {

	buf := make([]rune, 0, len(text))
	attribute := getVisualDistinction(c)
	if c.Int64("seed") == 0 {
		rand.Seed(time.Now().UnixNano())
	} else {
		rand.Seed(c.Int64("seed"))
	}
	i := rand.Intn(256)

	freq := c.Float64("freq")
	spread := c.Float64("spread")

	for _, char := range text {
		if char == '\n' {
			buf = append(buf, char)
			continue
		}

		buf = append(buf, []rune(fmt.Sprintf("\033[38;5;%d%sm%c\033[0m", rgb(float64(i)/spread, freq), attribute, char))...)
		i++
	}

	return string(buf)

}

func rgb(i float64, freq float64) int {
	red := int(6*((math.Sin(freq*i+0)*127+128)/256)) * 36
	green := int(6*((math.Sin(freq*i+2*math.Pi/3)*127+128)/256)) * 6
	blue := int(6*((math.Sin(freq*i+4*math.Pi/3)*127+128)/256)) * 1
	return 16 + red + green + blue
}

func getVisualDistinction(c *cli.Context) string {
	if c.GlobalBool("bold") {
		return ";1"
	} else if c.GlobalBool("italic") {
		return ";3"
	} else if c.GlobalBool("underline") {
		return ";4"
	} else if c.GlobalBool("reverse") {
		return ";7"
	} else if c.GlobalBool("cancel") {
		return ";9"
	} else {
		return ""
	}
}

func getText(c *cli.Context) string {
	if !terminal.IsTerminal(0) {
		b, _ := ioutil.ReadAll(os.Stdin)
		return string(b)
	}
	if c.GlobalString("file") != "" {
		b, _ := ioutil.ReadFile(c.GlobalString("file"))
		return string(b)
	}
	return strings.Join(c.Args(), "\n")
}
