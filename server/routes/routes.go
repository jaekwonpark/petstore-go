/*
 * Generated file main.go.  Product version: 1.0.0-SNAPSHOT
 *
 * Part of the Petstore API project
 *
 * (c) 2020 Nutanix Inc.  All rights reserved
 *
 */

package routes

import (
  "fmt"
  "github.com/jaekwon.park/petstore-go/endpoints/pet"
  "github.com/jaekwon.park/petstore-go/endpoints/utils"
  "net/http"
)

var petApiImplWrapper = pet.PetApiImplWrapper{}
var utilsApiImplWrapper = utils.UtilsApiImplWrapper{}

type implAccessor interface {
  GetServiceImplementation() interface{}
  SetServiceImplementation(impl interface{})
}

type HandlerImplType int
const (
  PET_API_IMPL_WRAPPER HandlerImplType = 0
  UTILS_API_IMPL_WRAPPER HandlerImplType = 1
)

var handlerImplName = map[HandlerImplType] string {
  PET_API_IMPL_WRAPPER : "petApiImplWrapper",
  UTILS_API_IMPL_WRAPPER : "utilsApiImplWrapper",
}

var handlerImplMap = map[HandlerImplType] implAccessor {
  PET_API_IMPL_WRAPPER : &petApiImplWrapper,
  UTILS_API_IMPL_WRAPPER : &utilsApiImplWrapper,
}

type RouteKey struct {
  Path string
  Method string
}

func HandlerMap() map[RouteKey] func(http.ResponseWriter, *http.Request) {
  r := make(map[RouteKey]func(http.ResponseWriter, *http.Request))
  var route RouteKey
  route = RouteKey{ Path:"/petstore/v1.a1/defaultapi/pets/{petId}/downloadImage", Method:"Get" }
  r[route] = petApiImplWrapper.GetImageById
  route = RouteKey{ Path:"/petstore/v1.a1/defaultapi/pets/{petId}/uploadImage", Method:"Post" }
  r[route] = petApiImplWrapper.UploadFile
  route = RouteKey{ Path:"/petstore/v1.a1/defaultapi/pets/{petId}", Method:"Get" }
  r[route] = petApiImplWrapper.GetPetById
  route = RouteKey{ Path:"/petstore/v1.a1/defaultapi/pets", Method:"Post" }
  r[route] = petApiImplWrapper.AddPet
  route = RouteKey{ Path:"/petstore/v1.a1/utils/namespaces", Method:"Get" }
  r[route] = utilsApiImplWrapper.GetAllNamespaces
  route = RouteKey{ Path:"/petstore/v1.a1/utils/namespaces{id}", Method:"Get" }
  r[route] = utilsApiImplWrapper.GetNamespaceById


  return r
}

func ValidateReadinessToRun() {
  unimplementedInterfaces := make([] string, 0)

  for k := range handlerImplMap {
    if handlerImplMap[k].GetServiceImplementation() == nil {
      unimplementedInterfaces = append(unimplementedInterfaces, handlerImplName[k])
    }
  }
  if len(unimplementedInterfaces) > 0 {
    errMsg := fmt.Sprintf("The following interfaces have not been implemented:%v", unimplementedInterfaces)
    panic(errMsg)
  }
}

func SetServiceImplementation(implType HandlerImplType, impl interface{}) {
  handlerImplMap[implType].SetServiceImplementation(impl)
}