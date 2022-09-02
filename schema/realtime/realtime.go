package realtime

import (
	"fmt"
	"reflect"
	"time"

	"gitlab.com/ipron-cloud/ArchDB/schema/code"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/**
 * Realtime Account Table
 */

type Account struct {
	Id                   *primitive.ObjectID  `bson:"_id"                  json:"id,omitempty"`
	ConcurrentLoginUsers *int                 `bson:"concurrentLoginUsers" json:"concurrentLoginUsers,omitempty"`
	MediaUsage           []*AccountMediaUsage `bson:"mediaUsage"           json:"mediaUsage,omitempty"`
}

// 매체별 사용량
type AccountMediaUsage struct {
	Media code.MediaType `bson:"media"     json:"media,omitempty"    ipron:"primary"`
	Count *int           `bson:"count"     json:"count,omitempty"`
}

/**
 * Realtime Users Table
 */

type User struct {
	Id             *primitive.ObjectID   `bson:"_id"              json:"id,omitempty"`
	TntId          *primitive.ObjectID   `bson:"tntId"            json:"tntId,omitempty"`
	CurrentState   []*UserState          `bson:"currentState"     json:"currentState,omitempty"`
	CurrentCall    []*UserCall           `bson:"currentCall"      json:"currentCall,omitempty"`
	RecallState    []*UserRecallState    `bson:"recallState"      json:"recallState,omitempty"`
	AfterState     []*UserAfterState     `bson:"afterState"       json:"afterState,omitempty"`
	ACDAggregate   []*UserACDAggregate   `bson:"acdAggregate"     json:"acdAggregate,omitempty"`
	MediaAggregate []*UserMediaAggregate `bson:"mediaAggregate"   json:"mediaAggregate,omitempty"`
}

type UserCall struct {
	CallId      *primitive.ObjectID `bson:"callId"           json:"callId,omitempty"     ipron:"primary"`
	QueueId     *primitive.ObjectID `bson:"queueId"          json:"queueId,omitempty"`
	Ani         *string             `bson:"ani"              json:"ani,omitempty"`
	Dnis        *string             `bson:"dnis"             json:"dnis,omitempty"`
	ConnectTime *time.Time          `bson:"connectTime"      json:"connectTime,omitempty"`
	MediaType   code.MediaType      `bson:"mediaType"        json:"mediaType,omitempty"`
}

type UserState struct {
	MediaType    code.MediaType       `bson:"mediaType"        json:"mediaType,omitempty"    ipron:"primary"`
	State        code.UserState       `bson:"state"            json:"state,omitempty"`
	AssignState  code.UserAssignState `bson:"assignState"      json:"assignState,omitempty"`
	Cause        *string              `bson:"cause"            json:"cause,omitempty"`
	StateChgTime *time.Time           `bson:"stateChgTime"     json:"stateChgTime,omitempty"`
}

type UserRecallState struct {
	MediaType code.MediaType `bson:"mediaType"        json:"mediaType,omitempty"    ipron:"primary"`
	State     code.UserState `bson:"state"            json:"state,omitempty"`
	Cause     *string        `bson:"cause"            json:"cause,omitempty"`
}

type UserAfterState struct {
	MediaType code.MediaType `bson:"mediaType"        json:"mediaType,omitempty"    ipron:"primary"`
	State     code.UserState `bson:"state"            json:"state,omitempty"`
	Cause     *string        `bson:"cause"            json:"cause,omitempty"`
}

type UserACDAggregate struct {
	QueueId         *primitive.ObjectID `bson:"queueId"          json:"queueId,omitempty"    ipron:"primary"`
	AcdConnTotCalls *int                `bson:"acdConnTotCalls"  json:"acdConnTotCalls,omitempty"`
	AcdConnTotSec   *int                `bson:"acdConnTotSec"    json:"acdConnTotSec,omitempty"`
}

type UserMediaAggregate struct {
	MediaType       code.MediaType `bson:"mediaType"        json:"mediaType,omitempty"    ipron:"primary"`
	AcdConnTotCalls *int           `bson:"acdConnTotCalls"  json:"acdConnTotCalls,omitempty"`
	AcdConnTotSec   *int           `bson:"acdConnTotSec"    json:"acdConnTotSec,omitempty"`
	ReadyTotSec     *int           `bson:"readyTotSec"      json:"readyTotSec,omitempty"`
	BusyCount       *int           `bson:"busyCount"        json:"busyCount,omitempty"`
	LastLoginTime   *time.Time     `bson:"lastLoginTime"    json:"lastLoginTime,omitempty"`
}

/**
 * Realtime Calls Table
 */

type Call struct {
	Id           *primitive.ObjectID `bson:"_id"                   json:"id,omitempty"`
	TntId        *primitive.ObjectID `bson:"tntId"                 json:"tntId,omitempty"`
	SiteId       *primitive.ObjectID `bson:"siteId"                json:"siteId,omitempty"`
	Ani          *string             `bson:"ani"                   json:"ani,omitempty"`
	Dnis         *string             `bson:"dnis"                  json:"dnis,omitempty"`
	CreateTime   *time.Time          `bson:"createTime"            json:"createTime,omitempty"`
	Category     *string             `bson:"category"              json:"category,omitempty"`
	CallType     *string             `bson:"callType"              json:"callType,omitempty"`
	Media        code.MediaType      `bson:"media"                 json:"media,omitempty"`
	UUI          *string             `bson:"uui"                   json:"uui,omitempty"`
	UEI          *string             `bson:"uei"                   json:"uei,omitempty"`
	MediaURI     *string             `bson:"mediaURI"              json:"mediaURI,omitempty"`
	CallerId     *primitive.ObjectID `bson:"callerId"              json:"callerId,omitempty"`
	CallerType   *string             `bson:"callerType"            json:"callerType,omitempty"`
	CalleeId     *primitive.ObjectID `bson:"calleeId"              json:"calleeId,omitempty"`
	CalleeType   *string             `bson:"calleeType"            json:"calleeType,omitempty"`
	LastRingTime *time.Time          `bson:"lastRingTime"          json:"lastRingTime,omitempty"`
	LastConnTime *time.Time          `bson:"lastConnTime"          json:"lastConnTime,omitempty"`
	LastHoldTime *time.Time          `bson:"lastHoldTime"          json:"lastHoldTime,omitempty"`
	LastConfTime *time.Time          `bson:"lastConfTime"          json:"lastConfTime,omitempty"`
	Participant  []*CallParticipant  `bson:"participant,omitempty" json:"participant,omitempty"`
}

type CallParticipant struct {
	Id               *string                   `bson:"id"               json:"id,omitempty"           ipron:"primary"`
	Part             *CallParticipantPart      `bson:"part,omitempty"   json:"part,omitempty"`
	LocalRedirectNum *string                   `bson:"localRedirectNum" json:"localRedirectNum,omitempty"`
	OriginalNum      *string                   `bson:"originalNum"      json:"originalNum,omitempty"`
	RealNum          *string                   `bson:"realNum"          json:"realNum,omitempty"`
	AddTime          *time.Time                `bson:"addTime"          json:"addTime,omitempty"`
	InstanceURI      *string                   `bson:"instanceURI"      json:"instanceURI,omitempty"`
	State            code.CallParticipantState `bson:"state"            json:"state,omitempty"`
}

type CallParticipantPart struct {
	Type code.CallParticipantPartType `bson:"type"    json:"type,omitempty"`
	Id   *primitive.ObjectID          `bson:"id"      json:"id,omitempty"`
}

/**
 * Realtime Interactions Table
 */

type Interaction struct {
	Id      *primitive.ObjectID `bson:"_id"                    json:"id,omitempty"`
	TntId   *primitive.ObjectID `bson:"tntId"                  json:"tntId,omitempty"`
	CallId  *primitive.ObjectID `bson:"callId"                 json:"callId,omitempty"`
	ReqTime *time.Time          `bson:"reqTime"                json:"reqTime,omitempty"`
	//   ANI               string                  `bson:"ani"                    json:"ani"`
	QueueId       *primitive.ObjectID    `bson:"queueId"                json:"queueId,omitempty"`
	Priority      *int                   `bson:"priority"               json:"priority,omitempty"`
	RouteType     code.IrRouteType       `bson:"routeType"              json:"routeType,omitempty"`
	ChoiceType    code.AgentChoiceMethod `bson:"choiceType"             json:"choiceType,omitempty"`
	Skills        []*InteractionSkill    `bson:"skills,omitempty"       json:"skills,omitempty"`
	Groups        []*primitive.ObjectID  `bson:"groups,omitempty"       json:"groups,omitempty"`
	AcdKind       code.ACDKind           `bson:"acdKind"                json:"acdKind,omitempty"`
	AcdTarget     []string               `bson:"acdTarget"              json:"acdTarget,omitempty"`
	ServiceLevel  *int                   `bson:"serviceLevel"           json:"serviceLevel,omitempty"`
	AbandonMinSec *int                   `bson:"abandonMinSec"          json:"abandonMinSec,omitempty"`
	Media         code.MediaType         `bson:"media"                  json:"media,omitempty"`
	PodUri       *string                 `bson:"podUri"                 json:"podUri,omitempty"`
}

type InteractionSkill struct {
	SkillId *primitive.ObjectID `bson:"skillId"  json:"skillId,omitempty"	ipron:"primary"`
	SkillLv *int                `bson:"skillLv"  json:"skillLv,omitempty"`
}

/**
 * Realtime Conversations Table
 */

type Conversation struct {
	Id         *primitive.ObjectID    `bson:"_id"                   json:"id,omitempty"`
	TntId      *primitive.ObjectID    `bson:"tntId"                 json:"tntId,omitempty"`
	Ani        *string                `bson:"ani"                   json:"ani,omitempty"`
	Dnis       *string                `bson:"dnis"                  json:"dnis,omitempty"`
	CreateTime *time.Time             `bson:"createTime"            json:"createTime,omitempty"`
	Media      code.MediaType         `bson:"media"                 json:"media,omitempty"`
	Segment    []*ConversationSegment `bson:"segment,omitempty"     json:"segment,omitempty"`
}

type ConversationSegment struct {
	ConnId       *primitive.ObjectID `bson:"connId"                json:"connId,omitempty"			ipron:"primary"`
	EndpointType code.EndpointKind   `bson:"endpointType"          json:"endpointType,omitempty"`
	EndpointID   *primitive.ObjectID `bson:"endpointId"            json:"endpointId,omitempty"`
	Direction    code.Direction      `bson:"direction"             json:"direction,omitempty"`
	StartTime    *time.Time          `bson:"startTime"             json:"startTime,omitempty"`
	ConnectTime  *time.Time          `bson:"connectTime"           json:"connectTime,omitempty"`
	EndTime      *time.Time          `bson:"endTime"               json:"endTime,omitempty"`
	Duration     *int                `bson:"duration"              json:"duration,omitempty"`
	EndReason    code.EndReason      `bson:"endReason"             json:"endReason,omitempty"`
}

/**
 * Realtime Phones Table
 */

type Phone struct {
	Id          *primitive.ObjectID `bson:"_id"            json:"id,omitempty"`
	TntId       *primitive.ObjectID `bson:"tntId"          json:"tntId,omitempty"`
	InstanceURI *string             `bson:"instanceURI"    json:"instanceURI,omitempty"`
	State       code.PhoneState     `bson:"state"          json:"state,omitempty"`
	UserId      *primitive.ObjectID `bson:"userId"         json:"userId,omitempty"`
	Extension   *string             `bson:"extension"      json:"extension,omitempty"`
	Register    *PhoneRegister      `bson:"register"       json:"register,omitempty"`
}

type PhoneRegister struct {
	Ip         *string               `bson:"ip"             json:"ip,omitempty"`
	Port       *int                  `bson:"port"           json:"port,omitempty"`
	Protocol   code.ProtocolKind     `bson:"protocol"       json:"protocol,omitempty"`
	State      code.PhoneRegistState `bson:"state"          json:"state,omitempty"`
	RegiExpire *int                  `bson:"regiExpire"     json:"regiExpire,omitempty"`
	Date       *time.Time            `bson:"date"           json:"date,omitempty"`
}

type Trunk struct {
	Id               *primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	TntId            *primitive.ObjectID `bson:"tntId" json:"tntId,omitempty"`
	EdgeId           *primitive.ObjectID `bson:"edgeId" json:"edgeId,omitempty"`
	InstanceURI      *string             `bson:"instanceURI" json:"instanceURI,omitempty"`
	State            code.TrunkState     `bson:"state" json:"state,omitempty"`
	AliveChkRetryCnt *int                `bson:"aliveChkRetryCnt" json:"aliveChkRetryCnt,omitempty"`
	StartTime        *time.Time          `bson:"startTime" json:"startTime,omitempty"`
}

/*
type TrunkRegister struct {
	Id               *primitive.ObjectID   `bson:"_id" json:"id,omitempty"`
	RegisterNumber   *string               `bson:"registerNumber" json:"registerNumber,omitempty"`
	RegisterState    code.TrunkRegistState `bson:"registerState" json:"registerState,omitempty"`
	RegisterTime     *time.Time            `bson:"registerTime" json:"registerTime,omitempty"`
	RegisterInterval *int                  `bson:"registerInterval" json:"registerInterval,omitempty"`
}
*/

type DidPort struct {
	Id       *primitive.ObjectID `bson:"_id"      json:"id,omitempty"`
	TntId    *primitive.ObjectID `bson:"tntId"    json:"tntId,omitempty"`
	Dnis     *string             `bson:"dnis"     json:"dnis,omitempty"`
	RegDate  *time.Time          `bson:"regDate"  json:"regDate,omitempty"`
	RegState code.DidPortState   `bson:"regState" json:"regState,omitempty"`
}

type Media struct {
	Id         *primitive.ObjectID `bson:"_id"         json:"id,omitempty"`
	PodName    *string             `bson:"podName"     json:"podName,omitempty"`
	PodIp      *string             `bson:"podIp"       json:"podIp,omitempty"`
	MaxAudioCh *int                `bson:"maxAudioCh"  json:"maxAudioCh,omitempty"`
	UseAudioCh *int                `bson:"useAudioCh"  json:"useAudioCh,omitempty"`
	MaxVideoCh *int                `bson:"maxVideoCh"  json:"maxVideoCh,omitempty"`
	UseVideoCh *int                `bson:"useVideoCh"  json:"useVideoCh,omitempty"`
	Expire     *time.Time          `bson:"expire"      json:"expire,omitempty"`
}

func (s *User) Marshal() ([]byte, error) {
	return bson.Marshal(s)
}

func (s *Call) Marshal() ([]byte, error) {
	return bson.Marshal(s)
}

func (s *Interaction) Marshal() ([]byte, error) {
	return bson.Marshal(s)
}

func (s *Conversation) Marshal() ([]byte, error) {
	return bson.Marshal(s)
}

func (s *Phone) Marshal() ([]byte, error) {
	return bson.Marshal(s)
}

func (s *Trunk) Marshal() ([]byte, error) {
	return bson.Marshal(s)
}

func (s *DidPort) Marshal() ([]byte, error) {
	return bson.Marshal(s)
}

func (s *Media) Marshal() ([]byte, error) {
	return bson.Marshal(s)
}

//func (s *TrunkRegister) Mashal() ([]byte, error) {
//	return bson.Marshal(s)
//}

func T2String(inf interface{}) string {
	var output string
	value := reflect.ValueOf(inf)
	item := value.Type()

	switch item.Kind() {
	case reflect.Ptr:
		if value.IsZero() == false {
			return T2String(value.Elem().Interface())
		}
	case reflect.Struct:
		switch item.Name() {
		case "Time":
			return fmt.Sprintf("%v", inf)
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
	default:
		return fmt.Sprintf("%v", inf)
	}

	return output
}
