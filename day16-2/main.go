package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

type Packet struct {
	Version int
	TypeID  int
	LenType int
	Value   int64
	Packets []Packet
}

func Execute(packet Packet) int64 {
	var result int64
	switch packet.TypeID {
	case 0:
		// sum
		for _, p := range packet.Packets {
			result += Execute(p)
		}
		return result
	case 1:
		// product
		result = 1
		for _, p := range packet.Packets {
			result *= Execute(p)

		}
		return result
	case 2:
		// min
		result = math.MaxInt
		for _, p := range packet.Packets {
			value := Execute(p)
			if value < result {
				result = value
			}
		}
		return result
	case 3:
		// max
		result = 0
		for _, p := range packet.Packets {
			value := Execute(p)
			if value > result {
				result = value
			}
		}
		return result
	case 4:
		return packet.Value
	case 5:
		// gt
		v1 := Execute(packet.Packets[0])
		v2 := Execute(packet.Packets[1])
		if v1 > v2 {
			return 1
		}
	case 6:
		// lt
		v1 := Execute(packet.Packets[0])
		v2 := Execute(packet.Packets[1])
		if v1 < v2 {
			return 1
		}
	case 7:
		// eq
		v1 := Execute(packet.Packets[0])
		v2 := Execute(packet.Packets[1])
		if v1 == v2 {
			return 1
		}
	}
	return result
}

func NewPacketFromString(s string) Packet {
	input := make([]byte, 0)
	bytes, _ := hex.DecodeString(s)

	for _, b := range bytes {
		s := fmt.Sprintf("%08b", b)
		input = append(input, s...)
	}
	p, _ := NewPacketFromBinary(input)
	return p
}

func NewPacketFromBinary(input []byte) (Packet, []byte) {
	p := Packet{}

	version, _ := strconv.ParseInt(string(input[:3]), 2, 32)
	p.Version = int(version)

	id, _ := strconv.ParseInt(string(input[3:6]), 2, 32)
	p.TypeID = int(id)

	offset := 6

	if p.TypeID == 4 {
		num := make([]byte, 0)
		for {
			num = append(num, input[offset+1:offset+5]...)
			offset += 5
			if input[offset-5] == '0' {
				break
			}
		}
		value, _ := strconv.ParseInt(string(num), 2, 64)
		p.Value = value
		if len(input) < 11 {
			return p, []byte{}
		} else {
			return p, input[offset:]
		}
	} else {
		if input[offset] == '0' {
			p.LenType = 0
			lenght, _ := strconv.ParseInt(string(input[7:22]), 2, 32)

			input = input[22:]
			size := len(input) - int(lenght)

			for len(input) > size {
				sp, remaining := NewPacketFromBinary(input)
				p.Packets = append(p.Packets, sp)
				input = remaining
			}

			return p, input
		} else {
			p.LenType = 1
			lenght, _ := strconv.ParseInt(string(input[7:18]), 2, 32)
			input = input[18:]
			for lenght > 0 {
				lenght--
				sp, remaining := NewPacketFromBinary(input)
				p.Packets = append(p.Packets, sp)
				input = remaining
			}
			return p, input
		}

	}
}

func Solve(lines []string) int64 {
	packet := NewPacketFromString(lines[0])
	return Execute(packet)
}

func ReadInput(file *os.File) []string {
	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func main() {
	start := time.Now()
	defer func() {
		fmt.Println("- Elapsed Time: ", time.Now().Sub(start))
	}()
	fmt.Println("- Solution: ", Solve(ReadInput(os.Stdin)))
}
