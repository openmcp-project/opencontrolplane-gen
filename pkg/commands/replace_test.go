package commands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_replaceCommand_Execute(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		command      replaceCommand
		loc          string
		envVariables []testEnv
		want         string
		wantCommand  replaceCommand // expected command state post execution
	}{
		{
			name: "test successful activation with line removal",
			command: replaceCommand{
				active: false,
			},
			loc:          "// opencontrolplane-gen:replace a=ENV_A",
			envVariables: []testEnv{{key: "ENV_A", value: "aValue"}},
			// line gets removed on activation
			want: "",
			// arguments are properly mapped from env variable for actual replace
			wantCommand: replaceCommand{
				active: true,
				arguments: []searchAndReplace{{
					search:  "a",
					replace: "aValue",
				}}},
		},
		{
			name: "test successful activation with multiple arguments",
			command: replaceCommand{
				active: false,
			},
			loc: "// opencontrolplane-gen:replace a=ENV_A b=ENV_B",
			envVariables: []testEnv{
				{key: "ENV_A", value: "aValue"},
				{key: "ENV_B", value: "bValue"},
			},
			// line gets removed on activation
			want: "",
			// arguments are properly mapped from env variable for actual replace
			wantCommand: replaceCommand{
				active: true,
				arguments: []searchAndReplace{
					{
						search:  "a",
						replace: "aValue",
					},
					{
						search:  "b",
						replace: "bValue",
					},
				},
			},
		},
		{
			name: "test invalid env variable reference results in directive not being removed",
			command: replaceCommand{
				active: false,
			},
			loc:          "// opencontrolplane-gen:replace a=ENV_A",
			envVariables: []testEnv{{key: "NOT_ENV_A", value: "aValue"}},
			want:         "// opencontrolplane-gen:replace a=ENV_A",
			wantCommand: replaceCommand{
				active:    false,
				arguments: nil,
			},
		},
		{
			name: "test deactivation after replace excution",
			command: replaceCommand{
				active: true,
				arguments: []searchAndReplace{{
					search:  "namespace",
					replace: "customNamespace",
				}},
			},
			loc:  "var name, namespace string",
			want: "var name, customNamespace string",
			wantCommand: replaceCommand{
				active:    false,
				arguments: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, v := range tt.envVariables {
				t.Setenv(v.key, v.value)
			}
			got := tt.command.Execute(tt.loc)
			assert.Equal(t, tt.want, got)
			// assert command state
			assert.Equal(t, tt.command, tt.wantCommand)
		})
	}
}

type testEnv struct {
	key, value string
}
