// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"
	"time"

	"github.com/Bluestaq/udl-golang-sdk"
	"github.com/Bluestaq/udl-golang-sdk/internal/testutil"
	"github.com/Bluestaq/udl-golang-sdk/option"
)

func TestConjunctionGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Conjunctions.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.ConjunctionGetParams{
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

func TestConjunctionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Conjunctions.List(context.TODO(), unifieddatalibrary.ConjunctionListParams{
		Tca:         time.Now(),
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

func TestConjunctionCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Conjunctions.Count(context.TODO(), unifieddatalibrary.ConjunctionCountParams{
		Tca:         time.Now(),
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

func TestConjunctionNewUdlWithOptionalParams(t *testing.T) {
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
	err := client.Conjunctions.NewUdl(context.TODO(), unifieddatalibrary.ConjunctionNewUdlParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.ConjunctionNewUdlParamsDataModeTest,
		Source:                "Bluestaq",
		Tca:                   time.Now(),
		ConvertPosVel:         unifieddatalibrary.Bool(true),
		ID:                    unifieddatalibrary.String("CONJUNCTION-ID"),
		Ccir:                  unifieddatalibrary.String("Example notes"),
		CdAoM1:                unifieddatalibrary.Float(0.016386),
		CdAoM2:                unifieddatalibrary.Float(0.016386),
		CollisionProb:         unifieddatalibrary.Float(0.5),
		CollisionProbMethod:   unifieddatalibrary.String("FOSTER-1992"),
		Comments:              unifieddatalibrary.String("Example notes"),
		ConcernNotes:          unifieddatalibrary.String("Example notes"),
		CrAoM1:                unifieddatalibrary.Float(0.013814),
		CrAoM2:                unifieddatalibrary.Float(0.013814),
		Descriptor:            unifieddatalibrary.String("sample_descriptor here"),
		EphemName1:            unifieddatalibrary.String("MEME_SPCFT_ABC_2180000_ops_nomnvr_unclassified.oem"),
		EphemName2:            unifieddatalibrary.String("MEME_SPCFT_DEF_2170000_ops_nomnvr_unclassified.txt"),
		EsId1:                 unifieddatalibrary.String("a2ae2356-6d83-4e4b-896d-ddd1958800fa"),
		EsId2:                 unifieddatalibrary.String("6fa31433-8beb-4b9b-8bf9-326dbd041c3f"),
		EventID:               unifieddatalibrary.String("some.user"),
		IDStateVector1:        unifieddatalibrary.String("REF-STATEVECTOR1-ID"),
		IDStateVector2:        unifieddatalibrary.String("REF-STATEVECTOR2-ID"),
		LargeCovWarning:       unifieddatalibrary.Bool(false),
		LargeRelPosWarning:    unifieddatalibrary.Bool(false),
		LastObTime1:           unifieddatalibrary.Time(time.Now()),
		LastObTime2:           unifieddatalibrary.Time(time.Now()),
		MessageFor:            unifieddatalibrary.String("Message for space craft A"),
		MessageID:             unifieddatalibrary.String("MESSAGE-ID"),
		MissDistance:          unifieddatalibrary.Float(1.1),
		OrigIDOnOrbit1:        unifieddatalibrary.String("ORIGONORBIT1-ID"),
		OrigIDOnOrbit2:        unifieddatalibrary.String("ORIGONORBIT2-ID"),
		Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		Originator:            unifieddatalibrary.String("some.user"),
		OwnerContacted:        unifieddatalibrary.Bool(false),
		PenetrationLevelSigma: unifieddatalibrary.Float(1.1),
		RawFileUri:            unifieddatalibrary.String("Example link"),
		RelPosN:               unifieddatalibrary.Float(1.1),
		RelPosR:               unifieddatalibrary.Float(1.1),
		RelPosT:               unifieddatalibrary.Float(1.1),
		RelVelMag:             unifieddatalibrary.Float(1.1),
		RelVelN:               unifieddatalibrary.Float(1.1),
		RelVelR:               unifieddatalibrary.Float(1.1),
		RelVelT:               unifieddatalibrary.Float(1.1),
		SatNo1:                unifieddatalibrary.Int(1),
		SatNo2:                unifieddatalibrary.Int(2),
		ScreenEntryTime:       unifieddatalibrary.Time(time.Now()),
		ScreenExitTime:        unifieddatalibrary.Time(time.Now()),
		ScreenVolumeX:         unifieddatalibrary.Float(1.1),
		ScreenVolumeY:         unifieddatalibrary.Float(1.1),
		ScreenVolumeZ:         unifieddatalibrary.Float(1.1),
		SmallCovWarning:       unifieddatalibrary.Bool(false),
		SmallRelVelWarning:    unifieddatalibrary.Bool(false),
		StateDeptNotified:     unifieddatalibrary.Bool(false),
		StateVector1: unifieddatalibrary.ConjunctionNewUdlParamsStateVector1{
			ClassificationMarking: "U",
			DataMode:              "TEST",
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
			CovReferenceFrame:     "J2000",
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
			ReferenceFrame:        "J2000",
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
		StateVector2: unifieddatalibrary.ConjunctionNewUdlParamsStateVector2{
			ClassificationMarking: "U",
			DataMode:              "TEST",
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
			CovReferenceFrame:     "J2000",
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
			ReferenceFrame:        "J2000",
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
		Tags:          []string{"PROVIDER_TAG1", "PROVIDER_TAG1"},
		ThrustAccel1:  unifieddatalibrary.Float(0.033814),
		ThrustAccel2:  unifieddatalibrary.Float(0.033814),
		TransactionID: unifieddatalibrary.String("TRANSACTION-ID"),
		Type:          unifieddatalibrary.String("CONJUNCTION"),
		UvwWarn:       unifieddatalibrary.Bool(false),
		VolEntryTime:  unifieddatalibrary.Time(time.Now()),
		VolExitTime:   unifieddatalibrary.Time(time.Now()),
		VolShape:      unifieddatalibrary.String("ELLIPSOID"),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestConjunctionNewBulk(t *testing.T) {
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
	err := client.Conjunctions.NewBulk(context.TODO(), unifieddatalibrary.ConjunctionNewBulkParams{
		Body: []unifieddatalibrary.ConjunctionNewBulkParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			Source:                "Bluestaq",
			Tca:                   time.Now(),
			ID:                    unifieddatalibrary.String("CONJUNCTION-ID"),
			Ccir:                  unifieddatalibrary.String("Example notes"),
			CdAoM1:                unifieddatalibrary.Float(0.016386),
			CdAoM2:                unifieddatalibrary.Float(0.016386),
			CollisionProb:         unifieddatalibrary.Float(0.5),
			CollisionProbMethod:   unifieddatalibrary.String("FOSTER-1992"),
			Comments:              unifieddatalibrary.String("Example notes"),
			ConcernNotes:          unifieddatalibrary.String("Example notes"),
			CrAoM1:                unifieddatalibrary.Float(0.013814),
			CrAoM2:                unifieddatalibrary.Float(0.013814),
			Descriptor:            unifieddatalibrary.String("sample_descriptor here"),
			EphemName1:            unifieddatalibrary.String("MEME_SPCFT_ABC_2180000_ops_nomnvr_unclassified.oem"),
			EphemName2:            unifieddatalibrary.String("MEME_SPCFT_DEF_2170000_ops_nomnvr_unclassified.txt"),
			EsId1:                 unifieddatalibrary.String("a2ae2356-6d83-4e4b-896d-ddd1958800fa"),
			EsId2:                 unifieddatalibrary.String("6fa31433-8beb-4b9b-8bf9-326dbd041c3f"),
			EventID:               unifieddatalibrary.String("some.user"),
			IDStateVector1:        unifieddatalibrary.String("REF-STATEVECTOR1-ID"),
			IDStateVector2:        unifieddatalibrary.String("REF-STATEVECTOR2-ID"),
			LargeCovWarning:       unifieddatalibrary.Bool(false),
			LargeRelPosWarning:    unifieddatalibrary.Bool(false),
			LastObTime1:           unifieddatalibrary.Time(time.Now()),
			LastObTime2:           unifieddatalibrary.Time(time.Now()),
			MessageFor:            unifieddatalibrary.String("Message for space craft A"),
			MessageID:             unifieddatalibrary.String("MESSAGE-ID"),
			MissDistance:          unifieddatalibrary.Float(1.1),
			OrigIDOnOrbit1:        unifieddatalibrary.String("ORIGONORBIT1-ID"),
			OrigIDOnOrbit2:        unifieddatalibrary.String("ORIGONORBIT2-ID"),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			Originator:            unifieddatalibrary.String("some.user"),
			OwnerContacted:        unifieddatalibrary.Bool(false),
			PenetrationLevelSigma: unifieddatalibrary.Float(1.1),
			RawFileUri:            unifieddatalibrary.String("Example link"),
			RelPosN:               unifieddatalibrary.Float(1.1),
			RelPosR:               unifieddatalibrary.Float(1.1),
			RelPosT:               unifieddatalibrary.Float(1.1),
			RelVelMag:             unifieddatalibrary.Float(1.1),
			RelVelN:               unifieddatalibrary.Float(1.1),
			RelVelR:               unifieddatalibrary.Float(1.1),
			RelVelT:               unifieddatalibrary.Float(1.1),
			SatNo1:                unifieddatalibrary.Int(1),
			SatNo2:                unifieddatalibrary.Int(2),
			ScreenEntryTime:       unifieddatalibrary.Time(time.Now()),
			ScreenExitTime:        unifieddatalibrary.Time(time.Now()),
			ScreenVolumeX:         unifieddatalibrary.Float(1.1),
			ScreenVolumeY:         unifieddatalibrary.Float(1.1),
			ScreenVolumeZ:         unifieddatalibrary.Float(1.1),
			SmallCovWarning:       unifieddatalibrary.Bool(false),
			SmallRelVelWarning:    unifieddatalibrary.Bool(false),
			StateDeptNotified:     unifieddatalibrary.Bool(false),
			StateVector1: unifieddatalibrary.ConjunctionNewBulkParamsBodyStateVector1{
				ClassificationMarking: "U",
				DataMode:              "TEST",
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
				CovReferenceFrame:     "J2000",
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
				ReferenceFrame:        "J2000",
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
			StateVector2: unifieddatalibrary.ConjunctionNewBulkParamsBodyStateVector2{
				ClassificationMarking: "U",
				DataMode:              "TEST",
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
				CovReferenceFrame:     "J2000",
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
				ReferenceFrame:        "J2000",
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
			Tags:          []string{"PROVIDER_TAG1", "PROVIDER_TAG1"},
			ThrustAccel1:  unifieddatalibrary.Float(0.033814),
			ThrustAccel2:  unifieddatalibrary.Float(0.033814),
			TransactionID: unifieddatalibrary.String("TRANSACTION-ID"),
			Type:          unifieddatalibrary.String("CONJUNCTION"),
			UvwWarn:       unifieddatalibrary.Bool(false),
			VolEntryTime:  unifieddatalibrary.Time(time.Now()),
			VolExitTime:   unifieddatalibrary.Time(time.Now()),
			VolShape:      unifieddatalibrary.String("ELLIPSOID"),
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

func TestConjunctionGetHistoryWithOptionalParams(t *testing.T) {
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
	_, err := client.Conjunctions.GetHistory(context.TODO(), unifieddatalibrary.ConjunctionGetHistoryParams{
		Tca:         time.Now(),
		Columns:     unifieddatalibrary.String("columns"),
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

func TestConjunctionQueryhelp(t *testing.T) {
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
	_, err := client.Conjunctions.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestConjunctionTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Conjunctions.Tuple(context.TODO(), unifieddatalibrary.ConjunctionTupleParams{
		Columns:     "columns",
		Tca:         time.Now(),
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

func TestConjunctionUnvalidatedPublish(t *testing.T) {
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
	err := client.Conjunctions.UnvalidatedPublish(context.TODO(), unifieddatalibrary.ConjunctionUnvalidatedPublishParams{
		Body: []unifieddatalibrary.ConjunctionUnvalidatedPublishParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			Source:                "Bluestaq",
			Tca:                   time.Now(),
			ID:                    unifieddatalibrary.String("CONJUNCTION-ID"),
			Ccir:                  unifieddatalibrary.String("Example notes"),
			CdAoM1:                unifieddatalibrary.Float(0.016386),
			CdAoM2:                unifieddatalibrary.Float(0.016386),
			CollisionProb:         unifieddatalibrary.Float(0.5),
			CollisionProbMethod:   unifieddatalibrary.String("FOSTER-1992"),
			Comments:              unifieddatalibrary.String("Example notes"),
			ConcernNotes:          unifieddatalibrary.String("Example notes"),
			CrAoM1:                unifieddatalibrary.Float(0.013814),
			CrAoM2:                unifieddatalibrary.Float(0.013814),
			Descriptor:            unifieddatalibrary.String("sample_descriptor here"),
			EphemName1:            unifieddatalibrary.String("MEME_SPCFT_ABC_2180000_ops_nomnvr_unclassified.oem"),
			EphemName2:            unifieddatalibrary.String("MEME_SPCFT_DEF_2170000_ops_nomnvr_unclassified.txt"),
			EsId1:                 unifieddatalibrary.String("a2ae2356-6d83-4e4b-896d-ddd1958800fa"),
			EsId2:                 unifieddatalibrary.String("6fa31433-8beb-4b9b-8bf9-326dbd041c3f"),
			EventID:               unifieddatalibrary.String("some.user"),
			IDStateVector1:        unifieddatalibrary.String("REF-STATEVECTOR1-ID"),
			IDStateVector2:        unifieddatalibrary.String("REF-STATEVECTOR2-ID"),
			LargeCovWarning:       unifieddatalibrary.Bool(false),
			LargeRelPosWarning:    unifieddatalibrary.Bool(false),
			LastObTime1:           unifieddatalibrary.Time(time.Now()),
			LastObTime2:           unifieddatalibrary.Time(time.Now()),
			MessageFor:            unifieddatalibrary.String("Message for space craft A"),
			MessageID:             unifieddatalibrary.String("MESSAGE-ID"),
			MissDistance:          unifieddatalibrary.Float(1.1),
			OrigIDOnOrbit1:        unifieddatalibrary.String("ORIGONORBIT1-ID"),
			OrigIDOnOrbit2:        unifieddatalibrary.String("ORIGONORBIT2-ID"),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			Originator:            unifieddatalibrary.String("some.user"),
			OwnerContacted:        unifieddatalibrary.Bool(false),
			PenetrationLevelSigma: unifieddatalibrary.Float(1.1),
			RawFileUri:            unifieddatalibrary.String("Example link"),
			RelPosN:               unifieddatalibrary.Float(1.1),
			RelPosR:               unifieddatalibrary.Float(1.1),
			RelPosT:               unifieddatalibrary.Float(1.1),
			RelVelMag:             unifieddatalibrary.Float(1.1),
			RelVelN:               unifieddatalibrary.Float(1.1),
			RelVelR:               unifieddatalibrary.Float(1.1),
			RelVelT:               unifieddatalibrary.Float(1.1),
			SatNo1:                unifieddatalibrary.Int(1),
			SatNo2:                unifieddatalibrary.Int(2),
			ScreenEntryTime:       unifieddatalibrary.Time(time.Now()),
			ScreenExitTime:        unifieddatalibrary.Time(time.Now()),
			ScreenVolumeX:         unifieddatalibrary.Float(1.1),
			ScreenVolumeY:         unifieddatalibrary.Float(1.1),
			ScreenVolumeZ:         unifieddatalibrary.Float(1.1),
			SmallCovWarning:       unifieddatalibrary.Bool(false),
			SmallRelVelWarning:    unifieddatalibrary.Bool(false),
			StateDeptNotified:     unifieddatalibrary.Bool(false),
			StateVector1: unifieddatalibrary.ConjunctionUnvalidatedPublishParamsBodyStateVector1{
				ClassificationMarking: "U",
				DataMode:              "TEST",
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
				CovReferenceFrame:     "J2000",
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
				ReferenceFrame:        "J2000",
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
			StateVector2: unifieddatalibrary.ConjunctionUnvalidatedPublishParamsBodyStateVector2{
				ClassificationMarking: "U",
				DataMode:              "TEST",
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
				CovReferenceFrame:     "J2000",
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
				ReferenceFrame:        "J2000",
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
			Tags:          []string{"PROVIDER_TAG1", "PROVIDER_TAG1"},
			ThrustAccel1:  unifieddatalibrary.Float(0.033814),
			ThrustAccel2:  unifieddatalibrary.Float(0.033814),
			TransactionID: unifieddatalibrary.String("TRANSACTION-ID"),
			Type:          unifieddatalibrary.String("CONJUNCTION"),
			UvwWarn:       unifieddatalibrary.Bool(false),
			VolEntryTime:  unifieddatalibrary.Time(time.Now()),
			VolExitTime:   unifieddatalibrary.Time(time.Now()),
			VolShape:      unifieddatalibrary.String("ELLIPSOID"),
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

func TestConjunctionUploadConjunctionDataMessageWithOptionalParams(t *testing.T) {
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
	err := client.Conjunctions.UploadConjunctionDataMessage(
		context.TODO(),
		io.Reader(bytes.NewBuffer([]byte("some file contents"))),
		unifieddatalibrary.ConjunctionUploadConjunctionDataMessageParams{
			Classification: "classification",
			DataMode:       unifieddatalibrary.ConjunctionUploadConjunctionDataMessageParamsDataModeReal,
			Filename:       "filename",
			Source:         "source",
			Tags:           unifieddatalibrary.String("tags"),
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
