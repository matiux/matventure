package systems

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/matiux/matventure/entities"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

// Spritesheet contains the sprites for the city buildings, cars, and roads
var Spritesheet *common.Spritesheet

var cities = [...][12]int{
	{
		99, 100, 101,
		454, 269, 455,
		415, 195, 416,
		452, 306, 453,
	},
	{
		99, 100, 101,
		268, 269, 270,
		268, 269, 270,
		305, 306, 307,
	},
	{
		75, 76, 77,
		446, 261, 447,
		446, 261, 447,
		444, 298, 445,
	},
	{
		75, 76, 77,
		407, 187, 408,
		407, 187, 408,
		444, 298, 445,
	},
	{
		75, 76, 77,
		186, 150, 188,
		186, 150, 188,
		297, 191, 299,
	},
	{
		83, 84, 85,
		413, 228, 414,
		411, 191, 412,
		448, 302, 449,
	},
	{
		83, 84, 85,
		227, 228, 229,
		190, 191, 192,
		301, 302, 303,
	},
	{
		91, 92, 93,
		241, 242, 243,
		278, 279, 280,
		945, 946, 947,
	},
	{
		91, 92, 93,
		241, 242, 243,
		278, 279, 280,
		945, 803, 947,
	},
	{
		91, 92, 93,
		238, 239, 240,
		238, 239, 240,
		312, 313, 314,
	},
}

// type MouseTracker struct {
// 	ecs.BasicEntity
// 	common.MouseComponent
// }

type CityBuildingSystem struct {
	world *ecs.World
	// mouseTracker MouseTracker

	usedTiles          []int
	elapsed, buildTime float32
	built              int
}

// Remove is called whenever an Entity is removed from the World,
// in order to remove it from this sytem as well
func (*CityBuildingSystem) Remove(ecs.BasicEntity) {}

/*
La funzione Update del nostro System CityBuilding viene chiamata ogni frame dal mondo
`dt` è il tempo in secondi dall'ultimo fotogramma
*/

func (cb *CityBuildingSystem) Update(dt float32) {

	cb.elapsed += dt

	if cb.elapsed >= cb.buildTime {
		cb.generateCity()
		cb.elapsed = 0
		cb.updateBuildTime()
		cb.built++
	}
}

func (cb *CityBuildingSystem) New(w *ecs.World) {

	cb.world = w
	rand.Seed(time.Now().UnixNano())

	fmt.Println("CityBuildingSystem was added to the Scene")

	Spritesheet = common.NewSpritesheetWithBorderFromFile("textures/citySheet.png", 16, 16, 1, 1)

	cb.updateBuildTime()
}

func (cb *CityBuildingSystem) isTileUsed(tile int) bool {

	for _, t := range cb.usedTiles {
		if tile == t {
			return true
		}
	}

	return false
}

func (cb *CityBuildingSystem) generateCity() {

	x := rand.Intn(18)
	y := rand.Intn(18)
	t := x + y*18

	for cb.isTileUsed(t) {

		fmt.Printf("Tile usata: %v\n", t)

		if len(cb.usedTiles) > 300 {
			break //to avoid infinite loop
		}

		x = rand.Intn(18)
		y = rand.Intn(18)
		t = x + y*18
	}

	cb.usedTiles = append(cb.usedTiles, t)

	city := rand.Intn(len(cities))
	cityTiles := make([]*entities.City, 0)

	fmt.Printf("x: %v | y: %v | tile: %v | city: %v, \n", x, y, t, city)

	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {

			tile := &entities.City{BasicEntity: ecs.NewBasic()}

			x2 := float32(((x+1)*64)+8) + float32(i*16)
			y2 := float32((y+1)*64) + float32(j*16)

			tile.SpaceComponent.Position = engo.Point{
				X: x2,
				Y: y2,
			}

			index := cities[city][i+3*j]

			tile.RenderComponent.Drawable = Spritesheet.Cell(index)
			tile.RenderComponent.SetZIndex(1)
			cityTiles = append(cityTiles, tile)

			fmt.Printf("\t\tx2: %v | y2: %v | index: %v, \n", x2, y2, index)
		}
	}

	for _, system := range cb.world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			for _, v := range cityTiles {
				sys.Add(&v.BasicEntity, &v.RenderComponent, &v.SpaceComponent)
			}
		}
	}
}

func (cb *CityBuildingSystem) updateBuildTime() {
	switch {
	case cb.built < 2:
		cb.buildTime = 1*rand.Float32() + 3 // 10 to 15 seconds
	case cb.built < 8:
		cb.buildTime = 30*rand.Float32() + 60 // 60 to 90 seconds
	case cb.built < 18:
		cb.buildTime = 60*rand.Float32() + 30 // 30 to 90 seconds
	case cb.built < 28:
		cb.buildTime = 35*rand.Float32() + 30 // 30 to 65 seconds
	case cb.built < 33:
		cb.buildTime = 30*rand.Float32() + 30 // 30 to 60 seconds
	default:
		cb.buildTime = 20*rand.Float32() + 20 // 20 to 40 seconds
	}
}
