package promql

import (
	"math"
	"testing"
)

var (
	testValues     = []float64{123, 34, 44, 21, 54, 34, 99, 12, 44, 32, 34, 34}
	testTimestamps = []int64{5, 15, 24, 36, 49, 60, 78, 80, 97, 115, 120, 130}
)

func TestRemoveCounterResets(t *testing.T) {
	removeCounterResets(nil)

	values := append([]float64{}, testValues...)
	removeCounterResets(values)
	valuesExpected := []float64{123, 157, 167, 188, 221, 255, 320, 332, 364, 396, 398, 398}
	testRowsEqual(t, values, testTimestamps, valuesExpected, testTimestamps)

	// removeCounterResets doesn't expect negative values, so it doesn't work properly with them.
	values = []float64{-100, -200, -300, -400}
	removeCounterResets(values)
	valuesExpected = []float64{-100, -300, -600, -1000}
	timestampsExpected := []int64{0, 1, 2, 3}
	testRowsEqual(t, values, timestampsExpected, valuesExpected, timestampsExpected)

	// verify how jitter from `Prometheus HA pairs` is handled
	values = []float64{100, 95, 120, 140, 137, 50}
	removeCounterResets(values)
	valuesExpected = []float64{100, 100, 120, 140, 140, 190}
	timestampsExpected = []int64{0, 1, 2, 3, 4, 5}
	testRowsEqual(t, values, timestampsExpected, valuesExpected, timestampsExpected)
}

func TestDeltaValues(t *testing.T) {
	deltaValues(nil)

	values := []float64{123}
	deltaValues(values)
	valuesExpected := []float64{nan}
	testRowsEqual(t, values, testTimestamps[:1], valuesExpected, testTimestamps[:1])

	values = append([]float64{}, testValues...)
	deltaValues(values)
	valuesExpected = []float64{-89, 10, -23, 33, -20, 65, -87, 32, -12, 2, 0, nan}
	testRowsEqual(t, values, testTimestamps, valuesExpected, testTimestamps)

	// remove counter resets
	values = append([]float64{}, testValues...)
	removeCounterResets(values)
	deltaValues(values)
	valuesExpected = []float64{34, 10, 21, 33, 34, 65, 12, 32, 32, 2, 0, nan}
	testRowsEqual(t, values, testTimestamps, valuesExpected, testTimestamps)
}

func TestDerivValues(t *testing.T) {
	derivValues(nil, nil)

	values := []float64{123}
	derivValues(values, testTimestamps[:1])
	valuesExpected := []float64{nan}
	testRowsEqual(t, values, testTimestamps[:1], valuesExpected, testTimestamps[:1])

	values = append([]float64{}, testValues...)
	derivValues(values, testTimestamps)
	valuesExpected = []float64{-8900, 1111.111111111111, -1916.6666666666665, 2538.461538461538, -1818.1818181818182, 3611.111111111111,
		-43500, 1882.3529411764705, -666.6666666666666, 400, 0, nan}
	testRowsEqual(t, values, testTimestamps, valuesExpected, testTimestamps)

	// remove counter resets
	values = append([]float64{}, testValues...)
	removeCounterResets(values)
	derivValues(values, testTimestamps)
	valuesExpected = []float64{3400, 1111.111111111111, 1750, 2538.461538461538, 3090.909090909091, 3611.111111111111,
		6000, 1882.3529411764705, 1777.7777777777776, 400, 0, nan}
	testRowsEqual(t, values, testTimestamps, valuesExpected, testTimestamps)
}

func testRollupFunc(t *testing.T, funcName string, args []interface{}, meExpected *metricExpr, vExpected float64) {
	t.Helper()
	nrf := getRollupFunc(funcName)
	if nrf == nil {
		t.Fatalf("cannot obtain %q", funcName)
	}
	rf, err := nrf(args)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	var rfa rollupFuncArg
	rfa.prevValue = nan
	rfa.prevTimestamp = 0
	rfa.values = append(rfa.values, testValues...)
	rfa.timestamps = append(rfa.timestamps, testTimestamps...)
	if rollupFuncsRemoveCounterResets[funcName] {
		removeCounterResets(rfa.values)
	}
	for i := 0; i < 5; i++ {
		v := rf(&rfa)
		if math.IsNaN(vExpected) {
			if !math.IsNaN(v) {
				t.Fatalf("unexpected value; got %v; want %v", v, vExpected)
			}
		} else {
			if v != vExpected {
				t.Fatalf("unexpected value; got %v; want %v", v, vExpected)
			}
		}
	}
}

