package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
)

func main () {
	dirpath,_ := os.Getwd();
	datapath := path.Join(dirpath, "./png.png")
	file, err := os.Open(datapath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	res, err := http.Post("http://127.0.0.1:5050/upload","multipart/form-data", file)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	message, _ := ioutil.ReadAll(res.Body)
	fmt.Printf(string(message))
}


