package handler

import (
	"fmt"
	"io"
	"io/ioutil"
	"meta"
	"net/http"
	"os"
	"time"
	"util"
)

func Uploaded(w http.ResponseWriter, r *http.Request) {
	file, err := os.Create("6.txt")
	if err != nil {
		fmt.Println("ioutilfff.ReadFile哦", err)
		return
	}
	defer file.Close()
	if r.Method == "GET" {
		data, err := ioutil.ReadFile("./static/view/index.html")
		if err != nil {
			fmt.Println("ioutil.ReadFile", err)
			return
		}
		_, err = io.WriteString(w, string(data))
	} else if r.Method == "POST" {
		//接受文件
		file, head, err := r.FormFile("file")
		if err != nil {
			fmt.Println("f.FormFile", err)
			return
		}
		defer file.Close()
		fileMeta := meta.FileMeta{
			FileName: head.Filename,
			Location: "./static/tmp/" + head.Filename,
			UploadAt: time.Now().Format("2006-01-02 15:04:05"),
		}
		newFile, err := os.Create(fileMeta.Location)
		if err != nil {
			fmt.Println("os.Create", err)
			return
		}
		fileMeta.FileSize, err = io.Copy(newFile, file)
		if err != nil {
			fmt.Println("io.Copy", err)
			return
		}
		newFile.Seek(0, 0)
		fileMeta.FileSha1 = util.FileSha1(newFile)
		meta.UpdataFileMeta(fileMeta)
		// 页面按跳转
		http.Redirect(w, r, "/file/upload/suc", http.StatusFound)
	}
}
func UploadSuc(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "成功上传")
	if err != nil {
		fmt.Println("io.Copy", err)
		return
	}
}
