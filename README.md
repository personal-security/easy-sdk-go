# easy-sdk-go [![Build Status](https://travis-ci.com/personal-security/easy-sdk-go.svg?branch=main)](https://travis-ci.com/personal-security/easy-sdk-go)

This is a mini library that contains some functions that make our life easier. If you have something to add, make a pull request or write.

## Install

`go get -u "github.com/personal-security/easy-sdk-go"`

## Functions

Listing functions is dev.

* MattermostSendMessage - Send message with hook to mattermost service
* RunCMD - Run command in console and return resutl

## Import

```GO
import (
    easysdk "github.com/personal-security/easy-sdk-go"
)
```

## Example Rest Answer

```GO
package controllers

import (
    "net/http"
    "rest-api/models"

    "github.com/gorilla/mux"
    easysdk "github.com/personal-security/easy-sdk-go"
)

var StatusGetNow = func(w http.ResponseWriter, r *http.Request) {
    // CODE

    resp := &easysdk.RespondApi{}
    resp.Create(true, "Success")
    resp.Respond(w)
}
```

or

```GO
package controllers

import (
    "net/http"
    "rest-api/models"

    "github.com/gorilla/mux"
    easysdk "github.com/personal-security/easy-sdk-go"
)

var StatusGetNow = func(w http.ResponseWriter, r *http.Request) {
    // CODE

    easysdk.GenerateApiRespond(w,true,"Success",nil)
}
```

## Links

[pkg.go.dev](https://pkg.go.dev/github.com/personal-security/easy-sdk-go)
