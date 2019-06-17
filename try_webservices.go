package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
)

import (
	"html/template"
)

func upload(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		// GET
		t, _ := template.ParseFiles("upload.gtpl")

		t.Execute(w, nil)

	} else if r.Method == "POST" {
		// Post
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		//fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()

		io.Copy(f, file)
		path := handler.Filename
		var fpath = "./test/" + path
		segmentation(fpath, path)
		var segpath = "./data/splash_" + handler.Filename + ".png"
		final_result := detectDiease(segpath, r.PostFormValue("age"), r.PostFormValue("pain"), r.PostFormValue("h"))
		fmt.Println(final_result)
	} else {
		fmt.Println("Unknown HTTP " + r.Method + "  Method")
	}
}

func segmentation(img string, filename string) {
	app := "balloon.py"
	arg1 := "--weights=mask_rcnn_balloon_0030.h5"
	arg2 := "--image=" + img
	arg3 := "--png=" + filename
	//var stdout bytes.Buffer
	c := exec.Command("python", app, arg1, arg2, arg3)
	//c.Stdout = &stdout
	out, _ := c.Output()
	fmt.Println(string(out))
}
func detectDiease(img string, age string, pain string, h string) string {
	app := "TwoSecondInONe.py"
	arg1 := "--age=" + age
	arg4 := "--h=" + h
	arg5 := "--pain=" + pain
	arg6 := "--img=" + img
	c := exec.Command("python", app, arg1, arg2, arg4, arg5, arg6)
	out, _ := c.Output()
	fmt.Println("TheResult:" + string(out))
	return string(out)
}

func main() {
	http.HandleFunc("/upload", upload)
	http.ListenAndServe(":5000", nil) // setting listening port
}
