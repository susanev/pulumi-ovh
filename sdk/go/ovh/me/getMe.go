// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package me

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information about the current OVHcloud account.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/lbrlabs/pulumi-ovh/sdk/go/ovh/Me"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Me.GetMe(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetMe(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetMeResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetMeResult
	err := ctx.Invoke("ovh:Me/getMe:getMe", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getMe.
type GetMeResult struct {
	// Postal address of the account
	Address string `pulumi:"address"`
	// Area of the account
	Area string `pulumi:"area"`
	// City of birth
	BirthCity string `pulumi:"birthCity"`
	// Birth date
	BirthDay string `pulumi:"birthDay"`
	// City of the account
	City string `pulumi:"city"`
	// This is the national identification number of the company that possess this account
	CompanyNationalIdentificationNumber string `pulumi:"companyNationalIdentificationNumber"`
	// Type of corporation
	CorporationType string `pulumi:"corporationType"`
	// Country of the account
	Country    string          `pulumi:"country"`
	Currencies []GetMeCurrency `pulumi:"currencies"`
	// The customer code of this account (a numerical value used for identification when contacting support via phone call)
	CustomerCode string `pulumi:"customerCode"`
	// Email address
	Email string `pulumi:"email"`
	// Fax number
	Fax string `pulumi:"fax"`
	// First name
	Firstname string `pulumi:"firstname"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Italian SDI
	ItalianSdi string `pulumi:"italianSdi"`
	// Preferred language for this account
	Language string `pulumi:"language"`
	// Legal form of the account
	Legalform string `pulumi:"legalform"`
	// Name of the account holder
	Name string `pulumi:"name"`
	// National Identification Number of this account
	NationalIdentificationNumber string `pulumi:"nationalIdentificationNumber"`
	// Nic handle / customer identifier
	Nichandle string `pulumi:"nichandle"`
	// Name of the organisation for this account
	Organisation string `pulumi:"organisation"`
	// OVHcloud subsidiary
	OvhCompany string `pulumi:"ovhCompany"`
	// OVHcloud subsidiary
	OvhSubsidiary string `pulumi:"ovhSubsidiary"`
	// Phone number
	Phone string `pulumi:"phone"`
	// Country code of the phone number
	PhoneCountry string `pulumi:"phoneCountry"`
	// Gender of the account holder
	Sex string `pulumi:"sex"`
	// Backup email address
	SpareEmail string `pulumi:"spareEmail"`
	// State of the postal address
	State string `pulumi:"state"`
	// VAT number
	Vat string `pulumi:"vat"`
	// Zipcode of the address
	Zip string `pulumi:"zip"`
}
