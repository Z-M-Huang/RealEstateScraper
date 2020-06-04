package data

import "time"

//TruliaWebData trulia web data
type TruliaWebData struct {
	Props struct {
		AsPath      string             `json:"asPath"`
		HomeDetails *truliaHomeDetails `json:"homeDetails"`
	} `json:"props"`
}

type truliaHomeDetails struct {
	URL           string               `json:"url"`
	Media         *truliaMedia         `json:"media"`
	PageText      *truliaPageText      `json:"pageText"`
	CurrentStatus *truliaCurrentStatus `json:"currentStatus"`
	Location      *truliaLocation      `json:"location"`
	Price         *truliaPrice         `json:"price"`
	Bedrooms      *truliaBedrooms      `json:"bedrooms"`
	Bathrooms     *truliaBathrooms     `json:"bathrooms"`
	FloorSpace    *truliaFloorSpace    `json:"floorSpace"`
	HoaFee        string               `json:"hoaFee"`
	Features      struct {
		Attributes []struct {
		} `json:"attributes"`
	} `json:"features"`
	PublicRecord struct {
		Features *truliaFeatures `json:"features"`
	} `json:"publicRecord"`
	Description struct {
		Value string `json:"value"`
	} `json:"description"`
	PriceHistory         []*truliaPriceHistory   `json:"priceHistory"`
	LocalInfoSummary     *truliaLocalInfoSummary `json:"localInfoSummary"`
	ActiveForSaleListing struct {
		DateListed time.Time `json:"dateListed"`
	} `json:"activeForSaleListing"`
	AssignedSchools struct {
		Schools []*truliaSchools `json:"schools"`
	} `json:"assignedSchools"`
	PropertyType struct {
		BranchBannerPropertyType string `json:"branchBannerPropertyType"`
	} `json:"propertyType"`
}

type truliaPhotos struct {
	URL struct {
		HiDpiLargeSrc string `json:"hiDpiLargeSrc"`
	} `json:"url"`
}

type truliaMedia struct {
	Photos          []*truliaPhotos `json:"photos"`
	TotalPhotoCount int             `json:"totalPhotoCount"`
}

type truliaPageText struct {
	Title           string `json:"title"`
	MetaDescription string `json:"metaDescription"`
}

type truliaCurrentStatus struct {
	IsOffMarket      bool   `json:"isOffMarket"`
	IsRecentlySold   bool   `json:"isRecentlySold"`
	IsForeclosure    bool   `json:"isForeclosure"`
	IsActiveForRent  bool   `json:"isActiveForRent"`
	IsActiveForSale  bool   `json:"isActiveForSale"`
	IsRecentlyRented bool   `json:"isRecentlyRented"`
	Label            string `json:"label"`
}

type truliaCoordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type truliaLocation struct {
	StateCode                string             `json:"stateCode"`
	HomeFormattedAddress     string             `json:"homeFormattedAddress"`
	CityStateZipAddress      string             `json:"cityStateZipAddress"`
	SummaryFormattedLocation string             `json:"summaryFormattedLocation"`
	City                     string             `json:"city"`
	ZipCode                  string             `json:"zipCode"`
	NeighborhoodName         string             `json:"neighborhoodName"`
	Coordinates              *truliaCoordinates `json:"coordinates"`
	StreetAddress            string             `json:"streetAddress"`
	FormattedLocation        string             `json:"formattedLocation"`
}

type truliaPrice struct {
	Price             int    `json:"price"`
	CurrencyCode      string `json:"currencyCode"`
	FormattedPrice    string `json:"formattedPrice"`
	BranchBannerPrice string `json:"branchBannerPrice"`
}

type truliaBedrooms struct {
	SummaryBedrooms      string `json:"summaryBedrooms"`
	FormattedValue       string `json:"formattedValue"`
	BranchBannerBedrooms string `json:"branchBannerBedrooms"`
}

type truliaBathrooms struct {
	SummaryBathrooms      string `json:"summaryBathrooms"`
	FormattedValue        string `json:"formattedValue"`
	BranchBannerBathrooms string `json:"branchBannerBathrooms"`
}

type truliaFloorSpace struct {
	FormattedDimension     string `json:"formattedDimension"`
	BranchBannerFloorSpace string `json:"branchBannerFloorSpace"`
}

type truliaAttributes struct {
	Key            string `json:"key,omitempty"`
	FormattedValue string `json:"formattedValue"`
	FormattedName  string `json:"formattedName,omitempty"`
}

type truliaFeatures struct {
	Attributes []*truliaAttributes `json:"attributes"`
}

type truliaPriceHistory struct {
	FormattedDate string `json:"formattedDate"`
	Event         string `json:"event"`
	Price         struct {
		FormattedPrice string `json:"formattedPrice"`
	} `json:"price"`
	Attributes []*truliaAttributes `json:"attributes"`
}

type truliaLocalInfoSummary struct {
	Crime struct {
		CrimeSummaryDescription string `json:"crimeSummaryDescription"`
	} `json:"crime"`
	Schools struct {
		SchoolsSummaryDescriptionLines []string `json:"schoolsSummaryDescriptionLines"`
	} `json:"schools"`
	Amenities struct {
		AmenitiesSummaryDescriptionLines []string `json:"amenitiesSummaryDescriptionLines"`
	} `json:"amenities"`
}

type truliaSchools struct {
	ID             string   `json:"id"`
	Name           string   `json:"name"`
	DistrictName   string   `json:"districtName"`
	Categories     []string `json:"categories"`
	EnrollmentType string   `json:"enrollmentType"`
	Reviews        []struct {
		Rating       int       `json:"rating"`
		MaxRating    int       `json:"maxRating"`
		ReviewText   string    `json:"reviewText"`
		Type         string    `json:"type"`
		RelativeDate string    `json:"relativeDate"`
		Date         time.Time `json:"date"`
	} `json:"reviews"`
	ReviewCount    int `json:"reviewCount"`
	ProviderRating struct {
		Rating    int `json:"rating"`
		MaxRating int `json:"maxRating"`
	} `json:"providerRating"`
	StudentCount        int     `json:"studentCount"`
	GradesRange         string  `json:"gradesRange"`
	StreetAddress       string  `json:"streetAddress"`
	CityName            string  `json:"cityName"`
	StateCode           string  `json:"stateCode"`
	ZipCode             string  `json:"zipCode"`
	ProviderURL         string  `json:"providerUrl"`
	URL                 string  `json:"url"`
	Latitude            float64 `json:"latitude"`
	Longitude           float64 `json:"longitude"`
	AverageParentRating struct {
		Rating    int `json:"rating"`
		MaxRating int `json:"maxRating"`
	} `json:"averageParentRating"`
}
