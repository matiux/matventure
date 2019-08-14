package main

import (
	"bytes"
	"image"
	"image/color"

	"golang.org/x/image/font/gofont/gosmallcaps"

	"github.com/matiux/matventure/entities"

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

// Ogni Scene ha un solo World e questo World è una raccolta di System ed Entity.
type myScene struct{}

// Type definisce in modo univoco il tuo tipo di gioco
func (*myScene) Type() string { return "myGame" }

/*
Preload viene chiamato prima di caricare qualsiasi risorsa dal disco, per consentire di registrarli / accodarli
Il rendering di qualcosa consiste di tre cose (ECS):
The System (RenderSystem, nello specifico)
The Entity (puoi averne molte)
Two Component (RenderComponent e SpaceComponent).

L'idea è che ci sono molte entità (oggetti). Questi possono avere valori e variabili diversi: formano i diversi componenti che ogni entità può avere.
Le entità che hanno un componente SpaceComponent, ad esempio, hanno alcune informazioni sulla loro posizione (nel gamespace).
Queste entità e componenti non fanno nulla. Sono solo contenitori di dati.
Un componente può avere variabili (e valori per tali variabili) e un'entità è solo un []Component.
Effettivamente fare qualcosa è permesso solo per i System(s). Possono cambiare / aggiungere / rimuovere entità, nonché cambiare / aggiungere / rimuovere qualsiasi componente su tali entità.
Possono modificare valori come i valori di posizione all'interno di SpaceComponent.
Puoi avere più System(s) e ognuno avrà il proprio compito. Ogni frame, questi sistemi sono chiamati in modo che possano fare le cose.
Il RenderSystem è uno dei sistemi inclusi in Engo, che ha già molte chiamate OpenGL integrate.
*/
func (*myScene) Preload() {

	engo.Files.Load("textures/citySheet.png", "textures/city.png", "tilemap/TrafficMap.tmx")
	engo.Files.LoadReaderData("go.ttf", bytes.NewReader(gosmallcaps.TTF))
}

/*
Setup viene chiamato prima dell'avvio del ciclo principale. Ti permette di aggiungere Entity e System alla tua scena.
Per aggiungere il RenderSystem al nostro gioco, dobbiamo aggiungerlo all'interno della funzione Setup della nostra scena.
Nello specifico, un'istanza di RenderSystem viene aggiunta al mondo di questa scena specifica.
*/
func (*myScene) Setup(u engo.Updater) {

	world, _ := u.(*ecs.World)

	common.SetBackground(color.White)

	/*
		Da notare che abbiamo aggiunto gli altri sistemi prima del CityBuildingSystem. Questo per garantire che tutti i sistemi dai quali potremmo
		dipendere siano già inizializzati durante l'inizializzazione del CityBuildingSystem.
	*/
	world.AddSystem(&common.RenderSystem{})

	/*
		Per generare le Entity nel posto corretto (Sopra il mouse), dobbiamo sapere dove si trova il cursore. La nostra prima ipotesi potrebbe essere
		quella di utilizzare la struttura engo.Input.Mouse che è disponibile. Tuttavia, questo restituisce la posizione effettiva (X, Y) relativa alla
		dimensione dello schermo, non al sistema di griglie di gioco. Abbiamo uno speciale MouseSystem disponibile proprio per questo

		Il MouseSystem è principalmente scritto per tenere traccia degli eventi del mouse per le Entity; puoi verificare se la tua Entità è stata spostata,
		cliccata, trascinata, ecc. Per usarla, abbiamo quindi bisogno di una Entity che utilizza il MouseSystem. Questo deve contenere un componente Mouse,
		in cui verranno salvati i risultati / i dati.
	*/
	world.AddSystem(&common.MouseSystem{})

	kbs := common.NewKeyboardScroller(KeyboardScrollSpeed, engo.DefaultHorizontalAxis, engo.DefaultVerticalAxis)
	world.AddSystem(kbs)
	//world.AddSystem(&common.EdgeScroller{EdgeScrollSpeed, EdgeWidth})
	//world.AddSystem(&common.MouseZoomer{ZoomSpeed})

	world.AddSystem(&systems.CityBuildingSystem{})
	world.AddSystem(&systems.HUDTextSystem{})
	world.AddSystem(&systems.MoneySystem{})

	/*
		Poiché stiamo costruendo un HUD, non ci interessa quanto sia grande la nostra mappa di gioco.
		Siamo interessati solo a quanto è grande la finestra stessa. Dovremo usare WindowHeight() e sottrarre l'altezza del nostro HUD.
	*/
	hud := entities.HUD{BasicEntity: ecs.NewBasic()}
	hud.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{0, engo.WindowHeight() - 100},
		Width:    200,
		Height:   200,
	}

	hudImage := image.NewUniform(color.RGBA{205, 205, 205, 255})
	hudNRGBA := common.ImageToNRGBA(hudImage, 200, 200)
	hudImageObj := common.NewImageObject(hudNRGBA)
	hudTexture := common.NewTextureSingle(hudImageObj)

	hud.RenderComponent = common.RenderComponent{
		Drawable: hudTexture,
		Scale:    engo.Point{1, 1},
		Repeat:   common.Repeat,
	}

	hud.RenderComponent.SetShader(common.HUDShader)
	hud.RenderComponent.SetZIndex(1)

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

	tiles := make([]*entities.Tile, 0)
	for _, tileLayer := range levelData.TileLayers {

		for _, tileElement := range tileLayer.Tiles {

			if tileElement.Image != nil {

				tile := &entities.Tile{BasicEntity: ecs.NewBasic()}
				tile.RenderComponent = common.RenderComponent{
					Scale:    engo.Point{1, 1},
					Drawable: tileElement,
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
		Title:          "TrafficManager",
		Width:          800,
		Height:         800,
		StandardInputs: true,
	}

	engo.Run(opts, &myScene{})
}
