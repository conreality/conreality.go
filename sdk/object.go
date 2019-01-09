/* This is free and unencumbered software released into the public domain. */

package sdk

import uuid "github.com/satori/go.uuid"

// ObjectID TODO...
type ObjectID = uint64

// Object TODO...
type Object struct {
	UUID uuid.UUID
}

// NewObject TODO...
func NewObject(objectUUID string) *Object {
	var uuid, err = uuid.FromString(objectUUID)
	if err != nil {
		return nil
	}
	return &Object{UUID: uuid}
}