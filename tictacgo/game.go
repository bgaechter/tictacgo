package tictacgo

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	ScreenWidth  = 420
	ScreenHeight = 600
	boardSize    = 3
)

type State int

const (
	PlayerMove State = iota
	OpponentMove
	GameOver
	Won
	Quit
)

type Game struct {
	input      *Input
	board      *Board
	boardImage *ebiten.Image
	state      State
}

func NewGame() (*Game, error) {
	g := &Game{
		input: NewInput(),
		state: PlayerMove,
	}
	var err error
	g.board, err = NewBoard(boardSize)
	if err != nil {
		return nil, err
	}
	return g, nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Update() error {
	g.input.Update()

	switch g.state {
	case PlayerMove:
		playerMoveFinished, err := g.board.Update(g.input)
		if err != nil {
			return err
		}
		if playerMoveFinished {
			g.state = OpponentMove
		}
		if g.board.ThreeInARow("X") {
			g.state = Won
		}
	case OpponentMove:
		g.board.OpponentMove()
		if g.board.ThreeInARow("O") {
			g.state = GameOver
		} else {
			g.state = PlayerMove
		}
	case GameOver:
		if g.input.SpacePressed() {
			g.board.Reset()
			g.state = PlayerMove
		}
	case Won:
		if g.input.SpacePressed() {
			g.board.Reset()
			g.state = PlayerMove
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.boardImage == nil {
		g.boardImage = ebiten.NewImage(g.board.Size())
	}
	screen.Fill(color.Black)
	g.board.Draw(g.boardImage)
	op := &ebiten.DrawImageOptions{}
	sw, sh := screen.Bounds().Dx(), screen.Bounds().Dy()
	bw, bh := g.boardImage.Bounds().Dx(), g.boardImage.Bounds().Dy()
	x := (sw - bw) / 2
	y := (sh - bh) / 2
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(g.boardImage, op)
	if g.state == Won {
		textBackground := ebiten.NewImage(ScreenWidth, 170)
		textBackground.Fill(color.White)
		bgOptions := &ebiten.DrawImageOptions{}

		bgOptions.GeoM.Translate(0, ScreenHeight/2-75)

		screen.DrawImage(textBackground, bgOptions)
		text.Draw(screen, "You Win!", mplusBigFont, 20, ScreenHeight/2, color.Black)
		text.Draw(screen, "Press Space to continue", mplusNormalFont, 20, ScreenHeight/2+50, color.Black)
	}
	if g.state == GameOver {
		textBackground := ebiten.NewImage(ScreenWidth, 170)
		textBackground.Fill(color.White)
		bgOptions := &ebiten.DrawImageOptions{}
		bgOptions.GeoM.Translate(float64(0), ScreenHeight/2-75)

		screen.DrawImage(textBackground, bgOptions)
		text.Draw(screen, "You Lost!", mplusBigFont, 20, ScreenHeight/2, color.Black)
		text.Draw(screen, "Press Space to continue", mplusNormalFont, 20, ScreenHeight/2+50, color.Black)
	}
	text.Draw(screen, "Use Arrow Keys to select field", mplusNormalFont, 20, 50, color.White)
	text.Draw(screen, "Press space to mark it", mplusNormalFont, 20, 100, color.White)
}
