package found_handler

import (
	activitytypes "fgame/fgame/game/activity/types"
	"fgame/fgame/game/found/found"
	foundtypes "fgame/fgame/game/found/types"
	"fgame/fgame/game/player"
)

func init() {
	found.RegistFoundDataHandler(foundtypes.FoundResourceTypeJiuXiaoChengZhan, found.FoundObjDataHandlerFunc(getAllianceWarFoundParam))
}

func getAllianceWarFoundParam(pl player.Player) (resLevel int32, maxTimes int32, group int32) {
	return getParam(pl, activitytypes.ActivityTypeAlliance)
}

func getParam(pl player.Player, typ activitytypes.ActivityType) (resLevel int32, maxTimes int32, group int32) {
	group = int32(1)
	resLevel = pl.GetLevel()
	maxTimes = 0
	return
}
