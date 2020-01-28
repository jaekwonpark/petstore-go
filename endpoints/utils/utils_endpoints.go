/*
 * Generated file utils_endpoints.go.  Product version: 1.0.0-SNAPSHOT
 *
 * Part of the Petstore API project
 *
 * (c) 2020 Nutanix Inc.  All rights reserved
 *
 */

package utils

import (
  "net/http"
)

type UtilsApi interface {

    // Endpoint path: /petstore/v1.a1/utils/namespaces
    // Description: Returns a list of all namespaces in the API platform
    // HttpMethod: Get
    //
    // RETURN TYPE:  import2.NamespaceListApiResponse - The returned response must be of this type
    // Implementation class must make sure that this is the case.  The prefix of the import is the prefix
    // associated with the import statement.  You will need to replace it with your own prefix.
    GetAllNamespaces(w http.ResponseWriter, r *http.Request)

    // Endpoint path: /petstore/v1.a1/utils/namespaces/{id}
    // Description: Returns details for a single namespace identified by {id}
    // HttpMethod: Get
    // Parameters
    //    Name: id, Type: string, Description: The uuid of the namespaces
    //
    //
    // RETURN TYPE:  import2.NamespaceApiResponse - The returned response must be of this type
    // Implementation class must make sure that this is the case.  The prefix of the import is the prefix
    // associated with the import statement.  You will need to replace it with your own prefix.
    GetNamespaceById(w http.ResponseWriter, r *http.Request)


}

type UtilsApiImplWrapper struct {
  // create a binding to the actual interface implementation in the Init method
  // where this is used.
  svcImpl UtilsApi
}

// CREATE IMPLEMENTATION BINDINGS TO ALL THE METHODS IN THE INTERFACE

func (serviceWrapper UtilsApiImplWrapper) GetAllNamespaces(w http.ResponseWriter, r *http.Request) {
  serviceWrapper.svcImpl.GetAllNamespaces(w, r)
}


func (serviceWrapper UtilsApiImplWrapper) GetNamespaceById(w http.ResponseWriter, r *http.Request) {
  serviceWrapper.svcImpl.GetNamespaceById(w, r)
}


func (serviceWrapper *UtilsApiImplWrapper) SetServiceImplementation(apiImplementation interface{}) {
  val := apiImplementation.(UtilsApi)
  serviceWrapper.svcImpl = val
}

func (serviceWrapper UtilsApiImplWrapper) GetServiceImplementation() interface{} {
   return serviceWrapper.svcImpl
}


