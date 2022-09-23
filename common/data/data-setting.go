package data

import (
	"fmt"
	"math"
	"math/rand"
	"net"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/sony/sonyflake"
	"golang.org/x/crypto/bcrypt"
)

var snoflake *sonyflake.Sonyflake

func init() {
    snoflake = sonyflake.NewSonyflake(sonyflake.Settings{
        StartTime:      time.Time{},
        MachineID:      nil,
        CheckMachineID: nil,
    })
}

//雪花算法生成id
func NextSonyFlakeIdInt64() uint64 {
    snoyId, _ := snoflake.NextID()
    return snoyId
}

//格式化价格保留两位小数
func Decimal(value float64) float64 {
    value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
    return value
}

// string转int
func StringToInt(a string) int {
    d, _ := strconv.Atoi(a)
    return d
}

//Atoi是ParseInt(s, 10, 0)的简写。

// string转int64
func StringToInt64(a string) int64 {
    d, _ := strconv.ParseInt(a, 10, 64)
    return d
}

func StringToUint64(a string) uint64 {
    d, _ := strconv.ParseUint(a, 10, 64)
    return d
}

// int转string
func IntToString(a int) string {
    str := strconv.Itoa(a)
    return str
}

// int64转string
func Int64ToString(a int64) string {
    str := strconv.FormatInt(a, 10)
    return str
}

// float转string
func FloatToString(f float64) string {
    return strconv.FormatFloat(f, 'f', -1, 32)
}

// string转float
func StringToFloat(s string) float64 {
    f, _ := strconv.ParseFloat(s, 64)
    return f
}

func StructToMap(obj interface{}) map[string]interface{} {
    t := reflect.TypeOf(obj)
    v := reflect.ValueOf(obj)
    var data = make(map[string]interface{})
    for i := 0; i < t.NumField(); i++ {
        data[t.Field(i).Name] = v.Field(i).Interface()
    }
    return data
}

func ArrToString(arr []string) string {

    if len(arr) == 0 {
        return ""
    }

    var str = ""

    for _, i := range arr {
        str += "," + i
    }

    return str[1:]
}

func RemoveRepByLoop(slc []int) []int {
    result := []int{} // 存放结果
    for i := range slc {
        flag := true
        for j := range result {
            if slc[i] == result[j] {
                flag = false // 存在重复元素，标识为false
                break
            }
        }
        if flag { // 标识为false，不添加进结果
            result = append(result, slc[i])
        }
    }
    return result
}

func RemoveRepByLoopString(slc []string) []string {
    result := []string{} // 存放结果
    for i := range slc {
        flag := true
        for j := range result {
            if slc[i] == result[j] {
                flag = false // 存在重复元素，标识为false
                break
            }
        }
        if flag { // 标识为false，不添加进结果
            result = append(result, slc[i])
        }
    }
    return result
}

func GetRemoteClientIp(r *http.Request) string {
    remoteIp := r.RemoteAddr

    if ip := r.Header.Get("X-Real-IP"); ip != "" {
        remoteIp = ip
    } else if ip = r.Header.Get("X-Forwarded-For"); ip != "" {
        remoteIp = ip
    } else {
        remoteIp, _, _ = net.SplitHostPort(remoteIp)
    }

    //本地ip
    if remoteIp == "::1" {
        remoteIp = "127.0.0.1"
    }

    return remoteIp
}
func Page(limit, page int, count int64) (int, int) {
    pageSetNum := limit // 每页条数

    pageCount := math.Ceil((float64(count)) / (float64(pageSetNum))) // 总页数
    pageNum := page                                                  // 当前页码
    if pageNum > int(pageCount) {                                    // 如果传入的页码超出范围
        //pageNum = int(pageCount)
        return 0, 0
    }
    offset := pageSetNum * (pageNum - 1)
    if offset < 0 {
        offset = 0
    }
    return pageSetNum, offset
}

func RemoveTopStruct(fields map[string]string) string {
    rsp := ""
    for _, err := range fields {
        rsp += err
    }
    return rsp
}

func RandStr(len int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)

}


var (
    StatusName = map[int8]string{
        0: "开启",
        1: "关闭",
    }
    IsDeleteName = map[int8]string{
        0: "未删除",
        1: "已删除",
    }
)

// 加密密码
func HashAndSalt(pwd []byte) string {
    hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
    if err != nil {

    }
    return string(hash)
}

// 验证密码
func ValidatePasswords(hashedPwd string, plainPwd []byte) bool {
    byteHash := []byte(hashedPwd)
    err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
    if err != nil {
        return false
    }
    return true
}