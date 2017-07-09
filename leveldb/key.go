package leveldb

import "fmt"

func formatEventKey(actorName string, eventIndex int) string {
	return fmt.Sprintf("%s:event:%010d", actorName, eventIndex)
}

func formatSnapshotKey(actorName string) string {
	return fmt.Sprintf("%s:snapshot", actorName)
}

func formatNextEventKeyPrefix(actorName string) string {
	return fmt.Sprintf("%s:f", actorName)
}
