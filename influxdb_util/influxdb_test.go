package influxdb_util

import (
	"testing"
	"time"
)

func TestClient_CreateDB(t *testing.T) {
	var client Client
	config := Config{
		Host:     "192.168.21.129",
		Port:     8086,
		Username: "",
		Password: "",
	}

	err := client.SetConfig(config)
	if err != nil {
		t.Error(err)
		return
	}

	err = client.CreateDB("test3")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestClient_Insert(t *testing.T) {
	var client Client
	config := Config{
		Host:     "192.168.21.129",
		Port:     8086,
		Username: "",
		Password: "",
	}

	err := client.SetConfig(config)
	if err != nil {
		t.Error(err)
		return
	}

	tags := map[string]string{
		"name": "ai",
		"city": "G",
	}
	fields := map[string]interface{}{
		"age": 30,
	}

	err = client.Insert("test3", "nginx_log3", tags, fields, time.Now())
	if err != nil {
		t.Error(err)
	}

}
