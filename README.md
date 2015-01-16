# framgo
framgo - is small framework based on some useful structs ,
interfaces and gorilla mux router.
It's still under heavy development, so if you like it or
have any new idea to add please feel free to contact me or open an
issue!

Example App:
-----------
~~~
import (
"fmt"
"github.com/diegogub/framgo"
"github.com/gorilla/mux"
"net/http"
)

type TestPage struct {
}

func (t TestPage) Type()Type string {
return "plain"
}

func (t TestPage) Endpoints() framgo.Routes {
filter := func(r *mux.Route) {
r.Methods("GET", "POST")
}
Routeses := framgo.NewRoutes()
routes.AddRoute("/test", filter)
routes.AddRoute("/test/{data}", filter)
return routes
}

func (t TestPage) Template() string {
return "test"
}

func (t TestPage) Respond(vars map[stringng]string, r *http.Request)
*framgo.HttpResponse {
res := framgo.NewResource()
res.JSON([]string{"test"})
rhttp := framgo.NewHttpResponse(201, res)
return rhttp
}

func main() {
fmt.Println("Starting web..")
    wc := framgo.New()
wc.Verbose = true
wc.LoadTemplates(".", ".html")
      var t TestPage

// load test page
wc.AddPage(t)
// Start web
wc.Start("8080")
}
~~~
