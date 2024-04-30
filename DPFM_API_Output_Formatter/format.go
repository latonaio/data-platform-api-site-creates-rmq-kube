package dpfm_api_output_formatter

import (
	dpfm_api_input_reader "data-platform-api-site-creates-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_processing_formatter "data-platform-api-site-creates-rmq-kube/DPFM_API_Processing_Formatter"
	"data-platform-api-site-creates-rmq-kube/sub_func_complementer"
	"encoding/json"

	"golang.org/x/xerrors"
)

func ConvertToHeaderCreates(subfuncSDC *sub_func_complementer.SDC) (*Header, error) {
	data := subfuncSDC.Message.Header

	header, err := TypeConverter[*Header](data)
	if err != nil {
		return nil, err
	}

	return header, nil
}

func ConvertToPartnerCreates(subfuncSDC *sub_func_complementer.SDC) (*[]Partner, error) {
	partners := make([]Partner, 0)

	for _, data := range *subfuncSDC.Message.Partner {
		partner, err := TypeConverter[*Partner](data)
		if err != nil {
			return nil, err
		}

		partners = append(partners, *partner)
	}

	return &partners, nil
}

func ConvertToAddressCreates(subfuncSDC *sub_func_complementer.SDC) (*[]Address, error) {
	addresses := make([]Address, 0)

	for _, data := range *subfuncSDC.Message.Address {
		address, err := TypeConverter[*Address](data)
		if err != nil {
			return nil, err
		}

		addresses = append(addresses, *address)
	}

	return &addresses, nil
}

func ConvertToHeaderUpdates(headerData dpfm_api_input_reader.Header) (*Header, error) {
	data := headerData

	header, err := TypeConverter[*Header](data)
	if err != nil {
		return nil, err
	}

	return header, nil
}

func ConvertToPartnerUpdates(partnerUpdates *[]dpfm_api_processing_formatter.PartnerUpdates) (*[]Partner, error) {
	partners := make([]Partner, 0)

	for _, data := range *partnerUpdates {
		partner, err := TypeConverter[*Partner](data)
		if err != nil {
			return nil, err
		}

		partners = append(partners, *partner)
	}

	return &partners, nil
}

func ConvertToAddressUpdates(addressUpdates *[]dpfm_api_processing_formatter.AddressUpdates) (*[]Address, error) {
	addresses := make([]Address, 0)

	for _, data := range *addressUpdates {
		address, err := TypeConverter[*Address](data)
		if err != nil {
			return nil, err
		}

		addresses = append(addresses, *address)
	}

	return &addresses, nil
}

func ConvertToHeader(
	input *dpfm_api_input_reader.SDC,
	subfuncSDC *sub_func_complementer.SDC,
) *sub_func_complementer.SDC {
	subfuncSDC.Message.Header = &sub_func_complementer.Header{
		Site:                          *input.Header.Site,
		SiteType:                      input.Header.SiteType,
		SiteOwner:                     input.Header.SiteOwner,
		SiteOwnerBusinessPartnerRole:  input.Header.SiteOwnerBusinessPartnerRole,
		Brand:						   input.Header.Brand,
		PersonResponsible:             input.Header.PersonResponsible,
		ValidityStartDate:             input.Header.ValidityStartDate,
		ValidityStartTime:             input.Header.ValidityStartTime,
		ValidityEndDate:               input.Header.ValidityEndDate,
		ValidityEndTime:               input.Header.ValidityEndTime,
		DailyOperationStartTime:	   input.Header.DailyOperationStartTime,
		DailyOperationEndTime:		   input.Header.DailyOperationEndTime,
		Description:                   input.Header.Description,
		LongText:                      input.Header.LongText,
		Introduction:                  input.Header.Introduction,
		OperationRemarks:			   input.Header.OperationRemarks,
		PhoneNumber:			   	   input.Header.PhoneNumber,
		AvailabilityOfParking:		   input.Header.AvailabilityOfParking,
		NumberOfParkingSpaces:		   input.Header.NumberOfParkingSpaces,
		SuperiorSite:				   input.Header.SuperiorSite,
		Project:                       input.Header.Project,
		WBSElement:                    input.Header.WBSElement,
		Tag1:                          input.Header.Tag1,
		Tag2:                          input.Header.Tag2,
		Tag3:                          input.Header.Tag3,
		Tag4:                          input.Header.Tag4,
		PointConsumptionType:          input.Header.PointConsumptionType,
		CreationDate:                  input.Header.CreationDate,
		CreationTime:                  input.Header.CreationTime,
		LastChangeDate:                input.Header.LastChangeDate,
		LastChangeTime:                input.Header.LastChangeTime,
		IsReleased:                    input.Header.IsReleased,
		IsMarkedForDeletion:           input.Header.IsMarkedForDeletion,
	}

	return subfuncSDC
}

