package packet

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIdentifiedPacketDecode(t *testing.T) {
	pktBytes := []byte{
		byte(PUBACK << 4),
		2,
		0, // packet ID MSB
		7, // packet ID LSB
	}

	n, pid, err := identifiedPacketDecode(pktBytes, PUBACK)
	assert.NoError(t, err)
	assert.Equal(t, 4, n)
	assert.Equal(t, ID(7), pid)
}

func TestIdentifiedPacketDecodeError1(t *testing.T) {
	pktBytes := []byte{
		byte(PUBACK << 4),
		1, // < wrong remaining length
		0, // packet ID MSB
		7, // packet ID LSB
	}

	n, pid, err := identifiedPacketDecode(pktBytes, PUBACK)
	assert.Error(t, err)
	assert.Equal(t, 2, n)
	assert.Equal(t, ID(0), pid)
}

func TestIdentifiedPacketDecodeError2(t *testing.T) {
	pktBytes := []byte{
		byte(PUBACK << 4),
		2,
		7, // packet ID LSB
		// < insufficient bytes
	}

	n, pid, err := identifiedPacketDecode(pktBytes, PUBACK)
	assert.Error(t, err)
	assert.Equal(t, 2, n)
	assert.Equal(t, ID(0), pid)
}

func TestIdentifiedPacketDecodeError3(t *testing.T) {
	pktBytes := []byte{
		byte(PUBACK << 4),
		2,
		0, // packet ID LSB
		0, // packet ID MSB < zero id
	}

	n, pid, err := identifiedPacketDecode(pktBytes, PUBACK)
	assert.Error(t, err)
	assert.Equal(t, 4, n)
	assert.Equal(t, ID(0), pid)
}

func TestIdentifiedPacketEncode(t *testing.T) {
	pktBytes := []byte{
		byte(PUBACK << 4),
		2,
		0, // packet ID MSB
		7, // packet ID LSB
	}

	dst := make([]byte, identifiedPacketLen())
	n, err := identifiedPacketEncode(dst, 7, PUBACK)

	assert.NoError(t, err)
	assert.Equal(t, 4, n)
	assert.Equal(t, pktBytes, dst[:n])
}

func TestIdentifiedPacketEncodeError1(t *testing.T) {
	dst := make([]byte, 3) // < insufficient buffer
	n, err := identifiedPacketEncode(dst, 7, PUBACK)

	assert.Error(t, err)
	assert.Equal(t, 0, n)
}

func TestIdentifiedPacketEncodeError2(t *testing.T) {
	dst := make([]byte, identifiedPacketLen())
	n, err := identifiedPacketEncode(dst, 0, PUBACK) // < zero id

	assert.Error(t, err)
	assert.Equal(t, 0, n)
}

func TestIdentifiedPacketEqualDecodeEncode(t *testing.T) {
	pktBytes := []byte{
		byte(PUBACK << 4),
		2,
		0, // packet ID MSB
		7, // packet ID LSB
	}

	pkt := &PubackPacket{}
	n, err := pkt.Decode(pktBytes)

	assert.NoError(t, err)
	assert.Equal(t, 4, n)

	dst := make([]byte, 100)
	n2, err := identifiedPacketEncode(dst, 7, PUBACK)

	assert.NoError(t, err)
	assert.Equal(t, 4, n2)
	assert.Equal(t, pktBytes, dst[:n2])

	n3, pid, err := identifiedPacketDecode(pktBytes, PUBACK)
	assert.NoError(t, err)
	assert.Equal(t, 4, n3)
	assert.Equal(t, ID(7), pid)
}

func BenchmarkIdentifiedPacketEncode(b *testing.B) {
	pkt := &PubackPacket{}
	pkt.ID = 1

	buf := make([]byte, pkt.Len())

	for i := 0; i < b.N; i++ {
		_, err := pkt.Encode(buf)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkIdentifiedPacketDecode(b *testing.B) {
	pktBytes := []byte{
		byte(PUBACK << 4),
		2,
		0, // packet ID MSB
		1, // packet ID LSB
	}

	pkt := &PubackPacket{}

	for i := 0; i < b.N; i++ {
		_, err := pkt.Decode(pktBytes)
		if err != nil {
			panic(err)
		}
	}
}

func testIdentifiedPacketImplementation(t *testing.T, pkt GenericPacket) {
	assert.Equal(t, fmt.Sprintf("<%sPacket ID=1>", pkt.Type().String()), pkt.String())

	buf := make([]byte, pkt.Len())
	n, err := pkt.Encode(buf)
	assert.NoError(t, err)
	assert.Equal(t, 4, n)

	n, err = pkt.Decode(buf)
	assert.NoError(t, err)
	assert.Equal(t, 4, n)
}

func TestPubackImplementation(t *testing.T) {
	pkt := NewPubackPacket()
	pkt.ID = 1

	testIdentifiedPacketImplementation(t, pkt)
}

func TestPubcompImplementation(t *testing.T) {
	pkt := NewPubcompPacket()
	pkt.ID = 1

	testIdentifiedPacketImplementation(t, pkt)
}

func TestPubrecImplementation(t *testing.T) {
	pkt := NewPubrecPacket()
	pkt.ID = 1

	testIdentifiedPacketImplementation(t, pkt)
}

func TestPubrelImplementation(t *testing.T) {
	pkt := NewPubrelPacket()
	pkt.ID = 1

	testIdentifiedPacketImplementation(t, pkt)
}

func TestUnsubackImplementation(t *testing.T) {
	pkt := NewUnsubackPacket()
	pkt.ID = 1

	testIdentifiedPacketImplementation(t, pkt)
}
