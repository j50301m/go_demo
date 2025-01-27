// Code generated by ent, DO NOT EDIT.

package loginrecord

import (
	"hype-casino-platform/auth_service/internal/infrastructure/ent_impl/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldEQ(FieldUpdatedAt, v))
}

// Browser applies equality check predicate on the "browser" field. It's identical to BrowserEQ.
func Browser(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldEQ(FieldBrowser, v))
}

// BrowserVer applies equality check predicate on the "browser_ver" field. It's identical to BrowserVerEQ.
func BrowserVer(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldEQ(FieldBrowserVer, v))
}

// IP applies equality check predicate on the "ip" field. It's identical to IPEQ.
func IP(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldEQ(FieldIP, v))
}

// Os applies equality check predicate on the "os" field. It's identical to OsEQ.
func Os(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldEQ(FieldOs, v))
}

// Platform applies equality check predicate on the "platform" field. It's identical to PlatformEQ.
func Platform(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldEQ(FieldPlatform, v))
}

// Country applies equality check predicate on the "country" field. It's identical to CountryEQ.
func Country(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldEQ(FieldCountry, v))
}

// CountryCode applies equality check predicate on the "country_code" field. It's identical to CountryCodeEQ.
func CountryCode(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldEQ(FieldCountryCode, v))
}

// City applies equality check predicate on the "city" field. It's identical to CityEQ.
func City(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldEQ(FieldCity, v))
}

// Asp applies equality check predicate on the "asp" field. It's identical to AspEQ.
func Asp(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldEQ(FieldAsp, v))
}

// IsMobile applies equality check predicate on the "is_mobile" field. It's identical to IsMobileEQ.
func IsMobile(v bool) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldEQ(FieldIsMobile, v))
}

// IsSuccess applies equality check predicate on the "is_success" field. It's identical to IsSuccessEQ.
func IsSuccess(v bool) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldEQ(FieldIsSuccess, v))
}

// ErrMessage applies equality check predicate on the "err_message" field. It's identical to ErrMessageEQ.
func ErrMessage(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldEQ(FieldErrMessage, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldLTE(FieldUpdatedAt, v))
}

// BrowserEQ applies the EQ predicate on the "browser" field.
func BrowserEQ(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldEQ(FieldBrowser, v))
}

// BrowserNEQ applies the NEQ predicate on the "browser" field.
func BrowserNEQ(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldNEQ(FieldBrowser, v))
}

// BrowserIn applies the In predicate on the "browser" field.
func BrowserIn(vs ...string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldIn(FieldBrowser, vs...))
}

// BrowserNotIn applies the NotIn predicate on the "browser" field.
func BrowserNotIn(vs ...string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldNotIn(FieldBrowser, vs...))
}

// BrowserGT applies the GT predicate on the "browser" field.
func BrowserGT(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldGT(FieldBrowser, v))
}

// BrowserGTE applies the GTE predicate on the "browser" field.
func BrowserGTE(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldGTE(FieldBrowser, v))
}

// BrowserLT applies the LT predicate on the "browser" field.
func BrowserLT(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldLT(FieldBrowser, v))
}

// BrowserLTE applies the LTE predicate on the "browser" field.
func BrowserLTE(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldLTE(FieldBrowser, v))
}

// BrowserContains applies the Contains predicate on the "browser" field.
func BrowserContains(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldContains(FieldBrowser, v))
}

// BrowserHasPrefix applies the HasPrefix predicate on the "browser" field.
func BrowserHasPrefix(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldHasPrefix(FieldBrowser, v))
}

// BrowserHasSuffix applies the HasSuffix predicate on the "browser" field.
func BrowserHasSuffix(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldHasSuffix(FieldBrowser, v))
}

// BrowserEqualFold applies the EqualFold predicate on the "browser" field.
func BrowserEqualFold(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldEqualFold(FieldBrowser, v))
}

// BrowserContainsFold applies the ContainsFold predicate on the "browser" field.
func BrowserContainsFold(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldContainsFold(FieldBrowser, v))
}

// BrowserVerEQ applies the EQ predicate on the "browser_ver" field.
func BrowserVerEQ(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldEQ(FieldBrowserVer, v))
}

// BrowserVerNEQ applies the NEQ predicate on the "browser_ver" field.
func BrowserVerNEQ(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldNEQ(FieldBrowserVer, v))
}

// BrowserVerIn applies the In predicate on the "browser_ver" field.
func BrowserVerIn(vs ...string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldIn(FieldBrowserVer, vs...))
}

