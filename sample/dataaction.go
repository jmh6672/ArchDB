package main

import (
	"log"
	"fmt"
	"time"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"	

	"gitlab.com/ipron-cloud/ArchDB/schema/service"
	"gitlab.com/ipron-cloud/ArchDB/schema/realtime"
	"gitlab.com/ipron-cloud/ArchDB/schema/code"
)

func OpenConnect(user string, password string, url string) (client *mongo.Client){
   credential := options.Credential{
      Username: user,
      Password: password,
   }
   connStr := fmt.Sprintf("%s/?replicaSet=rs0&directConnection=true", url)
   clientOptions := options.Client().ApplyURI(connStr).SetAuth(credential)
   client, err := mongo.Connect(context.TODO(), clientOptions)
   if err != nil {
      log.Fatal(err)
   }

   // Check the connection
   err = client.Ping(context.TODO(), nil)
   if err != nil {
      log.Fatal(err)
   }

   return client
}

func OpenDatabase(conn *mongo.Client, dbName string) (db *mongo.Database) {
	db = conn.Database(dbName)
	return db
}


func InsertAccount(db *mongo.Database, name string, alias string, option service.AccountServiceOption) (string, error) {

	id := primitive.NewObjectID()

	account := service.Account{
		Id: &id,
		Name: name,
		Alias: alias,
		ServiceOption: &option,
	}

	item, err := bson.Marshal(&account)
	if err != nil {
		return "", err
	}

	_, err = db.Collection("account").InsertOne(context.TODO(), item)
	if err != nil {
		return "", err
	}

	return account.Id.Hex(), err
}


func UpdateAccount(db *mongo.Database, accountId string, limitCalls int) (error) {

	objId, err := primitive.ObjectIDFromHex(accountId)

	// 기존 데이터 존재시 업데이트
	filter := bson.M{ "_id": objId }
	update := bson.M{ "$set": bson.M{ "serviceOption.limitConcurrentCalls": limitCalls } }
	_, err = db.Collection("account").UpdateOne(context.TODO(), filter, update)

	return err
}


func InsertFlow(db *mongo.Database, tntId string, flowKind code.FlowKind, flowName string, tags []string, option service.FlowServiceOption) (string, error) {

	tntObjId, err := primitive.ObjectIDFromHex(tntId)
	if err != nil {
		return "", err
	}

	newId := primitive.NewObjectID()
	now := time.Now()

	flow := service.Flow{
		Id: &newId,
		TntId: &tntObjId,
		FlowKind: flowKind,
		FlowName: flowName,
		FlowVersion: "",
		FlowStatus: code.FlowUndefined,
		EditLock: false,
		EditUserId: "",
		Tags: tags,
		ServiceOption: &option,
		Reservation: &service.FlowVerReservation{
			Enable: false,
			FlowVersion: "",
			Status: "",
			TargetTime: &now,
		},
		// Versions: FlowVersion{},
	}

	item, err := flow.Marshal()
	_, err = db.Collection("flow").InsertOne(context.TODO(), item)
	if err != nil {
		return "", err
	}

	return flow.Id.Hex(), err
}

func InserFlowVersion(db *mongo.Database, flowId string, version string, status code.FlowVersionStatus, desc string) (string, error) {

	now := time.Now()

	flowVersion := service.FlowVersion{
		Version: version,
		Status: status,
		LastEditTime: &now,
		Desc: desc,
	}

	objId, err := primitive.ObjectIDFromHex(flowId)
	if err != nil {
		return "", err
	}

	filter := bson.M{ "_id": objId, "versions.version": version }
	update := bson.M{
		"$set": bson.M{ "versions.$": flowVersion },
	}
	result, err := db.Collection("flow").UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Printf("set fail\n")
		return "", err
	}
	if result.ModifiedCount == 1 {
		log.Printf("set success\n")
		return flowId, nil
	}

	filter = bson.M{ "_id": objId }
	update = bson.M{
		"$push": bson.M{ "versions": flowVersion },
	}
	result, err = db.Collection("flow").UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
		log.Printf("push fail\n")
		return "", err
	}

	return flowId, nil
