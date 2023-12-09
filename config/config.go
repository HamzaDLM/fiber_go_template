package config

import (
	"embed"
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Config represents the composition of yml settings.
type Config struct {
	App struct {
		Port string `default:"6969"`
	}
	Database struct {
		Driver    string `default:"sqlite3"`
		Host      string `default:"book.db"`
		Port      string
		Dbname    string
		Username  string
		Password  string
		Migration bool `default:"false"`
	}
	Redis struct {
		Enabled            bool `default:"false"`
		ConnectionPoolSize int  `yaml:"connection_pool_size" default:"10"`
		Host               string
		Port               string
	}
	Extension struct {
		MasterGenerator bool `yaml:"master_generator" default:"false"`
		CorsEnabled     bool `yaml:"cors_enabled" default:"false"`
		SecurityEnabled bool `yaml:"security_enabled" default:"false"`
	}
	// not used ftm
	Log struct {
		RequestLogFormat string `yaml:"request_log_format" default:"${remote_ip} ${account_name} ${uri} ${method} ${status} ${latency}"`
	}
	StaticContents struct {
		Enabled bool `default:"false"`
	}
	Swagger struct {
		Enabled bool `default:"false"`
		Path    string
	}
	Security struct {
		AuthPath    []string `yaml:"auth_path"`
		ExculdePath []string `yaml:"exclude_path"`
		UserPath    []string `yaml:"user_path"`
		AdminPath   []string `yaml:"admin_path"`
	}
}

const (
	// DEV represents development environment
	DEV = "dev"
	// PRD represents production environment
	PRD = "production"
	// DOC represents docker container
	DOC = "docker"
)

// Load reads the settings written to the yml file
func Load(appConfigFile embed.FS) (*Config, string) {
	var env *string
	if value := os.Getenv("WEB_APP_ENV"); value != "" {
		env = &value
	} else {
		env = flag.String("env", "dev", "To switch configurations.")
		flag.Parse()
	}

	file, err := appConfigFile.ReadFile("config/app." + *env + ".yaml")
	if err != nil {
		fmt.Printf("Failed to read app.%s.yml: %s", *env, err)
		os.Exit(2)
	}

	config := &Config{}
	if err := yaml.Unmarshal(file, config); err != nil {
		fmt.Printf("Failed to read app.%s.yml: %s", *env, err)
		os.Exit(2)
	}

	return config, *env
}
