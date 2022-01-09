package options

import (
	"errors"

	"github.com/EventStore/EventStore-Client-Go/esdb"
	"github.com/spf13/pflag"
)

type ESDbOptions struct {
	Uri string `json:"uri,omitempty"    mapstructure:"uri"`
}

func NewESDbOptions() *ESDbOptions {
	return &ESDbOptions{
		Uri: "esdb://localhost:2113?tls=false",
	}
}

func (o *ESDbOptions) Validate() []error {
	errs := []error{}

	if o.Uri == "" {
		errs = append(errs, errors.New("the uri o cannot be empty"))
	}

	return errs
}

func (o *ESDbOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.Uri, "esdb.uri", o.Uri, "The uri of esdb(format:esdb://localhost:2113?tls=false).")
}

func (o *ESDbOptions) NewClient() (*esdb.Client, error) {
	settings, err := esdb.ParseConnectionString(o.Uri)
	if err != nil {
		return nil, err
	}

	client, err := esdb.NewClient(settings)
	if err != nil {
		return nil, err
	}

	return client, nil
}
