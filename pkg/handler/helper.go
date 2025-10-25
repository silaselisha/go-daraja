package handler

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
)

func handlerHelper[T B2BReqParams | B2CReqParams | C2BReqParams | ExpressReqParams | BExpressCheckoutParams](payload T, url, method, authToken string) (*DarajaResParams, error) {
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

	defer func() {
		if err := res.Body.Close(); err != nil {
			log.Fatalf("failed to close response body %v\n", err)
		}
	}()

    buff, err = io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

    var result DarajaResParams
    if err := json.Unmarshal(buff, &result); err == nil {
        return &result, nil
    }

    // If the standard response shape failed to parse, try parsing known error fields
    var errPayload DarajaErrorParams
    if err := json.Unmarshal(buff, &errPayload); err == nil {
        result.DarajaErrorParams = errPayload
        return &result, nil
    }

    // Fallback: non-JSON (e.g., HTML error page). Map to a generic unreachable error
    result.ErrorCode = "500.003.1001"
    result.ErrorMessage = "Service is currently unreachable. Please try again later."
    // include a hint with truncated payload for debugging
    if len(buff) > 0 {
        preview := buff
        if len(preview) > 128 {
            preview = preview[:128]
        }
        _ = fmt.Sprintf("non-JSON response preview: %s", string(preview))
    }
    return &result, nil
}
