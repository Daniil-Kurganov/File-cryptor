package src

import (
	"encoding/base64"
	"encoding/binary"
	"file_crypter/src/gost2814789"
	"log"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	reuse "github.com/libp2p/go-reuseport"
)

func HTTPServer() {
	router := gin.Default()
	crypter := router.Group("crypter")
	{
		crypter.POST("encrypt", encrypt)
		crypter.POST("decrypt", decrypt)
	}
	var listener net.Listener
	var err error
	if listener, err = reuse.Listen("tcp", ServerHTTPServeSocket); err != nil {
		log.Fatalf("Error on creating listener: %s", err)
	}
	if err = router.RunListener(listener); err != nil {
		log.Fatalf("Error on starting HTTP-server: %s", err)
	}
}

func encrypt(gctx *gin.Context) {
	var openText string
	var ok bool
	if openText, ok = gctx.GetQuery("data"); !ok {
		log.Fatal("No data for encryption")
	}
	openTextBlocks := gost2814789.TextToUint64Slice(openText)
	closeText := gost2814789.Encryption(openTextBlocks)
	log.Print(closeText)
	log.Print(convertSliceUint64ToUint8(closeText))
	gctx.JSON(http.StatusOK, convertSliceUint64ToUint8(closeText))
}

func decrypt(gctx *gin.Context) {
	var closeText string
	var ok bool
	if closeText, ok = gctx.GetQuery("data"); !ok {
		log.Fatal("No data for encryption")
	}
	b, _ := base64.StdEncoding.DecodeString(closeText)
	openTextBlocks := gost2814789.Decryption(convertSliceUint8ToUint64(b))
	openText := gost2814789.Uint64SliceToText(openTextBlocks)
	log.Print(openText)
	gctx.JSON(http.StatusOK, openText)
}

func convertSliceUint64ToUint8(numbers []uint64) (result []uint8) {
	for _, currentUint64 := range numbers {
		currentBuffer := make([]byte, numberBytesInUint64)
		binary.LittleEndian.PutUint64(currentBuffer, currentUint64)
		result = append(result, currentBuffer...)
	}
	return
}

func convertSliceUint8ToUint64(numbers []uint8) (response []uint64) {
	counterSubText := 0
	for {
		rightBorder := (counterSubText + 1) * numberBytesInUint64
		leftBorder := rightBorder - numberBytesInUint64
		if rightBorder >= len(numbers) {
			byteTextTail := numbers[leftBorder:]
			for {
				if len(byteTextTail) == numberBytesInUint64 {
					break
				}
				byteTextTail = append([]uint8{0}, byteTextTail...)
			}
			response = append(response, binary.LittleEndian.Uint64(byteTextTail))
			break
		}
		response = append(response, binary.LittleEndian.Uint64(numbers[leftBorder:rightBorder]))
		counterSubText += 1
	}
	return
}
