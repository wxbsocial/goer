package options

import (
	"fmt"

	"github.com/spf13/pflag"
)

type GrpcOptions struct {
	BindAddress string `json:"bind-address" mapstructure:"bind-address"`
	BindPort    int    `json:"bind-port"    mapstructure:"bind-port"`
	MaxMsgSize  int    `json:"max-msg-size" mapstructure:"max-msg-size"`
}

func NewGrpcOptions() *GrpcOptions {
	return &GrpcOptions{
		BindAddress: "127.0.0.1",
		BindPort:    8080,
		MaxMsgSize:  4 * 1024 * 1024,
	}
}

func (s *GrpcOptions) Validate() []error {
	var errors []error

	if s.BindPort < 0 || s.BindPort > 65535 {
		errors = append(
			errors,
			fmt.Errorf(
				"--bind-port %v must be between 0 and 65535, inclusive.",
				s.BindPort,
			),
		)
	}

	return errors
}

func (s *GrpcOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&s.BindAddress, "grpc.bind-address", s.BindAddress, ""+
		"The IP address on which to serve (set to 0.0.0.0 for all IPv4 interfaces and :: for all IPv6 interfaces).")

	fs.IntVar(&s.BindPort, "grpc.bind-port", s.BindPort, ""+
		"The port on which to serve")

	fs.IntVar(&s.MaxMsgSize, "grpc.max-msg-size", s.MaxMsgSize, "gRPC max message size.")
}
