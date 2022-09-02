package service

import (
	"fmt"
	"reflect"
	"time"

	"gitlab.com/ipron-cloud/ArchDB/schema/code"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/**
 * COMMON Struct
 */

type MediaSchedule struct {
	MediaType  code.MediaType      `bson:"mediaType"  json:"mediaType,omitempty"     ipron:"primary"`
	ScheduleId *primitive.ObjectID `bson:"scheduleId" json:"scheduleId,omitempty"`
	Enable     *bool               `bson:"enable"     json:"enable,omitempty"`
}

type SkillSchedule struct {
	SkillId    *primitive.ObjectID `bson:"skillId"    json:"skillId,omitempty"      ipron:"primary"`
	ScheduleId *primitive.ObjectID `bson:"scheduleId" json:"scheduleId,omitempty"`
	Enable     *bool               `bson:"enable"     json:"enable,omitempty"`
}

type SIPServerAuth struct {
	MD5Enable *bool   `bson:"enable"    json:"enable,omitempty"`
	MD5Id     *string `bson:"md5Id"     json:"md5Id,omitempty"`
	MD5Pwd    *string `bson:"md5Pwd"    json:"md5Pwd,omitempty"`
}

type SIPClientAuth struct {
	MD5Enable *bool   `bson:"md5Enable" json:"md5Enable,omitempty"`
	MD5Id     *string `bson:"md5Id"     json:"md5Id,omitempty"`
	MD5Pwd    *string `bson:"md5Pwd"    json:"md5Pwd,omitempty"`
}

/**
 * Account Table
 */

type Account struct {
	Id            *primitive.ObjectID   `bson:"_id"            json:"_id,omitempty"`
	Name          *string               `bson:"name"           json:"name,omitempty"`
	Alias         *string               `bson:"alias"          json:"alias,omitempty"`
	ServiceOption *AccountServiceOption `bson:"serviceOption"  json:"serviceOption,omitempty"`
	Enable        *bool                 `bson:"enable"         json:"enable,omitempty"`
	ExpireDate    *time.Time            `bson:"expireDate"     json:"expireDate,omitempty"`
	Timezone      *string               `bson:"timezone"       json:"timezone,omitempty"` // v0.9
	IsMaster      *bool                 `bson:"isMaster"       json:"isMaster,omitempty"` // v0.9
}

type AccountServiceOption struct {
	TTSEngine *string `bson:"ttsEngine"            json:"ttsEngine,omitempty"`
	//	TTSEngine            *primitive.ObjectID            `bson:"ttsEngine"            json:"ttsEngine,omitempty"`
	STTEngine            *primitive.ObjectID         `bson:"sttEngine"            json:"sttEngine,omitempty"`
	DefaultLang          *string                     `bson:"defaultLang"          json:"defaultLang,omitempty"`
	AllowFreeSeating     *bool                       `bson:"allowfreeSeating"     json:"allowfreeSeating,omitempty"`
	FullRecording        *bool                       `bson:"fullRecording"        json:"fullRecording,omitempty"`
	LimitConcurrentCalls *int                        `bson:"limitConcurrentCalls" json:"limitConcurrentCalls,omitempty"`
	LimitConcurrentUsers *int                        `bson:"limitConcurrentUsers" json:"limitConcurrentUsers,omitempty"` // v0.9
	LimitCPS             *int                        `bson:"limitCPS"             json:"limitCPS,omitempty"`
	RoutePolicy          code.RoutePolicy            `bson:"routePolicy"          json:"routePolicy,omitempty"`
	DayResetTime         *time.Time                  `bson:"dayResetTime"         json:"dayResetTime,omitempty"` // v0.9
	Record               *AccountServiceOptionRecord `bson:"record"               json:"record,omitempty"`       // v0.9
}

type AccountServiceOptionRecord struct {
	Type          *string  `bson:"type"            json:"type,omitempty"`
	VoiceDivision *bool    `bson:"voiceDivision"   json:"voiceDivision,omitempty"`
	Servers       []string `bson:"servers"         json:"servers,omitempty"`
}

/**
 * Flow Table
 */

type Flow struct {
	Id            *primitive.ObjectID `bson:"_id"                json:"_id,omitempty"`
	TntId         *primitive.ObjectID `bson:"tntId"              json:"tntId,omitempty"`
	Name          *string             `bson:"name"               json:"name,omitempty"`
	Kind          code.FlowKind       `bson:"kind"               json:"kind,omitempty"`
	Version       *string             `bson:"version"            json:"version,omitempty"`
	Type          *string             `bson:"type"               json:"type,omitempty"`
	Status        code.FlowStatus     `bson:"status"             json:"status,omitempty"`
	EditType      *string             `bson:"editType"           json:"editType,omitempty"` // v0.9
	EditLock      *bool               `bson:"editLock"           json:"editLock,omitempty"`
	EditUserId    *string             `bson:"editUserId"         json:"editUserId,omitempty"`
	Tags          []string            `bson:"tags,omitempty"     json:"tags,omitempty"`
	ServiceOption *FlowServiceOption  `bson:"serviceOption"      json:"serviceOption,omitempty"`
	Versions      []*FlowVersion      `bson:"versions,omitempty" json:"versions,omitempty"`
	Reservation   *FlowVerReservation `bson:"reservation"        json:"reservation,omitempty"`
}

type FlowServiceOption struct {
	Fdt int `bson:"fdt" json:"fdt,omitempty"`
	Idt int `bson:"idt" json:"idt,omitempty"`
}

type FlowVersion struct {
	Version      *string                `bson:"version"      json:"version"            ipron:"primary"` // key
	Status       code.FlowVersionStatus `bson:"status"       json:"status,omitempty"`
	LastEditTime *time.Time             `bson:"lastEditTime" json:"lastEditTime,omitempty"`
	Desc         *string                `bson:"desc"         json:"desc,omitempty"`
}

type FlowVerReservation struct {
	Enable     *bool                      `bson:"enable"        json:"enable,omitempty"`
	Version    *string                    `bson:"version"       json:"version,omitempty"`
	Status     code.FlowReservationStatus `bson:"status"        json:"status,omitempty"`
	TargetTime *time.Time                 `bson:"targetTime"    json:"targetTime,omitempty"`
}

/**
 * Prompt Table
 */

type Prompt struct {
	Id      *primitive.ObjectID `bson:"_id"        json:"_id,omitempty"`
	TntId   *primitive.ObjectID `bson:"tntId"      json:"tntId,omitempty"`
	Name    *string             `bson:"name"       json:"name,omitempty"`
	Kind    code.PromptKind     `bson:"kind"       json:"kind,omitempty"`
	TTSText *string             `bson:"ttsText"    json:"ttsText,omitempty"`
	Desc    *string             `bson:"desc"       json:"desc,omitempty"`
}

/**
 * FlowData Table
 */

type FlowData struct {
	Id        *primitive.ObjectID `bson:"_id"                json:"_id,omitempty"`
	TntId     *primitive.ObjectID `bson:"tntId"              json:"tntId,omitempty"`
	Name      *string             `bson:"name"               json:"name,omitempty"`
	TableName *string             `bson:"tableName"          json:"tableName,omitempty"`
	Columns   []*FlowDataColumn   `bson:"columns,omitempty"  json:"columns,omitempty"`
}

type FlowDataColumn struct {
	Name *string `bson:"name"              json:"name,omitempty"   ipron:"primary"`
	Type *string `bson:"type"              json:"type,omitempty"`
	Size *int    `bson:"size"              json:"size,omitempty"`
}

/**
 * Group Table
 */

type Group struct {
	Id            *primitive.ObjectID `bson:"_id"                     json:"_id,omitempty"`
	TntId         *primitive.ObjectID `bson:"tntId"                   json:"tntId,omitempty"`
	Name          *string             `bson:"name"                    json:"name,omitempty"`
	Tags          []string            `bson:"tags,omitempty"          json:"tags,omitempty"`
	DidNum        *string             `bson:"didNum"                  json:"didNum,omitempty"`
	GroupPath     []string            `bson:"groupPath,omitempty"     json:"groupPath,omitempty"`
	UserAutoRules []*CustAttribute    `bson:"userAutoRules,omitempty" json:"userAutoRules,omitempty"`
	Schedule      *GroupSchedule      `bson:"schedule"                json:"schedule,omitempty"`
	Depth         *int                `bson:"depth"                   json:"depth,omitempty"`
	ParentId      *primitive.ObjectID `bson:"parentId"                json:"parentId,omitempty"`
	Enable        *bool               `bson:"enable"                  json:"enable,omitempty"` // v0.9
}

type CustAttribute struct {
	AttrName *string `bson:"attrName" json:"attrName,omitempty"  ipron:"primary"` // key
	AttrRule *string `bson:"attrRule" json:"attrRule,omitempty"`
}

type GroupSchedule struct {
	Media []*MediaSchedule `bson:"media,omitempty" json:"media,omitempty"`
	Skill []*SkillSchedule `bson:"skill,omitempty" json:"skill,omitempty"`
}

/**
 * Group Table
 */
// index add PhoneId
type Users struct {
	Id             *primitive.ObjectID `bson:"_id"                     json:"_id,omitempty"`
	TntId          *primitive.ObjectID `bson:"tntId"                   json:"tntId,omitempty"`
	Email          *string             `bson:"email"                   json:"email,omitempty"`
	Password       *string             `bson:"password"                json:"password,omitempty"`
	Name           *string             `bson:"name"                    json:"name,omitempty"`
	Tags           []string            `bson:"tags,omitempty"          json:"tags,omitempty"`
	AuthLevel      code.UserAuthLevel  `bson:"authLevel"               json:"authLevel,omitempty"`
	AccessAuthId   *string             `bson:"accessAuthId"            json:"accessAuthId,omitempty"` // v0.9
	GroupId        *primitive.ObjectID `bson:"groupId"                 json:"groupId,omitempty"`
	Extension      *string             `bson:"extension"               json:"extension,omitempty"`
	DidNum         *string             `bson:"didNum"                  json:"didNum,omitempty"`
	BillNum        *string             `bson:"billNum"                 json:"billNum,omitempty"` // v0.9
	PhoneId        *primitive.ObjectID `bson:"phoneId"                 json:"phoneId,omitempty"`
	DefaultSkillId *primitive.ObjectID `bson:"defaultSkillId"          json:"defaultSkillId,omitempty"`
	//  Contact        []Contact
	SkillSets     []*UserSkill       `bson:"skillSets,omitempty"     json:"skillSets,omitempty"`
	Schedule      *UserSchedule      `bson:"schedule"                json:"schedule,omitempty"`
	MediaOptions  []*UserMediaOption `bson:"mediaOptions,omitempty"  json:"mediaOptions,omitempty"`
	ServiceOption *UserServiceOption `bson:"serviceOption"           json:"serviceOption,omitempty"`
	ExpireDate    *time.Time         `bson:"expireDate"              json:"expireDate,omitempty"`
	LastLoginDate *time.Time         `bson:"lastLoginDate"           json:"lastLoginDate,omitempty"`
	Enable        *bool              `bson:"enable"                  json:"enable,omitempty"`
}

// type Contact struct {
//
// }

type UserSkill struct {
	SkillId  *primitive.ObjectID `bson:"skillId"  json:"skillId,omitempty"    ipron:"primary"`
	SkillLv  *int                `bson:"skillLv"  json:"skillLv,omitempty"`
	SkillPri *int                `bson:"skillPri" json:"skillPri,omitempty"`
	Enable   *bool               `bson:"enable"   json:"enable,omitempty"` // v0.9
}

type UserSchedule struct {
	Enable *bool                 `bson:"enable"           json:"enable,omitempty"`
	Type   code.UserScheduleType `bson:"type"             json:"type,omitempty"`
	Medias []*MediaSchedule      `bson:"medias,omitempty" json:"medias,omitempty"`
	Skills []*SkillSchedule      `bson:"skills,omitempty" json:"skills,omitempty"`
}

type UserMediaOption struct {
	MediaType    code.MediaType  `bson:"mediaType"    json:"mediaType,omitempty"       ipron:"primary"`
	PriorityRate *int            `bson:"priorityRate" json:"priorityRate,omitempty"`
	Concurrent   *int            `bson:"concurrent"   json:"concurrent,omitempty"`
	ACWTimeSec   *int            `bson:"acwTimeSec"   json:"acwTimeSec,omitempty"`
	AutoAnswer   *UserAutoAnswer `bson:"autoAnswer"   json:"autoAnswer,omitempty"`
	Enable       *bool           `bson:"enable"       json:"enable,omitempty"`
}

type UserAutoAnswer struct {
	Enable   *bool `bson:"enable"   json:"enable"`
	DelaySec *int  `bson:"delaySec" json:"delaySec"`
}

type UserServiceOption struct {
	ServiceTemplateId *string             `bson:"serviceTemplateId" json:"serviceTemplateId,omitempty"`
	OutgoingBlock     *bool               `bson:"outgoingBlock"     json:"outgoingBlock,omitempty"`
	IncomingBlock     *bool               `bson:"incomingBlock"     json:"incomingBlock,omitempty"`
	ForwardUse        *bool               `bson:"forwardUse"        json:"forwardUse,omitempty"` // v0.9
	Forward           *UserServiceForward `bson:"forward"           json:"forward,omitempty,omitempty"`
	ReleaseToneUse    *bool               `bson:"releaseToneUse"    json:"releaseToneUse,omitempty"`
	TransferedToneUse *bool               `bson:"transferedToneUse" json:"transferedToneUse,omitempty"`
	EnableCallWait    *bool               `bson:"enableCallWait"    json:"enableCallWait,omitempty"`
	Auth              *UserServiceAuth    `bson:"auth"              json:"auth,omitempty"`
}

type UserServiceForward struct {
	Type        code.SvcForwardType `bson:"type"        json:"type,omitempty"`
	NoAnswerSec *int                `bson:"noAnswerSec" json:"noAnswerSec,omitempty"`
	TransferNum *string             `bson:"transferNum" json:"transferNum,omitempty"`
}

type UserServiceAuth struct {
	Monitor      *bool `bson:"monitor"      json:"monitor,omitempty"`
	Coaching     *bool `bson:"coaching"     json:"coaching,omitempty"`
	AvoidMonitor *bool `bson:"avoidMonitor" json:"avoidMonitor,omitempty"`
}

/**
 * Skills Table
 */

type Skills struct {
	Id     *primitive.ObjectID `bson:"_id"      json:"_id,omitempty"`
	TntId  *primitive.ObjectID `bson:"tntId"    json:"tntId,omitempty"`
	Name   *string             `bson:"name"     json:"name,omitempty"`
	Desc   *string             `bson:"desc"     json:"desc,omitempty"`
	Enable *bool               `bson:"enable"   json:"enable,omitempty"` // v0.9
}

/**
 * Queue Table
 */

type Queue struct {
	Id                *primitive.ObjectID    `bson:"_id"                         json:"_id,omitempty"`
	TntId             *primitive.ObjectID    `bson:"tntId"                       json:"tntId,omitempty"`
	Name              *string                `bson:"name"                        json:"name,omitempty"`
	Desc              *string                `bson:"desc"                        json:"desc,omitempty"`
	Tags              []string               `bson:"tags,omitempty"              json:"tags,omitempty"`
	RouteKind         code.QueueRouteKind    `bson:"routeKind"                   json:"routeKind,omitempty"`
	ExpandRouteOpts   []*QueueExpandRouteOpt `bson:"expandRouteOpts,omitempty"   json:"expandRouteOpts,omitempty"`
	AgentChoiceMethod code.AgentChoiceMethod `bson:"agentChoiceMethod"           json:"agentChoiceMethod,omitempty"`
	RouteTarget       *QueueRouteTarget      `bson:"routeTarget"                 json:"routeTarget,omitempty"`
	Option            *QueueOption           `bson:"option"                      json:"option,omitempty"`
	ServiceLevel      *QueueServiceLevel     `bson:"serviceLevel"                json:"serviceLevel,omitempty"`
	Enable            *bool                  `bson:"enable"                      json:"enable,omitempty"` // v0.9
}

type QueueExpandRouteOpt struct {
	SkillId  *primitive.ObjectID `bson:"skillId"  json:"skillId,omitempty"   ipron:"primary"`
	DelaySec *int                `bson:"delaySec" json:"delaySec,omitempty"`
}

type QueueRouteTarget struct {
	TargetKind *string  `bson:"targetKind"           json:"targetKind,omitempty"`
	TargetList []string `bson:"targetList,omitempty" json:"targetList,omitempty"`
}

type QueueOption struct {
	NoAnswerSec        *int                `bson:"noAnswerSec"        json:"noAnswerSec,omitempty"`
	MaxWaitCalls       *int                `bson:"maxWaitCalls"       json:"maxWaitCalls,omitempty"`
	MaxWaitSec         *int                `bson:"maxWaitSec"         json:"maxWaitSec,omitempty"`
	MinAbandonSec      *int                `bson:"minAbandonSec"      json:"minAbandonSec,omitempty"`
	DefaultQueueFlowId *primitive.ObjectID `bson:"defaultQueueFlowId" json:"defaultQueueFlowId,omitempty"`
}

type QueueServiceLevel struct {
	StandardSec *int     `bson:"standardSec"           json:"standardSec,omitempty"`
	Enable      *bool    `bson:"enable"                json:"enable,omitempty"`
	GoalRate    *int     `bson:"goalRate"              json:"goalRate,omitempty"`
	ScheduleId  []string `bson:"scheduleId,omitempty"  json:"scheduleId,omitempty"`
}

/**
 * StatusCause Table
 */

type StatusCause struct {
	Id    *primitive.ObjectID  `bson:"_id"   json:"_id,omitempty"`
	TntId *primitive.ObjectID  `bson:"tntId" json:"tntId,omitempty"`
	Type  code.UserStateMaster `bson:"type"  json:"type,omitempty"`
	Code  *string              `bson:"code"  json:"code,omitempty"`
	Name  *string              `bson:"name"  json:"name,omitempty"`
	Desc  *string              `bson:"desc"  json:"desc,omitempty"`
}

/**
 * Site Table
 */

type Site struct {
	Id     *primitive.ObjectID `bson:"_id"      json:"_id,omitempty"`
	TntId  *primitive.ObjectID `bson:"tntId"    json:"tntId,omitempty"`
	Name   *string             `bson:"name"     json:"name,omitempty"`
	Desc   *string             `bson:"desc"     json:"desc,omitempty"`
	Domain *string             `bson:"domain"   json:"domain,omitempty"`
}

/**
 * Trunk Table
 */

type Trunks struct {
	Id             *primitive.ObjectID  `bson:"_id"                 json:"_id,omitempty"`
	TntId          *primitive.ObjectID  `bson:"tntId"               json:"tntId,omitempty"`
	Name           *string              `bson:"name"                json:"name,omitempty"`
	Enable         *bool                `bson:"enable"              json:"enable,omitempty"`
	SiteId         *primitive.ObjectID  `bson:"siteId"              json:"siteId,omitempty"`
	Location       *TrunkLocation       `bson:"location"            json:"location,omitempty"`
	Kind           code.TrunkKind       `bson:"kind"                json:"kind,omitempty"`
	Protocol       code.ProtocolKind    `bson:"protocol"            json:"protocol,omitempty"`
	Networks       []*TrunkNetwork      `bson:"networks,omitempty"  json:"networks,omitempty"`
	AliveCheck     *TrunkAliveCheck     `bson:"aliveCheck"          json:"aliveCheck,omitempty"`
	Auth           *TrunkAuth           `bson:"auth"                json:"auth,omitempty"`
	OutboundOption *TrunkOutboundOption `bson:"outboundOption"      json:"outboundOption,omitempty"`
}

type TrunkLocation struct {
	InService *bool               `bson:"inService" json:"inService,omitempty"`
	EdgeId    *primitive.ObjectID `bson:"edgeId"    json:"edgeId,omitempty"`
}

type TrunkNetwork struct {
	IpPort           *string              `bson:"ipPort"               json:"ipPort,omitempty"  ipron:"primary"`
	Enable           *bool                `bson:"enable"               json:"enable,omitempty"`
	Status           code.TrunkStatusKind `bson:"status"               json:"status,omitempty"`           // v0.9
	StatusChangeDate *time.Time           `bson:"statusChangeDate"     json:"statusChangeDate,omitempty"` // v0.9
}

type TrunkAliveCheck struct {
	Enable   *bool `bson:"enable"   json:"enable,omitempty"`
	Interval *int  `bson:"interval" json:"interval,omitempty"`
	Retry    *int  `bson:"retry"    json:"retry,omitempty"`
}

type TrunkAuth struct {
	MD5Enable *bool   `bson:"md5Enable"          json:"md5Enable,omitempty"`
	Realm     *string `bson:"realm"              json:"realm,omitempty"`
	MD5Id     *string `bson:"md5Id"              json:"md5Id,omitempty"`
	MD5Pwd    *string `bson:"md5Pwd"             json:"md5Pwd,omitempty"`
}

type TrunkOutboundOption struct {
	ANINum              *string `bson:"aniNum"              json:"aniNum,omitempty"`
	ANIName             *string `bson:"aniName"             json:"aniName,omitempty"`
	AlwaysANINumEnable  *bool   `bson:"alwaysAniNumEnable"  json:"alwaysAniNumEnable,omitempty"`
	AlwaysANINameEnable *bool   `bson:"alwaysAniNameEnable" json:"alwaysAniNameEnable,omitempty"`
}

/**
 * Phone Table
 */

// index add MacAddr
type Phones struct {
	Id             *primitive.ObjectID `bson:"_id"                 json:"_id,omitempty"`
	TntId          *primitive.ObjectID `bson:"tntId"               json:"tntId,omitempty"`
	Name           *string             `bson:"name"                json:"name,omitempty"`
	ModelId        *string             `bson:"modelId"             json:"modelId,omitempty"`
	SiteId         *primitive.ObjectID `bson:"siteId"              json:"siteId,omitempty"`
	MacAddr        *string             `bson:"macAddr"             json:"macAddr,omitempty"`
	Codecs         []string            `bson:"codecs,omitempty"    json:"codecs,omitempty"`
	Network        *PhoneNetwork       `bson:"network"             json:"network,omitempty"`
	Auth           *SIPServerAuth      `bson:"auth"                json:"auth,omitempty"`
	Standalone     *PhoneStandalone    `bson:"standalone"          json:"standalone,omitempty"`
	LineKeys       []*PhoneLineKey     `bson:"lineKeys,omitempty"  json:"lineKeys,omitempty"`
	ProvInfo       *PhoneProvInfo      `bson:"provInfo"            json:"provInfo,omitempty"`
	AssignedUserId *primitive.ObjectID `bson:"assignedUserId"      json:"assignedUserId,omitempty"`
}

type PhoneNetwork struct {
	Protocol   code.ProtocolKind `bson:"protocol"   json:"protocol,omitempty"`
	Port       *int              `bson:"port"       json:"port,omitempty"`
	RegiExpire *int              `bson:"regiExpire" json:"regiExpire,omitempty"`
}

type PhoneStandalone struct {
	Enable  *bool   `bson:"enable"     json:"enable,omitempty"`
	DIDNum  *string `bson:"didNum"     json:"didNum,omitempty"`
	BillNum *string `bson:"billNum"    json:"billNum,omitempty"` // v0.9
	ExtNum  *string `bson:"extNum"     json:"extNum,omitempty"`
}

type PhoneLineKey struct {
	Label    *string `bson:"label"    json:"label,omitempty"     ipron:"primary"`
	Function *string `bson:"function" json:"function,omitempty"`
	Value    *string `bson:"value"    json:"value,omitempty"`
}

type PhoneProvInfo struct {
	FirmwareVer    *string             `bson:"firmwareVer"           json:"firmwareVer,omitempty"`
	UpdateUse      *bool               `bson:"updateUse"             json:"updateUse,omitempty"`
	FirmwareId     *primitive.ObjectID `bson:"firmwareId,omitempty"  json:"firmwareId,omitempty"`
	LastUpdateTime *time.Time          `bson:"lastUpdateDate"        json:"lastUpdateDate,omitempty"`
}

/**
 * Certification Table
 */

type Certification struct {
	Id    *primitive.ObjectID `bson:"_id"   json:"_id,omitempty"`
	TntId *primitive.ObjectID `bson:"tntId" json:"tntId,omitempty"`
	Name  *string             `bson:"name"  json:"name,omitempty"`
	Key   *string             `bson:"key"   json:"key,omitempty"`
}

/**
 * DidPlan Table
 */

type DidPlan struct {
	Id          *primitive.ObjectID `bson:"_id"          json:"_id,omitempty"`
	TntId       *primitive.ObjectID `bson:"tntId"        json:"tntId,omitempty"`
	Name        *string             `bson:"name"         json:"name,omitempty"`
	ANIPattern  *string             `bson:"aniPattern"   json:"aniPattern,omitempty"`
	DNISPattern *string             `bson:"dnisPattern"  json:"dnisPattern,omitempty"`
	Desc        *string             `bson:"desc"         json:"desc,omitempty"`
	RouteKind   code.DidRouteKind   `bson:"routeKind"    json:"routeKind,omitempty"`
	RoutePoint  *string             `bson:"routePoint"   json:"routePoint,omitempty"`
}

/**
 * DidPorts Table
 */

type DidPorts struct {
	Id               *primitive.ObjectID    `bson:"_id"                  json:"_id,omitempty"`
	TntId            *primitive.ObjectID    `bson:"tntId"                json:"tntId,omitempty"`
	DNIS             *string                `bson:"dnis"                 json:"dnis,omitempty"`
	CarrierKind      code.CarrierKind       `bson:"carrierKind"          json:"carrierKind,omitempty"`
	TrunkId          *primitive.ObjectID    `bson:"trunkId"              json:"trunkId,omitempty"`
	Desc             *string                `bson:"desc"                 json:"desc,omitempty"`
	Auth             *SIPCarrierAuth        `bson:"auth"                 json:"auth,omitempty"`
	NetOrder         *int                   `bson:"netOrder"             json:"netOrder,omitempty"`         // v0.9
	Status           code.DidPortStatusKind `bson:"status"               json:"status,omitempty"`           // v0.9
	StatusChangeDate *time.Time             `bson:"statusChangeDate"     json:"statusChangeDate,omitempty"` // v0.9
}

type SIPCarrierAuth struct {
	Enable    *bool   `bson:"enable"           json:"enable,omitempty"` // v0.9
	MD5Id     *string `bson:"md5Id"            json:"md5Id,omitempty"`
	MD5Pwd    *string `bson:"md5Pwd"           json:"md5Pwd,omitempty"`
	OrderType *bool   `bson:"orderType"        json:"orderType,omitempty"` // v0.9
}

/**
 * CarrierRoute Table
 */

type CarrierRoute struct {
	Id          *primitive.ObjectID `bson:"_id"          json:"_id,omitempty"`
	TntId       *primitive.ObjectID `bson:"tntId"        json:"tntId,omitempty"`
	SiteId      *primitive.ObjectID `bson:"siteId"       json:"siteId,omitempty"`
	Name        *string             `bson:"name"         json:"name,omitempty"`
	CarrierKind code.CarrierKind    `bson:"carrierKind"  json:"carrierKind,omitempty"`
	Trunks      []string            `bson:"trunks"       json:"trunks,omitempty"`
	Auth        *CarrierAuth        `bson:"auth"         json:"auth,omitempty"`
}

type CarrierAuth struct {
	Enable       *bool   `bson:"enable"           json:"enable,omitempty"`
	AuthType     *string `bson:"carrierAuthType"  json:"carrierAuthType,omitempty"`
	AuthInterval *int    `bson:"regiInterval"     json:"regiInterval,omitempty"`
	AuthRetry    *int    `bson:"regiRetry"        json:"regiRetry,omitempty"`
}

/**
 * NumberPlan Table
 */

type NumberPlan struct {
	Id       *primitive.ObjectID     `bson:"_id"                    json:"_id,omitempty"`
	TntId    *primitive.ObjectID     `bson:"tntId"                  json:"tntId,omitempty"`
	SiteId   *primitive.ObjectID     `bson:"siteId"                 json:"siteId,omitempty"`
	Name     *string                 `bson:"name"                   json:"name,omitempty"`
	Type     code.NumPlanType        `bson:"type"                   json:"type,omitempty"`
	Number   *DialPlanNumberInfo     `bson:"number"                 json:"number,omitempty"`
	EditRule *DialPlanNumberEditRule `bson:"editRule"               json:"editRule,omitempty"`
	Target   *DialPlanTarget         `bson:"target"                 json:"target,omitempty"`
}

type DialPlanNumberInfo struct {
	Pattern *string `bson:"pattern" json:"pattern,omitempty"`
	Prefix  *string `bson:"prefix"  json:"prefix,omitempty"`
	MinLen  *int    `bson:"minLen"  json:"minLen,omitempty"`
	MaxLen  *int    `bson:"maxLen"  json:"maxLen,omitempty"`
}

type DialPlanNumberEditRule struct {
	KeepEnable     *bool   `bson:"keepEnable"       json:"keepEnable,omitempty"`
	KeepLen        *int    `bson:"keepLen"          json:"keepLen,omitempty"`
	PrefixDelLen   *int    `bson:"prefixDelLen"     json:"prefixDelLen,omitempty"`
	PrefixAddDigit *string `bson:"prefixAddDigit"   json:"prefixAddDigit,omitempty"`
	SuffixDelLen   *int    `bson:"suffixDelLen"     json:"suffixDelLen,omitempty"`
	SuffixAddDigit *string `bson:"suffixAddDigit"   json:"suffixAddDigit,omitempty"`
}

type DialPlanTarget struct {
	Classification code.NumPlanClassification `bson:"classification"  json:"classification,omitempty"`
	Point          *string                    `bson:"point"           json:"point,omitempty"`
}

type RoutePoints struct {
	Id          *primitive.ObjectID `bson:"_id"               json:"_id,omitempty"`
	TntId       *primitive.ObjectID `bson:"tntId"             json:"tntId,omitempty"`
	SiteId      *primitive.ObjectID `bson:"siteId"            json:"siteId,omitempty"`
	Name        *string             `bson:"name"              json:"name,omitempty"`
	RouteMethod code.RouteMethod    `bson:"routeMethod"       json:"routeMethod,omitempty"`
	Trunks      []string            `bson:"trunks,omitempty"  json:"trunks,omitempty"`
}

type FeatureCodes struct {
	Id     *primitive.ObjectID `bson:"_id"     json:"_id,omitempty"`
	TntId  *primitive.ObjectID `bson:"tntId"   json:"tntId,omitempty"`
	SiteId *primitive.ObjectID `bson:"siteId"  json:"siteId,omitempty"`
	Code   *string             `bson:"code"    json:"code,omitempty"`
	Prefix *string             `bson:"prefix"  json:"prefix,omitempty"`
	MinLen *int                `bson:"minLen"  json:"minLen,omitempty"`
	MaxLen *int                `bson:"maxLen"  json:"maxLen,omitempty"`
}

/**
 * Spam Table
 */

type Spam struct {
	Id     *primitive.ObjectID `bson:"_id"     json:"_id,omitempty"`
	TntId  *primitive.ObjectID `bson:"tntId"   json:"tntId,omitempty"`
	ANI    *string             `bson:"ani"     json:"ani,omitempty"`
	Desc   *string             `bson:"desc"    json:"desc,omitempty"`
	Enable *bool               `bson:"enable"  json:"enable,omitempty"`
}

/**
 * Schedule Table
 */

type Schedule struct {
	Id    *primitive.ObjectID `bson:"_id"              json:"_id,omitempty"`
	TntId *primitive.ObjectID `bson:"tntId"            json:"tntId,omitempty"`
	Name  *string             `bson:"name"             json:"name,omitempty"`
	Rules []*ScheduleRule     `bson:"rules,omitempty"  json:"rules,omitempty"`
}

type ScheduleRule struct {
	Name      *string    `bson:"name"             json:"name,omitempty"          ipron:"primary"`
	StartDate *string    `bson:"startDate"        json:"startDate,omitempty"`
	EndDate   *string    `bson:"endDate"          json:"endDate,omitempty"`
	StartTime *string    `bson:"startTime"        json:"startTime,omitempty"`
	EndTime   *string    `bson:"endTime"          json:"endTime,omitempty"`
	Week      []bool     `bson:"week,omitempty"   json:"week,omitempty"`
	Enable    *bool      `bson:"enable"           json:"enable,omitempty"`
}

/**
 * Bridge Table
 */

type Bridge struct {
	Id        *primitive.ObjectID   `bson:"_id"                  json:"_id,omitempty"`
	TntId     *primitive.ObjectID   `bson:"tntId"                json:"tntId,omitempty"`
	Name      *string               `bson:"name"                 json:"name,omitempty"`
	Type      code.BridgeType       `bson:"type"                 json:"type,omitempty"` // v0.9
	BaseURL   *string               `bson:"baseURL"              json:"baseURL,omitempty"`
	AuthType  code.BridgeAuthType   `bson:"authType"             json:"authType,omitempty"` // v0.9
	AuthToken *string               `bson:"authToken"            json:"authToken,omitempty"`
	Property  []*BridgePropertyItem `bson:"property"             json:"property,omitempty"` // v0.9
}

type BridgePropertyItem struct {
	Name  *string `bson:"name"   json:"name,omitempty"     ipron:"primary"` // v0.9
	Value *string `bson:"value"  json:"value,omitempty"`                    // v0.9
}

/**
 * Bridge Table
 */

type Tags struct {
	Id    *primitive.ObjectID `bson:"_id"      json:"_id,omitempty"`
	TntId *primitive.ObjectID `bson:"tntId"    json:"tntId,omitempty"`
	Name  *string             `bson:"name"     json:"name,omitempty"`
}

/**
 * Firmeare
 */

type Firmware struct {
	Id          *primitive.ObjectID `bson:"_id"                   json:"_id,omitempty"`
	Name        *string             `bson:"name"                  json:"name,omitempty"`
	ModelId     *string             `bson:"modelId"               json:"modelId,omitempty"`
	Version     *string             `bson:"version"               json:"version,omitempty"`
	VersionFile *string             `bson:"versionFile"           json:"versionFile,omitempty"`
	Files       []string            `bson:"files,omitempty"       json:"files,omitempty"`
	UploadDate  *time.Time          `bson:"uploadDate"            json:"uploadDate,omitempty"`
	Desc        *string             `bson:"desc"                  json:"desc,omitempty"`
}

/**
 * Holiday   v0.9
 */

type Holiday struct {
	Id    *primitive.ObjectID `bson:"_id"                   json:"_id,omitempty"`
	TntId *primitive.ObjectID `bson:"tntId"                 json:"tntId,omitempty"`
	Name  *string             `bson:"name"                  json:"name,omitempty"`
	Type      code.HolidayType       `bson:"type"                  json:"type,omitempty"`
	RepeatOpt code.HolidayRepeatType `bson:"repeatOpt"             json:"repeatOpt,omitempty"`
	IsLunar   *bool                  `bson:"isLunar"               json:"isLunar,omitempty"`
	StartDate *time.Time             `bson:"startDate"             json:"startDate,omitempty"`
	EndDate   *time.Time             `bson:"endDate"               json:"endDate,omitempty"`
	Enable    *bool                  `bson:"enable"                json:"enable,omitempty"`
	Desc      *string                `bson:"desc"                  json:"desc,omitempty"`
}

/**
 * API Route
 */

type ApiRoute struct {
	Id         *primitive.ObjectID  `bson:"_id"                      json:"_id"`
	RouteId    *string              `bson:"routeId"                  json:"routeId"`
	Orders     *int                 `bson:"orders"                   json:"orders"`
	Uri        *string              `bson:"uri"                      json:"uri"`
	Predicates []*ApiRoutePredicate `bson:"predicates,omitempty"     json:"predicates"`
	Filters    []*ApiRouteFilter    `bson:"filters,omitempty"        json:"filters"`
	Desc       *string              `bson:"desc"                     json:"desc"`
	EnableYn   *bool                `bson:"enableYn"                 json:"enableYn"`
}

type ApiRoutePredicate struct {
	Name *string               `bson:"name"                  json:"name"`
	Args *ApiRoutePredicateArg `bson:"args,omitempty"        json:"args"`
}

type ApiRoutePredicateArg struct {
	Pattern *string `bson:"pattern"        json:"pattern"`
	Method  *string `bson:"method"         json:"method"`
}

type ApiRouteFilter struct {
	Name *string            `bson:"name"           json:"name"`
	Args *ApiRouteFilterArg `bson:"args,omitempty" json:"args"`
}

type ApiRouteFilterArg struct {
	Regexp      *string `bson:"regexp"        json:"regexp"`
	Replacement *string `bson:"replacement"   json:"replacement"`
}

/**
 * Grant
 */

type Grant struct {
	Id        *primitive.ObjectID `bson:"_id"                     json:"_id,omitempty"`
	Pattern   *string             `bson:"pattern"                 json:"pattern,omitempty"`
	Method    []string            `bson:"method,omitempty"        json:"method,omitempty"`
	GrantRole []string            `bson:"grantRole,omitempty"     json:"grantRole,omitempty"`
}

/**
 * Policy
 */

type Policy struct {
	Id       *primitive.ObjectID `bson:"_id"                    json:"_id,omitempty"`
	TntId    *primitive.ObjectID `bson:"tntId"                  json:"tntId"`
	Password []*PolicyPassword   `bson:"password"               json:"password"`
	Login    *PolicyLogin        `bson:"login"                  json:"login"`
}

type PolicyPassword struct {
	InvalidChangeDate  *int     `bson:"invalidChangeDate"                  json:"invalidChangeDate"`
	InvalidTextTypeTag []string `bson:"invalidTextTypeTag,omitempty"       json:"invalidTextTypeTag,omitempty"`
	MinLength          *int     `bson:"minLength"                          json:"minLength"`
	MaxLength          *int     `bson:"maxLength"                          json:"maxLength"`
}

type PolicyLogin struct {
	NoLoginExpireDate *int `bson:"noLoginExpireDate"                 json:"noLoginExpireDate"`
}

/******
 *
 *  Marshal define
 *
 **/

func (s *Account) Marshal() ([]byte, error) {
	return bson.Marshal(s)
}

func (s *Flow) Marshal() ([]byte, error) {
	return bson.Marshal(s)
}

func (s *Prompt) Marshal() ([]byte, error) {
	return bson.Marshal(s)
}

func (s *Group) Marshal() ([]byte, error) {
	return bson.Marshal(s)
}

func (s *Users) Marshal() ([]byte, error) {
	return bson.Marshal(s)
}

func (s *Skills) Marshal() ([]byte, error) {
	return bson.Marshal(s)
}

func (s *Queue) Marshal() ([]byte, error) {
	return bson.Marshal(s)
}

func (s *Site) Marshal() ([]byte, error) {
	return bson.Marshal(s)
}

func (s *Trunks) Marshal() ([]byte, error) {
	return bson.Marshal(s)
}

func (s *Phones) Marshal() ([]byte, error) {
	return bson.Marshal(s)
}

func (s *NumberPlan) Marshal() ([]byte, error) {
	return bson.Marshal(s)
}

func (s *RoutePoints) Marshal() ([]byte, error) {
	return bson.Marshal(s)
}

func (s *FeatureCodes) Marshal() ([]byte, error) {
	return bson.Marshal(s)
}

func (s *DidPlan) Marshal() ([]byte, error) {
	return bson.Marshal(s)
}

func (s *DidPorts) Marshal() ([]byte, error) {
	return bson.Marshal(s)
}

func (s *CarrierRoute) Marshal() ([]byte, error) {
	return bson.Marshal(s)
}

func (s *Spam) Marshal() ([]byte, error) {
	return bson.Marshal(s)
}

func (s *Bridge) Marshal() ([]byte, error) {
	return bson.Marshal(s)
}

func (s *Tags) Marshal() ([]byte, error) {
	return bson.Marshal(s)
}

func (s *ApiRoute) Marshal() ([]byte, error) {
	return bson.Marshal(s)
}

func (s *Grant) Marshal() ([]byte, error) {
	return bson.Marshal(s)
}

func (s *Policy) Marshal() ([]byte, error) {
	return bson.Marshal(s)
}

func T2String(inf interface{}) string {
	var output string
	value := reflect.ValueOf(inf)
	item := value.Type()

	switch item.Kind() {
	case reflect.Ptr:
		if value.IsZero() == false {
			output = T2String(value.Elem().Interface())
		}
	case reflect.Struct:
		switch item.Name() {
		case "Time":
			output = fmt.Sprintf("%v", inf)
		default:
			output = "{"
			for i := 0; i < item.NumField(); i++ {
				sub := value.Type().Field(i)
				// log.Println(sub.Name)
				if i > 0 {
					output = output + ", "
				}
				output = output + fmt.Sprintf("%s:%s", sub.Name, T2String(value.Field(i).Interface()))
			}
			output += "}"
		}
	case reflect.Slice:
		output = "["
		for i := 0; i < value.Len(); i++ {
			if i > 0 {
				output = output + ","
			}
			output = output + T2String(value.Index(i).Interface())
		}
		output += "]"
	case reflect.Map:
		output = "{"
		for i, e := range value.MapKeys() {
			sub := value.MapIndex(e)
			// log.Println(sub.Name)
			if i > 0 {
				output = output + ", "
			}
			output = output + fmt.Sprintf("%s:%s", e, T2String(sub.Interface()))
		}
		output += "}"
	case reflect.Array:
		switch item.Name() {
		case "ObjectID": 
			output = inf.(primitive.ObjectID).Hex()
		}
	default:
		output = fmt.Sprintf("%v", inf)
	}

	return output
}
