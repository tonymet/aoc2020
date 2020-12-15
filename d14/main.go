package main

import (
	"flag"
	"fmt"
	"io"
	"strconv"
)

// good
// bitmask worked well

// bad
// too long to parse

type memType map[int64]int64
type memStorage struct {
	mem                     memType
	trueMask, outboundValue int64
}

func sumMem(mem memType) int64 {
	var sum int64
	for _, v := range mem {
		sum += v
	}
	return sum
}

func scanFile() memStorage {
	ms := memStorage{make(memType), 0, 0}

	for {
		var cmd, arg string
		_, err := fmt.Scanf("%s = %s", &cmd, &arg)
		if err == io.EOF {
			break
		}
		//fmt.Printf("cmd: %s, arg: %s \n", cmd, arg)
		switch cmd[0:3] {
		case "mem":
			memAddr := cmd[4 : len(cmd)-1]
			//fmt.Printf("memAddr: %s\n", memAddr)
			if memValueInt, err := strconv.ParseInt(arg, 10, 64); err != nil {
				panic(err)
			} else if memAddrInt, err := strconv.ParseInt(memAddr, 10, 64); err != nil {
				panic(err)
			} else {
				ms.mem[memAddrInt] = memValue(memValueInt, ms.outboundValue, ms.trueMask)
			}
		case "mas":
			//fmt.Printf("read mask %s\n", arg)
			ms.outboundValue, ms.trueMask = parseMask(arg)
		default:
			panic("wrong command")
		}
	}
	return ms
}

func parseMask(mask string) (int64, int64) {
	var outboundValue, trueMask int64
	for pos := 0; pos < len(mask); pos++ {
		// read char
		//fmt.Printf("c: %b", mask[pos])
		outboundValue <<= 1
		trueMask <<= 1
		switch mask[pos] {
		case '1':
			outboundValue |= 1
			// set 1 on outbound value
		case '0':
			outboundValue |= 0
			// set 0 on outbound value
		case 'X':
			trueMask |= 1
		default:
			panic("bad mask")
		}
	}
	return outboundValue, trueMask
}

func memValue(inboundValue, outboundValue, trueMask int64) int64 {
	// first apply outboundValue
	// then maks inbound & trueMask,
	// then outboundValue & that
	var memValue int64
	memValue |= outboundValue
	memValue |= (inboundValue & trueMask)
	return memValue
}

func main() {
	var pFlag = flag.Int("p", 1, "1 or 2")
	flag.Parse()
	switch *pFlag {
	case 1:
		ms := scanFile()
		//fmt.Printf("ms: %+v\n", ms)
		fmt.Printf("sum: %+v\n", sumMem(ms.mem))
	case 2:
		fmt.Printf("Part 2\n")
	case 3:
		{
			o, m := parseMask("X100110110X011000101000101XX11001X11")
			fmt.Printf("obValue: %b, truemask %b\n", o, m)
			o, m = parseMask("XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X")
			ibValue := int64(11)
			memValue := memValue(ibValue, o, m)
			fmt.Printf("obValue: %b, truemask %b, memValue: %b\n", o, m, memValue)
		}
		{
			o, m := parseMask("X100110110X011000101000101XX11001X11")
			fmt.Printf("obValue: %b, truemask %b\n", o, m)
			o, m = parseMask("XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X")
			ibValue := int64(11)
			memValue := memValue(ibValue, o, m)
			fmt.Printf("obValue: %b, truemask %b, memValue: %b\n", o, m, memValue)
		}
		//
		{
			o, m := parseMask("XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X")
			ibValue := int64(0)
			memValue := memValue(ibValue, o, m)
			fmt.Printf("obValue: %b, truemask %b, memValue: %b\n", o, m, memValue)
		}
	case 4:
		scanFile()
	default:
		panic("-p 1 or 2")
	}
}
