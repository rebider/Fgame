/*此类自动生成,请勿修改*/
package model

import logserverlog "fgame/fgame/logserver/log"

func init() {
	logserverlog.RegisterLogMsg((*PlayerAdditionSysTongLing)(nil))
}

/*附加系统通灵*/
type PlayerAdditionSysTongLing struct {
	PlayerLogMsg `bson:",inline"`

	//当前等级
	CurLev int32 `json:"curLev"`

	//变化前等级
	BeforeLev int32 `json:"beforeLev"`

	//进阶原因编号
	Reason int32 `json:"reason"`

	//进阶原因
	ReasonText string `json:"reasonText"`
}

func (c *PlayerAdditionSysTongLing) LogName() string {
	return "player_additionsys_tongling"
}