package main

import (
	"os"

	"github.com/vatsimnetwork/datafeed/pkg/server"
	"github.com/vatsimnetwork/datafeed/pkg/store"
)

func main() {
	st := store.New(
		os.Getenv("DATA_SERVER_URL"),
		os.Getenv("TRANSCEIVER_URL"),
		os.Getenv("ATIS_CLIENTS_URL"),
	)

	go st.Run()
	server.Run(st)
}
