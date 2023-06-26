package game

func piecesCollision(piece1, piece2 piece) bool {
	return piece1.X == piece2.X && piece1.Y == piece2.Y
}

func checkHitTheWall(head piece, table size) bool {
	return head.X <= -1 || head.Y <= -1 || head.X >= table.X || head.Y >= table.Y
}

func intersectBetweenWorms(w *worm, worms map[string]*worm) {
	for i, p := range w.pieces {
		for _, otherWorm := range worms {
			if w.uuid == otherWorm.uuid || otherWorm.IsDead() {
				continue
			}

			if wormIntersectHeadWithPiece(otherWorm, p) {
				if i == 0 {
					headCrush(w, otherWorm)
				} else {
					otherWorm.AddPieces(len(w.pieces) - 1)
					w.RemovePieces(i)
				}
			}
		}
	}
}

func wormIntersectHeadWithPiece(w *worm, p piece) bool {
	head := w.pieces[0]
	return piecesCollision(head, p)
}

func headCrush(w, otherWorm *worm) {
	w1l := len(w.pieces)
	w2l := len(otherWorm.pieces)

	if w1l == w2l {
		w.Kill()
		otherWorm.Kill()
	} else if w1l < w2l {
		w.Kill()
		otherWorm.AddPieces(w1l)
	} else if w2l < w1l {
		otherWorm.Kill()
		w.AddPieces(w2l)
	}
}
