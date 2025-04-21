package json

import (
	"encoding/json"
	"main/structs"
	"os"
)

type Output struct {
	Node_investments     []structs.Price
	MarketPrices         []structs.Price
	Node_Shares          []int
	Agg_Node_Shares      map[int][]int
	Agg_Node_Investments map[int][]structs.Price
}

func ConvertJson(out Output) error {
	json, err := json.Marshal(out)
	if err != nil {
		return err
	} else {
		os.WriteFile("model_1.json", json, os.ModePerm)
	}
	return err
}

func (out *Output) Init() {
	out.Agg_Node_Investments = map[int][]structs.Price{}
	out.Agg_Node_Shares = map[int][]int{}
}

func (out *Output) AppendAllNodeInvestments(ncl structs.NodeCollection) {
	for i := range ncl.Nodes {
		out.Node_investments = append(out.Node_investments, ncl.Nodes[i].Investment)
		out.Node_Shares = append(out.Node_Shares, int(ncl.Nodes[i].Shares))
	}
}

func (out *Output) AppendMarketPrices(mkt structs.Stock_Market) {
	for i := range mkt.PrevPrices {
		out.MarketPrices = append(out.MarketPrices, *mkt.PrevPrices[i])
	}
}

func (out *Output) AppendNodeStats(j int64, node *structs.Node) {
	i := int(j)
	out.Agg_Node_Investments[i] = append(out.Agg_Node_Investments[i], node.Investment)
	out.Agg_Node_Shares[i] = append(out.Agg_Node_Shares[i], node.Shares)
}
