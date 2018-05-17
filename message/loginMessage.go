package message

import (
	//. "project/elangshen/main"
	"time"
)

/*登陆信息包

终端 ID 8
类型识别码 2 //两个字节前三位代表机种，后一位代表分支。如GT700，类型识别码是0x70 0x00
时区语言 2  //最低一位是8则西时区，0为东时区，然后右移四位后然后除以100得到时区

*/

//TODO 现在的情况在mqtt支持下，如果直接tcp连接设备，这里需要考虑device_id的获取
type LoginMessage struct {
	DeviceId   string
	DeviceType string
	A_time     time.Time
}

func (r *LoginMessage) DealPackage(deviceId string, messageStr []string) {
	//TODO 处理设备ID
	r.DeviceId = deviceId
	//TODO 处理设备型号
	//TODO 处理时区
	r.A_time = time.Now()
	//更新数据到数据库

}
