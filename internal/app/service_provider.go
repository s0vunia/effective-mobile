package app

import (
	"context"
	"time"

	"github.com/s0vunia/effective-mobile/internal/api/songs"
	"github.com/s0vunia/effective-mobile/internal/logger"
	"github.com/s0vunia/platform_common/pkg/closer"
	"github.com/s0vunia/platform_common/pkg/db"
	"github.com/s0vunia/platform_common/pkg/db/pg"
	"go.uber.org/zap"

	"github.com/s0vunia/effective-mobile/internal/config"
	"github.com/s0vunia/effective-mobile/internal/config/env"
	"github.com/s0vunia/effective-mobile/internal/repository"
	groupRepository "github.com/s0vunia/effective-mobile/internal/repository/group"
	songRepository "github.com/s0vunia/effective-mobile/internal/repository/song"
	verseRepository "github.com/s0vunia/effective-mobile/internal/repository/verse"
	"github.com/s0vunia/effective-mobile/internal/service"
	songService "github.com/s0vunia/effective-mobile/internal/service/song"

	"github.com/avast/retry-go"
)

type serviceProvider struct {
	pgConfig     config.PGConfig
	httpConfig   config.HTTPConfig
	loggerConfig config.LoggerConfig

	dbClient db.Client

	songRepository  repository.SongRepository
	groupRepository repository.GroupRepository
	verseRepository repository.VerseRepository

	songService service.SongService

	songImpl *songs.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := env.NewPGConfig()
		if err != nil {
			logger.Fatal(
				"failed to get pg config",
				zap.Error(err),
			)
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) HTTPConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		cfg, err := env.NewHTTPConfig()
		if err != nil {
			logger.Fatal(
				"failed to get http config",
				zap.Error(err),
			)
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}

func (s *serviceProvider) LoggerConfig() config.LoggerConfig {
	if s.loggerConfig == nil {
		cfg, err := env.NewLoggerConfig()
		if err != nil {
			logger.Fatal(
				"failed to get logger config",
				zap.Error(err),
			)
		}
		s.loggerConfig = cfg
	}
	return s.loggerConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		var cl db.Client
		err := retry.Do(
			func() error {
				var err error
				cl, err = pg.New(ctx, s.PGConfig().DSN())
				return err
			},
			retry.Attempts(10),
			retry.Delay(500*time.Millisecond),
			retry.OnRetry(func(n uint, err error) {
				logger.Info("Retrying database connection", zap.Error(err))
			}),
		)
		if err != nil {
			logger.Fatal("Failed to connect to database", zap.Error(err))
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			logger.Fatal("Failed to ping database", zap.Error(err))
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}
	return s.dbClient
}

// Repositories
func (s *serviceProvider) SongRepository(ctx context.Context) repository.SongRepository {
	if s.songRepository == nil {
		s.songRepository = songRepository.NewRepository(s.DBClient(ctx))
	}
	return s.songRepository
}

func (s *serviceProvider) GroupRepository(ctx context.Context) repository.GroupRepository {
	if s.groupRepository == nil {
		s.groupRepository = groupRepository.NewRepository(s.DBClient(ctx))
	}
	return s.groupRepository
}

func (s *serviceProvider) VerseRepository(ctx context.Context) repository.VerseRepository {
	if s.verseRepository == nil {
		s.verseRepository = verseRepository.NewRepository(s.DBClient(ctx))
	}
	return s.verseRepository
}

// Services
func (s *serviceProvider) SongService(ctx context.Context) service.SongService {
	if s.songService == nil {
		s.songService = songService.NewService(
			s.SongRepository(ctx),
			s.GroupRepository(ctx),
			s.VerseRepository(ctx),
		)
	}
	return s.songService
}

// API Implementations
func (s *serviceProvider) SongImpl(ctx context.Context) *songs.Implementation {
	if s.songImpl == nil {
		s.songImpl = songs.NewImplementation(s.SongService(ctx))
	}
	return s.songImpl
}
