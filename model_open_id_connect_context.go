/*
 * Ory Oathkeeper API
 *
 * Documentation for all of Ory Oathkeeper's APIs. 
 *
 * API version: v1.11.7
 * Contact: hi@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OpenIDConnectContext struct for OpenIDConnectContext
type OpenIDConnectContext struct {
	// ACRValues is the Authentication AuthorizationContext Class Reference requested in the OAuth 2.0 Authorization request. It is a parameter defined by OpenID Connect and expresses which level of authentication (e.g. 2FA) is required.  OpenID Connect defines it as follows: > Requested Authentication AuthorizationContext Class Reference values. Space-separated string that specifies the acr values that the Authorization Server is being requested to use for processing this Authentication Request, with the values appearing in order of preference. The Authentication AuthorizationContext Class satisfied by the authentication performed is returned as the acr Claim Value, as specified in Section 2. The acr Claim is requested as a Voluntary Claim by this parameter.
	AcrValues []string `json:"acr_values,omitempty"`
	// Display is a string value that specifies how the Authorization Server displays the authentication and consent user interface pages to the End-User. The defined values are: page: The Authorization Server SHOULD display the authentication and consent UI consistent with a full User Agent page view. If the display parameter is not specified, this is the default display mode. popup: The Authorization Server SHOULD display the authentication and consent UI consistent with a popup User Agent window. The popup User Agent window should be of an appropriate size for a login-focused dialog and should not obscure the entire window that it is popping up over. touch: The Authorization Server SHOULD display the authentication and consent UI consistent with a device that leverages a touch interface. wap: The Authorization Server SHOULD display the authentication and consent UI consistent with a \"feature phone\" type display.  The Authorization Server MAY also attempt to detect the capabilities of the User Agent and present an appropriate display.
	Display *string `json:"display,omitempty"`
	// IDTokenHintClaims are the claims of the ID Token previously issued by the Authorization Server being passed as a hint about the End-User's current or past authenticated session with the Client.
	IdTokenHintClaims map[string]map[string]interface{} `json:"id_token_hint_claims,omitempty"`
	// LoginHint hints about the login identifier the End-User might use to log in (if necessary). This hint can be used by an RP if it first asks the End-User for their e-mail address (or other identifier) and then wants to pass that value as a hint to the discovered authorization service. This value MAY also be a phone number in the format specified for the phone_number Claim. The use of this parameter is optional.
	LoginHint *string `json:"login_hint,omitempty"`
	// UILocales is the End-User'id preferred languages and scripts for the user interface, represented as a space-separated list of BCP47 [RFC5646] language tag values, ordered by preference. For instance, the value \"fr-CA fr en\" represents a preference for French as spoken in Canada, then French (without a region designation), followed by English (without a region designation). An error SHOULD NOT result if some or all of the requested locales are not supported by the OpenID Provider.
	UiLocales []string `json:"ui_locales,omitempty"`
}

// NewOpenIDConnectContext instantiates a new OpenIDConnectContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenIDConnectContext() *OpenIDConnectContext {
	this := OpenIDConnectContext{}
	return &this
}

// NewOpenIDConnectContextWithDefaults instantiates a new OpenIDConnectContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenIDConnectContextWithDefaults() *OpenIDConnectContext {
	this := OpenIDConnectContext{}
	return &this
}

// GetAcrValues returns the AcrValues field value if set, zero value otherwise.
func (o *OpenIDConnectContext) GetAcrValues() []string {
	if o == nil || o.AcrValues == nil {
		var ret []string
		return ret
	}
	return o.AcrValues
}

// GetAcrValuesOk returns a tuple with the AcrValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIDConnectContext) GetAcrValuesOk() ([]string, bool) {
	if o == nil || o.AcrValues == nil {
		return nil, false
	}
	return o.AcrValues, true
}

// HasAcrValues returns a boolean if a field has been set.
func (o *OpenIDConnectContext) HasAcrValues() bool {
	if o != nil && o.AcrValues != nil {
		return true
	}

	return false
}

// SetAcrValues gets a reference to the given []string and assigns it to the AcrValues field.
func (o *OpenIDConnectContext) SetAcrValues(v []string) {
	o.AcrValues = v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *OpenIDConnectContext) GetDisplay() string {
	if o == nil || o.Display == nil {
		var ret string
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIDConnectContext) GetDisplayOk() (*string, bool) {
	if o == nil || o.Display == nil {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *OpenIDConnectContext) HasDisplay() bool {
	if o != nil && o.Display != nil {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given string and assigns it to the Display field.
func (o *OpenIDConnectContext) SetDisplay(v string) {
	o.Display = &v
}

// GetIdTokenHintClaims returns the IdTokenHintClaims field value if set, zero value otherwise.
func (o *OpenIDConnectContext) GetIdTokenHintClaims() map[string]map[string]interface{} {
	if o == nil || o.IdTokenHintClaims == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.IdTokenHintClaims
}

// GetIdTokenHintClaimsOk returns a tuple with the IdTokenHintClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIDConnectContext) GetIdTokenHintClaimsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.IdTokenHintClaims == nil {
		return nil, false
	}
	return o.IdTokenHintClaims, true
}

// HasIdTokenHintClaims returns a boolean if a field has been set.
func (o *OpenIDConnectContext) HasIdTokenHintClaims() bool {
	if o != nil && o.IdTokenHintClaims != nil {
		return true
	}

	return false
}

// SetIdTokenHintClaims gets a reference to the given map[string]map[string]interface{} and assigns it to the IdTokenHintClaims field.
func (o *OpenIDConnectContext) SetIdTokenHintClaims(v map[string]map[string]interface{}) {
	o.IdTokenHintClaims = v
}

// GetLoginHint returns the LoginHint field value if set, zero value otherwise.
func (o *OpenIDConnectContext) GetLoginHint() string {
	if o == nil || o.LoginHint == nil {
		var ret string
		return ret
	}
	return *o.LoginHint
}

// GetLoginHintOk returns a tuple with the LoginHint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIDConnectContext) GetLoginHintOk() (*string, bool) {
	if o == nil || o.LoginHint == nil {
		return nil, false
	}
	return o.LoginHint, true
}

// HasLoginHint returns a boolean if a field has been set.
func (o *OpenIDConnectContext) HasLoginHint() bool {
	if o != nil && o.LoginHint != nil {
		return true
	}

	return false
}

// SetLoginHint gets a reference to the given string and assigns it to the LoginHint field.
func (o *OpenIDConnectContext) SetLoginHint(v string) {
	o.LoginHint = &v
}

// GetUiLocales returns the UiLocales field value if set, zero value otherwise.
func (o *OpenIDConnectContext) GetUiLocales() []string {
	if o == nil || o.UiLocales == nil {
		var ret []string
		return ret
	}
	return o.UiLocales
}

// GetUiLocalesOk returns a tuple with the UiLocales field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIDConnectContext) GetUiLocalesOk() ([]string, bool) {
	if o == nil || o.UiLocales == nil {
		return nil, false
	}
	return o.UiLocales, true
}

// HasUiLocales returns a boolean if a field has been set.
func (o *OpenIDConnectContext) HasUiLocales() bool {
	if o != nil && o.UiLocales != nil {
		return true
	}

	return false
}

// SetUiLocales gets a reference to the given []string and assigns it to the UiLocales field.
func (o *OpenIDConnectContext) SetUiLocales(v []string) {
	o.UiLocales = v
}

func (o OpenIDConnectContext) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AcrValues != nil {
		toSerialize["acr_values"] = o.AcrValues
	}
	if o.Display != nil {
		toSerialize["display"] = o.Display
	}
	if o.IdTokenHintClaims != nil {
		toSerialize["id_token_hint_claims"] = o.IdTokenHintClaims
	}
	if o.LoginHint != nil {
		toSerialize["login_hint"] = o.LoginHint
	}
	if o.UiLocales != nil {
		toSerialize["ui_locales"] = o.UiLocales
	}
	return json.Marshal(toSerialize)
}

type NullableOpenIDConnectContext struct {
	value *OpenIDConnectContext
	isSet bool
}

func (v NullableOpenIDConnectContext) Get() *OpenIDConnectContext {
	return v.value
}

func (v *NullableOpenIDConnectContext) Set(val *OpenIDConnectContext) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenIDConnectContext) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenIDConnectContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenIDConnectContext(val *OpenIDConnectContext) *NullableOpenIDConnectContext {
	return &NullableOpenIDConnectContext{value: val, isSet: true}
}

func (v NullableOpenIDConnectContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenIDConnectContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


