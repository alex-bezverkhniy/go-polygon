package yaml

import (
	"reflect"
	"testing"
)

func Test_readFile(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		{
			"happy",
			args{"./sample.yml"},
			//map[string]interface{}{"firstName": "Alex", "lastName": "Bez", "cars": [...]string{"Mazda", "Subaru"}},
			map[string]interface{}{"firstName": string("Alex"), "lastName": "Bez"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readFile(tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("readFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readFile() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}
