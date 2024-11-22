package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/zinego/go-tools/utils/log"
)

func main() {
	log.Init(log.WithFileName(fmt.Sprintf("%s/static.log", os.TempDir())))
	var f = flag.String("f", "filepath", "filepath, format: a=/a,b=/b,c=/c")
	var p = flag.String("p", "port", "9999")
	flag.Parse()
	if _, err := strconv.Atoi(*p); err != nil {
		log.Errorf("invalid port %s", *p)
		os.Exit(1)
	}
	router := gin.Default()
	fileList := strings.Split(*f, ",")
	for _, kv := range fileList {
		if strings.Count(kv, "=") != 1 {
			log.Errorf("invalid file format: %s", kv)
			continue
		}
		var k = strings.Split(kv, "=")[0]
		var v = strings.Split(kv, "=")[1]
		router.StaticFS(k, http.Dir(v))
	}
	// 监听并在 0.0.0.0:8080 上启动服务
	router.Run(":" + *p)
}
