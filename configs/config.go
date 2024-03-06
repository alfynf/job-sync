package configs

type Config interface {
	Load() any
}
