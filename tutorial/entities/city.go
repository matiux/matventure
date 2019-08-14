package entities

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo/common"
)

/*
Questa struttura (Entity), consiste in una cosa standard (ecs.BasicEntity, che fornisce un identificatore univoco) e due componenti:
sono l'unico modo in cui puoi passare informazioni su diversi sistemi, come dire al RenderSystem cosa rendere.
RenderComponent contiene informazioni su cosa rendere (cioè quale texture), SpaceComponent contiene informazioni su dove dovrebbe essere eseguito il rendering.

Per creare correttamente un'istanza, è necessario assicurarsi che ecs.BasicEntity sia impostato su un nuovo identificatore univoco.
Possiamo farlo chiamando ecs.NewBasic() nella nostra funzione di installazione.
*/
type City struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}
