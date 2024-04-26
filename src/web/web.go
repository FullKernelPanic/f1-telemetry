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
	myhttp.HandleFunc("/favicon.ico", iconHandler)

	ws := Server{myhttp, make([]socketReader, 0), rand.Intn(100)}

	myhttp.HandleFunc("/socket", ws.socketReaderCreate)

	return &ws
}

func iconHandler(w http.ResponseWriter, r *http.Request) {

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("index.html").ParseFS(resources.Htmls(), "html/*.gohtml")

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	w.WriteHeader(200)

	err2 := t.ExecuteTemplate(w, "index.gohtml", WebParams{IpAddress: util.IpAddress(), UdpPort: ":20777"})

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

	//ptrSocketReader.startThread()
}

type socketReader struct {
	con *websocket.Conn
}

/*
func (i *socketReader) startThread() {
	i.writeMsg([]byte("Please write your name"))

	go func() {
		defer func() {
			err := recover()
			if err != nil {
				log.Println(err)
			}
			log.Println("thread socketreader finish")
		}()

		for {
			i.read()
		}

	}()
}*/

func (ws *Server) Broadcast(data []byte) {
	for _, g := range ws.savedsocketreader {
		g.writeMsg(data)
	}
}

/*
	func (i *socketReader) read() {
		_, b, er := i.con.ReadMessage()
		if er != nil {
			panic(er)
		}

		//i.broadcast(string(b))

		log.Println(i.name + " " + string(b))
	}
*/
func (i *socketReader) writeMsg(data []byte) {
	i.con.WriteMessage(websocket.TextMessage, data)
}

type WebParams struct {
	IpAddress net.IP
	UdpPort   string
}
