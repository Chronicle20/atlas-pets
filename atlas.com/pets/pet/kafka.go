package pet

const (
	EnvStatusEventTopic             = "EVENT_TOPIC_PET_STATUS"
	StatusEventTypeCreated          = "CREATED"
	StatusEventTypeDeleted          = "DELETED"
	StatusEventTypeSpawned          = "SPAWNED"
	StatusEventTypeDespawned        = "DESPAWNED"
	StatusEventTypeCommandResponse  = "COMMAND_RESPONSE"
	StatusEventTypeClosenessChanged = "CLOSENESS_CHANGED"
	StatusEventTypeFullnessChanged  = "FULLNESS_CHANGED"
	StatusEventTypeLevelChanged     = "LEVEL_CHANGED"
	StatusEventTypeSlotChanged      = "SLOT_CHANGED"
	StatusEventTypeExcludeChanged   = "EXCLUDE_CHANGED"

	DespawnReasonNormal  = "NORMAL"
	DespawnReasonHunger  = "HUNGER"
	DespawnReasonExpired = "EXPIRED"
)

type statusEvent[E any] struct {
	PetId   uint32 `json:"petId"`
	OwnerId uint32 `json:"ownerId"`
	Type    string `json:"type"`
	Body    E      `json:"body"`
}

type createdStatusEventBody struct {
}

type deletedStatusEventBody struct {
}

type spawnedStatusEventBody struct {
	TemplateId uint32 `json:"templateId"`
	Name       string `json:"name"`
	Slot       int8   `json:"slot"`
	Level      byte   `json:"level"`
	Closeness  uint16 `json:"closeness"`
	Fullness   byte   `json:"fullness"`
	X          int16  `json:"x"`
	Y          int16  `json:"y"`
	Stance     byte   `json:"stance"`
	FH         int16  `json:"fh"`
}

type despawnedStatusEventBody struct {
	TemplateId uint32 `json:"templateId"`
	Name       string `json:"name"`
	Slot       int8   `json:"slot"`
	Level      byte   `json:"level"`
	Closeness  uint16 `json:"closeness"`
	Fullness   byte   `json:"fullness"`
	OldSlot    int8   `json:"oldSlot"`
	Reason     string `json:"reason"`
}

type commandResponseStatusEventBody struct {
	Slot      int8   `json:"slot"`
	Closeness uint16 `json:"closeness"`
	Fullness  byte   `json:"fullness"`
	CommandId byte   `json:"commandId"`
	Success   bool   `json:"success"`
}

type closenessChangedStatusEventBody struct {
	Slot      int8   `json:"slot"`
	Closeness uint16 `json:"closeness"`
	Amount    int16  `json:"amount"`
}

type fullnessChangedStatusEventBody struct {
	Slot     int8 `json:"slot"`
	Fullness byte `json:"fullness"`
	Amount   int8 `json:"amount"`
}

type levelChangedStatusEventBody struct {
	Slot   int8 `json:"slot"`
	Level  byte `json:"level"`
	Amount int8 `json:"amount"`
}

type slotChangedStatusEventBody struct {
	OldSlot int8 `json:"oldSlot"`
	NewSlot int8 `json:"newSlot"`
}

type excludeChangedStatusEventBody struct {
	Items []uint32 `json:"items"`
}
