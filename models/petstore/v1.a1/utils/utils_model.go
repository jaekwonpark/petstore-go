/*
 * Generated file petstore/v1.a1/utils/utils_model.go.  Product version: 1.0.0-SNAPSHOT
 *
 * Part of the Petstore API project
 *
 * (c) 2020 Nutanix Inc.  All rights reserved
 *
 */
package utils
import (
    import1 "github.com/jaekwon.park/petstore-go/common/v1.a1/response"
)

type NamespaceApiResponse struct {
        Data *OneOfPetstoreV1A1UtilsNamespaceApiResponseData `json:"data,omitempty"`
        Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

type NamespaceListApiResponse struct {
        Data *OneOfPetstoreV1A1UtilsNamespaceListApiResponseData `json:"data,omitempty"`
        Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

//func (*utils.CommonUnversionedA1UtilsNamespaceList) import1.ApiResponseMetadata() {}
type OneOfPetstoreV1A1UtilsNamespaceApiResponseData interface {
    OneOfPetstoreV1A1UtilsNamespaceApiResponseData()
}

//func (*CommonUnversionedA1UtilsNamespace) OneOfPetstoreV1A1UtilsNamespaceApiResponseData() {}
type OneOfPetstoreV1A1UtilsNamespaceListApiResponseData interface {
    OneOfPetstoreV1A1UtilsNamespaceListApiResponseData()
}

//func (*CommonUnversionedA1UtilsNamespaceList) OneOfPetstoreV1A1UtilsNamespaceListApiResponseData() {}
type ApiResponseMetadata interface {
    ApiResponseMetadata()
}

//func (*CommonUnversionedA1UtilsNamespace) ApiResponseMetadata() {}

