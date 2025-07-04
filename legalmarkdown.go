// The LegalMarkdown package contains the parser functions to automate and simplify contracting.
// You can easily transform your contracting by using the functionality of how programmers work
// into your legal practice. By integrating such programming concepts as partials, templates,
// boolean triggers, and the like (which sound fancy but are very simple to explain), your
// transactional practice will be greatly simplified.
package main

import (
	"github.com/codegangsta/cli"
	"github.com/maxxieb/legalmarkdown/lmd"
	"log"
	"os"
)

// main parses the command line inputs and routes the commands to the appropriate wrapper
// functions.
func main() {
	legalmd := cli.NewApp()

	legalmd.Name = "legalmarkdown"
	legalmd.Usage = "Automate your contracting."
	legalmd.Version = "0.9.0"
	legalmd.Author = "Eris Industries, Ltd."
	legalmd.Email = "contact@erisindustries.com"

	legalmd.Commands = []cli.Command{

		{
			Name:      "assemble",
			ShortName: "a",
			Usage:     "create yaml frontmatter",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "t, template",
					Usage: "template file to be parsed",
				},
				cli.StringFlag{
					Name:  "p, parameters",
					Usage: "parameters file to be parsed",
				},
				cli.StringFlag{
					Name:  "o, output",
					Usage: "output file to be written",
				},
			},
			Action: cliMakeYAMLFrontMatter,
		},

		{
			Name:      "parse",
			ShortName: "p",
			Usage:     "parse to markdown",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "t, template",
					Usage: "template file to be parsed",
				},
				cli.StringFlag{
					Name:  "p, parameters",
					Usage: "parameters file to be parsed",
				},
				cli.StringFlag{
					Name:  "o, output",
					Usage: "output file to be written",
				},
			},
			Action: cliLegalToMarkdown,
		},

		{
			Name:      "render",
			ShortName: "r",
			Usage:     "render from lmd to pdf or markdown to pdf",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "t, template",
					Usage: "template file to be parsed",
				},
				cli.StringFlag{
					Name:  "p, parameters",
					Usage: "parameters file to be parsed",
				},
				cli.StringFlag{
					Name:  "o, output",
					Usage: "output file to be written",
				},
			},
			Action: cliMarkdownToPDF,
		},
	}

	legalmd.Run(os.Args)
}

func cliMakeYAMLFrontMatter(c *cli.Context) {

	if c.String("template") == "" {
		log.Fatal("Please specify a template file to parse with the --template or -t flag.")
	}

	contents := c.String("template")
	parameters := c.String("parameters")
	output := c.String("output")

	lmd.MakeYAMLFrontMatter(contents, parameters, output)
}

func cliLegalToMarkdown(c *cli.Context) {

	if c.String("template") == "" {
		log.Fatal("Please specify a template file to parse with the --template or -t flag.")
	}

	if c.String("output") == "" {
		log.Fatal("Please specify an output file to write to with the --output or -o flag.")
	}

	contents := c.String("template")
	parameters := c.String("parameters")
	output := c.String("output")

	lmd.LegalToMarkdown(contents, parameters, output)
}

func cliMarkdownToPDF(c *cli.Context) {

	if c.String("template") == "" {
		log.Fatal("Please specify a template file to parse with the --template or -t flag.")
	}

	if c.String("output") == "" {
		log.Fatal("Please specify an output file to write to with the --output or -o flag.")
	}

	contents := c.String("template")
	parameters := c.String("parameters")
	output := c.String("output")

	lmd.MarkdownToPDF(contents, parameters, output)
}
