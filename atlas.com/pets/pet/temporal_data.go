package pet

import "sync"

type temporalData struct {
	x      int16
	y      int16
	stance byte
	fh     int16
}

func (d *temporalData) UpdatePosition(x int16, y int16, fh int16) *temporalData {
	return &temporalData{
		x:      x,
		y:      y,
		stance: d.stance,
		fh:     fh,
	}
}

func (d *temporalData) Update(x int16, y int16, stance byte, fh int16) *temporalData {
	return &temporalData{
		x:      x,
		y:      y,
		stance: stance,
		fh:     fh,
	}
}

func (d *temporalData) UpdateStance(stance byte) *temporalData {
	return &temporalData{
		x:      d.x,
		y:      d.y,
		stance: stance,
		fh:     d.fh,
	}
}

func (d *temporalData) X() int16 {
	return d.x
}

func (d *temporalData) Y() int16 {
	return d.y
}

func (d *temporalData) Stance() byte {
	return d.stance
}

func (d *temporalData) FH() int16 {
	return d.fh
}

func NewTemporalData() *temporalData {
	return &temporalData{fh: 1}
}

type temporalRegistry struct {
	data     map[uint64]*temporalData
	mutex    *sync.RWMutex
	petLocks map[uint64]*sync.RWMutex
}

func (r *temporalRegistry) UpdatePosition(petId uint64, x int16, y int16, fh int16) {
	r.lockPet(petId)
	if val, ok := r.data[petId]; ok {
		r.data[petId] = val.UpdatePosition(x, y, fh)
	} else {
		r.data[petId] = NewTemporalData()
	}
	r.unlockPet(petId)
}

func (r *temporalRegistry) lockPet(petId uint64) {
	r.mutex.Lock()
	lock, exists := r.petLocks[petId]
	if !exists {
		lock = &sync.RWMutex{}
		r.petLocks[petId] = lock
	}
	r.mutex.Unlock()
	lock.Lock()
}

func (r *temporalRegistry) readLockPet(petId uint64) {
	r.mutex.Lock()
	lock, exists := r.petLocks[petId]
	if !exists {
		lock = &sync.RWMutex{}
		r.petLocks[petId] = lock
	}
	r.mutex.Unlock()
	lock.RLock()
}

func (r *temporalRegistry) unlockPet(petId uint64) {
	if val, ok := r.petLocks[petId]; ok {
		val.Unlock()
	}
}

func (r *temporalRegistry) readUnlockPet(petId uint64) {
	if val, ok := r.petLocks[petId]; ok {
		val.RUnlock()
	}
}

func (r *temporalRegistry) Update(petId uint64, x int16, y int16, stance byte, fh int16) {
	r.lockPet(petId)
	if val, ok := r.data[petId]; ok {
		r.data[petId] = val.Update(x, y, stance, fh)
	} else {
		r.data[petId] = NewTemporalData()
	}
	r.unlockPet(petId)
}

func (r *temporalRegistry) UpdateStance(petId uint64, stance byte) {
	r.lockPet(petId)
	if val, ok := r.data[petId]; ok {
		r.data[petId] = val.UpdateStance(stance)
	} else {
		r.data[petId] = NewTemporalData()
	}
	r.unlockPet(petId)
}

func (r *temporalRegistry) GetById(petId uint64) *temporalData {
	r.readLockPet(petId)
	defer r.readUnlockPet(petId)

	result := r.data[petId]
	if result != nil {
		return result
	}
	return NewTemporalData()
}

func (r *temporalRegistry) Remove(petId uint64) {
	r.lockPet(petId)
	defer r.unlockPet(petId)
	delete(r.data, petId)
}

var once sync.Once
var instance *temporalRegistry

func GetTemporalRegistry() *temporalRegistry {
	once.Do(func() {
		instance = &temporalRegistry{
			data:     make(map[uint64]*temporalData),
			mutex:    &sync.RWMutex{},
			petLocks: make(map[uint64]*sync.RWMutex),
		}
	})
	return instance
}
