const (
	body = 0b01111111
	flag = 0b10000000
)

func pack(v int) []byte {
    if v == 0 {
        return []byte{0b00000000}
    }
    res := []byte{}
    u := uint(v)
    for u != 0 {
        b := byte(u & body)
        u >>= 7
        if u != 0 {
            b |= flag
        }
        res = append(res, b)
    }
    return res
}

func unpack(data []byte) (value int, bytesRead int) {
    if len(data) == 0 {
        return 0, 0
    }

    var result uint
    for i, b := range data {
        result |= uint(b&body) << (7 * i)

        if b&flag == 0 {
            return int(result), i + 1
        }
    }
    return 0, 1
}

type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	res := []byte{}
	for _, str := range strs {
		res = append(res, pack(len(str))...)
		res = append(res, str...)
	}
	return string(res)
}

func (s *Solution) Decode(encoded string) []string {
	data := []byte(encoded)
	res := []string{}
	i := 0
	for i < len(data) {
		l, n := unpack(data[i:])
		if n == 0 {
			break
		}
		str := string(data[i+n : i+n+l])
		res = append(res, str)
		i += n + l
	}
	return res
}
