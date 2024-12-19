package defaultgame

import (
	"ChessShire/internal/characters/gamecharacter"
	"ChessShire/internal/common"
	"ChessShire/internal/factory"
	"ChessShire/internal/player"
	"fmt"
	"log"
	"sync"
)

const (
	WHITE_PLAYER = 1
	BLACK_PLAYER = 2
)

const (
	ROW_MULTIPLICATIVE_CONSTANT = 8
	MAX_BOARD_SIZE              = 64
)

type GamePiece struct {
	*gamecharacter.GameCharacter
	owner uint8
}

type DefaultChessShireGame struct {
	//shared bit board
	chessShireBoard uint64
	//map of positions to entities
	positionToEntities [MAX_BOARD_SIZE]*GamePiece

	//initial position layout
	initialLayout [MAX_BOARD_SIZE]uint8

	//factory
	gameFactory factory.MainFactory

	//Players
	whitePlayer *player.Player
	blackPlayer *player.Player

	lock sync.Mutex

	//sudo-enum gameState with iota (check in models.go)
	gameState uint8
}

func NewGame(players [2]*player.Player) *DefaultChessShireGame {
	return &DefaultChessShireGame{
		initialLayout: createInitialBoardLayout(),
		gameFactory:   factory.MainFactory{},
		whitePlayer:   players[0],
		blackPlayer:   players[1],
		gameState:     GAME_STATE_SETUP,
	}
}

func (dc *DefaultChessShireGame) PrintPositions() {
	for row := 7; row >= 0; row-- {
		fmt.Printf("%d| ", row+1)
		for col := 0; col < 8; col++ {
			pos := bitBoardLocation(uint8(row), uint8(col))

			fmt.Printf(" %d ", pos)
		}
		fmt.Println()
	}
	fmt.Println("   a b c d e f g h")
	fmt.Println()
}

func (dc *DefaultChessShireGame) PrintOwners() {
	for row := 7; row >= 0; row-- {
		fmt.Printf("%d| ", row+1)
		for col := 0; col < 8; col++ {
			pos := bitBoardLocation(uint8(row), uint8(col))

			piece, err := dc.getPieceFromPosition(pos)
			if err != nil {
				continue
			}
			fmt.Printf("%d ", piece.owner)
		}
		fmt.Println()
	}
	fmt.Println("   a b c d e f g h")
	fmt.Println()
}

func (dc *DefaultChessShireGame) Play() uint8 {
	for dc.gameState != GAME_STATE_GAME_OVER {
		switch dc.gameState {
		case GAME_STATE_SETUP:
			dc.gameInit()
			dc.gameState = GAME_STATE_PLAYER_WHITE_TURN
			break
		case GAME_STATE_PLAYER_WHITE_TURN:
			dc.printBoard()
			fmt.Println("player white's turn...")
			//do player inputs
			fromPosition, toPosition, err := getMoveInput()
			if err != nil {
				fmt.Printf("%s try again...\n", err)
				break
			}

			//check boundaries
			if err = dc.boundaryChecks(fromPosition, toPosition); err != nil {
				fmt.Printf("%s try again...\n", err)
				break
			}

			//pieces
			ownerPiece, err := dc.getPieceFromPosition(fromPosition)
			if err != nil {
				fmt.Printf("%s try again...\n", err)
				break
			}

			//check if player owns the piece
			if err = playerTurnOwnsPiece(ownerPiece.owner, WHITE_PLAYER); err != nil {
				fmt.Printf("%s try again...\n", err)
				break
			}

			//validate the owner pieces move
			if err = validateMoveByMovementType(ownerPiece.MovementType); err != nil {
				fmt.Printf("%s try again...\n", err)
				break
			}

			//check if to position is occupied; if it's not, move there
			if !dc.isSpaceOccupied(toPosition) {
				//move
				dc.removePiece(fromPosition)
				dc.setPiece(toPosition, ownerPiece)
				dc.gameState = GAME_STATE_PLAYER_BLACK_TURN
				break
			}

			//if it is occupied, check if we can move there GamePiece is not owned by player turn
			//do battle
			//TODO this is kind of bad here because we can't really guarantee there's not going to be an error. We may want to grab owning piece earlier
			toPiece, err := dc.getPieceFromPosition(toPosition)
			if err != nil {
				fmt.Printf("%s try again...\n", err)
				break
			}

			//TODO this doesn't take into account Castling...
			if toPiece.owner == ownerPiece.owner {
				fmt.Printf("can't move onto a square that you own already try again...\n")
				break
			}

			//simulate dumb battle
			dc.removePiece(toPosition)
			dc.removePiece(fromPosition)
			dc.setPiece(toPosition, ownerPiece)
			dc.gameState = GAME_STATE_PLAYER_BLACK_TURN
			break

		case GAME_STATE_PLAYER_BLACK_TURN:
			dc.printBoard()
			fmt.Println("player black's turn...")
			//do player inputs
			fromPosition, toPosition, err := getMoveInput()
			if err != nil {
				fmt.Printf("%s try again...\n", err)
				break
			}

			//check boundaries
			if err = dc.boundaryChecks(fromPosition, toPosition); err != nil {
				fmt.Printf("%s try again...\n", err)
				break
			}

			//pieces
			ownerPiece, err := dc.getPieceFromPosition(fromPosition)
			if err != nil {
				fmt.Printf("%s try again...\n", err)
				break
			}

			//check if player owns the piece
			if err = playerTurnOwnsPiece(ownerPiece.owner, BLACK_PLAYER); err != nil {
				fmt.Printf("%s try again...\n", err)
				break
			}

			//validate the owner pieces move
			if err = validateMoveByMovementType(ownerPiece.MovementType); err != nil {
				fmt.Printf("%s try again...\n", err)
				break
			}

			//check if to position is occupied; if it's not, move there
			if !dc.isSpaceOccupied(toPosition) {
				//move
				dc.removePiece(fromPosition)
				dc.setPiece(toPosition, ownerPiece)
				dc.gameState = GAME_STATE_PLAYER_WHITE_TURN
				break
			}

			//if it is occupied, check if we can move there GamePiece is not owned by player turn
			//do battle
			//TODO this is kind of bad here because we can't really guarantee there's not going to be an error. We may want to grab owning piece earlier
			toPiece, err := dc.getPieceFromPosition(toPosition)
			if err != nil {
				fmt.Printf("%s try again...\n", err)
				break
			}

			//TODO this doesn't take into account Castling...
			if toPiece.owner == ownerPiece.owner {
				fmt.Printf("can't move onto a square that you own already try again...\n")
				break
			}

			//simulate dumb battle
			dc.removePiece(toPosition)
			dc.removePiece(fromPosition)
			dc.setPiece(toPosition, ownerPiece)
			dc.gameState = GAME_STATE_PLAYER_BLACK_TURN
			break
		case GAME_STATE_GAME_OVER:
			//disconnect players
		default:
			return GAME_STATE_UNKNOWN
		}
	}

	return GAME_STATE_GAME_OVER
}

