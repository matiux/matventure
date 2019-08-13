package entities

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo/common"
)

/*
HUD = Heads-Up display
Potresti chiedere, cosa differenzia un HUD dalle città che abbiamo aggiunto nel tutorial precedente?
Quando avremo dimensioni della mappa più grandi della finestra del gioco, saremo in grado di scorrere la nostra mappa.
Alcuni componenti non vorremo spostarli quando spostiamo la fotocamera.
L'opzione shader HUD lo rende così questi componenti non si muovono mentre fai una panoramica nel mondo di gioco.
*/
type HUD struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}
