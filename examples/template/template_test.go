package template

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"os"
	"testing"
	"time"
)

// 测试define
func handler(w http.ResponseWriter, r *http.Request) {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	var t *template.Template
	if rand.Intn(5) > 2 {

		//解析模板文件
		t = template.Must(template.ParseFiles("./hello.html", "./content1.html"))
	} else {
		//解析模板文件
		t = template.Must(template.ParseFiles("./hello.html", "./content2.html"))
	}

	//执行模板
	err := t.ExecuteTemplate(w, "model", "")
	if err != nil {
		return
	}
}
func TestTemplate(t *testing.T) {
	dir, _ := os.Getwd()
	fmt.Println("dir:", dir)
	http.HandleFunc("/template/index", handler)
	http.Handle("/template", http.RedirectHandler("/template/index", http.StatusSeeOther))
	http.Handle("/template/", http.RedirectHandler("/template/index", http.StatusSeeOther))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
