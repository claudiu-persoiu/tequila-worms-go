package game

type repo struct {
	worms map[string]*worm
}

func (r *repo) AddWorm(uuid string, w *worm) {
	r.worms[uuid] = w
}

func (r *repo) GetWorm(uuid string) *worm {
	return r.worms[uuid]
}

func (r *repo) RemoveWorm(uuid string) {
	if r.worms[uuid] != nil {
		delete(r.worms, uuid)
	}
}
