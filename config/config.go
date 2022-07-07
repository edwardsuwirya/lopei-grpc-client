package config

import "os"

type LopeiGrpcConfig struct {
	Url string
}
type Config struct {
	LopeiGrpcConfig
}

func (c *Config) readConfig() {
	grpcUrl := os.Getenv("GRPC_URL")
	c.LopeiGrpcConfig = LopeiGrpcConfig{Url: grpcUrl}
}
func NewConfig() Config {
	cfg := Config{}
	cfg.readConfig()
	return cfg
}
