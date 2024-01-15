# VATSIM.net Datafeed

This is the server powering the Data API (https://data.vatsim.net).
It aggregates data from our various data sources
and provides a single endpoint for clients to access.

## Usage

Build the image using the provided Dockerfile
and deploy using your method of choice.

The container exposes port 8080.

### Environment Variables

* `DATA_SERVER_URL`: The URL of the data exposed by the DataServer.
* `TRANSCEIVER_URL`: The URL of the transceiver data exposed by the Voice API.
* `ATIS_CLIENTS_URL`: The URL of the ATIS clients data exposed by the DataServer.
