# Philote client for the Go programming language ![Build status](https://travis-ci.org/pote/philote-go.svg)

This package provides helper method to create [Philote](https://github.com/pote/philote) authentication tokens as well as a client to connect to and interact with a Philote instance.

## Creating auth tokens

To connect to Philote you'll first have to create an authorization token describing the connection permissions and salted with the secret shared between your consuming application and the Philote server.

```go
package main

import(
  "github.com/pote/philote-go"
)

func main() {
  token, _ := philote.NewToken(
    "yourSharedSecret",
    []string{"read-channel-1", "read-channel-2", "read-write-channel"},
    []string{"write-channel", "read-write-channel"}
  )
}
```

You can safely use this token from your browser-based Philtoe client or any other application.

## Send/receive messages

The `Client` struct implements convenience methods to interact with a Philote server.

```Go
  c, _ := NewClient("ws://localhost:6380", "yourAuthToken")

  // Publish a message to a given channel
  c.Publish(
    &Message{
      Channel: "test-channel",
      Data: "You can encode any kind of payload here, it'll be received by subscribers",
    }
  )

  // Basic receiving of messages
  message, _ := c.Receive()
  message.Data //=> "Hello! I come from another Philote connection"


  // You can also create a Go channel that receives the client's messages for
  // a more idiomatic usage.

  messages := c.NewPhilote()

  m := <- messages
  m.Data //=> "this is yet another message coming from beyond (the local network)"
```


## License

Released under MIT License, check LICENSE file for details.
