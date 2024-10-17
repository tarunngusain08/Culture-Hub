package config_test

import (
	"testing"

	"github.com/tarunngusain08/culturehub/config"
)

func TestStartup(t *testing.T) {
	t.Setenv("AppEnv", "test")
	e := config.Startup()
	if e != nil {
		t.Error(e)
	}
}

func TestReadConfig(t *testing.T) {
	t.Run("GetString", func(t *testing.T) {
		valBool := config.GetBool("TestConfig.ReadBool")
		if !valBool {
			t.Error("expected true, got:", valBool)
		}
	})

	t.Run("GetInt", func(t *testing.T) {
		valBool := config.GetBool("TestConfig.WrongBool")
		if valBool {
			t.Error("expected false, got:", valBool)
		}
	})

	t.Run("GetFloat", func(t *testing.T) {
		valStr := config.GetString("TestConfig.ReadString")
		if valStr != "string" {
			t.Error("expected 'test', got:", valStr)
		}
	})

	t.Run("GetBool", func(t *testing.T) {
		valInt := config.GetInt("TestConfig.ReadInt")
		if valInt != 1 {
			t.Error("expected 10, got:", valInt)
		}
	})
}

func TestWriteConfig(t *testing.T) {
	t.Run("WriteBool", func(t *testing.T) {
		config.Set("TestConfig.WriteBool", true)
		valBool := config.GetBool("TestConfig.WriteBool")
		if !valBool {
			t.Error("expected true, got:", valBool)
		}
	})

	t.Run("WriteString", func(t *testing.T) {
		config.Set("TestConfig.WriteString", "test")
		valStr := config.GetString("TestConfig.WriteString")
		if valStr != "test" {
			t.Error("expected 'test', got:", valStr)
		}
	})

	t.Run("WriteInt", func(t *testing.T) {
		config.Set("TestConfig.WriteInt", 10)
		valInt := config.GetInt("TestConfig.WriteInt")
		if valInt != 10 {
			t.Error("expected 10, got:", valInt)
		}
	})
}
