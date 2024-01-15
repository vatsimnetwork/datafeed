package server

import (
	"fmt"
	"net/http"

	"github.com/vatsimnetwork/datafeed/pkg/fsdservers"
	"github.com/vatsimnetwork/datafeed/pkg/store"
	"github.com/vatsimnetwork/datafeed/web"
)

func Run(st *store.Store) {
	fmt.Println("Running server...")

	mux := http.NewServeMux()

	mux.Handle("/vatsim-data.txt", stringHandler(&web.LegacyDataFile, "text/plain"))
	mux.Handle("/vatsim-servers.txt", stringHandler(&web.LegacyServersFile, "text/plain"))
	mux.Handle("/v3/afv-atis-data.json", stringHandler(&st.ATISClientsJSON, "application/json"))
	mux.Handle("/v3/all-servers.json", jsonMarshalHandler(fsdservers.AllServers))
	mux.Handle("/v3/sweatbox-servers.json", jsonMarshalHandler(fsdservers.SweatboxServers))
	mux.Handle("/v3/transceivers-data.json", stringHandler(&st.TransceiverJSON, "application/json"))
	mux.Handle("/v3/vatsim-data.json", stringHandler(&st.DataServerJSON, "application/json"))
	mux.Handle("/v3/vatsim-servers.json", jsonMarshalHandler(fsdservers.LiveServers))

	http.ListenAndServe(":8080", mux)
}
