/*
 * Generated file common/unversioned.a1/utils/utils_model.go.  Product version: 1.0.0-RC
 *
 * Part of the Common API and Schema definitions
 *
 * (c) 2020 Nutanix Inc.  All rights reserved
 *
 */
package utils
import (
    import1 "github.com/jaekwon.park/petstore-go/common/v1.a1/response"
)

type Namespace struct {
        Reserved_ *string `json:"reserved_,omitempty"`
        Endpoints *[]string `json:"endpoints,omitempty"`
        ExtId *string `json:"ext_id,omitempty"`
        Links *[]import1.ApiLink `json:"links,omitempty"`
        ModelDependencies *[]string `json:"model_dependencies,omitempty"`
        Models *[]string `json:"models,omitempty"`
        Name *string `json:"name,omitempty"`
        ObjectType *string `json:"object_type,omitempty"`
        Version *string `json:"version,omitempty"`
}

type NamespaceList struct {
        Data *Namespace `json:"data,omitempty"`
}