/*
	filter := bson.M{ "_id": objId }	
	arrayFilter := options.ArrayFilters{
		Filters: bson.A{bson.M{"x.version": version}},
	}
	upsert := true
	opts := options.UpdateOptions{
		ArrayFilters: &arrayFilter,
		Upsert: &upsert,
	}
	update := bson.M{
		"$set": bson.M{ "versions.$[x]": flowVersion },
	}

	result, err := db.Collection("flow").UpdateOne(context.TODO(), filter, update, &opts)
	if err != nil {
		log.Fatal(err)
		log.Printf("push fail\n")
		return "", err
	}
	if result.ModifiedCount != 1 {
		log.Printf("update count fail : %d\n", result.ModifiedCount)
	}

	return flowId, nil
*/
}


func InserGroup(db *mongo.Database, tntId string, name string, tags []string, didNum string, groupPath []string, schedule service.GroupSchedule) (string, error) {

	tntObjId, err := primitive.ObjectIDFromHex(tntId)
	if err != nil {
		return "", err
	}

	newId := primitive.NewObjectID()

	group := service.Group{
		Id: &newId,
		TntId: &tntObjId,
		GroupName: name,
		Tags: tags,
		DidNum: didNum,
		GroupPath: groupPath,
		Schedule: &schedule,
	}

	item, err := group.Marshal()
	_, err = db.Collection("group").InsertOne(context.TODO(), item)
	if err != nil {
		return "", err
	}

	log.Printf("Inserted Group '%s' : %s\n", name, group.Id.Hex())

	return group.Id.Hex(), err	
}

func InserUser(db *mongo.Database, tntId string, email string, password string, name string, tags []string, authLv code.UserAuthLevel, accessAuth string, groupId string, extension string, didNum string) (string, error) {
	tntObjId, err := primitive.ObjectIDFromHex(tntId)
	if err != nil {
		return "", err
	}
	grpObjId, err := primitive.ObjectIDFromHex(groupId)
	if err != nil {
		return "", err
	}

	newId := primitive.NewObjectID()
	now := time.Now()

	user := service.Users{
		Id: &newId,
		TntId: &tntObjId,
		Email: email,
		Password: password,
		Name: name,
		Tags: tags,
		AuthLevel: authLv,
		AccessAuth: accessAuth,
		GroupId: &grpObjId,
		Extension: extension,
		DidNum: didNum,
		LockYn: false,
		ExpireDate: &now,
		LastLoginDate: &now,
	}

	item, err := user.Marshal()
	_, err = db.Collection("users").InsertOne(context.TODO(), item)
	if err != nil {
		return "", err
	}

	log.Printf("Inserted User '%s' : %s\n", name, user.Id.Hex())

	return user.Id.Hex(), err	
}


func AttachUserSkill(db *mongo.Database, userId string, skillId string, level int8, priority int8) error {

	userObjId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}
	skillObjId, err := primitive.ObjectIDFromHex(skillId)
	if err != nil {
		return err
	}

	skill := service.UserSkill{
		SkillId: skillObjId,
		SkillLv: int(level),
		SkillPri: int(priority),
	}

	// 기존 데이터 존재시 업데이트
	filter := bson.M{ "_id": userObjId, "skillSets.skillId": skillObjId }
	update := bson.M{ "$set": bson.M{ "skillSets.$": skill } }
	result, err := db.Collection("users").UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	if result.ModifiedCount == 1 {
		log.Printf("set success\n")
		return nil
	}

	// 기존 데이터 없으면 추가
	filter = bson.M{ "_id": userObjId }
	update = bson.M{ "$push": bson.M{ "skillSets": skill } }
	result, err = db.Collection("users").UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func AttachUserPhone(db *mongo.Database, userId string, phoneId string) error {
	userObjId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}
	phoneObjId, err := primitive.ObjectIDFromHex(phoneId)
	if err != nil {
		return err
	}

	// 기존 데이터 존재시 업데이트
	filter := bson.M{ "_id": userObjId }
	update := bson.M{ "$set": bson.M{ "phoneId": phoneObjId } }
	_, err = db.Collection("users").UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return nil
}


func UpdateUserSvcOption(db *mongo.Database, userId string, option service.UserServiceOptin) error {
	return nil
}


func InsertSkills(db *mongo.Database, tntId string, name string, desc string) (string, error) {

	tntObjId, err := primitive.ObjectIDFromHex(tntId)
	if err != nil {
		return "", err
	}

	newId := primitive.NewObjectID()

	skill := service.Skills{
		Id: &newId,
		TntId: &tntObjId,
		Name: name,
		Desc: desc,
	}

	item, err := skill.Marshal()
	_, err = db.Collection("skills").InsertOne(context.TODO(), item)
	if err != nil {
		return "", err
	}

	log.Printf("Inserted Skill '%s': %s\n", name, skill.Id.Hex())

	return skill.Id.Hex(), err
}

