package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	t "time"

	pb "event_api/pkg/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedEventServer
}

var eventsList []*pb.EventInfo
var lastUsedId int64 = 0

func (s *server) Test(ctx context.Context, in *pb.TestRequest) (*pb.TestResponse, error) {
	fmt.Printf("Received request: %v\n", in.ProtoReflect().Descriptor().FullName())
	return &pb.TestResponse{
		Answer: "Worked!!",
	}, nil
}

func (s *server) MakeEvent(ctx context.Context, in *pb.MakeEventRequest) (*pb.MakeEventResponse, error) {
	fmt.Printf("Received request: %v\n", in.ProtoReflect().Descriptor().FullName())
	lastUsedId++
	tmpEventInfo := &pb.EventInfo{
		Senderid: in.Senderid,
		Eventid:  lastUsedId,
		Name:     in.Name,
		Time:     in.Time,
	}
	eventsList = append(eventsList, tmpEventInfo)
	fmt.Printf("NewEvent {%v}", tmpEventInfo)
	return &pb.MakeEventResponse{
		Eventid: lastUsedId,
	}, nil
}

func (s *server) GetEvent(ctx context.Context, in *pb.GetEventRequest) (*pb.GetEventResponse, error) {
	fmt.Printf("Received request: %v\n", in.ProtoReflect().Descriptor().FullName())
	tmpEventInfo := getSenderEvents(in.Eventid, in.Senderid)
	if tmpEventInfo != nil {
		return &pb.GetEventResponse{
			EventInfo: tmpEventInfo,
		}, nil
	}

	return nil, status.Error(404, "NotFound")

}

func (s *server) DeleteEvent(ctx context.Context, in *pb.DeleteEventRequest) (*pb.DeleteEventResponse, error) {
	fmt.Printf("Received request: %v\n", in.ProtoReflect().Descriptor().FullName())
	for i := 0; i < len(eventsList); i++ {
		if eventsList[i].Eventid == in.Eventid {
			copy(eventsList[i:], eventsList[i+1:])
			//eventsList[len(eventsList)-1] = *pb.EventInfo{}
			eventsList = eventsList[:len(eventsList)-1]
			return &pb.DeleteEventResponse{
				Senderid: in.Senderid,
			}, nil
		}
	}
	return nil, status.Error(404, "NotFound")
}
func (s *server) GetEvents(ctx context.Context, in *pb.GetEventsRequest) (*pb.GetEventsResponse, error) {
	fmt.Printf("Received request: %v\n", in.ProtoReflect().Descriptor().FullName())
	var events []*pb.EventInfo = getSenderEventsByTime(in.Time, in.Senderid)
	if len(events) != 0 {
		return &pb.GetEventsResponse{
			EventInfo: events,
		}, nil
	}
	return nil, status.Error(404, "NotFound")
}

func (s *server) GetEventsStream(in *pb.GetEventsStreamRequest, stream pb.Event_GetEventsStreamServer) error {
	fmt.Printf("Received request: %v\n", in.ProtoReflect().Descriptor().FullName())
	var events []*pb.EventInfo = getSenderEventsByTime(in.Time, in.Senderid)
	if len(events) != 0 {
		for _, event := range events {
			err := stream.SendMsg(&pb.GetEventsStreamResponse{
				EventInfo: event,
			})
			if err != nil {
				return status.Error(502, "BadGateway")
			}
		}
	} else {
		return status.Error(404, "NotFound")
	}
	return nil

}

func main() {
	var serverIP string
	var serverPort string
	flag.StringVar(&serverIP, "h", "127.0.0.1", "Server ip address")
	flag.StringVar(&serverPort, "p", "8090", "Server port ")

	listener, err := net.Listen("tcp", serverIP+":"+serverPort)
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterEventServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func getSenderEventsByTime(time, senderid int64) []*pb.EventInfo {
	var events []*pb.EventInfo
	for _, eventInfo := range eventsList {
		if eventInfo.Senderid == senderid {
			if t.Now().Unix() < time && eventInfo.Time < time && eventInfo.Time > t.Now().Unix() {
				events = append(events, eventInfo)
			}
			if t.Now().Unix() > time && eventInfo.Time > time && eventInfo.Time < t.Now().Unix() {
				events = append(events, eventInfo)
			}
		}
	}
	return events
}

func getSenderEvents(eventid, senderid int64) *pb.EventInfo {
	for _, eventInfo := range eventsList {
		if eventInfo.Eventid == eventid && eventInfo.Senderid == senderid {
			return eventInfo
		}
	}
	return nil
}
