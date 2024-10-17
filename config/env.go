//go:generate stringer -type=Env -output=env_string.go
package config

import (
	"fmt"
	"os"

	"go.uber.org/zap"
)

type Env int

const (
	test Env = iota
	local
	dev
	prod
)

var GetEnv func() Env = func() Env {
	return test
}

func IsProd() bool {
	return GetEnv() == prod
}

func GetTestEnv() Env {
	return test
}

func GetAppPath() string {
	const appPathEnv = "AppPath"
	path := os.Getenv(appPathEnv)
	fmt.Println("path:!!!!", path)
	if path == "" {
		logger.Fatal("environment variable not set", zap.String("want", appPathEnv))
	}
	return path
}
