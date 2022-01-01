package goer

import "time"

type TimeProvider interface {
	Now() time.Time
}

type DefaultTimeProvider struct {
}

func (provider *DefaultTimeProvider) Now() time.Time {
	return time.Now()
}

var timeProvider TimeProvider = &DefaultTimeProvider{}

func Now() time.Time {
	return timeProvider.Now()
}

func SetTimeProvider(newTimeProvider TimeProvider) {
	timeProvider = newTimeProvider
}
