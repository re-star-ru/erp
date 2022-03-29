package usecase

import (
	"backend/pkg/item"
	"backend/pkg/item/repo"
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/rs/zerolog/log"
)

type IItemUsecase interface {
	UploadPricelists(limit int) error
}

type Renderer interface {
	Render(map[string]item.Item) error
}

type ItemUsecase struct {
	renders []Renderer
	repo1c  repo.IClient1c
}

const pol = `
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Principal": {
                "AWS": [
                    "*"
                ]
            },
            "Action": [
                "s3:GetBucketLocation",
                "s3:ListBucket",
                "s3:ListBucketMultipartUploads"
            ],
            "Resource": [
                "arn:aws:s3:::pricelists"
            ]
        },
        {
            "Effect": "Allow",
            "Principal": {
                "AWS": [
                    "*"
                ]
            },
            "Action": [
                "s3:GetObject"
            ],
            "Resource": [
                "arn:aws:s3:::pricelists/*"
            ]
        }
    ]
}

`

func NewItemUsecase(repo1c repo.IClient1c, m *minio.Client) *ItemUsecase {
	bucket := "pricelists"

	err := m.MakeBucket(context.Background(), bucket, minio.MakeBucketOptions{})
	if err != nil {
		exists, errBucketExists := m.BucketExists(context.Background(), bucket)
		if errBucketExists == nil && exists {
			log.Info().Msgf("we already own %s", bucket)
		} else {
			log.Fatal().Err(err).Send()
		}
	} else {
		log.Info().Msgf("suscessfully created %s", bucket)
	}

	if err := m.SetBucketPolicy(context.Background(), bucket, pol); err != nil {
		log.Fatal().Err(err).Msg("set policy error")
	}

	return &ItemUsecase{
		renders: []Renderer{NewDromRender(bucket, m)},
		repo1c:  repo1c,
	}
}

func (iu *ItemUsecase) UploadPricelists(limit int) error {
	products, err := iu.repo1c.Products(0, limit)
	if err != nil {
		return fmt.Errorf("cant get producst from 1c %w", err)
	}
	for _, v := range iu.renders {
		if err := v.Render(products); err != nil {
			return err
		}
	}

	return nil
}