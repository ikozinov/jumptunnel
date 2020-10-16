package endpoint

import (
	"fmt"
	"strconv"
	"strings"
)

type SSHEndpoint struct {
	Host string
	Port int
	User string
}

const DefaultSSHPort = 22

func ParseSSHConnect(s string) *SSHEndpoint {
	endpoint := &SSHEndpoint{
		Host: s,
	}

	if parts := strings.Split(endpoint.Host, "@"); len(parts) > 1 {
		endpoint.User = parts[0]
		endpoint.Host = parts[1]
	}

	if parts := strings.Split(endpoint.Host, ":"); len(parts) > 1 {
		endpoint.Host = parts[0]
		endpoint.Port, _ = strconv.Atoi(parts[1])
	}

	if endpoint.Port == 0 {
		endpoint.Port = DefaultSSHPort
	}

	return endpoint
}

func (endpoint *SSHEndpoint) String() string {
	return fmt.Sprintf("%s:%d", endpoint.Host, endpoint.Port)
}

type Endpoint struct {
	Host string
	Port int
}

func Parse(s string) *Endpoint {
	endpoint := &Endpoint{
		Host: s,
	}

	if parts := strings.Split(endpoint.Host, ":"); len(parts) > 1 {
		endpoint.Host = parts[0]
		endpoint.Port, _ = strconv.Atoi(parts[1])
	}

	return endpoint
}

func (endpoint *Endpoint) String() string {
	return fmt.Sprintf("%s:%d", endpoint.Host, endpoint.Port)
}
