package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	//"fyne.io/fyne/v2"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	//"github.com/yuin/goldmark/extension"
)

func showGalleryApp() {
	a := app.New()
	w := a.NewWindow("Hello")
	w.Resize(fyne.NewSize(800, 600))
	root_src := "/home/amit/Pictures/"
	files, err := ioutil.ReadDir(root_src)
	if err != nil {
		log.Fatal(err)
	}
	tabs := container.NewAppTabs()
	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())
		if file.IsDir() == false {
			extension := strings.Split(file.Name(), ".")
			if extension[1] == "png" || extension[1] == "jpeg" {

				image := canvas.NewImageFromFile(root_src + "//" + file.Name())
				image.FillMode = canvas.ImageFillContain
				tabs.Append(container.NewTabItem(file.Name(), image))
			}
		}
	}

	//image := canvas.NewImageFromFile(picsArr[0])
	tabs.SetTabLocation(container.TabLocationLeading)
	w.SetContent(tabs)
	w.ShowAndRun()

}
