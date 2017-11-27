package page

import (
	"bytes"
	"fmt"
	"regexp"

	"bitbucket.org/djr2/tldr/color"
	"bitbucket.org/djr2/tldr/config"
)

var (
	headerRxV1  = regexp.MustCompile(`^#\s`)
	exampleRxV1 = regexp.MustCompile(`^(-\s)`)
	syntaxRxV1  = regexp.MustCompile("^`(.+)`$")
)

type pagev1 struct {
	lns [][]byte
	buf *bytes.Buffer
}

func (p *pagev1) Lines() [][]byte {
	return p.lns
}

func (p *pagev1) Write(b []byte) {
	p.buf.Write(b)
}

func (p *pagev1) Print() {
	fmt.Println(p.buf.String() + color.Reset)
}

func (p *pagev1) Header() []byte {
	cfg := config.Config
	return headerRxV1.ReplaceAll(p.lns[0], to_b(color.ColorBold(cfg.HeaderColor)))
}

func (p *pagev1) Description(line []byte) []byte {
	cfg := config.Config
	if descRx.Match(line) {
		return descRx.ReplaceAll(line, to_b(color.Color(cfg.DescriptionColor)))
	}
	return nil
}

func (p *pagev1) Example(line []byte) []byte {
	if exampleRxV1.Match(line) {
		cfg := config.Config
		return exampleRxV1.ReplaceAll(line, to_b(color.Color(cfg.HypenColor)+"$1"+color.Color(cfg.ExampleColor)))
	}
	return nil
}

func (p *pagev1) Syntax(line []byte) []byte {
	if syntaxRxV1.Match(line) {
		cfg := config.Config
		return syntaxRxV1.ReplaceAll(line, to_b(color.Color(cfg.SyntaxColor)+"$1"))
	}
	return nil
}

func (p *pagev1) Variable(line []byte) []byte {
	cfg := config.Config
	return varRx.ReplaceAll(line, to_b(color.Color(cfg.VariableColor)+"$1"+color.Color(cfg.SyntaxColor)))
}
