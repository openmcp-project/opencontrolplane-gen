package runner_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/christophrj/opencontrolplane-gen/commands"
	"github.com/christophrj/opencontrolplane-gen/runner"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRunner_Run(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		envVariables []testEnv
		fpath        string
		want         string
	}{
		{
			name: "test succesfull execution -> returns modified source file",
			envVariables: []testEnv{
				{key: "RECONCILER_NAME", value: "Example"},
				{key: "OPTIONAL_FIELDS", value: "include"},
				{key: "FIELD_NAME", value: "myName"},
				{key: "FIELD_NAMESPACE", value: "myNamespace"},
			},
			fpath: "testdata/in.go",
			want:  "testdata/out.go",
		},
		{
			name:  "test execution with all env variables missing -> returns file where only the main //go:generate directive is removed",
			fpath: "testdata/in.go",
			want:  "testdata/out_without_envs.go",
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			for _, v := range tt.envVariables {
				t.Setenv(v.key, v.value)
			}
			r := runner.Runner{
				Commands: []commands.Command{
					commands.NewReplaceCommand(),
					commands.NewIfCommand(),
				},
			}
			got := r.Run(tt.fpath)
			// compare to expected out.go file
			expected, err := os.ReadFile(tt.want)
			require.NoError(t, err)
			assert.Equal(t, bytes.NewBuffer(expected).String(), got.String())
		})
	}
}

type testEnv struct {
	key, value string
}
