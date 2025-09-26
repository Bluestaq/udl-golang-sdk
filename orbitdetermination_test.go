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

func TestOrbitdeterminationNewWithOptionalParams(t *testing.T) {
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
	err := client.Orbitdetermination.New(context.TODO(), unifieddatalibrary.OrbitdeterminationNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.OrbitdeterminationNewParamsDataModeTest,
		EndTime:               time.Now(),
		InitialOd:             false,
		Method:                "BLS",
		Source:                "Bluestaq",
		StartTime:             time.Now(),
		ID:                    unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
		AcceptedObIDs:         []string{"EOOBSERVATION-ID1", "RADAROBSERVATION-ID1"},
		AcceptedObTyps:        []string{"EO", "RADAR"},
		AgomEst:               unifieddatalibrary.Bool(false),
		AgomModel:             unifieddatalibrary.String("RandomWalk"),
		AprioriElset: unifieddatalibrary.OrbitdeterminationNewParamsAprioriElset{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			Epoch:                 time.Now(),
			Source:                "Bluestaq",
			Agom:                  unifieddatalibrary.Float(0.0126),
			Algorithm:             unifieddatalibrary.String("Example algorithm"),
			Apogee:                unifieddatalibrary.Float(1.1),
			ArgOfPerigee:          unifieddatalibrary.Float(1.1),
			BallisticCoeff:        unifieddatalibrary.Float(0.00815),
			BStar:                 unifieddatalibrary.Float(1.1),
			Descriptor:            unifieddatalibrary.String("Example description"),
			Eccentricity:          unifieddatalibrary.Float(0.333),
			EphemType:             unifieddatalibrary.Int(1),
			IDElset:               unifieddatalibrary.String("ELSET-ID"),
			IDOrbitDetermination:  unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
			Inclination:           unifieddatalibrary.Float(45.1),
			MeanAnomaly:           unifieddatalibrary.Float(179.1),
			MeanMotion:            unifieddatalibrary.Float(1.1),
			MeanMotionDDot:        unifieddatalibrary.Float(1.1),
			MeanMotionDot:         unifieddatalibrary.Float(1.1),
			Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigObjectID:          unifieddatalibrary.String("ORIGOBJECT-ID"),
			Perigee:               unifieddatalibrary.Float(1.1),
			Period:                unifieddatalibrary.Float(1.1),
			Raan:                  unifieddatalibrary.Float(1.1),
			RawFileUri:            unifieddatalibrary.String("Example URI"),
			RevNo:                 unifieddatalibrary.Int(111),
			SatNo:                 unifieddatalibrary.Int(12),
			SemiMajorAxis:         unifieddatalibrary.Float(1.1),
			SourcedData:           []string{"OBSERVATION_UUID1", "OBSERVATION_UUID2"},
			SourcedDataTypes:      []string{"RADAR", "RF"},
			Tags:                  []string{"PROVIDER_TAG1", "PROVIDER_TAG2"},
			TransactionID:         unifieddatalibrary.String("TRANSACTION-ID"),
			Uct:                   unifieddatalibrary.Bool(false),
		},
		AprioriIDElset:       unifieddatalibrary.String("80e544b7-6a17-4554-8abf-7301e98f8e5d"),
		AprioriIDStateVector: unifieddatalibrary.String("6e291992-8ae3-4592-bb0f-055715bf4803"),
		AprioriStateVector: unifieddatalibrary.OrbitdeterminationNewParamsAprioriStateVector{
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
			MsgTs:                 unifieddatalibrary.Time(time.Now()),
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
		BallisticCoeffEst:      unifieddatalibrary.Bool(false),
		BallisticCoeffModel:    unifieddatalibrary.String("GaussMarkov"),
		BestPassWrms:           unifieddatalibrary.Float(0.975),
		Edr:                    unifieddatalibrary.Float(1.23),
		EffectiveFrom:          unifieddatalibrary.Time(time.Now()),
		EffectiveUntil:         unifieddatalibrary.Time(time.Now()),
		ErrorGrowthRate:        unifieddatalibrary.Float(1.23),
		FirstPassWrms:          unifieddatalibrary.Float(0.985),
		FitSpan:                unifieddatalibrary.Float(0.6),
		LastObEnd:              unifieddatalibrary.Time(time.Now()),
		LastObStart:            unifieddatalibrary.Time(time.Now()),
		MethodSource:           unifieddatalibrary.String("ASW"),
		NumIterations:          unifieddatalibrary.Int(8),
		Origin:                 unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		OrigObjectID:           unifieddatalibrary.String("ORIGOBJECT-ID"),
		PreviousWrms:           unifieddatalibrary.Float(1.02),
		RejectedObIDs:          []string{"DIFFOFARRIVAL-ID2", "RFOBSERVATION-ID2"},
		RejectedObTyps:         []string{"DOA", "RF"},
		RmsConvergenceCriteria: unifieddatalibrary.Float(0.001),
		SatNo:                  unifieddatalibrary.Int(54741),
		SensorIDs:              []string{"SENSOR-ID1", "SENSOR-ID2"},
		TimeSpan:               unifieddatalibrary.Float(3.5),
		Wrms:                   unifieddatalibrary.Float(0.991),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOrbitdeterminationListWithOptionalParams(t *testing.T) {
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
	_, err := client.Orbitdetermination.List(context.TODO(), unifieddatalibrary.OrbitdeterminationListParams{
		FirstResult: unifieddatalibrary.Int(0),
		IDOnOrbit:   unifieddatalibrary.String("idOnOrbit"),
		MaxResults:  unifieddatalibrary.Int(0),
		StartTime:   unifieddatalibrary.Time(time.Now()),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOrbitdeterminationCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Orbitdetermination.Count(context.TODO(), unifieddatalibrary.OrbitdeterminationCountParams{
		FirstResult: unifieddatalibrary.Int(0),
		IDOnOrbit:   unifieddatalibrary.String("idOnOrbit"),
		MaxResults:  unifieddatalibrary.Int(0),
		StartTime:   unifieddatalibrary.Time(time.Now()),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOrbitdeterminationNewBulk(t *testing.T) {
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
	err := client.Orbitdetermination.NewBulk(context.TODO(), unifieddatalibrary.OrbitdeterminationNewBulkParams{
		Body: []unifieddatalibrary.OrbitdeterminationNewBulkParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			EndTime:               time.Now(),
			InitialOd:             false,
			Method:                "BLS",
			Source:                "Bluestaq",
			StartTime:             time.Now(),
			ID:                    unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
			AcceptedObIDs:         []string{"EOOBSERVATION-ID1", "RADAROBSERVATION-ID1"},
			AcceptedObTyps:        []string{"EO", "RADAR"},
			AgomEst:               unifieddatalibrary.Bool(false),
			AgomModel:             unifieddatalibrary.String("RandomWalk"),
			AprioriElset: unifieddatalibrary.OrbitdeterminationNewBulkParamsBodyAprioriElset{
				ClassificationMarking: "U",
				DataMode:              "TEST",
				Epoch:                 time.Now(),
				Source:                "Bluestaq",
				Agom:                  unifieddatalibrary.Float(0.0126),
				Algorithm:             unifieddatalibrary.String("Example algorithm"),
				Apogee:                unifieddatalibrary.Float(1.1),
				ArgOfPerigee:          unifieddatalibrary.Float(1.1),
				BallisticCoeff:        unifieddatalibrary.Float(0.00815),
				BStar:                 unifieddatalibrary.Float(1.1),
				Descriptor:            unifieddatalibrary.String("Example description"),
				Eccentricity:          unifieddatalibrary.Float(0.333),
				EphemType:             unifieddatalibrary.Int(1),
				IDElset:               unifieddatalibrary.String("ELSET-ID"),
				IDOrbitDetermination:  unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
				Inclination:           unifieddatalibrary.Float(45.1),
				MeanAnomaly:           unifieddatalibrary.Float(179.1),
				MeanMotion:            unifieddatalibrary.Float(1.1),
				MeanMotionDDot:        unifieddatalibrary.Float(1.1),
				MeanMotionDot:         unifieddatalibrary.Float(1.1),
				Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
				OrigObjectID:          unifieddatalibrary.String("ORIGOBJECT-ID"),
				Perigee:               unifieddatalibrary.Float(1.1),
				Period:                unifieddatalibrary.Float(1.1),
				Raan:                  unifieddatalibrary.Float(1.1),
				RawFileUri:            unifieddatalibrary.String("Example URI"),
				RevNo:                 unifieddatalibrary.Int(111),
				SatNo:                 unifieddatalibrary.Int(12),
				SemiMajorAxis:         unifieddatalibrary.Float(1.1),
				SourcedData:           []string{"OBSERVATION_UUID1", "OBSERVATION_UUID2"},
				SourcedDataTypes:      []string{"RADAR", "RF"},
				Tags:                  []string{"PROVIDER_TAG1", "PROVIDER_TAG2"},
				TransactionID:         unifieddatalibrary.String("TRANSACTION-ID"),
				Uct:                   unifieddatalibrary.Bool(false),
			},
			AprioriIDElset:       unifieddatalibrary.String("80e544b7-6a17-4554-8abf-7301e98f8e5d"),
			AprioriIDStateVector: unifieddatalibrary.String("6e291992-8ae3-4592-bb0f-055715bf4803"),
			AprioriStateVector: unifieddatalibrary.OrbitdeterminationNewBulkParamsBodyAprioriStateVector{
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
				MsgTs:                 unifieddatalibrary.Time(time.Now()),
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
			BallisticCoeffEst:      unifieddatalibrary.Bool(false),
			BallisticCoeffModel:    unifieddatalibrary.String("GaussMarkov"),
			BestPassWrms:           unifieddatalibrary.Float(0.975),
			Edr:                    unifieddatalibrary.Float(1.23),
			EffectiveFrom:          unifieddatalibrary.Time(time.Now()),
			EffectiveUntil:         unifieddatalibrary.Time(time.Now()),
			ErrorGrowthRate:        unifieddatalibrary.Float(1.23),
			FirstPassWrms:          unifieddatalibrary.Float(0.985),
			FitSpan:                unifieddatalibrary.Float(0.6),
			LastObEnd:              unifieddatalibrary.Time(time.Now()),
			LastObStart:            unifieddatalibrary.Time(time.Now()),
			MethodSource:           unifieddatalibrary.String("ASW"),
			NumIterations:          unifieddatalibrary.Int(8),
			Origin:                 unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigObjectID:           unifieddatalibrary.String("ORIGOBJECT-ID"),
			PreviousWrms:           unifieddatalibrary.Float(1.02),
			RejectedObIDs:          []string{"DIFFOFARRIVAL-ID2", "RFOBSERVATION-ID2"},
			RejectedObTyps:         []string{"DOA", "RF"},
			RmsConvergenceCriteria: unifieddatalibrary.Float(0.001),
			SatNo:                  unifieddatalibrary.Int(54741),
			SensorIDs:              []string{"SENSOR-ID1", "SENSOR-ID2"},
			TimeSpan:               unifieddatalibrary.Float(3.5),
			Wrms:                   unifieddatalibrary.Float(0.991),
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

func TestOrbitdeterminationGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Orbitdetermination.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.OrbitdeterminationGetParams{
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

func TestOrbitdeterminationQueryhelp(t *testing.T) {
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
	_, err := client.Orbitdetermination.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOrbitdeterminationTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Orbitdetermination.Tuple(context.TODO(), unifieddatalibrary.OrbitdeterminationTupleParams{
		Columns:     "columns",
		FirstResult: unifieddatalibrary.Int(0),
		IDOnOrbit:   unifieddatalibrary.String("idOnOrbit"),
		MaxResults:  unifieddatalibrary.Int(0),
		StartTime:   unifieddatalibrary.Time(time.Now()),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOrbitdeterminationUnvalidatedPublish(t *testing.T) {
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
	err := client.Orbitdetermination.UnvalidatedPublish(context.TODO(), unifieddatalibrary.OrbitdeterminationUnvalidatedPublishParams{
		Body: []unifieddatalibrary.OrbitdeterminationUnvalidatedPublishParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			EndTime:               time.Now(),
			InitialOd:             false,
			Method:                "BLS",
			Source:                "Bluestaq",
			StartTime:             time.Now(),
			ID:                    unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
			AcceptedObIDs:         []string{"EOOBSERVATION-ID1", "RADAROBSERVATION-ID1"},
			AcceptedObTyps:        []string{"EO", "RADAR"},
			AgomEst:               unifieddatalibrary.Bool(false),
			AgomModel:             unifieddatalibrary.String("RandomWalk"),
			AprioriElset: unifieddatalibrary.OrbitdeterminationUnvalidatedPublishParamsBodyAprioriElset{
				ClassificationMarking: "U",
				DataMode:              "TEST",
				Epoch:                 time.Now(),
				Source:                "Bluestaq",
				Agom:                  unifieddatalibrary.Float(0.0126),
				Algorithm:             unifieddatalibrary.String("Example algorithm"),
				Apogee:                unifieddatalibrary.Float(1.1),
				ArgOfPerigee:          unifieddatalibrary.Float(1.1),
				BallisticCoeff:        unifieddatalibrary.Float(0.00815),
				BStar:                 unifieddatalibrary.Float(1.1),
				Descriptor:            unifieddatalibrary.String("Example description"),
				Eccentricity:          unifieddatalibrary.Float(0.333),
				EphemType:             unifieddatalibrary.Int(1),
				IDElset:               unifieddatalibrary.String("ELSET-ID"),
				IDOrbitDetermination:  unifieddatalibrary.String("026dd511-8ba5-47d3-9909-836149f87686"),
				Inclination:           unifieddatalibrary.Float(45.1),
				MeanAnomaly:           unifieddatalibrary.Float(179.1),
				MeanMotion:            unifieddatalibrary.Float(1.1),
				MeanMotionDDot:        unifieddatalibrary.Float(1.1),
				MeanMotionDot:         unifieddatalibrary.Float(1.1),
				Origin:                unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
				OrigObjectID:          unifieddatalibrary.String("ORIGOBJECT-ID"),
				Perigee:               unifieddatalibrary.Float(1.1),
				Period:                unifieddatalibrary.Float(1.1),
				Raan:                  unifieddatalibrary.Float(1.1),
				RawFileUri:            unifieddatalibrary.String("Example URI"),
				RevNo:                 unifieddatalibrary.Int(111),
				SatNo:                 unifieddatalibrary.Int(12),
				SemiMajorAxis:         unifieddatalibrary.Float(1.1),
				SourcedData:           []string{"OBSERVATION_UUID1", "OBSERVATION_UUID2"},
				SourcedDataTypes:      []string{"RADAR", "RF"},
				Tags:                  []string{"PROVIDER_TAG1", "PROVIDER_TAG2"},
				TransactionID:         unifieddatalibrary.String("TRANSACTION-ID"),
				Uct:                   unifieddatalibrary.Bool(false),
			},
			AprioriIDElset:       unifieddatalibrary.String("80e544b7-6a17-4554-8abf-7301e98f8e5d"),
			AprioriIDStateVector: unifieddatalibrary.String("6e291992-8ae3-4592-bb0f-055715bf4803"),
			AprioriStateVector: unifieddatalibrary.OrbitdeterminationUnvalidatedPublishParamsBodyAprioriStateVector{
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
				MsgTs:                 unifieddatalibrary.Time(time.Now()),
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
			BallisticCoeffEst:      unifieddatalibrary.Bool(false),
			BallisticCoeffModel:    unifieddatalibrary.String("GaussMarkov"),
			BestPassWrms:           unifieddatalibrary.Float(0.975),
			Edr:                    unifieddatalibrary.Float(1.23),
			EffectiveFrom:          unifieddatalibrary.Time(time.Now()),
			EffectiveUntil:         unifieddatalibrary.Time(time.Now()),
			ErrorGrowthRate:        unifieddatalibrary.Float(1.23),
			FirstPassWrms:          unifieddatalibrary.Float(0.985),
			FitSpan:                unifieddatalibrary.Float(0.6),
			LastObEnd:              unifieddatalibrary.Time(time.Now()),
			LastObStart:            unifieddatalibrary.Time(time.Now()),
			MethodSource:           unifieddatalibrary.String("ASW"),
			NumIterations:          unifieddatalibrary.Int(8),
			Origin:                 unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigObjectID:           unifieddatalibrary.String("ORIGOBJECT-ID"),
			PreviousWrms:           unifieddatalibrary.Float(1.02),
			RejectedObIDs:          []string{"DIFFOFARRIVAL-ID2", "RFOBSERVATION-ID2"},
			RejectedObTyps:         []string{"DOA", "RF"},
			RmsConvergenceCriteria: unifieddatalibrary.Float(0.001),
			SatNo:                  unifieddatalibrary.Int(54741),
			SensorIDs:              []string{"SENSOR-ID1", "SENSOR-ID2"},
			TimeSpan:               unifieddatalibrary.Float(3.5),
			Wrms:                   unifieddatalibrary.Float(0.991),
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