func ConvertToAddress(
	input *dpfm_api_input_reader.SDC,
	subfuncSDC *sub_func_complementer.SDC,
) *sub_func_complementer.SDC {
	var addresses []sub_func_complementer.Address

	addresses = append(
		addresses,
		sub_func_complementer.Address{
			Site:          *input.Header.Site,
			AddressID:      input.Header.Address[0].AddressID,
			PostalCode:     input.Header.Address[0].PostalCode,
			LocalSubRegion: input.Header.Address[0].LocalSubRegion,
			LocalRegion:    input.Header.Address[0].LocalRegion,
			Country:        input.Header.Address[0].Country,
			GlobalRegion:   input.Header.Address[0].GlobalRegion,
			TimeZone:       input.Header.Address[0].TimeZone,
			District:       input.Header.Address[0].District,
			StreetName:     input.Header.Address[0].StreetName,
			CityName:       input.Header.Address[0].CityName,
			Building:       input.Header.Address[0].Building,
			Floor:          input.Header.Address[0].Floor,
			Room:           input.Header.Address[0].Room,
			XCoordinate:    input.Header.Address[0].XCoordinate,
			YCoordinate:    input.Header.Address[0].YCoordinate,
			ZCoordinate:    input.Header.Address[0].ZCoordinate,
		},
	)

	subfuncSDC.Message.Address = &addresses

	return subfuncSDC
}

func ConvertToPartner(
	input *dpfm_api_input_reader.SDC,
	subfuncSDC *sub_func_complementer.SDC,
) *sub_func_complementer.SDC {
	var partners []sub_func_complementer.Partner

	partners = append(
		partners,
		sub_func_complementer.Partner{
			Site:          			 *input.Header.Site,
			PartnerFunction:         input.Header.Partner[0].PartnerFunction,
			BusinessPartner:         input.Header.Partner[0].BusinessPartner,
			BusinessPartnerFullName: input.Header.Partner[0].BusinessPartnerFullName,
			BusinessPartnerName:     input.Header.Partner[0].BusinessPartnerName,
			Organization:            input.Header.Partner[0].Organization,
			Country:                 input.Header.Partner[0].Country,
			Language:                input.Header.Partner[0].Language,
			Currency:                input.Header.Partner[0].Currency,
			ExternalDocumentID:      input.Header.Partner[0].ExternalDocumentID,
			AddressID:               input.Header.Partner[0].AddressID,
			EmailAddress:            input.Header.Partner[0].EmailAddress,
		},
	)

	subfuncSDC.Message.Partner = &partners

	return subfuncSDC
}

func TypeConverter[T any](data interface{}) (T, error) {
	var dist T
	b, err := json.Marshal(data)
	if err != nil {
		return dist, xerrors.Errorf("Marshal error: %w", err)
	}
	err = json.Unmarshal(b, &dist)
	if err != nil {
		return dist, xerrors.Errorf("Unmarshal error: %w", err)
	}
	return dist, nil
}
