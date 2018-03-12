package server

import (
	"fmt"
	"net/http"
	"path"

	"github.com/dbenque/kcodec/pkg/api/kcodec"
	"github.com/dbenque/kcodec/pkg/api/kcodec/v1"
	"github.com/dbenque/kcodec/pkg/api/kcodec/v1ext"
	"github.com/dbenque/kcodec/pkg/api/kcodec/v2"
	restful "github.com/emicklei/go-restful"
	"github.com/golang/glog"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
)

const (
	serverName = "server"
)

//Server is the root data object
type Server struct {
	BindAddress  string
	Handler      http.Handler
	codec        runtime.Codec
	codecFactory serializer.CodecFactory
}

// MarshalToYaml marshals an object into yaml.
func (s *Server) MarshalToYaml(obj runtime.Object, gv schema.GroupVersion) ([]byte, error) {
	return s.MarshalToYamlForCodecs(obj, gv, s.codecFactory)
}

// MarshalToYamlForCodecs marshals an object into yaml using the specified codec
func (s *Server) MarshalToYamlForCodecs(obj runtime.Object, gv schema.GroupVersion, codecs serializer.CodecFactory) ([]byte, error) {
	mediaType := "application/yaml"
	info, ok := runtime.SerializerInfoForMediaType(codecs.SupportedMediaTypes(), mediaType)
	if !ok {
		return []byte{}, fmt.Errorf("unsupported media type %q", mediaType)
	}

	encoder := codecs.EncoderForVersion(info.Serializer, gv)
	return runtime.Encode(encoder, obj)
}

// MarshalToJson marshals an object into yaml.
func (s *Server) MarshalToJson(obj runtime.Object, gv schema.GroupVersion) ([]byte, error) {
	return s.MarshalToJsonForCodecs(obj, gv, s.codecFactory)
}

// MarshalToJsonForCodecs an object into yaml using the specified codec
func (s *Server) MarshalToJsonForCodecs(obj runtime.Object, gv schema.GroupVersion, codecs serializer.CodecFactory) ([]byte, error) {
	mediaType := "application/json"
	info, ok := runtime.SerializerInfoForMediaType(codecs.SupportedMediaTypes(), mediaType)
	if !ok {
		return []byte{}, fmt.Errorf("unsupported media type %q", mediaType)
	}

	encoder := codecs.EncoderForVersion(info.Serializer, gv)
	return runtime.Encode(encoder, obj)
}
func (s *Server) Init() {
	container := restful.NewContainer()
	s.Handler = container
	s.codecFactory = serializer.NewCodecFactory(kcodec.Scheme)
	//s.codec = serializer.NewCodecFactory(kcodec.Scheme).LegacyCodec(v1.SchemeGroupVersion, v1ext.SchemeGroupVersion, v2.SchemeGroupVersion, kcodec.SchemeGroupVersion)

	// for k, v := range kcodec.Scheme.AllKnownTypes() {
	// 	fmt.Printf("Group:%v , Version:%v , Kind:%v : %v\n", k.Group, k.Version, k.Kind, v)
	// }

	// Add container filter to enable CORS
	cors := restful.CrossOriginResourceSharing{
		ExposeHeaders:  []string{},
		AllowedHeaders: []string{"Content-Type"},
		CookiesAllowed: false,
		Container:      container}
	container.Filter(cors.Filter)

	// Add container filter to respond to OPTIONS
	container.Filter(container.OPTIONSFilter)

	wsV1 := new(restful.WebService).Path("/v1").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)
	wsV2 := new(restful.WebService).Path("/v2").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)

	allWebServices := []*restful.WebService{wsV1, wsV2}
	for _, ws := range allWebServices {
		container.Add(ws)
	}

	s.RegisterItem(allWebServices)
}

func (s *Server) Run() {
	server := &http.Server{Addr: s.BindAddress, Handler: s.Handler}

	glog.Fatal(server.ListenAndServe())
}

func init() {
	kcodec.AddToScheme(kcodec.Scheme)
	v1.AddToScheme(kcodec.Scheme)
	v1ext.AddToScheme(kcodec.Scheme)
	v2.AddToScheme(kcodec.Scheme)
}

func (s *Server) RegisterItem(webservices []*restful.WebService) {

	getItem := func(request *restful.Request, response *restful.Response) {
		name := request.PathParameter("name")
		val := &v1ext.Item{}
		val.Name = name
		val.Kind = "Item"
		val.APIVersion = "kcodec-ext/v1"
		val.Annotations = map[string]string{"specvalue": "5"}

		valInternal := &kcodec.Item{}
		if err := kcodec.Scheme.Convert(val, valInternal, nil); err != nil {
			glog.Errorf("Convertion to internal error: %s", err)
		}

		val2 := &v2.Item{}
		if err := kcodec.Scheme.Convert(valInternal, val2, nil); err != nil {
			glog.Errorf("Convertion to v2 error: %s", err)
		}

		b, err := s.MarshalToJson(val2, v2.SchemeGroupVersion)
		if err != nil {
			glog.Errorf("Marshalling error: %s", err)
		}
		_, err = response.Write(b)
		if err != nil {
			glog.Errorf("Encoding error: %s", err)
		}
	}

	for _, ws := range webservices {
		ws.Route(ws.GET(path.Join("item", "{name}")).To(getItem).
			Doc("Retrieve the resource that matches the given name.").
			Param(ws.PathParameter("name", "name of the object").DataType("string").Required(true)).
			Writes(v2.Item{}))
	}
}
