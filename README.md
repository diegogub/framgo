# framgo
framgo - is small framework based on some useful structs ,
interfaces and gorilla mux router.
It's still under heavy development, so if you like it or
have any new idea to add please feel free to contact me or open an
issue!

Example App:
-----------
~~~
package main

import (
	"fmt"
	"github.com/diegogub/framgo"
	"net/http"
)

type TestPage struct {
	framgo.Mapper
}

func NewTest() *TestPage {
	var t TestPage
	t.Register("GET", t.Get)
	t.Register("POST", t.Post)
	return &t
}

func (t *TestPage) Endpoints() []string {
	return []string{"/test"}
}

func (t *TestPage) Type() string {
	return "plain"
}

func (t *TestPage) Get(vars map[string]string, r *http.Request) *framgo.HttpResponse {
	re := framgo.Redirect("http://google.com", 302)
	return re
}

func (t *TestPage) Post(vars map[string]string, r *http.Request) *framgo.HttpResponse {
	res := framgo.NewResource()
	res.Plain = []byte("post action")
	re := framgo.NewHttpResponse(200, res, "", "plain")
	return re
}

func (t *TestPage) Template() string {
	return "test"
}

func main() {
	fmt.Println("Starting web..")
	wc := framgo.New()
	wc.Verbose = true
	wc.LoadTemplates(".", ".html")

	t := NewTest()
	wc.AddPage(t)
	wc.Start("4545")
}
~~~
