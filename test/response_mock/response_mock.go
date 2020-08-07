package responsemock

import (
	"io/ioutil"
	"path/filepath"
	"runtime"
)

// JSON reads json file as a response body string
func JSON(path string) string {
	_, b, _, _ := runtime.Caller(0)
	path = filepath.Dir(b) + path
	result, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(result)
}
