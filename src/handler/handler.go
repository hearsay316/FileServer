package handler

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func Uploaded(w http.ResponseWriter, r *http.Request) {
	file, err := os.Create("6.txt")
	if err != nil {
		fmt.Println("ioutilfff.ReadFile哦", err)
	}
	defer file.Close()
	if r.Method == "GET" {
		data, err := ioutil.ReadFile("./static/view/index.html")
		if err != nil {
			fmt.Println("ioutil.ReadFile", err)
		}
		_, err = io.WriteString(w, string(data))
	} else if r.Method == "POST" {
		//接受文件
	}
}
