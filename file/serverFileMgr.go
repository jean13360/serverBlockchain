package FileMgr

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type serverWs struct {
	Name        string `json:"name"`
	Server      string `json:"server"`
	Description string `json:"description"`
}

type serverList struct {
	Servers []serverWs
}

//ReadFile return server List
func ReadFile(pathFile string) (data []serverWs) {
	file, e := ioutil.ReadFile(pathFile)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	fmt.Printf("%s\n", string(file))

	json.Unmarshal(file, &data)
	return data
}
