package main

import "github.com/webview/webview"
import "os"
import "log"
import "io/ioutil"
import "net/url"
import "fmt"

func main() {
	debug := true
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("Test 0.0.1")
	w.SetSize(550, 500, webview.HintNone)
        path, err := os.Getwd()
        if err != nil {
           log.Println(err)
        }
        log.Println(path)
        w.Bind("signIn", func(returned string) {
                log.Println("Sign In button clicked!")
                go signIn(w, returned)
        })
	dat, err := ioutil.ReadFile(path+"/index.html")
	w.Navigate("data:text/html,"+url.QueryEscape(string(dat)))
	w.Run()
}

func signIn(p webview.WebView, returned string) {
	js := fmt.Sprintf("'Sign In button event received from %s'", returned)

	p.Dispatch(func() {
		p.Eval("document.getElementById('label').innerHTML = "+js)
	})
}
