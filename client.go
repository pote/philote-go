package philote

import(
  "net/url"

  "github.com/gorilla/websocket"
)

type Client struct {
  Conn *websocket.Conn
  Read []string
  Write []string
  Server *url.URL
  token string
}

func NewClient(server, jwtSecret string, read, write []string) (*Client, error) {
  var err error
  c := &Client{}

  c.Read = read
  c.Write = write

  c.token, err =  NewToken(jwtSecret, read, write); if err != nil {
    return c, err
  }

  header := map[string][]string{
    "Authorization": []string{"Bearer " + c.token},
  }

  c.Server, err = url.Parse(server); if err != nil {
    return c, err
  }

  c.Conn, _, err = websocket.DefaultDialer.Dial(c.Server.String(), header); if err != nil {
    return c, err
  }

  return c, nil
}

func (c *Client) Publish(message *Message) (error) {
  return c.Conn.WriteJSON(message)
}

func (c *Client) Receive() (*Message, error) {
  m := &Message{}
  err := c.Conn.ReadJSON(m); if err != nil {
    return m, err
  }

  return m, nil
}

func (c *Client) NewPhilote() (chan *Message) {
  messages := make(chan *Message)

  go func () {
    m := &Message{}
    for {
      err := c.Conn.ReadJSON(m)
      if err == nil {
        messages <- m
      }
    }
  }()

  return messages
}
