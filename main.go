/*
 * @Author: bingo
 * @Date: 2018-11-16 10:55:02
 */
package main

import (
	"bts-helloworld/logger"
	"fmt"
	"net/http"
	"html/template"
)

func Hello(response http.ResponseWriter, request *http.Request) {
	logger.Info(fmt.Sprintf("Hello enter, request:%v.",request))
	tmpl, err := template.ParseFiles("./views/hello.html")
	if err != nil {
		fmt.Println("Error happened..")
	}

	tmpl.Execute(response, nil)
}

func MyHello(response http.ResponseWriter, request *http.Request) {
	logger.Info(fmt.Sprintf("MyHello enter, request:%v.",request))
	tmpl, err := template.ParseFiles("./views/myhello.html")
	if err != nil {
		fmt.Println("Error happened..")
	}

	tmpl.Execute(response, nil)
}

func Login(response http.ResponseWriter, request *http.Request) {
	logger.Info(fmt.Sprintf("Login enter, request:%v.",request))
	tmpl, err := template.ParseFiles("./views/demo.html")
	if err != nil {
		fmt.Println("Error happened..")
	}

	tmpl.Execute(response, nil)
}

func Common(response http.ResponseWriter, request *http.Request) {
	logger.Info(fmt.Sprintf("Common enter, request:%v.",request))
	tmpl, err := template.ParseFiles("./views/common.html")
	if err != nil {
		fmt.Println("Error happened..")
	}

	tmpl.Execute(response, nil)
}

func TableContentNav(response http.ResponseWriter, request *http.Request) {
	logger.Info(fmt.Sprintf("TableContentNav enter, request:%v.",request))
	tmpl, err := template.ParseFiles("./views/table-content-nav.html")
	if err != nil {
		fmt.Println("Error happened..")
	}

	tmpl.Execute(response, nil)
}

func Crontab(response http.ResponseWriter, request *http.Request) {
	logger.Info(fmt.Sprintf("Crontab enter, request:%v.",request))
	tmpl, err := template.ParseFiles("./views/crontab.html")
	if err != nil {
		fmt.Println("Error happened..")
	}

	tmpl.Execute(response, nil)
}

func RemJs(response http.ResponseWriter, request *http.Request) {
	logger.Info(fmt.Sprintf("RemJs enter, request:%v.",request))
	tmpl, err := template.ParseFiles("./views/rem-js.html")
	if err != nil {
		fmt.Println("Error happened..")
	}

	tmpl.Execute(response, nil)
}

func MyLogin(response http.ResponseWriter, request *http.Request) {
	logger.Info(fmt.Sprintf("MyLogin enter, request:%v.",request))
	tmpl, err := template.ParseFiles("./views/mylogin.html")
	if err != nil {
		fmt.Println("Error happened..")
	}

	tmpl.Execute(response, nil)
}

func MyAuth(response http.ResponseWriter, request *http.Request) {
	logger.Info(fmt.Sprintf("MyAuth enter, request:%v.",request))
	tmpl, err := template.ParseFiles("./views/myauth.html")
	if err != nil {
		fmt.Println("Error happened..")
	}

	tmpl.Execute(response, nil)
}

func main() {
	logger.Error(fmt.Sprintf("bts-helloworld start..."))

	http.Handle("/css/", http.StripPrefix("/static/css/", http.FileServer(http.Dir("/static/css"))))
	http.Handle("/jquery-3.3.1/", http.FileServer(http.Dir("static/jquery-3.3.1")))
	http.Handle("/js/", http.FileServer(http.Dir("static/js")))

	http.HandleFunc("/login",Login)
	http.HandleFunc("/hello",Hello)
	http.HandleFunc("/myhello",MyHello)
	http.HandleFunc("/common",Common)
	http.HandleFunc("/nav",TableContentNav)
	http.HandleFunc("/cron",Crontab)
	http.HandleFunc("/remjs",RemJs)
	http.HandleFunc("/mylogin",MyLogin)
	http.HandleFunc("/myauth",MyAuth)

	http.ListenAndServe(":8001", nil)
}
