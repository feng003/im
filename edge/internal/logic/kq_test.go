package logic

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"testing"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/stretchr/testify/assert"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/cmdline"
	"github.com/zeromicro/go-zero/core/logx"

	"math/rand"

	"encoding/json"
)

func TestNewPusher(t *testing.T) {
	addrs := []string{"127.0.0.1:9092"}
	topic := "test-topic"

	t.Run("DefaultOptions", func(t *testing.T) {
		pusher := kq.NewPusher(addrs, topic)
		assert.NotNil(t, pusher)
		err := pusher.Push(context.Background(),"test")
		assert.Nil(t, err)
		fmt.Printf("pusher: %+v\n", pusher.Name())
	})

	t.Run("WithChunkSize", func(t *testing.T) {
		pusher := kq.NewPusher(addrs, topic, kq.WithChunkSize(100))
		assert.NotNil(t, pusher)
		for i := 0; i < 10; i++ {
			err := pusher.Push(context.Background(),"test" + strconv.Itoa(i+10))
			assert.Nil(t, err)
		}
		fmt.Printf("pusher: %+v\n", pusher.Name())
	})

	t.Run("WithFlushInterval", func(t *testing.T) {
		pusher := kq.NewPusher(addrs, topic, kq.WithFlushInterval(time.Second))
		assert.NotNil(t, pusher)
		for i := 0; i < 10; i++ {
			err := pusher.Push(context.Background(),"test" + strconv.Itoa(i))
			assert.Nil(t, err)
		}
		fmt.Printf("pusher: %+v\n", pusher.Name())
	})

	t.Run("WithAllowAutoTopicCreation", func(t *testing.T) {
		pusher := kq.NewPusher(addrs, topic, kq.WithAllowAutoTopicCreation())
		assert.NotNil(t, pusher)
	})
}

func TestKafka(t *testing.T) {
	// Kafka配置
	broker := "localhost:9092"
	topic := "test-topic"

	// 创建生产者
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": broker})
	if err != nil {
		logx.Errorf("创建生产者失败: %s\n", err)
	}
	defer producer.Close()

	// 生产消息
	for _, word := range []string{"Hello", "Kafka", "from", "Golang"} {
		err = producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(word),
		}, nil)
		if err != nil {
			logx.Errorf("生产消息失败: %s\n", err)
		}
		logx.Infof("生产消息: %s\n", word)
	}

	// 等待所有消息都被交付
	producer.Flush(5 * 1000)

	// 创建消费者
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
		"group.id":          "test-group",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		logx.Errorf("创建消费者失败: %s\n", err)
	}
	defer consumer.Close()

	_ = consumer.SubscribeTopics([]string{topic}, nil)

	//消费消息
	run := true
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for run {
			msg, err := consumer.ReadMessage(-1)
			if err == nil {
				fmt.Printf("收到消息: %s\n", string(msg.Value))
			} else {
				fmt.Printf("消费消息失败: %v (%v)\n", err, msg)
			}
		}
	}()

	<-sigchan
	run = false
	//
	// 删除主题
	//adminClient, err := kafka.NewAdminClientFromProducer(producer)
	//if err != nil {
	//	log.Fatalf("创建AdminClient失败: %s\n", err)
	//}
	//defer adminClient.Close()
	//
	//timeout := 10 * time.Second
	//_, err = adminClient.DeleteTopics(context.Background(), []string{topic}, kafka.SetAdminOperationTimeout(timeout))
	//if err != nil {
	//	log.Fatalf("删除主题失败: %s\n", err)
	//}
	//fmt.Printf("主题 %s 删除成功\n", topic)
}

type message struct {
	Key     string `json:"key"`
	Value   string `json:"value"`
	Payload string `json:"message"`
}

func TestKq(t *testing.T) {

	kafkaBrokers := []string{"localhost:9092"}
	//topic := "test-topic"
	topic := "topic-test-01"

	pusher := kq.NewPusher(kafkaBrokers, topic)
	ticker := time.NewTicker(time.Millisecond)
	for round := 0; round < 3; round++ {
		select {
		case <-ticker.C:
			count := rand.Intn(100)
			m := message{
				Key:     strconv.FormatInt(time.Now().UnixNano(), 10),
				Value:   fmt.Sprintf("%d,%d", round, count),
				Payload: fmt.Sprintf("%d,%d", round, count),
			}
			body, err := json.Marshal(m)
			if err != nil {
				logx.Errorf("marshal error: %v", err)
			}

			fmt.Println(string(body))
			if err := pusher.Push(context.Background(),string(body)); err != nil {
				logx.Errorf("push error: %v", err)
			}
		}
	}
	cmdline.EnterToContinue()

	// 定义消费者配置
	//consumerConf := kq.KqConf{
	//	Brokers: kafkaBrokers,
	//	Topic:   topic,
	//	Group:   "test-group",
	//	Offset:  "first",
	//}
	//q := kq.MustNewQueue(consumerConf, kq.WithHandle(func(k, v string) error {
	//	fmt.Printf("k => %s, v => %+v\n", k, v)
	//	return nil
	//}))
	//defer q.Stop()
	//
	//fmt.Println("消费者启动成功")
	//q.Start()

	// 为了测试，阻塞主线程
	time.Sleep(time.Second * 10)
}
