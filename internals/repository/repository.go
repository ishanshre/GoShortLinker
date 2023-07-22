package repository

import "github.com/ishanshre/GoShortLinker/internals/types"

type DbRepository interface {
	InsertUrl(u *types.Url) (*types.Url, error)
	GetUrl(urlCode string) (*types.Url, error)
	UrlCodeExists(urlCode string) (bool, error)
}
