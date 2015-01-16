package framgo

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// Defines a whole app, with pages,apis and handlers
type WebContainer struct {
	// lock
	lock sync.Mutex
	// Global resources, to inject in every request
	GlobalKey string
	Global    *Resource
	// Pages, could be json responses, html or just http
	Pages []WebPager
	// Logger interface
	log RequestLogger
	// main router
	router  *mux.Router
	tmpl    *template.Template
	Verbose bool
}

// New WebContainer
func New() *WebContainer {
	var wc WebContainer
	wc.Pages = make([]WebPager, 0)
	wc.router = mux.NewRouter()
	wc.Global = NewResource()
	wc.GlobalKey = "Global"
	return &wc
}

func (wc *WebContainer) LoadTemplates(dir, suffix string) error {
	if dir == "" || suffix == "" {
		return errors.New("Invalid suffix or dir")
	}
	tmps := template.New("tmp")

	fmt.Println("Load templates:")
	// Load all templates in every request! . We don't need performace,just to load template
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, suffix) {
			if wc.Verbose {
				fmt.Println("Matched template file - ", path)
			}
			tmps = template.Must(tmps.ParseFiles(path))
		}
		return nil
	})
	wc.tmpl = tmps
	return nil
}

func (wc *WebContainer) Start(port string) {
	// build routes
	wc.buildRouter()
	err := http.ListenAndServe(":"+port, wc.router)
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func (wc *WebContainer) SetGlobalKey(s string) {
	if s != "" {
		wc.GlobalKey = s
	}
}

// Register webpage into webcontainer
func (wc *WebContainer) AddPage(wp WebPager) {
	wc.lock.Lock()
	defer wc.lock.Unlock()

	if wp != nil {
		wc.Pages = append(wc.Pages, wp)
	}
}

// Write Response to the wire
func (wc *WebContainer) Write(w http.ResponseWriter, httr *HttpResponse, kind string, template string) {
	// set all cookies
	for _, c := range httr.Cookies {
		http.SetCookie(w, c)
	}
	// set all headers
	for k, v := range httr.Headers {
		if k != "" {
			w.Header().Set(k, v)
		}
	}

	if httr.Code > 0 {
		w.WriteHeader(httr.Code)
	}
	switch kind {
	//try to render template and pass resource
	case "html":
		if template == "" {
			panic("Empty template")
		}

		if httr.Res != nil {
			err := wc.tmpl.ExecuteTemplate(w, template, httr.Res.Data)
			if err != nil {
				panic(err)
			}
		}

	case "plain":
		if httr.Res != nil && len(httr.Res.Plain) > 0 {
			w.Write(httr.Res.Plain)
		} else {
			w.Write([]byte("Empty"))
		}
	}
}

func (wc *WebContainer) buildRouter() error {
	if len(wc.Pages) == 0 {
		return errors.New("No endpoints registered!")
	}

	for _, wp := range wc.Pages {
		var auxr *mux.Route
		for endpoint, filter := range wp.Endpoints() {
			auxr = wc.router.HandleFunc(endpoint, buildHandler(wp, wc))
			if filter != nil {
				filter(auxr)
			}
		}
	}

	return nil
}

// Return http handler from webpager, based in the webcontainer
func buildHandler(wp WebPager, wc *WebContainer) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if e := recover(); e != nil {
				ErrorHTML(e, w)
			}
		}()
		// get vars from request
		res := wp.Respond(mux.Vars(r), r)
		// merge global resources into key
		res.MergeResource(wc.GlobalKey, wc.Global)
		wc.Write(w, res, wp.Type(), wp.Template())
	}
}
