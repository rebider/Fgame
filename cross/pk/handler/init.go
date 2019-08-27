package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	crosscodec "fgame/fgame/cross/codec"
)

func init() {
	initCodec()

}

func initCodec() {
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_CS_PK_STATE_SWITCH_TYPE), (*uipb.CSPkStateSwitch)(nil))
	crosscodec.RegisterMsg(codec.MessageType(uipb.MessageType_SC_PK_STATE_SWITCH_TYPE), (*uipb.SCPkStateSwitch)(nil))
}
