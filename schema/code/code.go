package code

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FlowKind string

const (
	CallInbound  = FlowKind("call-inbound")
	CallOutbound = FlowKind("call-outbound")
	CallInQueue  = FlowKind("call-inqueue")
	ChatInbound  = FlowKind("chat-inbound")
	CallBot      = FlowKind("bot")
)

type FlowStatus string

const (
	FlowUndefined = FlowStatus("undefined")
	FlowPublished = FlowStatus("published")
)

type FlowVersionStatus string

const (
	FlowVersionUndefined  = FlowVersionStatus("undefined")
	FlowVersionEditing    = FlowVersionStatus("editing")
	FlowVersionPublishing = FlowVersionStatus("publishing")
	FlowVersionPublished  = FlowVersionStatus("published")
	FlowVersionFailed     = FlowVersionStatus("failed")
)

type FlowReservationStatus string

const (
	FlowReservationUndefined = FlowReservationStatus("undefined")
	FlowReservationWaiting   = FlowReservationStatus("waiting")
	FlowReservationPublished = FlowReservationStatus("published")
)

type PromptKind string

const (
	PromptFile = PromptKind("file")
	PromptTTS  = PromptKind("tts")
)

type MediaType string

const (
	MediaVoice = MediaType("voice")
	MediaVideo = MediaType("video")
	MediaChat  = MediaType("chat")
	MediaEmail = MediaType("email")
)

type RoutePolicy string

const (
	RoutePolicyInteractionOrder = RoutePolicy("interactionorder")
	RoutePolicySkillPriority    = RoutePolicy("skillpriority")
)

type UserAuthLevel string

const (
	UserAuthSuperAdmin  = UserAuthLevel("superadmin")
	UserAuthAdmin       = UserAuthLevel("admin")
	UserAuthSvcManager  = UserAuthLevel("svcmanager")
	UserAuthUserManager = UserAuthLevel("usermanager")
	UserAuthUser        = UserAuthLevel("user")
)

type UserScheduleType string

const (
	UserScheduleGroup = UserScheduleType("group")
	UserScheduleUser  = UserScheduleType("user")
)

type SvcForwardType string

const (
	SvcForwardUncondition = SvcForwardType("uncondition")
	SvcForwardNoAnswer    = SvcForwardType("noanswer")
	SvcForwardBusy        = SvcForwardType("busy")
)

type QueueRouteKind string

const (
	QueueRouteStandard = QueueRouteKind("standard")
	QueueRouteExpend   = QueueRouteKind("expend")
)

type AgentChoiceMethod string

const (
	AgentChoiceLastLongWait = AgentChoiceMethod("lastLongWait")
	AgentChoiceAccLongWait  = AgentChoiceMethod("accLongWait")
	AgentChoiceTotMinCall   = AgentChoiceMethod("totMinCall")
	AgentChoiceQueMinCall   = AgentChoiceMethod("queMinCall")
	AgentChoiceTotMinTime   = AgentChoiceMethod("totMinTime")
	AgentChoiceQueMinTime   = AgentChoiceMethod("queMinTime")
	AgentChoiceRoundrobin   = AgentChoiceMethod("roundrobin")
)

type UserStateMaster string

const (
	UserStateMstACW  = UserStateMaster("acw")
	UserStateMstIDEL = UserStateMaster("idel")
)

type TrunkKind string

const (
	TrunkCarrier = TrunkKind("carrier")
	TrunkPBX     = TrunkKind("pbx")
	TrunkNormal  = TrunkKind("normal")
)

type ProtocolKind string

const (
	ProtoUDP    = ProtocolKind("udp")
	ProtoTCP    = ProtocolKind("tcp")
	ProtoTLS    = ProtocolKind("tls")
	ProtoWebRTC = ProtocolKind("webrtc")
)

type TrunkStatusKind string

const (
	TrunkStatusNone  = TrunkStatusKind("none")
	TrunkStatusAlive = TrunkStatusKind("alive")
	TrunkStatusDead  = TrunkStatusKind("dead")
)

type DidPortStatusKind string

const (
	DidPortStatusNone  = DidPortStatusKind("none")
	DidPortStatusAlive = DidPortStatusKind("alive")
	DidPortStatusDead  = DidPortStatusKind("dead")
)

type CarrierKind string

