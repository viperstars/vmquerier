package promql

import (
	"testing"
)

func TestParseMetricSelectorSuccess(t *testing.T) {
	f := func(s string) {
		t.Helper()
		tfs, err := ParseMetricSelector(s)
		if err != nil {
			t.Fatalf("unexpected error when parsing %q: %s", s, err)
		}
		if tfs == nil {
			t.Fatalf("expecting non-nil tfs when parsing %q", s)
		}
	}
	f("foo")
	f(":foo")
	f("  :fo:bar.baz")
	f(`a{}`)
	f(`{foo="bar"}`)
	f(`{:f:oo=~"bar.+"}`)
	f(`foo {bar != "baz"}`)
	f(` foo { bar !~ "^ddd(x+)$", a="ss", __name__="sffd"}  `)
	f(`(foo)`)
}

func TestParseMetricSelectorError(t *testing.T) {
	f := func(s string) {
		t.Helper()
		tfs, err := ParseMetricSelector(s)
		if err == nil {
			t.Fatalf("expecting non-nil error when parsing %q", s)
		}
		if tfs != nil {
			t.Fatalf("expecting nil tfs when parsing %q", s)
		}
	}
	f("")
	f(`{}`)
	f(`foo bar`)
	f(`foo+bar`)
	f(`sum(bar)`)
	f(`x{y}`)
	f(`x{y+z}`)
	f(`foo[5m]`)
	f(`foo offset 5m`)
}

