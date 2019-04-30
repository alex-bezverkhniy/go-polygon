package template

import (
	"bytes"
	"testing"
)

func Test_renderTemplate(t *testing.T) {
	type args struct {
		data User
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
		wantErr bool
	}{
		{
			"Full name and prefix",
			args{data: User{FirstName: "Alex", LastName: "Bez", Prefix: "Mr"}}, "Hello Mr Alex Bez!", false,
		},
		{
			"First name and prefix",
			args{data: User{FirstName: "Alex", Prefix: "Mr"}}, "Hello Mr Alex !", false,
		},
		{
			"Last name and prefix",
			args{data: User{LastName: "Bez", Prefix: "Mr"}}, "Hello Mr  Bez!", false,
		},
		{
			"Full name and NO prefix",
			args{data: User{FirstName: "Alex", LastName: "Bez"}}, "Hello  Alex Bez!", false,
		},
		{
			"First name and NO prefix",
			args{data: User{FirstName: "Alex"}}, "Hello  Alex !", false,
		},
		{
			"Last name and No prefix",
			args{data: User{LastName: "Bez"}}, "Hello   Bez!", false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			if err := renderTemplate(out, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("renderTemplate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOut := out.String(); gotOut != tt.wantOut {
				t.Errorf("renderTemplate() = '%v', want '%v'", gotOut, tt.wantOut)
			}
		})
	}
}

func Test_renderTemplateByMap(t *testing.T) {
	type args struct {
		data map[string]string
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
		wantErr bool
	}{
		{
			"Full name and prefix",
			args{data: map[string]string{"FirstName": "Alex", "LastName": "Bez", "Prefix": "Mr"}}, "Hello Mr Alex Bez!", false,
		},
		{
			"First name and prefix",
			args{data: map[string]string{"FirstName": "Alex", "Prefix": "Mr"}}, "Hello Mr Alex !", false,
		},
		{
			"Last name and prefix",
			args{data: map[string]string{"LastName": "Bez", "Prefix": "Mr"}}, "Hello Mr  Bez!", false,
		},
		{
			"Full name and NO prefix",
			args{data: map[string]string{"FirstName": "Alex", "LastName": "Bez"}}, "Hello  Alex Bez!", false,
		},
		{
			"First name and NO prefix",
			args{data: map[string]string{"FirstName": "Alex"}}, "Hello  Alex !", false,
		},
		{
			"Last name and No prefix",
			args{data: map[string]string{"LastName": "Bez"}}, "Hello   Bez!", false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			if err := renderTemplateByMap(out, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("renderTemplateByMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOut := out.String(); gotOut != tt.wantOut {
				t.Errorf("renderTemplateByMap() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
