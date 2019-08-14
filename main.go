package main

import (
	"image/color"

	"github.com/matiux/matventure/entity"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

var (
	heroWidth  = 52
	heroHeight = 73
)

// Ogni Scene ha un solo World e questo World è una raccolta di System ed Entity.
type MainScene struct{}

// Type definisce in modo univoco il tuo tipo di gioco
func (*MainScene) Type() string { return "MainScene" }

func (*MainScene) Preload() {

	engo.Files.Load("world.tmx", "hero.png")
}

func (scene *MainScene) Setup(u engo.Updater) {

	world, _ := u.(*ecs.World)
	common.SetBackground(color.White)

	world.AddSystem(&common.RenderSystem{})

	scene.setupHero(world)

	//resource, err := engo.Files.Resource("world.tmx")
	//if err != nil {
	//	panic(err)
	//}
	//
	//tmxResource := resource.(common.TMXResource)
	//levelData := tmxResource.Level
	//
	//tiles := make([]*entity.Tile, 0)
	//
	//for _, tileLayer := range levelData.TileLayers {
	//
	//	name := tileLayer.Name
	//	fmt.Println("Layer: " + name)
	//	//fmt.Println(tileLayer.Properties)
	//
	//	for _, tileElement := range tileLayer.Tiles {
	//
	//		if tileElement.Image != nil {
	//
	//			tile := &entity.Tile{BasicEntity: ecs.NewBasic()}
	//
	//			tile.RenderComponent = common.RenderComponent{
	//				Scale:    engo.Point{1, 1},
	//				Drawable: tileElement,
	//			}
	//
	//			tile.SpaceComponent = common.SpaceComponent{
	//				Position: tileElement.Point,
	//				Width:    0,
	//				Height:   0,
	//			}
	//
	//			if tileLayer.Name == "world" {
	//				tile.RenderComponent.SetZIndex(0)
	//			}
	//
	//			tiles = append(tiles, tile)
	//		}
	//	}
	//
	//	break
	//}
	//
	//for _, system := range world.Systems() {
	//	switch sys := system.(type) {
	//	case *common.RenderSystem:
	//		for _, v := range tiles {
	//			sys.Add(&v.BasicEntity, &v.RenderComponent, &v.SpaceComponent)
	//		}
	//	}
	//}
	//common.CameraBounds = levelData.Bounds()
}

func (scene *MainScene) setupHero(world *ecs.World) {

	spriteSheet := common.NewSpritesheetFromFile("hero.png", heroWidth, heroHeight)

	x := engo.GameWidth() / 2
	y := engo.GameHeight() / 2

	hero := scene.createHero(
		engo.Point{x, y},
		spriteSheet,
	)

	hero.RenderComponent.SetZIndex(1)

	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(
				&hero.BasicEntity,
				&hero.RenderComponent,
				&hero.SpaceComponent,
			)
		}
	}
}

func (scene *MainScene) createHero(point engo.Point, spriteSheet *common.Spritesheet) *entity.Hero {

	hero := &entity.Hero{BasicEntity: ecs.NewBasic()}

	hero.SpaceComponent = common.SpaceComponent{
		Position: point,
		Width:    float32(heroWidth),
		Height:   float32(heroHeight),
	}

	hero.RenderComponent = common.RenderComponent{
		Drawable: spriteSheet.Cell(0),
		Scale:    engo.Point{1, 1},
	}

	//hero.SpeedComponent = SpeedComponent{}
	//hero.AnimationComponent = common.NewAnimationComponent(spriteSheet.Drawables(), 0.1)

	//hero.AnimationComponent.AddAnimations(actions)
	//hero.AnimationComponent.SelectAnimationByName("downstop")

	return hero
}

func main() {
	opts := engo.RunOptions{
		Title:          "Matventure",
		Width:          1024,
		Height:         768,
		StandardInputs: true,
	}
	engo.Run(opts, new(MainScene))
}
