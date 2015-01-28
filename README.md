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
}

func NewTest() *TestPage {
	var t TestPage
	return &t
}

func (t *TestPage) Endpoints() *framgo.RequestMap {
	rm := framgo.NewRequestMap()
	rm.Add("/test", t.Get)
	rm.Add("/admin", t.Post)
	return rm
}

func (t *TestPage) Get(rc *framgo.RequestContext, r *http.Request) *framgo.HttpResponse {
  // we check if it's AJAX. We can define diffent response
	if rc.AJAX {
		res := framgo.NewResource()
		res.JSON(map[string]interface{}{"type": "AJAX REQUEST"})
		re := framgo.NewHttpResponse(201, res, "", "plain")
		return re
	} else {
		res := framgo.NewResource()
// we can define links into template
		res.AddLink("CSS", "http:google.com/")
		re := framgo.NewHttpResponse(200, res, "test", "html")
		return re
	}
}

func (t *TestPage) Post(rc *framgo.RequestContext, r *http.Request) *framgo.HttpResponse {
	res := framgo.NewResource()
	res.Plain = []byte("hola me gusta Golang")
	re := framgo.NewHttpResponse(201, res, "", "plain")
	return re
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
test.html template
~~~
{{ define "test" }}
<html>
<head>
  <script src="//code.jquery.com/jquery-1.11.2.min.js"></script>
</head>
<p>test 2</p>
<p id="ajax"></p>
{{ range .Links.CSS}}
<p>{{ . }}</p>
{{ end }}

<script>
  $.ajax({
    url: "/test",
    success: function( data ) {
      $( "#ajax" ).html( data );
    }
  });
</script>
</html>
{{ end }}

~~~
