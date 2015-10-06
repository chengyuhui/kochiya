package main

type config struct {
	Output     output
	CustomArgs string
	Videos     []video
}

type output struct {
	Path   string
	Height int
}

type video struct {
	VideoPath     string
	ImagePath     string
	StartFrame    int // 开始帧数
	Length        int // 持续帧数
	OverlayLength int // 覆盖层持续帧数
	FadeInLength  int // 淡出效果持续帧数
	FadeOutLength int // 淡出效果持续帧数
}
