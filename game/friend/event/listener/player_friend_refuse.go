package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	friendeventtypes "fgame/fgame/game/friend/event/types"
	"fgame/fgame/game/friend/pbutil"
	"fgame/fgame/game/player"
)

func playerFriendInviteRefuse(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	friendId, ok := data.(int64)
	if !ok {
		return
	}

	scFriendAdd := pbutil.BuildSCFriendAdd(friendId, 0, nil, false)
	pl.SendMsg(scFriendAdd)

	fr := player.GetOnlinePlayerManager().GetPlayerById(friendId)
	if fr == nil {
		return
	}
	name := pl.GetName()
	scFriendInviteRefusePushPeer := pbutil.BuildSCFriendInviteRefusePushPeer(pl.GetId(), name)
	fr.SendMsg(scFriendInviteRefusePushPeer)
	return
}

func init() {
	gameevent.AddEventListener(friendeventtypes.EventTypeFriendAddRefuse, event.EventListenerFunc(playerFriendInviteRefuse))
}
