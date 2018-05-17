package main

import (
	"net/http"
	"fmt"
	"strings"
	"log"
	"html/template"
	"net/url"
)

func sayHelloName(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k,v := range r.Form{
		fmt.Println("key", k)
		fmt.Println("value:" , strings.Join(v,""))
	}
	fmt.Fprintf(w, "Hello Chaos")
}

func login(w http.ResponseWriter, r *http.Request){
	fmt.Println("method", r.Method)
	if r.Method == "GET"{
		v := url.Values{}
		fmt.Println(v)
		t,_ := template.ParseFiles("login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		r.ParseForm()
		fmt.Println("username:" , r.Form["username"])
		fmt.Println("password", r.Form["password"])
	}
}
func main(){
	http.HandleFunc("/", sayHelloName)
	http.HandleFunc("/login", login)
	error := http.ListenAndServe(":9090", nil)
	if error != nil{
		log.Fatal("ListenAndServe:", error)
	}
}
