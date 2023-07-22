package repository

import "github.com/ishanshre/GoShortLinker/internals/types"

type DbRepository interface {
	InsertUrl(u *types.Url) (*types.Url, error)
	GetUrl(url string) (*types.Url, error)
}
