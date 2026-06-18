package commands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ifCommand_Execute(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		command      ifCommand
		loc          string
		envVariables []testEnv
		want         string
		wantCommand  ifCommand // expected command state post execution
	}{
		{
			name: "test successful activation with line removal",
			command: ifCommand{
				active: false,
			},
			loc:          "// opencontrolplane-gen:if ENV_A=a",
			envVariables: []testEnv{{key: "ENV_A", value: "a"}},
			// line gets removed on activation
			want: "",
			// arguments are properly mapped from env variable for actual replace
			wantCommand: ifCommand{
				active:      true,
				includeLine: true,
			},
		},
		{
			name: "test invalid env variable reference results in directive not being removed",
			command: ifCommand{
				active: false,
			},
			loc:          "// opencontrolplane-gen:if ENV_A=a",
			envVariables: []testEnv{{key: "NOT_ENV_A", value: "a"}},
			want:         "// opencontrolplane-gen:if ENV_A=a",
			wantCommand: ifCommand{
				active:      false,
				includeLine: false,
			},
		},
		{
			name: "test include",
			command: ifCommand{
				active:      true,
				includeLine: true,
			},
			loc:  "var name, namespace string",
			want: "var name, namespace string",
			wantCommand: ifCommand{
				active:      true,
				includeLine: true,
			},
		},
		{
			name: "test remove",
			command: ifCommand{
				active:      true,
				includeLine: false,
			},
			loc:  "var name, namespace string",
			want: "",
			wantCommand: ifCommand{
				active:      true,
				includeLine: false,
			},
		},
		{
			name: "test deactivation",
			command: ifCommand{
				active:      true,
				includeLine: false,
			},
			loc:  "// opencontrolplane-gen:fi",
			want: "",
			wantCommand: ifCommand{
				active:      false,
				includeLine: false,
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
