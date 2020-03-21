package amqp_util

import (
	"fmt"
	"github.com/streadway/amqp"
)

type Config struct {
	Account      string
	Password     string
	Host         string
	Port         string
	Exchange     string
	ExchangeKind string
	RoutingKey   string
	Queue        string
}

func (c *Config) ConnStr() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%s/", c.Account, c.Password, c.Host, c.Port)
}

type Client struct {
	Channel    *amqp.Channel
	Connection *amqp.Connection
	Queue      *amqp.Queue
	config     *Config
}

func GetClient(config *Config) (*Client, error) {
	conn, err := amqp.Dial(config.ConnStr())
	if err != nil {
		return nil, err
	}
	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	if err = channel.ExchangeDeclare(
		config.Exchange,
		config.ExchangeKind,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return nil, err
	}
	if err = channel.Qos(
		1,
		0,
		false,
	); err != nil {
		return nil, err
	}

	var queue amqp.Queue
	if config.Queue != "" {
		if queue, err = channel.QueueDeclare(
			config.Queue,
			true,
			false,
			false,
			false,
			nil,
		); err != nil {
			return nil, err
		}
	}

	return &Client{
		Channel:    channel,
		Connection: conn,
		Queue:      &queue,
		config:     config,
	}, nil
}

func (c *Client) Publish(body []byte) error {
	return c.Channel.Publish(
		c.config.Exchange,
		c.config.RoutingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
}

/*
这个写法是在使用的使用可以直接传入一个匿名函数
例如:
c.Consume(func(body []byte) error {
		log.Logger.Info(string(body))
		return nil
	}
)
这样起到一个Log记录的作用，中间件的作用
*/

type ProcessMQMessage func(body []byte) error

// Consume 消耗
func (c *Client) Consume(process ProcessMQMessage) error {
	if err := c.Channel.QueueBind(
		c.Queue.Name,
		c.config.RoutingKey,
		c.config.Exchange,
		false,
		nil,
	); err != nil {
		return err
	}

	messages, err := c.Channel.Consume(
		c.Queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil
	}

	for message := range messages {
		process(message.Body)
	}
	return nil
}

func (c *Client) Close() error {
	if c.Channel != nil {
		if err := c.Channel.Close(); err != nil {
			return err
		}
	}

	if c.Connection != nil {
		if err := c.Connection.Close(); err != nil {
			return err
		}
	}
	return nil
}
