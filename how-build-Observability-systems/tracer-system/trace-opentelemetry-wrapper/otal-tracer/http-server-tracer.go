package otal_tracer

import (
	"encoding/json"
	"expvar"
	"fmt"
	"io"
	"net/http"
	"path"
	"time"

	"github.com/jaegertracing/jaeger/examples/hotrod/pkg/httperr"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

type HttpServerTracer struct {
	tracer   trace.TracerProvider
}

func NewHttpServerTracer(tracerItem trace.TracerProvider) *HttpServerTracer {
	return &HttpServerTracer{
		tracer: tracerItem,
	}
}

func (s *HttpServerTracer) createServeMux() http.Handler  {
	mux := NewServeMux(true, s.tracer)
	p := path.Join("/", "")
	mux.Handle(path.Join(p, "/dispatch"), http.HandlerFunc(s.dispatch))
	mux.Handle(path.Join(p, "/debug/vars"), expvar.Handler()) // expvar

	return mux
}


type responseData struct {
	A int 
	B string
}

type RequestData struct {
	CustomerId int `json:"customer"`
}

func (s *HttpServerTracer) dispatch(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Printf("HTTP request received, method: %v, url: %v\n", r.Method, r.URL)
	_ = ctx

	// if err := r.ParseForm(); httperr.HandleError(w, err, http.StatusBadRequest) {
	// 	fmt.Println("bad request, err: ", err)
	// 	return
	// }
	 // 读取请求体
    body, err := io.ReadAll(r.Body)
    if err!= nil {
        http.Error(w, "读取请求体失败", http.StatusInternalServerError)
        return
    }
    defer r.Body.Close()
	 // 解析JSON数据
    var requestData RequestData
    err = json.Unmarshal(body, &requestData)
    if err!= nil {
        http.Error(w, "解析JSON失败", http.StatusBadRequest)
        return
    }


	// customer := r.Form.Get("customer")
	// if customer == "" {
	// 	http.Error(w, "Missing required 'customer' parameter", http.StatusBadRequest)
	// 	return
	// }
	// customerID, err := strconv.Atoi(customer)
	// if err != nil {
	// 	http.Error(w, "Parameter 'customer' is not an integer", http.StatusBadRequest)
	// 	return
	// }
	fmt.Println("customerId: ", requestData.CustomerId)

	// TODO distinguish between user errors (such as invalid customer ID) and server failures
	// response, err := s.bestETA.Get(ctx, customerID)
	// if httperr.HandleError(w, err, http.StatusInternalServerError) {
	// 	fmt.Println("request failed", err)
	// 	return
	// }
	// s.writeResponse(response, w, r)

	var rData = &responseData{
		A: 100,
		B: "hello world",
	}
	s.writeResponse(rData, w, r)
}

func (s *HttpServerTracer) writeResponse(response any, w http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(response)
	if httperr.HandleError(w, err, http.StatusInternalServerError) {
		fmt.Println("cannot marshal response, err: ", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func (h *HttpServerTracer) Run() error {
	mux := h.createServeMux()
	fmt.Printf("Starting, address: http://%v\n", "0.0.0.0:5656")
	server := &http.Server{
		Addr:              "0.0.0.0:5656",
		Handler:           mux,
		ReadHeaderTimeout: 3 * time.Second,
	}
	return server.ListenAndServe()
}



// NewServeMux creates a new TracedServeMux.
func NewServeMux(copyBaggage bool, tracer trace.TracerProvider) *TracedServeMux {
	return &TracedServeMux{
		mux:         http.NewServeMux(),
		copyBaggage: copyBaggage,
		tracer:      tracer,
	}
}

// TracedServeMux is a wrapper around http.ServeMux that instruments handlers for tracing.
type TracedServeMux struct {
	mux         *http.ServeMux
	copyBaggage bool
	tracer      trace.TracerProvider
}

// Handle implements http.ServeMux#Handle, which is used to register new handler.
func (tm *TracedServeMux) Handle(pattern string, handler http.Handler) {
	fmt.Println("registering traced handler, endpoint: ", pattern)

	middleware := otelhttp.NewHandler(
		otelhttp.WithRouteTag(pattern, traceResponseHandler(handler)),
		pattern,
		otelhttp.WithTracerProvider(tm.tracer))

	tm.mux.Handle(pattern, middleware)
}

// ServeHTTP implements http.ServeMux#ServeHTTP.
func (tm *TracedServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tm.mux.ServeHTTP(w, r)
}

// Returns a handler that generates a traceresponse header.
// https://github.com/w3c/trace-context/blob/main/spec/21-http_response_header_format.md
func traceResponseHandler(handler http.Handler) http.Handler {
	// We use the standard TraceContext propagator, since the formats are identical.
	// But the propagator uses "traceparent" header name, so we inject it into a map
	// `carrier` and then use the result to set the "tracereponse" header.
	var prop propagation.TraceContext
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		carrier := make(map[string]string)
		prop.Inject(r.Context(), propagation.MapCarrier(carrier))
		w.Header().Add("traceresponse", carrier["traceparent"])
		handler.ServeHTTP(w, r)
	})
}
