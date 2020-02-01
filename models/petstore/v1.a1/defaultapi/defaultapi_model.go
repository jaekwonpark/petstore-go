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
  "errors"
  "fmt"
  import1 "github.com/jaekwon.park/petstore-go/models/common/v1.a1/config"
        import2 "github.com/jaekwon.park/petstore-go/models/common/v1.a1/response"
)

type Category struct {
        Id *int64 `json:"id,omitempty"`
        Name *string `json:"name,omitempty"`
}

type IntVal struct {
        Val *int32 `json:"int_val,omitempty"`
}

type StrVal struct {
        Val *string `json:"str_val,omitempty"`
}

type BoolVal struct {
        Val *bool `json:"bool_val,omitempty"`
}

type OneofIntValStrValBoolVal struct {
  I IntVal
  S StrVal
  B BoolVal
}

func (o *OneofIntValStrValBoolVal) SetValue(v interface{}) error {
       switch v.(type) {
       case IntVal:
               o.I = v.(IntVal)
       case StrVal:
               o.S = v.(StrVal)
       case BoolVal:
               o.B = v.(BoolVal)
       default:
         return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
       }
       return nil
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
        Data OneOfPetApiResponseData `json:"data,omitempty"`
        Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

type Point struct {
        TimestampEpoch *int32 `json:"timestamp_epoch,omitempty"`
        Value *OneOfPetstoreV1A1DefaultapiPointValue `json:"value,omitempty"`
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
        Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
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


type OneOfPetApiResponseData struct {
  Messages *import1.Messages
  Pet *Pet
}

func (o *OneOfPetApiResponseData) SetValue (v interface {}) error {
  switch v.(type) {
  case Pet:
    if nil == o.Pet {
      o.Pet = new(Pet)
    }
    *o.Pet = v.(Pet)
  case import1.Messages:
    if nil == o.Messages {
      o.Messages = new(import1.Messages)
    }
    *o.Messages = v.(import1.Messages)
  default:
    return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

/*type OneOfPetApiResponseData interface {
  Messages()
}
func (*Pet) Messages() {}
 */

func (*Pet) ApiResponseMetadata() {}

func (*Url) ApiResponseMetadata() {}