// BrowserVerNotIn applies the NotIn predicate on the "browser_ver" field.
func BrowserVerNotIn(vs ...string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldNotIn(FieldBrowserVer, vs...))
}

// BrowserVerGT applies the GT predicate on the "browser_ver" field.
func BrowserVerGT(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldGT(FieldBrowserVer, v))
}

// BrowserVerGTE applies the GTE predicate on the "browser_ver" field.
func BrowserVerGTE(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldGTE(FieldBrowserVer, v))
}

// BrowserVerLT applies the LT predicate on the "browser_ver" field.
func BrowserVerLT(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldLT(FieldBrowserVer, v))
}

// BrowserVerLTE applies the LTE predicate on the "browser_ver" field.
func BrowserVerLTE(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldLTE(FieldBrowserVer, v))
}

// BrowserVerContains applies the Contains predicate on the "browser_ver" field.
func BrowserVerContains(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldContains(FieldBrowserVer, v))
}

// BrowserVerHasPrefix applies the HasPrefix predicate on the "browser_ver" field.
func BrowserVerHasPrefix(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldHasPrefix(FieldBrowserVer, v))
}

// BrowserVerHasSuffix applies the HasSuffix predicate on the "browser_ver" field.
func BrowserVerHasSuffix(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldHasSuffix(FieldBrowserVer, v))
}

// BrowserVerEqualFold applies the EqualFold predicate on the "browser_ver" field.
func BrowserVerEqualFold(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldEqualFold(FieldBrowserVer, v))
}

// BrowserVerContainsFold applies the ContainsFold predicate on the "browser_ver" field.
func BrowserVerContainsFold(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldContainsFold(FieldBrowserVer, v))
}

// IPEQ applies the EQ predicate on the "ip" field.
func IPEQ(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldEQ(FieldIP, v))
}

// IPNEQ applies the NEQ predicate on the "ip" field.
func IPNEQ(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldNEQ(FieldIP, v))
}

// IPIn applies the In predicate on the "ip" field.
func IPIn(vs ...string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldIn(FieldIP, vs...))
}

// IPNotIn applies the NotIn predicate on the "ip" field.
func IPNotIn(vs ...string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldNotIn(FieldIP, vs...))
}

// IPGT applies the GT predicate on the "ip" field.
func IPGT(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldGT(FieldIP, v))
}

// IPGTE applies the GTE predicate on the "ip" field.
func IPGTE(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldGTE(FieldIP, v))
}

// IPLT applies the LT predicate on the "ip" field.
func IPLT(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldLT(FieldIP, v))
}

// IPLTE applies the LTE predicate on the "ip" field.
func IPLTE(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldLTE(FieldIP, v))
}

// IPContains applies the Contains predicate on the "ip" field.
func IPContains(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldContains(FieldIP, v))
}

// IPHasPrefix applies the HasPrefix predicate on the "ip" field.
func IPHasPrefix(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldHasPrefix(FieldIP, v))
}

// IPHasSuffix applies the HasSuffix predicate on the "ip" field.
func IPHasSuffix(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldHasSuffix(FieldIP, v))
}

// IPEqualFold applies the EqualFold predicate on the "ip" field.
func IPEqualFold(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldEqualFold(FieldIP, v))
}

// IPContainsFold applies the ContainsFold predicate on the "ip" field.
func IPContainsFold(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldContainsFold(FieldIP, v))
}

// OsEQ applies the EQ predicate on the "os" field.
func OsEQ(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldEQ(FieldOs, v))
}

// OsNEQ applies the NEQ predicate on the "os" field.
func OsNEQ(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldNEQ(FieldOs, v))
}

// OsIn applies the In predicate on the "os" field.
func OsIn(vs ...string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldIn(FieldOs, vs...))
}

// OsNotIn applies the NotIn predicate on the "os" field.
func OsNotIn(vs ...string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldNotIn(FieldOs, vs...))
}

// OsGT applies the GT predicate on the "os" field.
func OsGT(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldGT(FieldOs, v))
}

// OsGTE applies the GTE predicate on the "os" field.
func OsGTE(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldGTE(FieldOs, v))
}

// OsLT applies the LT predicate on the "os" field.
func OsLT(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldLT(FieldOs, v))
}

// OsLTE applies the LTE predicate on the "os" field.
func OsLTE(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldLTE(FieldOs, v))
}

