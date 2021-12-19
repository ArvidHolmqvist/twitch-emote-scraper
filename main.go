package main

import (
	"fmt"
	"net"
	//"emote_scraper/emotes"
	"bufio"
	"net/textproto"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func getConnnection() net.Conn {
	conn, err := net.Dial("tcp", "irc.chat.twitch.tv:6667")
	if err != nil {
		panic(err)
	}
	return conn
}

func disconnect(conn net.Conn) {
	conn.Close()
}

func logon(conn net.Conn) {
	sendData(conn, "PASS oauth:95ajeukcnfk38v3rhog4fvwgir4hzq")
	sendData(conn, "NICK emote_scraper")
}

func sendData(conn net.Conn, message string) {
	fmt.Fprintf(conn, "%s\r\n", message)
}

func main() {
	// emoteSet := emotes.FetchEmotes(62300805)

	streamers := []string{"gorgc"}
	channel := make(chan string)

	for _, name := range streamers {
		println(name)
		conn := getConnnection()

		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		go func() {
			<-c
			disconnect(conn)
		}()
		go scrapeChat(conn, name, channel)
	}
	fmt.Println(<-channel)

}

func scrapeChat(conn net.Conn, streamName string, channel chan string) {
	logon(conn)
	tp := textproto.NewReader(bufio.NewReader(conn))
	sendData(conn, "JOIN #"+streamName)

	for {
		status, err := tp.ReadLine()
		if err != nil {
			panic(err)
		}
		fmt.Println(status)
		if strings.HasSuffix(status, "list") {
			break
		}
	}

	for {
		status, err := tp.ReadLine()
		if err != nil {
			panic(err)
		}

		if strings.HasPrefix(status, "PING") {
			sendData(conn, "PONG")
			continue
		}

		msg := strings.SplitN(status, ":", 3)

		fmt.Println("#"+streamName, time.Now().UnixMilli(), strings.Split(msg[1], "!")[0], ":", msg[2])
		// msg_arr := strings.Split(msg[2], " ")
		// for _, str := range msg_arr {
		// 	if emoteSet[str] {
		// 		fmt.Println("Emotes: ", str)
		// 	}
		// }
		fmt.Println("-----------------------------------------------------")
	}
	channel <- streamName + " done"
}
