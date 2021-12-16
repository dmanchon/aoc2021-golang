package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Packet struct {
	Version int
	TypeID  int
	LenType int
	Value   int
	Packets []Packet
}

func Traverse(packet Packet) int {
	sum := 0
	packets := make([]Packet, 1)
	packets[0] = packet
	for len(packets) > 0 {
		packet = packets[0]
		sum += int(packet.Version)
		packets = append(packets[1:], packet.Packets...)
	}
	return sum
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
		value, _ := strconv.ParseInt(string(num), 2, 32)
		p.Value = int(value)
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

func Solve(lines []string) int {
	packet := NewPacketFromString(lines[0])
	return Traverse(packet)
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
