package models

type ShareList struct {
	TvSeries *TvSeries `yaml:"电视剧"`
	Movies   *Movies   `yaml:"电影"`
}

type TvSeries struct {
	XiYouJi      string `yaml:"西游记86版"`
	ShiZongZui   string `yaml:"十宗罪"`
	FaYiQingMing string `yaml:"法医秦明"`
}

type Movies struct {
	ManWei               string `yaml:"漫威合集"`
	XinHaiChengGongQiJun string `yaml:"新海诚&宫崎骏合集"`
	LinZhengYing         string `yaml:"林正英合集"`
}
