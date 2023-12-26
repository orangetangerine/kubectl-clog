/**
 * Author: Orange
 * Date: 2023/12/25
 */

package plugin

import (
	"strings"

	"github.com/buger/jsonparser"
)

type pureJsonLogFilter struct{}

func (f *pureJsonLogFilter) FilteringLine(s string, log Logger) (int, error) {
	const jsonScanMax = 20

	l := len(s)

	if s[0] == '{' && s[l-1] == '}' {
		return f.pickWriterViaJson(s, log)
	}
	jsonStartCharIdx := strings.Index(s[:min(jsonScanMax, l)], "{")
	if jsonStartCharIdx != -1 && s[l-1] == '}' {
		return f.pickWriterViaJson(s[jsonStartCharIdx:], log)
	}

	/*
		// something + json
		for _, sub := range strings.Split(s[:min(jsonScanMax, l)], " ") {
			switch sub {
			case "ERROR":
				return c.errorFg.WriteString(s)
			case "WARN":
				return c.warnFg.WriteString(s)
			case "INFO":
				return c.infoFg.WriteString(s)
			case "DEBUG":
				return c.debugFg.WriteString(s)
			default:
				continue
			}
		}
	*/

	return 0, nil
}

func (f *pureJsonLogFilter) pickWriterViaJson(s string, log Logger) (int, error) {
	level, _ := jsonparser.GetString([]byte(s), "level")
	switch strings.ToLower(level) {
	case "error":
		return log.Errorln(s)
	case "warn":
		return log.Warnln(s)
	case "info":
		return log.Infoln(s)
	case "", "debug":
		return log.Debugln(s)
	default:
		return log.Warnln(s)
	}
}
