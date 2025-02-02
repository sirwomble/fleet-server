module github.com/elastic/fleet-server/v7

go 1.15

require (
	github.com/aleksmaus/generate v0.0.0-20201213151810-c5bc68a6a42f
	github.com/dgraph-io/ristretto v0.0.3
	github.com/elastic/beats/v7 v7.11.1
	github.com/elastic/elastic-agent-client/v7 v7.0.0-20200709172729-d43b7ad5833a
	github.com/elastic/go-elasticsearch/v8 v8.0.0-20200728144331-527225d8e836
	github.com/elastic/go-ucfg v0.8.3
	github.com/gofrs/uuid v3.3.0+incompatible
	github.com/google/go-cmp v0.4.0
	github.com/hashicorp/go-cleanhttp v0.5.1
	github.com/hashicorp/golang-lru v0.5.2-0.20190520140433-59383c442f7d
	github.com/julienschmidt/httprouter v1.3.0
	github.com/pkg/errors v0.9.1
	github.com/rs/xid v1.2.1
	github.com/rs/zerolog v1.19.0
	github.com/spf13/cobra v0.0.5
	github.com/stretchr/testify v1.6.1
	go.uber.org/zap v1.14.0
	golang.org/x/net v0.0.0-20200822124328-c89045814202
	golang.org/x/sync v0.0.0-20200625203802-6e8e738ad208
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e
)

replace (
	github.com/Shopify/sarama => github.com/elastic/sarama v1.19.1-0.20200629123429-0e7b69039eec
	github.com/dop251/goja => github.com/andrewkroh/goja v0.0.0-20190128172624-dd2ac4456e20
	github.com/fsnotify/fsevents => github.com/elastic/fsevents v0.0.0-20181029231046-e1d381a4d270
)
