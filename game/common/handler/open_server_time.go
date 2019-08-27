package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/common/pbutil"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/welfare/welfare"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_OPEN_SERVER_TIME_TYPE), dispatch.HandlerFunc(handleGetOpenServerTime))
}

//获取开服时间
func handleGetOpenServerTime(s session.Session, msg interface{}) error {
	log.Debug("common:处理获取开服时间信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	openServerTime := global.GetGame().GetServerTime()
	activityOpenServerTime := welfare.GetWelfareService().GetServerStartTime() //global.GetGame().GetServerTime()
	scOpenServerTime := pbutil.BuildSCOpenServerTime(openServerTime, activityOpenServerTime)
	err := tpl.SendMsg(scOpenServerTime)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("common:处理获取开服时间信息,错误")
		return err
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("common:处理获取开服时间信息完成")
	return nil
}
