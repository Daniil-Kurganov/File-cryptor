package gost2814789

import (
	"encoding/binary"
	"log"
	"math"
	"math/bits"
)

func firstStep(juniorDataBlock, roundKey uint32) uint32 {
	return uint32(moduloReduction(juniorDataBlock+roundKey, 1, int(math.Pow(2, 32))))
}

func secondStep(inputResult uint32) (result uint32) {
	for counterIteration := 0; counterIteration < countBlocksForSBlock; counterIteration++ { // 2
		currentSlideValue := lengthBlockForSBlock * (countBlocksForSBlock - counterIteration - 1)
		currentSBlockValue := ((inputResult >> uint32(currentSlideValue) << uint32(currentSlideValue)) ^
			(inputResult >> uint32(currentSlideValue+lengthBlockForSBlock) << uint32(currentSlideValue+lengthBlockForSBlock))) >>
			uint32(currentSlideValue)
		result ^= sBlock[counterIteration][currentSBlockValue] << uint32(currentSlideValue)
	}
	return
}

func thirdStep(inputResult uint32) uint32 {
	return bits.RotateLeft32(inputResult, countBitsShift)
}

func fourthStep(intputResult uint32, seniorBlock uint32) uint32 {
	return intputResult ^ seniorBlock
}

func fifthStep(juniorBlock, inputResult uint32) uint64 {
	return uint64(juniorBlock)<<(dataBlockLengthBits/2) ^ uint64(inputResult)
}

func mainStepOfCryptoTransformation(dataBlock uint64, sessionKeyBlock []byte) (result uint64) {
	log.Printf("\nMain step of crypto transformation:\n data = %b\n round key = %v\n\n", dataBlock, sessionKeyBlock)
	roundKey := binary.LittleEndian.Uint32(sessionKeyBlock)
	seniorDataBlock := uint32(dataBlock >> (dataBlockLengthBits / 2))
	juniorDataBlock := uint32(dataBlock)
	log.Printf(" senior block: %b", seniorDataBlock)
	log.Printf(" junior block: %b", juniorDataBlock)
	log.Printf(" round key:    %b", roundKey)
	firstStepResult := firstStep(juniorDataBlock, roundKey)
	log.Print(" steps:")
	log.Printf("  1) %b", firstStepResult)
	secondStepResult := secondStep(firstStepResult)
	log.Printf("  2) %b", secondStepResult)
	thirdStepResult := thirdStep(secondStepResult)
	log.Printf("  3) %b", thirdStepResult)
	fourthStepResult := fourthStep(thirdStepResult, seniorDataBlock)
	log.Printf("  4) %b", fourthStepResult)
	result = fifthStep(juniorDataBlock, fourthStepResult)
	log.Printf("  5) %b", result)
	return
}