func (dc *DefaultChessShireGame) performTurn(fromPos, toPos uint8, turnColor uint8) error {

	return nil
}

func (dc *DefaultChessShireGame) printBoard() {
	for row := 7; row >= 0; row-- {
		fmt.Printf("%d| ", row+1)
		for col := 0; col < 8; col++ {
			pos := bitBoardLocation(uint8(row), uint8(col))

			piece := dc.positionToEntities[pos]
			if piece != nil {
				fmt.Printf("%d ", piece.PieceType)
				continue
			}

			fmt.Printf(". ")
		}
		fmt.Println()
	}
	fmt.Println("   a b c d e f g h")
	fmt.Println()
}

func (dc *DefaultChessShireGame) gameInit() {
	for pos := uint8(0); pos <= uint8(MAX_BOARD_SIZE-1); pos++ {
		piece := dc.initialLayout[pos]
		if piece != 0 {
			if pos <= 15 { //white player
				playerPieceType, found := dc.whitePlayer.Pieces[piece]
				if !found {
					//TODO move this to a validation pre-game
					log.Fatalf("white player does not have piece type %d", piece)
				}
				dc.positionToEntities[pos] = &GamePiece{
					GameCharacter: dc.gameFactory.CreatePieceFromType(playerPieceType),
					owner:         WHITE_PLAYER,
				}
			}
			if pos >= 48 { //black player
				playerPieceType, found := dc.blackPlayer.Pieces[piece]
				if !found {
					//TODO move this to a validation pre-game
					log.Fatalf("black player does not have piece type %d", piece)
				}
				dc.positionToEntities[pos] = &GamePiece{
					GameCharacter: dc.gameFactory.CreatePieceFromType(playerPieceType),
					owner:         BLACK_PLAYER,
				}
			}
		}
	}
}

func bitBoardLocation(row, col uint8) uint8 {
	return row*ROW_MULTIPLICATIVE_CONSTANT + col
}

func createInitialBoardLayout() [MAX_BOARD_SIZE]uint8 {
	return [MAX_BOARD_SIZE]uint8{
		0:  common.PIECE_TYPE_ROOK,
		1:  common.PIECE_TYPE_KNIGHT,
		2:  common.PIECE_TYPE_BISHOP,
		3:  common.PIECE_TYPE_QUEEN,
		4:  common.PIECE_TYPE_KING,
		5:  common.PIECE_TYPE_BISHOP,
		6:  common.PIECE_TYPE_KNIGHT,
		7:  common.PIECE_TYPE_ROOK,
		8:  common.PIECE_TYPE_PAWN,
		9:  common.PIECE_TYPE_PAWN,
		10: common.PIECE_TYPE_PAWN,
		11: common.PIECE_TYPE_PAWN,
		12: common.PIECE_TYPE_PAWN,
		13: common.PIECE_TYPE_PAWN,
		14: common.PIECE_TYPE_PAWN,
		15: common.PIECE_TYPE_PAWN,
		48: common.PIECE_TYPE_PAWN,
		49: common.PIECE_TYPE_PAWN,
		50: common.PIECE_TYPE_PAWN,
		51: common.PIECE_TYPE_PAWN,
		52: common.PIECE_TYPE_PAWN,
		53: common.PIECE_TYPE_PAWN,
		54: common.PIECE_TYPE_PAWN,
		55: common.PIECE_TYPE_PAWN,
		56: common.PIECE_TYPE_ROOK,
		57: common.PIECE_TYPE_KNIGHT,
		58: common.PIECE_TYPE_BISHOP,
		59: common.PIECE_TYPE_QUEEN,
		60: common.PIECE_TYPE_KING,
		61: common.PIECE_TYPE_BISHOP,
		62: common.PIECE_TYPE_KNIGHT,
		63: common.PIECE_TYPE_ROOK,
	}
}
