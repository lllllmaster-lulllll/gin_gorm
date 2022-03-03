package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func f1(w http.ResponseWriter, r *http.Request) {
	//定义一个函数 kua
	//要么只有一个返回值,要么有两个返回值,第二个返回值必须是 error 类型
	kua := func(name string) (string, error) {
		return name + "真帅", nil
	}
	//定义模板
	t := template.New("f.tmpl")
	//告诉模板引擎,我现在多了一个自定义的函数 kua
	t.Funcs(template.FuncMap{
		"kua": kua,
	})

	//解析模板
	_, err := t.ParseFiles("./f.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err:%v\n", err)
	}
	name := "小王子"
	//渲染模板
	t.Execute(w, name)
}
func demo1(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	t, err := template.ParseFiles("./t.tmpl", "./ul.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err:%v\n", err)
		return
	}
	name := "小王子"
	//渲染模板
	t.Execute(w, name)

}
func main() {
	http.HandleFunc("/", f1)
	http.HandleFunc("/tmplDemo", demo1)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("HTTP server start failed, err:%v\n", err)
		return
	}
}
