package handlers

import "testing"

func TestGetReferentsURI(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test should return correct URI for the given ID",
			args: args{
				"10347",
			},
			want: "https://api.genius.com/referents?web_page_id=10347",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetReferentsURI(tt.args.key); got != tt.want {
				t.Errorf("GetReferentsURI() = %v, want %v", got, tt.want)
			}
		})
	}
}
