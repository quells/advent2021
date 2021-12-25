package main

import (
	"encoding/hex"
	"io"
	"math"
)

func parse(input string) (bs bitstream) {
	decoded, _ := hex.DecodeString(input)
	for _, b := range decoded {
		for i := 7; i >= 0; i-- {
			mask := byte(1 << i)
			bit := (b & mask) >> i
			bs = append(bs, bit)
		}
	}
	return
}

type bitstream []byte

func (bs *bitstream) shift(n int) bitstream {
	shifted := (*bs)[:n]
	*bs = (*bs)[n:]
	return shifted
}

func (bs *bitstream) take(n int) (v uint64) {
	taken := bs.shift(n)
	for _, t := range taken {
		v <<= 1
		v |= uint64(t)
	}
	return
}

type packet struct {
	version int
	id      int

	literal int

	subpackets []packet
}

func (p packet) versionSum() int {
	return p.version + versionSum(p.subpackets)
}

func versionSum(packets []packet) int {
	sum := 0
	for _, p := range packets {
		sum += p.versionSum()
	}
	return sum
}

func (p packet) value() int {
	switch p.id {
	case 0:
		sum := 0
		for _, sp := range p.subpackets {
			sum += sp.value()
		}
		return sum

	case 1:
		product := 1
		for _, sp := range p.subpackets {
			product *= sp.value()
		}
		return product

	case 2:
		min := math.MaxInt
		for _, sp := range p.subpackets {
			v := sp.value()
			if v < min {
				min = v
			}
		}
		return min

	case 3:
		max := 0
		for _, sp := range p.subpackets {
			v := sp.value()
			if v > max {
				max = v
			}
		}
		return max

	case 4:
		return p.literal

	case 5:
		a := p.subpackets[0].value()
		b := p.subpackets[1].value()
		if a > b {
			return 1
		}
		return 0

	case 6:
		a := p.subpackets[0].value()
		b := p.subpackets[1].value()
		if a < b {
			return 1
		}
		return 0

	case 7:
		a := p.subpackets[0].value()
		b := p.subpackets[1].value()
		if a == b {
			return 1
		}
		return 0

	default:
		panic("unhandled packet id")
	}
}

func (bs *bitstream) next() (p packet, err error) {
	if len(*bs) < 11 {
		err = io.EOF
		return
	}

	version := int(bs.take(3))
	id := int(bs.take(3))

	if id == 4 {
		return bs.nextLiteral(version, id), nil
	}
	p.version = version
	p.id = id

	lengthTypeID := bs.take(1)
	if lengthTypeID == 0 {
		subpacketLength := bs.take(15)
		substream := bs.shift(int(subpacketLength))
		for {
			if len(substream) == 0 {
				break
			}
			var sp packet
			sp, err = substream.next()
			if err != nil {
				return
			}
			p.subpackets = append(p.subpackets, sp)
		}
	} else {
		subpacketCount := bs.take(11)
		for i := 0; i < int(subpacketCount); i++ {
			var sp packet
			sp, err = bs.next()
			if err != nil {
				return
			}
			p.subpackets = append(p.subpackets, sp)
		}
	}

	return
}

func (bs *bitstream) nextLiteral(version, id int) (p packet) {
	p.version = version
	p.id = id

	for {
		p.literal <<= 4

		chunk := bs.take(5)
		p.literal |= int(chunk & 0xF)

		if chunk&0x10 == 0 {
			break
		}
	}

	return
}

func (bs *bitstream) packets() ([]packet, error) {
	var packets []packet
	for {
		p, err := bs.next()
		if err != nil {
			if err == io.EOF {
				err = nil
				break
			}
			return nil, err
		}

		packets = append(packets, p)
	}

	return packets, nil
}
