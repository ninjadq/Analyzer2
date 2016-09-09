package models

import (
	"time"
)

// Event 是OpenFalcon的一次事件，可以是problem也可以是OK
type Event struct {
	ID           string `json:"id"`
	StrategyID   int    `json:"strategy"`
	ExpressionID int    `json:"expression"`
	TplID        int
	ActionID     int
	Status       string            `json:"status"` // OK or PROBLEM
	Endpoint     string            `json:"endpoint"`
	LeftValue    float64           `json:"leftValue"`
	CurrentStep  int               `json:"currentStep"`
	EventTime    int64             `json:"eventTime"`
	PushedTags   map[string]string `json:"pushedTags"`
	Priority     int
	Note         string
	Metric       string
	RightValue   string
	Operator     string
	Func         string
	MaxStep      int
	Counter      string
}

// CounterStates is stroe the current status of a counter
type CounterStatus struct {
	ID         int
	PK         string // md5(endpoint+metric+tags+strategy)
	Endpoint   string
	Metric     string
	Tags       string
	Strategy   int
	Expression int
	Status     string
	StartTime  time.Time
}

type GroupAlarmClock struct {
	LastAlarmTime time.Time

	CounterStatus int
	Group         int
}

// StateChangeHistory is histroy  of state changes
type CounterStatusHistory struct {
	ID            int
	CounterStatus int
	State         string
	StartTime     time.Time
	EndTime       time.Time
}

// SendingRecord is store the AlarmContent and related users
type AlarmReceiveRecord struct {
	ID         int
	State      string
	SendMethod int
	AlarmType  string //HOST, MODULE, CONVERGENCE
	Content    AlarmContent
	Users      User
	SendTime   time.Time
	ReadTime   time.Time
}

// AlarmContent used to store the content of a AlarmType
// TODO: This should be a interface and all types of alarm content(ep. Host, Module) can represented by it.
type AlarmContent struct {
	Title      string
	Note       string
	Endpoint   string
	Metric     string
	LeftValue  float64
	Func       string
	Operator   string
	RightValue string
	Tags       map[string]string
}

// User is user
type User struct {
	ID       int
	Username string
}
