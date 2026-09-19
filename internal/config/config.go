package config

import "github.com/kelseyhightower/envconfig"
// Config represents the application configuration loaded from environment variables.
type Config struct {
	AppPort		string	`default:"8082" envconfig:"APP_PORT"`
	ServiceName	string	`default:"bookmark_service" envconfig:"SERVICE_NAME"`
	InstanceID	string	`default:"" envconfig:"INSTANCE_ID"`
}
// NewConfig creates a new Config instance populated with environment variables
// prefixed with "api". It returns an error if the environment variable processing fails.
func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := envconfig.Process("api", cfg)
	if err != nil {
		return nil, err
	}
	return cfg, err
}
