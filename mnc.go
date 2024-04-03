package mnc

import (
	"encoding/json"
	"errors"
	"os"
	"time"

	"github.com/nats-io/nats.go"
)

func Call[R any](topic string, payload any) (*R, error) {
	var res *R
	b, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	br, err := Request(topic, b)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(br, &res)
	if err != nil {
		return nil, errors.Join(err, errors.New("Failed to unmarshal response"))
	}
	return res, nil
}

func Request(topic string, payload []byte) ([]byte, error) {
	server := nats.DefaultURL
	if path, ok := os.LookupEnv("NATS_URL"); ok {
		server = path
	}
	nc, err := nats.Connect(server)
	if err != nil {
		return nil, err
	}
	defer nc.Close()
	// err = nc.Publish(topic, payload)
	var msg = new(nats.Msg)
	msg.Subject = topic
	msg.Data = payload
	responseMsg, err := nc.RequestMsg(msg, 60*time.Second)
	if err != nil {
		return nil, err
	}
	return responseMsg.Data, nil
}
