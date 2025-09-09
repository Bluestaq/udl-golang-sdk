// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/Bluestaq/udl-golang-sdk"
	"github.com/Bluestaq/udl-golang-sdk/internal/testutil"
	"github.com/Bluestaq/udl-golang-sdk/option"
)

func TestStateVectorNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := unifieddatalibrary.NewClient(
		option.WithBaseURL(baseURL),
		option.WithPassword("My Password"),
		option.WithUsername("My Username"),
	)
	err := client.StateVector.New(context.TODO(), unifieddatalibrary.StateVectorNewParams{
		StateVectorIngest: unifieddatalibrary.StateVectorIngestParam{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.StateVectorIngestDataModeTest,
			Epoch:                 time.Now(),
			Source:                "Bluestaq",
			ActualOdSpan:          unifieddatalibrary.Float(3.5),
			Algorithm:             unifieddatalibrary.String("SAMPLE_ALGORITHM"),
			Alt1ReferenceFrame:    unifieddatalibrary.String("TEME"),
			Alt2ReferenceFrame:    unifieddatalibrary.String("EFG/TDR"),
			Area:                  unifieddatalibrary.Float(5.065),
			BDot:                  unifieddatalibrary.Float(1.23),
			CmOffset:              unifieddatalibrary.Float(1.23),
			Cov:                   []float64{1.1, 2.4, 3.8, 4.2, 5.5, 6},
			CovMethod:             unifieddatalibrary.String("CALCULATED"),
			CovReferenceFrame:     unifieddatalibrary.StateVectorIngestCovReferenceFrameJ2000,
			Descriptor:            unifieddatalibrary.String("descriptor"),
			DragArea:              unifieddatalibrary.Float(4.739),
			DragCoeff:             unifieddatalibrary.Float(0.0224391269775),
			DragModel:             unifieddatalibrary.String("JAC70"),
			Edr:                   unifieddatalibrary.Float(1.23),
			EqCov:                 []float64{1.1, 2.2},
			ErrorControl:          unifieddatalibrary.Float(1.23),
			FixedStep:             unifieddatalibrary.Bool(true),
			GeopotentialModel:     unifieddatalibrary.String("EGM-96"),
			Iau1980Terms:          unifieddatalibrary.Int(4),
			IDOrbitDetermination:  unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
			IDStateVector:         unifieddatalibrary.String("STATEVECTOR-ID"),
			IntegratorMode:        unifieddatalibrary.String("integratorMode"),
			InTrackThrust:         unifieddatalibrary.Bool(true),
			LastObEnd:             unifieddatalibrary.Time(time.Now()),
			LastObStart:           unifieddatalibrary.Time(time.Now()),
			LeapSecondTime:        unifieddatalibrary.Time(time.Now()),
			LunarSolar:            unifieddatalibrary.Bool(true),
			Mass:                  unifieddatalibrary.Float(164.5),
			ObsAvailable:          unifieddatalibrary.Int(376),
			ObsUsed:               unifieddatalibrary.Int(374),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigObjectID:          unifieddatalibrary.String("ORIGOBJECT-ID"),
			Partials:              unifieddatalibrary.String("ANALYTIC"),
			Pedigree:              unifieddatalibrary.String("CONJUNCTION"),
			PolarMotionX:          unifieddatalibrary.Float(1.23),
			PolarMotionY:          unifieddatalibrary.Float(1.23),
			PosUnc:                unifieddatalibrary.Float(0.333399744452),
			RawFileUri:            unifieddatalibrary.String("rawFileURI"),
			RecOdSpan:             unifieddatalibrary.Float(3.5),
			ReferenceFrame:        unifieddatalibrary.StateVectorIngestReferenceFrameJ2000,
			ResidualsAcc:          unifieddatalibrary.Float(99.5),
			RevNo:                 unifieddatalibrary.Int(7205),
			Rms:                   unifieddatalibrary.Float(0.991),
			SatNo:                 unifieddatalibrary.Int(12),
			SigmaPosUvw:           []float64{1.23, 4.56},
			SigmaVelUvw:           []float64{1.23, 4.56},
			SolarFluxApAvg:        unifieddatalibrary.Float(1.23),
			SolarFluxF10:          unifieddatalibrary.Float(1.23),
			SolarFluxF10Avg:       unifieddatalibrary.Float(1.23),
			SolarRadPress:         unifieddatalibrary.Bool(true),
			SolarRadPressCoeff:    unifieddatalibrary.Float(0.0244394),
			SolidEarthTides:       unifieddatalibrary.Bool(true),
			SourcedData:           []string{"DATA1", "DATA2"},
			SourcedDataTypes:      []string{"RADAR"},
			SrpArea:               unifieddatalibrary.Float(4.311),
			StepMode:              unifieddatalibrary.String("AUTO"),
			StepSize:              unifieddatalibrary.Float(1.23),
			StepSizeSelection:     unifieddatalibrary.String("AUTO"),
			Tags:                  []string{"TAG1", "TAG2"},
			TaiUtc:                unifieddatalibrary.Float(1.23),
			ThrustAccel:           unifieddatalibrary.Float(1.23),
			TracksAvail:           unifieddatalibrary.Int(163),
			TracksUsed:            unifieddatalibrary.Int(163),
			TransactionID:         unifieddatalibrary.String("transactionId"),
			Uct:                   unifieddatalibrary.Bool(true),
			Ut1Rate:               unifieddatalibrary.Float(1.23),
			Ut1Utc:                unifieddatalibrary.Float(1.23),
			VelUnc:                unifieddatalibrary.Float(0.000004),
			Xaccel:                unifieddatalibrary.Float(-2.12621392),
			Xpos:                  unifieddatalibrary.Float(-1118.577381),
			XposAlt1:              unifieddatalibrary.Float(-1145.688502),
			XposAlt2:              unifieddatalibrary.Float(-1456.915926),
			Xvel:                  unifieddatalibrary.Float(-4.25242784),
			XvelAlt1:              unifieddatalibrary.Float(-4.270832252),
			XvelAlt2:              unifieddatalibrary.Float(-1.219814294),
			Yaccel:                unifieddatalibrary.Float(2.645553717),
			Ypos:                  unifieddatalibrary.Float(3026.231084),
			YposAlt1:              unifieddatalibrary.Float(3020.729572),
			YposAlt2:              unifieddatalibrary.Float(-2883.540406),
			Yvel:                  unifieddatalibrary.Float(5.291107434),
			YvelAlt1:              unifieddatalibrary.Float(5.27074276),
			YvelAlt2:              unifieddatalibrary.Float(-6.602080212),
			Zaccel:                unifieddatalibrary.Float(-1.06310696),
			Zpos:                  unifieddatalibrary.Float(6167.831808),
			ZposAlt1:              unifieddatalibrary.Float(6165.55187),
			ZposAlt2:              unifieddatalibrary.Float(6165.55187),
			Zvel:                  unifieddatalibrary.Float(-3.356493869),
			ZvelAlt1:              unifieddatalibrary.Float(-3.365155181),
			ZvelAlt2:              unifieddatalibrary.Float(-3.365155181),
		},
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestStateVectorListWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := unifieddatalibrary.NewClient(
		option.WithBaseURL(baseURL),
		option.WithPassword("My Password"),
		option.WithUsername("My Username"),
	)
	_, err := client.StateVector.List(context.TODO(), unifieddatalibrary.StateVectorListParams{
		Epoch:       time.Now(),
		FirstResult: unifieddatalibrary.Int(0),
		MaxResults:  unifieddatalibrary.Int(0),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestStateVectorCountWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := unifieddatalibrary.NewClient(
		option.WithBaseURL(baseURL),
		option.WithPassword("My Password"),
		option.WithUsername("My Username"),
	)
	_, err := client.StateVector.Count(context.TODO(), unifieddatalibrary.StateVectorCountParams{
		Epoch:       time.Now(),
		FirstResult: unifieddatalibrary.Int(0),
		MaxResults:  unifieddatalibrary.Int(0),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestStateVectorNewBulk(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := unifieddatalibrary.NewClient(
		option.WithBaseURL(baseURL),
		option.WithPassword("My Password"),
		option.WithUsername("My Username"),
	)
	err := client.StateVector.NewBulk(context.TODO(), unifieddatalibrary.StateVectorNewBulkParams{
		Body: []unifieddatalibrary.StateVectorIngestParam{{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.StateVectorIngestDataModeTest,
			Epoch:                 time.Now(),
			Source:                "Bluestaq",
			ActualOdSpan:          unifieddatalibrary.Float(3.5),
			Algorithm:             unifieddatalibrary.String("SAMPLE_ALGORITHM"),
			Alt1ReferenceFrame:    unifieddatalibrary.String("TEME"),
			Alt2ReferenceFrame:    unifieddatalibrary.String("EFG/TDR"),
			Area:                  unifieddatalibrary.Float(5.065),
			BDot:                  unifieddatalibrary.Float(1.23),
			CmOffset:              unifieddatalibrary.Float(1.23),
			Cov:                   []float64{1.1, 2.4, 3.8, 4.2, 5.5, 6},
			CovMethod:             unifieddatalibrary.String("CALCULATED"),
			CovReferenceFrame:     unifieddatalibrary.StateVectorIngestCovReferenceFrameJ2000,
			Descriptor:            unifieddatalibrary.String("descriptor"),
			DragArea:              unifieddatalibrary.Float(4.739),
			DragCoeff:             unifieddatalibrary.Float(0.0224391269775),
			DragModel:             unifieddatalibrary.String("JAC70"),
			Edr:                   unifieddatalibrary.Float(1.23),
			EqCov:                 []float64{1.1, 2.2},
			ErrorControl:          unifieddatalibrary.Float(1.23),
			FixedStep:             unifieddatalibrary.Bool(true),
			GeopotentialModel:     unifieddatalibrary.String("EGM-96"),
			Iau1980Terms:          unifieddatalibrary.Int(4),
			IDOrbitDetermination:  unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
			IDStateVector:         unifieddatalibrary.String("STATEVECTOR-ID"),
			IntegratorMode:        unifieddatalibrary.String("integratorMode"),
			InTrackThrust:         unifieddatalibrary.Bool(true),
			LastObEnd:             unifieddatalibrary.Time(time.Now()),
			LastObStart:           unifieddatalibrary.Time(time.Now()),
			LeapSecondTime:        unifieddatalibrary.Time(time.Now()),
			LunarSolar:            unifieddatalibrary.Bool(true),
			Mass:                  unifieddatalibrary.Float(164.5),
			ObsAvailable:          unifieddatalibrary.Int(376),
			ObsUsed:               unifieddatalibrary.Int(374),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigObjectID:          unifieddatalibrary.String("ORIGOBJECT-ID"),
			Partials:              unifieddatalibrary.String("ANALYTIC"),
			Pedigree:              unifieddatalibrary.String("CONJUNCTION"),
			PolarMotionX:          unifieddatalibrary.Float(1.23),
			PolarMotionY:          unifieddatalibrary.Float(1.23),
			PosUnc:                unifieddatalibrary.Float(0.333399744452),
			RawFileUri:            unifieddatalibrary.String("rawFileURI"),
			RecOdSpan:             unifieddatalibrary.Float(3.5),
			ReferenceFrame:        unifieddatalibrary.StateVectorIngestReferenceFrameJ2000,
			ResidualsAcc:          unifieddatalibrary.Float(99.5),
			RevNo:                 unifieddatalibrary.Int(7205),
			Rms:                   unifieddatalibrary.Float(0.991),
			SatNo:                 unifieddatalibrary.Int(12),
			SigmaPosUvw:           []float64{1.23, 4.56},
			SigmaVelUvw:           []float64{1.23, 4.56},
			SolarFluxApAvg:        unifieddatalibrary.Float(1.23),
			SolarFluxF10:          unifieddatalibrary.Float(1.23),
			SolarFluxF10Avg:       unifieddatalibrary.Float(1.23),
			SolarRadPress:         unifieddatalibrary.Bool(true),
			SolarRadPressCoeff:    unifieddatalibrary.Float(0.0244394),
			SolidEarthTides:       unifieddatalibrary.Bool(true),
			SourcedData:           []string{"DATA1", "DATA2"},
			SourcedDataTypes:      []string{"RADAR"},
			SrpArea:               unifieddatalibrary.Float(4.311),
			StepMode:              unifieddatalibrary.String("AUTO"),
			StepSize:              unifieddatalibrary.Float(1.23),
			StepSizeSelection:     unifieddatalibrary.String("AUTO"),
			Tags:                  []string{"TAG1", "TAG2"},
			TaiUtc:                unifieddatalibrary.Float(1.23),
			ThrustAccel:           unifieddatalibrary.Float(1.23),
			TracksAvail:           unifieddatalibrary.Int(163),
			TracksUsed:            unifieddatalibrary.Int(163),
			TransactionID:         unifieddatalibrary.String("transactionId"),
			Uct:                   unifieddatalibrary.Bool(true),
			Ut1Rate:               unifieddatalibrary.Float(1.23),
			Ut1Utc:                unifieddatalibrary.Float(1.23),
			VelUnc:                unifieddatalibrary.Float(0.000004),
			Xaccel:                unifieddatalibrary.Float(-2.12621392),
			Xpos:                  unifieddatalibrary.Float(-1118.577381),
			XposAlt1:              unifieddatalibrary.Float(-1145.688502),
			XposAlt2:              unifieddatalibrary.Float(-1456.915926),
			Xvel:                  unifieddatalibrary.Float(-4.25242784),
			XvelAlt1:              unifieddatalibrary.Float(-4.270832252),
			XvelAlt2:              unifieddatalibrary.Float(-1.219814294),
			Yaccel:                unifieddatalibrary.Float(2.645553717),
			Ypos:                  unifieddatalibrary.Float(3026.231084),
			YposAlt1:              unifieddatalibrary.Float(3020.729572),
			YposAlt2:              unifieddatalibrary.Float(-2883.540406),
			Yvel:                  unifieddatalibrary.Float(5.291107434),
			YvelAlt1:              unifieddatalibrary.Float(5.27074276),
			YvelAlt2:              unifieddatalibrary.Float(-6.602080212),
			Zaccel:                unifieddatalibrary.Float(-1.06310696),
			Zpos:                  unifieddatalibrary.Float(6167.831808),
			ZposAlt1:              unifieddatalibrary.Float(6165.55187),
			ZposAlt2:              unifieddatalibrary.Float(6165.55187),
			Zvel:                  unifieddatalibrary.Float(-3.356493869),
			ZvelAlt1:              unifieddatalibrary.Float(-3.365155181),
			ZvelAlt2:              unifieddatalibrary.Float(-3.365155181),
		}},
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestStateVectorGetWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := unifieddatalibrary.NewClient(
		option.WithBaseURL(baseURL),
		option.WithPassword("My Password"),
		option.WithUsername("My Username"),
	)
	_, err := client.StateVector.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.StateVectorGetParams{
			FirstResult: unifieddatalibrary.Int(0),
			MaxResults:  unifieddatalibrary.Int(0),
		},
	)
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestStateVectorQueryhelp(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := unifieddatalibrary.NewClient(
		option.WithBaseURL(baseURL),
		option.WithPassword("My Password"),
		option.WithUsername("My Username"),
	)
	_, err := client.StateVector.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestStateVectorTupleWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := unifieddatalibrary.NewClient(
		option.WithBaseURL(baseURL),
		option.WithPassword("My Password"),
		option.WithUsername("My Username"),
	)
	_, err := client.StateVector.Tuple(context.TODO(), unifieddatalibrary.StateVectorTupleParams{
		Columns:     "columns",
		Epoch:       time.Now(),
		FirstResult: unifieddatalibrary.Int(0),
		MaxResults:  unifieddatalibrary.Int(0),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestStateVectorUnvalidatedPublish(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := unifieddatalibrary.NewClient(
		option.WithBaseURL(baseURL),
		option.WithPassword("My Password"),
		option.WithUsername("My Username"),
	)
	err := client.StateVector.UnvalidatedPublish(context.TODO(), unifieddatalibrary.StateVectorUnvalidatedPublishParams{
		Body: []unifieddatalibrary.StateVectorIngestParam{{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.StateVectorIngestDataModeTest,
			Epoch:                 time.Now(),
			Source:                "Bluestaq",
			ActualOdSpan:          unifieddatalibrary.Float(3.5),
			Algorithm:             unifieddatalibrary.String("SAMPLE_ALGORITHM"),
			Alt1ReferenceFrame:    unifieddatalibrary.String("TEME"),
			Alt2ReferenceFrame:    unifieddatalibrary.String("EFG/TDR"),
			Area:                  unifieddatalibrary.Float(5.065),
			BDot:                  unifieddatalibrary.Float(1.23),
			CmOffset:              unifieddatalibrary.Float(1.23),
			Cov:                   []float64{1.1, 2.4, 3.8, 4.2, 5.5, 6},
			CovMethod:             unifieddatalibrary.String("CALCULATED"),
			CovReferenceFrame:     unifieddatalibrary.StateVectorIngestCovReferenceFrameJ2000,
			Descriptor:            unifieddatalibrary.String("descriptor"),
			DragArea:              unifieddatalibrary.Float(4.739),
			DragCoeff:             unifieddatalibrary.Float(0.0224391269775),
			DragModel:             unifieddatalibrary.String("JAC70"),
			Edr:                   unifieddatalibrary.Float(1.23),
			EqCov:                 []float64{1.1, 2.2},
			ErrorControl:          unifieddatalibrary.Float(1.23),
			FixedStep:             unifieddatalibrary.Bool(true),
			GeopotentialModel:     unifieddatalibrary.String("EGM-96"),
			Iau1980Terms:          unifieddatalibrary.Int(4),
			IDOrbitDetermination:  unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
			IDStateVector:         unifieddatalibrary.String("STATEVECTOR-ID"),
			IntegratorMode:        unifieddatalibrary.String("integratorMode"),
			InTrackThrust:         unifieddatalibrary.Bool(true),
			LastObEnd:             unifieddatalibrary.Time(time.Now()),
			LastObStart:           unifieddatalibrary.Time(time.Now()),
			LeapSecondTime:        unifieddatalibrary.Time(time.Now()),
			LunarSolar:            unifieddatalibrary.Bool(true),
			Mass:                  unifieddatalibrary.Float(164.5),
			ObsAvailable:          unifieddatalibrary.Int(376),
			ObsUsed:               unifieddatalibrary.Int(374),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigObjectID:          unifieddatalibrary.String("ORIGOBJECT-ID"),
			Partials:              unifieddatalibrary.String("ANALYTIC"),
			Pedigree:              unifieddatalibrary.String("CONJUNCTION"),
			PolarMotionX:          unifieddatalibrary.Float(1.23),
			PolarMotionY:          unifieddatalibrary.Float(1.23),
			PosUnc:                unifieddatalibrary.Float(0.333399744452),
			RawFileUri:            unifieddatalibrary.String("rawFileURI"),
			RecOdSpan:             unifieddatalibrary.Float(3.5),
			ReferenceFrame:        unifieddatalibrary.StateVectorIngestReferenceFrameJ2000,
			ResidualsAcc:          unifieddatalibrary.Float(99.5),
			RevNo:                 unifieddatalibrary.Int(7205),
			Rms:                   unifieddatalibrary.Float(0.991),
			SatNo:                 unifieddatalibrary.Int(12),
			SigmaPosUvw:           []float64{1.23, 4.56},
			SigmaVelUvw:           []float64{1.23, 4.56},
			SolarFluxApAvg:        unifieddatalibrary.Float(1.23),
			SolarFluxF10:          unifieddatalibrary.Float(1.23),
			SolarFluxF10Avg:       unifieddatalibrary.Float(1.23),
			SolarRadPress:         unifieddatalibrary.Bool(true),
			SolarRadPressCoeff:    unifieddatalibrary.Float(0.0244394),
			SolidEarthTides:       unifieddatalibrary.Bool(true),
			SourcedData:           []string{"DATA1", "DATA2"},
			SourcedDataTypes:      []string{"RADAR"},
			SrpArea:               unifieddatalibrary.Float(4.311),
			StepMode:              unifieddatalibrary.String("AUTO"),
			StepSize:              unifieddatalibrary.Float(1.23),
			StepSizeSelection:     unifieddatalibrary.String("AUTO"),
			Tags:                  []string{"TAG1", "TAG2"},
			TaiUtc:                unifieddatalibrary.Float(1.23),
			ThrustAccel:           unifieddatalibrary.Float(1.23),
			TracksAvail:           unifieddatalibrary.Int(163),
			TracksUsed:            unifieddatalibrary.Int(163),
			TransactionID:         unifieddatalibrary.String("transactionId"),
			Uct:                   unifieddatalibrary.Bool(true),
			Ut1Rate:               unifieddatalibrary.Float(1.23),
			Ut1Utc:                unifieddatalibrary.Float(1.23),
			VelUnc:                unifieddatalibrary.Float(0.000004),
			Xaccel:                unifieddatalibrary.Float(-2.12621392),
			Xpos:                  unifieddatalibrary.Float(-1118.577381),
			XposAlt1:              unifieddatalibrary.Float(-1145.688502),
			XposAlt2:              unifieddatalibrary.Float(-1456.915926),
			Xvel:                  unifieddatalibrary.Float(-4.25242784),
			XvelAlt1:              unifieddatalibrary.Float(-4.270832252),
			XvelAlt2:              unifieddatalibrary.Float(-1.219814294),
			Yaccel:                unifieddatalibrary.Float(2.645553717),
			Ypos:                  unifieddatalibrary.Float(3026.231084),
			YposAlt1:              unifieddatalibrary.Float(3020.729572),
			YposAlt2:              unifieddatalibrary.Float(-2883.540406),
			Yvel:                  unifieddatalibrary.Float(5.291107434),
			YvelAlt1:              unifieddatalibrary.Float(5.27074276),
			YvelAlt2:              unifieddatalibrary.Float(-6.602080212),
			Zaccel:                unifieddatalibrary.Float(-1.06310696),
			Zpos:                  unifieddatalibrary.Float(6167.831808),
			ZposAlt1:              unifieddatalibrary.Float(6165.55187),
			ZposAlt2:              unifieddatalibrary.Float(6165.55187),
			Zvel:                  unifieddatalibrary.Float(-3.356493869),
			ZvelAlt1:              unifieddatalibrary.Float(-3.365155181),
			ZvelAlt2:              unifieddatalibrary.Float(-3.365155181),
		}},
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
