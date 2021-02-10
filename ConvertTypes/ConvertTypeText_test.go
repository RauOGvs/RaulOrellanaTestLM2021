package ConvertTypes

import "testing"

func TestConvertTypeTextToBinaryImplementation_ConvertByType(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Text To Binary",
			args: args{input: "RaUl"},
			want: "01010010011000010101010101101100",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ConvertTypeTextToBinaryImplementation{}
			if got := c.ConvertByType(tt.args.input); got != tt.want {
				t.Errorf("ConvertByType() = %v, want %v", got, tt.want)
			}
		})
	}
}
