package app

import (
	"context"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/s0vunia/effective-mobile/internal/logger"
	"github.com/s0vunia/effective-mobile/internal/middlewares"
	"github.com/s0vunia/effective-mobile/pkg/validator"
	"github.com/s0vunia/platform_common/pkg/closer"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/s0vunia/effective-mobile/internal/config"
)

var (
	configPath string
)

func init() {
	configPath = os.Getenv("CONFIG_PATH")
}

type App struct {
	serviceProvider *serviceProvider
	httpServer      *http.Server
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	logger.TestInit()
	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run(ctx context.Context) error {
	logger.Debug("Starting the application...")

	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	logger.Info(a.serviceProvider.LoggerConfig().FileName())

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		err := a.runHTTPServer(ctx)
		if err != nil {
			logger.Fatal(
				"failed to run HTTP server",
				zap.Error(err),
			)
		}
	}()

	wg.Wait()
	logger.Debug("Application has finished running.")
	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initHTTPServer,
		a.initLogger,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	if configPath == "" {
		return nil
	}
	err := config.Load(configPath)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	if a.serviceProvider == nil {
		a.serviceProvider = newServiceProvider()
	}
	return nil
}

func (a *App) initHTTPServer(ctx context.Context) error {
	e := echo.New()

	e.Validator = validator.NewCustomValidator()

	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
			http.MethodPatch,
		},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
			echo.HeaderAccessControlAllowHeaders,
			echo.HeaderAccessControlAllowOrigin,
		},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	e.Use(middleware.BodyLimit("2M"))

	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: 30 * time.Second,
	}))

	a.serviceProvider.SongImpl(ctx).RegisterHandlers(e, middlewares.LogMiddleware())

	a.httpServer = &http.Server{
		Addr:              a.serviceProvider.HTTPConfig().Address(),
		Handler:           e,
		ReadHeaderTimeout: a.serviceProvider.HTTPConfig().ReadHeaderTimeout(),
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       120 * time.Second,
	}

	return nil
}

func (a *App) initLogger(_ context.Context) error {
	logger.Init(a.getCore(a.getAtomicLevel()))
	return nil
}

func (a *App) runHTTPServer(ctx context.Context) error {
	logger.Info("HTTP server is running",
		zap.String("address", a.serviceProvider.HTTPConfig().Address()),
	)

	err := a.httpServer.ListenAndServe()
	closer.Add(func() error {
		return a.httpServer.Shutdown(context.Background())
	})
	if err != nil {
		return err
	}

	return nil
}

func (a *App) getCore(level zap.AtomicLevel) zapcore.Core {
	stdout := zapcore.AddSync(os.Stdout)
	isFile := a.serviceProvider.LoggerConfig().MaxSize() != 0

	var file zapcore.WriteSyncer
	if isFile {
		file = zapcore.AddSync(&lumberjack.Logger{
			Filename:   a.serviceProvider.LoggerConfig().FileName(),
			MaxSize:    a.serviceProvider.LoggerConfig().MaxSize(), // megabytes
			MaxBackups: a.serviceProvider.LoggerConfig().MaxBackups(),
			MaxAge:     a.serviceProvider.LoggerConfig().MaxAge(), // days
		})
	}
	productionCfg := zap.NewProductionEncoderConfig()
	productionCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	developmentCfg := zap.NewDevelopmentEncoderConfig()
	developmentCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder

	consoleEncoder := zapcore.NewConsoleEncoder(developmentCfg)

	var fileEncoder zapcore.Encoder
	if isFile {
		fileEncoder = zapcore.NewJSONEncoder(productionCfg)
	}

	if isFile {
		return zapcore.NewTee(
			zapcore.NewCore(consoleEncoder, stdout, level),
			zapcore.NewCore(fileEncoder, file, level),
		)
	}

	return zapcore.NewCore(consoleEncoder, stdout, level)
}

func (a *App) getAtomicLevel() zap.AtomicLevel {
	var level zapcore.Level
	if err := level.Set(a.serviceProvider.LoggerConfig().Level()); err != nil {
		logger.Fatal("failed to set log level", zap.Error(err))
	}

	return zap.NewAtomicLevelAt(level)
}
