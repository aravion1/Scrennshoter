package service

import (
	"context"

	"github.com/aravion1/Scrennshoter/structs"
	"github.com/chromedp/chromedp"
)

type Image struct{}

func NewImageGenerator() *Image {
	return &Image{}
}

var img []byte
var err error

func (ig *Image) GetImage(params structs.Params) ([]byte, error) {

	if params.IsFull {
		err = ig.screenshot(ig.getFullPageImage(params, &img))
	} else {
		err = ig.screenshot(ig.getPageImage(params, &img))
	}

	if err != nil {
		return nil, err
	}

	return img, nil
}

func (ig *Image) GetElementImageByUrl(p structs.Params) ([]byte, error) {
	ig.screenshot(ig.getElementImage(p, &img))
	if err != nil {
		return nil, err
	}
	return img, nil
}

func (ig *Image) getFullPageImage(p structs.Params, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(p.Url),
		chromedp.EmulateViewport(p.Width, p.Height),
		chromedp.FullScreenshot(res, 100),
	}
}

func (ig *Image) getPageImage(p structs.Params, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(p.Url),
		chromedp.EmulateViewport(p.Width, p.Height),
		chromedp.CaptureScreenshot(res),
	}
}

func (ig *Image) getElementImage(p structs.Params, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(p.Url),
		chromedp.EmulateViewport(p.Width, p.Height),
		chromedp.Screenshot(p.Sel, res),
	}
}

func (ig *Image) screenshot(tasks chromedp.Tasks) error {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	if err := chromedp.Run(ctx, tasks); err != nil {
		return err
	}
	return nil
}
