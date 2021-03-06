package san

import (
	"io/ioutil"
)

// Load is a sortcut for Readfile + Unmarshal
// use it like: err = san.Load("myfile.san", &myStruct)
func Load(file string, v interface{}) error {
	byt, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	err = Unmarshal(byt, v)
	return err
}
