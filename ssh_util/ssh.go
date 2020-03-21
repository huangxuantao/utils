package ssh_util

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"net"
)

type Client struct {
	client *ssh.Client
}

type Config struct {
	Network  string
	Host     string
	Port     int
	Username string
	Password string
}

func (c *Client) SetConfig(config Config) error {
	var err error
	if c.client, err = ssh.Dial(
		config.Network,
		fmt.Sprintf("%s:%d", config.Host, config.Port),
		&ssh.ClientConfig{
			User: config.Username,
			Auth: []ssh.AuthMethod{
				ssh.Password(config.Password),
			},
			HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
				return nil
			},
		},
	); err != nil {
		return err
	} else {
		return nil
	}
}

func (c *Client) GetNewSession() (*ssh.Session, error) {
	return c.client.NewSession()
}
