package book

import (
	"crypto/sha1"
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/olahol/go-imageupload"
)

func HundleUplond(c *gin.Context) {
	file, _ := c.FormFile("upfile")
	img, err := imageupload.Process(c.Request, "upfile")
	if err != nil {
		panic(err)
	}

	thumb, err := imageupload.ThumbnailPNG(img, 300, 300)
	if err != nil {
		panic(err)
	}
	h := sha1.Sum(thumb.Data)
	thumb.Save(fmt.Sprintf("/usr/local/go/godotask/static/images/%s_%x.png",
		time.Now().Format("20060102150405"), h[4]))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": "unable to save file",
		})
		return
	}

	extension := filepath.Ext(file.Filename)
	newFileName := "t" + extension

	if err := c.SaveUploadedFile(file, "/usr/local/go/godotask/static/uploads/"+newFileName); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file",
		})
		return
	}

	// c.JSON(http.StatusOK, gin.H{
	// 	"message": "Successfully upload",
	// })

}
