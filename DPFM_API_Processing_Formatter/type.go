package dpfm_api_processing_formatter

type HeaderUpdates struct {
	Site							int		`json:"Site"`
	SiteType						string	`json:"SiteType"`
	SiteOwner						int		`json:"SiteOwner"`
	SiteOwnerBusinessPartnerRole	string	`json:"SiteOwnerBusinessPartnerRole"`
	Brand							*int	`json:"Brand"`
	PersonResponsible				string	`json:"PersonResponsible"`
	ValidityStartDate				string	`json:"ValidityStartDate"`
	ValidityStartTime				string	`json:"ValidityStartTime"`
	ValidityEndDate					string	`json:"ValidityEndDate"`
	ValidityEndTime					string	`json:"ValidityEndTime"`
	DailyOperationStartTime			string	`json:"DailyOperationStartTime"`
	DailyOperationEndTime			string	`json:"DailyOperationEndTime"`
	Description						string	`json:"Description"`
	LongText						string	`json:"LongText"`
	Introduction					*string	`json:"Introduction"`
	OperationRemarks				*string	`json:"OperationRemarks"`
	PhoneNumber						*string	`json:"PhoneNumber"`
	AvailabilityOfParking			*bool	`json:"AvailabilityOfParking"`
	NumberOfParkingSpaces			*int	`json:"NumberOfParkingSpaces"`
	SuperiorSite					*int	`json:"SuperiorSite"`
	Project							*int	`json:"Project"`
	WBSElement						*int	`json:"WBSElement"`
	Tag1							*string	`json:"Tag1"`
	Tag2							*string	`json:"Tag2"`
	Tag3							*string	`json:"Tag3"`
	Tag4							*string	`json:"Tag4"`
	PointConsumptionType      		string  `json:"PointConsumptionType"`
	LastChangeDate					string	`json:"LastChangeDate"`
	LastChangeTime					string	`json:"LastChangeTime"`
	LastChangeUser					int		`json:"LastChangeUser"`
}

type PartnerUpdates struct {
	Site                 	int     `json:"Site"`
	PartnerFunction         string  `json:"PartnerFunction"`
	BusinessPartner         int     `json:"BusinessPartner"`
	BusinessPartnerFullName *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName     *string `json:"BusinessPartnerName"`
	Organization            *string `json:"Organization"`
	Country                 *string `json:"Country"`
	Language                *string `json:"Language"`
	Currency                *string `json:"Currency"`
	ExternalDocumentID      *string `json:"ExternalDocumentID"`
	AddressID               *int    `json:"AddressID"`
	EmailAddress            *string `json:"EmailAddress"`
}

type AddressUpdates struct {
	Site     		int     	`json:"Site"`
	AddressID   	int     	`json:"AddressID"`
	PostalCode  	string 		`json:"PostalCode"`
	LocalSubRegion 	string 		`json:"LocalSubRegion"`
	LocalRegion 	string 		`json:"LocalRegion"`
	Country     	string 		`json:"Country"`
	GlobalRegion   	string 		`json:"GlobalRegion"`
	TimeZone   		string 		`json:"TimeZone"`
	District    	*string 	`json:"District"`
	StreetName  	*string 	`json:"StreetName"`
	CityName    	*string 	`json:"CityName"`
	Building    	*string 	`json:"Building"`
	Floor       	*int		`json:"Floor"`
	Room        	*int		`json:"Room"`
	XCoordinate 	*float32	`json:"XCoordinate"`
	YCoordinate 	*float32	`json:"YCoordinate"`
	ZCoordinate 	*float32	`json:"ZCoordinate"`
}
