

package main


import (
	"os"
	//"io/ioutil"
)


func checkPathExists(path string) (bool) { //(bool, error)
	_, err := os.Stat(path)
	if err == nil {
		return true//, nil
	}
	if os.IsNotExist(err) {
		return false//, nil
	}
	return true//, err
}

