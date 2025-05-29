package etherum

import "fmt"

type EventPairCreate struct {
}

func NewEventPairCreate() *EventPairCreate {
	return &EventPairCreate{}
}

func (e *EventPairCreate) Run() {
	fmt.Println("etherum task run")
}
