package utils

// Direction enum
type Direction int

const (
	NotSet Direction = iota
	Up
	UpRight
	Right
	DownRight
	Down
	DownLeft
	Left
	UpLeft
)

func (d Direction) String() string {
	return [...]string{"NotSet", "Up", "UpRight", "Right", "DownRight", "Down", "DownLeft", "Left", "UpLeft"}[d]
}

func GetAllDirections() []Direction {
	return []Direction{Up, UpRight, Right, DownRight, Down, DownLeft, Left, UpLeft}
}

func (d Direction) Turn(turn Direction) Direction {
	if turn != Left && turn != Right {
		panic("should be left or right")
	}

	switch d {
	case Up:
		return turn
	case Down:
		switch turn {
		case Left:
			return Right
		case Right:
			return Left
		}
	case Left:
		switch turn {
		case Left:
			return Down
		case Right:
			return Up
		}
	case Right:
		switch turn {
		case Left:
			return Up
		case Right:
			return Down
		}
	}

	panic("not handled")
}
