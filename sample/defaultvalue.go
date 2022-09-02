package main

import (
	"log"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gitlab.com/ipron-cloud/ArchDB/schema/service"
	"gitlab.com/ipron-cloud/ArchDB/schema/realtime"
	"gitlab.com/ipron-cloud/ArchDB/schema/code"

	"go.mongodb.org/mongo-driver/mongo"
	"encoding/json"
)

func main() {
	conn := OpenConnect("root", "wlwndgo", "mongodb://100.100.103.163:30071")
	db   := OpenDatabase(conn, "ipron")

	var account service.Account
	SelectAccount(db, "61276d4c7790a01da4bac0fd", &account)

	log.Println(account)

	by, _ := json.Marshal(&account)
	log.Println(string(by))
	

	// for i:=0; i<100; i++ {
	// 	UpdateAccount(db, "6167768724181009e98a2d4b", i)
	// }
}

func main2() {

	conn := OpenConnect("root", "wlwndgo", "mongodb://100.100.103.163:30071")
	db   := OpenDatabase(conn, "ipron")

	masterTntId := "face0000face0000face0000"
	// tntId := "61276d4c7790a01da4bac0fd"
	tntId := AccountAdd(db)

	// flowId := "61276d4c7790a01da4bac0fe"
	flowId := FlowAdd(db, tntId)

	FlowVersionAdd(db, flowId, "1.0.0")
	FlowVersionAdd(db, flowId, "1.0.1")

	// skillId1 := "612db3190cdd2b31e7f5b39e"
	// skillId2 := "612db31a0cdd2b31e7f5b39f"
	skillId1, _ := InsertSkills(db, tntId, "일반응대sk", "")
	skillId2, _ := InsertSkills(db, tntId, "VIP응대sk", "")


	groupId, _  := InserGroup(db, tntId, "최상위그룹", []string{}, "", []string{}, service.GroupSchedule{})
	groupId1, _ := InserGroup(db, tntId, "서브그룹1", []string{"1팀"}, "", []string{groupId}, service.GroupSchedule{})
	groupId2, _ := InserGroup(db, tntId, "서브그룹2", []string{"1팀"}, "", []string{groupId}, service.GroupSchedule{})
	               InserGroup(db, tntId, "서브그룹3", []string{"2팀"}, "1588-1002", []string{groupId}, service.GroupSchedule{})
	groupId1_1, _ := InserGroup(db, tntId, "서브그룹1-1", []string{}, "1588-1001", []string{groupId, groupId1}, service.GroupSchedule{})
	groupId2_1, _ := InserGroup(db, tntId, "서브그룹2-1", []string{}, "1588-1002", []string{groupId, groupId2}, service.GroupSchedule{})

	// groupId3 := "612d79f11e9f447e0e046ec7"
	// groupId1_1 := "612d79f11e9f447e0e046ec8"
	// groupId2_1 := "612d79f11e9f447e0e046ec9"


	// user1 := "612da21a9a29cfeea9a7a74d"
	// user2 := "612da21a9a29cfeea9a7a74e"
	user1, _ := InserUser(db, tntId, "gildong@bridgetec.co.kr", "dnflth", "홍길동", []string{}, code.UserAuthAdmin, "1Team-manager", groupId1_1, "3000", "02-3430-3000")
	AttachUserSkill(db, user1, skillId1, 5, 1)
	AttachUserSkill(db, user1, skillId2, 7, 2)
	user2, _ := InserUser(db, tntId, "mugae@bridgetec.co.kr", "dnflth", "아무개", []string{}, code.UserAuthAdmin, "2Team-manager", groupId2_1, "3001", "02-3430-3001")
	AttachUserSkill(db, user2, skillId1, 5, 2)


	// queue1 := "612dc12ea45efb224a986ce5"
	queue1, _ := InsertQueue(db, tntId, "일반상담큐", "", []string{"1팀"})
	UpdateQueueOption(db, queue1, service.QueueOption{NoAnswerSec: 10, MaxWaitCalls: 100, MaxWaitSec: 600, MinAbandonSec: 3})
	UpdateQueueRouteOptStandard(db, queue1, code.AgentChoiceLastLongWait)

	// queue2 := "612dc12ea45efb224a986ce6"
	queue2, _ := InsertQueue(db, tntId, "VIP상담큐", "", []string{"2팀"})
	UpdateQueueOption(db, queue2, service.QueueOption{NoAnswerSec: 10, MaxWaitCalls: 100, MaxWaitSec: 600, MinAbandonSec: 3})
	UpdateQueueRouteOptExpend(db, queue2, code.AgentChoiceTotMinCall, []string{skillId1, skillId2}, []int16{10, 30})


	// masterSiteId := "612dcd5908234b511b85f6c8"
	masterSiteId, _ := InsertSite(db, masterTntId, "기본Site", "")

	// siteId := "612dc85c04822b5f45f7a20f"
	siteId, _ := InsertSite(db, tntId, "여의도", "")

	// trunkId := "612dcd7420bbe3fdf0fa3e2e"
	trunkId, _ := InsertTrunk(db, masterTntId, "국선트렁크1", true, masterSiteId, code.TrunkCarrier, code.ProtoUDP)
	AttachTrunkNetwork(db, trunkId, "192.168.101.21:5060", true)

	// phoneId1 := "612dd7108826dc98124a9b57"
	// phoneId2 := "612dd7ec9c7d1e61217ae8d9"
	phoneId1, _ := InsertPhone(db, tntId, "user1-phone", "audiocode-VVX201", siteId, "0004f22035a6", []string{"G.711Mu","G.711A","G.729AB"})
	phoneId2, _ := InsertPhone(db, tntId, "user2-phone", "audiocode-VVX201", siteId, "0004f2208dce", []string{"G.711Mu","G.711A","G.729AB"})

	AttachUserPhone(db, user1, phoneId1)
	AttachUserPhone(db, user2, phoneId2)

	// planId1 := "612de0e83da7658cf7599989"
	InsertNumberPlan(db, tntId, siteId, "Outbound#1", 
		code.NumPlanPrefix, 
		service.DialPlanNumberInfo{Prefix: "9", MinLen:3, MaxLen: 12}, 
		service.DialPlanNumberEditRule{KeepEnable:false, PrefixDelLen:1},
		service.DialPlanTarget{Classification: code.NumPlanClassNetwork, Point: "DOD-Route1"},
	)
	InsertRoutePoint(db, tntId, siteId, "DOD-Route1", code.RouteSequence, []string{"common-trunk"})

	// didPlan1 := "612de39431b01c9edd8de98a"
	InsertDIDPlan(db, tntId, "15881000", "시험용 DID번호", code.DidRouteToUser, user1)

	return

	/**
	 * REAL-TIME DATA
	 */

	dbRealtime   := OpenDatabase(conn, "ipron-realtime")

	RealtimeUserAdd(dbRealtime, user1, tntId)
	// call1 := "61441fece4a5b141a8c244d9"
	call1 := RealtimeCallAdd(dbRealtime, tntId, "01012345678", "15881000", code.MediaVoice, "uui-data", "uei-data", "media-1.ipron:8080")
	//ir1 := "61441fece4a5b141a8c244da"
	ir1 := RealtimeInteractionAdd(dbRealtime, tntId, call1, queue1, 5, code.IrRouteSkill, code.AgentChoiceLastLongWait, 20, 5, 
		[]realtime.InteractionSkill{
			realtime.InteractionSkill{SkillId: GetObjId(skillId1), SkillLv: 50},
		  	realtime.InteractionSkill{SkillId: GetObjId(skillId2), SkillLv: 70},
		},
		[]primitive.ObjectID{ GetObjId(groupId1_1) },
	)

	log.Printf("Inserted Realtime Interaction : %s\n", ir1)



	log.Println("Finished")
}



