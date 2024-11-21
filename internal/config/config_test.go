package config

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestLoadOrDefaultRun(t *testing.T) {

	// write a fake env file
	cfg := &Config{IP: "127.1.2.3", Port: "9123"}
	dotenvFile := "testdata/.env"

	_, err := cfg.toDotenv(dotenvFile)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	defer os.Remove(dotenvFile)

	testCases := []struct {
		filename string
		want     *Config
	}{
		{"", &defaultConfig}, // default configuration
		{dotenvFile, cfg},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("[%s]", tc.filename), func(t *testing.T) {
			got, err := loadOrDefault(tc.filename)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Expected %+v, but got %+v", tc.want, got)
			}
		})
	}

}
