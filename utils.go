package fastgo

import (
	"log"
	"mime"
	"path/filepath"
)

func Log(v ...interface{}) {
	log.Println(v...)
}


func MimeType(file string) string {
	return mime.TypeByExtension(filepath.Ext(file))
}
