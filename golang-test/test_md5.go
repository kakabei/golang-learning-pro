package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	// 计算BodyMd5

	data := []byte("runtime_pool_v39/cloud-desk-os-2014-1244-1")
	fmt.Printf("%v", hex.EncodeToString(md5.Sum(data)))

}