// OsContains applies the Contains predicate on the "os" field.
func OsContains(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldContains(FieldOs, v))
}

// OsHasPrefix applies the HasPrefix predicate on the "os" field.
func OsHasPrefix(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldHasPrefix(FieldOs, v))
}

// OsHasSuffix applies the HasSuffix predicate on the "os" field.
func OsHasSuffix(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldHasSuffix(FieldOs, v))
}

// OsEqualFold applies the EqualFold predicate on the "os" field.
func OsEqualFold(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldEqualFold(FieldOs, v))
}

// OsContainsFold applies the ContainsFold predicate on the "os" field.
func OsContainsFold(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldContainsFold(FieldOs, v))
}

// PlatformEQ applies the EQ predicate on the "platform" field.
func PlatformEQ(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldEQ(FieldPlatform, v))
}

// PlatformNEQ applies the NEQ predicate on the "platform" field.
func PlatformNEQ(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldNEQ(FieldPlatform, v))
}

// PlatformIn applies the In predicate on the "platform" field.
func PlatformIn(vs ...string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldIn(FieldPlatform, vs...))
}

// PlatformNotIn applies the NotIn predicate on the "platform" field.
func PlatformNotIn(vs ...string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldNotIn(FieldPlatform, vs...))
}

// PlatformGT applies the GT predicate on the "platform" field.
func PlatformGT(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldGT(FieldPlatform, v))
}

// PlatformGTE applies the GTE predicate on the "platform" field.
func PlatformGTE(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldGTE(FieldPlatform, v))
}

// PlatformLT applies the LT predicate on the "platform" field.
func PlatformLT(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldLT(FieldPlatform, v))
}

// PlatformLTE applies the LTE predicate on the "platform" field.
func PlatformLTE(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldLTE(FieldPlatform, v))
}

// PlatformContains applies the Contains predicate on the "platform" field.
func PlatformContains(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldContains(FieldPlatform, v))
}

// PlatformHasPrefix applies the HasPrefix predicate on the "platform" field.
func PlatformHasPrefix(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldHasPrefix(FieldPlatform, v))
}

// PlatformHasSuffix applies the HasSuffix predicate on the "platform" field.
func PlatformHasSuffix(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldHasSuffix(FieldPlatform, v))
}

// PlatformEqualFold applies the EqualFold predicate on the "platform" field.
func PlatformEqualFold(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldEqualFold(FieldPlatform, v))
}

// PlatformContainsFold applies the ContainsFold predicate on the "platform" field.
func PlatformContainsFold(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldContainsFold(FieldPlatform, v))
}

// CountryEQ applies the EQ predicate on the "country" field.
func CountryEQ(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldEQ(FieldCountry, v))
}

// CountryNEQ applies the NEQ predicate on the "country" field.
func CountryNEQ(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldNEQ(FieldCountry, v))
}

// CountryIn applies the In predicate on the "country" field.
func CountryIn(vs ...string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldIn(FieldCountry, vs...))
}

// CountryNotIn applies the NotIn predicate on the "country" field.
func CountryNotIn(vs ...string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldNotIn(FieldCountry, vs...))
}

// CountryGT applies the GT predicate on the "country" field.
func CountryGT(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldGT(FieldCountry, v))
}

// CountryGTE applies the GTE predicate on the "country" field.
func CountryGTE(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldGTE(FieldCountry, v))
}

// CountryLT applies the LT predicate on the "country" field.
func CountryLT(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldLT(FieldCountry, v))
}

// CountryLTE applies the LTE predicate on the "country" field.
func CountryLTE(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldLTE(FieldCountry, v))
}

// CountryContains applies the Contains predicate on the "country" field.
func CountryContains(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldContains(FieldCountry, v))
}

// CountryHasPrefix applies the HasPrefix predicate on the "country" field.
func CountryHasPrefix(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldHasPrefix(FieldCountry, v))
}

// CountryHasSuffix applies the HasSuffix predicate on the "country" field.
func CountryHasSuffix(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldHasSuffix(FieldCountry, v))
}

// CountryEqualFold applies the EqualFold predicate on the "country" field.
func CountryEqualFold(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldEqualFold(FieldCountry, v))
}

// CountryContainsFold applies the ContainsFold predicate on the "country" field.
func CountryContainsFold(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldContainsFold(FieldCountry, v))
}

// CountryCodeEQ applies the EQ predicate on the "country_code" field.
func CountryCodeEQ(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldEQ(FieldCountryCode, v))
}

