package defaultgame

import "fmt"

func (dc *DefaultChessShireGame) boundaryChecks(fromPos, toPos uint8) error {
	//position checks
	if err := positionOutOfBounds(fromPos); err != nil {
		return err
	}
	if err := positionOutOfBounds(toPos); err != nil {
		return err
	}

	return nil
}

func (dc *DefaultChessShireGame) move(fromPos, toPos uint8, piece *GamePiece) {
	dc.positionToEntities[fromPos] = nil
	dc.positionToEntities[toPos] = piece
}

func getMoveInput() (uint8, uint8, error) {
	var from, to string

	//source input
	fmt.Print("Enter source position (e.g., e2): ")
	if _, err := fmt.Scan(&from); err != nil {
		return 0, 0, err
	}

	fromRow, fromCol, err := parseMoveInput(from)
	if err != nil {
		return 0, 0, err
	}

	fromBitPosition := bitBoardLocation(fromRow, fromCol)

	//destination input
	fmt.Print("Enter destination position (e.g., f2): ")
	if _, err = fmt.Scan(&to); err != nil {
		return 0, 0, err
	}

	toRow, toCol, err := parseMoveInput(to)
	if err != nil {
		return 0, 0, err
	}

	toBitPosition := bitBoardLocation(toRow, toCol)

	return fromBitPosition, toBitPosition, nil
}

func parseMoveInput(input string) (uint8, uint8, error) {
	if len(input) != 2 {
		return 0, 0, fmt.Errorf("invalid input: %s must be colum letter and numerical row. Example: e2", input)
	}

	col := input[0] - 'a' //ascii -- so a is 97. Used to map Chess board columns conversions
	row := input[1] - '1' // 0 index based board, but displayed at 1 index based. Subtract the 1 to make it 0 based

	if !isInBounds(row, col) {
		return 0, 0, fmt.Errorf("row: %d col %d is out of bounds, try again", row, col)
	}

	return row, col, nil
}
