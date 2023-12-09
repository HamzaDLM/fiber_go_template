package container

import (
	"github.com/HamzaDLM/go_vue/config"
	"github.com/HamzaDLM/go_vue/database"
	"go.uber.org/zap"
)

// Container represents data which should be shared across the app
type Container interface {
	GetConfig() *config.Config
	GetLogger() *zap.Logger
	GetEnv() string
	GetDb() *database.Database
}

type container struct {
	config *config.Config
	logger *zap.Logger
	env    string
	db     *database.Database
}

// logger logger.Logger
func NewContainer(config *config.Config, env string, logger *zap.Logger, db *database.Database) Container {
	return &container{
		config: config,
		logger: logger,
		env:    env,
		db:     db,
	}
}

func (c *container) GetConfig() *config.Config {
	return c.config
}

func (c *container) GetLogger() *zap.Logger {
	return c.logger
}

func (c *container) GetEnv() string {
	return c.env
}

func (c *container) GetDb() *database.Database {
	return c.db
}
