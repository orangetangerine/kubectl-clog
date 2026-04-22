# Adding a New Log Format Parser

`kubectl-clog` colorizes log output by running each line through a chain of
**content filters**. Adding support for a new log format means implementing
a new filter and registering it in the chain.

## How the filter chain works

Each log line is passed through `ContentFilter` implementations in order
(defined in `colorize.go`). The first filter that successfully handles a line
"wins" - it writes colourised output and returns a positive byte count. If no
filter matches, the line is printed as-is.


## Step 1 - Create the filter file

Create `pkg/plugin/filter_<name>.go`. Declare a struct and implement the
`ContentFilter` interface:

```go
package plugin

var _ ContentFilter = (*myFormatFilter)(nil)

type myFormatFilter struct{}

func (f *myFormatFilter) FilteringLine(s string, logger Logger) (int, error) {
    // Return (0, nil) if this line doesn't match your format.
    // This lets the next filter in the chain try.
    if !looksLikeMyFormat(s) {
        return 0, nil
    }

    // Use logger methods to colourise parts of the line, then
    // call logger.Println to write the final composed string.
    return logger.Println(logger.WrapFgInfo("matched: ") + s)
}
```

The `Logger` interface (defined in `types.go`) provides:

| Method | Effect |
|---|---|
| `Errorln(msg)` | Print whole line in error colour |
| `Warnln(msg)` | Print whole line in warn colour |
| `Infoln(msg)` | Print whole line in info colour |
| `Debugln(msg)` | Print whole line in debug colour |
| `Println(msg)` | Print composed string with no additional colouring |
| `WrapFgError/Warn/Info/Debug(s)` | Wrap a substring in foreground colour |
| `WrapBgError/Warn/Info/Debug(s)` | Wrap a substring in background colour |

Use `Println` with `Wrap*` calls when you want to colourise only parts of a
line (as `envoyLogFilter` does). Use `Errorln`/`Warnln`/etc. when you want the
whole line coloured (as `pureJsonLogFilter` does).

## Step 2 - Register the filter

Open `pkg/plugin/colorize.go` and add your filter to the `contentFilters`
slice in `defaultWriter()`. **Order matters**: put more specific formats
before more general ones so they get first pick.

```go
contentFilters: []ContentFilter{
    &pureJsonLogFilter{},
    &envoyLogFilter{},
    &istioLogFilter{},
    &myFormatFilter{},   // ← add your filter
},
```

## Step 3 - Test it

Run the existing tests to make sure nothing is broken:

```sh
make test
```

Add a test file `pkg/plugin/filter_<name>_test.go` with cases that:

- confirm matching lines are handled (return `n > 0`)
- confirm non-matching lines are skipped (return `n == 0, err == nil`)

Look at the existing `*_test.go` files in `pkg/plugin/` for examples of how
tests are structured.

## Tips

- Keep detection cheap. `FilteringLine` is called for every line of every log
  stream, so avoid allocations in the "not my format" fast path.
- Prefer scanning a small prefix of the line (e.g. `const scanMax = 50`)
  rather than the whole string when possible - see `envoyLogFilter` and
  `istioLogFilter` for examples.
