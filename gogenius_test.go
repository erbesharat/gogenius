package gogenius

import "testing"

func TestReferents(t *testing.T) {
	type args struct {
		webPageID string
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
			if got := Referents(tt.args.webPageID); got != tt.want {
				t.Errorf("Referents() = %v, want %v", got, tt.want)
			}
		})
	}
}
