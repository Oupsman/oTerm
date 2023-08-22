package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/fyne-io/terminal"
	"strconv"
)

var Tabs *container.AppTabs
var TabCount int

func NewTab() {
	TabCount++
	tabTitle := "Local Terminal " + strconv.Itoa(TabCount)
	Tabs.Append(container.NewTabItem(tabTitle, NewTerm()))
}

func EditProfiles() {

}

func NewTerm() *terminal.Terminal {
	t := terminal.New()
	go func() {
		_ = t.RunLocalShell()
	}()
	return t
}

func main() {
	oTerm := app.New()
	w := oTerm.NewWindow("oTerm")
	appLayout := layout.NewGridLayoutWithRows(2)
	appContent := container.New(appLayout)
	Tabs = container.NewAppTabs()
	TabCount = 1
	newTabButton := widget.NewButton("New tab", NewTab)
	profilesButton := widget.NewButton("Profiles", EditProfiles)
	leftColumn := container.NewGridWithColumns(2, newTabButton, profilesButton)
	appContent.Add(leftColumn)
	Tabs.Append(container.NewTabItem("Local Terminal 1", NewTerm()))
	Tabs.SetTabLocation(container.TabLocationTop)
	appContent.Add(Tabs)
	w.SetContent(appContent)

	w.ShowAndRun()
}
