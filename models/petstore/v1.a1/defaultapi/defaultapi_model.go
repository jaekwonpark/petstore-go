/*
 * Generated file petstore/v1.a1/defaultapi/defaultapi_model.go.  Product version: 1.0.0-SNAPSHOT
 *
 * Part of the Petstore API project
 *
 * (c) 2020 Nutanix Inc.  All rights reserved
 *
 */
package defaultapi
import (
    import1 "github.com/jaekwon.park/petstore-go/models/common/v1.a1/response"
)

type Category struct {
        Id *int64 `json:"id,omitempty"`
        Name *string `json:"name,omitempty"`
}

type IntVal struct {
        IntVal *int32 `json:"int_val,omitempty"`
}

type Pet struct {
        Category *Category `json:"category,omitempty"`
        Id *int64 `json:"id,omitempty"`
        Name *string `json:"name"`
        PhotoFiles *[]string `json:"photo_files,omitempty"`
        Point *Point `json:"point,omitempty"`
        Status *string `json:"status,omitempty"`
        Tags *[]Tag `json:"tags,omitempty"`
}

type PetApiResponse struct {
        Data *OneOfPetApiResponseData `json:"data,omitempty"`
        Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

type Point struct {
        TimestampEpoch *int32 `json:"timestamp_epoch,omitempty"`
        Value *OneOfPetstoreV1A1DefaultapiPointValue `json:"value,omitempty"`
}

type StrVal struct {
        StrVal *string `json:"str_val,omitempty"`
}

type Tag struct {
        Id *int64 `json:"id,omitempty"`
        Name *string `json:"name,omitempty"`
}

type Url struct {
        Url *string `json:"url,omitempty"`
}

type UrlApiResponse struct {
        Data *OneOfPetstoreV1A1DefaultapiUrlApiResponseData `json:"data,omitempty"`
        Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

type OneOfPetstoreV1A1DefaultapiPointValue interface {
    OneOfPetstoreV1A1DefaultapiPointValue()
}

func (*StrVal) OneOfPetstoreV1A1DefaultapiPointValue() {}
func (*IntVal) OneOfPetstoreV1A1DefaultapiPointValue() {}
type OneOfPetstoreV1A1DefaultapiUrlApiResponseData interface {
    OneOfPetstoreV1A1DefaultapiUrlApiResponseData()
}

func (*Url) OneOfPetstoreV1A1DefaultapiUrlApiResponseData() {}


type OneOfPetApiResponseData interface {
    OneOfPetApiResponseData()
}
func (*Pet) OneOfPetApiResponseData() {}


func (*Pet) ApiResponseMetadata() {}

func (*Url) ApiResponseMetadata() {}

