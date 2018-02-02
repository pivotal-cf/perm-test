package cmd

import (
	"os"

	"code.cloudfoundry.org/lager"
)

type LoadDataConfig struct {
	LogLevel              string                `yaml:"log_level"`
	CloudControllerConfig CloudControllerConfig `yaml:"cloud_controller"`
	TestDataConfig        TestDataConfig        `yaml:"test_data"`
}

type CloudControllerConfig struct {
	URL          string `yaml:"url"`
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
	CACert       string `yaml:"ca_cert"`
}

type TestDataConfig struct {
	AppsPerSpaceCount int `yaml:"apps_per_space_count"`
	SpacesPerOrgCount int `yaml:"spaces_per_org_count"`

	TestEnvironmentConfig     TestEnvironmentConfig     `yaml:"test_environment"`
	ExternalEnvironmentConfig ExternalEnvironmentConfig `yaml:"external_environment"`
}

type TestEnvironmentConfig struct {
	UserGUID string `yaml:"user_guid"`
	OrgCount int    `yaml:"org_count"`
}

type ExternalEnvironmentConfig struct {
	OrgCount  int `yaml:"org_count"`
	UserCount int `yaml:"user_count"`
}

func (c *LoadDataConfig) NewLogger(component string) lager.Logger {
	var l lager.LogLevel

	switch c.LogLevel {
	case "debug":
		l = lager.DEBUG
	case "info":
		l = lager.INFO
	case "error":
		l = lager.ERROR
	case "fatal":
		l = lager.FATAL
	default:
		l = lager.INFO
	}

	sink := lager.NewWriterSink(os.Stdout, l)
	logger := lager.NewLogger(component)
	logger.RegisterSink(sink)

	return logger
}
