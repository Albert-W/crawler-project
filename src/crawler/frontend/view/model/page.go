package model

type SearchResult struct {
	Hits int64
	Start int
	Items []interface{}
	//Items []engine.Item
	Query string
	PrevFrom int
	NextFrom int
}
