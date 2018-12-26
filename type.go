package event

import (
	"context"
	"time"
)

// CDC of an event will be represented by this.
type Event interface {
	GetEventId() int64
	GetOperationId() int64
	GetAuthorizerId() string
	GetGlobalContextId() int64
	GetOrgId() string
	GetEnvironment() string
	GetEventData() string
	GetIngestionTimestamp() time.Time
	GetEventOffset() int64
}

// CanalEvent implements Event interface. Processors should not use CanalEvent directly,
// instead they must use the methods provided by the Event interface
type CanalEvent struct {
	Id                 int64  `json:"id"`
	OperationId        int64  `json:"operation_id"`
	GlobalContextId    int64  `json:"global_context_id"`
	AuthorizerId       string `json:"authorizer_id"`
	OrgId              string `json:"org_id"`
	Env                string `json:"env"`
	EventData          string `json:"event_data"`
	IngestionTimestamp time.Time
	Offset             int64
}

func (ce *CanalEvent) GetEventId() int64                { return ce.Id }
func (ce *CanalEvent) GetOperationId() int64            { return ce.OperationId }
func (ce *CanalEvent) GetAuthorizerId() string          { return ce.AuthorizerId }
func (ce *CanalEvent) GetGlobalContextId() int64        { return ce.GlobalContextId }
func (ce *CanalEvent) GetOrgId() string                 { return ce.OrgId }
func (ce *CanalEvent) GetEnvironment() string           { return ce.Env }
func (ce *CanalEvent) GetEventData() string             { return ce.EventData }
func (ce *CanalEvent) GetIngestionTimestamp() time.Time { return ce.IngestionTimestamp }
func (ce *CanalEvent) GetEventOffset() int64            { return ce.Offset }

type ResponseRecord struct {
	Id       int64
	Response string
	UserId   string
	SeriesId string
	EntityId string
}

type EventResponse struct {
	LastOffset int64
	Records    []*ResponseRecord
}

type Processor interface {
	//takes a list of record and respond with number of event consumed along with error
	Run(context.Context, []Event) (*EventResponse, error)
}
