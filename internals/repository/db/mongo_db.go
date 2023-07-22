package db

import (
	"context"
	"errors"

	"github.com/ishanshre/GoShortLinker/internals/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

func (m *mongoDbRepo) GetUrl(urlCode string) (*types.Url, error) {
	ctx, cancel := context.WithTimeout(m.Ctx, timeout)
	defer cancel()

	res := m.DB.GetUrlCollections().FindOne(ctx, bson.M{"url_code": urlCode})
	getUrl := types.Url{}
	if err := res.Decode(&getUrl); err != nil {
		return nil, errors.New("error in fetching url")
	}
	return &getUrl, nil
}

func (m *mongoDbRepo) UrlCodeExists(urlCode string) (bool, error) {
	ctx, cancel := context.WithTimeout(m.Ctx, timeout)
	defer cancel()

	filter := bson.M{"url_code": urlCode}
	res := m.DB.GetUrlCollections().FindOne(ctx, filter)
	if err := res.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (m *mongoDbRepo) DeleteUrlCode(urlCode string) error {
	ctx, cancel := context.WithTimeout(m.Ctx, timeout)
	defer cancel()

	filter := bson.M{"url_code": urlCode}
	res, err := m.DB.GetUrlCollections().DeleteOne(ctx, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.New("url does not exists")
		}
		return errors.New("error in deleting url")
	}
	if res.DeletedCount == 0 {
		return errors.New("error in deleting url or url does not exists")
	}
	return nil
}
