package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

//NOTE: There are likely some race conditions with subscribe (subs) and unsubscribe (subs).
// Could use mutex to fix this.

type queueService struct {
	subs map[uuid.UUID]*subscriber
}

func newQueueService() *queueService {
	return &queueService{subs: make(map[uuid.UUID]*subscriber)}
}

func (qs *queueService) subscribe(topic string) *subscriber {
	s := newSubscriber(topic)
	qs.subs[s.id] = s
	// Listen for unsubscribe
	go func() {
		<-s.u
		delete(qs.subs, s.id)
	}()
	return s
}

func (qs *queueService) publish(msg msg) {
	for _, s := range qs.subs {
		if s.topic == msg.topic {
			go s.publish(msg)
		}
	}
}

type msg struct {
	topic   string
	payload interface{}
}

type subscriber struct {
	id    uuid.UUID
	topic string
	m     chan msg      //chan that messages will be published on
	u     chan struct{} // unsubscribe chan
}

func newSubscriber(topic string) *subscriber {
	return &subscriber{id: uuid.New(), topic: topic, m: make(chan msg), u: make(chan struct{})}
}

func (s *subscriber) unsubscribe() {
	s.u <- struct{}{}
}

func (s *subscriber) publish(msg msg) {
	s.m <- msg
}

func listenForMessages(s *subscriber) {
	for {
		select {
		case msg := <-s.m:
			fmt.Printf("id: %s, msg: %s\n", s.id.String(), msg.payload)
		case <-s.u:
			return
		}

	}
}

func main() {
	qs := newQueueService()

	dogs1 := qs.subscribe("dogs")
	dogs2 := qs.subscribe("dogs")
	cats1 := qs.subscribe("cats")
	cats2 := qs.subscribe("cats")

	// Listen for n messages
	go listenForMessages(dogs1)
	go listenForMessages(dogs2)
	go listenForMessages(cats1)
	go listenForMessages(cats2)

	n := 10
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			qs.publish(msg{topic: "dogs", payload: "Doggy"})
			continue
		}
		qs.publish(msg{topic: "cats", payload: "Cats"})
	}

	time.Sleep(time.Second)

	dogs1.unsubscribe()
	dogs2.unsubscribe()
	cats1.unsubscribe()
	cats2.unsubscribe()
}