func AccountAdd(db *mongo.Database) string {

	serviceOption := service.AccountServiceOption{
		TTSEngine: "",
		DefaultLang: "",
		AllowFreeSeating: true,
		FullRecording: false,
		LimitConcurrentCalls: 100,
		LimitCPS: 5,
	}

	tntId, err := InsertAccount(db, "bridgetec", "bt", serviceOption)
	if err != nil {
		return ""
	}
	log.Printf("Inserted Account : %s\n", tntId)

	return tntId
}


func FlowAdd(db *mongo.Database, tntId string) string {

	flowOption := service.FlowServiceOption{
		Fdt: 5,
		Idt: 2,
	}

	flowId, err := InsertFlow(db, tntId, code.CallInbound, "sample-inbound", []string{}, flowOption)
	if err != nil {
		return ""
	}

	log.Printf("Inserted Flow : %s\n", flowId)

	return flowId
}

func FlowVersionAdd(db *mongo.Database, flowId string, version string) string {

	editId, err := InserFlowVersion(db, flowId, version, code.FlowVersionUndefined, "sample flow")
	if err != nil {
		return ""
	}

	log.Printf("Inserted Flow Version : %s:%s\n", flowId, version)

	return editId
}


func RealtimeUserAdd(db *mongo.Database, userId string, tntId string) string {
	userId, err := InsertRealtimeUser(db, tntId, userId, code.UserStateReady, code.UserReadyInbound, 10, 70*10, 600)
	if err != nil {
		return ""
	}

	log.Printf("Inserted Realtime User : %s\n", userId)

	return userId
}

func RealtimeCallAdd(db *mongo.Database, tntId string, ani string, dnis string, media code.MediaType, uui string, uei string, mediaURI string) string {
	callId, err := InsertRealtimeCall(db, tntId, ani, dnis, media, uui, uei, mediaURI)
	if err != nil {
		return ""
	}

	log.Printf("Inserted Realtime Call : %s\n", callId)

	return callId
}

func RealtimeInteractionAdd(db *mongo.Database, tntId string, callId string, queueId string, priority int8, routeType code.IrRouteType, choiceType code.AgentChoiceMethod, svcLv int16, abandon int16, skills []realtime.InteractionSkill, groups []primitive.ObjectID) string {
	irId, err := InsertRealtimeInteraction(db, tntId, callId, queueId, priority, routeType, choiceType, skills, groups, svcLv, abandon)
	if err != nil {
		return ""
	}

	log.Printf("Inserted Realtime IR : %s\n", irId)

	return irId
}

