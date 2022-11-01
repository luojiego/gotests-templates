package f

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_f1(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		a int
	}
	tests := []struct {
		name string
		args args
		err  string
	}{
		{
			name: "simple",
			args: args{
				a: 0,
			},
			err: "",
		},
		{
			name: "err",
			args: args{
				a: 3,
			},
			err: "err 3 ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := f1(tt.args.a); err != nil {
				if tt.err == "" {
					t.Errorf("f1() return error: %v, but want nil", err)
				} else {
					assert.Contains(err.Error(), tt.err, "test: ["+tt.name+"] error message not eq")
				}
			}
		})
	}
}

func Test_f2(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		b int
	}
	tests := []struct {
		name string
		args args
		want int
		err  string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := f2(tt.args.b)
			if err != nil {
				if tt.err == "" {
					t.Errorf("f2() return error: %v, but want nil", err)
				} else {
					assert.Contains(err.Error(), tt.err, "test: ["+tt.name+"] error message not eq")
				}
				return
			}
			assert.Equalf(got, tt.want, "f2() = %v, want %v", got, tt.want)
		})
	}
}

func Test_f3(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		b int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 *D
		err   string
	}{
		{
			name:  "",
			args:  args{},
			want:  0,
			want1: &D{name: "foo"},
			err:   "empty",
		},
		{
			name: "b > 0",
			args: args{
				b: 2,
			},
			want:  0,
			want1: &D{name: "foo"},
			err:   "empty",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := f3(tt.args.b)
			if err != nil {
				if tt.err == "" {
					t.Errorf("f3() return error: %v, but want nil", err)
				} else {
					assert.Contains(err.Error(), tt.err, "test: ["+tt.name+"] error message not eq")
				}
				return
			}
			assert.Equalf(got, tt.want, "f3() got = %v, want %v", got, tt.want)
			assert.Equalf(got1, tt.want1, "f3() got1 = %v, want %v", got1, tt.want1)
		})
	}
}

func Test_f4(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		b int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 string
		want2 *D
		want3 *E
		err   string
	}{
		{
			name:  "simple",
			args:  args{},
			want:  0,
			want1: "",
			want2: &D{name: "foo"},
			want3: &E{
				id:   0,
				name: "foo-E",
				F:    &F{id: 100},
			},
			err: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, got3, err := f4(tt.args.b)
			if err != nil {
				if tt.err == "" {
					t.Errorf("f4() return error: %v, but want nil", err)
				} else {
					assert.Contains(err.Error(), tt.err, "test: ["+tt.name+"] error message not eq")
				}
				return
			}
			assert.Equalf(got, tt.want, "f4() got = %v, want %v", got, tt.want)
			assert.Equalf(got1, tt.want1, "f4() got1 = %v, want %v", got1, tt.want1)
			assert.Equalf(got2, tt.want2, "f4() got2 = %v, want %v", got2, tt.want2)
			assert.Equalf(got3, tt.want3, "f4() got3 = %v, want %v", got3, tt.want3)
		})
	}
}
