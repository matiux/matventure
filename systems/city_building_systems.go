package systems

import (
	"fmt"

	"github.com/EngoEngine/engo"

	"github.com/EngoEngine/ecs"
)

type CityBuildingSystem struct{}

// Remove is called whenever an Entity is removed from the World,
// in order to remove it from this sytem as well
func (*CityBuildingSystem) Remove(ecs.BasicEntity) {}

// Update is ran every frame, with `dt` being
// the time in seconds since the last frame
func (*CityBuildingSystem) Update(dt float32) {

	if engo.Input.Button("AddCity").JustPressed() {
		fmt.Println("The gamer pressed F1")
	}
}

func (*CityBuildingSystem) New(*ecs.World) {
	fmt.Println("CityBuildingSystem was added to the Scene")
}
