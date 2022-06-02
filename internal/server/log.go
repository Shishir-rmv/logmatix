package server

import (
	"fmt"
	"sync"
)

var ErrOffsetNotFound = fmt.Errorf("offset not found")

type Record struct {
	Value  []byte `json:"value"`
	Offset uint64 `json:"offset"`
}
type Log struct {
	mu      sync.Mutex
	records []Record
}

func NewLog() *Log {
	return &Log{}
}

func (m *Log) Append(record Record) (uint64, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	record.Offset = uint64(len(m.records))
	m.records = append(m.records, record)
	return record.Offset, nil
}

func (m *Log) Read(offset uint64) (Record, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if offset >= uint64(len(m.records)) {
		return Record{}, ErrOffsetNotFound
	}
	return m.records[offset], nil
}
