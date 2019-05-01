package yaml

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func readFile(fileName string) (map[string]string, error) {
	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var m map[string]string
	m = make(map[string]string)
	err = yaml.Unmarshal(yamlFile, &m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
