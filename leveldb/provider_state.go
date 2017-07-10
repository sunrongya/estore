package leveldb

import (
	"fmt"
	"sync"

	proto "github.com/golang/protobuf/proto"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

type providerState struct {
	*Provider
	_db     *leveldb.DB
	_closed bool
	_wg     sync.WaitGroup
}

func newProviderState(provider *Provider) *providerState {
	return &providerState{
		Provider: provider,
		_closed:  true,
	}
}

func (state *providerState) Restart() {
	//wait for any pending writes to complete
	state._wg.Wait()

	if !state._closed {
		state.close()
	}

	state.open()
}

func (provider *Provider) GetSnapshotInterval() int {
	return provider._snapshotInterval
}

func (state *providerState) GetSnapshot(actorName string) (snapshot interface{}, eventIndex int, ok bool) {
	value, err := state._db.Get([]byte(formatSnapshotKey(actorName)), nil)
	if err != nil || len(value) == 0 {
		return nil, 0, false
	}
	envelope, err := state.parseEnvelope(value)
	if err != nil {
		return nil, 0, false
	}

	return envelope.message(), envelope.EventIndex, true
}

func (state *providerState) PersistSnapshot(actorName string, eventIndex int, snapshot proto.Message) {
	key := formatSnapshotKey(actorName)
	envelope := newEnvelope(snapshot, "snapshot", eventIndex)
	state.persistEnvelope(key, envelope)
}

func (state *providerState) GetEvents(actorName string, eventIndexStart int, callback func(e interface{})) {
	iter := state._db.NewIterator(&util.Range{
		Start: []byte(formatEventKey(actorName, eventIndexStart)),
		Limit: []byte(formatNextEventKeyPrefix(actorName)),
	}, nil)
	defer iter.Release()

	for iter.Next() {
		fmt.Printf("%s => %s\n", iter.Key(), iter.Value())
		envelope, err := state.parseEnvelope(iter.Value())
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		callback(envelope.message())
	}
}

func (state *providerState) PersistEvent(actorName string, eventIndex int, event proto.Message) {
	key := formatEventKey(actorName, eventIndex)
	envelope := newEnvelope(event, "event", eventIndex)
	state.persistEnvelope(key, envelope)
}

func (state *providerState) open() {
	db, err := leveldb.OpenFile(state.Path(), nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	state._db = db
	state._closed = false
}

func (state *providerState) close() {
	state._db.Close()
	state._closed = true
}

func (state *providerState) persistEnvelope(key string, envelope *envelope) {
	value, err := state._coder.Encode(envelope)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if err := state._db.Put([]byte(key), value, nil); err != nil {
		fmt.Println(err.Error())
	}
}

func (state *providerState) parseEnvelope(value []byte) (*envelope, error) {
	result := &envelope{}
	err := state._coder.Decode(value, result)
	return result, err
}
