package main

import (
	"fmt"
	"image"
	"image/color"

	"github.com/matiux/matventure/systems"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

const (
	KeyboardScrollSpeed = 400
	EdgeScrollSpeed     = KeyboardScrollSpeed
	EdgeWidth           = 20
	ZoomSpeed           = -0.125
)

// Ogni Scene ha un solo World. E questo World è una raccolta di System ed Entity.
type myScene struct{}

type HUD struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

type Tile struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

// Type definisce in modo univoco il tuo tipo di gioco
func (*myScene) Type() string { return "myGame" }

// Preload viene chiamato prima di caricare qualsiasi risorsa dal disco,
// per consentire di registrarli / accodarli
func (*myScene) Preload() {

	engo.Files.Load("textures/city.png", "tilemap/TrafficMap.tmx")
}

// Setup viene chiamato prima dell'avvio del ciclo principale. Ti permette di aggiungere Entity e System alla tua scena.
func (*myScene) Setup(u engo.Updater) {
	world, _ := u.(*ecs.World)
	common.SetBackground(color.White)

	// An instance of the RenderSystem is added to the World of this specific Scene.
	world.AddSystem(&common.RenderSystem{})
	world.AddSystem(&common.MouseSystem{})

	kbs := common.NewKeyboardScroller(KeyboardScrollSpeed, engo.DefaultHorizontalAxis, engo.DefaultVerticalAxis)
	world.AddSystem(kbs)
	//world.AddSystem(&common.EdgeScroller{EdgeScrollSpeed, EdgeWidth})
	world.AddSystem(&common.MouseZoomer{ZoomSpeed})

	// An instance of the CityBuildingSystem is added to the World of this specific Scene.
	world.AddSystem(&systems.CityBuildingSystem{})

	hud := HUD{BasicEntity: ecs.NewBasic()}
	hud.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{0, engo.WindowHeight() - 40},
		Width:    200,
		Height:   40,
	}

	hudImage := image.NewUniform(color.RGBA{205, 205, 205, 255})
	hudNRGBA := common.ImageToNRGBA(hudImage, 200, 40)
	hudImageObj := common.NewImageObject(hudNRGBA)
	hudTexture := common.NewTextureSingle(hudImageObj)

	hud.RenderComponent = common.RenderComponent{
		Drawable: hudTexture,
		Scale:    engo.Point{1, 1},
		Repeat:   common.Repeat,
	}

	hud.RenderComponent.SetShader(common.HUDShader)
	hud.RenderComponent.SetZIndex(1)

	// And finally add it to the world:
	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&hud.BasicEntity, &hud.RenderComponent, &hud.SpaceComponent)
		}
	}

	resource, err := engo.Files.Resource("tilemap/TrafficMap.tmx")
	if err != nil {
		panic(err)
	}
	tmxResource := resource.(common.TMXResource)
	levelData := tmxResource.Level

	tiles := make([]*Tile, 0)

	for _, tileLayer := range levelData.TileLayers {

		fmt.Println(fmt.Sprintf("Carico la mappa: %s.", tileLayer.Name))

		for _, tileElement := range tileLayer.Tiles {

			fmt.Println(fmt.Sprintf("Tile in posizione X: %f - Y %f", tileElement.X, tileElement.Y))

			if tileElement.Image != nil {
				tile := &Tile{BasicEntity: ecs.NewBasic()}
				tile.RenderComponent = common.RenderComponent{
					Drawable: tileElement,
					Scale:    engo.Point{1, 1},
				}
				tile.SpaceComponent = common.SpaceComponent{
					Position: tileElement.Point,
					Width:    0,
					Height:   0,
				}
				tiles = append(tiles, tile)
			}
		}
	}

	// add the tiles to the RenderSystem
	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			for _, v := range tiles {
				sys.Add(&v.BasicEntity, &v.RenderComponent, &v.SpaceComponent)
			}
		}
	}

	common.CameraBounds = levelData.Bounds()
}

func main() {
	opts := engo.RunOptions{
		Title:          "Hello World",
		Width:          800,
		Height:         641,
		StandardInputs: true,
	}
	engo.Run(opts, &myScene{})
}
