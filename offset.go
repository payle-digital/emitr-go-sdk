package emitr

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type ackRequest struct {
	Listener string `json:"listener"`
	Offset   int64  `json:"offset"`
}

func (c *Client) CommitOffset(topic, listener string, offset int64) error {
	url := fmt.Sprintf("%s/clusters/%s/ack/%s", c.baseURL, c.apiKey, topic)

	body, _ := json.Marshal(ackRequest{
		Listener: listener,
		Offset:   offset,
	})

	resp, err := c.client.Post(url, "application/json", bytes.NewReader(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("erro ao confirmar offset, status %d", resp.StatusCode)
	}
	return nil
}
