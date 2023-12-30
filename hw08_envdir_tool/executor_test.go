package main

import (
	"testing"

	"github.com/stretchr/testify/require" //nolint:depguard
)

func TestRunCmd(t *testing.T) {
	expectedEnv := make(Environment)
	expectedEnv["HELLO"] = EnvValue{
		Value:      "\"hello\"",
		NeedRemove: false,
	}
	expectedEnv["BAR"] = EnvValue{
		Value:      "bar",
		NeedRemove: false,
	}
	expectedEnv["FOO"] = EnvValue{
		Value:      "   foo\nwith new line",
		NeedRemove: false,
	}
	expectedEnv["UNSET"] = EnvValue{
		Value:      "",
		NeedRemove: true,
	}
	expectedEnv["EMPTY"] = EnvValue{
		Value:      "",
		NeedRemove: false,
	}
	t.Run("run true", func(t *testing.T) {
		returnCode := RunCmd([]string{"true"}, expectedEnv)
		require.Equal(t, 0, returnCode)
	})
	t.Run("run false", func(t *testing.T) {
		returnCode := RunCmd([]string{"false"}, expectedEnv)
		require.Equal(t, 1, returnCode)
	})
}
