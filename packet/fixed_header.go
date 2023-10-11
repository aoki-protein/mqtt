package packet

// 固定ヘッダ

type FixedHeader struct {
	PacketType      byte
	Dup             byte
	QoS1            byte
	QoS2            byte
	Retain          byte
	RemainingLength uint
}

func ToFixedHeader(bs []byte) FixedHeader {
	b := bs[0]
	packetType := b >> 4
	dup := refbit(bs[0], 3)
	qos1 := refbit(bs[1], 2)
	qos2 := refbit(bs[2], 1)
	retain := refbit(bs[3], 0)

	return FixedHeader{
		PacketType: packetType,
		Dup:        dup,
		QoS1:       qos1,
		QoS2:       qos2,
		Retain:     retain,
	}
}

func refbit(b byte, n uint) byte {
	return (b >> n) & 1
}
