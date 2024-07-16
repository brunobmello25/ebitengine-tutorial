package main

import (
	"fmt"
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Sprite struct {
	Img  *ebiten.Image
	X, Y float64
}

type Player struct {
	*Sprite
	Health uint
}

type Enemy struct {
	*Sprite
	FollowsPlayer bool
}

type Potion struct {
	*Sprite
	AmountHeal uint
}

type Game struct {
	player  *Player
	enemies []*Enemy
	potions []*Potion
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.player.X += 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.player.X -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.player.Y -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.player.Y += 2
	}

	for _, enemy := range g.enemies {
		if !enemy.FollowsPlayer {
			continue
		}

		if enemy.X < g.player.X {
			enemy.X += 1
		} else if enemy.X > g.player.X {
			enemy.X -= 1
		}

		if enemy.Y < g.player.Y {
			enemy.Y += 1
		} else if enemy.Y > g.player.Y {
			enemy.Y -= 1
		}
	}

	for _, potion := range g.potions {
		if g.player.X > potion.X {
			g.player.Health += potion.AmountHeal
			fmt.Printf("Picked up potion! Health: %d\n", g.player.Health)
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{120, 180, 255, 255})

	opts := ebiten.DrawImageOptions{}

	opts.GeoM.Translate(g.player.X, g.player.Y)
	screen.DrawImage(
		g.player.Img.SubImage(
			image.Rect(0, 0, 16, 16),
		).(*ebiten.Image),
		&opts,
	)
	opts.GeoM.Reset()

	for _, enemy := range g.enemies {
		opts.GeoM.Translate(enemy.X, enemy.Y)

		screen.DrawImage(
			enemy.Img.SubImage(
				image.Rect(0, 0, 16, 16),
			).(*ebiten.Image),
			&opts,
		)

		opts.GeoM.Reset()
	}

	for _, potion := range g.potions {
		opts.GeoM.Translate(potion.X, potion.Y)

		screen.DrawImage(
			potion.Img.SubImage(
				image.Rect(0, 0, 16, 16),
			).(*ebiten.Image),
			&opts,
		)

		opts.GeoM.Reset()
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	playerImg, _, err := ebitenutil.NewImageFromFile("assets/images/ninja.png")
	if err != nil {
		log.Fatal(err)
	}

	skeletonImg, _, err := ebitenutil.NewImageFromFile("assets/images/skeleton.png")
	if err != nil {
		log.Fatal(err)
	}

	potionImg, _, err := ebitenutil.NewImageFromFile("assets/images/potion.png")
	if err != nil {
		log.Fatal(err)
	}

	game := Game{
		player: &Player{
			Sprite: &Sprite{
				Img: playerImg,
				X:   50.0,
				Y:   50.0,
			},
		},
		potions: []*Potion{
			{
				Sprite: &Sprite{
					Img: potionImg,
					X:   210.0,
					Y:   100.0,
				},
				AmountHeal: 10,
			},
		},
		enemies: []*Enemy{
			{
				Sprite: &Sprite{
					Img: skeletonImg,
					X:   100.0,
					Y:   100.0,
				},
				FollowsPlayer: true,
			},
			{
				Sprite: &Sprite{
					Img: skeletonImg,
					X:   150.0,
					Y:   150.0,
				},
				FollowsPlayer: true,
			},
			{
				Sprite: &Sprite{
					Img: skeletonImg,
					X:   75.0,
					Y:   75.0,
				},
				FollowsPlayer: false,
			},
		},
	}

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
