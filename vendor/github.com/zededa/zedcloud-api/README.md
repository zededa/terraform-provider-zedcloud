# github.com/zededa/zedcloud-api

zedcloud-api provides the open source API for zedcontrol (zedcloud). Currently, it provides a
Go-Lang API. Support for API for Python and other languages can be added as needed.

API Swagger files are downloaded from zedcontrol. Go-Lang API is then generated using
*quay.io/goswagger/swagger generate client* command. *make* in the root directory
downloads the swagger files and generates the corresponding GO API - swagger_models and swagger_client.

1. *swagger_models*
    - This directory has the generated Go structures for all API structs.
2. swagger_client
    - This directory the corresponding generated Client wrapper.
3. client.go
    - This file provides a wrapper *Client* object to interact with zedcloud. The Client has
    CSREF support.
4. zedcloud_client
    - This has a sample Go-lang CLI client utility that uses the zedcloud-api. This is
      a tool mainly intended to serve as an example for using the API. Also used to
      manually test any changes to the API.
    - Do *make client* from zedcloud-api directory to build zedcloud_client
    - *zedcloud_client/zedcloud_client -h* for more info on using zedcloud_client.
