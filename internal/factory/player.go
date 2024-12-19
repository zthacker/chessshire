package factory

import (
	"ChessShire/internal/common"
	"ChessShire/internal/player"
	"fmt"
	"math/rand"
)

func NewDefaultPlayer() *player.Player {
	name := fmt.Sprintf("name%d", rand.Intn(1000))

	pieceMap := map[uint8]uint64{
		common.PIECE_TYPE_PAWN:   common.GC_DEFAULT_PAWN,
		common.PIECE_TYPE_BISHOP: common.GC_DEFAULT_BISHOP,
		common.PIECE_TYPE_KNIGHT: common.GC_DEFAULT_KNIGHT,
		common.PIECE_TYPE_ROOK:   common.GC_DEFAULT_ROOK,
		common.PIECE_TYPE_QUEEN:  common.GC_DEFAULT_QUEEN,
		common.PIECE_TYPE_KING:   common.GC_DEFAULT_KING,
	}

	return &player.Player{
		Name:   name,
		Pieces: pieceMap,
	}
}
