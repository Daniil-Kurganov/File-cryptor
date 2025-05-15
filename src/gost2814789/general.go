package gost2814789

var sessionKey [][]byte

func Encryption(openTextBlocks []uint64) (closeTextBlocks []uint64) {
	sessioinKeyGeneration()
	for _, currentOpenTextBlock := range openTextBlocks {
		currentCloseBlock := currentOpenTextBlock
		for counterIteration := 0; counterIteration < numberIteraionStraitKeySplit; counterIteration++ {
			for counterIndex := 0; counterIndex < numberRoundKeys; counterIndex++ {
				currentCloseBlock = mainStepOfCryptoTransformation(currentCloseBlock, sessionKey[counterIndex])
			}
		}
		for counterIndex := numberRoundKeys - 1; counterIndex >= 0; counterIndex-- {
			currentCloseBlock = mainStepOfCryptoTransformation(currentCloseBlock, sessionKey[counterIndex])
		}
		currentCloseBlock = halfBitsSwap(currentCloseBlock)
		closeTextBlocks = append(closeTextBlocks, currentCloseBlock)
	}
	return
}

func Decryption(closeTextBlocks []uint64) (openTextBlocks []uint64) {
	for _, currentCloseBlock := range closeTextBlocks {
		currentOpenBlock := currentCloseBlock
		for counterIndex := 0; counterIndex < numberRoundKeys; counterIndex++ {
			currentOpenBlock = mainStepOfCryptoTransformation(currentOpenBlock, sessionKey[counterIndex])
		}
		for counterIteration := 0; counterIteration < numberIteraionStraitKeySplit; counterIteration++ {
			for counterIndex := numberRoundKeys - 1; counterIndex >= 0; counterIndex-- {
				currentOpenBlock = mainStepOfCryptoTransformation(currentOpenBlock, sessionKey[counterIndex])
			}
		}
		currentOpenBlock = halfBitsSwap(currentOpenBlock)
		openTextBlocks = append(openTextBlocks, currentOpenBlock)
	}
	return
}
