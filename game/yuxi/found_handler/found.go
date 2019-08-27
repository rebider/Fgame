package found_handler

import (
	"fgame/fgame/game/found/found"
	foundtypes "fgame/fgame/game/found/types"
	"fgame/fgame/game/player"
)

func init() {
	found.RegistFoundDataHandler(foundtypes.FoundResourceTypePeachParty, found.FoundObjDataHandlerFunc(getYuXiFoundParam))
}

func getYuXiFoundParam(pl player.Player) (resLevel int32, maxTimes int32, group int32) {
	group = int32(1)
	resLevel = pl.GetLevel()
	maxTimes = 0
	return
}
