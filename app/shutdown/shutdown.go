package shutdown

import "sync"

type ShutdownCallback interface {
	OnShutdown(string) error
}

type ShutdownFunc func(string) error

func (f ShutdownFunc) OnShutdown(shutdownManager string) error {
	return f(shutdownManager)
}

type ErrorHandler interface {
	OnError(err error)
}

type ErrorFunc func(err error)

func (f ErrorFunc) OnError(err error) {
	f(err)
}

type ShutdownManager interface {
	Name() string
	Start(sd Shutdown) error
	ShutdownStart() error
	ShutdownFinish() error
}

type Shutdown interface {
	Start() error
	SetErrorHandler(handler ErrorHandler)
	AddShutdownManager(mgr ShutdownManager)
	AddShutdownCallback(callback ShutdownCallback)
	StartShutdown(mgr ShutdownManager)
}

type gracefulShutdown struct {
	callbacks    []ShutdownCallback
	managers     []ShutdownManager
	errorHandler ErrorHandler
}

func New() Shutdown {
	return &gracefulShutdown{
		callbacks: make([]ShutdownCallback, 0, 10),
		managers:  make([]ShutdownManager, 0, 3),
	}
}

func (gs *gracefulShutdown) AddShutdownManager(mgr ShutdownManager) {
	gs.managers = append(gs.managers, mgr)
}

func (gs *gracefulShutdown) AddShutdownCallback(callback ShutdownCallback) {
	gs.callbacks = append(gs.callbacks, callback)
}

func (gs *gracefulShutdown) SetErrorHandler(handler ErrorHandler) {
	gs.errorHandler = handler
}

func (gs *gracefulShutdown) Start() error {
	for _, mgr := range gs.managers {
		if err := mgr.Start(gs); err != nil {
			return err
		}
	}

	return nil
}

func (gs *gracefulShutdown) StartShutdown(mgr ShutdownManager) {
	gs.reportError(mgr.ShutdownStart())

	var wg sync.WaitGroup
	for _, callback := range gs.callbacks {
		wg.Add(1)
		go func(callback ShutdownCallback) {
			defer wg.Done()

			gs.reportError(callback.OnShutdown(mgr.Name()))
		}(callback)
	}
	wg.Wait()

	gs.reportError(mgr.ShutdownFinish())

}

func (gs *gracefulShutdown) reportError(err error) {
	if err != nil && gs.errorHandler != nil {
		gs.errorHandler.OnError(err)
	}
}
