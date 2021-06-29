package store

type Store struct {
	config *StoreConfig
}

func New(config *StoreConfig) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {
	return nil
}

func (s *Store) Close() {
}
