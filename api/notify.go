package api

import (
	"fmt"
)

type NotifyRequest struct {
	Text       string `json:"text,omitempty"`
	TextCase   int    `json:"textCase,omitempty"`
	TopText    bool   `json:"topText,omitempty"`
	TextOffset int    `json:"textOffset,omitempty"`
	Center     bool   `json:"center,omitempty"`
	Color      string `json:"color,omitempty"`
	BlinkText  int    `json:"blinkText,omitempty"`
	FadeText   int    `json:"fadeText,omitempty"`
	Background string `json:"background,omitempty"`
	Rainbow    bool   `json:"rainbow,omitempty"`
	Icon       string `json:"icon,omitempty"`
	PushIcon   int    `json:"pushIcon,omitempty"`
	Repeat     int    `json:"repeat,omitempty"`
	Duration   int    `json:"duration,omitempty"`
	Hold       bool   `json:"hold,omitempty"`
	Sound      string `json:"sound,omitempty"`
	Rtttl      string `json:"rtttl,omitempty"`
	LoopSound  bool   `json:"loopSound,omitempty"`
	Bar        []int  `json:"bar,omitempty"`
	Line       []any  `json:"line,omitempty"`
	AutoScale  bool   `json:"autoScale,omitempty"`
	Progress   int    `json:"progress,omitempty"`
	ProgressC  string `json:"progressC,omitempty"`
	ProgressBC string `json:"progressBC,omitempty"`
	Draw       []struct {
		Dc []any `json:"dc,omitempty"`
		Dr []any `json:"dr,omitempty"`
		Dt []any `json:"dt,omitempty"`
	} `json:"draw"`
	Stack          bool   `json:"stack"`
	Wakeup         bool   `json:"wakeup,omitempty"`
	NoScroll       bool   `json:"noScroll,omitempty"`
	Clients        string `json:"clients,omitempty"`
	ScrollSpeed    int    `json:"scrollSpeed,omitempty"`
	Effect         string `json:"effect,omitempty"`
	EffectSettings struct {
		Speed   int    `json:"speed,omitempty"`
		Palette string `json:"palette,omitempty"`
		Blend   bool   `json:"blend,omitempty"`
	} `json:"effectSettings,omitempty"`
}


func (c *Client) SendNotification(message string, color string, hold bool) {

	var duration int

	if hold {
		duration = 1000
	} else {
		duration = 10
	}

	fmt.Println("Sending message to", c.baseUrl)
	notifyRequest := NotifyRequest{
		Text:     message,
		TextCase: 1,
		Duration: duration,
		Stack:    false,
		Color:    color,
	}
	c.PostRequest("/api/notify", notifyRequest)
}

func (c *Client) DismissNotification() {
	fmt.Println("Dismissing notification")

	c.PostRequest("/api/notify/dismiss", nil)
}