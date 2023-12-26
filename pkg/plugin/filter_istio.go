/**
 * Author: Orange
 * Date: 2023/12/25
 */

package plugin

import (
	"strconv"
	"strings"
)

var _ ContentFilter = (*istioLogFilter)(nil)

type istioLogFilter struct{}

func (f *istioLogFilter) FilteringLine(s string, log Logger) (int, error) {
	const istioScanMax = 200

	ss := strings.Split(s[:min(istioScanMax, len(s))], " ")
	if len(ss) <= 4 {
		return 0, nil
	}

	if !strings.HasPrefix(ss[3], "HTTP") {
		return 0, nil
	}

	httpCode := ss[4]
	if _, err := strconv.ParseInt(httpCode, 10, 64); err != nil {
		return 0, err
	}

	if strings.HasPrefix(httpCode, "2") || strings.HasPrefix(httpCode, "3") {
		s = strings.Join(ss[:4], " ") + " " + log.WrapBgInfo(httpCode) + " " + strings.Join(ss[5:], " ")
		return log.Println(s)
	}
	s = strings.Join(ss[:4], " ") + " " + log.WrapBgError(httpCode) + " " + strings.Join(ss[5:], " ")
	return log.Println(s)
}
