package responses

import "encoding/json"

type Int64ListResponse struct {
	Data   []uint64
	Status int
	hideMe string
}

func (r Int64ListResponse) Json() []byte {
	rj, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return rj
}
