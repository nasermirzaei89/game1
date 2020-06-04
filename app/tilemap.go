package app

import (
	"encoding/xml"
	"github.com/pkg/errors"
	"github.com/veandco/go-sdl2/sdl"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type TileMap interface {
}

type tileMap struct {
	texture    *sdl.Texture
	mapWidth   int
	mapHeight  int
	tileWidth  int
	tileHeight int
	xOffset    int
	yOffset    int
	xSpacing   int
	ySpacing   int
	columns    int
	data       []int
}

type xmlMap struct {
	XMLName          xml.Name `xml:"map"`
	Version          string   `xml:"version,attr"`
	TiledVersion     string   `xml:"tiledversion,attr"`
	Orientation      string   `xml:"orientation,attr"`
	RenderOrder      string   `xml:"renderorder,attr"`
	CompressionLevel string   `xml:"compressionlevel,attr"`
	Width            int      `xml:"width,attr"`
	Height           int      `xml:"height,attr"`
	TileWidth        int      `xml:"tilewidth,attr"`
	TileHeight       int      `xml:"tileheight,attr"`
	Infinite         string   `xml:"infinite,attr"`
	NextLayerID      string   `xml:"nextlayerid,attr"`
	NextObjectID     string   `xml:"nextobjectid,attr"`
	TileSet          struct {
		FirstGID string `xml:"firstgid,attr"`
		Source   string `xml:"source,attr"`
	} `xml:"tileset"`
	Layer struct {
		ID     string `xml:"id,attr"`
		Name   string `xml:"name,attr"`
		Width  string `xml:"width,attr"`
		Height string `xml:"height,attr"`
		Data   struct {
			Encoding string `xml:"encoding,attr"`
			Text     string `xml:",chardata"`
		} `xml:"data"`
	} `xml:"layer"`
}

type xmlTileSet struct {
	XMLName      xml.Name `xml:"tileset"`
	Version      string   `xml:"version,attr"`
	TiledVersion string   `xml:"tiledversion,attr"`
	Name         string   `xml:"name,attr"`
	TileWidth    int      `xml:"tilewidth,attr"`
	TileHeight   int      `xml:"tileheight,attr"`
	Spacing      int      `xml:"spacing,attr"`
	Margin       int      `xml:"margin,attr"`
	TileCount    int      `xml:"tilecount,attr"`
	Columns      int      `xml:"columns,attr"`
	Image        struct {
		Source           string `xml:"source,attr"`
		TransparentColor string `xml:"trans,attr"`
		Width            int    `xml:"width,attr"`
		Height           int    `xml:"height,attr"`
	} `xml:"image"`
}

func NewTileMap(filename string, game Game) (*tileMap, error) {
	f1, err := os.Open(filename)
	if err != nil {
		return nil, errors.Wrap(err, "error on open map file")
	}
	defer func() { _ = f1.Close() }()

	var map1 xmlMap
	err = xml.NewDecoder(f1).Decode(&map1)
	if err != nil {
		return nil, errors.Wrap(err, "error on decode map file xml")
	}

	f2, err := os.Open(filepath.Join(filepath.Dir(filename), map1.TileSet.Source))
	if err != nil {
		return nil, errors.Wrap(err, "error on open tile set file")
	}
	defer func() { _ = f2.Close() }()

	var tileSet1 xmlTileSet
	err = xml.NewDecoder(f2).Decode(&tileSet1)
	if err != nil {
		return nil, errors.Wrap(err, "error on decode tile set file xml")
	}

	texture, err := game.LoadTexture(filepath.Join(filepath.Dir(filename), tileSet1.Image.Source))
	if err != nil {
		return nil, errors.Wrap(err, "error on load texture")
	}

	res := tileMap{
		texture:    texture,
		mapWidth:   map1.Width,
		mapHeight:  map1.Height,
		tileWidth:  tileSet1.TileWidth,
		tileHeight: tileSet1.TileHeight,
		xOffset:    tileSet1.Margin,
		yOffset:    tileSet1.Margin,
		xSpacing:   tileSet1.Spacing,
		ySpacing:   tileSet1.Spacing,
		columns:    tileSet1.Columns,
		data:       make([]int, map1.Width*map1.Height),
	}

	data := strings.Split(map1.Layer.Data.Text, ",")
	for i := range data {
		res.data[i], err = strconv.Atoi(strings.TrimSpace(data[i]))
		if err != nil {
			return nil, errors.Wrapf(err, "error on convert `%s` to int", data[i])
		}
	}

	return &res, nil
}

func (t *tileMap) Update(deltaTime float32) error {
	return nil
}

func (t *tileMap) Render(renderer *sdl.Renderer) error {
	for i := 0; i < t.mapWidth; i++ {
		for j := 0; j < t.mapHeight; j++ {
			index := t.data[j*t.mapWidth+i] - 1
			src := sdl.Rect{
				X: int32(t.xOffset + (index%t.columns)*t.tileWidth + (index%t.columns)*t.xSpacing),
				Y: int32(t.yOffset + (index/t.columns)*t.tileHeight + (index/t.columns)*t.ySpacing),
				W: int32(t.tileWidth),
				H: int32(t.tileHeight),
			}
			dst := sdl.Rect{
				X: int32(i * t.tileWidth),
				Y: int32(j * t.tileHeight),
				W: int32(t.tileWidth),
				H: int32(t.tileHeight),
			}

			err := renderer.Copy(t.texture, &src, &dst)
			if err != nil {
				return errors.Wrap(err, "error on copy texture to renderer")
			}
		}
	}
	return nil
}
