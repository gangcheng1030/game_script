package model

import (
	"time"
)

type EventStatus int

const (
	EVENT_STATUS_INIT = iota
	EVENT_STATUS_PROCESSING
)

type EventType int

const (
	EVENT_TYPE_COMMON = iota
	EVENT_TYPE_FIXEDID
	EVENT_TYPE_FIXEDGROUP
)

type Event struct {
	Id              int64
	Status          EventStatus
	Type            EventType
	NodeId          string
	FollowerNodeIds []string
	GroupId         string

	MeiRi   string
	Account Account
}

type Account struct {
	AccountName string
	Password    string
	Roles       []Role

	FollowerAddrs        []string
	FollowerAccountNames []string
	FollowerPasswords    []string
}

type Role struct {
	Id                   int
	FollowerRoleIds      []int
	FollowerJunTuanNames []string
	Fubens               []FuBen

	DisablePreClearBag bool
	PostClearBag       bool
}

type FuBen struct {
	Name       string
	Difficulty int
}

type NodeStatus int

const (
	NODE_STATUS_IDLE = iota
	NODE_STATUS_BUSY
)

type Node struct {
	Id              string
	Addr            string
	Status          NodeStatus
	LastUpdatedTime time.Time
	Config          NodeConfig
}

func FindEventIndex(events []*Event, eventId int64) int {
	for i, tEvent := range events {
		if tEvent.Id == eventId {
			return i
		}
	}

	return -1
}

type AccountConfig struct {
	AccountName string
	Password    string
	JunTuanName string
}

type NodeConfig struct {
	Group          string
	CanNotBeLeader bool
	FullScreenMode int
	DirName        string
}

func NewDefaultNodeConfig() *NodeConfig {
	return &NodeConfig{
		Group:          "1",
		CanNotBeLeader: false,
		FullScreenMode: 0,
		DirName:        "D:\\Netease\\CJDMJ",
	}
}

type UpdateNodeStatusReq struct {
	Group     string
	Id        string
	OldStatus NodeStatus
	NewStatus NodeStatus
}

type AccountStatus int

const (
	ACCOUNT_STATUS_IDLE = iota
	ACCOUNT_STATUS_BUSY
)

type AccountWithStatus struct {
	AccountName string
	Status      AccountStatus
}
