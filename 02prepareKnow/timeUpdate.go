//记录所有session的最后报文时间，检查超时并删除节点
//查找session的设备ID，以便写入到数据库
package prepareKnow

import (
	"time"
)

type Session struct {
	DeviceID  string
	TimeStamp time.Time
}

var TimeGroup = make(map[string]*Session)

func updateSessonTime(sessionID string) {
	// TimeGroup := make(map[string]time.Time)
	//如果session不定义成指针就会出错了
	k := time.Now()
	TimeGroup[sessionID].TimeStamp = k
	// tmp := TimeGroup[sessionID]
	// tmp.time = k
	// TimeGroup[sessionID] = tmp
}

func checkSessionAlive() {
	k := time.Now()
	d, _ := time.ParseDuration("-3s")
	k = k.Add(d)
	for i, t := range TimeGroup {
		if t.TimeStamp.Before(k) {
			delete(TimeGroup, i)
		}
	}
}
