package gogenius

import (
	"testing"
)

func TestGetReferents(t *testing.T) {
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
			want: "something",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetReferents(tt.args.webPageID); got != tt.want {
				t.Errorf("GetReferents() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSong(t *testing.T) {
	type args struct {
		songID string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test should return correct URI for the given ID",
			args: args{
				"378195",
			},
			want: "something",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSong(tt.args.songID); got != tt.want {
				t.Errorf("GetSong() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetArtist(t *testing.T) {
	type args struct {
		artistID string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test should return correct URI for the given ID",
			args: args{
				"16775",
			},
			want: "something",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetArtist(tt.args.artistID); got != tt.want {
				t.Errorf("GetArtist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetArtistSongs(t *testing.T) {
	type args struct {
		artistID string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test should return correct URI for the given ID",
			args: args{
				"16775",
			},
			want: "something",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetArtistSongs(tt.args.artistID); got != tt.want {
				t.Errorf("GetArtistSongs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetWebpage(t *testing.T) {
	type args struct {
		pageURL string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test should return correct URI for the given ID",
			args: args{
				"https://docs.genius.com",
			},
			want: "something",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetWebpage(tt.args.pageURL); got != tt.want {
				t.Errorf("GetWebpage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearch(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test should return correct URI for the given ID",
			args: args{
				"Another%20brick%20in%20the%20wall",
			},
			want: "something",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Search(tt.args.query); got != tt.want {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAnnotation(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test should return correct URI for the given ID",
			args: args{
				"10225840",
			},
			want: "something",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAnnotation(tt.args.id); got != tt.want {
				t.Errorf("GetAnnotation() = %v, want %v", got, tt.want)
			}
		})
	}
}
