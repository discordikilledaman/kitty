package kitty

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/acdenisSK/gol"
)

// Logger complies with the `gol.Logger` interface
type Logger struct {
	Level gol.Level
}

// Enabled f
func (l Logger) Enabled(metadata *gol.Metadata) bool {
	return metadata.Level <= l.Level
}

// Log f
func (l Logger) Log(record *gol.Record) {
	if !l.Enabled(record.Metadata) {
		return
	}
	tim := time.Now()
	years, m, days := tim.Date()
	months := int(m)
	hours, minutes, seconds := tim.Clock()
	res := modifiedTime(minutes, seconds, hours, months, days)
	timeresult := fmt.Sprintf("%d/%s/%s %s:%s:%s", years, res[0], res[1], res[2], res[3], res[4])
	funcname := record.Location.FuncName
	for _, name := range []string{"func", "1"} {
		if strings.HasPrefix(funcname, name) {
			funcname = "closure"
		}
	}
	fmt.Fprintf(os.Stderr, "[%s] %s %s() %s:%d %s\n", strings.ToUpper(record.Level.String()), timeresult, funcname, record.Location.File, record.Location.Line, record.Args)
}

func modifiedTime(date ...int) []string {
	result := []string{}
	for _, da := range date {
		d := strconv.Itoa(da)
		if len(d) == 1 {
			d = "0" + d
		}
		result = append(result, d)
	}
	return result
}
