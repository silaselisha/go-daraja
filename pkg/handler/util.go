package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func handlerHelper[T B2BReqParams | B2CReqParams | C2BReqParams | ExpressReqParams| QRReqParams | BExpressCheckoutParams](payload T, url, method, authToken string) ([]byte, error) {
	fmt.Println(payload)
	buff, err := json.Marshal(&payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(buff))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+authToken)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func(){
		if err := res.Body.Close(); err != nil {
			log.Fatalf("failed to close response body %v\n", err)
		}
	}()
	return io.ReadAll(res.Body)
}
