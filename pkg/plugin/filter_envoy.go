/**
 * Author: Orange
 * Date: 2023/12/26
 */

package plugin

import (
	"strings"
)

var _ ContentFilter = (*envoyLogFilter)(nil)

type envoyLogFilter struct{}

func (e *envoyLogFilter) FilteringLine(s string, logger Logger) (int, error) {
	const envoyScanMax = 50
	l := len(s)
	if l < envoyScanMax {
		return 0, nil
	}

	ss := strings.SplitN(s[:envoyScanMax], "\t", 3)
	if len(ss) < 2 {
		return 0, nil
	}

	originalLevelTag := ss[1]
	levelTag := ss[1]
	switch originalLevelTag {
	case "critical":
		levelTag = logger.WrapBgError(levelTag)
	case "error":
		levelTag = logger.WrapFgError(levelTag)
	case "warning":
		levelTag = logger.WrapBgWarn(levelTag)
	case "info":
		levelTag = logger.WrapBgInfo(levelTag)
	case "trace", "debug":
		levelTag = logger.WrapBgDebug(levelTag)
	default:
		return 0, nil
	}
	idx := strings.Index(s, originalLevelTag)
	return logger.Println(s[:idx] + levelTag + s[idx+len(originalLevelTag):])
}
