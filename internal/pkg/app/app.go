package app

import (
	"context"
	"errors"
	"fmt"
	"geek-graduate/internal/pkg/transport"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type App struct {
	ctx    context.Context
	cancel func()
	srvs   []transport.Server
}

func New(httpserver transport.Server, grpcserver transport.Server) *App {
	ctx, cancel := context.WithCancel(context.Background())
	return &App{
		ctx:    ctx,
		cancel: cancel,
		srvs:   []transport.Server{httpserver, grpcserver},
	}
}

func (a *App) Run() error {
	eg, ctx := errgroup.WithContext(a.ctx)
	wg := sync.WaitGroup{}

	for _, srv := range a.srvs {
		srv := srv

		eg.Go(func() error {
			<-ctx.Done()
			ctx, cancel := context.WithTimeout(ctx, time.Second*2)
			defer cancel()
			return srv.Stop(ctx)
		})

		wg.Add(1)
		eg.Go(func() error {
			wg.Done()
			return srv.Start(a.ctx)
		})
	}
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)

	eg.Go(func() error {
		for {
			select {
			// 捕获到系统退出信号, 调用cancel停止所有子goroutine
			case <-quit:
				return a.Stop()
			case <-ctx.Done(): //子goroutine退出
				return nil
			}
		}
	})

	if err := eg.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return fmt.Errorf("errgroup.Wait:%v", err)
	}
	fmt.Println("all servers shutdown.")
	return nil
}

func (a *App) Stop() error {
	if a.cancel != nil {
		a.cancel()
	}
	return nil
}
