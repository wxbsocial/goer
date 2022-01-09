package options

import (
	"errors"

	"github.com/spf13/pflag"
)

type MongoDbOptions struct {
	Uri string `json:"uri,omitempty"    mapstructure:"uri"`
}

func NewMongoDbOptions() *MongoDbOptions {
	return &MongoDbOptions{
		Uri: "",
	}
}

func (o *MongoDbOptions) Validate() []error {
	errs := []error{}

	if o.Uri == "" {
		errs = append(errs, errors.New("the uri cannot be empty"))
	}

	return errs
}

func (o *MongoDbOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.Uri, "mongo.uri", o.Uri, "The uri of mongo(format:mongodb://mongodb0.example.com:27017).")
}
