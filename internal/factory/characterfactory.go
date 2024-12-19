package factory

import (
	"ChessShire/internal/characters/gamecharacter"
	"ChessShire/internal/common"
)

type MainFactory struct{}

func (mf MainFactory) CreatePieceFromType(pieceType uint64) *gamecharacter.GameCharacter {
	switch pieceType {
	case common.GC_DEFAULT_PAWN:
		return NewDefaultPawn()
	case common.GC_DEFAULT_ROOK:
		return NewDefaultRook()
	case common.GC_DEFAULT_KNIGHT:
		return NewDefaultKnight()
	case common.GC_DEFAULT_BISHOP:
		return NewDefaultBishop()
	case common.GC_DEFAULT_QUEEN:
		return NewDefaultQueen()
	case common.GC_DEFAULT_KING:
		return NewDefaultKing()
	default:
		return nil
	}
}
