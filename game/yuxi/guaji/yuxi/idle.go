package gate

import (
	"fgame/fgame/game/battle/battle"

	coreutils "fgame/fgame/core/utils"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	yuxitemplate "fgame/fgame/game/yuxi/template"

	log "github.com/Sirupsen/logrus"
)

func init() {
	scene.RegisterGuaJiActionFactory(scenetypes.GuaJiTypeYuXi, battle.PlayerStateIdle, scene.GuaJiActionFactoryFunc(newIdleAction))
}

type idleAction struct {
	*scene.DummyGuaJiAction
}

const (
	yuxiZhouMaxDistance = 30
)

func (a *idleAction) GuaJi(p scene.Player) {
	s := p.GetScene()
	if s == nil {
		return
	}
	constantTemp := yuxitemplate.GetYuXiTemplateService().GetYuXiConstTemplate()
	yuxiPos := constantTemp.GetYuXiPos()

	//玉玺周围查找敌人
	if coreutils.Distance(p.GetPos(), yuxiPos) <= yuxiZhouMaxDistance {
		//获取玉玺位置
		e := scenelogic.FindHatestEnemy(p)
		if e != nil {
			p.SetAttackTarget(e.BattleObject)
			p.GuaJiTrace()
			return
		}
		//查找默认目标
		bo := p.GetDefaultAttackTarget()
		if bo != nil {
			p.SetAttackTarget(bo)
			p.GuaJiTrace()
			return
		}
	}

	//判断是否正在移动
	if p.IsMove() {
		return
	}
	pos := s.MapTemplate().GetMap().RandomPosition()
	flag := p.SetDestPosition(pos)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": p.GetId(),
			}).Warn("yuxi:挂机找不到路")
		return
	}

	return
}

func (a *idleAction) OnExit() {
}

func newIdleAction() scene.GuaJiAction {
	a := &idleAction{}
	a.DummyGuaJiAction = scene.NewDummyGuaJiAction()
	return a
}