// CountryCodeNEQ applies the NEQ predicate on the "country_code" field.
func CountryCodeNEQ(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldNEQ(FieldCountryCode, v))
}

// CountryCodeIn applies the In predicate on the "country_code" field.
func CountryCodeIn(vs ...string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldIn(FieldCountryCode, vs...))
}

// CountryCodeNotIn applies the NotIn predicate on the "country_code" field.
func CountryCodeNotIn(vs ...string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldNotIn(FieldCountryCode, vs...))
}

// CountryCodeGT applies the GT predicate on the "country_code" field.
func CountryCodeGT(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldGT(FieldCountryCode, v))
}

// CountryCodeGTE applies the GTE predicate on the "country_code" field.
func CountryCodeGTE(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldGTE(FieldCountryCode, v))
}

// CountryCodeLT applies the LT predicate on the "country_code" field.
func CountryCodeLT(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldLT(FieldCountryCode, v))
}

// CountryCodeLTE applies the LTE predicate on the "country_code" field.
func CountryCodeLTE(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldLTE(FieldCountryCode, v))
}

// CountryCodeContains applies the Contains predicate on the "country_code" field.
func CountryCodeContains(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldContains(FieldCountryCode, v))
}

// CountryCodeHasPrefix applies the HasPrefix predicate on the "country_code" field.
func CountryCodeHasPrefix(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldHasPrefix(FieldCountryCode, v))
}

// CountryCodeHasSuffix applies the HasSuffix predicate on the "country_code" field.
func CountryCodeHasSuffix(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldHasSuffix(FieldCountryCode, v))
}

// CountryCodeEqualFold applies the EqualFold predicate on the "country_code" field.
func CountryCodeEqualFold(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldEqualFold(FieldCountryCode, v))
}

// CountryCodeContainsFold applies the ContainsFold predicate on the "country_code" field.
func CountryCodeContainsFold(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldContainsFold(FieldCountryCode, v))
}

// CityEQ applies the EQ predicate on the "city" field.
func CityEQ(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldEQ(FieldCity, v))
}

// CityNEQ applies the NEQ predicate on the "city" field.
func CityNEQ(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldNEQ(FieldCity, v))
}

// CityIn applies the In predicate on the "city" field.
func CityIn(vs ...string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldIn(FieldCity, vs...))
}

// CityNotIn applies the NotIn predicate on the "city" field.
func CityNotIn(vs ...string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldNotIn(FieldCity, vs...))
}

// CityGT applies the GT predicate on the "city" field.
func CityGT(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldGT(FieldCity, v))
}

// CityGTE applies the GTE predicate on the "city" field.
func CityGTE(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldGTE(FieldCity, v))
}

// CityLT applies the LT predicate on the "city" field.
func CityLT(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldLT(FieldCity, v))
}

// CityLTE applies the LTE predicate on the "city" field.
func CityLTE(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldLTE(FieldCity, v))
}

// CityContains applies the Contains predicate on the "city" field.
func CityContains(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldContains(FieldCity, v))
}

// CityHasPrefix applies the HasPrefix predicate on the "city" field.
func CityHasPrefix(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldHasPrefix(FieldCity, v))
}

// CityHasSuffix applies the HasSuffix predicate on the "city" field.
func CityHasSuffix(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldHasSuffix(FieldCity, v))
}

// CityEqualFold applies the EqualFold predicate on the "city" field.
func CityEqualFold(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldEqualFold(FieldCity, v))
}

// CityContainsFold applies the ContainsFold predicate on the "city" field.
func CityContainsFold(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldContainsFold(FieldCity, v))
}

// AspEQ applies the EQ predicate on the "asp" field.
func AspEQ(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldEQ(FieldAsp, v))
}

// AspNEQ applies the NEQ predicate on the "asp" field.
func AspNEQ(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldNEQ(FieldAsp, v))
}

// AspIn applies the In predicate on the "asp" field.
func AspIn(vs ...string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldIn(FieldAsp, vs...))
}

// AspNotIn applies the NotIn predicate on the "asp" field.
func AspNotIn(vs ...string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldNotIn(FieldAsp, vs...))
}

// AspGT applies the GT predicate on the "asp" field.
func AspGT(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldGT(FieldAsp, v))
}

// AspGTE applies the GTE predicate on the "asp" field.
func AspGTE(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldGTE(FieldAsp, v))
}

