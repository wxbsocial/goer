package app

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/marmotedu/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wxbsocial/goer/cli/flag"
)

type Options interface {
	Flags() flag.NamedFlagSets
	Validate() []error
}

type RunFunc func(basename string) error

type App struct {
	basename    string
	name        string
	description string
	options     Options
	runFunc     RunFunc
	cmd         *cobra.Command
	args        cobra.PositionalArgs
	noConfig    bool
}

type Option func(*App)

func WithOptions(opts Options) Option {
	return func(a *App) {
		a.options = opts
	}
}

func WithRunFunc(run RunFunc) Option {
	return func(a *App) {
		a.runFunc = run
	}
}

func WithDescription(desc string) Option {
	return func(a *App) {
		a.description = desc
	}
}

func WithValidArgs(args cobra.PositionalArgs) Option {
	return func(a *App) {
		a.args = args
	}
}

func WithDefaultValidArgs() Option {
	return func(a *App) {
		a.args = func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}

			return nil
		}
	}
}

func WithNoConfig() Option {
	return func(a *App) {
		a.noConfig = true
	}
}

func NewApp(name string, basename string, opts ...Option) *App {

	a := &App{
		name:     name,
		basename: basename,
	}

	for _, o := range opts {
		o(a)
	}

	a.buildCommand()

	return a
}

func (a *App) Run() {
	if err := a.cmd.Execute(); err != nil {
		fmt.Printf("%v %v\n", color.RedString("Error:"), err)
		os.Exit(1)
	}
}

func (a *App) buildCommand() {
	cmd := &cobra.Command{
		Use:   a.basename,
		Short: a.name,
		Long:  a.description,
		Args:  a.args,
	}
	cmd.SetOut(os.Stdout)
	cmd.SetErr(os.Stderr)
	cmd.Flags().SortFlags = true
	flag.InitFlags(cmd.Flags())

	if a.runFunc != nil {
		cmd.RunE = a.runCommand
	}

	cmdFlagSet := cmd.Flags()
	var namedFlagSets flag.NamedFlagSets
	if a.options != nil {
		namedFlagSets = a.options.Flags()
		for _, fs := range namedFlagSets.FlagSets {
			cmdFlagSet.AddFlagSet(fs)
		}
	}

	globalFlagSet := namedFlagSets.FlagSet("global")
	if !a.noConfig {
		addConfigFlag(a.basename, globalFlagSet)
	}
	cmdFlagSet.AddFlagSet(globalFlagSet)

	a.cmd = cmd
}

func (a *App) runCommand(cmd *cobra.Command, args []string) error {

	if !a.noConfig {
		if err := viper.BindPFlags(cmd.Flags()); err != nil {
			return err
		}

		printConfig()

		if err := viper.Unmarshal(a.options); err != nil {
			return err
		}
	}

	if a.options != nil {
		if err := a.validateOptions(); err != nil {
			return err
		}
	}

	if a.runFunc != nil {
		return a.runFunc(a.basename)
	}

	return nil
}

func (a *App) validateOptions() error {

	if errs := a.options.Validate(); len(errs) != 0 {
		return errors.NewAggregate(errs)
	}

	return nil
}