func InsertQueue(db *mongo.Database, tntId string, name string, desc string, tags []string) (string, error) {

	tntObjId, err := primitive.ObjectIDFromHex(tntId)
	if err != nil {
		return "", err
	}

	newId := primitive.NewObjectID()

	queue := service.Queue{
		Id: &newId,
		TntId: &tntObjId,
		Name: name,
		Desc: desc,
		Tags: tags,
	}

	item, err := queue.Marshal()
	_, err = db.Collection("queue").InsertOne(context.TODO(), item)
	if err != nil {
		return "", err
	}

	log.Printf("Inserted Queue '%s' : %s\n", name, queue.Id.Hex())

	return queue.Id.Hex(), err
}

func UpdateQueueOption(db *mongo.Database, queueId string, option service.QueueOption) error {

	queueObjId, err := primitive.ObjectIDFromHex(queueId)
	if err != nil {
		return err
	}

	// 기존 데이터 존재시 업데이트
	filter := bson.M{ "_id": queueObjId }
	update := bson.M{ "$set": bson.M{ "option": option } }
	_, err = db.Collection("queue").UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return nil	
}


func UpdateQueueRouteOptStandard(db *mongo.Database, queueId string, agentChoice code.AgentChoiceMethod) error {

	queueObjId, err := primitive.ObjectIDFromHex(queueId)
	if err != nil {
		return err
	}

	// 기존 데이터 존재시 업데이트
	filter := bson.M{ "_id": queueObjId }
	update := bson.M{ "$set": bson.M{ "routeKind": code.QueueRouteStandard, "agentChoiceMethod": agentChoice } }
	_, err = db.Collection("queue").UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return nil	
}

func UpdateQueueRouteOptExpend(db *mongo.Database, queueId string, agentChoice code.AgentChoiceMethod, skillId []string, delay []int16) error {

	queueObjId, err := primitive.ObjectIDFromHex(queueId)
	if err != nil {
		return err
	}

	var expendOpt []service.QueueExternRouteOpt

	for i, v := range skillId {
		if v == "" { break }		
		objId, _ := primitive.ObjectIDFromHex(v)
		expendOpt = append(expendOpt, service.QueueExternRouteOpt{SkillId: objId, DelaySec: int(delay[i])})
	}

	// 기존 데이터 존재시 업데이트
	filter := bson.M{ "_id": queueObjId }
	update := bson.M{ "$set": bson.M{ "routeKind": code.QueueRouteExpend, "agentChoiceMethod": agentChoice, "externRouteOpts": expendOpt } }
	_, err = db.Collection("queue").UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return nil	
}


func InsertSite(db *mongo.Database, tntId string, name string, desc string) (string, error) {

	tntObjId, err := primitive.ObjectIDFromHex(tntId)
	if err != nil {
		return "", err
	}

	newId := primitive.NewObjectID()

	site := service.Site{
		Id: &newId,
		TntId: &tntObjId,
		Name: name,
		Desc: desc,
	}

	item, err := site.Marshal()
	_, err = db.Collection("site").InsertOne(context.TODO(), item)
	if err != nil {
		return "", err
	}

	log.Printf("Inserted Site '%s' : %s\n", name, site.Id.Hex())

	return site.Id.Hex(), err
}

func InsertTrunk(db *mongo.Database, tntId string, name string, enable bool, siteId string, kind code.TrunkKind, proto code.ProtocolKind) (string, error) {

	tntObjId, _ := primitive.ObjectIDFromHex(tntId)
	siteObjId, _ := primitive.ObjectIDFromHex(siteId)

	newId := primitive.NewObjectID()

	trunk := service.Trunks{
		Id: &newId,
		TntId: &tntObjId,
		Name: name,
		Enable: enable,
		SiteId: &siteObjId,
		Kind: kind,
		Protocol: proto,
		AliveCheck: &service.TrunkAliveCheck{
			AliveCheckEnable: false,
			AliveCheckInterval: 0,
			AliveCheckRetry: 3,
		},
		Auth: &service.TrunkAuth{
			MD5Enable: false,
			Realm: "",
			MD5Id: "",
			MD5Pwd: "",
		},
		OutboundOption: &service.TrunkOutboundOption{
			ANINum: "",
			ANIName: "",
			AlwaysANINumEnable: false,
			AlwaysANINameEnable: false,
		},
	}

	item, err := trunk.Marshal()
	_, err = db.Collection("trunks").InsertOne(context.TODO(), item)
	if err != nil {
		return "", err
	}

	log.Printf("Inserted Trunk '%s' : %s\n", name, trunk.Id.Hex())

	return trunk.Id.Hex(), err
}

