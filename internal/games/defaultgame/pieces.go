package defaultgame

func (dc *DefaultChessShireGame) pieceChecks() error {

	return nil
}

func (dc *DefaultChessShireGame) removePiece(pos uint8) {
	dc.lock.Lock()
	defer dc.lock.Unlock()
	dc.positionToEntities[pos] = nil
}

func (dc *DefaultChessShireGame) setPiece(pos uint8, piece *GamePiece) {
	dc.lock.Lock()
	defer dc.lock.Unlock()
	dc.positionToEntities[pos] = piece
}
