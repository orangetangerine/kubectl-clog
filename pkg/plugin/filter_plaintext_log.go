package plugin

import (
	"strings"
)

var _ ContentFilter = (*plaintextLogFilter)(nil)

// plaintextLogFilter handles logs in the format:
//
//	2026-03-08 16:05:00.919 INFO  some message here
//	YYYY-MM-DD HH:MM:SS[.mmm] LEVEL message...
type plaintextLogFilter struct{}

func (f *plaintextLogFilter) FilteringLine(s string, logger Logger) (int, error) {
	dateTime, rest, ok := parseDateTimePrefix(s)
	if !ok {
		return 0, nil
	}

	// Split off the level token (first word of rest)
	spaceIdx := strings.IndexByte(rest, ' ')
	var level, message string
	if spaceIdx == -1 {
		level = rest
		message = ""
	} else {
		level = rest[:spaceIdx]
		message = rest[spaceIdx:]
	}

	var levelTag string
	switch strings.ToUpper(level) {
	case "INFO":
		levelTag = logger.WrapFgInfo(level)
	case "DEBUG", "TRACE":
		levelTag = logger.WrapFgDebug(level)
	case "WARN", "WARNING":
		levelTag = logger.WrapFgWarn(level)
	case "ERR", "ERROR", "FATAL", "CRITICAL":
		levelTag = logger.WrapFgError(level)
	default:
		return 0, nil
	}

	return logger.Println(logger.WrapFgTimestamp(dateTime) + " " + levelTag + message)
}

// parseDateTimePrefix checks whether s begins with "YYYY-MM-DD HH:MM:SS[.ddd+]" and
// returns the timestamp prefix, the remainder of the string (after the trailing space),
// and true on success.
func parseDateTimePrefix(s string) (dateTime, rest string, ok bool) {
	// Minimum length: "2006-01-02 15:04:05 X"
	if len(s) < 21 {
		return "", "", false
	}

	// Date: YYYY-MM-DD
	if !allDigits(s[0:4]) || s[4] != '-' ||
		!allDigits(s[5:7]) || s[7] != '-' ||
		!allDigits(s[8:10]) {
		return "", "", false
	}

	// Space between date and time
	if s[10] != ' ' {
		return "", "", false
	}

	// Time: HH:MM:SS
	if !allDigits(s[11:13]) || s[13] != ':' ||
		!allDigits(s[14:16]) || s[16] != ':' ||
		!allDigits(s[17:19]) {
		return "", "", false
	}

	// Optional fractional seconds: .ddd+
	pos := 19
	if pos < len(s) && s[pos] == ',' {
		// Some formats use comma as decimal separator
		pos++
		for pos < len(s) && isASCIIDigit(s[pos]) {
			pos++
		}
	} else if pos < len(s) && s[pos] == '.' {
		pos++
		for pos < len(s) && isASCIIDigit(s[pos]) {
			pos++
		}
	}

	// Must be followed by a space and at least one more character
	if pos >= len(s) || s[pos] != ' ' {
		return "", "", false
	}

	return s[:pos], s[pos+1:], true
}

func allDigits(b string) bool {
	for i := 0; i < len(b); i++ {
		if !isASCIIDigit(b[i]) {
			return false
		}
	}
	return true
}

func isASCIIDigit(c byte) bool {
	return c >= '0' && c <= '9'
}
