package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	t.Run("TEST01_NO", func(t *testing.T) {
		env, err := ReadDir("mypath/path")
		require.Error(t, err)
		require.Nil(t, env)
	})

	t.Run("TEST02_YES", func(t *testing.T) {
		env, err := ReadDir("testdata/env")
		expected := Environment{
			"BAR":   {Value: "bar", NeedRemove: false},
			"EMPTY": {Value: "", NeedRemove: false},
			"FOO":   {Value: "   foo\nwith new line", NeedRemove: false},
			"HELLO": {Value: "\"hello\"", NeedRemove: false},
			"UNSET": {Value: "", NeedRemove: true},
		}
		require.NoError(t, err)
		require.Equal(t, expected, env)
	})
}