func TestRollupQuantileOverTime(t *testing.T) {
	f := func(phi, vExpected float64) {
		t.Helper()
		phis := []*timeseries{{
			Values:     []float64{phi},
			Timestamps: []int64{123},
		}}
		var me metricExpr
		args := []interface{}{phis, &rollupExpr{Expr: &me}}
		testRollupFunc(t, "quantile_over_time", args, &me, vExpected)
	}

	f(-123, 12)
	f(-0.5, 12)
	f(0, 12)
	f(0.1, 21)
	f(0.5, 34)
	f(0.9, 99)
	f(1, 123)
	f(234, 123)
}

func TestRollupPredictLinear(t *testing.T) {
	f := func(sec, vExpected float64) {
		t.Helper()
		secs := []*timeseries{{
			Values:     []float64{sec},
			Timestamps: []int64{123},
		}}
		var me metricExpr
		args := []interface{}{&rollupExpr{Expr: &me}, secs}
		testRollupFunc(t, "predict_linear", args, &me, vExpected)
	}

	f(0e-3, 63.739757761102624)
	f(50e-3, 50.39682764539959)
	f(100e-3, 37.053897529696556)
	f(200e-3, 10.368037298290488)
}

func TestRollupHoltWinters(t *testing.T) {
	f := func(sf, tf, vExpected float64) {
		t.Helper()
		sfs := []*timeseries{{
			Values:     []float64{sf},
			Timestamps: []int64{123},
		}}
		tfs := []*timeseries{{
			Values:     []float64{tf},
			Timestamps: []int64{123},
		}}
		var me metricExpr
		args := []interface{}{&rollupExpr{Expr: &me}, sfs, tfs}
		testRollupFunc(t, "holt_winters", args, &me, vExpected)
	}

	f(-1, 0.5, nan)
	f(0, 0.5, nan)
	f(1, 0.5, nan)
	f(2, 0.5, nan)
	f(0.5, -1, nan)
	f(0.5, 0, nan)
	f(0.5, 1, nan)
	f(0.5, 2, nan)
	f(0.5, 0.5, 34.97794532775879)
	f(0.1, 0.5, -131.30529492371622)
	f(0.1, 0.1, -397.3307790780296)
	f(0.5, 0.1, -5.791530520284198)
	f(0.5, 0.9, 25.498906408926757)
	f(0.9, 0.9, 33.99637566941818)
}

func TestRollupNewRollupFuncSuccess(t *testing.T) {
	f := func(funcName string, vExpected float64) {
		t.Helper()
		var me metricExpr
		args := []interface{}{&rollupExpr{Expr: &me}}
		testRollupFunc(t, funcName, args, &me, vExpected)
	}

	f("default_rollup", 34)
	f("changes", 10)
	f("delta", -89)
	f("deriv", -266.85860231406065)
	f("idelta", 0)
	f("increase", 275)
	f("irate", 0)
	f("rate", 2200)
	f("resets", 5)
	f("avg_over_time", 47.083333333333336)
	f("min_over_time", 12)
	f("max_over_time", 123)
	f("sum_over_time", 565)
	f("count_over_time", 12)
	f("stddev_over_time", 30.752935722554287)
	f("stdvar_over_time", 945.7430555555555)
	f("first_over_time", 123)
	f("last_over_time", 34)
	f("integrate", 61.0275)
}

func TestRollupNewRollupFuncError(t *testing.T) {
	if nrf := getRollupFunc("non-existing-func"); nrf != nil {
		t.Fatalf("expecting nil func; got %p", nrf)
	}

	f := func(funcName string, args []interface{}) {
		t.Helper()

		nrf := getRollupFunc(funcName)
		rf, err := nrf(args)
		if err == nil {
			t.Fatalf("expecting non-nil error")
		}
		if rf != nil {
			t.Fatalf("expecting nil rf; got %p", rf)
		}
	}

	// Invalid number of args
	f("default_rollup", nil)
	f("holt_winters", nil)
	f("predict_linear", nil)
	f("quantile_over_time", nil)

	// Invalid arg type
	scalarTs := []*timeseries{{
		Values:     []float64{321},
		Timestamps: []int64{123},
	}}
	me := &metricExpr{}
	f("holt_winters", []interface{}{123, 123, 321})
	f("holt_winters", []interface{}{me, 123, 321})
	f("holt_winters", []interface{}{me, scalarTs, 321})
	f("predict_linear", []interface{}{123, 123})
	f("predict_linear", []interface{}{me, 123})
	f("quantile_over_time", []interface{}{123, 123})
}

