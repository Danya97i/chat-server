package models

// Action is the type of action
type Action string

const (
	ActionCreate Action = "create" // ActionCreate - create chat
	ActionDelete Action = "delete" // ActionDelete - delete chat
)

// LogInfo is the log information
type LogInfo struct {
	ChatID int64
	Action Action
}
