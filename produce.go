package emitr

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) Produce(topic string, message string) error {
	url := fmt.Sprintf("%s/clusters/%s/topics/%s/messages", c.baseURL, c.apiKey, topic)

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(message)))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/octet-stream")

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 && resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("erro ao produzir mensagem: %s", body)
	}
	return nil
}
