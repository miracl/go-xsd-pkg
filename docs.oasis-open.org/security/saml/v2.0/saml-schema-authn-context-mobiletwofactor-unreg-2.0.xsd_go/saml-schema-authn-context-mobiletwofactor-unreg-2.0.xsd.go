// Package goSamlAuthnContextMobiletwofactorUnreg20 : autogenerated from XSD schema and manually adjusted (Nicola Asuni - 2016-11-03)
package goSamlAuthnContextMobiletwofactorUnreg20

import (
	sac "github.com/miracl/go-xsd-pkg/docs.oasis-open.org/security/saml/v2.0/saml-schema-authn-context-2.0.xsd_go"
	xsdt "github.com/miracl/go-xsd-pkg/xsdt"
)

// XAttrIDXsdtID defines attribute Id
type XAttrIDXsdtID struct {
	ID xsdt.ID `xml:"ID,attr,omitempty"`
}

// TAuthnContextDeclarationBaseType defines type AuthnContextDeclarationBaseType
type TAuthnContextDeclarationBaseType struct {
	sac.XElemAuthnMethod
	sac.XElemGoverningAgreements
	sac.XElemsExtension
	XAttrIDXsdtID
	*TAuthnContextDeclarationBaseType
	sac.XElemIdentification
	sac.XElemTechnicalProtection
	sac.XElemOperationalProtection
}

