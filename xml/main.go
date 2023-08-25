package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/antchfx/xmlquery"
)

const (
	UpdateOperation  = "updated"
	DestroyOperation = "destroyed"
	CreateOperation  = "created"
)

func OperationType(value string) string {
	switch value {
	case "1":
		return UpdateOperation
	case "2":
		return DestroyOperation
	case "3":
		return CreateOperation
	default:
		return "unknown"
	}
}

type CertificateChange struct {
	Certificate Certificate
	Period      CertificateDescriptionPeriod
	Description CertificateDescription
}

type Certificate struct {
	UpdateType        string
	TypeCode          string
	Code              string
	ValidityStartDate string
	ValidityEndDate   string
}

type CertificateDescriptionPeriod struct {
	UpdateType        string
	SID               string
	TypeCode          string
	Code              string
	ValidityStartDate string
	ValidityEndDate   string
}

type CertificateDescription struct {
	UpdateType  string
	SID         string
	LangID      string
	TypeCode    string
	Code        string
	Description string
}

func getElementText(node *xmlquery.Node, elementName string) string {
	el := node.SelectElement(elementName)
	if el != nil {
		return el.InnerText()
	}
	return ""
}

func populateCertificate(node *xmlquery.Node) Certificate {
	return Certificate{
		UpdateType:        OperationType(getElementText(node.Parent, "oub:update.type")),
		TypeCode:          getElementText(node, "oub:certificate.type.code"),
		Code:              getElementText(node, "oub:certificate.code"),
		ValidityStartDate: getElementText(node, "oub:validity.start.date"),
		ValidityEndDate:   getElementText(node, "oub:validity.end.date"),
	}
}

func populateCertificateDescriptionPeriod(node *xmlquery.Node) CertificateDescriptionPeriod {
	return CertificateDescriptionPeriod{
		UpdateType:        OperationType(getElementText(node.Parent, "oub:update.type")),
		SID:               getElementText(node, "oub:certificate.description.period.sid"),
		TypeCode:          getElementText(node, "oub:certificate.type.code"),
		Code:              getElementText(node, "oub:certificate.code"),
		ValidityStartDate: getElementText(node, "oub:validity.start.date"),
		ValidityEndDate:   getElementText(node, "oub:validity.end.date"),
	}
}

func populateCertificateDescription(node *xmlquery.Node) CertificateDescription {
	return CertificateDescription{
		UpdateType:  OperationType(getElementText(node.Parent, "oub:update.type")),
		SID:         getElementText(node, "oub:certificate.description.period.sid"),
		LangID:      getElementText(node, "oub:language.id"),
		TypeCode:    getElementText(node, "oub:certificate.type.code"),
		Code:        getElementText(node, "oub:certificate.code"),
		Description: getElementText(node, "oub:description"),
	}
}

func main() {
	var certificateChanges []CertificateChange
	xml, err := os.ReadFile("test.xml")

	if err != nil {
		panic(err)
	}

	doc, err := xmlquery.Parse(strings.NewReader(string(xml)))

	if err != nil {
		panic(err)
	}

	transactions := xmlquery.Find(doc, "//env:envelope/env:transaction")

	for _, t := range transactions {
		certificateNode := xmlquery.FindOne(t, ".//oub:certificate")
		descriptionPeriodNode := xmlquery.FindOne(t, ".//oub:certificate.description.period")
		descriptionNode := xmlquery.FindOne(t, ".//oub:certificate.description")

		if certificateNode != nil && descriptionPeriodNode != nil && descriptionNode != nil {
			certificateChanges = append(
				certificateChanges,
				CertificateChange{
					Certificate: populateCertificate(certificateNode),
					Period:      populateCertificateDescriptionPeriod(descriptionPeriodNode),
					Description: populateCertificateDescription(descriptionNode),
				},
			)
		}

	}

	fmt.Printf("%+v\n", certificateChanges)
}