func TestRollupNoWindowNoPoints(t *testing.T) {
	t.Run("beforeStart", func(t *testing.T) {
		rc := rollupConfig{
			Func:   rollupFirst,
			Start:  0,
			End:    4,
			Step:   1,
			Window: 0,
		}
		rc.Timestamps = getTimestamps(rc.Start, rc.End, rc.Step)
		values := rc.Do(nil, testValues, testTimestamps)
		valuesExpected := []float64{nan, nan, nan, nan, nan}
		timestampsExpected := []int64{0, 1, 2, 3, 4}
		testRowsEqual(t, values, rc.Timestamps, valuesExpected, timestampsExpected)
	})
	t.Run("afterEnd", func(t *testing.T) {
		rc := rollupConfig{
			Func:   rollupDelta,
			Start:  120,
			End:    144,
			Step:   4,
			Window: 0,
		}
		rc.Timestamps = getTimestamps(rc.Start, rc.End, rc.Step)
		values := rc.Do(nil, testValues, testTimestamps)
		valuesExpected := []float64{2, 2, 2, 0, 0, 0, nan}
		timestampsExpected := []int64{120, 124, 128, 132, 136, 140, 144}
		testRowsEqual(t, values, rc.Timestamps, valuesExpected, timestampsExpected)
	})
}

func TestRollupWindowNoPoints(t *testing.T) {
	t.Run("beforeStart", func(t *testing.T) {
		rc := rollupConfig{
			Func:   rollupFirst,
			Start:  0,
			End:    4,
			Step:   1,
			Window: 3,
		}
		rc.Timestamps = getTimestamps(rc.Start, rc.End, rc.Step)
		values := rc.Do(nil, testValues, testTimestamps)
		valuesExpected := []float64{nan, nan, nan, nan, nan}
		timestampsExpected := []int64{0, 1, 2, 3, 4}
		testRowsEqual(t, values, rc.Timestamps, valuesExpected, timestampsExpected)
	})
	t.Run("afterEnd", func(t *testing.T) {
		rc := rollupConfig{
			Func:   rollupFirst,
			Start:  141,
			End:    171,
			Step:   10,
			Window: 3,
		}
		rc.Timestamps = getTimestamps(rc.Start, rc.End, rc.Step)
		values := rc.Do(nil, testValues, testTimestamps)
		valuesExpected := []float64{34, 34, nan, nan}
		timestampsExpected := []int64{141, 151, 161, 171}
		testRowsEqual(t, values, rc.Timestamps, valuesExpected, timestampsExpected)
	})
}

func TestRollupNoWindowPartialPoints(t *testing.T) {
	t.Run("beforeStart", func(t *testing.T) {
		rc := rollupConfig{
			Func:   rollupFirst,
			Start:  0,
			End:    20,
			Step:   5,
			Window: 0,
		}
		rc.Timestamps = getTimestamps(rc.Start, rc.End, rc.Step)
		values := rc.Do(nil, testValues, testTimestamps)
		valuesExpected := []float64{nan, 123, 123, 123, 123}
		timestampsExpected := []int64{0, 5, 10, 15, 20}
		testRowsEqual(t, values, rc.Timestamps, valuesExpected, timestampsExpected)
	})
	t.Run("afterEnd", func(t *testing.T) {
		rc := rollupConfig{
			Func:   rollupFirst,
			Start:  100,
			End:    160,
			Step:   20,
			Window: 0,
		}
		rc.Timestamps = getTimestamps(rc.Start, rc.End, rc.Step)
		values := rc.Do(nil, testValues, testTimestamps)
		valuesExpected := []float64{12, 44, 34, 34}
		timestampsExpected := []int64{100, 120, 140, 160}
		testRowsEqual(t, values, rc.Timestamps, valuesExpected, timestampsExpected)
	})
	t.Run("middle", func(t *testing.T) {
		rc := rollupConfig{
			Func:   rollupFirst,
			Start:  -50,
			End:    150,
			Step:   50,
			Window: 0,
		}
		rc.Timestamps = getTimestamps(rc.Start, rc.End, rc.Step)
		values := rc.Do(nil, testValues, testTimestamps)
		valuesExpected := []float64{nan, nan, 123, 54, 44}
		timestampsExpected := []int64{-50, 0, 50, 100, 150}
		testRowsEqual(t, values, rc.Timestamps, valuesExpected, timestampsExpected)
	})
}

