package PeerToPeer

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"

	"blockchain/file"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:6040", "http service address")
var upgrader = websocket.Upgrader{} // use default options
type Hub struct {
	broadcast []string
}

var hub = new(Hub)

/*
func test() bool {
	fmt.Println("ici")
	return true
}*/
// InitServer creates and returns objects of
// the unexported type alertCounter.
func InitServer() {
	fmt.Println("InitServer")
	flag.Parse()
	log.SetFlags(0)
	InitServerForTest()
	//http.HandleFunc("/echo", echo)
	http.HandleFunc("/", home)
	http.HandleFunc("/message", messageHandler)
	log.Fatal(http.ListenAndServe(*addr, nil))

}
func InitServerForTest() {
	serverList := FileMgr.ReadFile("./client.json")
	for _, server := range serverList {
		fmt.Println("InitServerForTest")
		var addrTmp = flag.String(server.Name, server.Server, server.Description)
		u := url.URL{Scheme: "ws", Host: *addrTmp, Path: "/message"}
		hub.broadcast = append(hub.broadcast, u.String())
		fmt.Println("Server added")
	}
}
func messageHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer conn.Close()

	for {
		messageType, data, err := conn.ReadMessage()
		if err != nil {
			return
		}
		if messageType >= 0 {
			log.Printf("(%d) recv: %s", messageType, data)
		}
	}
}
func SendMessageToHub(data string) {
	fmt.Println("SendMessageToHub")

	for _, client := range hub.broadcast {
		SendMessageToClient(client, data)
	}
}

//SendMessage data ngn rmg,
func SendMessageToClient(client string, data string) {
	c, _, err := websocket.DefaultDialer.Dial(client, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	err2 := c.WriteMessage(websocket.TextMessage, []byte(data))

	if err2 != nil {
		log.Println("write:", err2)
		return
	}
	c.Close()
}

//SendMessage data ngn rmg,
func SendMessage(data string) {

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/message"}
	log.Printf("connecting to %s", u.String())
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	err2 := c.WriteMessage(websocket.TextMessage, []byte(data))

	if err2 != nil {
		log.Println("write:", err2)
		return
	}
	c.Close()

}
func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		log.Println("read:", err)
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	homeTemplate.Execute(w, "ws://"+r.Host+"/echo")
}

/*
var homeTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<head>
<meta charset="utf-8">
</head>
<body>
<table>
<tr><td valign="top" width="50%">
<p>It's alive! It's alive<p>
<form>
<button id="open">Open</button>
<button id="close">Close</button>
<p><input id="input" type="text" value="Hello world!">
<button id="send">Send</button>
</form>

</td></tr></table>
</body>
</html>
`))
*/
var homeTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<head>
<meta charset="utf-8">
<script>  
window.addEventListener("load", function(evt) {
    var output = document.getElementById("output");
    var input = document.getElementById("input");
    var ws;
    var print = function(message) {
        var d = document.createElement("div");
        d.innerHTML = message;
        output.appendChild(d);
    };
    document.getElementById("open").onclick = function(evt) {
        if (ws) {
            return false;
        }
        ws = new WebSocket("{{.}}");
        ws.onopen = function(evt) {
            print("OPEN");
        }
        ws.onclose = function(evt) {
            print("CLOSE");
            ws = null;
        }
        ws.onmessage = function(evt) {
            print("RESPONSE: " + evt.data);
        }
        ws.onerror = function(evt) {
            print("ERROR: " + evt.data);
        }
        return false;
    };
    document.getElementById("send").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        print("SEND: " + input.value);
        ws.send(input.value);
        return false;
    };
    document.getElementById("close").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        ws.close();
        return false;
    };
});
</script>
</head>
<body>
<table>
<tr><td valign="top" width="50%">
<p>
It's alive! It's alive

Click "Open" to create a connection to the server, 
"Send" to send a message to the server and "Close" to close the connection. 
You can change the message and send multiple times.
<p>
<form>
<button id="open">Open</button>
<button id="close">Close</button>
<p><input id="input" type="text" value="Hello world!">
<button id="send">Send</button>
</form>
</td><td valign="top" width="50%">
<div id="output"></div>
</td></tr></table>
</body>
</html>
`))
