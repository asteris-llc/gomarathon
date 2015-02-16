package gomarathon

import (
	log "github.com/Sirupsen/logrus"
	"strings"
)

/*
 * Server Miscellaneous requests
 * https://mesosphere.github.io/marathon/docs/rest-api.html#miscellaneous
 */

//Ping the server
//Returns 'pong'
func (c *Client) Ping(host string) (string, error) {
	data, _, err := c.do("GET", "/ping", nil)

	if err != nil {
		return "error", err
	}

	body := strings.TrimSpace(string(data))
	if body != "pong" {
		log.Error("Did not recieve a 'pong', got ", body)
	}

	return body, nil
}
