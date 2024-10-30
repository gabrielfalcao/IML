//// #""#                     dP                       dP
//// #  #                     88                       88
//// #  # 88d888b. .d8888b. d8888P .d8888b. 88d888b. d8888P
//// #  # 88'  `88 Y8ooooo.   88   88'  `88 88'  `88   88
//// #  # 88    88       88   88   88.  .88 88    88   88
//// #  # dP    dP `88888P'   dP   `88888P8 dP    dP   dP
//// ####
////
//// #"""""`'"""`Y#
//// #  mm.  mm.  #
//// #  ###  ###  # .d8888b. .d8888b. .d8888b. .d8888b. 88d888b. .d8888b. .d8888b. 88d888b.
//// #  ###  ###  # 88ooood8 Y8ooooo. Y8ooooo. 88ooood8 88'  `88 88'  `88 88ooood8 88'  `88
//// #  ###  ###  # 88.  ...       88       88 88.  ... 88    88 88.  .88 88.  ... 88
//// #  ###  ###  # `88888P' `88888P' `88888P' `88888P' dP    dP `8888P88 `88888P' dP
//// ##############                                                   .88
////                                                              d8888P
////
//// #""########
//// #  ########
//// #  ########
//// #  ########
//// #  ########
//// #         #
//// ###########
////
/*
xmpp_echo is a demo client that connect on an XMPP server and echo message received back to original sender.
*/

package main

import (
	"fmt"
	"log"
	"os"

	xmpp "gosrc.io/xmpp"
	"gosrc.io/xmpp/stanza"
)

func main() {
	config := xmpp.Config{
		TransportConfiguration: xmpp.TransportConfiguration{
			Address: "localhost:5222",
		},
		Jid:          "test@localhost",
		Credential:   xmpp.Password("test"),
		StreamLogger: os.Stdout,
		Insecure:     true,
		// TLSConfig: tls.Config{InsecureSkipVerify: true},
	}

	router := xmpp.NewRouter()
	router.HandleFunc("message", handleMessage)

	client, err := xmpp.NewClient(&config, router, errorHandler)
	if err != nil {
		log.Fatalf("%+v", err)
	}

	// If you pass the client to a connection manager, it will handle the reconnect policy
	// for you automatically.
	cm := xmpp.NewStreamManager(client, nil)
	log.Fatal(cm.Run())
}

func handleMessage(s xmpp.Sender, p stanza.Packet) {
	msg, ok := p.(stanza.Message)
	if !ok {
		_, _ = fmt.Fprintf(os.Stdout, "Ignoring packet: %T\n", p)
		return
	}

	_, _ = fmt.Fprintf(os.Stdout, "Body = %s - from = %s\n", msg.Body, msg.From)
	reply := stanza.Message{Attrs: stanza.Attrs{To: msg.From}, Body: msg.Body}
	_ = s.Send(reply)
}

func errorHandler(err error) {
	fmt.Println(err.Error())
}
