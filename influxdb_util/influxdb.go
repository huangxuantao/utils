package influxdb_util

import (
	"fmt"
	"github.com/influxdata/influxdb/client/v2"
	"time"
)

type Client struct {
	client client.Client
}

type Config struct {
	Host     string
	Port     int
	Username string
	Password string
}

func (c *Client) SetConfig(config Config) (err error) {
	c.client, err = client.NewHTTPClient(client.HTTPConfig{
		Addr:     fmt.Sprintf("http://%s:%d", config.Host, config.Port),
		Username: config.Username,
		Password: config.Password,
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) CreateDB(db string) (err error) {
	sql := client.Query{
		Command:  fmt.Sprintf("CREATE DATABASE %s", db),
		Database: db,
	}
	if response, err := c.client.Query(sql); err == nil {
		if response.Error() != nil {
			return response.Error()
		}
	} else {
		return err
	}
	return nil
}

func (c *Client) Query(db string, cmd string) (rst []client.Result, err error) {
	q := client.Query{
		Command:  cmd,
		Database: db,
	}
	if response, err := c.client.Query(q); err == nil {
		if response.Error() != nil {
			return rst, response.Error()
		}
		rst = response.Results
	} else {
		return rst, err
	}
	return rst, nil
}

func (c *Client) Insert(db string, measurement string, tags map[string]string, fields map[string]interface{}, timeSet time.Time) (err error) {
	points, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database: db,
	})
	if err != nil {
		return err
	}
	point, err := client.NewPoint(measurement, tags, fields, timeSet)
	if err != nil {
		return err
	}

	points.AddPoint(point)
	err = c.client.Write(points)
	if err != nil {
		return err
	}
	return nil
}
