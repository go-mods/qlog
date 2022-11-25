package plain

import (
	"github.com/rs/zerolog"
	"os"
)

var Message = zerolog.New(
	zerolog.ConsoleWriter{
		Out:     os.Stdout,
		NoColor: true,
		PartsOrder: []string{
			zerolog.MessageFieldName,
		},
	},
)
