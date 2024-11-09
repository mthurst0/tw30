package appenv

import (
	"net/url"
	"os"
	"strings"

	"go.uber.org/zap"

	"github.com/joho/godotenv"
)

var (
	devMode    bool
	devModeSet bool
)

func loadIfExists(filename string) error {
	if _, err := os.Stat(filename); err == nil {
		return godotenv.Load(filename)
	}
	return nil
}

func init() {
	// See
	// https://github.com/joho/godotenv?tab=readme-ov-file#precedence--conventions
	// for a discussion of .env precedence.
	env := os.Getenv("TW30_ENV")
	if "" == env {
		env = "development"
	}
	err := loadIfExists(".env." + env + ".local")
	if err != nil {
		panicDuringInit("error loading .env."+env+".local", zap.Error(err))
	}
	if "test" != env {
		err = loadIfExists(".env.local")
		if err != nil {
			panicDuringInit("error loading .env.local", zap.Error(err))
		}
	}
	err = loadIfExists(".env." + env)
	if err != nil {
		panicDuringInit("error loading .env."+env, zap.Error(err))
	}
	err = loadIfExists(".env")
	if err != nil {
		panicDuringInit("error loading .env", zap.Error(err))
	}
	if DevMode() {
		MustConfigureDevelopmentLogging()
	}
}

// DevMode returns true if the TW30_DEV_MODE environment variable is set to
// "true". This is used to determine whether to enable development mode features.
// See README.md for more information on development mode features.
func DevMode() bool {
	if !devModeSet {
		s := os.Getenv("TW30_DEV_MODE")
		devMode = strings.ToLower(s) == "true"
		devModeSet = true
	}
	return devMode
}

type Settings struct {
	url *url.URL
}

func (s Settings) URL() *url.URL {
	return s.url
}

func (s Settings) AppName() string {
	return "TW30"
}

func MustSettings() Settings {
	u := os.Getenv("TW30_APP_URL")
	if u == "" {
		zap.L().Fatal("TW30_APP_URL not set")
	}
	up, err := url.Parse(u)
	if err != nil {
		zap.L().Fatal("TW30_APP_URL is not a valid URL", zap.Error(err))
	}
	return Settings{url: up}
}
