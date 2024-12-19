package main

import (
	"ChessShire/internal/factory"
	"ChessShire/internal/games/defaultgame"
	"ChessShire/internal/player"
)

func main() {

	//<bitboard> |= 1 << pos modifies the <bitboard> to mark the presence of a piece at the square corresponding to pos
	//00000000 |= 1 << 0 = 00000001  flip position 0 to "on" or 1
	//00000000 |= 1 << 2 = 00000100  flip position 2 to "on" or 1

	//var a uint64
	//a |= 1 << 9
	//010000010
	//pos := 0x0000000000000042
	//fmt.Printf("%064b\n", pos)
	//
	//if (pos>>4)&1 == 1 {
	//	fmt.Println("yay")
	//}

	//pawns := 0x000000000000FF00

	//fmt.Printf("%64b\n", pawns)
	//if (pawns>>8)&1 == 1 {
	//	fmt.Println("that's a pawn")
	//} else {
	//	fmt.Println("not a pawn")
	//}
	//fmt.Printf("%0x\n", 0x000000000000FF00)
	//fmt.Printf("%d\n", 0x000000000000FF00)

	//blackPawns := 0x00FF000000000000
	//fmt.Printf("%064b", blackPawns)
	//var move string
	//fmt.Scan(&move)
	//fmt.Println(move[0] - 'a')
	//fmt.Println(move[1] - '0')

	//whitePlayer := map[uint8]*gamecharacter.GameCharacter{
	//	common.PIECE_TYPE_PAWN: factory.NewDefaultPawn(),
	//}

	playerOne := factory.NewDefaultPlayer()
	playerTwo := factory.NewDefaultPlayer()

	players := [2]*player.Player{}
	players[0] = playerOne
	players[1] = playerTwo

	//game := board.NewChessShireGame(players)
	//game.Play()

	defaultGame := defaultgame.NewGame(players)
	defaultGame.Play()

	//game.PrintBoard()

	//game.PrintBitBoards()
	//game.PrintBitPositions()

	//testSimpleFight()

}

func getPosition(row int, col int) uint64 {
	return uint64(1) << (row*8 + col)
}

//func testSimpleFight() {
//	createdFighter := factory.NewDefaultFighter("Epic", 0)
//
//	createdGoblin := &gamecharacter.GameCharacter{
//		BaseCharacter: &basecharacter.BaseCharacter{
//			Name:     "EpicGoblin",
//			Position: 0,
//		},
//		Components: make(map[string]components.Component),
//	}
//
//	createdGoblin.Components[common.ATTRIBUTES] = attributes.NewAttributeComponent(map[string]int{
//		common.HEALTH:   20,
//		common.MANA:     100,
//		common.STRENGTH: 20,
//		common.RES_FIRE: 5,
//	})
//
//	goblinAttributes := createdGoblin.Components[common.ATTRIBUTES].(classbehaviors.Attributes)
//
//	for !goblinAttributes.IsZero(common.HEALTH) {
//		fmt.Println(fmt.Sprintf("%s is attacking %s", createdFighter.Name, createdGoblin.Name))
//		fighterAbilities := createdFighter.GetAbilityComponent()
//		fighterRes := fighterAbilities.UseAbility(createdFighter.GetAttributeComponent(), "fireball", []classbehaviors.Attributes{goblinAttributes})
//		fmt.Println(fighterRes)
//
//		time.Sleep(1 * time.Second)
//
//	}
//}
