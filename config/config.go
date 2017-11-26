package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"bitbucket.org/djr2/tldr/color"
	"github.com/mitchellh/go-homedir"
)

var Config Variables

type Variables struct {
	PagesURI         string `json:"pages_uri"`
	ZipURI           string `json:"zip_uri"`
	BannerColor1     int    `json:"banner_color_1"`
	BannerColor2     int    `json:"banner_color_2"`
	TLDRColor        int    `json:"tldr_color"`
	HeaderColor      int    `json:"header_color"`
	HeaderDecorColor int    `json:"header_decor_color"`
	PlatformColor    int    `json:"platform_color"`
	DescriptionColor int    `json:"description_color"`
	ExampleColor     int    `json:"example_color"`
	HypenColor       int    `json:"hypen_color"`
	SyntaxColor      int    `json:"syntax_color"`
	VariableColor    int    `json:"variable_color"`
}

func Load() {
	h, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}

	f := h + "/" + ".tldr/config.json"
	_, err = os.Stat(f)
	if err != nil {
		create(f)
	}

	file, err := os.Open(f)
	if err != nil {
		log.Println(err)
		return
	}

	b, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
		return
	}

	if json.Unmarshal(b, &Config) != nil {
		log.Println(err)
	}
}

func create(f string) {
	vars := Variables{
		PagesURI:         "",
		ZipURI:           "",
		BannerColor1:     color.Cyan,
		BannerColor2:     color.Blue,
		TLDRColor:        color.White,
		HeaderColor:      color.Blue,
		HeaderDecorColor: color.White,
		PlatformColor:    color.DarkGray,
		DescriptionColor: color.Normal,
		ExampleColor:     color.Cyan,
		HypenColor:       color.Normal,
		SyntaxColor:      color.Red,
		VariableColor:    color.Normal,
	}

	file, err := os.Create(f)
	if err != nil {
		log.Fatal(err)
	}

	j, err := json.MarshalIndent(vars, "", "")
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.Write(j)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
}
