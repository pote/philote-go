package philote

import(
  "testing"
  "time"
)

//
// Tests require a local Philote instance running with an empty SECRET
// and on the default 6380 port.
//

func TestNewClient(t *testing.T) {
  auth := []string{"test-channel"}

  _, err := NewClient("ws://localhost:6380", "", auth, auth); if err != nil {
    t.Fatal(err)
  }
}

func TestBasicPublish(t *testing.T) {
  auth := []string{"test-channel"}

  c, err := NewClient("ws://localhost:6380", "", auth, auth); if err != nil {
    t.Fatal(err)
  }

  err = c.Publish(&Message{Channel: "test-channel", Data: "yo!"}); if err != nil {
    t.Error(err)
  }
}

func TestReceive(t *testing.T) {
  auth := []string{"test-channel"}

  c1, err := NewClient("ws://localhost:6380", "", auth, auth); if err != nil {
    t.Fatal(err)
  }

  c2, err := NewClient("ws://localhost:6380", "", auth, auth); if err != nil {
    t.Fatal(err)
  }

  go func() { time.Sleep(time.Second); c1.Publish(&Message{Channel: "test-channel", Data: "yo!"}) }()

  message, err := c2.Receive(); if err != nil {
    t.Error(err)
  }

  if message.Data != "yo!" {
    t.Error("wrong data in received message")
  }
}

func TestNewPhilote(t *testing.T) {
  auth := []string{"test-channel"}
  c1, err := NewClient("ws://localhost:6380", "", auth, auth); if err != nil {
    t.Fatal(err)
  }
  c2, err := NewClient("ws://localhost:6380", "", auth, auth); if err != nil {
    t.Fatal(err)
  }

  philote := c1.NewPhilote()
  go func() { time.Sleep(time.Second); c2.Publish(&Message{Channel: "test-channel", Data: "yo!"}) }()

  select {
  case m := <- philote:
    if m.Data != "yo!" {
      t.Error("wrong data in message")
    }
  case <- time.Tick(2 * time.Second):
    t.Error("timed out waiting for message")
  }
}
