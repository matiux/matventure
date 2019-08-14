package entity

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo/common"
)

type Hero struct {
	ecs.BasicEntity
	//common.AnimationComponent
	common.RenderComponent
	common.SpaceComponent
	//component.ControlComponent
	//component.SpeedComponent
}
