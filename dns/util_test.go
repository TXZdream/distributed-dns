package dns

import "testing"

import "fmt"

func Test_CalculateHash(t *testing.T) {
	type args struct {
		raw string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				raw: "hello",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(CalculateHash(tt.args.raw))
		})
	}
}
