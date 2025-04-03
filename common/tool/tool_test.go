package tool

import (
	"container/list"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"testing"
)

func TestUploadFile(t *testing.T) {

	getenv := os.Getenv("HOME")
	fmt.Println(getenv)
}

func TestFloat64ToString(t *testing.T) {

	str := Float64ToString(0.10)
	fmt.Printf("str: %s\n", str)
}

func TestGetPath(t *testing.T) {

	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(exePath)

	dir := path.Dir(exePath)
	fmt.Println(dir)

	getwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(getwd)
}

type Deduction struct {
	DeductionAmount string `json:"deductionAmount"`
	DeductionDate   string `json:"deductionDate"`
	Period          string `json:"period"`
	Status          string `json:"status"`
}

func TestStringToJson(t *testing.T) {

	str := `[{"deductionAmount":"0.1","deductionDate":"2024-01-26","period":"1","status":"REFUNDED"},{"deductionAmount":"0.1","deductionDate":"2024-01-26","period":"1","status":"REFUNDED"},{"deductionAmount":"0.1","deductionDate":"2024-01-26","period":"1","status":"REFUNDED"},{"deductionAmount":"0.1","deductionDate":"2024-01-26","period":"1","status":"REFUNDED"}]`

	var deductions []Deduction

	err := json.Unmarshal([]byte(str), &deductions)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(deductions)
}

func TestSturctToString(t *testing.T) {
	m := map[string]interface{}{
		"deductionDate":   "2024-01-26",
		"period":          "1",
		"deductionAmount": "0.1",
		"status":          "REFUNDED",
	}

	encode, err := JSONEncode(m)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(encode))
}

func TestDate2Timestamp(t *testing.T) {
	str := "2024-02-28 00:00:00"

	fmt.Println(DateTime2Timestamp(str))
	fmt.Println(DateTimeToTimestamp(str))
}

func TestMd5ByString(t *testing.T) {
	str := "123456"
	fmt.Println(Md5ByString(str))

	salt := "salt"
	fmt.Println(Md5ByStringWithSalt(str, salt))
}

func TestDaysDifference(t *testing.T) {
	var dayTimestamp int64 = 1725844308
	day := DaysDifference(dayTimestamp)
	fmt.Println("day = ", day)
}

func TestGetMaxAndMinDate(t *testing.T) {
	//min, max := GetMaxAndMinDate(1)
	//fmt.Println(min, max)
	//fmt.Println(min.Unix(), max.Unix())

	min, _ := GetMaxAndMinDate(0)
	overDueDate := min.Unix() - 24*60*60

	fmt.Println(overDueDate, 1716652800-24*60*60)
}

func TestContainer(t *testing.T) {

	l := list.New()
	f := l.PushFront("first")
	fmt.Printf("front1 %+v \n", []*list.Element{f})
	f2 := l.PushFront("second")
	fmt.Printf("front2 %+v \n", []*list.Element{f2})
	f3 := l.PushBack("back")

	es := []*list.Element{f, f2, f3}
	for k, v := range es {
		fmt.Printf("k is %v and v is %v and %+v and next is %v \n", k, v, v.Value, v.Next())
	}
}