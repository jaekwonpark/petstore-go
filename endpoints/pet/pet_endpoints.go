/*
 * Generated file pet_endpoints.go.  Product version: 1.0.0-SNAPSHOT
 *
 * Part of the Petstore API project
 *
 * (c) 2020 Nutanix Inc.  All rights reserved
 *
 */

package pet

import (
  "net/http"
)

type PetApi interface {

    // Endpoint path: /petstore/v1.a1/defaultapi/pets
    // Description: 
    // HttpMethod: Post
    // Parameters
    //    Name: body, Type: petstore.v1.a1.defaultapi.Pet, Description: Create a new pet in the store
    //
    //
    // RETURN TYPE:  import1.PetApiResponse - The returned response must be of this type
    // Implementation class must make sure that this is the case.  The prefix of the import is the prefix
    // associated with the import statement.  You will need to replace it with your own prefix.
    AddPet(w http.ResponseWriter, r *http.Request)

    // Endpoint path: /petstore/v1.a1/defaultapi/pets/{petId}/downloadImage
    // Description: Downloads a pet image file
    // HttpMethod: Get
    // Parameters
    //    Name: petId, Type: int64, Description: ID of pet to download the image with
    //
    //
    // RETURN TYPE:   - The returned response must be of this type
    // Implementation class must make sure that this is the case.  The prefix of the import is the prefix
    // associated with the import statement.  You will need to replace it with your own prefix.
    GetImageById(w http.ResponseWriter, r *http.Request)

    // Endpoint path: /petstore/v1.a1/defaultapi/pets/{petId}
    // Description: Returns a single pet
    // HttpMethod: Get
    // Parameters
    //    Name: petId, Type: int64, Description: ID of pet to return
    //
    //
    // RETURN TYPE:   - The returned response must be of this type
    // Implementation class must make sure that this is the case.  The prefix of the import is the prefix
    // associated with the import statement.  You will need to replace it with your own prefix.
    GetPetById(w http.ResponseWriter, r *http.Request)

    // Endpoint path: /petstore/v1.a1/defaultapi/pets/{petId}/uploadImage
    // Description: 
    // HttpMethod: Post
    // Parameters
    //    Name: body, Type: Object, Description: Uploads an image
    //    Name: petId, Type: int64, Description: ID of pet to update
    //    Name: additionalMetadata, Type: string, Description: Additional Metadata
    //
    //
    // RETURN TYPE:   - The returned response must be of this type
    // Implementation class must make sure that this is the case.  The prefix of the import is the prefix
    // associated with the import statement.  You will need to replace it with your own prefix.
    UploadFile(w http.ResponseWriter, r *http.Request)


}

type PetApiImplWrapper struct {
  // create a binding to the actual interface implementation in the Init method
  // where this is used.
  svcImpl PetApi
}

// CREATE IMPLEMENTATION BINDINGS TO ALL THE METHODS IN THE INTERFACE

func (serviceWrapper PetApiImplWrapper) AddPet(w http.ResponseWriter, r *http.Request) {
  serviceWrapper.svcImpl.AddPet(w, r)
}


func (serviceWrapper PetApiImplWrapper) GetImageById(w http.ResponseWriter, r *http.Request) {
  serviceWrapper.svcImpl.GetImageById(w, r)
}


func (serviceWrapper PetApiImplWrapper) GetPetById(w http.ResponseWriter, r *http.Request) {
  serviceWrapper.svcImpl.GetPetById(w, r)
}


func (serviceWrapper PetApiImplWrapper) UploadFile(w http.ResponseWriter, r *http.Request) {
  serviceWrapper.svcImpl.UploadFile(w, r)
}


func (serviceWrapper *PetApiImplWrapper) SetServiceImplementation(apiImplementation interface{}) {
  val := apiImplementation.(PetApi)
  serviceWrapper.svcImpl = val
}

func (serviceWrapper PetApiImplWrapper) GetServiceImplementation() interface{} {
   return serviceWrapper.svcImpl
}