func TestRollupWindowPartialPoints(t *testing.T) {
	t.Run("beforeStart", func(t *testing.T) {
		rc := rollupConfig{
			Func:   rollupLast,
			Start:  0,
			End:    20,
			Step:   5,
			Window: 8,
		}
		rc.Timestamps = getTimestamps(rc.Start, rc.End, rc.Step)
		values := rc.Do(nil, testValues, testTimestamps)
		valuesExpected := []float64{nan, 123, 123, 34, 34}
		timestampsExpected := []int64{0, 5, 10, 15, 20}
		testRowsEqual(t, values, rc.Timestamps, valuesExpected, timestampsExpected)
	})
	t.Run("afterEnd", func(t *testing.T) {
		rc := rollupConfig{
			Func:   rollupLast,
			Start:  100,
			End:    160,
			Step:   20,
			Window: 18,
		}
		rc.Timestamps = getTimestamps(rc.Start, rc.End, rc.Step)
		values := rc.Do(nil, testValues, testTimestamps)
		valuesExpected := []float64{44, 34, 34, nan}
		timestampsExpected := []int64{100, 120, 140, 160}
		testRowsEqual(t, values, rc.Timestamps, valuesExpected, timestampsExpected)
	})
	t.Run("middle", func(t *testing.T) {
		rc := rollupConfig{
			Func:   rollupLast,
			Start:  0,
			End:    150,
			Step:   50,
			Window: 19,
		}
		rc.Timestamps = getTimestamps(rc.Start, rc.End, rc.Step)
		values := rc.Do(nil, testValues, testTimestamps)
		valuesExpected := []float64{nan, 54, 44, nan}
		timestampsExpected := []int64{0, 50, 100, 150}
		testRowsEqual(t, values, rc.Timestamps, valuesExpected, timestampsExpected)
	})
}

