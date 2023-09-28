package gomksservo42c

func AppendChecksum(input []uint8) []uint8 {
	checksum := CalcChecksum(input)

	return append(input, checksum)
}

func CheckChecksum(input []uint8) bool {
	len := len(input)
	checksum := CalcChecksum(input[:len-1])

	return checksum == input[len-1]
}

func CalcChecksum(input []uint8) uint8 {
	var sum uint64

	for _, value := range input {
		sum += uint64(value) //sum up the values
	}

	return uint8(sum & 0xFF) //bitwise AND
}
