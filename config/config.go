package config

import (
	"strings"

	"github.com/ikozinov/jumptunnel/endpoint"
	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh"
)

func init() {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	_ = err

	// common
	viper.SetDefault("tunnel", "")
	viper.SetDefault("listen_addr", "")
	viper.SetDefault("listen_port", 0)
	viper.SetDefault("private_key", "")
	viper.SetDefault("passphrase", "")
	viper.SetDefault("destination", "")
}

type Config struct {
	// ListenAddr is a network interface for the incoming connections
	ListenAddr string `mapstructure:"listen_addr"`
	// ListenPort is a port number for the incoming connections
	ListenPort int `mapstructure:"listen_port"`
	// Tunnel is a ssh-connect string with user, host and port like "ec2-user@jumpbox.us-east-1.mydomain.com"
	// If port is not specified, it will default to port 22
	Tunnel string `mapstructure:"tunnel"`
	// PrivateKey is PEM-encoded private key for ssh connection
	PrivateKey string `mapstructure:"private_key"`
	Passphrase string `mapstructure:"passphrase"`
	// Destination is colon delimited string with host and port for connection from jump host to remote connection
	// If port omitted it will default to ListenPort
	Destination string `mapstructure:"destination"`
}

func (c Config) Server() *endpoint.SSHEndpoint {
	return endpoint.ParseSSHConnect(c.Tunnel)

}

func (c Config) Listen() *endpoint.Endpoint {
	return &endpoint.Endpoint{
		Host: c.ListenAddr,
		Port: c.ListenPort,
	}
}

func (c Config) Remote() *endpoint.Endpoint {

	endpoint := endpoint.Parse(c.Destination)

	if endpoint.Port == 0 {
		endpoint.Port = c.ListenPort
	}

	return endpoint
}

func (c Config) ParsePrivateKey() (ssh.AuthMethod, error) {
	var (
		signer ssh.Signer
		err    error
	)

	privatekey := []byte(c.PrivateKey)

	if c.Passphrase != "" {
		signer, err = ssh.ParsePrivateKeyWithPassphrase(privatekey, []byte(c.Passphrase))
	} else {
		signer, err = ssh.ParsePrivateKey(privatekey)
	}

	if err != nil {
		return nil, err
	}

	authMethod := ssh.PublicKeys(signer)

	return authMethod, nil
}

func New() (*Config, error) {
	var conf Config
	err := viper.Unmarshal(&conf)
	return &conf, err
}
