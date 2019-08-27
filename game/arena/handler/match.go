package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	arenalogic "fgame/fgame/game/arena/logic"
	"fgame/fgame/game/arena/pbutil"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ARENA_MATCH_TYPE), dispatch.HandlerFunc(handleArenaMatch))
}

//处理3v3匹配
func handleArenaMatch(s session.Session, msg interface{}) (err error) {
	log.Debug("arena:处理3v3匹配")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = arenaMatch(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("arena:处理3v3匹配,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("arena:处理3v3匹配,完成")
	return nil

}

//TODO 取消这个接口
//3v3匹配
func arenaMatch(pl player.Player) (err error) {
	//判断活动时间

	//判断是否有组队
	canMatch, err := arenalogic.ArenaMatch(pl)
	if err != nil {
		return
	}
	if canMatch {
		scArenaMatch := pbutil.BuildSCArenaMatch()
		pl.SendMsg(scArenaMatch)
	}
	return
}