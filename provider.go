package estore

import (
	proto "github.com/golang/protobuf/proto"
)

type Provider interface {
	GetState() ProviderState
}

type ProviderState interface {
	Restart()
	GetSnapshotInterval() int
	GetSnapshot(actorName string) (snapshot interface{}, eventIndex int, ok bool)
	GetEvents(actorName string, eventIndexStart int, callback func(e interface{}))
	PersistEvent(actorName string, eventIndex int, event proto.Message)
	PersistSnapshot(actorName string, eventIndex int, snapshot proto.Message)
}