const (
	CarrierKT   = CarrierKind("KT")
	CarrierLG   = CarrierKind("LGU+")
	CarrierSKTL = CarrierKind("SKTelink")
	CarrierETC  = CarrierKind("etc")
)

type DidRouteKind string

const (
	DidRouteToUser  = DidRouteKind("user")
	DidRouteToPhone = DidRouteKind("phone")
	DidRouteToFlow  = DidRouteKind("flow")
	DidRouteToQueue = DidRouteKind("queue")
)

type EndpointKind string

const (
	EndpointTrunk = EndpointKind("trunk")
	EndpointUser  = EndpointKind("user")
	EndpointPhone = EndpointKind("phone")
	EndpointFlow  = EndpointKind("flow")
	EndpointQueue = EndpointKind("queue")
)

type NumPlanType string

const (
	NumPlanPattern = NumPlanType("pattern")
	NumPlanPrefix  = NumPlanType("prefix")
)

type NumPlanClassification string

const (
	NumPlanClassEmergency      = NumPlanClassification("emergency")
	NumPlanClassExtension      = NumPlanClassification("extension")
	NumPlanClassLocal          = NumPlanClassification("local")
	NumPlanClassInternational  = NumPlanClassification("international")
	NumPlanClassNetwork        = NumPlanClassification("network")
	NumPlanClassFeature        = NumPlanClassification("feature")
	NumPlanClassFlow           = NumPlanClassification("flow")
	NumPlanClassCompatibleFlow = NumPlanClassification("compatibleflow")
)

type RouteMethod string

const (
	RouteSequence = RouteMethod("sequence")
	RouteBalance  = RouteMethod("balance")
)

type UserAssignState string

const (
	UserAssignStateNone     = UserAssignState("none")
	UserAssignStateAssigned = UserAssignState("assigned")
	UserAssignStateReserved = UserAssignState("reserved")
)

type UserState string

const (
	UserStateLogin    = UserState("login")
	UserStateReady    = UserState("ready")
	UserStateInReady  = UserState("inready")
	UserStateOutReady = UserState("outready")
	UserStateNotReady = UserState("notready")
	UserStateACW      = UserState("afterwork")
	UserStateBusy     = UserState("online")
	UserStateLogout   = UserState("logout")
	UserStateDialing  = UserState("dialing")
	UserStateRinging  = UserState("ringing")
)

type UserStateCause string

const (
	UserNotReadyBusy   = UserStateCause("busy")
	UserNotReadyAway   = UserStateCause("away")
	UserNotReadyBreak  = UserStateCause("break")
	UserNotReadyIdel   = UserStateCause("idel")
	UserBusyInbound    = UserStateCause("inbound")
	UserBusyOutbound   = UserStateCause("outbound")
	UserLogoutWorkoff  = UserStateCause("workoff")
	UserLogoutVacation = UserStateCause("vacation")
	UserLogoutOutside  = UserStateCause("outside")
)

type CallCategory string

const (
	CallCategoryIn       = CallCategory("in")       // 국선 인입 통화
	CallCategoryOut      = CallCategory("out")      // 내선 (발신) 통화
	CallCategoryInternal = CallCategory("internal") // 국선 발신 통화
)

type CallType string

const (
	CallTypeNormal       = CallType("normal")       // 일반 콜
	CallTypeUserTransfer = CallType("usertransfer") // user로부터 호 전환 된 콜
	CallTypeConsult      = CallType("consult")      // 협의 콜
	CallTypeFeature      = CallType("feature")      // 기능 코드 설정 콜
)

type CallParticipantState string

const (
	CallParticipantStateNull       = CallParticipantState("null")
	CallParticipantStateOriginated = CallParticipantState("originated")
	CallParticipantStateDialing    = CallParticipantState("dialing")
	CallParticipantStateAlert      = CallParticipantState("alerting")
	CallParticipantStateConnected  = CallParticipantState("connected")
	CallParticipantStateHold       = CallParticipantState("hold")
	CallParticipantStateQueued     = CallParticipantState("queued")
	CallParticipantStateFailed     = CallParticipantState("failed")
)

type CallParticipantPartType string

