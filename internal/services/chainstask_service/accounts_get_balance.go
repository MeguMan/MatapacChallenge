package chainstask_service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"time"
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
	start := time.Now()

	bodyBytes, err := json.Marshal(getMultipleAccountsRequestBody{
		ID:      1,
		JsonRpc: "2.0",
		Method:  "getMultipleAccounts",
		Params:  [][]string{publicKeys},
	})
	if err != nil {
		return nil, err
	}
	elapsed := time.Since(start)
	log.Printf("bodyReady took %s", elapsed)

	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, s.chainstackURL, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	elapsed = time.Since(start)
	log.Printf("req took %s", elapsed)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	elapsed = time.Since(start)
	log.Printf("resp arrived %s", elapsed)

	bodyBytes, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var decodedResponse getMultipleAccountsResponseBody
	err = json.Unmarshal(bodyBytes, &decodedResponse)
	if err != nil {
		return nil, err
	}

	elapsed = time.Since(start)
	log.Printf("resp decoded %s", elapsed)

	result := make([]SolAccount, 0, len(publicKeys))
	if len(decodedResponse.Result.Value) != len(publicKeys) {
		return nil, errors.New("invalid response length")
	}

	for i, key := range publicKeys {
		result = append(result, SolAccount{
			PublicKey: key,
			Sol:       float64(decodedResponse.Result.Value[i].Lamports) / 1000000000,
		})
	}

	elapsed = time.Since(start)
	log.Printf("done %s", elapsed)

	return result, nil
}
