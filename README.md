# websocketsrv-demo-go

This is a simple WebSocket server written in Golang that sends random strings to connected clients at regular intervals. The server generates secure random strings using the `crypto/rand` package and sends them over a WebSocket connection. The server stops sending messages when the client disconnects.

## Features

- **WebSocket Communication:** Handles WebSocket connections using the `gorilla/websocket` package.
- **Secure Random String Generation:** Generates secure random strings using the `crypto/rand` package.
- **Automatic Disconnection Handling:** Stops sending messages when the client disconnects.

## Installation

```bash
$ go mod tidy
```

## Usage

1. Run the server:

    ```bash
    $ go run main.go
    Server started on :8080
    ```

2. The server will start on `http://localhost:8080`. You can connect to it using a WebSocket client.

3. Once connected, the server will send a random string of 10 characters to the client every 2 seconds.

4. The server will automatically stop sending messages if the client disconnects.

    ```bash
    $ go run main.go
    Server started on :8080
    Sent: NFfbxlXkb2
    Sent: 3n3thdPGvO
    Sent: GYAEc7XclO
    Sent: xiYGu8mrCn
    Sent: 3GqIhsCXhA
    Error reading message: websocket: close 1005 (no status)
    WebSocket connection closed.
    ^C
    ```

## Notes

The server does not implement origin checks, so it will accept connections from any origin.
This is suitable for a simple demo or internal testing.

## License

This project is licensed under the MIT License - see the [LICENSE](https://opensource.org/license/mit) for details.
