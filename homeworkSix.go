package main

/*
发送数据过程：应用程序发送消息包，消息包以数据流的形式放入缓冲区，等缓冲区的数据流到达一定阈值后，再发送到网络上

接受数据过程：接受到网络过来的数据流，放入缓冲区，缓冲区的数据流到达一定阈值后，通知应用程序进行读取数据

在数据发送和接受的过程中，都是对数据流进行操作，而数据流本身没有任何开始和结束的边界。因此正确地解析出数据包，就要知道数据在流中的开始和结束位置。

fix length frame decoder
数据发送方每次发送固定长度的数据，且不超出缓冲区，接收方获取同样长度的数据来解码拼成一个数据包。

delimiter based frame decoder
数据发送方在数据中添加特殊的分隔符来标记边界，接收方读到分隔符时解码拼成一个数据包。

length field based frame decoder
数据发送方在消息包头添加长度信息，接收方获取指定长度的数据来解码拼成一个数据包。
*/

import (
	"context"
	"encoding/binary"
	"fmt"
)

type homeworkSix struct {
}

/*
goim 协议结构
4bytes PacketLen 包长度，在数据流传输过程中，先写入整个包的长度，方便整个包的数据读取。
2bytes HeaderLen 头长度，在处理数据时，会先解析头部，可以知道具体业务操作。
2bytes Version 协议版本号，主要用于上行和下行数据包按版本号进行解析。
4bytes Operation 业务操作码，可以按操作码进行分发数据包到具体业务当中。
4bytes Sequence 序列号，数据包的唯一标记，可以做具体业务处理，或者数据包去重。
PacketLen-HeaderLen Body 实际业务数据，在业务层中会进行数据解码和编码。
*/

func (h *homeworkSix) homeworkDoSomething(ctx context.Context) error {
	ss := "hello word"
	zz := encoder([]byte(ss))
	decoder(zz)
	return nil
}

func decoder(data []byte) {
	if len(data) < 16 {
		fmt.Println("data is error")
	}
	fmt.Println(binary.BigEndian.Uint32(data[0:4]))
	fmt.Println(binary.BigEndian.Uint16(data[4:6]))
	fmt.Println(binary.BigEndian.Uint16(data[6:8]))
	fmt.Println(binary.BigEndian.Uint32(data[8:12]))
	fmt.Println(binary.BigEndian.Uint32(data[12:16]))
	fmt.Println(string(data[16:]))

}

func encoder(data []byte) []byte {
	packLen := 16 + uint32(len(data))
	ret := make([]byte, packLen)
	binary.BigEndian.PutUint32(ret[0:4], packLen)
	binary.BigEndian.PutUint16(ret[4:6], uint16(1))
	binary.BigEndian.PutUint16(ret[6:8], uint16(1))
	binary.BigEndian.PutUint32(ret[8:12], 123)
	binary.BigEndian.PutUint32(ret[12:16], 234235)
	ret = append(ret, data...)
	return ret
}

func (h *homeworkSix) getName() string {
	return "homeworkSix"
}
