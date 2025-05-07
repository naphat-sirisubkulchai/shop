package omise
import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func CreateCharge(amount int, currency, token string) error {
	url := "https://api.omise.co/charges"

	payload := map[string]interface{}{
		"amount":   amount,
		"currency": currency,
		"card":     token,
	}

	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.SetBasicAuth(os.Getenv("OMISE_SECRET_KEY"), "")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println("Status:", resp.Status)
	return nil
}