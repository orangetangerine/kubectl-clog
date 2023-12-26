/**
 * Author: Orange
 * Date: 2023/12/25
 */

package plugin

import (
	"bytes"
	"io"
	"os"
	"strings"

	"github.com/fatih/color"
)

type colorize struct {
	color.Color
}

func (c *colorize) WriteString(msg string) (int, error) {
	return c.Color.Fprintln(os.Stdout, msg)
}

var _ io.Writer = (*ColorizeWriter)(nil)

type ColorizeWriter struct {
	noColorWriter colorize

	errorFg colorize
	warnFg  colorize
	infoFg  colorize
	debugFg colorize

	errorBg colorize
	warnBg  colorize
	infoBg  colorize
	debugBg colorize

	contentFilters []ContentFilter
}

func defaultWriter() *ColorizeWriter {
	return &ColorizeWriter{
		errorFg: colorize{*color.New(color.FgHiRed)},
		warnFg:  colorize{*color.New(color.FgHiYellow)},
		infoFg:  colorize{*color.New(color.FgHiCyan)},
		debugFg: colorize{ /*Color: *color.New(color.FgHiWhite*/ },
		errorBg: colorize{*color.New(color.BgRed, color.Bold)},
		warnBg:  colorize{*color.New(color.BgYellow, color.Bold)},
		infoBg:  colorize{*color.New(color.BgCyan, color.Bold)},
		debugBg: colorize{*color.New(color.BgWhite, color.Bold)},
		contentFilters: []ContentFilter{
			&pureJsonLogFilter{},
			&envoyLogFilter{},
			&istioLogFilter{},
		},
	}
}

func (c *ColorizeWriter) Write(p []byte) (n int, err error) {
	return c.colorize(bytes.NewBuffer(p))
}

func (c *ColorizeWriter) Errorln(s string) (int, error) { return c.errorFg.WriteString(s) }
func (c *ColorizeWriter) Warnln(s string) (int, error)  { return c.warnFg.WriteString(s) }
func (c *ColorizeWriter) Infoln(s string) (int, error)  { return c.infoFg.WriteString(s) }
func (c *ColorizeWriter) Debugln(s string) (int, error) { return c.debugFg.WriteString(s) }
func (c *ColorizeWriter) Println(s string) (int, error) { return c.noColorWriter.WriteString(s) }

func (c *ColorizeWriter) WrapBgError(s string) string { return c.errorBg.Sprint(s) }
func (c *ColorizeWriter) WrapBgWarn(s string) string  { return c.warnBg.Sprint(s) }
func (c *ColorizeWriter) WrapBgInfo(s string) string  { return c.infoBg.Sprint(s) }
func (c *ColorizeWriter) WrapBgDebug(s string) string { return c.debugBg.Sprint(s) }

func (c *ColorizeWriter) WrapFgError(s string) string { return c.errorFg.Color.Sprint(s) }
func (c *ColorizeWriter) WrapFgWarn(s string) string  { return c.warnFg.Color.Sprint(s) }
func (c *ColorizeWriter) WrapFgInfo(s string) string  { return c.infoFg.Color.Sprint(s) }
func (c *ColorizeWriter) WrapFgDebug(s string) string { return c.debugFg.Color.Sprint(s) }

func (c *ColorizeWriter) colorize(buf *bytes.Buffer) (int, error) {
	ss := strings.Split(buf.String(), "\n")
	var total int = 0
	for _, s := range ss {
		n, err := c.filter(s)
		if err != nil {
			return total, err
		}
		total += n
	}
	return total, nil
}

func (c *ColorizeWriter) filter(s string) (int, error) {
	if len(s) == 0 {
		return 0, nil
	}

	for _, filter := range c.contentFilters {
		if n, err := filter.FilteringLine(s, c); n > 0 && err == nil {
			return n, err
		}
	}

	// default
	return c.noColorWriter.WriteString(s)
}
