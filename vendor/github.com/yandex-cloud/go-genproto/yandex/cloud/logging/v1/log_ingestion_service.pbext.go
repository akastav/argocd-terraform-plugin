// Code generated by protoc-gen-goext. DO NOT EDIT.

package logging

import (
	status "google.golang.org/genproto/googleapis/rpc/status"
)

func (m *WriteRequest) SetDestination(v *Destination) {
	m.Destination = v
}

func (m *WriteRequest) SetResource(v *LogEntryResource) {
	m.Resource = v
}

func (m *WriteRequest) SetEntries(v []*IncomingLogEntry) {
	m.Entries = v
}

func (m *WriteRequest) SetDefaults(v *LogEntryDefaults) {
	m.Defaults = v
}

func (m *WriteResponse) SetErrors(v map[int64]*status.Status) {
	m.Errors = v
}