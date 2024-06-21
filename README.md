# KrakenD Streaming Proxy (Proof of Concept)

This project is a proof of concept that demonstrates how to use the [KrakenD](https://www.krakend.io/) API Gateway to proxy a server-sent events (SSE) server. It's important to note that **KrakenD does not natively support long-lived connections, such as SSE or any other form of streaming**. However, this project showcases a workaround using a custom server plugin to handle the SSE connection and forward events to the client.

## Limitations

Due to the nature of long-lived connections and KrakenD's architecture, there are certain limitations to keep in mind:

- The server plugin hijacks the connection to handle the SSE stream, which means that additional functionality of KrakenD is not available for the hijacked connection.
- The plugin is solely responsible for managing the connection and forwarding events to the client, bypassing KrakenD's standard request/response flow.

## Features

- Demonstrates how to proxy SSE connections using KrakenD API Gateway with a custom server plugin
- Provides examples for running the SSE server, API gateway, and client
- Highlights the limitations and considerations when using KrakenD for streaming scenarios

## Prerequisites

- Go programming language installed
- KrakenD source code (for building from source)

## Getting Started

### 1. Build KrakenD from Source

To avoid dependency issues with the plugin, build KrakenD from source:

```bash
make build-krakend
```

### 2. Build the Plugin

Navigate to the plugin directory and build the plugin:

```bash
cd krakend-server-plugin-streaming && go build -buildmode=plugin -o krakend-server-plugin-streaming.so . && cd ..
```

### 3. Test the Plugin

Test the plugin with the KrakenD binary:

```bash
./krakend test-plugin  -s ./krakend-server-plugin-streaming/krakend-server-plugin-streaming.so
```

### 4. Start the SSE Server

Run the SSE server:

```bash
go run server.go
```

### 5. Start the API Gateway

Start the KrakenD API gateway:

```bash
./krakend run -dc krakend.json
```

### 6. Run the Client

Run the client to consume events from the API gateway:

```bash
go run client/client.go
```

Alternatively, you can use `curl` to consume events:

```bash
curl -N -H "Accept: text/event-stream" http://localhost:8080/events
```

### Calling the SSE Server Directly

To call the SSE server directly without the API gateway, use the following command:

```bash
curl -N -H "Accept: text/event-stream" http://localhost:9081/events-stream
```

