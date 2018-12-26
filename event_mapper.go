package event

import (
	"encoding/json"
	"github.com/puneetdev0509/logger"
	"golang.org/x/net/context"
)

type OldEvent struct {
	Id         int64                  `json:"event_id"`
	New        map[string]interface{} `json:"data"`
	Offset     int64
	Meta       Meta
	Skipped    bool
	SkipReason string
}

type Meta struct {
	GlobalContextId int64
	OperationId     int64
	AuthorizerId    string
	OrgId           string
}

func GetOldEventFromNewEvent(ctx context.Context, record Event) (*OldEvent, error) {
	parseEventData := make(map[string]interface{})
	eventData := []byte(record.GetEventData())

	err := json.Unmarshal(eventData, &parseEventData)
	if err != nil {
		logger.Info("Could not convert data %s to map due to %s", eventData, err.Error())
		return nil, err
	}

	oldEvent := &OldEvent{
		Id:     record.GetEventId(),
		New:    parseEventData["data"].(map[string]interface{}),
		Offset: record.GetEventOffset(),
		Meta: Meta{
			GlobalContextId: record.GetGlobalContextId(),
			OperationId:     record.GetOperationId(),
			AuthorizerId:    record.GetAuthorizerId(),
			OrgId:           record.GetOrgId(),
		},
	}

	return oldEvent, nil
}
