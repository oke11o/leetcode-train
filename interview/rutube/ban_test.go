package rutube

import (
	"github.com/stretchr/testify/require"
	"math"
	"strconv"
	"strings"
	"testing"
)

// написать функцию осуществляющую поиск по забанненным подсетям/ip-адресам
// 192.168.1.0/24 - ban
// 10.1.22.0/31 - ban
// 11.14.0.0/16 - ban
// 192.169.1.0/28 - ban
// 1.1.1.1 - ban
// 2.2.2.2 - ban
// 192 11010010

type Node struct {
	isLeaf   bool
	children [2]*Node
}

type List struct {
	root *Node
}

func (n *Node) GetChild(bit bool) *Node {
	return n.children[n.idx(bit)]
}

func (n *Node) Add(bit bool) *Node {
	//n.m.Lock()
	//defer n.m.Unlock()
	res := &Node{}
	n.children[n.idx(bit)] = res
	return res
}

func (n *Node) idx(bit bool) int {
	idx := 0
	if bit {
		idx = 1
	}
	return idx
}

func (l *List) AddIp(IP string) {
	if l.root == nil {
		l.root = &Node{}
	}
	ip, mask := IP2Bits(IP)

	ptr := uint32(math.MaxUint32)
	ptr = ptr >> 1
	ptr++
	tmp := strconv.FormatInt(int64(ptr), 2)
	_ = tmp
	node := l.root
	for ptr > 0 {
		isBit := ip & ptr
		isMask := mask & ptr
		_ = isMask
		bit := false
		if isBit > 0 {
			bit = true
		}
		ch := node.Add(bit)
		if isMask == 0 {
			ch.isLeaf = true
			break
		}
		node = ch
		ptr = ptr >> 1
	}
}

func (l *List) Exists(ip uint32) bool {
	ptr := uint32(math.MaxUint32)
	ptr = ptr >> 1
	ptr++
	node := l.root
	for ptr > 0 {
		isBit := ip & ptr
		bit := false
		if isBit > 0 {
			bit = true
		}
		n := node.GetChild(bit)
		if n == nil {
			return node.isLeaf
		}
		if n.isLeaf {
			return true
		}
		node = n
	}

	return false
}

func ValidateIP(IP string) bool {
	ip, _ := IP2Bits(IP)
	ok := BlackList.Exists(ip)
	return !ok
}

func IP2Bits(input string) (ip uint32, mask uint32) {
	inputL := strings.Split(input, "/")
	maskShift := 32
	if len(inputL) > 1 {
		maskShift, _ = strconv.Atoi(inputL[1])
	}
	mask = uint32(math.MaxUint32)
	mask = mask << (32 - maskShift)
	//tmp := strconv.FormatInt(int64(mask), 2)
	//_ = tmp
	ips := strings.Split(inputL[0], ".")
	for i, ipPart := range ips {
		ipParsed, _ := strconv.Atoi(ipPart)
		ip += uint32(ipParsed)
		if i < 3 {
			ip = ip << 8
		}
	}

	return ip, mask
}

var BlackList *List

func Test_IP2Bits(t *testing.T) {

	var tests = []struct {
		name     string
		input    string
		wantIP   uint32
		wantMask uint32
	}{
		{
			name:     "255.255.255.255/24",
			input:    "255.255.255.255/24",
			wantIP:   4294967295,
			wantMask: 4294967040,
		},
		{
			name:     "192.168.0.1/24",
			input:    "192.168.0.1/24",
			wantIP:   3232235521,
			wantMask: 4294967040,
		},
		{
			name:     "192.168.0.1",
			input:    "192.168.0.1",
			wantIP:   3232235521,
			wantMask: 4294967295,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIP, gotMask := IP2Bits(tt.input)
			require.Equal(t, tt.wantIP, gotIP)
			require.Equal(t, tt.wantMask, gotMask)
		})
	}
}

func Test_Ban(t *testing.T) {
	blackIPs := []string{
		"127.10.0.1/24",
		"192.168.0.1/24",
		"192.168.0.1",
		"255.255.255.255/24",
	}
	bl := &List{}
	for _, ip := range blackIPs {
		bl.AddIp(ip)
	}

	ip := "127.10.0.2"
	i, _ := IP2Bits(ip)

	ok := bl.Exists(i)

	println(ok)
}