// Walk : if the WalkHandlers.TAuthnContextDeclarationBaseType function is not nil (ie. was set by outside code), calls it with this TAuthnContextDeclarationBaseType instance as the single argument. Then calls the Walk() method on 1/8 embed(s) and 0/0 field(s) belonging to this TAuthnContextDeclarationBaseType instance.
func (me *TAuthnContextDeclarationBaseType) Walk() (err error) {
	if fn := WalkHandlers.TAuthnContextDeclarationBaseType; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if err = me.TAuthnContextDeclarationBaseType.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
			return
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

// TAuthnMethodBaseType defines type AuthnMethodBaseType
type TAuthnMethodBaseType struct {
	*TAuthnMethodBaseType
	sac.XElemPrincipalAuthenticationMechanism
	sac.XElemAuthenticator
	sac.XElemAuthenticatorTransportProtocol
	sac.XElemsExtension
}

// Walk : if the WalkHandlers.TAuthnMethodBaseType function is not nil (ie. was set by outside code), calls it with this TAuthnMethodBaseType instance as the single argument. Then calls the Walk() method on 1/5 embed(s) and 0/0 field(s) belonging to this TAuthnMethodBaseType instance.
func (me *TAuthnMethodBaseType) Walk() (err error) {
	if fn := WalkHandlers.TAuthnMethodBaseType; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if err = me.TAuthnMethodBaseType.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
			return
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

// TechnicalProtectionBaseType defines type echnicalProtectionBaseType
type TechnicalProtectionBaseType struct {
	*TechnicalProtectionBaseType
	sac.XElemPrivateKeyProtection
	sac.XElemSecretKeyProtection
	sac.XElemsExtension
}

// Walk : if the WalkHandlers.TechnicalProtectionBaseType function is not nil (ie. was set by outside code), calls it with this TechnicalProtectionBaseType instance as the single argument. Then calls the Walk() method on 1/4 embed(s) and 0/0 field(s) belonging to this TechnicalProtectionBaseType instance.
func (me *TechnicalProtectionBaseType) Walk() (err error) {
	if fn := WalkHandlers.TechnicalProtectionBaseType; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if err = me.TechnicalProtectionBaseType.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
			return
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

// TxsdRedefineKeyStorageTypeComplexContentRestrictionMedium defines type xsdRedefineKeyStorageTypeComplexContentRestrictionMedium
type TxsdRedefineKeyStorageTypeComplexContentRestrictionMedium sac.TmediumType

// IsMobileAuthCard : Returns true if the value of this enumerated TxsdRedefineKeyStorageTypeComplexContentRestrictionMedium is "MobileAuthCard".
func (me TxsdRedefineKeyStorageTypeComplexContentRestrictionMedium) IsMobileAuthCard() bool {
	return me.String() == "MobileAuthCard"
}

// IsSmartcard : Returns true if the value of this enumerated TxsdRedefineKeyStorageTypeComplexContentRestrictionMedium is "smartcard".
func (me TxsdRedefineKeyStorageTypeComplexContentRestrictionMedium) IsSmartcard() bool {
	return me.String() == "smartcard"
}

// Set : Since TxsdRedefineKeyStorageTypeComplexContentRestrictionMedium is just a simple String type, this merely sets the current value from the specified string.
func (me *TxsdRedefineKeyStorageTypeComplexContentRestrictionMedium) Set(s string) {
	(*sac.TmediumType)(me).Set(s)
}

// String : Since TxsdRedefineKeyStorageTypeComplexContentRestrictionMedium is just a simple String type, this merely returns the current string value.
func (me TxsdRedefineKeyStorageTypeComplexContentRestrictionMedium) String() string {
	return sac.TmediumType(me).String()
}

// ToTmediumType : This convenience method just performs a simple type conversion to TxsdRedefineKeyStorageTypeComplexContentRestrictionMedium's alias type sac.TmediumType.
func (me TxsdRedefineKeyStorageTypeComplexContentRestrictionMedium) ToTmediumType() sac.TmediumType {
	return sac.TmediumType(me)
}

// IsMobileDevice : Returns true if the value of this enumerated TxsdRedefineKeyStorageTypeComplexContentRestrictionMedium is "MobileDevice".
func (me TxsdRedefineKeyStorageTypeComplexContentRestrictionMedium) IsMobileDevice() bool {
	return me.String() == "MobileDevice"
}

// XAttrMediumTxsdRedefineKeyStorageTypeComplexContentRestrictionMedium defines attribute MediumTxsdRedefineKeyStorageTypeComplexContentRestrictionMedium
type XAttrMediumTxsdRedefineKeyStorageTypeComplexContentRestrictionMedium struct {
	Medium TxsdRedefineKeyStorageTypeComplexContentRestrictionMedium `xml:"medium,attr,omitempty"`
}

// TKeyStorageType defines type KeyStorageType
type TKeyStorageType struct {
	XAttrMediumTxsdRedefineKeyStorageTypeComplexContentRestrictionMedium
	*TKeyStorageType
}

// Walk : if the WalkHandlers.TKeyStorageType function is not nil (ie. was set by outside code), calls it with this TKeyStorageType instance as the single argument. Then calls the Walk() method on 1/2 embed(s) and 0/0 field(s) belonging to this TKeyStorageType instance.
func (me *TKeyStorageType) Walk() (err error) {
	if fn := WalkHandlers.TKeyStorageType; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if err = me.TKeyStorageType.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
			return
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

// TSecurityAuditType defines type SecurityAuditType
type TSecurityAuditType struct {
	*TSecurityAuditType
	sac.XElemSwitchAudit
	sac.XElemsExtension
}

// Walk : if the WalkHandlers.TSecurityAuditType function is not nil (ie. was set by outside code), calls it with this TSecurityAuditType instance as the single argument. Then calls the Walk() method on 1/3 embed(s) and 0/0 field(s) belonging to this TSecurityAuditType instance.
func (me *TSecurityAuditType) Walk() (err error) {
	if fn := WalkHandlers.TSecurityAuditType; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if err = me.TSecurityAuditType.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
			return
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

// TxsdRedefineIdentificationTypeComplexContentRestrictionNym defines type xsdRedefineIdentificationTypeComplexContentRestrictionNym
type TxsdRedefineIdentificationTypeComplexContentRestrictionNym sac.TnymType

// IsAnonymity : Returns true if the value of this enumerated TxsdRedefineIdentificationTypeComplexContentRestrictionNym is "anonymity".
func (me TxsdRedefineIdentificationTypeComplexContentRestrictionNym) IsAnonymity() bool {
	return me.String() == "anonymity"
}

// IsPseudonymity : Returns true if the value of this enumerated TxsdRedefineIdentificationTypeComplexContentRestrictionNym is "pseudonymity".
func (me TxsdRedefineIdentificationTypeComplexContentRestrictionNym) IsPseudonymity() bool {
	return me.String() == "pseudonymity"
}

// Set : Since TxsdRedefineIdentificationTypeComplexContentRestrictionNym is just a simple String type, this merely sets the current value from the specified string.
func (me *TxsdRedefineIdentificationTypeComplexContentRestrictionNym) Set(s string) {
	(*sac.TnymType)(me).Set(s)
}

// String : Since TxsdRedefineIdentificationTypeComplexContentRestrictionNym is just a simple String type, this merely returns the current string value.
func (me TxsdRedefineIdentificationTypeComplexContentRestrictionNym) String() string {
	return sac.TnymType(me).String()
}

// ToTnymType : This convenience method just performs a simple type conversion to TxsdRedefineIdentificationTypeComplexContentRestrictionNym's alias type sac.TnymType.
func (me TxsdRedefineIdentificationTypeComplexContentRestrictionNym) ToTnymType() sac.TnymType {
	return sac.TnymType(me)
}

// TAuthenticatorBaseType defines type AuthenticatorBaseType
type TAuthenticatorBaseType struct {
	*TAuthenticatorBaseType
	sac.XElemSharedSecretDynamicPlaintext
	sac.XElemAsymmetricDecryption
	sac.XElemAsymmetricKeyAgreement
	sac.XElemDigSig
	sac.XElemZeroKnowledge
	sac.XElemSharedSecretChallengeResponse
	sac.XElemComplexAuthenticator
	sac.XElemsExtension
}

// Walk : if the WalkHandlers.TAuthenticatorBaseType function is not nil (ie. was set by outside code), calls it with this TAuthenticatorBaseType instance as the single argument. Then calls the Walk() method on 1/9 embed(s) and 0/0 field(s) belonging to this TAuthenticatorBaseType instance.
func (me *TAuthenticatorBaseType) Walk() (err error) {
	if fn := WalkHandlers.TAuthenticatorBaseType; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if err = me.TAuthenticatorBaseType.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
			return
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

// TComplexAuthenticatorType defines type ComplexAuthenticatorType
type TComplexAuthenticatorType struct {
	*TComplexAuthenticatorType
	sac.XElemSharedSecretChallengeResponse
	sac.XElemSharedSecretDynamicPlaintext
	sac.XElemPassword
}

// Walk : if the WalkHandlers.TComplexAuthenticatorType function is not nil (ie. was set by outside code), calls it with this TComplexAuthenticatorType instance as the single argument. Then calls the Walk() method on 1/4 embed(s) and 0/0 field(s) belonging to this TComplexAuthenticatorType instance.
func (me *TComplexAuthenticatorType) Walk() (err error) {
	if fn := WalkHandlers.TComplexAuthenticatorType; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if err = me.TComplexAuthenticatorType.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
			return
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

// TAuthenticatorTransportProtocolType defines type AuthenticatorTransportProtocolType
type TAuthenticatorTransportProtocolType struct {
	sac.XElemWtls
	sac.XElemsExtension
	*TAuthenticatorTransportProtocolType
	sac.XElemSsl
	sac.XElemMobileNetworkNoEncryption
	sac.XElemMobileNetworkRadioEncryption
	sac.XElemMobileNetworkEndToEndEncryption
}

// Walk : if the WalkHandlers.TAuthenticatorTransportProtocolType function is not nil (ie. was set by outside code), calls it with this TAuthenticatorTransportProtocolType instance as the single argument. Then calls the Walk() method on 1/7 embed(s) and 0/0 field(s) belonging to this TAuthenticatorTransportProtocolType instance.
func (me *TAuthenticatorTransportProtocolType) Walk() (err error) {
	if fn := WalkHandlers.TAuthenticatorTransportProtocolType; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if err = me.TAuthenticatorTransportProtocolType.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
			return
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

// TPrivateKeyProtectionType defines type PrivateKeyProtectionType
type TPrivateKeyProtectionType struct {
	*TPrivateKeyProtectionType
	sac.XElemKeyActivation
	sac.XElemKeyStorage
	sac.XElemsExtension
}

// Walk : if the WalkHandlers.TPrivateKeyProtectionType function is not nil (ie. was set by outside code), calls it with this TPrivateKeyProtectionType instance as the single argument. Then calls the Walk() method on 1/4 embed(s) and 0/0 field(s) belonging to this TPrivateKeyProtectionType instance.
func (me *TPrivateKeyProtectionType) Walk() (err error) {
	if fn := WalkHandlers.TPrivateKeyProtectionType; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if err = me.TPrivateKeyProtectionType.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
			return
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

// XCdata defines type CDATA
type XCdata struct {
	XCDATA string `xml:",chardata"`
}

// Walk : if the WalkHandlers.XCdata function is not nil (ie. was set by outside code), calls it with this XCdata instance as the single argument. Then calls the Walk() method on 0/0 embed(s) and 0/1 field(s) belonging to this XCdata instance.
func (me *XCdata) Walk() (err error) {
	if fn := WalkHandlers.XCdata; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

// TOperationalProtectionType defines type OperationalProtectionType
type TOperationalProtectionType struct {
	*TOperationalProtectionType
	sac.XElemSecurityAudit
	sac.XElemDeactivationCallCenter
	sac.XElemsExtension
}

// Walk : if the WalkHandlers.TOperationalProtectionType function is not nil (ie. was set by outside code), calls it with this TOperationalProtectionType instance as the single argument. Then calls the Walk() method on 1/4 embed(s) and 0/0 field(s) belonging to this TOperationalProtectionType instance.
func (me *TOperationalProtectionType) Walk() (err error) {
	if fn := WalkHandlers.TOperationalProtectionType; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if err = me.TOperationalProtectionType.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
			return
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

// TSecretKeyProtectionType defines type SecretKeyProtectionType
type TSecretKeyProtectionType struct {
	*TSecretKeyProtectionType
	sac.XElemKeyActivation
	sac.XElemKeyStorage
	sac.XElemsExtension
}

// Walk : if the WalkHandlers.TSecretKeyProtectionType function is not nil (ie. was set by outside code), calls it with this TSecretKeyProtectionType instance as the single argument. Then calls the Walk() method on 1/4 embed(s) and 0/0 field(s) belonging to this TSecretKeyProtectionType instance.
func (me *TSecretKeyProtectionType) Walk() (err error) {
	if fn := WalkHandlers.TSecretKeyProtectionType; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if err = me.TSecretKeyProtectionType.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
			return
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

// XAttrNymTxsdRedefineIdentificationTypeComplexContentRestrictionNym defines attribute NymTxsdRedefineIdentificationTypeComplexContentRestrictionNym
type XAttrNymTxsdRedefineIdentificationTypeComplexContentRestrictionNym struct {
	Nym TxsdRedefineIdentificationTypeComplexContentRestrictionNym `xml:"nym,attr,omitempty"`
}

// TIdentificationType defines type IdentificationType
type TIdentificationType struct {
	*TIdentificationType
	sac.XElemGoverningAgreements
	sac.XElemsExtension
	XAttrNymTxsdRedefineIdentificationTypeComplexContentRestrictionNym
}

// Walk : if the WalkHandlers.TIdentificationType function is not nil (ie. was set by outside code), calls it with this TIdentificationType instance as the single argument. Then calls the Walk() method on 1/4 embed(s) and 0/0 field(s) belonging to this TIdentificationType instance.
func (me *TIdentificationType) Walk() (err error) {
	if fn := WalkHandlers.TIdentificationType; me != nil {
		if fn != nil {
			if err = fn(me, true); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
		if err = me.TIdentificationType.Walk(); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
			return
		}
		if fn != nil {
			if err = fn(me, false); xsdt.OnWalkError(&err, &WalkErrors, WalkContinueOnError, WalkOnError) {
				return
			}
		}
	}
	return
}

var (
	// WalkContinueOnError can be set to false to break a Walk() immediately as soon as the first error is returned by a custom handler function.
	// If true, Walk() proceeds and accumulates all errors in the WalkErrors slice.
	WalkContinueOnError = true
	// WalkErrors contains all errors accumulated during Walk()s. If you're using this, you need to reset this yourself as needed prior to a fresh Walk().
	WalkErrors []error
	// WalkOnError is your custom error-handling function, if required.
	WalkOnError func(error)
	// WalkHandlers Provides 13 strong-typed hooks for your own custom handler functions to be invoked when the Walk() method is called on any instance of any (non-attribute-related) struct type defined in this package.
	// If your custom handler does get called at all for a given struct instance, then it always gets called twice, first with the 'enter' bool argument set to true, then (after having Walk()ed all subordinate struct instances, if any) once again with it set to false.
	WalkHandlers = &XWalkHandlers{}
)

// XWalkHandlers Provides 13 strong-typed hooks for your own custom handler functions to be invoked when the Walk() method is called on any instance of any (non-attribute-related) struct type defined in this package.
// If your custom handler does get called at all for a given struct instance, then it always gets called twice, first with the 'enter' bool argument set to true, then (after having Walk()ed all subordinate struct instances, if any) once again with it set to false.
type XWalkHandlers struct {
	TAuthnContextDeclarationBaseType    func(*TAuthnContextDeclarationBaseType, bool) error
	TKeyStorageType                     func(*TKeyStorageType, bool) error
	TSecurityAuditType                  func(*TSecurityAuditType, bool) error
	TAuthenticatorBaseType              func(*TAuthenticatorBaseType, bool) error
	TComplexAuthenticatorType           func(*TComplexAuthenticatorType, bool) error
	TOperationalProtectionType          func(*TOperationalProtectionType, bool) error
	TAuthnMethodBaseType                func(*TAuthnMethodBaseType, bool) error
	TechnicalProtectionBaseType         func(*TechnicalProtectionBaseType, bool) error
	TAuthenticatorTransportProtocolType func(*TAuthenticatorTransportProtocolType, bool) error
	TPrivateKeyProtectionType           func(*TPrivateKeyProtectionType, bool) error
	XCdata                              func(*XCdata, bool) error
	TSecretKeyProtectionType            func(*TSecretKeyProtectionType, bool) error
	TIdentificationType                 func(*TIdentificationType, bool) error
}
