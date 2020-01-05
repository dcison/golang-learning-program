package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/test", handler)
	http.HandleFunc("/404", notFount)
	http.HandleFunc("/pic", LookImg)
	http.ListenAndServe("localhost:8222", nil)
}

func notFount(res http.ResponseWriter, req *http.Request) {
	http.NotFoundHandler().ServeHTTP(res, req)
}

func indexHandler(res http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(res, "这里是首页")
}

func handler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "接受到的参数是 %q\n", req.URL.Path)
}

func LookImg(w http.ResponseWriter, r *http.Request) {
	fmt.Printf(r.URL.Scheme, r.URL.User)
	path := strings.Split(r.URL.Path, "/")
	var name string
	if len(path) > 1 {
		name = path[len(path)-1]
	} else {
		name = "null.png"
	}
	fmt.Printf(name)
	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Disposition", fmt.Sprintf("inline; filename=\"%s\"", name))

	file, err := ioutil.ReadFile("../../syllabus/go语言学习大纲.png")
	if err != nil {
		fmt.Fprintf(w, "查无此图片")
		return
	}
	w.Write(file)
}
