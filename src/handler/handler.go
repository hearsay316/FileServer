package handler

import (
	"encoding/json"
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
		defer newFile.Close()
		fileMeta.FileSize, err = io.Copy(newFile, file)
		if err != nil {
			fmt.Println("io.Copy", err)
			return
		}
		_, _ = newFile.Seek(0, 0)
		fileMeta.FileSha1 = util.FileSha1(newFile)
		meta.UpdataFileMeta(fileMeta)
		fmt.Println(fileMeta)
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

/*func FileQueryHandle(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	filehash:= r.Form.Get("filehash")
	data,err:=json.Marshal(filehash)
	if err!=nil{
		w.WriteHeader()
	}
}*/
// GetFileMetaHandle获取hash
func GetFileMetaHandler(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	filehash := r.Form["filehash"][0]
	fMeta := meta.GetFileMeta(filehash)
	data, err := json.Marshal(fMeta)
	fmt.Println("err", fMeta)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("err")
		return
	}
	_, _ = w.Write(data)
}
func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	filehash := r.Form["filehash"][0]
	fMetafile := meta.GetFileMeta(filehash)
	f, err := os.Open(fMetafile.Location)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()
	data, err := ioutil.ReadAll(f)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "application/octect-stream")
	w.Header().Set("content-disposition", "attachment;filename=\""+fMetafile.FileName+"\"")
	_, _ = w.Write(data)
}
func FileUpdateMetaHandler(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	opType := r.Form.Get("op")
	fileHash := r.Form.Get("filehash")
	newName := r.Form.Get("filename")
	if opType != "0" {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	fileMetas := meta.GetFileMeta(fileHash)
	fileMetas.FileName = newName
	meta.UpdataFileMeta(fileMetas)
	data, err := json.Marshal(fileMetas)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(data)
}
func FileDeleteHandler(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	fileHash := r.Form.Get("filehash")
	fmt.Println(fileHash)
	fMeta := meta.GetFileMeta(fileHash)
	err := os.Remove(fMeta.Location)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	meta.RemoveFileMeta(fileHash)
	w.WriteHeader(http.StatusOK)
}
