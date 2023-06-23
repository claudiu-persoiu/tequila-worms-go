package game

type repo struct {
	worms map[string]*worm
}

func NewRepo() *repo {
	return &repo{worms: map[string]*worm{}}
}

func (r *repo) AddWorm(uuid string, w *worm) {
	r.worms[uuid] = w
}

func (r *repo) GetWorm(uuid string) *worm {
	return r.worms[uuid]
}

func (r *repo) RemoveWorm(uuid string) *worm {
	if r.worms[uuid] != nil {
		w := r.worms[uuid]
		delete(r.worms, uuid)
		return w
	}
	return nil
}

type wormNames struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

func (r *repo) GetWormsList() []wormNames {
	wn := []wormNames{}
	for _, w := range r.worms {
		wn = append(wn, wormNames{
			Name:  w.name,
			Color: w.color,
		})
	}
	return wn
}

type wormData struct {
	Name   string  `json:"name"`
	Color  string  `json:"color"`
	Pieces []piece `json:"pieces"`
}

func (r *repo) GetWormsDataList() []wormData {
	wn := []wormData{}
	for _, w := range r.worms {
		wn = append(wn, wormData{
			Name:   w.name,
			Color:  w.color,
			Pieces: w.pieces,
		})
	}
	return wn
}

func (r *repo) GetSize() int {
	return len(r.worms)
}

func (r *repo) GetWorms() map[string]*worm {
	return r.worms
}
