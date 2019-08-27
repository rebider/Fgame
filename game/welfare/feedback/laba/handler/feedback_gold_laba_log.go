package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/welfare/pbutil"
	"fgame/fgame/game/welfare/welfare"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_OPEN_ACTIVITY_LABA_LOG_INCR_TYPE), dispatch.HandlerFunc(handleLabaLogIncr))
}

//处理元宝拉霸日志请求
func handleLabaLogIncr(s session.Session, msg interface{}) (err error) {
	log.Debug("laba:处理获取元宝拉霸日志请求")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSOpenActivityLaBaLogIncr)
	groupId := csMsg.GetGroupId()
	logTime := csMsg.GetLogTime()

	err = labaLogIncr(tpl, groupId, logTime)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"logTime":  logTime,
				"error":    err,
			}).Error("laba:处理获取元宝拉霸日志请求,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"logTime":  logTime,
		}).Debug("laba:处理获取元宝拉霸日志请求完成")
	return nil

}

//获取元宝拉霸界面信息逻辑
func labaLogIncr(pl player.Player, groupId int32, logTime int64) (err error) {
	logList := welfare.GetWelfareService().GetLaBaLogByTime(groupId, logTime)
	if len(logList) < 1 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"logTime":  logTime,
			}).Info("laba:处理获取元宝拉霸日志请求,日志增量列表为空")
	}

	scMsg := pbutil.BuildSCOpenActivityLaBaLogIncr(logList, groupId)
	pl.SendMsg(scMsg)
	return
}
