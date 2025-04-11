package emitr

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func (c *Client) Consume(topic, listener string, callback func(IncomingMessage) error) error {
	offset := int64(0)

	for {
		url := fmt.Sprintf("%s/clusters/%s/consume/%s?listener=%s&offset=%d&timeout=10000", c.baseURL, c.apiKey, topic, listener, offset)

		resp, err := c.client.Get(url)
		if err != nil {
			fmt.Printf("[eventstream] erro na requisição: %v\n", err)
			time.Sleep(1 * time.Second)
			continue
		}

		if resp.StatusCode == http.StatusNoContent {
			continue
		}

		if resp.StatusCode != http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			fmt.Printf("[eventstream] erro %d: %s\n", resp.StatusCode, body)
			_ = resp.Body.Close()
			time.Sleep(2 * time.Second)
			continue
		}

		var msgs []IncomingMessage
		if err := json.NewDecoder(resp.Body).Decode(&msgs); err != nil {
			fmt.Println("[eventstream] erro ao decodificar JSON:", err)
			_ = resp.Body.Close()
			continue
		}
		_ = resp.Body.Close()

		for _, msg := range msgs {
			if err := callback(msg); err != nil {
				fmt.Printf("[eventstream] erro no callback: %v\n", err)
				continue
			}

			if err := c.CommitOffset(topic, listener, msg.Offset); err != nil {
				fmt.Printf("[eventstream] erro ao confirmar offset: %v\n", err)
			}

			offset = msg.Offset + 1
		}
	}
}
