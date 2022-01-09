package manager

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/wxbsocial/goer/app/shutdown"
)

type PosixSignalManager struct {
	signals []os.Signal
}

func NewPosixSignalManager(sig ...os.Signal) *PosixSignalManager {
	if len(sig) == 0 {
		sig = make([]os.Signal, 2)
		sig[0] = os.Interrupt
		sig[1] = syscall.SIGTERM
	}

	return &PosixSignalManager{
		signals: sig,
	}
}

func (mgr *PosixSignalManager) Name() string {
	return "PosixSignalManager"
}

func (mgr *PosixSignalManager) Start(sd shutdown.Shutdown) error {

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, mgr.signals...)

		<-c

		sd.StartShutdown(mgr)
	}()

	return nil
}

func (mgr *PosixSignalManager) ShutdownStart() error {
	return nil
}

func (mgr *PosixSignalManager) ShutdownFinish() error {
	os.Exit(0)
	return nil
}
