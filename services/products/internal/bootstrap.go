package internal

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	"github.com/pkg/errors"
	"github.com/takumi-software/lets-work-together-my-theresa/protos/go/my-theresa/products"
	"github.com/takumi-software/lets-work-together-my-theresa/services/products/internal/adapter"
	"github.com/takumi-software/lets-work-together-my-theresa/services/products/internal/application"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

type GeneralRequest struct {
	UserId string
}

const Name = "promotions"

type Config struct {
	GrpcServer string `mapstructure:"grpc_server"`
	HTTP       string `mapstructure:"http"`
}

type Products struct {
	ErrHandler func(error, interface{})
}

func Bootstrap(ctx context.Context, cfg Config, logger *zap.Logger) error {
	be, err := New(ctx, cfg, logger)
	if err != nil {
		return errors.Wrapf(err, "%s initialization failed ", Name)
	}

	be.ErrHandler = func(err error, value interface{}) {
		logger.Error("err in "+Name, zap.Error(err), zap.Any("value", value))
	}

	g, errGroupContext := errgroup.WithContext(ctx)
	starter := func(name string, f func(context.Context) error) {
		g.Go(func() error {
			defer func() {
				if r := recover(); r != nil {
					be.ErrHandler(errors.New("panic on "+name), r)
				}
			}()
			return errors.Wrap(f(errGroupContext), "error on "+name)
		})
	}

	starter("grpc-server", func(ctx context.Context) error {
		return be.Serve(ctx, cfg, logger)
	})

	starter("http-server", func(ctx context.Context) error {
		return runHttp(ctx, cfg, logger)
	})

	return g.Wait()
}

func New(ctx context.Context, cfg Config, logger *zap.Logger) (Products, error) {
	//TODO we can add generic context, config and logger in the future
	return Products{}, nil
}

func (bo *Products) Serve(ctx context.Context, cfg Config, logger *zap.Logger) error {
	grpcServer := grpc.NewServer()
	go func() {
		<-ctx.Done()
		grpcServer.GracefulStop()
	}()

	listener, err := StartGRPCListener(&cfg)

	if err != nil {
		return errors.Wrap(err, "Unable to initialize listener")
	}

	productsService := application.NewService()
	server, err := adapter.NewGRPCServer(productsService)
	products.RegisterProductListingServer(grpcServer, &server)

	logger.Info("GRPC Server tap", zap.String("tpc_addr", cfg.GrpcServer))
	err = grpcServer.Serve(listener)
	return err
}

func StartGRPCListener(cfg *Config) (net.Listener, error) {
	l, err := net.Listen("tcp", cfg.GrpcServer)
	if err != nil {
		return nil, err
	}
	cfg.GrpcServer = l.Addr().String()
	return l, nil
}

func runHttp(ctx context.Context, cfg Config, logger *zap.Logger) error {
	logger.Info("Waiting to add gateway for http service")
	time.Sleep(5 * time.Second)
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := products.RegisterProductListingHandlerFromEndpoint(ctx, mux, cfg.GrpcServer, opts)
	if err != nil {
		return err
	}
	logger.Info("We are rerady to go!!")
	err = http.ListenAndServe(cfg.HTTP, mux)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
