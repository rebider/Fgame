package api

import (
	gmError "fgame/fgame/gm/gamegm/error"
	errhttp "fgame/fgame/gm/gamegm/error/utils"
	gmUserService "fgame/fgame/gm/gamegm/gm/user/service"
	gmhttp "fgame/fgame/gm/gamegm/pkg/httputils"
	"net/http"

	userremote "fgame/fgame/gm/gamegm/remote/service"

	monitor "fgame/fgame/gm/gamegm/monitor"

	log "github.com/Sirupsen/logrus"
	"github.com/xozrc/pkg/httputils"
)

type kickOut struct {
	PlatformId int32  `json:"centerPlatformId"`
	ServerId   int32  `json:"centerServerId"`
	PlayerId   string `json:"playerId"`
	Reason     string `json:"reason"`
}

func handleKickOut(rw http.ResponseWriter, req *http.Request) {
	form := &kickOut{}
	err := httputils.Bind(req, form)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("踢出玩家，解析异常")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	service := userremote.UserRemoteServiceInContext(req.Context())
	if service == nil {
		log.Error("踢出玩家，服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	//获得用户信息：这边先从库拿之后再优化
	us := gmUserService.GmUserServiceInContext(req.Context())
	// userid := gmUserService.GmUserIdInContext(req.Context())
	if us == nil {
		log.Error("踢出玩家，用户服务为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	// userInfo, err := us.GetUserInfo(userid)
	// if err != nil {
	// 	log.WithFields(log.Fields{
	// 		"error":  err,
	// 		"userid": userid,
	// 	}).Error("踢出玩家，获取用户信息失败")
	// 	rw.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }

	playerId := changeStringToInt64(form.PlayerId)

	centerService := monitor.CenterServerServiceInContext(req.Context())

	serverid := centerService.GetCenterServerDBId(form.PlatformId, form.ServerId)
	if serverid < 1 {
		log.WithFields(log.Fields{
			"serverid": serverid,
		}).Error("踢出玩家，获得服务器id为空")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = service.KickOut(int32(serverid), playerId, form.Reason)
	if err != nil {
		log.WithFields(log.Fields{
			"error":    err,
			"serverid": serverid,
		}).Error("踢出玩家，踢出玩家异常")
		// rw.WriteHeader(http.StatusInternalServerError)
		// return
		codeErr := gmError.GetError(gmError.ErrorCodeDefaultRemoteUser)
		errhttp.ResponseWithErrorMessage(rw, codeErr, err.Error())
		return
	}

	rr := gmhttp.NewSuccessResult(nil)
	httputils.WriteJSON(rw, http.StatusOK, rr)
}
