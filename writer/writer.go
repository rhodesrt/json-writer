package writer

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/google/uuid"
)

type JsonWriter struct {
	Loop int
}

type UuidJson struct {
	Uuid []string `json:"uuid"`
}

func (j JsonWriter) Write() error {
	file, err := os.Create("output.json")
	if err != nil {
		return fmt.Errorf("Error creating file")
	}
	defer file.Close()

	u := UuidJson{
		Uuid: []string{},
	}
	for i := 0; i < j.Loop; i++ {
		u.Uuid = append(u.Uuid, uuid.NewString())
	}

	data, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return fmt.Errorf("Error marshaling JSON")
	}

	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("Error writing to file: %s", file.Name())
	}

	return nil
}
