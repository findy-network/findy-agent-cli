package utils

import "testing"

func TestParseLoggingArgs(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantUse bool
		wantLvl int
	}{
		{"", args{""}, false, 0},
		{"-v 5", args{"-v 5"}, false, 0},
		{"-v=5 -logtostderr=true", args{"-v=5 -logtostderr"}, true, 5},
		{"-v=3 -logtostderr", args{"-v=3 -logtostderr"}, true, 3},
		{"-v 3 -logtostderr", args{"-v 3 -logtostderr"}, true, 3},
		{"-logtostderr -v=3", args{"-logtostderr -v=3"}, true, 3},
		{"-logtostderr -v 3", args{"-logtostderr -v 3"}, true, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUse, gotLvl := ParseLoggingArgs(tt.args.s)
			if gotUse != tt.wantUse {
				t.Errorf("ParseLoggingArgs() gotUse = %v, want %v", gotUse, tt.wantUse)
			}
			if gotLvl != tt.wantLvl {
				t.Errorf("ParseLoggingArgs() gotLvl = %v, want %v", gotLvl, tt.wantLvl)
			}
		})
	}
}
