// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package unifieddatalibrary

import (
	"context"
	"net/http"

	"github.com/Bluestaq/udl-golang-sdk/internal/apijson"
	"github.com/Bluestaq/udl-golang-sdk/internal/requestconfig"
	"github.com/Bluestaq/udl-golang-sdk/option"
	"github.com/Bluestaq/udl-golang-sdk/packages/respjson"
)

// SupportingDataQueryHelpService contains methods and other services that help
// with interacting with the unifieddatalibrary API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSupportingDataQueryHelpService] method instead.
type SupportingDataQueryHelpService struct {
	Options []option.RequestOption
}

// NewSupportingDataQueryHelpService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSupportingDataQueryHelpService(opts ...option.RequestOption) (r SupportingDataQueryHelpService) {
	r = SupportingDataQueryHelpService{}
	r.Options = opts
	return
}

// Service operation to provide detailed information on available dynamic query
// parameters for a particular data type.
func (r *SupportingDataQueryHelpService) Get(ctx context.Context, opts ...option.RequestOption) (res *SupportingDataQueryHelpGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "udl/dataowner/queryhelp"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type SupportingDataQueryHelpGetResponse struct {
	AodrSupported         bool                                          `json:"aodrSupported"`
	ClassificationMarking string                                        `json:"classificationMarking"`
	Description           string                                        `json:"description"`
	HistorySupported      bool                                          `json:"historySupported"`
	Name                  string                                        `json:"name"`
	Parameters            []SupportingDataQueryHelpGetResponseParameter `json:"parameters"`
	RequiredRoles         []string                                      `json:"requiredRoles"`
	RestSupported         bool                                          `json:"restSupported"`
	SortSupported         bool                                          `json:"sortSupported"`
	TypeName              string                                        `json:"typeName"`
	Uri                   string                                        `json:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AodrSupported         respjson.Field
		ClassificationMarking respjson.Field
		Description           respjson.Field
		HistorySupported      respjson.Field
		Name                  respjson.Field
		Parameters            respjson.Field
		RequiredRoles         respjson.Field
		RestSupported         respjson.Field
		SortSupported         respjson.Field
		TypeName              respjson.Field
		Uri                   respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SupportingDataQueryHelpGetResponse) RawJSON() string { return r.JSON.raw }
func (r *SupportingDataQueryHelpGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SupportingDataQueryHelpGetResponseParameter struct {
	ClassificationMarking string `json:"classificationMarking"`
	Derived               bool   `json:"derived"`
	Description           string `json:"description"`
	ElemMatch             bool   `json:"elemMatch"`
	Format                string `json:"format"`
	HistQuerySupported    bool   `json:"histQuerySupported"`
	HistTupleSupported    bool   `json:"histTupleSupported"`
	Name                  string `json:"name"`
	Required              bool   `json:"required"`
	RestQuerySupported    bool   `json:"restQuerySupported"`
	RestTupleSupported    bool   `json:"restTupleSupported"`
	Type                  string `json:"type"`
	UnitOfMeasure         string `json:"unitOfMeasure"`
	UtcDate               bool   `json:"utcDate"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClassificationMarking respjson.Field
		Derived               respjson.Field
		Description           respjson.Field
		ElemMatch             respjson.Field
		Format                respjson.Field
		HistQuerySupported    respjson.Field
		HistTupleSupported    respjson.Field
		Name                  respjson.Field
		Required              respjson.Field
		RestQuerySupported    respjson.Field
		RestTupleSupported    respjson.Field
		Type                  respjson.Field
		UnitOfMeasure         respjson.Field
		UtcDate               respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SupportingDataQueryHelpGetResponseParameter) RawJSON() string { return r.JSON.raw }
func (r *SupportingDataQueryHelpGetResponseParameter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
