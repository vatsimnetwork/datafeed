package fsdservers

type FSDServer struct {
	Identifier     string `json:"ident"`
	Hostname       string `json:"hostname_or_ip"`
	Location       string `json:"location"`
	Name           string `json:"name"`
	ConnectableInt int    `json:"clients_connection_allowed"`
	Connectable    bool   `json:"client_connections_allowed"`
	Sweatbox       bool   `json:"is_sweatbox"`
}

var (
	Automatic = FSDServer{
		Identifier:     "AUTOMATIC",
		Hostname:       "fsd.connect.vatsim.net",
		Location:       "Everywhere",
		Name:           "AUTOMATIC",
		ConnectableInt: 1,
		Connectable:    true,
		Sweatbox:       false,
	}

	UKSweatbox = FSDServer{
		Identifier:     "SWEATBOX",
		Hostname:       "209.97.133.228",
		Location:       "London, UK",
		Name:           "SWEATBOX",
		ConnectableInt: 1,
		Connectable:    true,
		Sweatbox:       true,
	}

	USASweatbox = FSDServer{
		Identifier:     "SWEATBOX-2",
		Hostname:       "45.55.54.56",
		Location:       "New York, USA",
		Name:           "SWEATBOX-2",
		ConnectableInt: 1,
		Connectable:    true,
		Sweatbox:       true,
	}

	AllServers = []FSDServer{
		Automatic,
		UKSweatbox,
		USASweatbox,
	}

	LiveServers = []FSDServer{
		Automatic,
	}

	SweatboxServers = []FSDServer{
		UKSweatbox,
		USASweatbox,
	}
)
