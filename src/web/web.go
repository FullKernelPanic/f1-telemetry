package web

import (
	"f1telemetry/resources"
	"f1telemetry/src/util"
	"log"
	"math/rand"
	"net"
	"net/http"
	"text/template"

	"github.com/gorilla/websocket"
)

type Server struct {
	mux               *http.ServeMux
	savedsocketreader []socketReader
	id                int
}

func (w *Server) ListenAndServe(port string) {
	http.ListenAndServe(port, w.mux)
}

func CreateServer() *Server {
	myhttp := http.NewServeMux()
	myhttp.HandleFunc("/", homeHandler)
	myhttp.HandleFunc("/js", jsHandler)
	myhttp.HandleFunc("/css", cssHandler)
	myhttp.HandleFunc("/favicon.ico", iconHandler)

	ws := Server{myhttp, make([]socketReader, 0), rand.Intn(100)}

	myhttp.HandleFunc("/socket", ws.socketReaderCreate)

	return &ws
}

func cssHandler(w http.ResponseWriter, r *http.Request) {
	res, err := resources.Css()

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))

		return
	}

	w.Header().Add("Content-Type", "text/css")
	w.Write(res)
}

func jsHandler(w http.ResponseWriter, r *http.Request) {
	res, err := resources.Js()

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))

		return
	}

	w.Header().Add("Content-Type", "text/javascript")
	w.Write(res)
}

func iconHandler(w http.ResponseWriter, r *http.Request) {

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("index.html").ParseFS(resources.Html(), "html/*.gohtml")

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	w.WriteHeader(200)

	err2 := t.ExecuteTemplate(w, "index.gohtml", WebParams{IpAddress: util.IpAddress(), UdpPort: "20777"})

	if err2 != nil {
		log.Fatalln(err2)
		return
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (ws *Server) socketReaderCreate(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			log.Println(err)
		}
		r.Body.Close()

	}()

	con, _ := upgrader.Upgrade(w, r, nil)

	ptrSocketReader := socketReader{
		con: con,
	}
	ws.savedsocketreader = append(ws.savedsocketreader, ptrSocketReader)
}

type socketReader struct {
	con *websocket.Conn
}

func (ws *Server) Broadcast(data []byte) {
	for _, g := range ws.savedsocketreader {
		g.writeMsg(data)
	}
}

func (i *socketReader) writeMsg(data []byte) {
	err := i.con.WriteMessage(websocket.TextMessage, data)

	if err != nil {
		log.Println(err.Error())
	}
}

type WebParams struct {
	IpAddress net.IP
	UdpPort   string
}
