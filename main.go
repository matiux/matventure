package main

import (
	"image/color"

	"github.com/matiux/matventure/systems"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

/**
Scenes are the backbone of engo: they contain pretty much all other things within the game.
You should have multiple scenes and switch between them: e.g. one for the main menu, another for the loading screen,
and yet another for the in-game experience. Specifically, a Scene contains one World, a collection
of multiple Systems and a magnitude of Entitys
*/

type City struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

type myScene struct{}

// Type uniquely defines your game type
func (*myScene) Type() string { return "myGame" }

// Preload is called before loading any assets from the disk,
// to allow you to register / queue them
func (*myScene) Preload() {
	engo.Files.Load("textures/city.png")
}

// Setup is called before the main loop starts. It allows you
// to add entities and systems to your Scene.
func (*myScene) Setup(u engo.Updater) {

	engo.Input.RegisterButton("AddCity", engo.KeyF1)

	common.SetBackground(color.White)

	world, _ := u.(*ecs.World)
	world.AddSystem(&common.RenderSystem{})
	world.AddSystem(&systems.CityBuildingSystem{})

	//city := City{BasicEntity: ecs.NewBasic()}
	//city.SpaceComponent = common.SpaceComponent{
	//	Position: engo.Point{10, 10},
	//	Width:    303,
	//	Height:   641,
	//}
	//
	//texture, err := common.LoadedSprite("textures/city.png")
	//
	//if err != nil {
	//	log.Println("Unable to load texture: " + err.Error())
	//}
	//
	//city.RenderComponent = common.RenderComponent{
	//	Drawable: texture,
	//	Scale:    engo.Point{1, 1},
	//}
	//
	//for _, system := range world.Systems() {
	//	switch sys := system.(type) {
	//	case *common.RenderSystem:
	//		sys.Add(&city.BasicEntity, &city.RenderComponent, &city.SpaceComponent)
	//	}
	//}
}

func main() {
	opts := engo.RunOptions{
		Title:  "Hello World",
		Width:  1024,
		Height: 768,
	}
	engo.Run(opts, &myScene{})
}
