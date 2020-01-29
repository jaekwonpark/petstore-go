package pet

import (
  "context"
  "encoding/json"
  "fmt"
  "github.com/gorilla/mux"
  "github.com/jaekwon.park/petstore-go/models/common/v1.a1/config"
  "github.com/jaekwon.park/petstore-go/models/common/v1.a1/response"
  "github.com/jaekwon.park/petstore-go/models/petstore/v1.a1/defaultapi"
  "go.mongodb.org/mongo-driver/bson"
  "github.com/google/uuid"
  "log"
  "net/http"
  "strconv"

  //"go.mongodb.org/mongo-driver/bson/primitive"
  "go.mongodb.org/mongo-driver/mongo"
)

type Trainer struct {
  Name string
  Age  int
  City string
}

type Service struct {
  Mc *mongo.Collection
}

func (s Service) GetImageById(w http.ResponseWriter, r *http.Request) {
  defaultResponse(w, http.StatusOK, []byte("{\"Said\":\"Hello\"}"))
}

func (s Service) GetPetById(w http.ResponseWriter, r *http.Request) {

  vars := mux.Vars(r)
  sId := vars["petId"]
  id, err := strconv.ParseInt(sId, 10, 64)
  res := &defaultapi.PetApiResponse{}
  var iMsgs defaultapi.OneOfPetApiResponseData
  if err != nil {
    msg := fmt.Sprintf("Can not convert %s to int64", sId)
    iMsgs = addMessage(msg)
    log.Printf(msg)
    res.Data = &iMsgs
  } else {
    pet := &defaultapi.Pet{}
    //to get by mongodb document id, use _id, for example, filter := bson.M{"_id": id}
    filter := bson.M{"id": id}
    err := s.Mc.FindOne(context.TODO(), filter).Decode(&pet)
    if err != nil {
      iMsgs = addMessage(err.Error())
      log.Printf(err.Error())
      res.Data = &iMsgs
    } else {
      var iPet defaultapi.OneOfPetApiResponseData
      iPet = pet
      res.Data = &iPet
    }
  }

  metadata := &response.ApiResponseMetadata{}
  res.Metadata = metadata
  bRes, err := json.Marshal(res)
  if err != nil {
    log.Fatal(err)
  }
  defaultResponse(w, http.StatusOK, bRes)
}

func (s Service) UploadFile(w http.ResponseWriter, r *http.Request) {
  defaultResponse(w, http.StatusOK, []byte("{\"Said\":\"Hello\"}"))
}

func (s Service) AddPet(w http.ResponseWriter, r *http.Request) {
  oneDoc := &defaultapi.Pet{}
  err := json.NewDecoder(r.Body).Decode(&oneDoc)
  res := &defaultapi.PetApiResponse{}
  var iMsgs defaultapi.OneOfPetApiResponseData
  if err != nil {
    iMsgs = addMessage(err.Error())
    log.Printf(err.Error())
    res.Data = &iMsgs
  } else {
    var iPet defaultapi.OneOfPetApiResponseData
    _, err := create(oneDoc, s.Mc)
    if err != nil {
      iMsgs = addMessage(err.Error())
      log.Printf(err.Error())
      res.Data = &iMsgs
    } else {
      iPet = oneDoc
      res.Data = &iPet
    }
  }

  metadata := &response.ApiResponseMetadata{}
  res.Metadata = metadata
  bRes, err := json.Marshal(res)
  if err != nil {
    log.Fatal(err)
  }
  defaultResponse(w, http.StatusOK, bRes)
}

func create(p *defaultapi.Pet, mc *mongo.Collection) (*mongo.InsertOneResult, error) {
  uuid := uuid.New()
  val := int64(uuid.ID())
  p.Id = &val
  return mc.InsertOne(context.TODO(), p)
}

func defaultResponse(w http.ResponseWriter, code int, resp []byte) {
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(code)
  w.Write(resp)
}

func addMessage(msg string) *config.Messages {
  errMsgs := &config.Messages{}
  errMsg :=  &config.Message{}
  //msg := fmt.Sprintf("Can not convert %s to int64", sId)
  errMsg.Message = &msg
  errMsgs.MessageList = &[]config.Message{}
  *errMsgs.MessageList =  append(*errMsgs.MessageList, *errMsg)

  return errMsgs
}