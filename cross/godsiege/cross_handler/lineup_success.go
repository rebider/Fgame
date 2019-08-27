package handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/cross/player/player"
	"fgame/fgame/cross/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(crosspb.MessageType_SI_GODSIEGE_LINEUP_SUCCESS_TYPE), dispatch.HandlerFunc(handleGodSiegeLineUpSuccess))
}

//处理神兽攻城排队成功
func handleGodSiegeLineUpSuccess(s session.Session, msg interface{}) (err error) {
	log.Debug("godsiege:处理神兽攻城排队成功")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(*player.Player)

	siGodSiegeLineUpSuccess := msg.(*crosspb.SIGodSiegeLineUpSuccess)
	godType := siGodSiegeLineUpSuccess.GetGodType()

	err = godSiegeLineUpSuccess(tpl, godType)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"godType":  godType,
			}).Error("godsiege:处理神兽攻城排队成功,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("godsiege:处理神兽攻城排队成功,完成")
	return nil
}

//神兽攻城排队成功
func godSiegeLineUpSuccess(pl *player.Player, godType int32) (err error) {
	return
}
