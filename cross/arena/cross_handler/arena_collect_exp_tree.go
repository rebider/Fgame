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
	processor.Register(codec.MessageType(crosspb.MessageType_SI_ARENA_COLLECT_EXP_TREE_TYPE), dispatch.HandlerFunc(handleArenaCollectExpTree))
}

//采集经验树成功
func handleArenaCollectExpTree(s session.Session, msg interface{}) (err error) {
	log.Debug("arena:采集经验树成功")
	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(*player.Player)

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("arena:采集经验树成功,完成")
	return nil
}
