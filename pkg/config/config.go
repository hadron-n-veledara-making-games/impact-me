package baseconfig

type Config interface {
	ReadConfig(path string) error
}
