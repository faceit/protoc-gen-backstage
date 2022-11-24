package main

import (
	"reflect"
	"regexp"
	"testing"

	plugin_go "github.com/golang/protobuf/protoc-gen-go/plugin"
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
		args: args{name: "spacer-quite_a_long_name_this-api-v1-WhatAnUnnecessaryLengthOfName"},
		want: "cer-quite_a_long_name_this-api-v1-WhatAnUnnecessaryLengthOfName",
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

func Test_regexTemplateName(t *testing.T) {
	type args struct {
		name  string
		regex *regexp.Regexp
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "prefix_before_-",
			args: args{
				name:  "justme-hello",
				regex: regexp.MustCompile("(.+)-.*")},
			want: "justme",
		},
		{
			name: "space_transform_test",
			args: args{
				name:  "space.test.api.V3.niceService",
				regex: regexp.MustCompile("(.+\\.)(.+\\.)api\\.(.+\\.)(.*)(?:Service)|(.+\\.)(.+\\.)api\\.(.+\\.)(.*)")},
			want: "space.test.V3.nice",
		},
		{
			name: "space_transform_NiceService",
			args: args{
				name:  "space.proto_name.api.v1.NiceService",
				regex: regexp.MustCompile("(.+\\.)(.+\\.)api\\.(.+\\.)(.*)(?:Service)|(.+\\.)(.+\\.)api\\.(.+\\.)(.*)")},
			want: "space.proto_name.v1.Nice",
		},
		{
			name: "space_transform_Nice_No_Service",
			args: args{
				name:  "space.proto_name.api.v1alpha1.Nice",
				regex: regexp.MustCompile("(.+\\.)(.+\\.)api\\.(.+\\.)(.*)(?:Service)|(.+\\.)(.+\\.)api\\.(.+\\.)(.*)")},
			want: "space.proto_name.v1alpha1.Nice",
		},
		{
			name: "keep original",
			args: args{
				name:  "space.csgo_backupfile_parser.api.v1.CsgoBackupfileParserService",
				regex: regexp.MustCompile("(.*)")},
			want: "space.csgo_backupfile_parser.api.v1.CsgoBackupfileParserService",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := regexTemplateName(tt.args.name, tt.args.regex); got != tt.want {
				t.Errorf("regexTemplateName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func stringPointer(original string) *string {
	return &original
}

func TestParseOptions(t *testing.T) {
	type args struct {
		req *plugin_go.CodeGeneratorRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *PluginOptions
		wantErr bool
	}{
		{
			name: "1 option",
			args: args{
				req: &plugin_go.CodeGeneratorRequest{
					Parameter: stringPointer("github.com/org/repo/"),
				},
			},
			want: &PluginOptions{
				OutputFileName:    "catalog-info.yaml",
				URLRoot:           "github.com/org/repo/",
				RegexNameTemplate: regexp.MustCompile("(.*)"),
			},
			wantErr: false,
		},
		{
			name: "2 options",
			args: args{
				req: &plugin_go.CodeGeneratorRequest{
					Parameter: stringPointer("github.com/org/repo/,my-info.yaml"),
				},
			},
			want: &PluginOptions{
				OutputFileName:    "my-info.yaml",
				URLRoot:           "github.com/org/repo/",
				RegexNameTemplate: regexp.MustCompile("(.*)"),
			},
			wantErr: false,
		},
		{
			name: "3 options",
			args: args{
				req: &plugin_go.CodeGeneratorRequest{
					Parameter: stringPointer("github.com/org/repo/,my-info.yaml,(.*)-(.*)"),
				},
			},
			want: &PluginOptions{
				OutputFileName:    "my-info.yaml",
				URLRoot:           "github.com/org/repo/",
				RegexNameTemplate: regexp.MustCompile("(.*)-(.*)"),
			},
			wantErr: false,
		},
		{
			name: "3 options missing 2nd",
			args: args{
				req: &plugin_go.CodeGeneratorRequest{
					Parameter: stringPointer("github.com/org/repo/,,(.*)-(.*)"),
				},
			},
			want: &PluginOptions{
				OutputFileName:    "catalog-info.yaml",
				URLRoot:           "github.com/org/repo/",
				RegexNameTemplate: regexp.MustCompile("(.*)-(.*)"),
			},
			wantErr: false,
		},
		{
			name: "3 options bad regex",
			args: args{
				req: &plugin_go.CodeGeneratorRequest{
					Parameter: stringPointer("github.com/org/repo/,,(.*)-(.*"),
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseOptions(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseOptions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseOptions() got = %v, want %v", got, tt.want)
			}
		})
	}
}
