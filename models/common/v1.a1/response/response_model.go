/*
 * Generated file common/v1.a1/response/response_model.go.  Product version: 1.0.0-RC
 *
 * Part of the Common API and Schema definitions
 *
 * (c) 2020 Nutanix Inc.  All rights reserved
 *
 */
package response
import (
    import2 "github.com/jaekwon.park/petstore-go/models/common/v1.a1/config"
)



type ApiResponseMetadata struct {
        Flags *[]import2.Flag `json:"flags,omitempty"`
        Links *[]import2.ApiLink `json:"links,omitempty"`
        Messages *[]import2.Message `json:"messages,omitempty"`
}


func (a *ApiResponseMetadata) OneOfPetApiResponseData() {}

