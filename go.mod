module AMPCortex

go 1.16

// Override since git.apache.org is down.  The docs say to fetch from github.
replace git.apache.org/thrift.git => github.com/apache/thrift v0.0.0-20180902110319-2566ecd5d999

replace k8s.io/client-go => k8s.io/client-go v0.20.4

replace k8s.io/api => k8s.io/api v0.20.4

// Use fork of gocql that has gokit logs and Prometheus metrics.
replace github.com/gocql/gocql => github.com/grafana/gocql v0.0.0-20200605141915-ba5dc39ece85

// Using a 3rd-party branch for custom dialer - see https://github.com/bradfitz/gomemcache/pull/86
replace github.com/bradfitz/gomemcache => github.com/themihai/gomemcache v0.0.0-20180902122335-24332e2d58ab

require (
	github.com/cortexproject/cortex v1.9.1-0.20210525181333-f16afd84e611
	github.com/pkg/errors v0.9.1
	github.com/weaveworks/common v0.0.0-20210419092856-009d1eebd624
	gopkg.in/yaml.v2 v2.4.0
)
