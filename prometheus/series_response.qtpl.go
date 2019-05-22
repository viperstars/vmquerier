// Code generated by qtc from "series_response.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line app/vmselect/prometheus/series_response.qtpl:1
package prometheus

//line app/vmselect/prometheus/series_response.qtpl:1
import (
	"github.com/valyala/quicktemplate"
)

// SeriesResponse generates response for /api/v1/series.See https://prometheus.io/docs/prometheus/latest/querying/api/#finding-series-by-label-matchers

//line app/vmselect/prometheus/series_response.qtpl:8
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line app/vmselect/prometheus/series_response.qtpl:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line app/vmselect/prometheus/series_response.qtpl:8
func StreamSeriesResponse(qw422016 *qt422016.Writer, resultsCh <-chan *quicktemplate.ByteBuffer) {
//line app/vmselect/prometheus/series_response.qtpl:8
	qw422016.N().S(`{"status":"success","data":[`)
//line app/vmselect/prometheus/series_response.qtpl:12
	bb, ok := <-resultsCh

//line app/vmselect/prometheus/series_response.qtpl:13
	if ok {
//line app/vmselect/prometheus/series_response.qtpl:14
		qw422016.N().Z(bb.B)
//line app/vmselect/prometheus/series_response.qtpl:15
		quicktemplate.ReleaseByteBuffer(bb)

//line app/vmselect/prometheus/series_response.qtpl:16
		for bb := range resultsCh {
//line app/vmselect/prometheus/series_response.qtpl:16
			qw422016.N().S(`,`)
//line app/vmselect/prometheus/series_response.qtpl:17
			qw422016.N().Z(bb.B)
//line app/vmselect/prometheus/series_response.qtpl:18
			quicktemplate.ReleaseByteBuffer(bb)

//line app/vmselect/prometheus/series_response.qtpl:19
		}
//line app/vmselect/prometheus/series_response.qtpl:20
	}
//line app/vmselect/prometheus/series_response.qtpl:20
	qw422016.N().S(`]}`)
//line app/vmselect/prometheus/series_response.qtpl:23
}

//line app/vmselect/prometheus/series_response.qtpl:23
func WriteSeriesResponse(qq422016 qtio422016.Writer, resultsCh <-chan *quicktemplate.ByteBuffer) {
//line app/vmselect/prometheus/series_response.qtpl:23
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmselect/prometheus/series_response.qtpl:23
	StreamSeriesResponse(qw422016, resultsCh)
//line app/vmselect/prometheus/series_response.qtpl:23
	qt422016.ReleaseWriter(qw422016)
//line app/vmselect/prometheus/series_response.qtpl:23
}

//line app/vmselect/prometheus/series_response.qtpl:23
func SeriesResponse(resultsCh <-chan *quicktemplate.ByteBuffer) string {
//line app/vmselect/prometheus/series_response.qtpl:23
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmselect/prometheus/series_response.qtpl:23
	WriteSeriesResponse(qb422016, resultsCh)
//line app/vmselect/prometheus/series_response.qtpl:23
	qs422016 := string(qb422016.B)
//line app/vmselect/prometheus/series_response.qtpl:23
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmselect/prometheus/series_response.qtpl:23
	return qs422016
//line app/vmselect/prometheus/series_response.qtpl:23
}
