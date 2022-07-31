package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	t.Run("TEST01_CHECK_ENV", func(t *testing.T) {
		os.Setenv("BAR", "nebar")
		os.Setenv("UNSET", "it's magic")
		env := Environment{
			"BAR":   {Value: "bar", NeedRemove: false},
			"EMPTY": {Value: "", NeedRemove: false},
			"FOO":   {Value: "   foo\nwith new line", NeedRemove: false},
			"HELLO": {Value: "\"hello\"", NeedRemove: false},
			"UNSET": {Value: "", NeedRemove: true},
		}
		code := RunCmd([]string{"ls"}, env)
		require.Equal(t, "bar", os.Getenv("BAR"))
		require.Equal(t, 0, code)
		require.Equal(t, "   foo\nwith new line", os.Getenv("FOO"))
		require.Equal(t, "", os.Getenv("EMPTY"))
		require.Equal(t, "", os.Getenv("UNSET"))

		code = RunCmd([]string{"echo", "hello"}, Environment{})
		require.Equal(t, 0, code)
	})
}
