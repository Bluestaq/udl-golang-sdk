// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/unifieddatalibrary-go"
	"github.com/stainless-sdks/unifieddatalibrary-go/internal/testutil"
	"github.com/stainless-sdks/unifieddatalibrary-go/option"
)

func TestSensorplanNewWithOptionalParams(t *testing.T) {
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
	err := client.Sensorplan.New(context.TODO(), unifieddatalibrary.SensorplanNewParams{
		ClassificationMarking: "U",
		DataMode:              unifieddatalibrary.SensorplanNewParamsDataModeTest,
		RecType:               "COLLECT",
		Source:                "Bluestaq",
		StartTime:             time.Now(),
		Type:                  "PLAN",
		ID:                    unifieddatalibrary.String("SENSORPLAN-ID"),
		CollectRequests: []unifieddatalibrary.SensorplanNewParamsCollectRequest{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			Source:                "Bluestaq",
			StartTime:             time.Now(),
			Type:                  "DWELL",
			ID:                    unifieddatalibrary.String("COLLECTREQUEST-ID"),
			Alt:                   unifieddatalibrary.Float(1.1),
			ArgOfPerigee:          unifieddatalibrary.Float(1.1),
			Az:                    unifieddatalibrary.Float(1.1),
			Customer:              unifieddatalibrary.String("Bluestaq"),
			Dec:                   unifieddatalibrary.Float(1.1),
			Duration:              unifieddatalibrary.Int(11),
			DwellID:               unifieddatalibrary.String("DWELL-ID"),
			Eccentricity:          unifieddatalibrary.Float(1.1),
			El:                    unifieddatalibrary.Float(1.1),
			Elset: unifieddatalibrary.SensorplanNewParamsCollectRequestElset{
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
			EndTime:         unifieddatalibrary.Time(time.Now()),
			Epoch:           unifieddatalibrary.Time(time.Now()),
			EsID:            unifieddatalibrary.String("ES-ID"),
			ExtentAz:        unifieddatalibrary.Float(1.1),
			ExtentEl:        unifieddatalibrary.Float(1.1),
			ExtentRange:     unifieddatalibrary.Float(1.1),
			ExternalID:      unifieddatalibrary.String("EXTERNAL-ID"),
			FrameRate:       unifieddatalibrary.Float(1.1),
			Freq:            unifieddatalibrary.Float(1.1),
			FreqMax:         unifieddatalibrary.Float(1.1),
			FreqMin:         unifieddatalibrary.Float(1.1),
			IDElset:         unifieddatalibrary.String("REF-ELSET-ID"),
			IDManifold:      unifieddatalibrary.String("REF-MANIFOLD-ID"),
			IDParentReq:     unifieddatalibrary.String("da98671b-34db-47bf-8c8d-7c668b92c800"),
			IDPlan:          unifieddatalibrary.String("REF-PLAN-ID"),
			IDSensor:        unifieddatalibrary.String("REF-SENSOR-ID"),
			IDStateVector:   unifieddatalibrary.String("STATEVECTOR-ID"),
			Inclination:     unifieddatalibrary.Float(1.1),
			IntegrationTime: unifieddatalibrary.Float(1.1),
			Iron:            unifieddatalibrary.Int(3),
			Irradiance:      unifieddatalibrary.Float(1.1),
			Lat:             unifieddatalibrary.Float(1.1),
			Lon:             unifieddatalibrary.Float(1.1),
			MsgCreateDate:   unifieddatalibrary.Time(time.Now()),
			MsgType:         unifieddatalibrary.String("SU67"),
			Notes:           unifieddatalibrary.String("Example notes"),
			NumFrames:       unifieddatalibrary.Int(6),
			NumObs:          unifieddatalibrary.Int(9),
			NumTracks:       unifieddatalibrary.Int(3),
			ObType:          unifieddatalibrary.String("RADAR"),
			OrbitRegime:     unifieddatalibrary.String("GEO"),
			OrientAngle:     unifieddatalibrary.Float(1.1),
			Origin:          unifieddatalibrary.String("Example source"),
			OrigObjectID:    unifieddatalibrary.String("ORIGOBJECT-ID"),
			OrigSensorID:    unifieddatalibrary.String("ORIGSENSOR-ID"),
			PlanIndex:       unifieddatalibrary.Int(8),
			Polarization:    unifieddatalibrary.String("H"),
			Priority:        unifieddatalibrary.String("EMERGENCY"),
			Ra:              unifieddatalibrary.Float(1.1),
			Raan:            unifieddatalibrary.Float(1.1),
			Range:           unifieddatalibrary.Float(1.1),
			Rcs:             unifieddatalibrary.Float(1.1),
			RcsMax:          unifieddatalibrary.Float(1.1),
			RcsMin:          unifieddatalibrary.Float(1.1),
			Reflectance:     unifieddatalibrary.Float(1.1),
			SatNo:           unifieddatalibrary.Int(101),
			Scenario:        unifieddatalibrary.String("Example direction"),
			SemiMajorAxis:   unifieddatalibrary.Float(1.1),
			SpectralModel:   unifieddatalibrary.String("Example Model"),
			SrchInc:         unifieddatalibrary.Float(1.1),
			SrchPattern:     unifieddatalibrary.String("SCAN"),
			StateVector: unifieddatalibrary.SensorplanNewParamsCollectRequestStateVector{
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
			StopAlt:       unifieddatalibrary.Float(1.1),
			StopLat:       unifieddatalibrary.Float(1.1),
			StopLon:       unifieddatalibrary.Float(1.1),
			Suffix:        unifieddatalibrary.String("T"),
			Tags:          []string{"PROVIDER_TAG1", "PROVIDER_TAG2"},
			TargetSize:    unifieddatalibrary.Float(1.1),
			TaskCategory:  unifieddatalibrary.Int(4),
			TaskGroup:     unifieddatalibrary.String("729"),
			TaskID:        unifieddatalibrary.String("TASK-ID"),
			TransactionID: unifieddatalibrary.String("TRANSACTION-ID"),
			TrueAnomoly:   unifieddatalibrary.Float(1.1),
			UctFollowUp:   unifieddatalibrary.Bool(false),
			VisMag:        unifieddatalibrary.Float(1.1),
			VisMagMax:     unifieddatalibrary.Float(1.1),
			VisMagMin:     unifieddatalibrary.Float(1.1),
			XAngle:        unifieddatalibrary.Float(1.1),
			YAngle:        unifieddatalibrary.Float(1.1),
		}},
		Customer:     unifieddatalibrary.String("CUSTOMER"),
		EndTime:      unifieddatalibrary.Time(time.Now()),
		IDSensor:     unifieddatalibrary.String("REF-SENSOR-ID"),
		Name:         unifieddatalibrary.String("EXAMPLE NAME"),
		Origin:       unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
		OrigSensorID: unifieddatalibrary.String("ORIGSENSOR-ID"),
		Purpose:      unifieddatalibrary.String("Example purpose"),
		ReqTotal:     unifieddatalibrary.Int(2),
		SenNetwork:   unifieddatalibrary.String("NETWORK"),
		Status:       unifieddatalibrary.String("ACCEPTED"),
		Tags:         []string{"PROVIDER_TAG1", "PROVIDER_TAG2"},
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSensorplanUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Sensorplan.Update(
		context.TODO(),
		"id",
		unifieddatalibrary.SensorplanUpdateParams{
			ClassificationMarking: "U",
			DataMode:              unifieddatalibrary.SensorplanUpdateParamsDataModeTest,
			RecType:               "COLLECT",
			Source:                "Bluestaq",
			StartTime:             time.Now(),
			Type:                  "PLAN",
			ID:                    unifieddatalibrary.String("SENSORPLAN-ID"),
			CollectRequests: []unifieddatalibrary.SensorplanUpdateParamsCollectRequest{{
				ClassificationMarking: "U",
				DataMode:              "TEST",
				Source:                "Bluestaq",
				StartTime:             time.Now(),
				Type:                  "DWELL",
				ID:                    unifieddatalibrary.String("COLLECTREQUEST-ID"),
				Alt:                   unifieddatalibrary.Float(1.1),
				ArgOfPerigee:          unifieddatalibrary.Float(1.1),
				Az:                    unifieddatalibrary.Float(1.1),
				Customer:              unifieddatalibrary.String("Bluestaq"),
				Dec:                   unifieddatalibrary.Float(1.1),
				Duration:              unifieddatalibrary.Int(11),
				DwellID:               unifieddatalibrary.String("DWELL-ID"),
				Eccentricity:          unifieddatalibrary.Float(1.1),
				El:                    unifieddatalibrary.Float(1.1),
				Elset: unifieddatalibrary.SensorplanUpdateParamsCollectRequestElset{
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
				EndTime:         unifieddatalibrary.Time(time.Now()),
				Epoch:           unifieddatalibrary.Time(time.Now()),
				EsID:            unifieddatalibrary.String("ES-ID"),
				ExtentAz:        unifieddatalibrary.Float(1.1),
				ExtentEl:        unifieddatalibrary.Float(1.1),
				ExtentRange:     unifieddatalibrary.Float(1.1),
				ExternalID:      unifieddatalibrary.String("EXTERNAL-ID"),
				FrameRate:       unifieddatalibrary.Float(1.1),
				Freq:            unifieddatalibrary.Float(1.1),
				FreqMax:         unifieddatalibrary.Float(1.1),
				FreqMin:         unifieddatalibrary.Float(1.1),
				IDElset:         unifieddatalibrary.String("REF-ELSET-ID"),
				IDManifold:      unifieddatalibrary.String("REF-MANIFOLD-ID"),
				IDParentReq:     unifieddatalibrary.String("da98671b-34db-47bf-8c8d-7c668b92c800"),
				IDPlan:          unifieddatalibrary.String("REF-PLAN-ID"),
				IDSensor:        unifieddatalibrary.String("REF-SENSOR-ID"),
				IDStateVector:   unifieddatalibrary.String("STATEVECTOR-ID"),
				Inclination:     unifieddatalibrary.Float(1.1),
				IntegrationTime: unifieddatalibrary.Float(1.1),
				Iron:            unifieddatalibrary.Int(3),
				Irradiance:      unifieddatalibrary.Float(1.1),
				Lat:             unifieddatalibrary.Float(1.1),
				Lon:             unifieddatalibrary.Float(1.1),
				MsgCreateDate:   unifieddatalibrary.Time(time.Now()),
				MsgType:         unifieddatalibrary.String("SU67"),
				Notes:           unifieddatalibrary.String("Example notes"),
				NumFrames:       unifieddatalibrary.Int(6),
				NumObs:          unifieddatalibrary.Int(9),
				NumTracks:       unifieddatalibrary.Int(3),
				ObType:          unifieddatalibrary.String("RADAR"),
				OrbitRegime:     unifieddatalibrary.String("GEO"),
				OrientAngle:     unifieddatalibrary.Float(1.1),
				Origin:          unifieddatalibrary.String("Example source"),
				OrigObjectID:    unifieddatalibrary.String("ORIGOBJECT-ID"),
				OrigSensorID:    unifieddatalibrary.String("ORIGSENSOR-ID"),
				PlanIndex:       unifieddatalibrary.Int(8),
				Polarization:    unifieddatalibrary.String("H"),
				Priority:        unifieddatalibrary.String("EMERGENCY"),
				Ra:              unifieddatalibrary.Float(1.1),
				Raan:            unifieddatalibrary.Float(1.1),
				Range:           unifieddatalibrary.Float(1.1),
				Rcs:             unifieddatalibrary.Float(1.1),
				RcsMax:          unifieddatalibrary.Float(1.1),
				RcsMin:          unifieddatalibrary.Float(1.1),
				Reflectance:     unifieddatalibrary.Float(1.1),
				SatNo:           unifieddatalibrary.Int(101),
				Scenario:        unifieddatalibrary.String("Example direction"),
				SemiMajorAxis:   unifieddatalibrary.Float(1.1),
				SpectralModel:   unifieddatalibrary.String("Example Model"),
				SrchInc:         unifieddatalibrary.Float(1.1),
				SrchPattern:     unifieddatalibrary.String("SCAN"),
				StateVector: unifieddatalibrary.SensorplanUpdateParamsCollectRequestStateVector{
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
				StopAlt:       unifieddatalibrary.Float(1.1),
				StopLat:       unifieddatalibrary.Float(1.1),
				StopLon:       unifieddatalibrary.Float(1.1),
				Suffix:        unifieddatalibrary.String("T"),
				Tags:          []string{"PROVIDER_TAG1", "PROVIDER_TAG2"},
				TargetSize:    unifieddatalibrary.Float(1.1),
				TaskCategory:  unifieddatalibrary.Int(4),
				TaskGroup:     unifieddatalibrary.String("729"),
				TaskID:        unifieddatalibrary.String("TASK-ID"),
				TransactionID: unifieddatalibrary.String("TRANSACTION-ID"),
				TrueAnomoly:   unifieddatalibrary.Float(1.1),
				UctFollowUp:   unifieddatalibrary.Bool(false),
				VisMag:        unifieddatalibrary.Float(1.1),
				VisMagMax:     unifieddatalibrary.Float(1.1),
				VisMagMin:     unifieddatalibrary.Float(1.1),
				XAngle:        unifieddatalibrary.Float(1.1),
				YAngle:        unifieddatalibrary.Float(1.1),
			}},
			Customer:     unifieddatalibrary.String("CUSTOMER"),
			EndTime:      unifieddatalibrary.Time(time.Now()),
			IDSensor:     unifieddatalibrary.String("REF-SENSOR-ID"),
			Name:         unifieddatalibrary.String("EXAMPLE NAME"),
			Origin:       unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigSensorID: unifieddatalibrary.String("ORIGSENSOR-ID"),
			Purpose:      unifieddatalibrary.String("Example purpose"),
			ReqTotal:     unifieddatalibrary.Int(2),
			SenNetwork:   unifieddatalibrary.String("NETWORK"),
			Status:       unifieddatalibrary.String("ACCEPTED"),
			Tags:         []string{"PROVIDER_TAG1", "PROVIDER_TAG2"},
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

func TestSensorplanListWithOptionalParams(t *testing.T) {
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
	_, err := client.Sensorplan.List(context.TODO(), unifieddatalibrary.SensorplanListParams{
		StartTime:   time.Now(),
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

func TestSensorplanCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Sensorplan.Count(context.TODO(), unifieddatalibrary.SensorplanCountParams{
		StartTime:   time.Now(),
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

func TestSensorplanGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Sensorplan.Get(
		context.TODO(),
		"id",
		unifieddatalibrary.SensorplanGetParams{
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

func TestSensorplanQueryhelp(t *testing.T) {
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
	err := client.Sensorplan.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSensorplanTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.Sensorplan.Tuple(context.TODO(), unifieddatalibrary.SensorplanTupleParams{
		Columns:     "columns",
		StartTime:   time.Now(),
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

func TestSensorplanUnvalidatedPublish(t *testing.T) {
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
	err := client.Sensorplan.UnvalidatedPublish(context.TODO(), unifieddatalibrary.SensorplanUnvalidatedPublishParams{
		Body: []unifieddatalibrary.SensorplanUnvalidatedPublishParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			RecType:               "COLLECT",
			Source:                "Bluestaq",
			StartTime:             time.Now(),
			Type:                  "PLAN",
			ID:                    unifieddatalibrary.String("SENSORPLAN-ID"),
			CollectRequests: []unifieddatalibrary.SensorplanUnvalidatedPublishParamsBodyCollectRequest{{
				ClassificationMarking: "U",
				DataMode:              "TEST",
				Source:                "Bluestaq",
				StartTime:             time.Now(),
				Type:                  "DWELL",
				ID:                    unifieddatalibrary.String("COLLECTREQUEST-ID"),
				Alt:                   unifieddatalibrary.Float(1.1),
				ArgOfPerigee:          unifieddatalibrary.Float(1.1),
				Az:                    unifieddatalibrary.Float(1.1),
				Customer:              unifieddatalibrary.String("Bluestaq"),
				Dec:                   unifieddatalibrary.Float(1.1),
				Duration:              unifieddatalibrary.Int(11),
				DwellID:               unifieddatalibrary.String("DWELL-ID"),
				Eccentricity:          unifieddatalibrary.Float(1.1),
				El:                    unifieddatalibrary.Float(1.1),
				Elset: unifieddatalibrary.SensorplanUnvalidatedPublishParamsBodyCollectRequestElset{
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
				EndTime:         unifieddatalibrary.Time(time.Now()),
				Epoch:           unifieddatalibrary.Time(time.Now()),
				EsID:            unifieddatalibrary.String("ES-ID"),
				ExtentAz:        unifieddatalibrary.Float(1.1),
				ExtentEl:        unifieddatalibrary.Float(1.1),
				ExtentRange:     unifieddatalibrary.Float(1.1),
				ExternalID:      unifieddatalibrary.String("EXTERNAL-ID"),
				FrameRate:       unifieddatalibrary.Float(1.1),
				Freq:            unifieddatalibrary.Float(1.1),
				FreqMax:         unifieddatalibrary.Float(1.1),
				FreqMin:         unifieddatalibrary.Float(1.1),
				IDElset:         unifieddatalibrary.String("REF-ELSET-ID"),
				IDManifold:      unifieddatalibrary.String("REF-MANIFOLD-ID"),
				IDParentReq:     unifieddatalibrary.String("da98671b-34db-47bf-8c8d-7c668b92c800"),
				IDPlan:          unifieddatalibrary.String("REF-PLAN-ID"),
				IDSensor:        unifieddatalibrary.String("REF-SENSOR-ID"),
				IDStateVector:   unifieddatalibrary.String("STATEVECTOR-ID"),
				Inclination:     unifieddatalibrary.Float(1.1),
				IntegrationTime: unifieddatalibrary.Float(1.1),
				Iron:            unifieddatalibrary.Int(3),
				Irradiance:      unifieddatalibrary.Float(1.1),
				Lat:             unifieddatalibrary.Float(1.1),
				Lon:             unifieddatalibrary.Float(1.1),
				MsgCreateDate:   unifieddatalibrary.Time(time.Now()),
				MsgType:         unifieddatalibrary.String("SU67"),
				Notes:           unifieddatalibrary.String("Example notes"),
				NumFrames:       unifieddatalibrary.Int(6),
				NumObs:          unifieddatalibrary.Int(9),
				NumTracks:       unifieddatalibrary.Int(3),
				ObType:          unifieddatalibrary.String("RADAR"),
				OrbitRegime:     unifieddatalibrary.String("GEO"),
				OrientAngle:     unifieddatalibrary.Float(1.1),
				Origin:          unifieddatalibrary.String("Example source"),
				OrigObjectID:    unifieddatalibrary.String("ORIGOBJECT-ID"),
				OrigSensorID:    unifieddatalibrary.String("ORIGSENSOR-ID"),
				PlanIndex:       unifieddatalibrary.Int(8),
				Polarization:    unifieddatalibrary.String("H"),
				Priority:        unifieddatalibrary.String("EMERGENCY"),
				Ra:              unifieddatalibrary.Float(1.1),
				Raan:            unifieddatalibrary.Float(1.1),
				Range:           unifieddatalibrary.Float(1.1),
				Rcs:             unifieddatalibrary.Float(1.1),
				RcsMax:          unifieddatalibrary.Float(1.1),
				RcsMin:          unifieddatalibrary.Float(1.1),
				Reflectance:     unifieddatalibrary.Float(1.1),
				SatNo:           unifieddatalibrary.Int(101),
				Scenario:        unifieddatalibrary.String("Example direction"),
				SemiMajorAxis:   unifieddatalibrary.Float(1.1),
				SpectralModel:   unifieddatalibrary.String("Example Model"),
				SrchInc:         unifieddatalibrary.Float(1.1),
				SrchPattern:     unifieddatalibrary.String("SCAN"),
				StateVector: unifieddatalibrary.SensorplanUnvalidatedPublishParamsBodyCollectRequestStateVector{
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
				StopAlt:       unifieddatalibrary.Float(1.1),
				StopLat:       unifieddatalibrary.Float(1.1),
				StopLon:       unifieddatalibrary.Float(1.1),
				Suffix:        unifieddatalibrary.String("T"),
				Tags:          []string{"PROVIDER_TAG1", "PROVIDER_TAG2"},
				TargetSize:    unifieddatalibrary.Float(1.1),
				TaskCategory:  unifieddatalibrary.Int(4),
				TaskGroup:     unifieddatalibrary.String("729"),
				TaskID:        unifieddatalibrary.String("TASK-ID"),
				TransactionID: unifieddatalibrary.String("TRANSACTION-ID"),
				TrueAnomoly:   unifieddatalibrary.Float(1.1),
				UctFollowUp:   unifieddatalibrary.Bool(false),
				VisMag:        unifieddatalibrary.Float(1.1),
				VisMagMax:     unifieddatalibrary.Float(1.1),
				VisMagMin:     unifieddatalibrary.Float(1.1),
				XAngle:        unifieddatalibrary.Float(1.1),
				YAngle:        unifieddatalibrary.Float(1.1),
			}},
			Customer:     unifieddatalibrary.String("CUSTOMER"),
			EndTime:      unifieddatalibrary.Time(time.Now()),
			IDSensor:     unifieddatalibrary.String("REF-SENSOR-ID"),
			Name:         unifieddatalibrary.String("EXAMPLE NAME"),
			Origin:       unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigSensorID: unifieddatalibrary.String("ORIGSENSOR-ID"),
			Purpose:      unifieddatalibrary.String("Example purpose"),
			ReqTotal:     unifieddatalibrary.Int(2),
			SenNetwork:   unifieddatalibrary.String("NETWORK"),
			Status:       unifieddatalibrary.String("ACCEPTED"),
			Tags:         []string{"PROVIDER_TAG1", "PROVIDER_TAG2"},
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
