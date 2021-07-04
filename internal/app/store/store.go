package store

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type Store struct {
	config         *StoreConfig
	db             *sql.DB
	logger         *logrus.Logger
	userRepository *UserRepository
}

func New(config *StoreConfig) *Store {
	return &Store{
		config: config,
		logger: logrus.New(),
	}
}

func (s *Store) Open() error {
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.logger.Info(fmt.Sprintf("attempting connection to %s", s.config.DBParams))
	db, err := sql.Open("postgres", s.config.DBParams)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	s.db = db
	s.logger.Info("database connection established")
	return nil
}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)

	s.logger.Formatter = &logrus.TextFormatter{
		ForceColors: true,
	}

	return nil
}

func (s *Store) User() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}
