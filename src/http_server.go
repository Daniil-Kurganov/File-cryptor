package src

import (
	"encoding/json"
	"file_crypter/src/gost2814789"
	"io/ioutil"
	"log"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	reuse "github.com/libp2p/go-reuseport"
)

type testRequest struct {
	Data []uint64 `json:"data"`
}

func HTTPServer() {
	router := gin.Default()
	crypter := router.Group("crypter")
	{
		crypter.GET("start", start)
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

func start(gctx *gin.Context) {
	var err error
	var requestBytes []byte
	if requestBytes, err = ioutil.ReadAll(gctx.Request.Body); err != nil {
		log.Fatal(err.Error())
	}
	var request testRequest
	json.Unmarshal(requestBytes, &request)
	gctx.JSON(http.StatusOK, gost2814789.Encryption(request.Data))
}
