package tictacgo

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Board struct {
	tiles      map[*Tile]struct{}
	activeTile *Tile
	size       int
}

func NewBoard(size int) (*Board, error) {
	b := &Board{
		size:  size,
		tiles: map[*Tile]struct{}{},
	}
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			t := NewTile("", x, y)
			b.tiles[t] = struct{}{}
		}
	}

	t := tileAt(b.tiles, 0, 0)
	t.isActive = true
	b.activeTile = t

	return b, nil
}

func tileAt(tiles map[*Tile]struct{}, x, y int) *Tile {
	var result *Tile
	for t := range tiles {
		if t.current.x != x || t.current.y != y {
			continue
		}
		if result != nil {
			panic("no such tile")
		}
		result = t
	}
	return result
}

func (b *Board) tileAt(x, y int) *Tile {
	return tileAt(b.tiles, x, y)
}

func (b *Board) Update(input *Input) (bool, error) {
	for t := range b.tiles {
		if err := t.Update(); err != nil {
			return false, err
		}
	}
	if dir, ok := input.Dir(); ok {
		if err := b.Move(dir); err != nil {
			return false, err
		}
	}
	if pressed := input.SpacePressed(); pressed {
		b.MarkTile()
		return true, nil
	}
	return false, nil
}

func (b *Board) Move(dir Dir) error {
	x, y := dir.Vector()
	new_x := b.activeTile.current.x + x
	new_y := b.activeTile.current.y + y
	if new_x < b.size && new_x >= 0 && new_y < b.size && new_y >= 0 {
		b.activeTile.isActive = false
		b.activeTile = b.tileAt(new_x, new_y)
		b.activeTile.isActive = true
	}

	return nil
}

func (b *Board) ThreeInARow(val string) bool {
	if b.tileAt(2, 0).current.value == val && b.tileAt(2, 1).current.value == val && b.tileAt(2, 2).current.value == val {
		return true
	}
	if b.tileAt(1, 0).current.value == val && b.tileAt(1, 1).current.value == val && b.tileAt(1, 2).current.value == val {
		return true
	}
	if b.tileAt(0, 0).current.value == val && b.tileAt(0, 1).current.value == val && b.tileAt(0, 2).current.value == val {
		return true
	}
	if b.tileAt(0, 0).current.value == val && b.tileAt(1, 1).current.value == val && b.tileAt(2, 2).current.value == val {
		return true
	}
	if b.tileAt(0, 0).current.value == val && b.tileAt(1, 0).current.value == val && b.tileAt(2, 0).current.value == val {
		return true
	}
	if b.tileAt(0, 1).current.value == val && b.tileAt(1, 1).current.value == val && b.tileAt(2, 1).current.value == val {
		return true
	}
	if b.tileAt(0, 2).current.value == val && b.tileAt(1, 2).current.value == val && b.tileAt(2, 2).current.value == val {
		return true
	}
	return false
}

func (b *Board) MarkTile() {
	if b.activeTile.current.value != "X" && b.activeTile.current.value != "O" {
		b.activeTile.current.value = "X"
	}
}

func (b *Board) GameWon() {
	b.Reset()
}

func (b *Board) GameLost() {
	b.Reset()
}

func (b *Board) Reset() {
	for t := range b.tiles {
		t.current.value = ""
	}
}

func (b *Board) OpponentMove() {
	var tiles []*Tile
	for t := range b.tiles {
		if t.current.value != "X" && t.current.value != "O" {
			tiles = append(tiles, t)
		}
	}

	if len(tiles) > 0 {
		tiles[0].current.value = "O"
	}
}

func (b *Board) Size() (int, int) {
	x := b.size*tileSize + (b.size+1)*tileMargin
	y := x
	return x, y // Board is a square
}

func (b *Board) Draw(boardImage *ebiten.Image) {
	boardImage.Fill(frameColor)

	for t := range b.tiles {
		op := &ebiten.DrawImageOptions{}
		x := t.current.x*tileSize + (t.current.x+1)*tileMargin
		y := t.current.y*tileSize + (t.current.y+1)*tileMargin
		op.GeoM.Translate(float64(x), float64(y))
		if t.isActive {
			op.ColorScale.ScaleWithColor(tileBackgroundColor(128))
		} else {
			op.ColorScale.ScaleWithColor(tileBackgroundColor(0))
		}

		boardImage.DrawImage(t.image, op)
		t.Draw(boardImage)
	}
}
