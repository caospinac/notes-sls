package util

import "testing"

func Pointer[T comparable](value T) *T {
	return &value
}

func Test_getBodyString(t *testing.T) {
	type dummyStruct struct {
		Foo int `json:"foo"`
	}
	type args struct {
		body any
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "string",
			args: args{
				body: "hello world",
			},
			want: "hello world",
		},
		{
			name: "nil",
			args: args{
				body: nil,
			},
			want: "",
		},
		{
			name: "false",
			args: args{
				body: false,
			},
			want: "false",
		},
		{
			name: "array",
			args: args{
				body: []string{},
			},
			want: "[]",
		},
		{
			name: "empty bytes",
			args: args{
				body: []byte{},
			},
			want: "",
		},
		{
			name: "map",
			args: args{
				body: map[string]any{
					"foo": 123,
				},
			},
			want: `{"foo":123}`,
		},
		{
			name: "zero",
			args: args{
				body: 0,
			},
			want: "0",
		},
		{
			name: "zero pointer",
			args: args{
				body: Pointer(0),
			},
			want: "0",
		},
		{
			name: "pointer string",
			args: args{
				body: Pointer("string pointer"),
			},
			want: "string pointer",
		},
		{
			name: "struct",
			args: args{
				body: dummyStruct{123},
			},
			want: `{"foo":123}`,
		},
		{
			name: "struct pointer",
			args: args{
				body: &dummyStruct{123},
			},
			want: `{"foo":123}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getBodyString(tt.args.body); got != tt.want {
				t.Errorf("getBodyString() = %v, want %v", got, tt.want)
			}
		})
	}
}
