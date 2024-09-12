package chainstask_service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"math"
	"net/http"
)

type getMultipleAccountsRequestBody struct {
	ID      int64      `json:"id"`
	JsonRpc string     `json:"jsonrpc"`
	Method  string     `json:"method"`
	Params  [][]string `json:"params"`
}

type getMultipleAccountsResponseBody struct {
	Result struct {
		Value []struct {
			Lamports int64 `json:"lamports"`
		} `json:"value"`
	} `json:"result"`
}

func (s *service) GetAccountsBalance(ctx context.Context, publicKeys []string) ([]SolAccount, error) {
	bodyBytes, err := json.Marshal(getMultipleAccountsRequestBody{
		ID:      1,
		JsonRpc: "2.0",
		Method:  "getMultipleAccounts",
		Params:  [][]string{publicKeys},
	})
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, s.chainstackURL, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	bodyBytes, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var decodedResponse getMultipleAccountsResponseBody
	err = json.Unmarshal(bodyBytes, &decodedResponse)
	if err != nil {
		return nil, err
	}

	result := make([]SolAccount, 0, len(publicKeys))
	if len(decodedResponse.Result.Value) != len(publicKeys) {
		return nil, errors.New("invalid response length")
	}

	for i, key := range publicKeys {
		result = append(result, SolAccount{
			PublicKey: key,
			Sol:       math.Round(float64(decodedResponse.Result.Value[i].Lamports)/1000000000*100) / 100,
		})
	}

	return result, nil
}
