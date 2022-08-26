package blut

import (
	"context"
	"fmt"
)

type Namespace struct {
	engine *Engine
	Name   string   `json:"name"` // 命名空间的名称
	Schema []Schema `json:"schema"`
}

type SchemaType int8

const (
	SchemaString SchemaType = iota + 1
	SchemaNumber
	SchemaBool
	SchemaObject
	SchemaArray
)

var (
	defaultString = ""
	defaultNumber = 0
	defaultBool   = false
	defaultObject = struct{}{}
	defaultArray  = make([]interface{}, 0)
)

type Schema struct {
	Name        string      `json:"name"`    // 字段名
	Type        SchemaType  `json:"type"`    // 数据类型
	Default     interface{} `json:"default"` // 默认值
	Pk          bool        `json:"pk"`      // 是否是主键
	Index       bool        `json:"index"`   // 是否是索引字段.
	CanSetEmpty bool        `json:"can_set_empty"`
}

func (e *Engine) NewNamespace(name string) *Namespace {
	return &Namespace{
		engine: e,
		Name:   name,
		Schema: make([]Schema, 0, 8),
	}
}

func (n *Namespace) SetPK(name string, typ SchemaType) *Namespace {
	n.Schema = append(n.Schema, Schema{
		Name:        name,
		Type:        typ,
		Pk:          true,
		Index:       true,
		CanSetEmpty: false,
	})
	return n
}

func (n *Namespace) AddField(name string, typ SchemaType, index bool) *Namespace {
	n.Schema = append(n.Schema, Schema{
		Name:        name,
		Type:        typ,
		Default:     getDefault(typ),
		Pk:          false,
		Index:       index,
		CanSetEmpty: true,
	})
	return n
}

func getDefault(typ SchemaType) interface{} {
	switch typ {
	case SchemaBool:
		return defaultBool
	case SchemaArray:
		return defaultArray
	case SchemaNumber:
		return defaultNumber
	case SchemaObject:
		return defaultObject
	case SchemaString:
		return defaultString
	}
	panic(fmt.Sprintf("unsupport typ: %d", typ))
}

func (n *Namespace) validate() error {
	var (
		hasPk bool
	)
	for i, s := range n.Schema {
		if s.Name == "" {
			return fmt.Errorf("field name is empty: %d", i+1)
		}
		if s.Pk {
			hasPk = true
		}
	}
}

func (n *Namespace) Create(ctx context.Context) error {
	if err := n.validate(); err != nil {
		return err
	}

	return n.engine.createNs(ctx, n)
}
