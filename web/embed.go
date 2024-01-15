package web

import (
	_ "embed"
)

//go:embed vatsim-data.txt
var LegacyDataFile string

//go:embed vatsim-servers.txt
var LegacyServersFile string