func TestRollupFuncsNoWindow(t *testing.T) {
	t.Run("first", func(t *testing.T) {
		rc := rollupConfig{
			Func:   rollupFirst,
			Start:  0,
			End:    160,
			Step:   40,
			Window: 0,
		}
		rc.Timestamps = getTimestamps(rc.Start, rc.End, rc.Step)
		values := rc.Do(nil, testValues, testTimestamps)
		valuesExpected := []float64{nan, 123, 21, 12, 34}
		timestampsExpected := []int64{0, 40, 80, 120, 160}
		testRowsEqual(t, values, rc.Timestamps, valuesExpected, timestampsExpected)
	})
	t.Run("count", func(t *testing.T) {
		rc := rollupConfig{
			Func:   rollupCount,
			Start:  0,
			End:    160,
			Step:   40,
			Window: 0,
		}
		rc.Timestamps = getTimestamps(rc.Start, rc.End, rc.Step)
		values := rc.Do(nil, testValues, testTimestamps)
		valuesExpected := []float64{nan, 4, 4, 3, 1}
		timestampsExpected := []int64{0, 40, 80, 120, 160}
		testRowsEqual(t, values, rc.Timestamps, valuesExpected, timestampsExpected)
	})
	t.Run("min", func(t *testing.T) {
		rc := rollupConfig{
			Func:   rollupMin,
			Start:  0,
			End:    160,
			Step:   40,
			Window: 0,
		}
		rc.Timestamps = getTimestamps(rc.Start, rc.End, rc.Step)
		values := rc.Do(nil, testValues, testTimestamps)
		valuesExpected := []float64{nan, 21, 12, 32, 34}
		timestampsExpected := []int64{0, 40, 80, 120, 160}
		testRowsEqual(t, values, rc.Timestamps, valuesExpected, timestampsExpected)
	})
	t.Run("max", func(t *testing.T) {
		rc := rollupConfig{
			Func:   rollupMax,
			Start:  0,
			End:    160,
			Step:   40,
			Window: 0,
		}
		rc.Timestamps = getTimestamps(rc.Start, rc.End, rc.Step)
		values := rc.Do(nil, testValues, testTimestamps)
		valuesExpected := []float64{nan, 123, 99, 44, 34}
		timestampsExpected := []int64{0, 40, 80, 120, 160}
		testRowsEqual(t, values, rc.Timestamps, valuesExpected, timestampsExpected)
	})
	t.Run("sum", func(t *testing.T) {
		rc := rollupConfig{
			Func:   rollupSum,
			Start:  0,
			End:    160,
			Step:   40,
			Window: 0,
		}
		rc.Timestamps = getTimestamps(rc.Start, rc.End, rc.Step)
		values := rc.Do(nil, testValues, testTimestamps)
		valuesExpected := []float64{nan, 222, 199, 110, 34}
		timestampsExpected := []int64{0, 40, 80, 120, 160}
		testRowsEqual(t, values, rc.Timestamps, valuesExpected, timestampsExpected)
	})
	t.Run("delta", func(t *testing.T) {
		rc := rollupConfig{
			Func:   rollupDelta,
			Start:  0,
			End:    160,
			Step:   40,
			Window: 0,
		}
		rc.Timestamps = getTimestamps(rc.Start, rc.End, rc.Step)
		values := rc.Do(nil, testValues, testTimestamps)
		valuesExpected := []float64{nan, -102, -9, 22, 0}
		timestampsExpected := []int64{0, 40, 80, 120, 160}
		testRowsEqual(t, values, rc.Timestamps, valuesExpected, timestampsExpected)
	})
	t.Run("idelta", func(t *testing.T) {
		rc := rollupConfig{
			Func:   rollupIdelta,
			Start:  10,
			End:    130,
			Step:   40,
			Window: 0,
		}
		rc.Timestamps = getTimestamps(rc.Start, rc.End, rc.Step)
		values := rc.Do(nil, testValues, testTimestamps)
		valuesExpected := []float64{nan, 33, -87, 0}
		timestampsExpected := []int64{10, 50, 90, 130}
		testRowsEqual(t, values, rc.Timestamps, valuesExpected, timestampsExpected)
	})
	t.Run("changes", func(t *testing.T) {
		rc := rollupConfig{
			Func:   rollupChanges,
			Start:  0,
			End:    160,
			Step:   40,
			Window: 0,
		}
		rc.Timestamps = getTimestamps(rc.Start, rc.End, rc.Step)
		values := rc.Do(nil, testValues, testTimestamps)
		valuesExpected := []float64{nan, 3, 4, 3, 0}
		timestampsExpected := []int64{0, 40, 80, 120, 160}
		testRowsEqual(t, values, rc.Timestamps, valuesExpected, timestampsExpected)
	})
	t.Run("resets", func(t *testing.T) {
		rc := rollupConfig{
			Func:   rollupResets,
			Start:  0,
			End:    160,
			Step:   40,
			Window: 0,
		}
		rc.Timestamps = getTimestamps(rc.Start, rc.End, rc.Step)
		values := rc.Do(nil, testValues, testTimestamps)
		valuesExpected := []float64{nan, 2, 2, 1, 0}
		timestampsExpected := []int64{0, 40, 80, 120, 160}
		testRowsEqual(t, values, rc.Timestamps, valuesExpected, timestampsExpected)
	})
	t.Run("avg", func(t *testing.T) {
		rc := rollupConfig{
			Func:   rollupAvg,
			Start:  0,
			End:    160,
			Step:   40,
			Window: 0,
		}
		rc.Timestamps = getTimestamps(rc.Start, rc.End, rc.Step)
		values := rc.Do(nil, testValues, testTimestamps)
		valuesExpected := []float64{nan, 55.5, 49.75, 36.666666666666664, 34}
		timestampsExpected := []int64{0, 40, 80, 120, 160}
		testRowsEqual(t, values, rc.Timestamps, valuesExpected, timestampsExpected)
	})
	t.Run("deriv", func(t *testing.T) {
		rc := rollupConfig{
			Func:   rollupDerivSlow,
			Start:  0,
			End:    160,
			Step:   40,
			Window: 0,
		}
		rc.Timestamps = getTimestamps(rc.Start, rc.End, rc.Step)
		values := rc.Do(nil, testValues, testTimestamps)
		valuesExpected := []float64{nan, -2879.310344827587, 558.0608793686592, 422.84569138276544, 0}
		timestampsExpected := []int64{0, 40, 80, 120, 160}
		testRowsEqual(t, values, rc.Timestamps, valuesExpected, timestampsExpected)
	})
	t.Run("ideriv", func(t *testing.T) {
		rc := rollupConfig{
			Func:   rollupIderiv,
			Start:  0,
			End:    160,
			Step:   40,
			Window: 0,
		}
		rc.Timestamps = getTimestamps(rc.Start, rc.End, rc.Step)
		values := rc.Do(nil, testValues, testTimestamps)
		valuesExpected := []float64{nan, -1916.6666666666665, -43500, 400, 0}
		timestampsExpected := []int64{0, 40, 80, 120, 160}
		testRowsEqual(t, values, rc.Timestamps, valuesExpected, timestampsExpected)
	})
	t.Run("stddev", func(t *testing.T) {
		rc := rollupConfig{
			Func:   rollupStddev,
			Start:  0,
			End:    160,
			Step:   40,
			Window: 0,
		}
		rc.Timestamps = getTimestamps(rc.Start, rc.End, rc.Step)
		values := rc.Do(nil, testValues, testTimestamps)
		valuesExpected := []float64{nan, 39.81519810323691, 32.080952292598795, 5.2493385826745405, 0}
		timestampsExpected := []int64{0, 40, 80, 120, 160}
		testRowsEqual(t, values, rc.Timestamps, valuesExpected, timestampsExpected)
	})
	t.Run("integrate", func(t *testing.T) {
		rc := rollupConfig{
			Func:   rollupIntegrate,
			Start:  0,
			End:    160,
			Step:   40,
			Window: 0,
		}
		rc.Timestamps = getTimestamps(rc.Start, rc.End, rc.Step)
		values := rc.Do(nil, testValues, testTimestamps)
		valuesExpected := []float64{nan, 4.6035, 4.3934999999999995, 2.166, 0.34}
		timestampsExpected := []int64{0, 40, 80, 120, 160}
		testRowsEqual(t, values, rc.Timestamps, valuesExpected, timestampsExpected)
	})
	t.Run("distinct", func(t *testing.T) {
		rc := rollupConfig{
			Func:   rollupDistinct,
			Start:  0,
			End:    160,
			Step:   40,
			Window: 0,
		}
		rc.Timestamps = getTimestamps(rc.Start, rc.End, rc.Step)
		values := rc.Do(nil, testValues, testTimestamps)
		valuesExpected := []float64{nan, 4, 4, 3, 1}
		timestampsExpected := []int64{0, 40, 80, 120, 160}
		testRowsEqual(t, values, rc.Timestamps, valuesExpected, timestampsExpected)
	})
}

