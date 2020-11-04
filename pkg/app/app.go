// Package app defines common utility functions for building an application.
package app

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
	"reflect"
	"strings"
	"syscall"

	"github.com/asaskevich/govalidator"
	"github.com/octago/sflags"
	"github.com/octago/sflags/gen/gpflag"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// NewCobraApp creates new Cobra application and sets up root logger.
// It supports graceful shutdown: the returned context is canceled if SIGTERM or SIGINT are received.
// Use the returned context as the root context in the application.
// It will panic if logger creation fails.
// The caller must call .Sync() on the logger at the end of the application lifecycle.
func NewCobraApp(name string, cfg interface{}, opt ...Option) (*cobra.Command, *zap.Logger, context.Context) {
	var opts options
	for _, o := range opt {
		o(&opts)
	}

	log, err := NewZapLogger(opts.LogOpts...)
	if err != nil {
		panic(err)
	}

	cmd := bootstrap(cfg, &cobra.Command{
		Use: name,
	})

	ctx, cancel := context.WithCancel(context.Background())

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)

	cleanups := []func(){
		cancel,
	}

	go func() {
		sig := <-ch
		log.Named("app").Info("Shutdown signal received", zap.String("signal", sig.String()))
		signal.Stop(ch)
		for _, c := range cleanups {
			if c != nil {
				c()
			}
		}
	}()

	return cmd, log, ctx
}

// ErrSignalReceived indicates OS Signal was received to stop application gracefully.
//
// Deprecated: Use NewCobraApp() instead.
var ErrSignalReceived = errors.New("stop signal received")

// Bootstrap sets common things for a root cmd and returns the same cmd.
// It assumes that cfg is a pointer to a struct. Struct fields should have tag named mapstructure
// that would be used to generate flags at runtime for each field.
// Multi-word tag values should use underscore as a separator.
// In the end the command will fill the passed cfg struct with values fetched from config file,
// environment or command line flags.
//
// Deprecated: Use New() instead.
func Bootstrap(cfg interface{}, v *viper.Viper, cmd *cobra.Command) *cobra.Command {
	return bootstrap(cfg, cmd)
}

func bootstrap(cfg interface{}, cmd *cobra.Command) *cobra.Command {
	v := viper.New()
	if err := gpflag.ParseTo(cfg, cmd.PersistentFlags(), sflags.FlagDivider("."), sflags.FlagTag("mapstructure")); err != nil {
		panic(err)
	}

	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := v.BindPFlags(cmd.PersistentFlags()); err != nil {
		panic(err)
	}

	cmd.SetGlobalNormalizationFunc(func(fs *pflag.FlagSet, name string) pflag.NormalizedName {
		return pflag.NormalizedName(strings.Replace(name, "_", "-", -1))
	})

	var cfgFile string

	cmd.PersistentFlags().StringVar(&cfgFile, "config", "", "path to the config file")

	cmd.SilenceUsage = true
	cmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		if cfgFile != "" {
			v.SetConfigFile(cfgFile)
		} else {
			v.SetConfigName(cmd.Name())
			v.AddConfigPath(".")
			v.AddConfigPath("/etc/")
		}

		if err := v.ReadInConfig(); err == nil {
			fmt.Fprintln(os.Stderr, "Using config file:", v.ConfigFileUsed())
		}

		if err := v.Unmarshal(cfg); err != nil {
			return err
		}

		if _, err := govalidator.ValidateStruct(cfg); err != nil {
			return err
		}

		return nil
	}

	return cmd
}

// Close attempts to solve a common mistake of not checking the error returned from io.Closer
// when executed as a deferred statement. This function should be executed in the deferred statement
// instead of calling Close() on the io.Closer directly.
//
// Example: `defer app.Close(kafkaProducer, logger)`
func Close(c io.Closer, log *zap.Logger) {
	err := c.Close()
	log.Info("Component closed",
		zap.String("component", reflect.TypeOf(c).String()),
		zap.Error(err),
	)
}
