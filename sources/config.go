package sources

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"time"
)

type BoundaryConfig struct {
	Lat1Key string
	Lat2Key string
	Lon1Key string
	Lon2Key string
}

func (c *BoundaryConfig) GetBoundary() ([][][]float64, error) {

	lat1, err := getFloat64(c.Lat1Key)
	if err != nil {
		return nil, err
	}
	lon1, err := getFloat64(c.Lon1Key)
	if err != nil {
		return nil, err
	}
	lat2, err := getFloat64(c.Lat2Key)
	if err != nil {
		return nil, err
	}
	lon2, err := getFloat64(c.Lon2Key)
	if err != nil {
		return nil, err
	}
	bounds := [][][]float64{{{lat1, lon1}, {lat2, lon2}}}
	return bounds, nil
}

type Config struct {
	Name           string
	Protocol       string
	AddressKey     string
	ApiKey         string
	TimeoutSecsKey string
	RetrySecsKey   string
	Boundary       *BoundaryConfig
}

func (c *Config) GetAddress() (string, error) {
	return getString(c.AddressKey)
}

func (c *Config) GetApiKey() (string, error) {
	return getString(c.ApiKey)
}

func (c *Config) GetTimeout() (time.Duration, error) {
	secs, err := getInt(c.TimeoutSecsKey)
	if err != nil {
		return 0, err
	}
	return time.Duration(secs), nil
}

func (c *Config) GetRetry() (time.Duration, error) {
	secs, err := getInt(c.RetrySecsKey)
	if err != nil {
		return 0, err
	}
	return time.Duration(secs), nil
}

func (c *Config) GetBoundary() ([][][]float64, error) {
	if c.Boundary == nil {
		return nil, fmt.Errorf("booundary is not set")
	}
	return c.Boundary.GetBoundary()
}

func logError(key string, err error) error {
	if err == nil {
		err = fmt.Errorf("key '%s' not found", key)
	}
	slog.Error("error retrieving environment var", "key", key, "error", err)
	return err
}

func getString(key string) (string, error) {
	value := os.Getenv(key)
	if len(value) == 0 {
		return value, logError(key, nil)
	}
	return value, nil
}

func getFloat64(key string) (float64, error) {
	value, err := strconv.ParseFloat(os.Getenv(key), 64)
	if err != nil {
		return value, logError(key, err)
	}
	return value, nil
}

func getInt(key string) (int64, error) {
	value, err := strconv.ParseInt(os.Getenv(key), 10, 64)
	if err != nil {
		return value, logError(key, err)
	}
	return value, nil
}