func TestParsePromQLSuccess(t *testing.T) {
	another := func(s string, sExpected string) {
		t.Helper()

		e, err := parsePromQL(s)
		if err != nil {
			t.Fatalf("unexpected error when parsing %q: %s", s, err)
		}
		res := e.AppendString(nil)
		if string(res) != sExpected {
			t.Fatalf("unexpected string constructed;\ngot\n%q\nwant\n%q", res, sExpected)
		}
	}
	same := func(s string) {
		t.Helper()
		another(s, s)
	}

	// metricExpr
	same(`{}`)
	same(`{}[5m]`)
	same(`{}[5m:]`)
	same(`{}[:]`)
	another(`{}[: ]`, `{}[:]`)
	same(`{}[:3s]`)
	another(`{}[: 3s ]`, `{}[:3s]`)
	same(`{}[5m:3s]`)
	another(`{}[ 5m : 3s ]`, `{}[5m:3s]`)
	same(`{} offset 5m`)
	same(`{}[5m] offset 10y`)
	same(`{}[5.3m:3.4s] offset 10y`)
	same(`{}[:3.4s] offset 10y`)
	same(`{Foo="bAR"}`)
	same(`{foo="bar"}`)
	same(`{foo="bar"}[5m]`)
	same(`{foo="bar"}[5m:]`)
	same(`{foo="bar"}[5m:3s]`)
	same(`{foo="bar"} offset 10y`)
	same(`{foo="bar"}[5m] offset 10y`)
	same(`{foo="bar"}[5m:3s] offset 10y`)
	another(`{foo="bar"}[5m] oFFSEt 10y`, `{foo="bar"}[5m] offset 10y`)
	same("METRIC")
	same("metric")
	same("m_e:tri44:_c123")
	another("-metric", "0 - metric")
	same(`metric offset 10h`)
	same("metric[5m]")
	same("metric[5m:3s]")
	same("metric[5m] offset 10h")
	same("metric[5m:3s] offset 10h")
	same("metric[5i:3i] offset 10i")
	same(`metric{foo="bar"}`)
	same(`metric{foo="bar"} offset 10h`)
	same(`metric{foo!="bar"}[2d]`)
	same(`metric{foo="bar"}[2d] offset 10h`)
	same(`metric{foo="bar", b="sdfsdf"}[2d:3h] offset 10h`)
	another(`  metric  {  foo  = "bar"  }  [  2d ]   offset   10h  `, `metric{foo="bar"}[2d] offset 10h`)
	// metric name matching keywords
	same("rate")
	same("RATE")
	same("by")
	same("BY")
	same("bool")
	same("BOOL")
	same("unless")
	same("UNLESS")
	same("Ignoring")
	same("with")
	same("WITH")
	same("With")
	same("alias")
	same(`alias{foo="bar"}`)
	same(`aLIas{alias="aa"}`)
	another(`al\ias`, `alias`)
	// identifiers with with escape chars
	same(`foo\ bar`)
	same(`foo\-bar\{{baz\+bar="aa"}`)
	another(`\x2E\x2ef\oo{b\xEF\ar="aa"}`, `\x2e.foo{b\xefar="aa"}`)
	// Duplicate filters
	same(`foo{__name__="bar"}`)
	same(`foo{a="b", a="c", __name__="aaa", b="d"}`)
	// Metric filters ending with comma
	another(`m{foo="bar",}`, `m{foo="bar"}`)
	// String concat in tag value
	another(`m{foo="bar" + "baz"}`, `m{foo="barbaz"}`)

	// Valid regexp
	same(`foo{bar=~"x"}`)
	same(`foo{bar=~"^x"}`)
	same(`foo{bar=~"^x$"}`)
	same(`foo{bar=~"^(a[bc]|d)$"}`)
	same(`foo{bar!~"x"}`)
	same(`foo{bar!~"^x"}`)
	same(`foo{bar!~"^x$"}`)
	same(`foo{bar!~"^(a[bc]|d)$"}`)

	// stringExpr
	same(`""`)
	same(`"\n\t\r 12:{}[]()44"`)
	another(`''`, `""`)
	another("``", `""`)
	another("   `foo\"b'ar`  ", "\"foo\\\"b'ar\"")
	another(`  'foo\'bar"BAZ'  `, `"foo'bar\"BAZ"`)
	// string concat
	another(`"foo"+'bar'`, `"foobar"`)

	// numberExpr
	same(`1`)
	same(`1.23`)
	same(`0.23`)
	same(`1.2e+45`)
	same(`1.2e-45`)
	same(`-1`)
	same(`-1.23`)
	same(`-0.23`)
	same(`-1.2e+45`)
	same(`-1.2e-45`)
	same(`-1.2e-45`)
	another(`12.5E34`, `1.25e+35`)
	another(`-.2`, `-0.2`)
	another(`-.2E-2`, `-0.002`)
	same(`NaN`)
	another(`nan`, `NaN`)
	another(`NAN`, `NaN`)
	another(`nAN`, `NaN`)
	another(`Inf`, `+Inf`)
	another(`INF`, `+Inf`)
	another(`inf`, `+Inf`)
	another(`+Inf`, `+Inf`)
	another(`-Inf`, `-Inf`)
	another(`-inF`, `-Inf`)

	// binaryOpExpr
	another(`nan == nan`, `NaN`)
	another(`nan ==bool nan`, `1`)
	another(`nan !=bool nan`, `0`)
	another(`nan !=bool 2`, `1`)
	another(`2 !=bool nan`, `1`)
	another(`nan >bool nan`, `0`)
	another(`nan <bool nan`, `0`)
	another(`1 ==bool nan`, `0`)
	another(`NaN !=bool 1`, `1`)
	another(`inf >=bool 2`, `1`)
	another(`-1 >bool -inf`, `1`)
	another(`-1 <bool -inf`, `0`)
	another(`nan + 2 *3 * inf`, `NaN`)
	another(`INF - Inf`, `NaN`)
	another(`Inf + inf`, `+Inf`)
	another(`1/0`, `+Inf`)
	another(`0/0`, `NaN`)
	another(`-m`, `0 - m`)
	same(`m + ignoring () n[5m]`)
	another(`M + IGNORING () N[5m]`, `M + ignoring () N[5m]`)
	same(`m + on (foo) n[5m]`)
	another(`m + ON (Foo) n[5m]`, `m + on (Foo) n[5m]`)
	same(`m + ignoring (a, b) n[5m]`)
	another(`1 or 2`, `1`)
	another(`1 and 2`, `1`)
	another(`1 unless 2`, `NaN`)
	another(`1 default 2`, `1`)
	another(`1 default NaN`, `1`)
	another(`NaN default 2`, `2`)
	another(`1 > 2`, `NaN`)
	another(`1 > bool 2`, `0`)
	another(`3 >= 2`, `3`)
	another(`3 <= bool 2`, `0`)
	another(`1 + -2 - 3`, `-4`)
	another(`1 / 0 + 2`, `+Inf`)
	another(`2 + -1 / 0`, `-Inf`)
	another(`-1 ^ 0.5`, `NaN`)
	another(`512.5 - (1 + 3) * (2 ^ 2) ^ 3`, `256.5`)
	another(`1 == bool 1 != bool 24 < bool 4 > bool -1`, `1`)
	another(`1 == bOOl 1 != BOOL 24 < Bool 4 > booL -1`, `1`)
	another(`m1+on(foo)group_left m2`, `m1 + on (foo) group_left () m2`)
	another(`M1+ON(FOO)GROUP_left M2`, `M1 + on (FOO) group_left () M2`)
	same(`m1 + on (foo) group_right () m2`)
	same(`m1 + on (foo, bar) group_right (x, y) m2`)
	another(`m1 + on (foo, bar,) group_right (x, y,) m2`, `m1 + on (foo, bar) group_right (x, y) m2`)
	same(`m1 == bool on (foo, bar) group_right (x, y) m2`)
	another(`5 - 1 + 3 * 2 ^ 2 ^ 3 - 2  OR Metric {Bar= "Baz", aaa!="bb",cc=~"dd" ,zz !~"ff" } `,
		`770 or Metric{Bar="Baz", aaa!="bb", cc=~"dd", zz!~"ff"}`)
	same(`"foo" + bar()`)
	same(`"foo" + bar{x="y"}`)
	same(`("foo"[3s] + bar{x="y"})[5m:3s] offset 10s`)
	same(`("foo"[3s] + bar{x="y"})[5i:3i] offset 10i`)
	same(`bar + "foo" offset 3s`)
	same(`bar + "foo" offset 3i`)
	another(`1+2 if 2>3`, `NaN`)
	another(`1+4 if 2<3`, `5`)
	another(`2+6 default 3 if 2>3`, `8`)
	another(`2+6 if 2>3 default NaN`, `NaN`)
	another(`42 if 3>2 if 2+2<5`, `42`)
	another(`42 if 3>2 if 2+2>=5`, `NaN`)
	another(`1+2 ifnot 2>3`, `3`)
	another(`1+4 ifnot 2<3`, `NaN`)
	another(`2+6 default 3 ifnot 2>3`, `8`)
	another(`2+6 ifnot 2>3 default NaN`, `8`)
	another(`42 if 3>2 ifnot 2+2<5`, `NaN`)
	another(`42 if 3>2 ifnot 2+2>=5`, `42`)

	// parensExpr
	another(`(-foo + ((bar) / (baz))) + ((23))`, `((0 - foo) + (bar / baz)) + 23`)
	another(`(FOO + ((Bar) / (baZ))) + ((23))`, `(FOO + (Bar / baZ)) + 23`)
	same(`(foo, bar)`)
	another(`((foo, bar),(baz))`, `((foo, bar), baz)`)
	same(`(foo, (bar, baz), ((x, y), (z, y), xx))`)
	another(`1+(foo, bar,)`, `1 + (foo, bar)`)
	another(`((foo(bar,baz)), (1+(2)+(3,4)+()))`, `(foo(bar, baz), (3 + (3, 4)) + ())`)
	same(`()`)

	// funcExpr
	same(`f()`)
	another(`f(x,)`, `f(x)`)
	another(`-f()-Ff()`, `(0 - f()) - Ff()`)
	same(`F()`)
	another(`+F()`, `F()`)
	another(`++F()`, `F()`)
	another(`--F()`, `0 - (0 - F())`)
	same(`f(http_server_request)`)
	same(`f(http_server_request)[4s:5m] offset 10m`)
	same(`f(http_server_request)[4i:5i] offset 10i`)
	same(`F(HttpServerRequest)`)
	same(`f(job, foo)`)
	same(`F(Job, Foo)`)
	another(` FOO (bar) + f  (  m  (  ),ff(1 + (  2.5)) ,M[5m ]  , "ff"  )`, `FOO(bar) + f(m(), ff(3.5), M[5m], "ff")`)
	// funcName matching keywords
	same(`by(2)`)
	same(`BY(2)`)
	same(`or(2)`)
	same(`OR(2)`)
	same(`bool(2)`)
	same(`BOOL(2)`)
	same(`rate(rate(m))`)
	same(`rate(rate(m[5m]))`)
	same(`rate(rate(m[5m])[1h:])`)
	same(`rate(rate(m[5m])[1h:3s])`)
	// funcName with escape chars
	same(`foo\(ba\-r()`)

	// aggrFuncExpr
	same(`sum(http_server_request) by ()`)
	same(`sum(http_server_request) by (job)`)
	same(`sum(http_server_request) without (job, foo)`)
	another(`sum(x,y,) without (a,b,)`, `sum(x, y) without (a, b)`)
	another(`sum by () (xx)`, `sum(xx) by ()`)
	another(`sum by (s) (xx)[5s]`, `(sum(xx) by (s))[5s]`)
	another(`SUM BY (ZZ, aa) (XX)`, `sum(XX) by (ZZ, aa)`)
	another(`sum without (a, b) (xx,2+2)`, `sum(xx, 4) without (a, b)`)
	another(`Sum WIthout (a, B) (XX,2+2)`, `sum(XX, 4) without (a, B)`)
	same(`sum(a) or sum(b)`)
	same(`sum(a) by () or sum(b) without (x, y)`)
	same(`sum(a) + sum(b)`)
	same(`sum(x) * (1 + sum(a))`)

	// All the above
	another(`Sum(Ff(M) * M{X=""}[5m] Offset 7m - 123, 35) BY (X, y) * F2("Test")`,
		`sum((Ff(M) * M{X=""}[5m] offset 7m) - 123, 35) by (X, y) * F2("Test")`)
	another(`# comment
		Sum(Ff(M) * M{X=""}[5m] Offset 7m - 123, 35) BY (X, y) # yet another comment
		* F2("Test")`,
		`sum((Ff(M) * M{X=""}[5m] offset 7m) - 123, 35) by (X, y) * F2("Test")`)

	// withExpr
	another(`with () x`, `x`)
	another(`with (x=1,) x`, `1`)
	another(`with (x = m offset 5h) x + x`, `m offset 5h + m offset 5h`)
	another(`with (x = m offset 5i) x + x`, `m offset 5i + m offset 5i`)
	another(`with (foo = bar{x="x"}) 1`, `1`)
	another(`with (foo = bar{x="x"}) "x"`, `"x"`)
	another(`with (f="x") f`, `"x"`)
	another(`with (foo = bar{x="x"}) x{x="y"}`, `x{x="y"}`)
	another(`with (foo = bar{x="x"}) 1+1`, `2`)
	another(`with (foo = bar{x="x"}) f()`, `f()`)
	another(`with (foo = bar{x="x"}) sum(x)`, `sum(x)`)
	another(`with (foo = bar{x="x"}) baz{foo="bar"}`, `baz{foo="bar"}`)
	another(`with (foo = bar) baz`, `baz`)
	another(`with (foo = bar) foo + foo{a="b"}`, `bar + bar{a="b"}`)
	another(`with (foo = bar, bar=baz + f()) test`, `test`)
	another(`with (ct={job="test"}) a{ct} + ct() + f({ct="x"})`, `(a{job="test"} + {job="test"}) + f({ct="x"})`)
	another(`with (ct={job="test", i="bar"}) ct + {ct, x="d"} + foo{ct, ct} + ctx(1)`,
		`(({job="test", i="bar"} + {job="test", i="bar", x="d"}) + foo{job="test", i="bar"}) + ctx(1)`)
	another(`with (foo = bar) {__name__=~"foo"}`, `{__name__=~"foo"}`)
	another(`with (foo = bar) foo{__name__="foo"}`, `bar`)
	another(`with (foo = bar) {__name__="foo", x="y"}`, `bar{x="y"}`)
	another(`with (foo(bar) = {__name__!="bar"}) foo(x)`, `{__name__!="bar"}`)
	another(`with (foo(bar) = bar{__name__="bar"}) foo(x)`, `x`)
	another(`with (foo\-bar(baz) = baz + baz) foo\-bar((x,y))`, `(x, y) + (x, y)`)
	another(`with (foo\-bar(baz) = baz + baz) foo\-bar(x*y)`, `(x * y) + (x * y)`)
	another(`with (foo\-bar(baz) = baz + baz) foo\-bar(x\*y)`, `x\*y + x\*y`)
	another(`with (foo\-bar(b\ az) = b\ az + b\ az) foo\-bar(x\*y)`, `x\*y + x\*y`)
	// override ttf to something new.
	another(`with (ttf = a) ttf + b`, `a + b`)
	// override ttf to ru
	another(`with (ttf = ru(m, n)) ttf`, `(clamp_min(n - clamp_min(m, 0), 0) / clamp_min(n, 0)) * 100`)

	// Verify withExpr recursion and forward reference
	another(`with (x = x+y, y = x+x) y ^ 2`, `((x + y) + (x + y)) ^ 2`)
	another(`with (f1(x)=f2(x), f2(x)=f1(x)^2) f1(foobar)`, `f2(foobar)`)
	another(`with (f1(x)=f2(x), f2(x)=f1(x)^2) f2(foobar)`, `f2(foobar) ^ 2`)

	// Verify withExpr funcs
	another(`with (x() = y+1) x`, `y + 1`)
	another(`with (x(foo) = foo+1) x(a)`, `a + 1`)
	another(`with (x(a, b) = a + b) x(foo, bar)`, `foo + bar`)
	another(`with (x(a, b) = a + b) x(foo, x(1, 2))`, `foo + 3`)
	another(`with (x(a) = sum(a) by (b)) x(xx) / x(y)`, `sum(xx) by (b) / sum(y) by (b)`)
	another(`with (f(a,f,x)=ff(x,f,a)) f(f(x,y,z),1,2)`, `ff(2, 1, ff(z, y, x))`)
	another(`with (f(x)=1+f(x)) f(foo{bar="baz"})`, `1 + f(foo{bar="baz"})`)
	another(`with (a=foo, y=bar, f(a)= a+a+y) f(x)`, `(x + x) + bar`)
	another(`with (f(a, b) = m{a, b}) f({a="x", b="y"}, {c="d"})`, `m{a="x", b="y", c="d"}`)
	another(`with (xx={a="x"}, f(a, b) = m{a, b}) f({xx, b="y"}, {c="d"})`, `m{a="x", b="y", c="d"}`)
	another(`with (x() = {b="c"}) foo{x}`, `foo{b="c"}`)
	another(`with (f(x)=x{foo="bar"} offset 5m) f(m offset 10m)`, `(m{foo="bar"} offset 10m) offset 5m`)
	another(`with (f(x)=x{foo="bar",bas="a"}[5m]) f(m[10m] offset 3s)`, `(m{foo="bar", bas="a"}[10m] offset 3s)[5m]`)
	another(`with (f(x)=x{foo="bar"}[5m] offset 10m) f(m{x="y"})`, `m{x="y", foo="bar"}[5m] offset 10m`)
	another(`with (f(x)=x{foo="bar"}[5m] offset 10m) f({x="y", foo="bar", foo="bar"})`, `{x="y", foo="bar"}[5m] offset 10m`)
	another(`with (f(m, x)=m{x}[5m] offset 10m) f(foo, {})`, `foo[5m] offset 10m`)
	another(`with (f(m, x)=m{x, bar="baz"}[5m] offset 10m) f(foo, {})`, `foo{bar="baz"}[5m] offset 10m`)
	another(`with (f(x)=x[5m] offset 3s) f(foo[3m]+bar)`, `(foo[3m] + bar)[5m] offset 3s`)
	another(`with (f(x)=x[5m:3s] oFFsEt 1.5m) f(sum(s) by (a,b))`, `(sum(s) by (a, b))[5m:3s] offset 1.5m`)
	another(`with (x="a", y=x) y+"bc"`, `"abc"`)
	another(`with (x="a", y="b"+x) "we"+y+"z"+f()`, `"webaz" + f()`)
	another(`with (f(x) = m{foo=x+"y", bar="y"+x, baz=x} + x) f("qwe")`, `m{foo="qwey", bar="yqwe", baz="qwe"} + "qwe"`)
	another(`with (f(a)=a) f`, `f`)
	another(`with (f\q(a)=a) f\q`, `fq`)

	// Verify withExpr for aggr func modifiers
	another(`with (f(x) = x, y = sum(m) by (f)) y`, `sum(m) by (f)`)
	another(`with (f(x) = sum(m) by (x)) f(foo)`, `sum(m) by (foo)`)
	another(`with (f(x) = sum(m) by (x)) f((foo, bar, foo))`, `sum(m) by (foo, bar)`)
	another(`with (f(x) = sum(m) without (x,y)) f((a, b))`, `sum(m) without (a, b, y)`)
	another(`with (f(x) = sum(m) without (y,x)) f((a, y))`, `sum(m) without (y, a)`)
	another(`with (f(x,y) = a + on (x,y) group_left (y,bar) b) f(foo,())`, `a + on (foo) group_left (bar) b`)
	another(`with (f(x,y) = a + on (x,y) group_left (y,bar) b) f((foo),())`, `a + on (foo) group_left (bar) b`)
	another(`with (f(x,y) = a + on (x,y) group_left (y,bar) b) f((foo,xx),())`, `a + on (foo, xx) group_left (bar) b`)

	// Verify nested with exprs
	another(`with (f(x) = (with(x=y) x) + x) f(z)`, `y + z`)
	another(`with (x=foo) f(a, with (y=x) y)`, `f(a, foo)`)
	another(`with (x=foo) a * x + (with (y=x) y) / y`, `(a * foo) + (foo / y)`)
	another(`with (x = with (y = foo) y + x) x/x`, `(foo + x) / (foo + x)`)
	another(`with (
		x = {foo="bar"},
		q = m{x, y="1"},
		f(x) =
			with (
				z(y) = x + y * q
			)
			z(foo) / f(x)
	)
	f(a)`, `(a + (foo * m{foo="bar", y="1"})) / f(a)`)

	// complex withExpr
	another(`WITH (
		treshold = (0.9),
		commonFilters = {job="cacher", instance=~"1.2.3.4"},
		hits = rate(cache{type="hit", commonFilters}[5m]),
		miss = rate(cache{type="miss", commonFilters}[5m]),
		sumByInstance(arg) = sum(arg) by (instance),
		hitRatio = sumByInstance(hits) / sumByInstance(hits + miss)
	)
	hitRatio < treshold`,
		`(sum(rate(cache{type="hit", job="cacher", instance=~"1.2.3.4"}[5m])) by (instance) / sum(rate(cache{type="hit", job="cacher", instance=~"1.2.3.4"}[5m]) + rate(cache{type="miss", job="cacher", instance=~"1.2.3.4"}[5m])) by (instance)) < 0.9`)
	another(`WITH (
		x2(x) = x^2,
		f(x, y) = x2(x) + x*y + x2(y)
	)
	f(a, 3)
	`, `((a ^ 2) + (a * 3)) + 9`)
	another(`WITH (
		x2(x) = x^2,
		f(x, y) = x2(x) + x*y + x2(y)
	)
	f(2, 3)
	`, `19`)
	another(`WITH (
		commonFilters = {instance="foo"},
		timeToFuckup(currv, maxv) = (maxv - currv) / rate(currv)
	)
	timeToFuckup(diskUsage{commonFilters}, maxDiskSize{commonFilters})`,
		`(maxDiskSize{instance="foo"} - diskUsage{instance="foo"}) / rate(diskUsage{instance="foo"})`)
	another(`WITH (
	       commonFilters = {job="foo", instance="bar"},
	       sumRate(m, cf) = sum(rate(m{cf})) by (job, instance),
	       hitRate(hits, misses) = sumRate(hits, commonFilters) / (sumRate(hits, commonFilters) + sumRate(misses, commonFilters))
	   )
	   hitRate(cacheHits, cacheMisses)`,
		`sum(rate(cacheHits{job="foo", instance="bar"})) by (job, instance) / (sum(rate(cacheHits{job="foo", instance="bar"})) by (job, instance) + sum(rate(cacheMisses{job="foo", instance="bar"})) by (job, instance))`)
	another(`with(y=123,z=5) union(with(y=3,f(x)=x*y) f(2) + f(3), with(x=5,y=2) x*y*z)`, `union(15, 50)`)
}

