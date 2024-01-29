package tictacgo

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var (
	mplusNormalFont font.Face
	mplusBigFont    font.Face
	jaKanjis        = []rune{}
)

func init() {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		log.Fatal(err)
	}
	mplusBigFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    48,
		DPI:     dpi,
		Hinting: font.HintingFull, // Use quantization to save glyph cache images.
	})
	if err != nil {
		log.Fatal(err)
	}

	// Adjust the line height.
	mplusBigFont = text.FaceWithLineHeight(mplusBigFont, 54)
}

type TileData struct {
	value string
	x     int
	y     int
}

type Tile struct {
	image    *ebiten.Image
	current  TileData
	isActive bool
}

func (t *Tile) Pos() (int, int) {
	return t.current.x, t.current.y
}

func (t *Tile) Value() string {
	return t.current.value
}

func (t *Tile) Update() error {
	// nothing to do?
	return nil
}

const (
	tileSize   = 80
	tileMargin = 4
)

func (t *Tile) Draw(boardImage *ebiten.Image) {
	x := t.current.x*tileSize + (t.current.x+1)*tileMargin + tileSize/2 - 3*tileMargin
	y := t.current.y*tileSize + (t.current.y+1)*tileMargin + tileSize/2 + 3*tileMargin
	text.Draw(boardImage, t.current.value, mplusNormalFont, x, y, color.Black)
}

func NewTile(value string, x int, y int) *Tile {
	image := ebiten.NewImage(tileSize, tileSize)
	image.Fill(color.White)
	return &Tile{
		current: TileData{
			value: value,
			x:     x,
			y:     y,
		},
		image: image,
	}
}
