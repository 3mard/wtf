package pocket

import (
	"fmt"
	"io/ioutil"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/logger"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/view"
	"gopkg.in/yaml.v2"
)

const HelpText = `
 Keyboard commands for NewTextWidget:
`

type Widget struct {
	view.ScrollableWidget
	view.KeyboardWidget

	settings *Settings
	client   *Client
}

func NewWidget(app *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		KeyboardWidget:   view.NewKeyboardWidget(app, pages, settings.common),
		ScrollableWidget: view.NewScrollableWidget(app, settings.common),
		settings:         settings,
		client:           NewClient(settings.consumerKey, "http://localhost"),
	}

	widget.CommonSettings()

	widget.View.SetInputCapture(widget.InputCapture)

	widget.SetRenderFunction(widget.Render)
	widget.View.SetScrollable(true)
	widget.View.SetRegions(true)
	widget.KeyboardWidget.SetView(widget.View)
	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Render() {
	// The last call should always be to the display function
	widget.Redraw(widget.content)
}

func (widget *Widget) Refresh() {
	widget.Redraw(widget.content)
}

/* -------------------- Unexported Functions -------------------- */

type PocketMetaData struct {
	AccessToken *string
}

func writeMetaDataToDisk(token PocketMetaData) error {

	fileData, err := yaml.Marshal(token)
	if err != nil {
		return fmt.Errorf("Could not write token to disk %v", err)
	}

	wtfConfigDir, err := cfg.WtfConfigDir()

	if err != nil {
		return nil
	}

	filePath := fmt.Sprintf("%s/%s", wtfConfigDir, "pocket.data")
	err = ioutil.WriteFile(filePath, fileData, 0644)

	return err
}

func readMetaDataFromDisk() (PocketMetaData, error) {
	wtfConfigDir, err := cfg.WtfConfigDir()
	var metaData PocketMetaData
	if err != nil {
		return metaData, err
	}
	filePath := fmt.Sprintf("%s/%s", wtfConfigDir, "pocket.data")
	fileData, err := utils.ReadFileBytes(filePath)

	if err != nil {
		return metaData, err
	}

	err = yaml.Unmarshal(fileData, &metaData)

	return metaData, err

}

func (widget *Widget) authorizeWorkFlow() (string, string, bool) {
	title := widget.CommonSettings().Title

	if widget.settings.requestKey == nil {
		requestToken, err := widget.client.ObtainRequestToken()

		if err != nil {
			logger.Log(err.Error())
			return title, err.Error(), true
		}
		widget.settings.requestKey = &requestToken
		redirectURL := widget.client.CreateRedirectLink(requestToken)
		content := fmt.Sprintf("Please click on %s to Authorize the app", redirectURL)
		return title, content, true
	}

	if widget.settings.accessToken == nil {
		accessToken, err := widget.client.getAccessToken(*widget.settings.requestKey)
		if err != nil {
			logger.Log(err.Error())
			redirectURL := widget.client.CreateRedirectLink(*widget.settings.requestKey)
			content := fmt.Sprintf("Please click on %s to Authorize the app", redirectURL)
			return title, content, true
		}
		content := "Authorized"
		widget.settings.accessToken = &accessToken

		metaData := PocketMetaData{
			AccessToken: &accessToken,
		}

		err = writeMetaDataToDisk(metaData)
		if err != nil {
			content = err.Error()
		}

		return title, content, true
	}

	content := "Authorized"
	return title, content, true

}

func (widget *Widget) content() (string, string, bool) {
	title := widget.CommonSettings().Title
	if widget.settings.accessToken == nil {
		metaData, err := readMetaDataFromDisk()
		if err != nil || metaData.AccessToken == nil {
			return widget.authorizeWorkFlow()
		}
		widget.settings.accessToken = metaData.AccessToken
	}

	content := "done" + *widget.settings.accessToken
	return title, content, true
}