func AttachTrunkNetwork(db *mongo.Database, trunkId string, ipPort string, status bool) error {

	trunkObjId, err := primitive.ObjectIDFromHex(trunkId)
	if err != nil {
		return err
	}

	network := service.TrunkNetwork{
		IpPort: ipPort,
		Status: status,
	}

	// 기존 데이터 존재시 업데이트
	filter := bson.M{ "_id": trunkObjId, "networks.ipPort": ipPort }
	update := bson.M{ "$set": bson.M{ "networks.$": network } }
	result, err := db.Collection("trunks").UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	if result.ModifiedCount == 1 {
		log.Printf("set success\n")
		return nil
	}

	// 기존 데이터 없으면 추가
	filter = bson.M{ "_id": trunkObjId }
	update = bson.M{ "$push": bson.M{ "networks": network } }
	result, err = db.Collection("trunks").UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

// db.trunks.updateOne({_id: ObjectId("6167768724181009e98a2d5b")}, {$set : { "networks.0.ipPort": "192.168.101.21:5060"}})
func UpdateTrunkIpPort(db *mongo.Database, trunkId string, pos int16, ipPort string) error {

	trunkObjId, err := primitive.ObjectIDFromHex(trunkId)
	if err != nil {
		return err
	}
/*
	network := service.TrunkNetwork{
		IpPort: ipPort,
		Status: status,
	}
*/

	// 기존 데이터 존재시 업데이트
	filter := bson.M{ "_id": trunkObjId }
	update := bson.M{ "$set": bson.M{ fmt.Sprintf("networks.%d.ipPort", pos): ipPort } }
	result, err := db.Collection("trunks").UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	if result.ModifiedCount == 1 {
		log.Printf("set success\n")
		return nil
	}

	return nil
}


func InsertPhone(db *mongo.Database, tntId string, name string, modelId string, siteId string, macAddr string, codec []string) (string, error) {

	tntObjId, _ := primitive.ObjectIDFromHex(tntId)
	siteObjId, _ := primitive.ObjectIDFromHex(siteId)

	newId := primitive.NewObjectID()

	phone := service.Phones{
		Id: &newId,
		TntId: &tntObjId,
		Name: name,
		ModelId: modelId,
		SiteId: &siteObjId,
		MacAddr: macAddr,
		Codecs: codec,
	}

	item, err := phone.Marshal()
	_, err = db.Collection("phones").InsertOne(context.TODO(), item)
	if err != nil {
		return "", err
	}

	log.Printf("Inserted Phone '%s' : %s\n", name, phone.Id.Hex())

	return phone.Id.Hex(), err
}



func InsertNumberPlan(db *mongo.Database, tntId string, siteId string, name string, planType code.NumPlanType, match service.DialPlanNumberInfo, edit service.DialPlanNumberEditRule, target service.DialPlanTarget) (string, error) {

	tntObjId, _ := primitive.ObjectIDFromHex(tntId)
	siteObjId, _ := primitive.ObjectIDFromHex(siteId)

	plan := service.NumberPlan{
		Id: primitive.NewObjectID(),
		TntId: tntObjId,
		SiteId: siteObjId,
		Name: name,
		Type: planType,
		Number: match,
		EditRule: edit,
		Target: target,
	}

	item, err := plan.Marshal()
	_, err = db.Collection("numberPlan").InsertOne(context.TODO(), item)
	if err != nil {
		return "", err
	}

	log.Printf("Inserted Plan '%s' : %s\n", name, plan.Id.Hex())

	return plan.Id.Hex(), err
}


func InsertRoutePoint(db *mongo.Database, tntId string, siteId string, name string, method code.RouteMethod, trunks []string) (string, error) {

	tntObjId, _  := primitive.ObjectIDFromHex(tntId)
	siteObjId, _ := primitive.ObjectIDFromHex(siteId)

	route := service.RoutePoints{
		Id: primitive.NewObjectID(),
		TntId: tntObjId,
		SiteId: siteObjId,
		Name: name,
		RouteMethod: method,
		Trunks: trunks,
	}

	item, err := route.Marshal()
	_, err = db.Collection("routePoints").InsertOne(context.TODO(), item)
	if err != nil {
		return "", err
	}

	log.Printf("Inserted Route '%s' : %s\n", name, route.Id.Hex())

	return route.Id.Hex(), err
}

func InsertDIDPlan(db *mongo.Database, tntId string, dnis string, desc string, routeKind code.DidRouteKind, point string) (string, error) {

	tntObjId, _ := primitive.ObjectIDFromHex(tntId)

	newId := primitive.NewObjectID()

	didPlan := service.DidPlan{
		Id: &newId,
		TntId: &tntObjId,
		DNISPattern: dnis,
		Desc: desc,
		RouteKind: routeKind,
		RoutePoint: point,
	}

	item, err := didPlan.Marshal()
	_, err = db.Collection("didPlan").InsertOne(context.TODO(), item)
	if err != nil {
		return "", err
	}

	log.Printf("Inserted didPlan '%s' : %s\n", dnis, didPlan.Id.Hex())

	return didPlan.Id.Hex(), err
}


func InsertRealtimeUser(db *mongo.Database, tntId string, userId string, state code.UserState, cause code.UserStateCause, calls int16, sec int16, ready int16) (string, error) {

	usrObjId, _ := primitive.ObjectIDFromHex(userId)
	tntObjId, _ := primitive.ObjectIDFromHex(tntId)

	users := realtime.User{
		Id: usrObjId,
		TntId: tntObjId,
		State: state,
		Cause: string(cause),
		StateChgTime: time.Now(),
		AcdConnTotCalls: calls,
		AcdConnTotSec: sec,
		ReadyTotSec: ready,
	}

	item, err := users.Marshal()
	_, err = db.Collection("users").InsertOne(context.TODO(), item)
	if err != nil {
		return "", err
	}

	return users.Id.Hex(), err
}


func InsertRealtimeCall(db *mongo.Database, tntId string, ani string, dnis string, media code.MediaType, uui string, uei string, mediaURI string) (string, error) {

	tntObjId, _ := primitive.ObjectIDFromHex(tntId)

	call := realtime.Call{
		Id: primitive.NewObjectID(),
		TntId: tntObjId,
		Ani: ani,
		Dnis: dnis,
		Media: media,
		UUI: uui,
		UEI: uei,
		MediaURI: mediaURI,
	}

	item, err := call.Marshal()
	_, err = db.Collection("calls").InsertOne(context.TODO(), item)
	if err != nil {
		return "", err
	}

	return call.Id.Hex(), err
}


func InsertRealtimeInteraction(db *mongo.Database, tntId string, callId string, queueId string, priority int8, routeType code.IrRouteType, choiceType code.AgentChoiceMethod, skills []realtime.InteractionSkill, groups []primitive.ObjectID, svcLv int16, abandon int16) (string, error) {

	tntObjId, _ := primitive.ObjectIDFromHex(tntId)
	callObjId, _ := primitive.ObjectIDFromHex(callId)
	queueObjId, _ := primitive.ObjectIDFromHex(queueId)

	ir := realtime.Interaction{
		Id: primitive.NewObjectID(),
		TntId: tntObjId,
		CallId: callObjId,
		ReqTime: time.Now(),
		QueueId: queueObjId,
		Priority: priority,
		RouteType: routeType,
		ChoiceType: choiceType,
		Skills: skills,
		Groups: groups,
		ServiceLevel: svcLv,
		AbandonMinSec: abandon,
	}

	item, err := ir.Marshal()
	_, err = db.Collection("interactions").InsertOne(context.TODO(), item)
	if err != nil {
		return "", err
	}

	return ir.Id.Hex(), err
}

func SelectAccount(db *mongo.Database, accountId string, account *service.Account) (error) {

  id, _ := primitive.ObjectIDFromHex(accountId)

  err := db.Collection("account").FindOne(context.TODO(), bson.M{"_id": id}).Decode(account)

  return err
}

func GetObjId(hex string) (primitive.ObjectID) {
  id, _ := primitive.ObjectIDFromHex(hex)
  return id
}

