package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:contact@pierre.io\">contact@pierre.io</a>.")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1>FAQ Page</h1>
	<p>Here is a list of frequently asked questions:</p>
	<ul>
	<li>Is there a free version?</li>
	<p>Nope, gotta pay us!</p>
	<li>What are your support hours?</li>
	<p>Support ? Hahaha</p>
	<li>How do I contact support?</li>
	<p>Look, we don't really do this \"support\" thing you describe. Sounds interesting though!</p>
	</ul>`)
}

func paramTestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<p>The param of this page is <b>"+chi.URLParam(r, "testParam")+"</b></p>")

}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Page Not Found", http.StatusNotFound)
}


func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", homeHandler)
	router.Get("/contact", contactHandler)
	router.Get("/faq", faqHandler)
	router.Get("/param/{testParam}", paramTestHandler)
	router.NotFound(notFoundHandler)

	fmt.Println("starting the server on :3000...")
	http.ListenAndServe(":3000", router)
}
