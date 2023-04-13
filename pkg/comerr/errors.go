package comerr

import "errors"

var (
	// ErrUnauthorized TODO
	ErrUnauthorized = errors.New("user unauthorized")
	// ErrPermission TODO
	ErrPermission = errors.New("user without permission")
	// ErrNoRecord TODO
	ErrNoRecord = errors.New("no record")
	// ErrRPC TODO
	ErrRPC = errors.New("RPC call failed")
	// ErrDatabase TODO
	ErrDatabase = errors.New("database error")
)
