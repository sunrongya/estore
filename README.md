# estore
事件存储库

## 功能列表

从实现[protoactor-go](https://github.com/AsynkronIT/protoactor-go)的
[persistence_provider.go](https://github.com/AsynkronIT/protoactor-go/blob/dev/persistence/persistence_provider.go)
开始完善事件存储功能

```go
//Provider is the abstraction used for persistence
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
```

### 支持的第三方存储库
- LevelDB
- BoltDB
- Redis
