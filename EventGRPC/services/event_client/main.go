package main

import (
	"context"
	"flag"
	"fmt"
	"io"

	pb "event_api/pkg/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var destinationIP string
	var destinationPort string
	var senderid int64
	flag.StringVar(&destinationIP, "dst", "127.0.0.1", "Server ip address")
	flag.StringVar(&destinationPort, "p", "8090", "Server port ")
	flag.Int64Var(&senderid, "sender-id", 1, "Your id")
	flag.Parse()
	conn, err := grpc.Dial(destinationIP+":"+destinationPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	client := pb.NewEventClient(conn)
	/*
		p := prompt.New(
			Executor,
			completer,
			prompt.OptionTitle("Event shell client"),
			prompt.OptionPrefix(">>>"),
			prompt.OptionInputTextColor(prompt.Yellow),
		)
		p.Run()
	*/
	for {
		var command string
		fmt.Scan(&command)
		switch command {
		case "MakeEvent":
			MakeEvent(client, senderid)
		case "GetEvent":
			GetEvent(client, senderid)
		case "DeleteEvent":
			DeleteEvent(client, senderid)
		case "GetEvents":
			GetEvents(client, senderid)
		case "GetEventsStream":
			GetEventsStream(client, senderid)
		default:
			fmt.Println("Error. Command not recognized")
		}
	}
}

func MakeEvent(client pb.EventClient, senderid int64) {
	var time int64
	var name string
	fmt.Scan(&time)
	fmt.Scan(&name)
	eventid, err := client.MakeEvent(context.Background(), &pb.MakeEventRequest{
		Senderid: senderid,
		Time:     time,
		Name:     name,
	})
	if err != nil {
		fmt.Printf("failed to make event: %v", err)
	}
	fmt.Printf("Created{%v}\n", eventid)
}
func GetEvent(client pb.EventClient, senderid int64) {
	var eventid int64
	fmt.Scan(&eventid)
	eventInfo, err := client.GetEvent(context.Background(), &pb.GetEventRequest{
		Senderid: senderid,
		Eventid:  eventid,
	})
	if err != nil {
		fmt.Printf("failed to get event: %v", err)
	}
	fmt.Printf("Event{%v}\n", eventInfo)
}
func DeleteEvent(client pb.EventClient, senderid int64) {
	var eventid int64
	fmt.Scan(&eventid)
	delete, err := client.DeleteEvent(context.Background(), &pb.DeleteEventRequest{
		Senderid: senderid,
		Eventid:  eventid,
	})
	if err != nil {
		fmt.Printf("failed to delete event: %v", err)
	}
	fmt.Printf("Deleted{%v}\n", delete)
}

func GetEvents(client pb.EventClient, senderid int64) {
	var time int64
	fmt.Scan(&time)
	events, err := client.GetEvents(context.Background(), &pb.GetEventsRequest{
		Senderid: senderid,
		Time:     time,
	})
	if err != nil {
		fmt.Printf("failed to get events: %v", err)
	}
	fmt.Printf("Events{%v}\n", events)
}
func GetEventsStream(client pb.EventClient, senderid int64) {
	var time int64
	fmt.Scan(&time)
	stream_events, _ := client.GetEventsStream(context.Background(), &pb.GetEventsStreamRequest{
		Senderid: senderid,
		Time:     time,
	})
	for {
		events, err := stream_events.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("failed to get events: %v", err)
			break
		}
		fmt.Printf("Events{%v}\n", events)
	}
}

/*func Executor(s string) {
	s = strings.TrimSpace(s)
	setCommand := strings.Split(s, " ")
	switch setCommand[0] {
	case "exit", "quit", "q":
		fmt.Println("Close")
		os.Exit(0)
		return
	}
}

func completer(d prompt.Document) []prompt.Suggest {
	var s []prompt.Suggest
	switch d.Text {
	case "M", "Ma", "Mak", "Make", "MakeE", "MakeEv", "MakeEve", "MakeEven", "MakeEvent":
		s = []prompt.Suggest{
			{Text: "MakeEvent", Description: "MakeEvent <sender-id> <time> <name>"},
		}
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}*/
