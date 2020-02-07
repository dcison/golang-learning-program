package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func uploadHandler(w http.ResponseWriter, request *http.Request) {
	dirPath,_ := os.Getwd();
	newFile, err := os.Create(dirPath + "/newFile.png")
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(newFile, request.Body)
	if err != nil {
		panic(err)
	}
	w.Write([]byte("upload success"))


}

func main() {
	fmt.Println("服务启动")
	http.HandleFunc("/upload", uploadHandler)
	http.ListenAndServe(":5050", nil)
}