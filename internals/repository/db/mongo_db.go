package db

import (
	"context"
	"errors"

	"github.com/ishanshre/GoShortLinker/internals/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (m *mongoDbRepo) InsertUrl(u *types.Url) (*types.Url, error) {
	ctx, cancel := context.WithTimeout(m.Ctx, timeout)
	defer cancel()

	res, err := m.DB.GetUrlCollections().InsertOne(ctx, u)
	if err != nil {
		return nil, errors.New("error in inserting new url")
	}
	u.ID = res.InsertedID.(primitive.ObjectID)
	return u, nil
}

func (m *mongoDbRepo) GetUrl(url string) (*types.Url, error) {
	ctx, cancel := context.WithTimeout(m.Ctx, timeout)
	defer cancel()

	res := m.DB.GetUrlCollections().FindOne(ctx, bson.M{"long_url": url})
	getUrl := types.Url{}
	if err := res.Decode(&getUrl); err != nil {
		return nil, errors.New("error in fetching url")
	}
	return &getUrl, nil
}
