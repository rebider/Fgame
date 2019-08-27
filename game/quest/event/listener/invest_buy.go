package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
	newleveleventtypes "fgame/fgame/game/welfare/invest/new_level/event/types"
)

//投资计划购买
func investNewLevelBuy(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	err = questlogic.SetQuestData(pl, questtypes.QuestSubTypeBuyInvest, 0, 1)
	return
}

func init() {
	gameevent.AddEventListener(newleveleventtypes.EventTypeInvestNewLevelBuy, event.EventListenerFunc(investNewLevelBuy))
}
