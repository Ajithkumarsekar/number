package number

import "testing"

func TestIsEven(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "positive even number",
			args: args{num: 24},
			want: true,
		},
		{
			name: "negative even number",
			args: args{num: -24},
			want: true,
		},
		{
			name: "odd number",
			args: args{num: 73},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEven(tt.args.num); got != tt.want {
				t.Errorf("IsEven() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsOdd(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "positive odd number",
			args: args{num: 327},
			want: true,
		},
		{
			name: "negative odd number",
			args: args{num: -3247},
			want: true,
		},
		{
			name: "even number",
			args: args{num: 43756942},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsOdd(tt.args.num); got != tt.want {
				t.Errorf("IsOdd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPrime(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "prime number",
			args: args{num: 13},
			want: true,
		},
		{
			name: "2 should be prime number",
			args: args{num: 2},
			want: true,
		},
		{
			name: "any number below 2 is not a prime number!",
			args: args{num: 0},
			want: false,
		},
		{
			name: "not a prime",
			args: args{num: 3242},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPrime(tt.args.num); got != tt.want {
				t.Errorf("IsPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}
