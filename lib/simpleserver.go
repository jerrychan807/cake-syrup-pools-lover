package lib

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
)

func StartSimpleServer() {
	AllConfig := GetConfig()
	publicFolderPath := filepath.Join(AllConfig.ProjectFolder, "/public/")
	//fmt.Printf("[*] publicFolderPath: %s ", publicFolderPath)
	fmt.Println("[*] start simple http server")
	// 启动web服务，监听9090端口
	err := http.ListenAndServe("127.0.0.1:9090", http.FileServer(http.Dir(publicFolderPath)))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
