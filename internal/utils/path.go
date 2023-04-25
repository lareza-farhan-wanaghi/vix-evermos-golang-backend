package utils

import (
	"fmt"
	"runtime"
)

// GetFunctionPath returns the path of the caller
func GetFunctionPath() string {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		return fmt.Sprintf("%s:%d", file, line)
	}
	return ""
}

const ProdukImagesPath = "./static/images/produk/"
const TokoImagesPath = "./static/images/toko/"
