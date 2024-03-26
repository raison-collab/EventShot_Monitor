package video_maker

import (
	"EventShot_Monitor/config"
	"EventShot_Monitor/errors"
	"EventShot_Monitor/utils"
	"github.com/icza/mjpeg"
	"log"
	"os"
)

func RenderVideo(config config.Config) error {
	hasFiles, err := utils.HasFilesInScreenshotDir(config)
	if err != nil {
		return err
	}

	if !hasFiles {
		log.Printf("[DIR] Дирректория пуста")
		return &errors.EmptyDirError{Message: "Пустая дирректория со скриншотами"}
	}

	fileNames, err := utils.GetScreenshotsFilenames(config)
	if err != nil {
		return err
	}

	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}

	aw, err := mjpeg.New(currentDir+config.VideoDir+"/avi_01.avi", 200, 100, 2)
	if err != nil {
		return err
	}

	for _, fileName := range fileNames {
		data, err := os.ReadFile(currentDir + config.ScreenshotDir + "/" + fileName)
		if err != nil {
			return err
		}
		err = aw.AddFrame(data)
		if err != nil {
			return err
		}
	}

	err = aw.Close()
	if err != nil {
		return err
	}

	return nil
}
