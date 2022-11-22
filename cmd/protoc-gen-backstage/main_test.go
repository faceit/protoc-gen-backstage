package main

import (
	"testing"
)

func Test_truncateTo63Chars(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{{
		name: "truncate 64 length string",
		args: args{name: "faceit-csgo_backupfile_parser-api-v1-CsgoBackupfileParserService"},
		want: "aceit-csgo_backupfile_parser-api-v1-CsgoBackupfileParserService",
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := truncateTo63Chars(tt.args.name); got != tt.want {
				t.Errorf("truncateTo63Chars() = %v, want %v", got, tt.want)
			}
		})
	}
}
