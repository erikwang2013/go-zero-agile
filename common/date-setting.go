package common

import "time"

var defaultTime = "2006-01-02 15:04:05"

func GetTodayTime() string {
    return time.Now().Format(defaultTime)
}

func GetStampToDate(TimeStamp int64) string {
    return time.Unix(TimeStamp, 0).Format(defaultTime)
}

func GetStampToDates(TimeStamp int64) string {
    return time.Unix(TimeStamp, 0).Format("2006-01-02")
}

// GetSubDate 动态改变多少天前，或后的日期
func GetSubDate(Year, Month, Date int) string {
    return time.Now().AddDate(Year, Month, Date).Format("2006-01-02")
}

func GetDateToTimeStamp(Date string) int64 {
    stamp, _ := time.ParseInLocation(defaultTime, Date, time.Local)
    return stamp.Unix()
}

//指定时间格式后，格式化时间
func DataToData(dataType string, data string) string {
    time, _ := time.Parse(dataType, data)
    return GetStampToDate(time.Unix())
}

//获得当前月的初始和结束日期
func GetMonthDay(types int) (string, string) {
    now := time.Now()
    currentYear, currentMonth, _ := now.Date()
    currentLocation := now.Location()

    firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
    lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
    f := firstOfMonth.Unix()
    l := lastOfMonth.Unix()
    if types == 1 {
        return time.Unix(f, 0).Format("2006-01-02"), time.Unix(l, 0).Format("2006-01-02")
    }
    return time.Unix(f, 0).Format("2006-01-02") + " 00:00:00", time.Unix(l, 0).Format("2006-01-02") + " 23:59:59"
}

//获取当月的第一天和上个月的最后一天
func GetLastMonthDayAndMontDay(types int) (string, string) {
    now := time.Now()
    currentYear, currentMonth, _ := now.Date()
    currentLocation := now.Location()

    firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
    lastOfMonth := firstOfMonth.AddDate(0, 0, -1)
    f := firstOfMonth.Unix()
    l := lastOfMonth.Unix()
    if types == 1 {
        return time.Unix(l, 0).Format("2006-01-02"), time.Unix(f, 0).Format("2006-01-02")
    }
    return time.Unix(l, 0).Format("2006-01-02") + " 23:59:59", time.Unix(f, 0).Format("2006-01-02") + " 00:00:00"
}

//获取当周的开始时间和结束时间
func GetWeekDay(types int) (string, string) {
    now := time.Now()
    offset := int(time.Monday - now.Weekday())
    //周日做特殊判断 因为time.Monday = 0
    if offset > 0 {
        offset = -6
    }

    lastoffset := int(time.Saturday - now.Weekday())
    //周日做特殊判断 因为time.Monday = 0
    if lastoffset == 6 {
        lastoffset = -1
    }

    firstOfWeek := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
    lastOfWeeK := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, lastoffset+1)
    f := firstOfWeek.Unix()
    l := lastOfWeeK.Unix()
    if types == 1 {
        return time.Unix(f, 0).Format("2006-01-02"), time.Unix(l, 0).Format("2006-01-02")
    }
    return time.Unix(f, 0).Format("2006-01-02") + " 00:00:00", time.Unix(l, 0).Format("2006-01-02") + " 23:59:59"
}

//获得当前季度的初始和结束日期
func GetQuarterDay() (string, string) {
    year := time.Now().Format("2006")
    month := int(time.Now().Month())
    var firstOfQuarter string
    var lastOfQuarter string
    if month >= 1 && month <= 3 {
        //1月1号
        firstOfQuarter = year + "-01-01 00:00:00"
        lastOfQuarter = year + "-03-31 23:59:59"
    } else if month >= 4 && month <= 6 {
        firstOfQuarter = year + "-04-01 00:00:00"
        lastOfQuarter = year + "-06-30 23:59:59"
    } else if month >= 7 && month <= 9 {
        firstOfQuarter = year + "-07-01 00:00:00"
        lastOfQuarter = year + "-09-30 23:59:59"
    } else {
        firstOfQuarter = year + "-10-01 00:00:00"
        lastOfQuarter = year + "-12-31 23:59:59"
    }
    return firstOfQuarter, lastOfQuarter
}

//获取开始日期和结束日期计算出时间段内所有日期
func GetBetweenDates(sdate, edate string) []string {
    d := []string{}
    timeFormatTpl := "2006-01-02 15:04:05"
    if len(timeFormatTpl) != len(sdate) {
        timeFormatTpl = timeFormatTpl[0:len(sdate)]
    }
    date, err := time.Parse(timeFormatTpl, sdate)
    if err != nil {
        // 时间解析，异常
        return d
    }
    date2, err := time.Parse(timeFormatTpl, edate)
    if err != nil {
        // 时间解析，异常
        return d
    }
    if date2.Before(date) {
        // 如果结束时间小于开始时间，异常
        return d
    }
    // 输出日期格式固定
    timeFormatTpl = "2006-01-02"
    date2Str := date2.Format(timeFormatTpl)
    d = append(d, date.Format(timeFormatTpl))
    for {
        date = date.AddDate(0, 0, 1)
        dateStr := date.Format(timeFormatTpl)
        d = append(d, dateStr)
        if dateStr == date2Str {
            break
        }
    }
    return d
}

//获取前一天
func BeforeData(days int) (startTime, endTime int64) {
    dateNow := time.Now().AddDate(0, 0, -days)
    startTime = time.Date(dateNow.Year(), dateNow.Month(), dateNow.Day(), 0, 0, 0, 0, dateNow.Location()).Unix()
    dateNows := time.Now()
    endTime = time.Date(dateNows.Year(), dateNows.Month(), dateNows.Day(), 23, 59, 59, 0, dateNows.Location()).Unix()
    return startTime, endTime
}

func BeforeDataStartAndEnd(days int) (startTime, endTime int64) {
    dateNow := time.Now().AddDate(0, 0, days)
    startTime = time.Date(dateNow.Year(), dateNow.Month(), dateNow.Day(), 0, 0, 0, 0, dateNow.Location()).Unix()
    endTime = time.Date(dateNow.Year(), dateNow.Month(), dateNow.Day(), 23, 59, 59, 0, dateNow.Location()).Unix()
    return startTime, endTime
}

//获取昨天和今天
func YesterDayAndToday(types int) (string, string) {
    ts := time.Now().AddDate(0, 0, -1)
    timeYesterday := time.Date(ts.Year(), ts.Month(), ts.Day(), 0, 0, 0, 0, ts.Location()).Unix()
    timeStr := time.Unix(timeYesterday, 0).Format("2006-01-02")

    t := time.Now()
    addTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
    timeStrTd := addTime.Format("2006-01-02")
    if types == 1 {
        return timeStr, timeStrTd
    }
    return timeStr + " 00:00:00", timeStrTd + " 23:59:59"
}
