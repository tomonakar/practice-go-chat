package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

// templ represents a single template
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServeHTTP handles the HTTP request.
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, r)
}

func main() {
	var addr = flag.String("addr", ":8080", "アプリケーションのアドレス")
	flag.Parse()
	r := newRoom()
	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)
	// チャットルームを開始します
	go r.run()
	// Webサーバーを起動します
	log.Println("Webサーバを開始します。Port", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

// こんな書き方もできる
// func main() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
// 		w.Write([]byte(`
// 		<html>
// 			<head>
// 				<title>チャット</title>
// 			</head>
// 			<body>
// 				チャットしましょう!
// 			</body>
// 		</html>
// 		`))
// 	})
// 	if err := http.ListenAndServe(":8080", nil); err != nil {
// 		log.Fatal("ListenAndServe:", err)
// 	}
// }
