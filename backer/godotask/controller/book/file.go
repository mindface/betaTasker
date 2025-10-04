package book

import (
	"crypto/sha1"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/olahol/go-imageupload"
)
type UploadInput struct {
    Description string `form:"description" binding:"max=200"`
}

func HundleUplond(c *gin.Context) {
	// バリデーション: ファイルがアップロードされているか
	var input UploadInput
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid form data",
			"detail":  err.Error(),
		})
		return
	}

	// ファイルがアップロードされているか
	file, err := c.FormFile("upfile")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "No file selected",
		})
		return
	}

	// 拡張子チェック（小文字に変換して比較）
	ext := strings.ToLower(filepath.Ext(file.Filename))
	allowedExt := map[string]bool{".jpg": true, ".jpeg": true, ".png": true}
	if !allowedExt[ext] {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Unsupported file type (only jpg, jpeg, png allowed)",
		})
		return
	}

	// 画像処理
	img, err := imageupload.Process(c.Request, "upfile")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Failed to process image",
			"detail":  err.Error(),
		})
		return
	}

	thumb, err := imageupload.ThumbnailPNG(img, 300, 300)
	if err != nil {
		panic(err)
	}
	h := sha1.Sum(thumb.Data)
	thumbPath := fmt.Sprintf("./static/images/%s_%x.png",
	time.Now().Format("20060102150405"), h[4])
	if err := thumb.Save(thumbPath); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"msg": "unable to save file",
		})
		return
	}

	// extension := filepath.Ext(file.Filename)
	newFileName := "t" + ext
	uploadPath := "./static/uploads/" + newFileName

	if err := c.SaveUploadedFile(file, uploadPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "ファイルの保存に失敗しました",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":      "success",
		"message":     "ファイルをアップロードしました",
		"uploadPath":  uploadPath,
		"thumbPath":   thumbPath,
		"originalName": file.Filename,
	})

	// c.JSON(http.StatusOK, gin.H{
	// 	"message": "Successfully upload",
	// })

}
