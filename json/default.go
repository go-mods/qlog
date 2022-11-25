package json

import (
	"github.com/rs/zerolog"
	"os"
)

// Default is the default global logger
var Default = zerolog.New(os.Stderr).With().Timestamp().Logger()
