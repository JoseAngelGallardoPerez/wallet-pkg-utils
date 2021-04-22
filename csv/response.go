package csv

import (
	"github.com/gin-gonic/gin"
)

// Send sends file with fileName
func Send(file *File, writer gin.ResponseWriter) error {
	bytes, err := file.Bytes()
	if err != nil {
		return err
	}

	writer.Header().Set("Content-Type", "text/csv")
	writer.Header().Set("Content-Disposition", "attachment;filename="+file.Name)
	writer.Header().Set("Access-Control-Expose-Headers", "Content-Disposition")
	writer.Write(bytes)
	return nil
}
