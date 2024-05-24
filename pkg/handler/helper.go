package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/rs/zerolog"
)

func handlerHelper[T B2BReqParams | B2CReqParams | C2BReqParams | ExpressReqParams | BExpressCheckoutParams](logger zerolog.Logger, payload T, url, method, authToken string) (*DarajaResParams, error) {
	buff, err := json.Marshal(&payload)
	if err != nil {
		logger.Error().Msg(err.Error())
		return nil, err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(buff))
	if err != nil {
		logger.Error().Msg(err.Error())
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+authToken)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		logger.Error().Msg(err.Error())
		return nil, err
	}

	defer func() {
		if err := res.Body.Close(); err != nil {
			logger.Error().Msg(err.Error())
			log.Fatalf("failed to close response body %v\n", err)
		}
	}()

	buff, err = io.ReadAll(res.Body)
	if err != nil {
		logger.Error().Msg(err.Error())
		return nil, err
	}

	logger.Info().Msg("Bytes From Helper: " + string(buff))
	var result DarajaResParams
	if err := json.Unmarshal(buff, &result); err != nil {
		logger.Error().Msg(err.Error())
		return nil, err
	}
	return &result, nil
}