func TestParsePromQLError(t *testing.T) {
	f := func(s string) {
		t.Helper()

		e, err := parsePromQL(s)
		if err == nil {
			t.Fatalf("expecting non-nil error when parsing %q", s)
		}
		if e != nil {
			t.Fatalf("expecting nil expr when parsing %q", s)
		}
	}

	// an empty string
	f("")
	f("  \t\b\r\n  ")

	// invalid metricExpr
	f(`{__name__="ff"} offset 55`)
	f(`{__name__="ff"} offset -5m`)
	f(`foo[55]`)
	f(`m[-5m]`)
	f(`{`)
	f(`foo{`)
	f(`foo{bar`)
	f(`foo{bar=`)
	f(`foo{bar="baz"`)
	f(`foo{bar="baz",  `)
	f(`foo{123="23"}`)
	f(`foo{foo}`)
	f(`foo{,}`)
	f(`foo{,foo="bar"}`)
	f(`foo{foo=}`)
	f(`foo{foo="ba}`)
	f(`foo{"foo"="bar"}`)
	f(`foo{$`)
	f(`foo{a $`)
	f(`foo{a="b",$`)
	f(`foo{a="b"}$`)
	f(`[`)
	f(`[]`)
	f(`f[5m]$`)
	f(`[5m]`)
	f(`[5m] offset 4h`)
	f(`m[5m] offset $`)
	f(`m[5m] offset 5h $`)
	f(`m[]`)
	f(`m[-5m]`)
	f(`m[5m:`)
	f(`m[5m:-`)
	f(`m[5m:-1`)
	f(`m[5m:-1]`)
	f(`m[:`)
	f(`m[:-`)
	f(`m[:1]`)
	f(`m[:-1m]`)
	f(`m[5]`)
	f(`m[[5m]]`)
	f(`m[foo]`)
	f(`m["ff"]`)
	f(`m[10m`)
	f(`m[123`)
	f(`m["ff`)
	f(`m[(f`)
	f(`fd}`)
	f(`]`)
	f(`m $`)
	f(`m{,}`)
	f(`m{x=y}`)
	f(`m{x=y/5}`)
	f(`m{x=y+5}`)

	// Invalid regexp
	f(`foo{bar=~"x["}`)
	f(`foo{bar=~"x("}`)
	f(`foo{bar=~"x)"}`)
	f(`foo{bar!~"x["}`)
	f(`foo{bar!~"x("}`)
	f(`foo{bar!~"x)"}`)

	// invalid stringExpr
	f(`'`)
	f(`"`)
	f("`")
	f(`"foo`)
	f(`'foo`)
	f("`foo")
	f(`"foo\"bar`)
	f(`'foo\'bar`)
	f("`foo\\`bar")
	f(`"" $`)
	f(`"foo" +`)
	f(`n{"foo" + m`)

	// invalid numberExpr
	f(`12.`)
	f(`1.2e`)
	f(`23e-`)
	f(`23E+`)
	f(`.`)
	f(`-12.`)
	f(`-1.2e`)
	f(`-23e-`)
	f(`-23E+`)
	f(`-.`)
	f(`-1$$`)
	f(`-$$`)
	f(`+$$`)
	f(`23 $$`)

	// invalid binaryOpExpr
	f(`+`)
	f(`1 +`)
	f(`1 + 2.`)
	f(`3 unless`)
	f(`23 + on (foo)`)
	f(`m + on (,) m`)
	f(`3 * ignoring`)
	f(`m * on (`)
	f(`m * on (foo`)
	f(`m * on (foo,`)
	f(`m * on (foo,)`)
	f(`m * on (,foo)`)
	f(`m * on (,)`)
	f(`m == bool (bar) baz`)
	f(`m == bool () baz`)
	f(`m * by (baz) n`)
	f(`m + bool group_left m2`)
	f(`m + on () group_left (`)
	f(`m + on () group_left (,`)
	f(`m + on () group_left (,foo`)
	f(`m + on () group_left (foo,)`)
	f(`m + on () group_left (,foo)`)
	f(`m + on () group_left (foo)`)
	f(`m + on () group_right (foo) (m`)
	f(`m or ignoring () group_left () n`)
	f(`1 + bool 2`)
	f(`m % bool n`)
	f(`m * bool baz`)
	f(`M * BOoL BaZ`)
	f(`foo unless ignoring (bar) group_left xxx`)
	f(`foo or bool bar`)
	f(`foo == bool $$`)
	f(`"foo" + bar`)

	// invalid parensExpr
	f(`(`)
	f(`($`)
	f(`(+`)
	f(`(1`)
	f(`(m+`)
	f(`1)`)
	f(`(,)`)
	f(`(1)$`)

	// invalid funcExpr
	f(`f $`)
	f(`f($)`)
	f(`f[`)
	f(`f()$`)
	f(`f(`)
	f(`f(foo`)
	f(`f(f,`)
	f(`f(,`)
	f(`f(,)`)
	f(`f(,foo)`)
	f(`f(,foo`)
	f(`f(foo,$`)
	f(`f() by (a)`)
	f(`f without (x) (y)`)
	f(`f() foo (a)`)
	f(`f bar (x) (b)`)
	f(`f bar (x)`)

	// invalid aggrFuncExpr
	f(`sum(`)
	f(`sum $`)
	f(`sum [`)
	f(`sum($)`)
	f(`sum()$`)
	f(`sum(foo) ba`)
	f(`sum(foo) ba()`)
	f(`sum(foo) by`)
	f(`sum(foo) without x`)
	f(`sum(foo) aaa`)
	f(`sum(foo) aaa x`)
	f(`sum() by $`)
	f(`sum() by (`)
	f(`sum() by ($`)
	f(`sum() by (a`)
	f(`sum() by (a $`)
	f(`sum() by (a ]`)
	f(`sum() by (a)$`)
	f(`sum() by (,`)
	f(`sum() by (a,$`)
	f(`sum() by (,)`)
	f(`sum() by (,a`)
	f(`sum() by (,a)`)
	f(`sum() on (b)`)
	f(`sum() bool`)
	f(`sum() group_left`)
	f(`sum() group_right(x)`)
	f(`sum ba`)
	f(`sum ba ()`)
	f(`sum by (`)
	f(`sum by (a`)
	f(`sum by (,`)
	f(`sum by (,)`)
	f(`sum by (,a`)
	f(`sum by (,a)`)
	f(`sum by (a)`)
	f(`sum by (a) (`)
	f(`sum by (a) [`)
	f(`sum by (a) {`)
	f(`sum by (a) (b`)
	f(`sum by (a) (b,`)
	f(`sum by (a) (,)`)
	f(`avg by (a) (,b)`)
	f(`sum by (x) (y) by (z)`)
	f(`sum(m) by (1)`)

	// invalid withExpr
	f(`with $`)
	f(`with a`)
	f(`with a=b c`)
	f(`with (`)
	f(`with (x=b)$`)
	f(`with ($`)
	f(`with (foo`)
	f(`with (foo $`)
	f(`with (x y`)
	f(`with (x =`)
	f(`with (x = $`)
	f(`with (x= y`)
	f(`with (x= y $`)
	f(`with (x= y)`)
	f(`with (x=(`)
	f(`with (x=[)`)
	f(`with (x=() x)`)
	f(`with ($$)`)
	f(`with (x $$`)
	f(`with (x = $$)`)
	f(`with (x = foo) bar{x}`)
	f(`with (x = {foo="bar"}[5m]) bar{x}`)
	f(`with (x = {foo="bar"} offset 5m) bar{x}`)
	f(`with (x = a, x = b) c`)
	f(`with (x(a, a) = b) c`)
	f(`with (x=m{f="x"}) foo{x}`)
	f(`with (sum = x) y`)
	f(`with (rate(a) = b) c`)
	f(`with (clamp_min=x) y`)
	f(`with (f()`)
	f(`with (a=b c=d) e`)
	f(`with (f(x)=x^2) m{x}`)
	f(`with (f(x)=ff()) m{x}`)
	f(`with (f(x`)
	f(`with (x=m) a{x} + b`)
	f(`with (x=m) b + a{x}`)
	f(`with (x=m) f(b, a{x})`)
	f(`with (x=m) sum(a{x})`)
	f(`with (x=m) (a{x})`)
	f(`with (f(a)=a) f(1, 2)`)
	f(`with (f(x)=x{foo="bar"}) f(1)`)
	f(`with (f(x)=x{foo="bar"}) f(m + n)`)
	f(`with (f = with`)
	f(`with (,)`)
	f(`with (1) 2`)
	f(`with (f(1)=2) 3`)
	f(`with (f(,)=x) x`)
	f(`with (x(a) = {b="c"}) foo{x}`)
	f(`with (f(x) = m{foo=xx}) f("qwe")`)
	f(`a + with(f(x)=x) f(1,2)`)
	f(`with (f(x) = sum(m) by (x)) f({foo="bar"})`)
	f(`with (f(x) = sum(m) by (x)) f((xx(), {foo="bar"}))`)
	f(`with (f(x) = m + on (x) n) f(xx())`)
	f(`with (f(x) = m + on (a) group_right (x) n) f(xx())`)
}
