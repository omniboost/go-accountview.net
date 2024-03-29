module github.com/omniboost/go-accountview.net

go 1.20

require (
	github.com/fatih/structtag v1.2.0
	github.com/gorilla/schema v0.0.0-20171211162101-9fa3b6af65dc
	github.com/pkg/errors v0.9.1
	golang.org/x/oauth2 v0.0.0-20210628180205-a41e5a781914
	gopkg.in/guregu/null.v3 v3.5.0
)

require (
	github.com/golang/protobuf v1.4.2 // indirect
	golang.org/x/net v0.0.0-20200822124328-c89045814202 // indirect
	google.golang.org/appengine v1.6.6 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
)

replace github.com/gorilla/schema => github.com/omniboost/schema v1.1.1-0.20191030093734-a170fe1a7240
