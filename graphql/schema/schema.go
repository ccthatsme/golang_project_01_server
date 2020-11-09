package schema

import (
	"io/ioutil"
	"log"
)

func GetRootSchema(path string) string {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err.Error())
		return ""
	}

	return string(b)
}
