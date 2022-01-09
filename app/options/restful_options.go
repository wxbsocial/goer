package options

import (
	"fmt"

	"github.com/spf13/pflag"
)

type RestfulOptions struct {
	BindAddress string `json:"bind-address" mapstructure:"bind-address"`
	BindPort    int    `json:"bind-port"    mapstructure:"bind-port"`
}

func NewRestfulOptions() *RestfulOptions {
	return &RestfulOptions{
		BindAddress: "127.0.0.1",
		BindPort:    8082,
	}
}

func (s *RestfulOptions) Validate() []error {
	var errors []error

	if s.BindPort < 0 || s.BindPort > 65535 {
		errors = append(
			errors,
			fmt.Errorf(
				"--bind-port %d must be between 0 and 65535, inclusive",
				s.BindPort,
			),
		)
	}

	return errors
}

func (s *RestfulOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&s.BindAddress, "restful.bind-address", s.BindAddress, ""+
		"The IP address on which to serve (set to 0.0.0.0 for all IPv4 interfaces and :: for all IPv6 interfaces).")

	fs.IntVar(&s.BindPort, "restful.bind-port", s.BindPort, ""+
		"The port on which to serve")
}