const (
	CallParticipantPartTypeTrunk          = CallParticipantPartType("trunk")
	CallParticipantPartTypePhone          = CallParticipantPartType("phone")
	CallParticipantPartTypeUser           = CallParticipantPartType("user")
	CallParticipantPartTypeFlow           = CallParticipantPartType("flow")
	CallParticipantPartTypeQueue          = CallParticipantPartType("queue")
	CallParticipantPartTypeCompatibleFlow = CallParticipantPartType("compatibleflow")
)

type Direction string

const (
	DirectionInbound  = Direction("inbound")
	DirectionOutbound = Direction("outbound")
)

type EndReason string

const (
	EndReasonConferenced      = EndReason("conferenced")
	EndReasonConsult          = EndReason("consult")
	EndReasonEndpoint         = EndReason("endpoint")
	EndReasonClient           = EndReason("client")
	EndReasonForward          = EndReason("forward")
	EndReasonNoAnswer         = EndReason("noAnswer")
	EndReasonNotAvaliable     = EndReason("notAvaliable")
	EndReasonSpam             = EndReason("spam")
	EndReasonTimeout          = EndReason("timeout")
	EndReasonTransportFailure = EndReason("transportFailure")
	EndReasonPeer             = EndReason("peer")
)

type IrRouteType string

const (
	IrRouteSkill      = IrRouteType("skill")
	IrRouteGroup      = IrRouteType("group")
	IrRouteSkillGroup = IrRouteType("skillGroup")
	IrRouteACD        = IrRouteType("acd")
)

type PhoneState string

const (
	PhoneStateNone = PhoneState("none")
	PhoneStateIdle = PhoneState("idle")
	PhoneStateBusy = PhoneState("busy")
)

type PhoneRegistState string

const (
	PhoneRegistStateNone     = PhoneRegistState("none")
	PhoneRegistStateUnRegist = PhoneRegistState("unregist")
	PhoneRegistStateRegist   = PhoneRegistState("regist")
)

type TrunkRegistState string

const (
	TrunkRegistStateNone     = TrunkRegistState("none")
	TrunkRegistStateUnRegist = TrunkRegistState("unregist")
	TrunkRegistStateRegist   = TrunkRegistState("regist")
)

type TrunkState string

const (
	TrunkStateNormal = TrunkState("normal")
	TrunkStateError  = TrunkState("error")
)

type DidPortState string

const (
	DidPortStateUnregist = DidPortState("unregist")
	DidPortStateFirst    = DidPortState("first")
	DidPortStateSecond   = DidPortState("second")
)

type BridgeType string

const (
	BridgeTypeData = BridgeType("data-bridge")
	BridgeTypeTTS  = BridgeType("tts-bridge")
	BridgeTypeSTT  = BridgeType("stt-bridge")
	BridgeTypeNLU  = BridgeType("nlu-bridge")
)

type BridgeAuthType string

const (
	BridgeAuthTypeKey      = BridgeAuthType("key-auth")
	BridgeAuthTypeApp      = BridgeAuthType("app-auth")
	BridgeAuthTypeFile     = BridgeAuthType("file-auth")
	BridgeAuthTypeUserDef  = BridgeAuthType("userdef-auth")
)

type HolidayRepeatType string

const (
	HolidayRepeatNone      = HolidayRepeatType("none")
	HolidayRepeatYear      = HolidayRepeatType("year")
	HolidayRepeatMonth     = HolidayRepeatType("month")
)

type HolidayType string

const (
	HolidayWeekend     = HolidayType("weekend")
	HolidayCountry     = HolidayType("country")
	HolidayTemporary   = HolidayType("temporary")
	HolidayEtc         = HolidayType("etc")
)

type ACDKind string

const (
	ACDKindUNKNOWN    = ACDKind("none")
	ACDKindUSER       = ACDKind("user")
	ACDKindPHONE      = ACDKind("phone")
	ACDKindQUEUE      = ACDKind("queue")
)


func Bool(v bool) *bool {
	return &v
}

func Int(v int) *int {
	return &v
}

func String(v string) *string {
	return &v
}

func Time(v time.Time) *time.Time {
	return &v
}

func ObjId(v primitive.ObjectID) *primitive.ObjectID {
	return &v
}

func ObjStr(v *primitive.ObjectID) string {
	return v.Hex()
}

func HexObjId(v string) *primitive.ObjectID {
	id, _ := primitive.ObjectIDFromHex(v)
	return &id
}
