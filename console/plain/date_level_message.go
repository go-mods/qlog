package plain

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
	"strings"
)

var DateLevelMessage = zerolog.New(
	zerolog.ConsoleWriter{
		Out:        os.Stdout,
		NoColor:    true,
		TimeFormat: "2006-01-02",
		PartsOrder: []string{
			zerolog.TimestampFieldName,
			zerolog.LevelFieldName,
			zerolog.MessageFieldName,
		},
		FormatLevel: func(i interface{}) string {
			return strings.ToUpper(fmt.Sprintf("[%s]", i))
		},
	},
).Level(zerolog.TraceLevel).With().Timestamp().Logger()
