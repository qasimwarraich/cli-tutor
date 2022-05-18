package input

import (
	"reflect"
	"testing"
)

func TestInputFilter(t *testing.T) {
	type args struct {
		s          string
		vocabulary []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"Test Input Filter: No Arg",
			args{"whoami", []string{"whoami", "echo"}},
			[]string{"whoami"},
		},
		{
			"Test Input Filter: 1 Arg",
			args{"echo TEST", []string{"whoami", "echo"}},
			[]string{"echo", "TEST"},
		},
		{
			"Test Input Filter: 1 Arg (Quoted)",
			args{`echo "TEST"`, []string{"whoami", "echo"}},
			[]string{"echo", "\"TEST\""},
		},
		{
			"Test Input Filter: 2 Arg",
			args{"echo -e TEST", []string{"whoami", "echo"}},
			[]string{"echo", "-e", "TEST"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InputFilter(tt.args.s, tt.args.vocabulary); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InputFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRunCommand(t *testing.T) {
	type args struct {
		filtered_input []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Test Run Command (no output)",
			args{[]string{"echo"}},
			"\n",
		},
		{
			"Test Run Command",
			args{[]string{"echo", "TEST"}},
			"TEST\n",
		},
		{
			"Test Run Command with a flag",
			args{[]string{"echo", "-e", "TEST"}},
			"TEST\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RunCommand(tt.args.filtered_input); got != tt.want {
				t.Errorf("RunCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_contains(t *testing.T) {
	type args struct {
		s   string
		arr []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Test contains function (present)",
			args{"echo", []string{"echo", "TEST"}},
			true,
		},
		{
			"Test contains function (present, different element)",
			args{"TEST", []string{"echo", "TEST"}},
			true,
		},
		{
			"Test contains function (absent)",
			args{"test", []string{"echo", "TEST"}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := contains(tt.args.s, tt.args.arr); got != tt.want {
				t.Errorf("contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
