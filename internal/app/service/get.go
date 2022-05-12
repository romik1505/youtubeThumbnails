package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/romik1505/youtubeThumbnails/internal/app/model"
	"github.com/romik1505/youtubeThumbnails/internal/app/store"
	desc "github.com/romik1505/youtubeThumbnails/pkg/api/thumbnails"
	"io/ioutil"
	"log"
	"regexp"
)

func (ts ThumbnailService) Get(ctx context.Context, request *desc.GetRequest) (*desc.GetResponse, error) {
	url, err := ts.getIDByUrl(ctx, request.GetUrl())
	if err != nil {
		return nil, err
	}

	th, err := ts.Storage.GetThumbnail(ctx, url)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}

		img, err := ts.Loader.LoadImg(ctx, url)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}

		th, err = ts.Storage.InsertThumbnail(ctx, model.Thumbnail{
			IDVideo:        store.NewNullString(url),
			ThumbnailImage: img,
		})
		if err != nil {
			return nil, err
		}
	}

	log.Println("Image successfully returned")
	return &desc.GetResponse{
		Image: th.ThumbnailImage,
	}, nil
}

func (f FileLoader) LoadImg(ctx context.Context, url string) ([]byte, error) {
	imgUrl := fmt.Sprintf("https://img.youtube.com/vi/%s/mqdefault.jpg", url)

	resp, err := f.Client.Get(imgUrl)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	img, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return img, nil
}

var (
	YoutubeVideoRegexp = regexp.MustCompile(`https://www\.youtube\.com/watch\?v=([\w-_]+)`)
)

func (ts ThumbnailService) getIDByUrl(ctx context.Context, url string) (string, error) {
	res := YoutubeVideoRegexp.FindStringSubmatch(url)
	if len(res) != 2 {
		return "", fmt.Errorf("invalid url format")
	}

	return res[1], nil
}
