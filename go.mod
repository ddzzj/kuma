module github.com/Kong/kuma

go 1.12

require (
	github.com/Kong/kuma/api v0.0.0-00010101000000-000000000000
	github.com/Kong/kuma/pkg/plugins/resources/k8s/native v0.0.0-00010101000000-000000000000
	github.com/Masterminds/goutils v1.1.0 // indirect
	github.com/Masterminds/semver v1.4.2 // indirect
	github.com/Masterminds/sprig v2.20.0+incompatible
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/emicklei/go-restful v2.9.6+incompatible
	github.com/envoyproxy/go-control-plane v0.11.1-0.20230524094728-9239064ad72f
	github.com/envoyproxy/protoc-gen-validate v0.10.1
	github.com/ghodss/yaml v1.0.0
	github.com/go-logr/logr v0.1.0
	github.com/go-logr/zapr v0.1.0
	github.com/golang/protobuf v1.5.3
	github.com/hoisie/mustache v0.0.0-20160804235033-6375acf62c69
	github.com/huandu/xstrings v1.2.0 // indirect
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/lib/pq v1.1.1
	github.com/mxk/go-flowrate v0.0.0-20140419014527-cca7078d478f // indirect
	github.com/onsi/ginkgo v1.10.3
	github.com/onsi/gomega v1.7.1
	github.com/pkg/errors v0.9.1
	github.com/shurcooL/httpfs v0.0.0-20190707220628-8d4bc4ba7749
	github.com/spf13/cobra v0.0.5
	github.com/spiffe/go-spiffe v0.0.0-20190820222348-6adcf1eecbcc
	github.com/spiffe/spire v0.0.0-20190905203639-e85640baca1d
	go.uber.org/multierr v1.1.0
	go.uber.org/zap v1.9.1
	google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1
	google.golang.org/grpc v1.56.3
	gopkg.in/yaml.v2 v2.2.5
	k8s.io/api v0.0.0-20190409021203-6e4e0e4f393b
	k8s.io/apimachinery v0.0.0-20190404173353-6a84e37a896d
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	k8s.io/helm v2.14.3+incompatible
	sigs.k8s.io/controller-runtime v0.2.2
	sigs.k8s.io/testing_frameworks v0.1.1
)

replace (
	github.com/Kong/kuma/api => ./api
	github.com/Kong/kuma/pkg/plugins/resources/k8s/native => ./pkg/plugins/resources/k8s/native
)
