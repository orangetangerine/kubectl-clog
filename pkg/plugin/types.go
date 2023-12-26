/**
 * Author: Orange
 * Date: 2023/12/25
 */

package plugin

type ContentFilter interface {
	FilteringLine(string, Logger) (int, error)
}

// Logger Nested Color will not work properly.
// Use Errorln/Warnln/Infoln/Debugln only, or WrapXXXs + Println
type Logger interface {
	Errorln(msg string) (int, error) // return what io.Writer returns
	Warnln(msg string) (int, error)  // return what io.Writer returns
	Infoln(msg string) (int, error)  // return what io.Writer returns
	Debugln(msg string) (int, error) // return what io.Writer returns

	Println(msg string) (int, error) // return what io.Writer returns, but without color

	WrapBgError(msg string) string
	WrapBgWarn(msg string) string
	WrapBgInfo(msg string) string
	WrapBgDebug(msg string) string

	WrapFgError(msg string) string
	WrapFgWarn(msg string) string
	WrapFgInfo(msg string) string
	WrapFgDebug(msg string) string
}
