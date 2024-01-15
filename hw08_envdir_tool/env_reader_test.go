package main

import (
	"testing"

	"github.com/stretchr/testify/require" //nolint:depguard
)

func TestReadDir(t *testing.T) {
	t.Run("testdata", func(t *testing.T) {
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
		actualEnv, err := ReadDir("./testdata/env")
		require.NoError(t, err)
		require.Equal(t, expectedEnv, actualEnv)
	})

	t.Run("invalid env", func(t *testing.T) {
		actualEnv, err := ReadDir("./testdata/invalid_env")
		require.NoError(t, err)
		require.Empty(t, actualEnv)
	})

	t.Run("non exist dir", func(t *testing.T) {
		actualEnv, err := ReadDir("./testdata/non_exist_env")
		require.Error(t, err)
		require.Nil(t, actualEnv)
	})
}