// AspLT applies the LT predicate on the "asp" field.
func AspLT(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldLT(FieldAsp, v))
}

// AspLTE applies the LTE predicate on the "asp" field.
func AspLTE(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldLTE(FieldAsp, v))
}

// AspContains applies the Contains predicate on the "asp" field.
func AspContains(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldContains(FieldAsp, v))
}

// AspHasPrefix applies the HasPrefix predicate on the "asp" field.
func AspHasPrefix(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldHasPrefix(FieldAsp, v))
}

// AspHasSuffix applies the HasSuffix predicate on the "asp" field.
func AspHasSuffix(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldHasSuffix(FieldAsp, v))
}

// AspEqualFold applies the EqualFold predicate on the "asp" field.
func AspEqualFold(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldEqualFold(FieldAsp, v))
}

// AspContainsFold applies the ContainsFold predicate on the "asp" field.
func AspContainsFold(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldContainsFold(FieldAsp, v))
}

// IsMobileEQ applies the EQ predicate on the "is_mobile" field.
func IsMobileEQ(v bool) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldEQ(FieldIsMobile, v))
}

// IsMobileNEQ applies the NEQ predicate on the "is_mobile" field.
func IsMobileNEQ(v bool) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldNEQ(FieldIsMobile, v))
}

// IsSuccessEQ applies the EQ predicate on the "is_success" field.
func IsSuccessEQ(v bool) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldEQ(FieldIsSuccess, v))
}

// IsSuccessNEQ applies the NEQ predicate on the "is_success" field.
func IsSuccessNEQ(v bool) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldNEQ(FieldIsSuccess, v))
}

// ErrMessageEQ applies the EQ predicate on the "err_message" field.
func ErrMessageEQ(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldEQ(FieldErrMessage, v))
}

// ErrMessageNEQ applies the NEQ predicate on the "err_message" field.
func ErrMessageNEQ(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldNEQ(FieldErrMessage, v))
}

// ErrMessageIn applies the In predicate on the "err_message" field.
func ErrMessageIn(vs ...string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldIn(FieldErrMessage, vs...))
}

// ErrMessageNotIn applies the NotIn predicate on the "err_message" field.
func ErrMessageNotIn(vs ...string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldNotIn(FieldErrMessage, vs...))
}

// ErrMessageGT applies the GT predicate on the "err_message" field.
func ErrMessageGT(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldGT(FieldErrMessage, v))
}

// ErrMessageGTE applies the GTE predicate on the "err_message" field.
func ErrMessageGTE(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldGTE(FieldErrMessage, v))
}

// ErrMessageLT applies the LT predicate on the "err_message" field.
func ErrMessageLT(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldLT(FieldErrMessage, v))
}

// ErrMessageLTE applies the LTE predicate on the "err_message" field.
func ErrMessageLTE(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldLTE(FieldErrMessage, v))
}

// ErrMessageContains applies the Contains predicate on the "err_message" field.
func ErrMessageContains(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldContains(FieldErrMessage, v))
}

// ErrMessageHasPrefix applies the HasPrefix predicate on the "err_message" field.
func ErrMessageHasPrefix(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldHasPrefix(FieldErrMessage, v))
}

// ErrMessageHasSuffix applies the HasSuffix predicate on the "err_message" field.
func ErrMessageHasSuffix(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldHasSuffix(FieldErrMessage, v))
}

// ErrMessageEqualFold applies the EqualFold predicate on the "err_message" field.
func ErrMessageEqualFold(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldEqualFold(FieldErrMessage, v))
}

// ErrMessageContainsFold applies the ContainsFold predicate on the "err_message" field.
func ErrMessageContainsFold(v string) predicate.LoginRecord {
	return predicate.LoginRecord(sql.FieldContainsFold(FieldErrMessage, v))
}

// HasUsers applies the HasEdge predicate on the "users" edge.
func HasUsers() predicate.LoginRecord {
	return predicate.LoginRecord(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UsersTable, UsersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUsersWith applies the HasEdge predicate on the "users" edge with a given conditions (other predicates).
func HasUsersWith(preds ...predicate.User) predicate.LoginRecord {
	return predicate.LoginRecord(func(s *sql.Selector) {
		step := newUsersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.LoginRecord) predicate.LoginRecord {
	return predicate.LoginRecord(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.LoginRecord) predicate.LoginRecord {
	return predicate.LoginRecord(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.LoginRecord) predicate.LoginRecord {
	return predicate.LoginRecord(sql.NotPredicates(p))
}
