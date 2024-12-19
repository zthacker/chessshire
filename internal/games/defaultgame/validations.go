package defaultgame

import (
	"ChessShire/internal/common"
	"fmt"
)

func isInBounds(row, col uint8) bool {
	//since we're checking a uint8 (which can't be negative) we don't add that condition here
	return col <= 7 && row <= 7
}

func positionOutOfBounds(pos uint8) error {
	if pos > MAX_BOARD_SIZE-1 || pos < 0 {
		return fmt.Errorf("out of bounds")
	}
	return nil
}

func playerTurnOwnsPiece(owner, color uint8) error {
	if owner != color {
		return fmt.Errorf("you do not own that piece -- owner: %d color: %d", owner, color)
	}
	return nil
}

func (dc *DefaultChessShireGame) getPieceFromPosition(pos uint8) (*GamePiece, error) {
	piece := dc.positionToEntities[pos]
	if piece == nil {
		return nil, fmt.Errorf("piece not found at position %d", pos)
	}
	return piece, nil
}

func (dc *DefaultChessShireGame) isSpaceOccupied(pos uint8) bool {
	piece := dc.positionToEntities[pos]
	if piece == nil {
		return false
	}
	return true
}

func validateMoveByMovementType(moveType uint8) error {
	switch moveType {
	case common.MOVEMENT_TYPE_PAWN:
		return nil
	case common.MOVEMENT_TYPE_KNIGHT:
		return nil
	case common.MOVEMENT_TYPE_ROOK:
		return nil
	case common.MOVEMENT_TYPE_BISHOP:
		return nil
	case common.MOVEMENT_TYPE_KING:
		return nil
	case common.MOVEMENT_TYPE_QUEEN:
		return nil
	default:
		return fmt.Errorf("unknown movement type")
	}
}
