package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type CatBreed struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func ValidateBreed(breedName string) error {
	url := "https://api.thecatapi.com/v1/breeds"
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to fetch breeds: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var breeds []CatBreed
	err = json.Unmarshal(body, &breeds)
	if err != nil {
		return err
	}

	for _, breed := range breeds {
		if breed.Name == breedName {
			return nil
		}
	}

	return errors.New("invalid breed name")
}
