package commands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrefix(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		loc               string
		commandIdentifier string
		want              bool
	}{
		{
			name:              "matching prefix",
			loc:               "// opencontrolplane-gen:if a=b",
			commandIdentifier: ocpIf,
			want:              true,
		},
		{
			name:              "non-matching prefix",
			loc:               "// opencontrolplane-gen:if a=b",
			commandIdentifier: ocpReplace,
			want:              false,
		},
		{
			name:              "matching identifier but not prefix",
			loc:               "// abc opencontrolplane-gen:if a=b",
			commandIdentifier: ocpReplace,
			want:              false,
		},
		{
			name:              "abitrary input",
			loc:               "invalid line of code",
			commandIdentifier: "abc",
			want:              false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Prefix(tt.loc, tt.commandIdentifier)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_uncomment(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		loc  string
		want string
	}{
		{
			name: "uncomment comment",
			loc:  "// test comment",
			want: "test comment",
		},
		{
			name: "uncomment comment with identation",
			loc:  "   // test comment",
			want: "test comment",
		},
		{
			name: "uncomment comment with comment identifiert",
			loc:  "// test // comment",
			want: "test // comment",
		},
		{
			name: "not a comment with infix",
			loc:  "test // comment",
			want: "test // comment",
		},
		{
			name: "obviously not a comment",
			loc:  "test loc",
			want: "test loc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := uncomment(tt.loc)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_trimCommand(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		loc               string
		commandIdentifier string
		want              string
	}{
		{
			name:              "trim matching command",
			loc:               "// opencontrolplane-gen:if a=b",
			commandIdentifier: ocpIf,
			want:              "a=b",
		},
		{
			name:              "return unchanged line when not matching",
			loc:               "// opencontrolplane-gen:if a=b",
			commandIdentifier: ocpReplace,
			want:              "// opencontrolplane-gen:if a=b",
		},
		{
			name:              "abitrary input",
			loc:               "invalid line of code",
			commandIdentifier: "abc",
			want:              "invalid line of code",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := trimCommand(tt.loc, tt.commandIdentifier)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_arguments(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		loc               string
		commandIdentifier string
		want              []string
	}{
		{
			name:              "retrieve single argument",
			loc:               "// opencontrolplane-gen:if a=b",
			commandIdentifier: ocpIf,
			want:              []string{"a=b"},
		},
		{
			name:              "retrieve multiple argument",
			loc:               "// opencontrolplane-gen:if a=b c=d",
			commandIdentifier: ocpIf,
			want:              []string{"a=b", "c=d"},
		},
		{
			name:              "return empty result when not matching",
			loc:               "// opencontrolplane-gen:if a=b",
			commandIdentifier: ocpReplace,
			want:              []string{},
		},
		{
			name:              "return empty result on abitrary input",
			loc:               "invalid line of code",
			commandIdentifier: "abc",
			want:              []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := arguments(tt.loc, tt.commandIdentifier)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_assignments(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		loc               string
		commandIdentifier string
		want              []assignment
	}{
		{
			name:              "retrieve single argument",
			loc:               "// opencontrolplane-gen:if a=b",
			commandIdentifier: ocpIf,
			want:              []assignment{{left: "a", right: "b"}},
		},
		{
			name:              "retrieve multiple argument",
			loc:               "// opencontrolplane-gen:if a=b c=d",
			commandIdentifier: ocpIf,
			want: []assignment{
				{left: "a", right: "b"},
				{left: "c", right: "d"},
			},
		},
		{
			name:              "return empty result when not matching",
			loc:               "// opencontrolplane-gen:if a=b",
			commandIdentifier: ocpReplace,
			want:              []assignment{},
		},
		{
			name:              "return empty result on abitrary input",
			loc:               "invalid line of code",
			commandIdentifier: "abc",
			want:              []assignment{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := assignments(tt.loc, tt.commandIdentifier)
			assert.Equal(t, tt.want, got)
		})
	}
}
