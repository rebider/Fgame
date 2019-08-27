/*此类自动生成,请勿修改*/
package model

import logserverlog "fgame/fgame/logserver/log"

func init() {
	logserverlog.RegisterLogMsg((*PlayerShiHunFan)(nil))
}

/*噬魂幡*/
type PlayerShiHunFan struct {
	PlayerLogMsg `bson:",inline"`

	//当前阶数
	CurAdvancedNum int32 `json:"curAdvancedNum"`

	//变化前阶数
	BeforeAdvancedNum int32 `json:"beforeAdvancedNum"`

	//进阶阶数
	AdvancedNum int32 `json:"advancedNum"`

	//进阶原因编号
	Reason int32 `json:"reason"`

	//进阶原因
	ReasonText string `json:"reasonText"`
}

func (c *PlayerShiHunFan) LogName() string {
	return "player_shihunfan"
}
