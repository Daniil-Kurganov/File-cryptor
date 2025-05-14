package gost2814789

import (
	"crypto/rand"
	"math/big"
)

const (
	numberIteraionStraitKeySplit                               = 3
	lengthBlockForSBlock, lenghtRoundKeyByte                   = 4, 4
	dataBlockLengthByte, countBlocksForSBlock, numberRoundKeys = 8, 8, 8
	countBitsShift                                             = 11
	slideBits, numberKeyBlocks                                 = 32, 32
	dataBlockLengthBits                                        = 64
)

var sBlock = [][]uint32{ // [columns, rows]
	{1, 15, 13, 0, 5, 7, 10, 4, 9, 2, 3, 14, 6, 11, 8, 12},  // 8 in table -> index == 0
	{13, 11, 4, 1, 3, 15, 5, 9, 0, 10, 14, 7, 6, 8, 2, 12},  // 7 in table -> index == 1
	{4, 11, 10, 0, 7, 2, 1, 13, 3, 6, 8, 5, 9, 12, 15, 14},  // 6 in table -> index == 2
	{6, 12, 7, 1, 5, 15, 13, 8, 4, 10, 9, 14, 0, 3, 11, 2},  // 5 in table -> index == 3
	{7, 13, 10, 1, 0, 8, 9, 15, 14, 4, 6, 12, 11, 2, 5, 3},  // 4 in table -> index == 4
	{5, 8, 1, 13, 10, 3, 4, 2, 14, 15, 12, 7, 6, 0, 99, 11}, // 3 in table -> index == 5
	{14, 11, 4, 12, 6, 13, 15, 10, 2, 3, 8, 1, 0, 7, 5, 9},  // 2 in table -> index == 6
	{4, 10, 9, 2, 13, 8, 0, 14, 6, 11, 1, 12, 7, 15, 5, 3},  // 1 in table -> index == 7
}

func halfBitsSwap(number uint64) (result uint64) {
	leftPart := number >> slideBits
	rightPart := ((leftPart << slideBits) ^ number)
	result = (rightPart << slideBits) ^ leftPart
	return
}

func sessioinKeyGeneration() {
	buffer := make([]byte, numberKeyBlocks)
	rand.Read(buffer)
	for counterIteration := 0; counterIteration < numberRoundKeys; counterIteration++ {
		leftBorder := counterIteration * lenghtRoundKeyByte
		rightBorder := leftBorder + lenghtRoundKeyByte
		sessionKey = append(sessionKey, buffer[leftBorder:rightBorder])
	}
}

func moduloReduction[T uint64 | uint32 | int](numberInt, power T, module int) (remainder T) {
	numberBig := big.NewInt(int64(numberInt))
	numberInPower := numberBig.Exp(numberBig, big.NewInt(int64(power)), nil)
	_, resultBig := new(big.Int).DivMod(numberInPower, big.NewInt(int64(module)), new(big.Int))
	return T(resultBig.Int64())
}
