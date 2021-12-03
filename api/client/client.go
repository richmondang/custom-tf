package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/richmondang/custom-tf/api/server"
	"io"
	"net/http"
)

// Client holds all of the information required to connect to a server
type Client struct {
	hostname   string
	port       int
	authToken  string
	httpClient *http.Client
}

// NewClient returns a new client configured to communicate on a server with the
// given hostname and port and to send an Authorization Header with the value of
// token
func NewClient(hostname string, port int, token string) *Client {
	return &Client{
		hostname:   hostname,
		port:       port,
		authToken:  token,
		httpClient: &http.Client{},
	}
}

// Get all volumes info
func (c *Client) GetAll() (*map[string]server.Volume, error) {
	body, err := c.httpRequest("volume", "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	volumes := map[string]server.Volume{}
	err = json.NewDecoder(body).Decode(&volumes)
	if err != nil {
		return nil, err
	}
	return &volumes, nil
}

// Get volume by ID
func (c *Client) GetVolume(name string) (*server.Volume, error) {
	body, err := c.httpRequest(fmt.Sprintf("volume/%v", name), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	volume := &server.Volume{}
	err = json.NewDecoder(body).Decode(volume)
	if err != nil {
		return nil, err
	}
	return volume, nil
}

// Create a new volume
func (c *Client) NewVolume(volume *server.Volume) error {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(volume)
	if err != nil {
		return err
	}
	_, err = c.httpRequest("volume", "POST", buf)
	if err != nil {
		return err
	}
	return nil
}

// updates specific volume by ID
func (c *Client) UpdateVolume(volume *server.Volume) error {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(volume)
	if err != nil {
		return err
	}
	_, err = c.httpRequest(fmt.Sprintf("volume/%s", volume.ID), "PUT", buf)
	if err != nil {
		return err
	}
	return nil
}

// Delete volume by ID
func (c *Client) DeleteVolume(volumeName string) error {
	_, err := c.httpRequest(fmt.Sprintf("volume/%s", volumeName), "DELETE", bytes.Buffer{})
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) httpRequest(path, method string, body bytes.Buffer) (closer io.ReadCloser, err error) {
	req, err := http.NewRequest(method, c.requestPath(path), &body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", c.authToken)
	switch method {
	case "GET":
	case "DELETE":
	default:
		req.Header.Add("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		respBody := new(bytes.Buffer)
		_, err := respBody.ReadFrom(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("got a non 200 status code: %v", resp.StatusCode)
		}
		return nil, fmt.Errorf("got a non 200 status code: %v - %s", resp.StatusCode, respBody.String())
	}
	return resp.Body, nil
}

func (c *Client) requestPath(path string) string {
	return fmt.Sprintf("%s:%v/%s", c.hostname, c.port, path)
}
