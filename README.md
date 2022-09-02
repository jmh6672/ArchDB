# MongoDB 접근용 스키마 ORM 사용을 위한

## 사용방법

```golang
import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

    "{gitRepo 주소}/ArchDB/schema/service"
    "{gitRepo 주소}/ArchDB/schema/service/code"
)

func main() {
    user := service.Users{
        Id: code.ObjId(primitive.NewObjectID()),
        TntId: code.String("abc"),
        Email: code.String("test@gmail.com"),
    }
}
```

주) 데이터 업데이트하는 경우 변경할 항목만 선별적으로 업데이트하는 것을 지정하기 위해 항목들을 포인트타입으로 정의되어 있음

```golang
type AccountServiceOption struct {
    TTSEngine            *string          `bson:"ttsEngine"            json:"ttsEngine,omitempty"`
    DefaultLang          *string          `bson:"defaultLang"          json:"defaultLang,omitempty"`
    AllowFreeSeating     *bool            `bson:"allowfreeSeating"     json:"allowfreeSeating,omitempty"`
    FullRecording        *bool            `bson:"fullRecording"        json:"fullRecording,omitempty"`
    LimitConcurrentCalls *int             `bson:"limitConcurrentCalls" json:"limitConcurrentCalls,omitempty"`
    LimitCPS             *int             `bson:"limitCPS"             json:"limitCPS,omitempty"`
    RoutePolicy          code.RoutePolicy `bson:"routePolicy"          json:"routePolicy"`
}

```

따라서 해당항목의 값을 접근할때에는 기존 사용하는 방식으로 접근하면되고, 값을 설정하는 경우 다음의 유틸리티 함수를 제공합니다.

```golang
accountOpt := service.AccountServiceOption{
    TTSEngine: code.String("default"),
    AllowFreeSeating: code.Bool(true),
    LimitConcurrentCalls: code.Int(10),
}
```

- code.String() : `string` 타입을 `*string` 타입으로 변경
- code.Int() : `int` 타입을 `*int` 타입으로 변경
- code.Bool() : `bool` 타입을 `*bool` 타입으로 변경
- code.Time() : `time.Time` 타입을 `*time.Time` 타입으로 변경
- code.ObjId() : `primitive.ObjectID` 타입을 `*primitive.ObjectID` 타입으로 변경
- code.HexObjId() : `string` 타입을 `*primitive.ObjectID` 타입으로 변경


구조체가 포인트로 구성되어 있어 구조체 전체 출력시 포인트 정보만 출력되어 가독성이 떨어집니다. 이를 해소하기 위해 실데이터로 출력할수 있도록 출력용 문자열을 반환해주는 함수를 제공합니다.

```golang
accountOpt := service.AccountServiceOption{
    TTSEngine: code.String("default"),
    AllowFreeSeating: code.Bool(true),
    LimitConcurrentCalls: code.Int(10),
}

log.Println(service.T2String(accountOpt))
```
