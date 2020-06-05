package scraper

import "time"

type truliaRawObject struct {
	Props *truliaObject `json:"props"`
}

type truliaObject struct {
	Page struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Tracking    struct {
			PageLoad              string `json:"pageLoad"`
			ListingState          string `json:"listingState"`
			ListingCounty         string `json:"listingCounty"`
			ListingCity           string `json:"listingCity"`
			ListingNeighborhood   string `json:"listingNeighborhood"`
			ListingZip            string `json:"listingZip"`
			ListingPrice          string `json:"listingPrice"`
			ListingStatus         string `json:"listingStatus"`
			ListingType           string `json:"listingType"`
			PropertyType          string `json:"propertyType"`
			LocationLat           string `json:"locationLat"`
			LocationLon           string `json:"locationLon"`
			IsthreeDTourAvaliable string `json:"isthreeDTourAvaliable"`
			SiteSection           string `json:"siteSection"`
			PageType              string `json:"pageType"`
			PageDetails           string `json:"pageDetails"`
		} `json:"tracking"`
	} `json:"_page"`
	AsPath      string `json:"asPath"`
	HomeDetails struct {
		URL   string `json:"url"`
		Media struct {
			Photos []struct {
				URL struct {
					HiDpiExtraSmallSrc string `json:"hiDpiExtraSmallSrc"`
					Thumbnail          string `json:"thumbnail"`
					ExtraSmallSrc      string `json:"extraSmallSrc"`
					SmallSrc           string `json:"smallSrc"`
					MediumSrc          string `json:"mediumSrc"`
					LargeSrc           string `json:"largeSrc"`
					HiDpiSmallSrc      string `json:"hiDpiSmallSrc"`
					HiDpiMediumSrc     string `json:"hiDpiMediumSrc"`
					HiDpiLargeSrc      string `json:"hiDpiLargeSrc"`
				} `json:"url"`
				WebpURL struct {
					HiDpiExtraSmallWebpSrc string `json:"hiDpiExtraSmallWebpSrc"`
					ExtraSmallWebpSrc      string `json:"extraSmallWebpSrc"`
					SmallWebpSrc           string `json:"smallWebpSrc"`
					MediumWebpSrc          string `json:"mediumWebpSrc"`
					LargeWebpSrc           string `json:"largeWebpSrc"`
					HiDpiSmallWebpSrc      string `json:"hiDpiSmallWebpSrc"`
					HiDpiMediumWebpSrc     string `json:"hiDpiMediumWebpSrc"`
					HiDpiLargeWebpSrc      string `json:"hiDpiLargeWebpSrc"`
				} `json:"webpUrl"`
			} `json:"photos"`
			TotalPhotoCount int `json:"totalPhotoCount"`
			Videos          []struct {
				URL struct {
					Original string `json:"original"`
					Poster   struct {
						Large string `json:"large"`
					} `json:"poster"`
				} `json:"url"`
			} `json:"videos"`
			HasVideo bool `json:"hasVideo"`
		} `json:"media"`
		PageText struct {
			Title           string `json:"title"`
			MetaDescription string `json:"metaDescription"`
		} `json:"pageText"`
		CurrentStatus struct {
			IsOffMarket      bool   `json:"isOffMarket"`
			IsRecentlySold   bool   `json:"isRecentlySold"`
			IsForeclosure    bool   `json:"isForeclosure"`
			IsActiveForRent  bool   `json:"isActiveForRent"`
			IsActiveForSale  bool   `json:"isActiveForSale"`
			IsRecentlyRented bool   `json:"isRecentlyRented"`
			Label            string `json:"label"`
		} `json:"currentStatus"`
		Location struct {
			StateCode                string `json:"stateCode"`
			HomeFormattedAddress     string `json:"homeFormattedAddress"`
			CityStateZipAddress      string `json:"cityStateZipAddress"`
			SummaryFormattedLocation string `json:"summaryFormattedLocation"`
			City                     string `json:"city"`
			ZipCode                  string `json:"zipCode"`
			NeighborhoodName         string `json:"neighborhoodName"`
			Coordinates              struct {
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
			} `json:"coordinates"`
			StreetAddress     string `json:"streetAddress"`
			FormattedLocation string `json:"formattedLocation"`
		} `json:"location"`
		Price struct {
			Price          int    `json:"price"`
			CurrencyCode   string `json:"currencyCode"`
			FormattedPrice string `json:"formattedPrice"`
		} `json:"price"`
		Bedrooms struct {
			SummaryBedrooms      string `json:"summaryBedrooms"`
			FormattedValue       string `json:"formattedValue"`
			BranchBannerBedrooms string `json:"branchBannerBedrooms"`
		} `json:"bedrooms"`
		Bathrooms struct {
			SummaryBathrooms      string `json:"summaryBathrooms"`
			FormattedValue        string `json:"formattedValue"`
			BranchBannerBathrooms string `json:"branchBannerBathrooms"`
		} `json:"bathrooms"`
		FloorSpace struct {
			FormattedDimension     string `json:"formattedDimension"`
			BranchBannerFloorSpace string `json:"branchBannerFloorSpace"`
		} `json:"floorSpace"`
		Surroundings struct {
			NdpActive    bool   `json:"ndpActive"`
			NdpURL       string `json:"ndpUrl"`
			LocalInfoURL string `json:"localInfoUrl"`
			Name         string `json:"name"`
			Media        struct {
				HeroImage struct {
					URL struct {
						Path string `json:"path"`
					} `json:"url"`
				} `json:"heroImage"`
			} `json:"media"`
			LocalFacts struct {
				ForSaleStats struct {
					Min int `json:"min"`
					Max int `json:"max"`
				} `json:"forSaleStats"`
				HomesForSaleCount int `json:"homesForSaleCount"`
				ForRentStats      struct {
					Min int `json:"min"`
					Max int `json:"max"`
				} `json:"forRentStats"`
				HomesForRentCount int `json:"homesForRentCount"`
			} `json:"localFacts"`
		} `json:"surroundings"`
		HoaFee       string `json:"hoaFee"`
		MortgageInfo struct {
			CtaUrls struct {
				PurchaseRates string `json:"purchaseRates"`
				PostLead      string `json:"postLead"`
			} `json:"ctaUrls"`
			MortgageInquiryCta struct {
				Text string `json:"text"`
				URL  string `json:"url"`
			} `json:"mortgageInquiryCta"`
			Rates struct {
				DefaultRates struct {
					TaxRate      float64 `json:"taxRate"`
					InterestRate float64 `json:"interestRate"`
				} `json:"defaultRates"`
				TaxRate       float64 `json:"taxRate"`
				InterestRates []struct {
					LoanDuration string  `json:"loanDuration"`
					Rate         float64 `json:"rate"`
					CreditScore  string  `json:"creditScore"`
					DisplayName  string  `json:"displayName"`
					LoanType     string  `json:"loanType"`
				} `json:"interestRates"`
			} `json:"rates"`
			Defaults struct {
				HomePrice struct {
					Price int `json:"price"`
				} `json:"homePrice"`
				LoanDuration          string `json:"loanDuration"`
				DownPaymentPercentage int    `json:"downPaymentPercentage"`
				Insurance             struct {
					Price int `json:"price"`
				} `json:"insurance"`
			} `json:"defaults"`
		} `json:"mortgageInfo"`
		Features struct {
			Attributes []struct {
				FormattedValue string `json:"formattedValue"`
				FormattedName  string `json:"formattedName,omitempty"`
			} `json:"attributes"`
		} `json:"features"`
		PublicRecord struct {
			Features struct {
				Attributes []struct {
					FormattedValue string `json:"formattedValue"`
					FormattedName  string `json:"formattedName,omitempty"`
				} `json:"attributes"`
			} `json:"features"`
		} `json:"publicRecord"`
		Description struct {
			Value string `json:"value"`
		} `json:"description"`
		TitleToPriceHistory string `json:"titleToPriceHistory"`
		PriceHistory        []struct {
			FormattedDate string `json:"formattedDate"`
			Event         string `json:"event"`
			Price         struct {
				FormattedPrice string `json:"formattedPrice"`
			} `json:"price"`
			Attributes []struct {
				Key                string `json:"key"`
				FormattedAttribute string `json:"formattedAttribute"`
			} `json:"attributes"`
		} `json:"priceHistory"`
		LocalProtections struct {
			Lede        string `json:"lede"`
			Protections []struct {
				Name        string `json:"name"`
				Description string `json:"description"`
				Covers      []struct {
					Key   string `json:"key"`
					Value string `json:"value"`
				} `json:"covers"`
			} `json:"protections"`
		} `json:"localProtections"`
		Taxes struct {
			HighlightedAssessments struct {
				Year     int `json:"year"`
				TaxValue struct {
					FormattedPrice string `json:"formattedPrice"`
				} `json:"taxValue"`
				TotalAssessment struct {
					FormattedPrice string `json:"formattedPrice"`
				} `json:"totalAssessment"`
				Assessments []struct {
					Type   string `json:"type"`
					Amount struct {
						FormattedPrice string `json:"formattedPrice"`
					} `json:"amount"`
				} `json:"assessments"`
			} `json:"highlightedAssessments"`
		} `json:"taxes"`
		LocalInfoSummary struct {
			Crime struct {
				CrimeSummaryDescription string `json:"crimeSummaryDescription"`
			} `json:"crime"`
			Schools struct {
				SchoolsSummaryDescriptionLines []string `json:"schoolsSummaryDescriptionLines"`
			} `json:"schools"`
			Commute struct {
				CommuteSummaryDescription string `json:"commuteSummaryDescription"`
			} `json:"commute"`
			Amenities struct {
				AmenitiesSummaryDescriptionLines []string `json:"amenitiesSummaryDescriptionLines"`
			} `json:"amenities"`
		} `json:"localInfoSummary"`
		ActiveForSaleListing struct {
			DateListed time.Time `json:"dateListed"`
		} `json:"activeForSaleListing"`
		AssignedSchools struct {
			Schools []struct {
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
			} `json:"schools"`
		} `json:"assignedSchools"`
	} `json:"homeDetails"`
}
