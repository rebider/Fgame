package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/foe/pbutil"
	playerfoe "fgame/fgame/game/foe/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_FOE_FEEDBACK_READ_TYPE), dispatch.HandlerFunc(handleFoeFeedbackRead))
}

//处理反馈查询
func handleFoeFeedbackRead(s session.Session, msg interface{}) (err error) {
	log.Debug("friend:处理反馈查询")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = foeFeedbackRead(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("friend:处理反馈查询,错误")
		return err
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("friend:处理反馈查询,完成")
	return nil
}

//处理反馈查询
func foeFeedbackRead(pl player.Player) (err error) {
	manager := pl.GetPlayerDataManager(playertypes.PlayerFoeDataManagerType).(*playerfoe.PlayerFoeDataManager)
	manager.ReadFeedback()

	feedbackList := manager.GetFoeFeedbackList()
	scMsg := pbutil.BuildSCFoeFeedbackRead(feedbackList)
	pl.SendMsg(scMsg)

	return
}