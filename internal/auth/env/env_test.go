package env

import (
    "os"
    "testing"

    "github.com/stretchr/testify/require"
)

func TestGetAPIKey(t *testing.T) {
	cases := []struct {
		name      string
        // value to set COINBASE_API_KEY to
        input    string
        // expected return string
		expected  string
        // true if an error is expected
		expectErr bool
        // expected error string
        expectErrStr string
	}{
		{
			"set",
			"abc",
            "abc",
			false,
            "",
		},
        {
            "empty",
            "",
            "",
            true,
            "Expected environment variable COINBASE_API_KEY to be set but it was empty",
        },
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
            os.Setenv(envAPIKey, tc.input)
            e := Env{}
            out, err := e.GetAPIKey()
            if tc.expectErr {
                require.Error(t, err)
                require.Equal(t, tc.expectErrStr, err.Error())
                return
            }
			require.NoError(t, err)
			require.Equal(t, tc.expected, out)
		})
	}
}

func TestGetSecret(t *testing.T) {
	cases := []struct {
		name      string
        // value to set COINBASE_API_KEY to
        input    string
        // expected return string
		expected  string
        // true if an error is expected
		expectErr bool
        // expected error string
        expectErrStr string
	}{
		{
			"set",
			"abc",
            "abc",
			false,
            "",
		},
        {
            "empty",
            "",
            "",
            true,
            "Expected environment variable COINBASE_SECRET to be set but it was empty",
        },
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
            os.Setenv(envSecret, tc.input)
            e := Env{}
            out, err := e.GetSecret()
            if tc.expectErr {
                require.Error(t, err)
                require.Equal(t, tc.expectErrStr, err.Error())
                return
            }
			require.NoError(t, err)
			require.Equal(t, tc.expected, out)
		})
	}
}
