package json

import (
	"encoding/json"
	"main/structs"
	"os"
)

type Output struct {
	Node_investments []float32
}

func ConvertJson(out Output) error {
	json, err := json.Marshal(out)
	if err != nil {
		return err
	} else {
		os.WriteFile("model.json", json, os.ModePerm)
	}
	return err
}

func (out *Output) AppendAllNodeInvestments(ncl structs.NodeCollection) {
	for i := range ncl.Nodes {
		out.Node_investments = append(out.Node_investments, ncl.Nodes[i].Investment)
	}
}
