package structs

import (
	"encoding/json"
	"fmt"
	"log"
)

// JSON은 struct를 marshal(stringify)화 할 때 대문자 필드만 변환함
type movie struct {
	ID    int64  `json:",string"`
	Title string `json:"title"` // json 옵션 추가시 직렬화 필드 이름 정의 가능
	Year  int    `json:"year"`
}

/*

ID의 ,string은 중요함. Go 언어 자체만이 아닌 웹 JS에서 숫자형은 8바이트이기 때문에 정확도가 떨어짐.
따라서 64비트 정수를 주고 받을 때는 `json:",string"` 옵션을 추가함
string 형이지만 역직렬화시 Go에서 정수형으로 변환되기 때문에 괜찮음

*/

// ShowMarshalJSON shows json marshal
func ShowMarshalJSON() {
	m := movie{
		1,
		"바람",
		2010,
	}
	b, err := json.Marshal(m)
	if err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Println(string(b))
}

// ShowUnmarshalJSON shows json unmarshal
func ShowUnmarshalJSON() {
	b := []byte(`{"Title":"바람","Year":2010}`)
	m := movie{}
	err := json.Unmarshal(b, &m)
	if err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Println(m.Title, m.Year)
}

/*

func main() {
	structs.ShowMarshalJSON()
	structs.ShowUnmarshalJSON()
}

*/