func testRowsEqual(t *testing.T, values []float64, timestamps []int64, valuesExpected []float64, timestampsExpected []int64) {
	t.Helper()
	if len(values) != len(valuesExpected) {
		t.Fatalf("unexpected len(values); got %d; want %d\nvalues=\n%v\nvaluesExpected=\n%v",
			len(values), len(valuesExpected), values, valuesExpected)
	}
	if len(timestamps) != len(timestampsExpected) {
		t.Fatalf("unexpected len(timestamps); got %d; want %d\ntimestamps=\n%v\ntimestampsExpected=\n%v",
			len(timestamps), len(timestampsExpected), timestamps, timestampsExpected)
	}
	if len(values) != len(timestamps) {
		t.Fatalf("len(values) doesn't match len(timestamps); got %d vs %d", len(values), len(timestamps))
	}
	for i, v := range values {
		ts := timestamps[i]
		tsExpected := timestampsExpected[i]
		if ts != tsExpected {
			t.Fatalf("unexpected timestamp at timestamps[%d]; got %d; want %d\ntimestamps=\n%v\ntimestampsExpected=\n%v",
				i, ts, tsExpected, timestamps, timestampsExpected)
		}
		vExpected := valuesExpected[i]
		if math.IsNaN(v) {
			if !math.IsNaN(vExpected) {
				t.Fatalf("unexpected nan value at values[%d]; want %f\nvalues=\n%v\nvaluesExpected=\n%v",
					i, vExpected, values, valuesExpected)
			}
			continue
		}
		if v != vExpected {
			t.Fatalf("unexpected value at values[%d]; got %f; want %f\nvalues=\n%v\nvaluesExpected=\n%v",
				i, v, vExpected, values, valuesExpected)
		}
	}
}
