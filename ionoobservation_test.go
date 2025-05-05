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

func TestIonOobservationListWithOptionalParams(t *testing.T) {
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
	_, err := client.IonOobservation.List(context.TODO(), unifieddatalibrary.IonOobservationListParams{
		StartTimeUtc: time.Now(),
		FirstResult:  unifieddatalibrary.Int(0),
		MaxResults:   unifieddatalibrary.Int(0),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestIonOobservationCountWithOptionalParams(t *testing.T) {
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
	_, err := client.IonOobservation.Count(context.TODO(), unifieddatalibrary.IonOobservationCountParams{
		StartTimeUtc: time.Now(),
		FirstResult:  unifieddatalibrary.Int(0),
		MaxResults:   unifieddatalibrary.Int(0),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestIonOobservationNewBulk(t *testing.T) {
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
	err := client.IonOobservation.NewBulk(context.TODO(), unifieddatalibrary.IonOobservationNewBulkParams{
		Body: []unifieddatalibrary.IonOobservationNewBulkParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			Source:                "Bluestaq",
			StartTimeUtc:          time.Now(),
			StationID:             "STATION-ID",
			System:                "Example hardware type",
			SystemInfo:            "Example settings",
			ID:                    unifieddatalibrary.String("IONOOBSERVATION-ID"),
			Amplitude: unifieddatalibrary.IonOobservationNewBulkParamsBodyAmplitude{
				Data:          [][][][][][][]float64{{{{{{{0.02, 0.034}}}}}}},
				DimensionName: []string{"NAME1", "NAME2"},
				Dimensions:    []int64{4, 5},
				Notes:         unifieddatalibrary.String("NOTES"),
			},
			AntennaElementPosition: unifieddatalibrary.IonOobservationNewBulkParamsBodyAntennaElementPosition{
				Data:          [][]float64{{1.23, 0.123}},
				DimensionName: []string{"NAME1", "NAME2"},
				Dimensions:    []int64{3, 4},
			},
			AntennaElementPositionCoordinateSystem: "J2000",
			ArtistFlags:                            []int64{1, 2, 3},
			Azimuth: unifieddatalibrary.IonOobservationNewBulkParamsBodyAzimuth{
				Data:          [][][][][][][]float64{{{{{{{0.02, 0.034}}}}}}},
				DimensionName: []string{"NAME1", "NAME2"},
				Dimensions:    []int64{1, 2},
				Notes:         unifieddatalibrary.String("NOTES"),
			},
			B0: unifieddatalibrary.Float(68.07),
			B1: unifieddatalibrary.Float(1.87),
			CharAtts: []unifieddatalibrary.IonOobservationNewBulkParamsBodyCharAtt{{
				CharName:                unifieddatalibrary.String("hprimeF2"),
				ClimateModelInputParams: []string{"ISSN1 88.1", "Option 2"},
				ClimateModelName:        unifieddatalibrary.String("IRI"),
				ClimateModelOptions:     []string{"URSI-88", "No storm"},
				D:                       unifieddatalibrary.String("K"),
				LowerBound:              unifieddatalibrary.Float(1.26),
				Q:                       unifieddatalibrary.String("T"),
				UncertaintyBoundType:    unifieddatalibrary.String("1sigma"),
				UpperBound:              unifieddatalibrary.Float(2.57),
				UrsiID:                  unifieddatalibrary.String("04"),
			}},
			D:  unifieddatalibrary.Float(1.1),
			D1: unifieddatalibrary.Float(1.94),
			Datum: unifieddatalibrary.IonOobservationNewBulkParamsBodyDatum{
				Data:  []float64{1.1, 2.1, 3.1},
				Notes: unifieddatalibrary.String("NOTES"),
			},
			DeltafoF2: unifieddatalibrary.Float(1.1),
			DensityProfile: unifieddatalibrary.IonOobservationNewBulkParamsBodyDensityProfile{
				Iri: unifieddatalibrary.IonOobservationNewBulkParamsBodyDensityProfileIri{
					B0:          unifieddatalibrary.Float(245.1),
					B1:          unifieddatalibrary.Float(3.45),
					Chi:         unifieddatalibrary.Float(35.1),
					D1:          unifieddatalibrary.Float(0),
					Description: unifieddatalibrary.String("Full altitude range of 1D vertical plasma density profile for D, E, and F regions of the ionosphere provided by IRI, in which model parameters are replaced with observed ionogram-derived data where available."),
					Fp1:         unifieddatalibrary.Float(0.0474732023322638),
					Fp2:         unifieddatalibrary.Float(-0.00112685246984002),
					Fp30:        unifieddatalibrary.Float(0.00827559450035957),
					Fp3U:        unifieddatalibrary.Float(0.000201178475411428),
					Ha:          unifieddatalibrary.Float(50.1),
					Hdx:         unifieddatalibrary.Float(85.6),
					HmD:         unifieddatalibrary.Float(81.1),
					HmE:         unifieddatalibrary.Float(99.8),
					HmF1:        unifieddatalibrary.Float(210.3),
					HmF2:        unifieddatalibrary.Float(265.42),
					HValTop:     unifieddatalibrary.Float(116.2),
					Hz:          unifieddatalibrary.Float(142.7),
					NmD:         unifieddatalibrary.Float(937543116.1),
					NmE:         unifieddatalibrary.Float(154164.1),
					NmF1:        unifieddatalibrary.Float(210486),
					NmF2:        unifieddatalibrary.Float(313283.1),
					NValB:       unifieddatalibrary.Float(147025.1),
				},
				Parabolic: unifieddatalibrary.IonOobservationNewBulkParamsBodyDensityProfileParabolic{
					Description: unifieddatalibrary.String("Best-fit algorithm in NHPC software."),
					ParabolicItems: []unifieddatalibrary.IonOobservationNewBulkParamsBodyDensityProfileParabolicParabolicItem{{
						F:     unifieddatalibrary.Float(3.621),
						Layer: unifieddatalibrary.String("E"),
						Y:     unifieddatalibrary.Float(11.1),
						Z:     unifieddatalibrary.Float(110.2),
					}},
				},
				QuasiParabolic: unifieddatalibrary.IonOobservationNewBulkParamsBodyDensityProfileQuasiParabolic{
					Description: unifieddatalibrary.String("Array of the best-fit 3-parameter quasi-parabolas defined via A, B, C coefficients, f^2=A/r^2+B/r+C"),
					EarthRadius: unifieddatalibrary.Float(6370.1),
					QuasiParabolicSegments: []unifieddatalibrary.IonOobservationNewBulkParamsBodyDensityProfileQuasiParabolicQuasiParabolicSegment{{
						A:     unifieddatalibrary.Float(-550273940000),
						B:     unifieddatalibrary.Float(169837632),
						C:     unifieddatalibrary.Float(13104.63),
						Error: unifieddatalibrary.Float(0.0016),
						Index: unifieddatalibrary.Int(12),
						Rb:    unifieddatalibrary.Float(6460.1),
						Re:    unifieddatalibrary.Float(6480.001),
					}},
				},
				ShiftedChebyshev: unifieddatalibrary.IonOobservationNewBulkParamsBodyDensityProfileShiftedChebyshev{
					Description: unifieddatalibrary.String("Best-fit Huang-Reinisch formalism based on shifted Chebyshev expansion."),
					ShiftedChebyshevs: []unifieddatalibrary.IonOobservationNewBulkParamsBodyDensityProfileShiftedChebyshevShiftedChebyshev{{
						Coeffs:     []float64{-11.2, -10.536, 3.357, -0.888, 0.155},
						Error:      unifieddatalibrary.Float(0.02),
						Fstart:     unifieddatalibrary.Float(0.2),
						Fstop:      unifieddatalibrary.Float(1.39),
						Layer:      unifieddatalibrary.String("F2"),
						N:          unifieddatalibrary.Int(5),
						Peakheight: unifieddatalibrary.Float(110.1),
						ZhalfNm:    unifieddatalibrary.Float(210.1),
					}},
				},
				TopsideExtensionChapmanConst: unifieddatalibrary.IonOobservationNewBulkParamsBodyDensityProfileTopsideExtensionChapmanConst{
					Chi:         unifieddatalibrary.Float(35.1),
					Description: unifieddatalibrary.String("Constant scale height Chapman topside layer determined using bottomside peak density parameters and placed chi km above the ionosonde-determined peak height."),
					HmF2:        unifieddatalibrary.Float(265.42),
					NmF2:        unifieddatalibrary.Float(313283.1),
					ScaleF2:     unifieddatalibrary.Float(45.191),
				},
				TopsideExtensionVaryChap: unifieddatalibrary.IonOobservationNewBulkParamsBodyDensityProfileTopsideExtensionVaryChap{
					Alpha:       unifieddatalibrary.Float(30.1),
					Beta:        unifieddatalibrary.Float(30.1),
					Chi:         unifieddatalibrary.Float(30.1),
					Description: unifieddatalibrary.String("Constant scale height Chapman topside layer determined using bottomside peak density parameters and placed chi km above the ionosonde-determined peak height."),
					HmF2:        unifieddatalibrary.Float(265.42),
					Ht:          unifieddatalibrary.Float(30.1),
					NmF2:        unifieddatalibrary.Float(313283.1),
					ScaleF2:     unifieddatalibrary.Float(45.191),
				},
				ValleyModelCoeffs:      []float64{39.597, 0.1777},
				ValleyModelDescription: unifieddatalibrary.String("2-parameter UMLCAR model with width W in km and depth D in MHz, no fitting."),
			},
			Doppler: unifieddatalibrary.IonOobservationNewBulkParamsBodyDoppler{
				Data:          [][][][][][][]float64{{{{{{{0.02, 0.034}}}}}}},
				DimensionName: []string{"NAME1", "NAME2"},
				Dimensions:    []int64{1, 2},
				Notes:         unifieddatalibrary.String("NOTES"),
			},
			DownE:                      unifieddatalibrary.Float(1.1),
			DownEs:                     unifieddatalibrary.Float(1.1),
			DownF:                      unifieddatalibrary.Float(1.1),
			ElectronDensity:            []float64{1.1, 2.1, 3.1},
			ElectronDensityUncertainty: []float64{0.8, 0.2, 0.5},
			Elevation: unifieddatalibrary.IonOobservationNewBulkParamsBodyElevation{
				Data:          [][][][][][][]float64{{{{{{{0.02, 0.034}}}}}}},
				DimensionName: []string{"NAME1", "NAME2"},
				Dimensions:    []int64{2, 3},
				Notes:         unifieddatalibrary.String("NOTES"),
			},
			FbEs:      unifieddatalibrary.Float(34.867),
			Fe:        unifieddatalibrary.Float(1.23),
			Ff:        unifieddatalibrary.Float(0.075),
			FhprimeF:  unifieddatalibrary.Float(2.5),
			FhprimeF2: unifieddatalibrary.Float(2.5),
			Fmin:      unifieddatalibrary.Float(2.025),
			FminE:     unifieddatalibrary.Float(2.025),
			FminEs:    unifieddatalibrary.Float(1.1),
			FminF:     unifieddatalibrary.Float(2.7),
			Fmuf:      unifieddatalibrary.Float(1.1),
			FoE:       unifieddatalibrary.Float(1.1),
			FoEa:      unifieddatalibrary.Float(45.764),
			FoEp:      unifieddatalibrary.Float(1.1),
			FoEs:      unifieddatalibrary.Float(2.35),
			FoF1:      unifieddatalibrary.Float(8.2),
			FoF1p:     unifieddatalibrary.Float(1.1),
			FoF2:      unifieddatalibrary.Float(6.75),
			FoF2p:     unifieddatalibrary.Float(1.1),
			FoP:       unifieddatalibrary.Float(87.21),
			Frequency: unifieddatalibrary.IonOobservationNewBulkParamsBodyFrequency{
				Data:          [][][][][][][]float64{{{{{{{0.02, 0.034}}}}}}},
				DimensionName: []string{"NAME1", "NAME2"},
				Dimensions:    []int64{1, 2},
				Notes:         unifieddatalibrary.String("NOTES"),
			},
			FxE:              unifieddatalibrary.Float(1.1),
			FxF1:             unifieddatalibrary.Float(4.2),
			FxF2:             unifieddatalibrary.Float(4.5),
			FxI:              unifieddatalibrary.Float(7.525),
			Height:           []float64{1.1, 2.1, 3.1},
			HmE:              unifieddatalibrary.Float(215.643),
			HmF1:             unifieddatalibrary.Float(230.128),
			HmF2:             unifieddatalibrary.Float(240.498),
			HprimeE:          unifieddatalibrary.Float(98.47),
			HprimeEa:         unifieddatalibrary.Float(102.6),
			HprimeEs:         unifieddatalibrary.Float(95),
			HprimeF:          unifieddatalibrary.Float(238.5),
			HprimeF1:         unifieddatalibrary.Float(230.1),
			HprimeF2:         unifieddatalibrary.Float(238.5),
			HprimefMuf:       unifieddatalibrary.Float(1.1),
			HprimeP:          unifieddatalibrary.Float(89.45),
			IDSensor:         unifieddatalibrary.String("SENSOR-ID"),
			Luf:              unifieddatalibrary.Float(1.1),
			Md:               unifieddatalibrary.Float(1.1),
			Mufd:             unifieddatalibrary.Float(1.1),
			NeProfileName:    unifieddatalibrary.String("NH"),
			NeProfileVersion: unifieddatalibrary.Float(4.32),
			Origin:           unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigSensorID:     unifieddatalibrary.String("ORIGSENSOR-ID"),
			Phase: unifieddatalibrary.IonOobservationNewBulkParamsBodyPhase{
				Data:          [][][][][][][]float64{{{{{{{0.02, 0.034}}}}}}},
				DimensionName: []string{"NAME1", "NAME2"},
				Dimensions:    []int64{1, 2},
				Notes:         unifieddatalibrary.String("NOTES"),
			},
			PlasmaFrequency:            []float64{1.1, 2.1, 3.1},
			PlasmaFrequencyUncertainty: []float64{0.8, 0.2, 0.5},
			PlatformName:               unifieddatalibrary.String("Millstone Hill"),
			Polarization: unifieddatalibrary.IonOobservationNewBulkParamsBodyPolarization{
				Data:          [][][][][][][]string{{{{{{{"X", "O"}}}}}}},
				DimensionName: []string{"NAME1", "NAME2"},
				Dimensions:    []int64{1, 2},
				Notes:         unifieddatalibrary.String("NOTES"),
			},
			Power: unifieddatalibrary.IonOobservationNewBulkParamsBodyPower{
				Data:          [][][][][][][]float64{{{{{{{0.02, 0.034}}}}}}},
				DimensionName: []string{"NAME1", "NAME2"},
				Dimensions:    []int64{1, 2},
				Notes:         unifieddatalibrary.String("NOTES"),
			},
			Qe: unifieddatalibrary.Float(0.95),
			Qf: unifieddatalibrary.Float(1.83),
			Range: unifieddatalibrary.IonOobservationNewBulkParamsBodyRange{
				Data:          [][][][][][][]float64{{{{{{{0.02, 0.034}}}}}}},
				DimensionName: []string{"NAME1", "NAME2"},
				Dimensions:    []int64{1, 2},
				Notes:         unifieddatalibrary.String("NOTES"),
			},
			ReceiveCoordinates:       [][]float64{{45.5, 179.3, 35.6}, {-80.2, -20.5, 43.2}},
			ReceiveSensorType:        "Mobile",
			RestrictedFrequency:      []float64{12.5, 34.5, 45.3},
			RestrictedFrequencyNotes: unifieddatalibrary.String("Example notes"),
			ScaleHeightF2Peak:        unifieddatalibrary.Float(35.613),
			ScalerInfo: unifieddatalibrary.IonOobservationNewBulkParamsBodyScalerInfo{
				ConfidenceLevel: unifieddatalibrary.Int(11),
				ConfidenceScore: unifieddatalibrary.Int(75),
				Name:            unifieddatalibrary.String("ARTIST-4"),
				Organization:    unifieddatalibrary.String("UML"),
				Type:            unifieddatalibrary.String("MANUAL"),
				Version:         unifieddatalibrary.Float(500200.1),
			},
			Stokes: unifieddatalibrary.IonOobservationNewBulkParamsBodyStokes{
				Data:          [][][][][][][]float64{{{{{{{0.02, 0.034}}}}}}},
				DimensionName: []string{"NAMES1", "NAMES2"},
				Dimensions:    []int64{2, 3},
				Notes:         unifieddatalibrary.String("NOTES"),
				S:             []float64{1, 2},
			},
			SystemNotes:    unifieddatalibrary.String("Example notes"),
			Tec:            unifieddatalibrary.Float(24.673),
			TidAzimuth:     []float64{1.1, 2.1, 3.1},
			TidPeriods:     []float64{1.1, 2.1, 3.1},
			TidPhaseSpeeds: []float64{1.1, 2.1, 3.1},
			Time: unifieddatalibrary.IonOobservationNewBulkParamsBodyTime{
				Data:          [][][][][][][]float64{{{{{{{0.02, 0.034}}}}}}},
				DimensionName: []string{"NAME1", "NAME2"},
				Dimensions:    []int64{1, 2},
				Notes:         unifieddatalibrary.String("NOTES"),
			},
			TraceGeneric: unifieddatalibrary.IonOobservationNewBulkParamsBodyTraceGeneric{
				Data:          [][][]float64{{{1.23, 1.0903}}},
				DimensionName: []string{"NAME1", "NAME2"},
				Notes:         unifieddatalibrary.String("NOTES"),
			},
			TransmitCoordinates: [][]float64{{45.5, 179.3, 35.6}, {-80.2, -20.5, 43.2}},
			TransmitSensorType:  "Mobile",
			TypeEs:              unifieddatalibrary.String("Auroral"),
			YE:                  unifieddatalibrary.Float(1.722),
			YF1:                 unifieddatalibrary.Float(55.645),
			YF2:                 unifieddatalibrary.Float(62.178),
			ZhalfNm:             unifieddatalibrary.Float(240.498),
			ZmE:                 unifieddatalibrary.Float(91.744),
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

func TestIonOobservationQueryhelp(t *testing.T) {
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
	err := client.IonOobservation.Queryhelp(context.TODO())
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestIonOobservationTupleWithOptionalParams(t *testing.T) {
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
	_, err := client.IonOobservation.Tuple(context.TODO(), unifieddatalibrary.IonOobservationTupleParams{
		Columns:      "columns",
		StartTimeUtc: time.Now(),
		FirstResult:  unifieddatalibrary.Int(0),
		MaxResults:   unifieddatalibrary.Int(0),
	})
	if err != nil {
		var apierr *unifieddatalibrary.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestIonOobservationUnvalidatedPublish(t *testing.T) {
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
	err := client.IonOobservation.UnvalidatedPublish(context.TODO(), unifieddatalibrary.IonOobservationUnvalidatedPublishParams{
		Body: []unifieddatalibrary.IonOobservationUnvalidatedPublishParamsBody{{
			ClassificationMarking: "U",
			DataMode:              "TEST",
			Source:                "Bluestaq",
			StartTimeUtc:          time.Now(),
			StationID:             "STATION-ID",
			System:                "Example hardware type",
			SystemInfo:            "Example settings",
			ID:                    unifieddatalibrary.String("IONOOBSERVATION-ID"),
			Amplitude: unifieddatalibrary.IonOobservationUnvalidatedPublishParamsBodyAmplitude{
				Data:          [][][][][][][]float64{{{{{{{0.02, 0.034}}}}}}},
				DimensionName: []string{"NAME1", "NAME2"},
				Dimensions:    []int64{4, 5},
				Notes:         unifieddatalibrary.String("NOTES"),
			},
			AntennaElementPosition: unifieddatalibrary.IonOobservationUnvalidatedPublishParamsBodyAntennaElementPosition{
				Data:          [][]float64{{1.23, 0.123}},
				DimensionName: []string{"NAME1", "NAME2"},
				Dimensions:    []int64{3, 4},
			},
			AntennaElementPositionCoordinateSystem: "J2000",
			ArtistFlags:                            []int64{1, 2, 3},
			Azimuth: unifieddatalibrary.IonOobservationUnvalidatedPublishParamsBodyAzimuth{
				Data:          [][][][][][][]float64{{{{{{{0.02, 0.034}}}}}}},
				DimensionName: []string{"NAME1", "NAME2"},
				Dimensions:    []int64{1, 2},
				Notes:         unifieddatalibrary.String("NOTES"),
			},
			B0: unifieddatalibrary.Float(68.07),
			B1: unifieddatalibrary.Float(1.87),
			CharAtts: []unifieddatalibrary.IonOobservationUnvalidatedPublishParamsBodyCharAtt{{
				CharName:                unifieddatalibrary.String("hprimeF2"),
				ClimateModelInputParams: []string{"ISSN1 88.1", "Option 2"},
				ClimateModelName:        unifieddatalibrary.String("IRI"),
				ClimateModelOptions:     []string{"URSI-88", "No storm"},
				D:                       unifieddatalibrary.String("K"),
				LowerBound:              unifieddatalibrary.Float(1.26),
				Q:                       unifieddatalibrary.String("T"),
				UncertaintyBoundType:    unifieddatalibrary.String("1sigma"),
				UpperBound:              unifieddatalibrary.Float(2.57),
				UrsiID:                  unifieddatalibrary.String("04"),
			}},
			D:  unifieddatalibrary.Float(1.1),
			D1: unifieddatalibrary.Float(1.94),
			Datum: unifieddatalibrary.IonOobservationUnvalidatedPublishParamsBodyDatum{
				Data:  []float64{1.1, 2.1, 3.1},
				Notes: unifieddatalibrary.String("NOTES"),
			},
			DeltafoF2: unifieddatalibrary.Float(1.1),
			DensityProfile: unifieddatalibrary.IonOobservationUnvalidatedPublishParamsBodyDensityProfile{
				Iri: unifieddatalibrary.IonOobservationUnvalidatedPublishParamsBodyDensityProfileIri{
					B0:          unifieddatalibrary.Float(245.1),
					B1:          unifieddatalibrary.Float(3.45),
					Chi:         unifieddatalibrary.Float(35.1),
					D1:          unifieddatalibrary.Float(0),
					Description: unifieddatalibrary.String("Full altitude range of 1D vertical plasma density profile for D, E, and F regions of the ionosphere provided by IRI, in which model parameters are replaced with observed ionogram-derived data where available."),
					Fp1:         unifieddatalibrary.Float(0.0474732023322638),
					Fp2:         unifieddatalibrary.Float(-0.00112685246984002),
					Fp30:        unifieddatalibrary.Float(0.00827559450035957),
					Fp3U:        unifieddatalibrary.Float(0.000201178475411428),
					Ha:          unifieddatalibrary.Float(50.1),
					Hdx:         unifieddatalibrary.Float(85.6),
					HmD:         unifieddatalibrary.Float(81.1),
					HmE:         unifieddatalibrary.Float(99.8),
					HmF1:        unifieddatalibrary.Float(210.3),
					HmF2:        unifieddatalibrary.Float(265.42),
					HValTop:     unifieddatalibrary.Float(116.2),
					Hz:          unifieddatalibrary.Float(142.7),
					NmD:         unifieddatalibrary.Float(937543116.1),
					NmE:         unifieddatalibrary.Float(154164.1),
					NmF1:        unifieddatalibrary.Float(210486),
					NmF2:        unifieddatalibrary.Float(313283.1),
					NValB:       unifieddatalibrary.Float(147025.1),
				},
				Parabolic: unifieddatalibrary.IonOobservationUnvalidatedPublishParamsBodyDensityProfileParabolic{
					Description: unifieddatalibrary.String("Best-fit algorithm in NHPC software."),
					ParabolicItems: []unifieddatalibrary.IonOobservationUnvalidatedPublishParamsBodyDensityProfileParabolicParabolicItem{{
						F:     unifieddatalibrary.Float(3.621),
						Layer: unifieddatalibrary.String("E"),
						Y:     unifieddatalibrary.Float(11.1),
						Z:     unifieddatalibrary.Float(110.2),
					}},
				},
				QuasiParabolic: unifieddatalibrary.IonOobservationUnvalidatedPublishParamsBodyDensityProfileQuasiParabolic{
					Description: unifieddatalibrary.String("Array of the best-fit 3-parameter quasi-parabolas defined via A, B, C coefficients, f^2=A/r^2+B/r+C"),
					EarthRadius: unifieddatalibrary.Float(6370.1),
					QuasiParabolicSegments: []unifieddatalibrary.IonOobservationUnvalidatedPublishParamsBodyDensityProfileQuasiParabolicQuasiParabolicSegment{{
						A:     unifieddatalibrary.Float(-550273940000),
						B:     unifieddatalibrary.Float(169837632),
						C:     unifieddatalibrary.Float(13104.63),
						Error: unifieddatalibrary.Float(0.0016),
						Index: unifieddatalibrary.Int(12),
						Rb:    unifieddatalibrary.Float(6460.1),
						Re:    unifieddatalibrary.Float(6480.001),
					}},
				},
				ShiftedChebyshev: unifieddatalibrary.IonOobservationUnvalidatedPublishParamsBodyDensityProfileShiftedChebyshev{
					Description: unifieddatalibrary.String("Best-fit Huang-Reinisch formalism based on shifted Chebyshev expansion."),
					ShiftedChebyshevs: []unifieddatalibrary.IonOobservationUnvalidatedPublishParamsBodyDensityProfileShiftedChebyshevShiftedChebyshev{{
						Coeffs:     []float64{-11.2, -10.536, 3.357, -0.888, 0.155},
						Error:      unifieddatalibrary.Float(0.02),
						Fstart:     unifieddatalibrary.Float(0.2),
						Fstop:      unifieddatalibrary.Float(1.39),
						Layer:      unifieddatalibrary.String("F2"),
						N:          unifieddatalibrary.Int(5),
						Peakheight: unifieddatalibrary.Float(110.1),
						ZhalfNm:    unifieddatalibrary.Float(210.1),
					}},
				},
				TopsideExtensionChapmanConst: unifieddatalibrary.IonOobservationUnvalidatedPublishParamsBodyDensityProfileTopsideExtensionChapmanConst{
					Chi:         unifieddatalibrary.Float(35.1),
					Description: unifieddatalibrary.String("Constant scale height Chapman topside layer determined using bottomside peak density parameters and placed chi km above the ionosonde-determined peak height."),
					HmF2:        unifieddatalibrary.Float(265.42),
					NmF2:        unifieddatalibrary.Float(313283.1),
					ScaleF2:     unifieddatalibrary.Float(45.191),
				},
				TopsideExtensionVaryChap: unifieddatalibrary.IonOobservationUnvalidatedPublishParamsBodyDensityProfileTopsideExtensionVaryChap{
					Alpha:       unifieddatalibrary.Float(30.1),
					Beta:        unifieddatalibrary.Float(30.1),
					Chi:         unifieddatalibrary.Float(30.1),
					Description: unifieddatalibrary.String("Constant scale height Chapman topside layer determined using bottomside peak density parameters and placed chi km above the ionosonde-determined peak height."),
					HmF2:        unifieddatalibrary.Float(265.42),
					Ht:          unifieddatalibrary.Float(30.1),
					NmF2:        unifieddatalibrary.Float(313283.1),
					ScaleF2:     unifieddatalibrary.Float(45.191),
				},
				ValleyModelCoeffs:      []float64{39.597, 0.1777},
				ValleyModelDescription: unifieddatalibrary.String("2-parameter UMLCAR model with width W in km and depth D in MHz, no fitting."),
			},
			Doppler: unifieddatalibrary.IonOobservationUnvalidatedPublishParamsBodyDoppler{
				Data:          [][][][][][][]float64{{{{{{{0.02, 0.034}}}}}}},
				DimensionName: []string{"NAME1", "NAME2"},
				Dimensions:    []int64{1, 2},
				Notes:         unifieddatalibrary.String("NOTES"),
			},
			DownE:                      unifieddatalibrary.Float(1.1),
			DownEs:                     unifieddatalibrary.Float(1.1),
			DownF:                      unifieddatalibrary.Float(1.1),
			ElectronDensity:            []float64{1.1, 2.1, 3.1},
			ElectronDensityUncertainty: []float64{0.8, 0.2, 0.5},
			Elevation: unifieddatalibrary.IonOobservationUnvalidatedPublishParamsBodyElevation{
				Data:          [][][][][][][]float64{{{{{{{0.02, 0.034}}}}}}},
				DimensionName: []string{"NAME1", "NAME2"},
				Dimensions:    []int64{2, 3},
				Notes:         unifieddatalibrary.String("NOTES"),
			},
			FbEs:      unifieddatalibrary.Float(34.867),
			Fe:        unifieddatalibrary.Float(1.23),
			Ff:        unifieddatalibrary.Float(0.075),
			FhprimeF:  unifieddatalibrary.Float(2.5),
			FhprimeF2: unifieddatalibrary.Float(2.5),
			Fmin:      unifieddatalibrary.Float(2.025),
			FminE:     unifieddatalibrary.Float(2.025),
			FminEs:    unifieddatalibrary.Float(1.1),
			FminF:     unifieddatalibrary.Float(2.7),
			Fmuf:      unifieddatalibrary.Float(1.1),
			FoE:       unifieddatalibrary.Float(1.1),
			FoEa:      unifieddatalibrary.Float(45.764),
			FoEp:      unifieddatalibrary.Float(1.1),
			FoEs:      unifieddatalibrary.Float(2.35),
			FoF1:      unifieddatalibrary.Float(8.2),
			FoF1p:     unifieddatalibrary.Float(1.1),
			FoF2:      unifieddatalibrary.Float(6.75),
			FoF2p:     unifieddatalibrary.Float(1.1),
			FoP:       unifieddatalibrary.Float(87.21),
			Frequency: unifieddatalibrary.IonOobservationUnvalidatedPublishParamsBodyFrequency{
				Data:          [][][][][][][]float64{{{{{{{0.02, 0.034}}}}}}},
				DimensionName: []string{"NAME1", "NAME2"},
				Dimensions:    []int64{1, 2},
				Notes:         unifieddatalibrary.String("NOTES"),
			},
			FxE:              unifieddatalibrary.Float(1.1),
			FxF1:             unifieddatalibrary.Float(4.2),
			FxF2:             unifieddatalibrary.Float(4.5),
			FxI:              unifieddatalibrary.Float(7.525),
			Height:           []float64{1.1, 2.1, 3.1},
			HmE:              unifieddatalibrary.Float(215.643),
			HmF1:             unifieddatalibrary.Float(230.128),
			HmF2:             unifieddatalibrary.Float(240.498),
			HprimeE:          unifieddatalibrary.Float(98.47),
			HprimeEa:         unifieddatalibrary.Float(102.6),
			HprimeEs:         unifieddatalibrary.Float(95),
			HprimeF:          unifieddatalibrary.Float(238.5),
			HprimeF1:         unifieddatalibrary.Float(230.1),
			HprimeF2:         unifieddatalibrary.Float(238.5),
			HprimefMuf:       unifieddatalibrary.Float(1.1),
			HprimeP:          unifieddatalibrary.Float(89.45),
			IDSensor:         unifieddatalibrary.String("SENSOR-ID"),
			Luf:              unifieddatalibrary.Float(1.1),
			Md:               unifieddatalibrary.Float(1.1),
			Mufd:             unifieddatalibrary.Float(1.1),
			NeProfileName:    unifieddatalibrary.String("NH"),
			NeProfileVersion: unifieddatalibrary.Float(4.32),
			Origin:           unifieddatalibrary.String("THIRD_PARTY_DATASOURCE"),
			OrigSensorID:     unifieddatalibrary.String("ORIGSENSOR-ID"),
			Phase: unifieddatalibrary.IonOobservationUnvalidatedPublishParamsBodyPhase{
				Data:          [][][][][][][]float64{{{{{{{0.02, 0.034}}}}}}},
				DimensionName: []string{"NAME1", "NAME2"},
				Dimensions:    []int64{1, 2},
				Notes:         unifieddatalibrary.String("NOTES"),
			},
			PlasmaFrequency:            []float64{1.1, 2.1, 3.1},
			PlasmaFrequencyUncertainty: []float64{0.8, 0.2, 0.5},
			PlatformName:               unifieddatalibrary.String("Millstone Hill"),
			Polarization: unifieddatalibrary.IonOobservationUnvalidatedPublishParamsBodyPolarization{
				Data:          [][][][][][][]string{{{{{{{"X", "O"}}}}}}},
				DimensionName: []string{"NAME1", "NAME2"},
				Dimensions:    []int64{1, 2},
				Notes:         unifieddatalibrary.String("NOTES"),
			},
			Power: unifieddatalibrary.IonOobservationUnvalidatedPublishParamsBodyPower{
				Data:          [][][][][][][]float64{{{{{{{0.02, 0.034}}}}}}},
				DimensionName: []string{"NAME1", "NAME2"},
				Dimensions:    []int64{1, 2},
				Notes:         unifieddatalibrary.String("NOTES"),
			},
			Qe: unifieddatalibrary.Float(0.95),
			Qf: unifieddatalibrary.Float(1.83),
			Range: unifieddatalibrary.IonOobservationUnvalidatedPublishParamsBodyRange{
				Data:          [][][][][][][]float64{{{{{{{0.02, 0.034}}}}}}},
				DimensionName: []string{"NAME1", "NAME2"},
				Dimensions:    []int64{1, 2},
				Notes:         unifieddatalibrary.String("NOTES"),
			},
			ReceiveCoordinates:       [][]float64{{45.5, 179.3, 35.6}, {-80.2, -20.5, 43.2}},
			ReceiveSensorType:        "Mobile",
			RestrictedFrequency:      []float64{12.5, 34.5, 45.3},
			RestrictedFrequencyNotes: unifieddatalibrary.String("Example notes"),
			ScaleHeightF2Peak:        unifieddatalibrary.Float(35.613),
			ScalerInfo: unifieddatalibrary.IonOobservationUnvalidatedPublishParamsBodyScalerInfo{
				ConfidenceLevel: unifieddatalibrary.Int(11),
				ConfidenceScore: unifieddatalibrary.Int(75),
				Name:            unifieddatalibrary.String("ARTIST-4"),
				Organization:    unifieddatalibrary.String("UML"),
				Type:            unifieddatalibrary.String("MANUAL"),
				Version:         unifieddatalibrary.Float(500200.1),
			},
			Stokes: unifieddatalibrary.IonOobservationUnvalidatedPublishParamsBodyStokes{
				Data:          [][][][][][][]float64{{{{{{{0.02, 0.034}}}}}}},
				DimensionName: []string{"NAMES1", "NAMES2"},
				Dimensions:    []int64{2, 3},
				Notes:         unifieddatalibrary.String("NOTES"),
				S:             []float64{1, 2},
			},
			SystemNotes:    unifieddatalibrary.String("Example notes"),
			Tec:            unifieddatalibrary.Float(24.673),
			TidAzimuth:     []float64{1.1, 2.1, 3.1},
			TidPeriods:     []float64{1.1, 2.1, 3.1},
			TidPhaseSpeeds: []float64{1.1, 2.1, 3.1},
			Time: unifieddatalibrary.IonOobservationUnvalidatedPublishParamsBodyTime{
				Data:          [][][][][][][]float64{{{{{{{0.02, 0.034}}}}}}},
				DimensionName: []string{"NAME1", "NAME2"},
				Dimensions:    []int64{1, 2},
				Notes:         unifieddatalibrary.String("NOTES"),
			},
			TraceGeneric: unifieddatalibrary.IonOobservationUnvalidatedPublishParamsBodyTraceGeneric{
				Data:          [][][]float64{{{1.23, 1.0903}}},
				DimensionName: []string{"NAME1", "NAME2"},
				Notes:         unifieddatalibrary.String("NOTES"),
			},
			TransmitCoordinates: [][]float64{{45.5, 179.3, 35.6}, {-80.2, -20.5, 43.2}},
			TransmitSensorType:  "Mobile",
			TypeEs:              unifieddatalibrary.String("Auroral"),
			YE:                  unifieddatalibrary.Float(1.722),
			YF1:                 unifieddatalibrary.Float(55.645),
			YF2:                 unifieddatalibrary.Float(62.178),
			ZhalfNm:             unifieddatalibrary.Float(240.498),
			ZmE:                 unifieddatalibrary.Float(91.744),
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
