package options

import (
	"errors"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/spf13/pflag"
)

type RedisOptions struct {
	Uri string `json:"uri,omitempty"    mapstructure:"uri"`

	// Maximum number of idle connections in the pool.
	MaxIdle int `json:"max-idle,omitempty"    mapstructure:"max-idle"`

	// Maximum number of connections allocated by the pool at a given time.
	// When zero, there is no limit on the number of connections in the pool.
	MaxActive int `json:"max-active,omitempty"    mapstructure:"max-active"`

	// Close connections after remaining idle for this duration. If the value
	// is zero, then idle connections are not closed. Applications should set
	// the timeout to a value less than the server's timeout.
	IdleTimeout time.Duration `json:"idle-timeout,omitempty"    mapstructure:"idle-timeout"`

	// If Wait is true and the pool is at the MaxActive limit, then Get() waits
	// for a connection to be returned to the pool before returning.
	Wait bool `json:"wait,omitempty"    mapstructure:"wait"`

	// Close connections older than this duration. If the value is zero, then
	// the pool does not close connections based on age.
	MaxConnLifetime time.Duration `json:"max-conn-lifetime,omitempty"    mapstructure:"max-conn-lifetime"`
}

func NewRedisOptions() *RedisOptions {
	return &RedisOptions{
		Uri:             "",
		MaxIdle:         2000,
		MaxActive:       4000,
		IdleTimeout:     0,
		Wait:            false,
		MaxConnLifetime: 0,
	}
}

func (o *RedisOptions) Validate() []error {
	errs := []error{}

	if o.Uri == "" {
		errs = append(errs, errors.New("the uri cannot be empty"))
	}

	return errs
}

func (o *RedisOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.Uri, "redis.uri", o.Uri, ""+
		"The uri of redis(format:redis://password@localhost:6379).")
	fs.IntVar(&o.MaxIdle, "redis.max-idle", o.MaxIdle, "Maximum number of idle connections in the pool.")
	fs.IntVar(&o.MaxActive, "redis.max-active", o.MaxActive, "Maximum number of connections allocated by the pool at a given time. When zero, there is no limit on the number of connections in the pool.")
	fs.DurationVar(&o.IdleTimeout, "redis.idle-timeout", o.IdleTimeout, "Close connections after remaining idle for this duration. If the value is zero, then idle connections are not closed. Applications should set the timeout to a value less than the server's timeout.")
	fs.DurationVar(&o.MaxConnLifetime, "redis.max-conn-lifetime", o.MaxConnLifetime, ""+
		"Close connections older than this duration. If the value is zero, then the pool does not close connections based on age.")
	fs.BoolVar(&o.Wait, "redis.wait", o.Wait, ""+
		"If Wait is true and the pool is at the MaxActive limit, then Get() waits for a connection to be returned to the pool before returning.")

}

func (o *RedisOptions) NewPool() (*redis.Pool, error) {
	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.DialURL(o.Uri)
		},
		MaxIdle:         o.MaxIdle,
		MaxActive:       o.MaxActive,
		IdleTimeout:     o.IdleTimeout,
		MaxConnLifetime: o.MaxConnLifetime,
		Wait:            o.Wait,
	}

	return pool, nil
}
