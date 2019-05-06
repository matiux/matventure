package systems

import (
	"fmt"
	"log"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

// Entity
type City struct {
	ecs.BasicEntity
	common.RenderComponent // holds information about what to render (i.e. which texture)
	common.SpaceComponent  // holds information about where it should be rendered.
}

type MouseTracker struct {
	ecs.BasicEntity
	common.MouseComponent
}

type CityBuildingSystem struct {
	world *ecs.World

	mouseTracker MouseTracker
}

// Remove is called whenever an Entity is removed from the World,
// in order to remove it from this sytem as well
func (*CityBuildingSystem) Remove(ecs.BasicEntity) {}

// Update is ran every frame, with `dt` being
// the time in seconds since the last frame
func (cb *CityBuildingSystem) Update(dt float32) {
	if engo.Input.Button("AddCity").JustPressed() {
		fmt.Println("The gamer pressed F1")

		// Entity
		city := City{BasicEntity: ecs.NewBasic()}

		// Entity - SpaceComponent
		city.SpaceComponent = common.SpaceComponent{
			Position: engo.Point{cb.mouseTracker.MouseX, cb.mouseTracker.MouseY},
			Width:    30,
			Height:   64,
		}

		// Entity - RenderComponent
		texture, err := common.LoadedSprite("textures/city.png")
		if err != nil {
			log.Println("Unable to load texture: " + err.Error())
		}

		city.RenderComponent = common.RenderComponent{
			Drawable: texture,
			Scale:    engo.Point{0.1, 0.1},
		}

		// Aggiungo l'Entity al System
		for _, system := range cb.world.Systems() {
			switch sys := system.(type) {
			case *common.RenderSystem:
				// We’re using the RenderSystem-specific Add method to add our City to that system.
				sys.Add(&city.BasicEntity, &city.RenderComponent, &city.SpaceComponent)
			}
		}
	}
}

func (cb *CityBuildingSystem) New(w *ecs.World) {

	cb.world = w

	fmt.Println("CityBuildingSystem was added to the Scene")

	cb.mouseTracker.BasicEntity = ecs.NewBasic()
	cb.mouseTracker.MouseComponent = common.MouseComponent{Track: true}

	engo.Input.RegisterButton("AddCity", engo.KeyF1)

	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.MouseSystem:
			sys.Add(&cb.mouseTracker.BasicEntity, &cb.mouseTracker.MouseComponent, nil, nil)
		}
	}
}
