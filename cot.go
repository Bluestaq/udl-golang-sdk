// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"net/http"
	"time"

	"github.com/stainless-sdks/unifieddatalibrary-go/internal/requestconfig"
	"github.com/stainless-sdks/unifieddatalibrary-go/option"
	"github.com/stainless-sdks/unifieddatalibrary-go/packages/param"
)

// CotService contains methods and other services that help with interacting with
// the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCotService] method instead.
type CotService struct {
	Options []option.RequestOption
}

// NewCotService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCotService(opts ...option.RequestOption) (r CotService) {
	r = CotService{}
	r.Options = opts
	return
}

// This service enables posting CoT messages to the UDL TAK server. CoT data will
// be persisted in the UDL POI schema as well as federated to connected TAK
// servers.
func (r *CotService) New(ctx context.Context, body CotNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "udl/cot"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type CotNewParams struct {
	// WGS-84 latitude of the POI, in degrees (+N, -S), -90 to 90.
	Lat float64 `json:"lat,required"`
	// WGS-84 longitude of the POI, in degrees (+E, -W), -180 to 180.
	Lon float64 `json:"lon,required"`
	// Point height above ellipsoid (WGS-84), in meters.
	Alt param.Opt[float64] `json:"alt,omitzero"`
	// Radius of circular area about lat/lon point, in meters (1-sigma, if representing
	// error).
	Ce param.Opt[float64] `json:"ce,omitzero"`
	// How the event point was generated, in CoT object heirarchy notation (optional,
	// CoT).
	How param.Opt[string] `json:"how,omitzero"`
	// Height above lat/lon point, in meters (1-sigma, if representing linear error).
	Le param.Opt[float64] `json:"le,omitzero"`
	// Identifier of the sender of the cot message which should remain the same on
	// subsequent POI records of the same point of interest.
	SenderUid param.Opt[string] `json:"senderUid,omitzero"`
	// Stale timestamp (optional), in ISO8601 UTC format.
	Stale param.Opt[time.Time] `json:"stale,omitzero" format:"date-time"`
	// Start time of event validity (optional), in ISO8601 UTC format.
	Start param.Opt[time.Time] `json:"start,omitzero" format:"date-time"`
	// Event type, in CoT object heirarchy notation (optional, CoT).
	Type param.Opt[string] `json:"type,omitzero"`
	// Optional list of call signs to send message to directly.
	CallSigns []string `json:"callSigns,omitzero"`
	// Schema for the CotChatData to post.
	CotChatData CotNewParamsCotChatData `json:"cotChatData,omitzero"`
	// Schema for the CotPositionData to post.
	CotPositionData CotNewParamsCotPositionData `json:"cotPositionData,omitzero"`
	// Optional set of groups to send message to specifically. If not specified, the
	// message will be sent to the default _ANON_ group.
	Groups []string `json:"groups,omitzero"`
	// Optional list of TAK user ids to send message to directly.
	Uids []string `json:"uids,omitzero"`
	paramObj
}

func (r CotNewParams) MarshalJSON() (data []byte, err error) {
	type shadow CotNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// Schema for the CotChatData to post.
type CotNewParamsCotChatData struct {
	// Contents of a chat message.
	ChatMsg param.Opt[string] `json:"chatMsg,omitzero"`
	// Chat room name to send chat message to.
	ChatRoom param.Opt[string] `json:"chatRoom,omitzero"`
	// Callsign of chat sender.
	ChatSenderCallSign param.Opt[string] `json:"chatSenderCallSign,omitzero"`
	paramObj
}

func (r CotNewParamsCotChatData) MarshalJSON() (data []byte, err error) {
	type shadow CotNewParamsCotChatData
	return param.MarshalObject(r, (*shadow)(&r))
}

// Schema for the CotPositionData to post.
//
// The properties CallSign, Team, TeamRole are required.
type CotNewParamsCotPositionData struct {
	// Name of the POI target Object.
	CallSign string `json:"callSign,required"`
	// Description of the POI target Object.
	Team string `json:"team,required"`
	// Team role (Team Member| Team Lead | HQ | Sniper | Medic | Forward Observer | RTO
	// | K9).
	TeamRole string `json:"teamRole,required"`
	paramObj
}

func (r CotNewParamsCotPositionData) MarshalJSON() (data []byte, err error) {
	type shadow CotNewParamsCotPositionData
	return param.MarshalObject(r, (*shadow)(&r))
}
