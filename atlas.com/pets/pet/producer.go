package pet

import (
	"github.com/Chronicle20/atlas-kafka/producer"
	"github.com/Chronicle20/atlas-model/model"
	"github.com/segmentio/kafka-go"
)

func spawnEventProvider(m Model, tm *temporalData) model.Provider[[]kafka.Message] {
	key := producer.CreateKey(int(m.OwnerId()))
	value := &statusEvent[spawnedStatusEventBody]{
		PetId:   m.Id(),
		OwnerId: m.OwnerId(),
		Type:    StatusEventTypeSpawned,
		Body: spawnedStatusEventBody{
			TemplateId: m.TemplateId(),
			Name:       m.Name(),
			Slot:       m.Slot(),
			Level:      m.Level(),
			Closeness:  m.Closeness(),
			Fullness:   m.Fullness(),
			X:          tm.X(),
			Y:          tm.Y(),
			Stance:     tm.Stance(),
			FH:         tm.FH(),
		},
	}
	return producer.SingleMessageProvider(key, value)
}

func despawnEventProvider(m Model, oldSlot int8, reason string) model.Provider[[]kafka.Message] {
	key := producer.CreateKey(int(m.OwnerId()))
	value := &statusEvent[despawnedStatusEventBody]{
		PetId:   m.Id(),
		OwnerId: m.OwnerId(),
		Type:    StatusEventTypeDespawned,
		Body: despawnedStatusEventBody{
			TemplateId: m.TemplateId(),
			Name:       m.Name(),
			Slot:       m.Slot(),
			Level:      m.Level(),
			Closeness:  m.Closeness(),
			Fullness:   m.Fullness(),
			OldSlot:    oldSlot,
			Reason:     reason,
		},
	}
	return producer.SingleMessageProvider(key, value)
}

func commandResponseEventProvider(m Model, commandId byte, success bool) model.Provider[[]kafka.Message] {
	key := producer.CreateKey(int(m.OwnerId()))
	value := &statusEvent[commandResponseStatusEventBody]{
		PetId:   m.Id(),
		OwnerId: m.OwnerId(),
		Type:    StatusEventTypeCommandResponse,
		Body: commandResponseStatusEventBody{
			Slot:      m.Slot(),
			Closeness: m.Closeness(),
			Fullness:  m.Fullness(),
			CommandId: commandId,
			Success:   success,
		},
	}
	return producer.SingleMessageProvider(key, value)
}

func closenessChangedEventProvider(m Model, amount int16) model.Provider[[]kafka.Message] {
	key := producer.CreateKey(int(m.OwnerId()))
	value := &statusEvent[closenessChangedStatusEventBody]{
		PetId:   m.Id(),
		OwnerId: m.OwnerId(),
		Type:    StatusEventTypeClosenessChanged,
		Body: closenessChangedStatusEventBody{
			Slot:      m.Slot(),
			Closeness: m.Closeness(),
			Amount:    amount,
		},
	}
	return producer.SingleMessageProvider(key, value)
}

func fullnessChangedEventProvider(m Model, amount int8) model.Provider[[]kafka.Message] {
	key := producer.CreateKey(int(m.OwnerId()))
	value := &statusEvent[fullnessChangedStatusEventBody]{
		PetId:   m.Id(),
		OwnerId: m.OwnerId(),
		Type:    StatusEventTypeFullnessChanged,
		Body: fullnessChangedStatusEventBody{
			Slot:     m.Slot(),
			Fullness: m.Fullness(),
			Amount:   amount,
		},
	}
	return producer.SingleMessageProvider(key, value)
}

func levelChangedEventProvider(m Model, amount int8) model.Provider[[]kafka.Message] {
	key := producer.CreateKey(int(m.OwnerId()))
	value := &statusEvent[levelChangedStatusEventBody]{
		PetId:   m.Id(),
		OwnerId: m.OwnerId(),
		Type:    StatusEventTypeLevelChanged,
		Body: levelChangedStatusEventBody{
			Slot:   m.Slot(),
			Level:  m.Level(),
			Amount: amount,
		},
	}
	return producer.SingleMessageProvider(key, value)
}

func slotChangedEventProvider(m Model, oldSlot int8) model.Provider[[]kafka.Message] {
	key := producer.CreateKey(int(m.OwnerId()))
	value := &statusEvent[slotChangedStatusEventBody]{
		PetId:   m.Id(),
		OwnerId: m.OwnerId(),
		Type:    StatusEventTypeSlotChanged,
		Body: slotChangedStatusEventBody{
			OldSlot: oldSlot,
			NewSlot: m.Slot(),
		},
	}
	return producer.SingleMessageProvider(key, value)
}

func excludeChangedEventProvider(m Model) model.Provider[[]kafka.Message] {
	items := make([]uint32, len(m.excludes))
	for i, e := range m.excludes {
		items[i] = e.ItemId()
	}

	key := producer.CreateKey(int(m.OwnerId()))
	value := &statusEvent[excludeChangedStatusEventBody]{
		PetId:   m.Id(),
		OwnerId: m.OwnerId(),
		Type:    StatusEventTypeExcludeChanged,
		Body: excludeChangedStatusEventBody{
			Items: items,
		},
	}
	return producer.SingleMessageProvider(key, value)
}
