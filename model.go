package dynago

import (
	"strings"
	"time"
)

type Model struct {
	PK         string `json:"-"`
	SK         string `json:"-"`
	Timestamps `json:"inline"`
}

type Timestamps struct {
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

func (t *Timestamps) BeforeUpdate() {
	t.UpdatedAt = time.Now()
}

func (t *Timestamps) BeforeInsert() {
	t.CreatedAt = time.Now()
}

func (m Model) GetPK() string {
	return strings.Split(m.PK, "#")[1]
}

func (m Model) GetSk() string {
	return strings.Split(m.SK, "#")[1]
}

type IModel interface {
	GetPK() string
	SetPK(id string)
	SetSK(id string)
	GetSk() string
	PkPrefix() string
	SkPrefix() string
}
