/*
 * Generated file common/v1.a1/config/config_model.go.  Product version: 1.0.0-RC
 *
 * Part of the Common API and Schema definitions
 *
 * (c) 2020 Nutanix Inc.  All rights reserved
 *
 */
package config

type AbstractModel struct {
        Reserved_ *string `json:"reserved_,omitempty"`
        ObjectType *string `json:"object_type,omitempty"`
}

type BaseEnum int

const(
    BASEENUM_UNKNOWN BaseEnum = 0
)

// returns the name of the enum given an ordinal number
func (e BaseEnum) name(index int) string {
    names := [...]string {
        "BASEENUM_UNKNOWN",
    }
    if index < 0 || index > len(names) {
       return "_UNKNOWN"
    }
    return names[index]
}
// returns the enum type given a string value
func (e BaseEnum) index(name string) BaseEnum {
   names := [...]string {
    "BASEENUM_UNKNOWN",
   }
   for idx := range names {
     if names[idx] == name {
        return BaseEnum(idx)
     }
   }

   return BASEENUM_UNKNOWN
}


type Flag struct {
        Name *string `json:"name,omitempty"`
        Value *bool `json:"value,omitempty"`
}

type KVStringPair struct {
        Name *string `json:"name,omitempty"`
        Value *string `json:"value,omitempty"`
}

type Message struct {
        Code *string `json:"code,omitempty"`
        Locale *string `json:"locale,omitempty"`
        Message *string `json:"message,omitempty"`
        Severity *MessageSeverity `json:"severity,omitempty"`
}

type MessageSeverity int

const(
    INFO MessageSeverity = 0
    WARNING MessageSeverity = 1
    ERROR MessageSeverity = 2
    MESSAGESEVERITY_UNKNOWN MessageSeverity = 3
)

// returns the name of the enum given an ordinal number
func (e MessageSeverity) name(index int) string {
    names := [...]string {
        "INFO",
        "WARNING",
        "ERROR",
        "MESSAGESEVERITY_UNKNOWN",
    }
    if index < 0 || index > len(names) {
       return "_UNKNOWN"
    }
    return names[index]
}
// returns the enum type given a string value
func (e MessageSeverity) index(name string) MessageSeverity {
   names := [...]string {
    "INFO",
    "WARNING",
    "ERROR",
    "MESSAGESEVERITY_UNKNOWN",
   }
   for idx := range names {
     if names[idx] == name {
        return MessageSeverity(idx)
     }
   }

   return MESSAGESEVERITY_UNKNOWN
}

type ApiLink struct {
  Href *string `json:"href,omitempty"`
  Rel *string `json:"rel,omitempty"`
}


type Messages []Message

type MetadataAbstractModel struct {
        Reserved_ *string `json:"reserved_,omitempty"`
        Links *[]ApiLink `json:"links,omitempty"`
        ObjectType *string `json:"object_type,omitempty"`
}

type Page struct {
        Limit_ *int32 `json:"limit_,omitempty"`
        Page_ *int32 `json:"page_,omitempty"`
}


