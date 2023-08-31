package apple

import (
	"githup.com/apibe/yifu/internal/warehouse/apple/argument"
	"githup.com/apibe/yifu/internal/warehouse/apple/cache"
	"githup.com/apibe/yifu/internal/warehouse/apple/formats"
	"githup.com/apibe/yifu/internal/warehouse/apple/roads"
)

type Apple struct {
	ID               string                 `json:"id,required"`
	Bucket           string                 `json:"bucket,required"`
	Name             string                 `json:"name,required"`
	Description      string                 `json:"description,omitempty"`
	Method           string                 `json:"method,required"`
	Url              string                 `json:"url,required"`
	ContentType      string                 `json:"Content-Type,required"`
	Payload          string                 `json:"payload,omitempty"`
	Argument         argument.Arguments     `json:"argument,omitempty"`
	RequestFunction  roads.RequestFunction  `json:"requestFunction,omitempty"`
	ResponseFunction roads.ResponseFunction `json:"responseFunction,omitempty"`
	Format           formats.Formats        `json:"format,omitempty"`
	Cache            cache.Config           `json:"cache,omitempty"`
	Claim            interface{}            `json:"claim"`
	Status           uint8                  `json:"status,omitempty"`
}
