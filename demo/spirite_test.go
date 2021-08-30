package demo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

//定义结构体
type Sprite struct {
	Name string   `json:"name"`
	Md5  string   `json:"md5"` //todo
	Type string   `json:"type"`
	Tags []string `json:"tags"`
	Info []int    `json:"info"`
	Json JsonObj  `json:"json"`
}
type JsonObj struct {
	ObjName             string     `json:"objName"`
	Sound               []Sound    `json:"sound"`
	Costumes            []Costume  `json:"costumes"`
	CurrentCostumeIndex int        `json:"currentCostumeIndex"`
	ScratchX            int        `json:"scratchX"`
	ScratchY            int        `json:"scratchY"`
	Scale               int        `json:"scale"`
	Direction           int        `json:"direction"`
	RotationStyle       string     `json:"rotationStyle"`
	IsDraggable         bool       `json:"isDraggable"`
	Visible             bool       `json:"visible"`
	SpriteInfo          SpriteInfo `json:"spriteInfo"`
}
type Sound struct {
	SoundName string `json:"soundName"`

	SoundID     int    `json:"soundID"`
	Md5         string `json:"md5"` //todo
	SampleCount int    `json:"sampleCount"`
	Rate        int    `json:"rate"`
	Format      string `json:"format"`
}

type Costume struct {
	CostumeName      string `json:"costumeName"`
	BaseLayerID      int    `json:"baseLayerID"`
	BaseLayerMD5     string `json:"baseLayerMD5"` //todo
	BitmapResolution int    `json:"bitmapResolution"`
	RotationCenterX  int    `json:"rotationCenterX"`
	RotationCenterY  int    `json:"rotationCenterY"`
}

type SpriteInfo struct {
}

var assetUrl = "https://assets.scratch.mit.edu/internalapi/asset/%s/get/"

func TestSprite(t *testing.T) {
	//读取json文件
	//spriteByte, err := ioutil.ReadFile("/Users/quon/go/src/github.com/Quons/matchmaker/demo/sprites.json")
	spriteByte, err := ioutil.ReadFile("/Users/quon/go/src/github.com/Quons/matchmaker/demo/sounds.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	var spriteList []Sprite
	err = json.Unmarshal(spriteByte, &spriteList)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(spriteList)
	//遍历并下载
	for _, s := range spriteList {
		err = DownloadFile(s.Md5, t)
		if err != nil {
			fmt.Println(err)
		}
		//遍历sounds costumes
		for _, sound := range s.Json.Sound {
			err = DownloadFile(sound.Md5, t)
			if err != nil {
				fmt.Println(err)
			}

		}
		for _, costume := range s.Json.Costumes {
			err = DownloadFile(costume.BaseLayerMD5, t)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func DownloadFile(md5 string, t *testing.T) error {
	t.Logf("download:%s", md5)
	resp, err := http.Get(fmt.Sprintf(assetUrl, md5))
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	//err = ioutil.WriteFile("/Users/quon/go/src/github.com/Quons/matchmaker/demo/asset/"+md5, body, 766)
	err = ioutil.WriteFile("/Users/quon/go/src/github.com/Quons/matchmaker/demo/sounds/"+md5, body, 766)
	if err != nil {
		return err
	}
	return nil
}
