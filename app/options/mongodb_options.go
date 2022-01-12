package options

import (
	"context"
	"errors"

	"github.com/spf13/pflag"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	fs.StringVar(&o.Uri, "mongo.uri", o.Uri, "The uri of mongo(format:mongodb://localhost:27017).")
}

func (o *MongoDbOptions) NewClient(ctx context.Context) (*mongo.Client, error) {

	clientOptions := options.Client().ApplyURI(o.Uri)
	clientOptions.SetMaxPoolSize(100)

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (o *MongoDbOptions) NewClientMust(ctx context.Context) *mongo.Client {
	client, err := o.NewClient(ctx)
	if err != nil {
		panic(err)
	}

	return client
}
