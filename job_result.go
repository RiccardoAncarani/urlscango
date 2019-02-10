package urlscango

import "time"

type JobResult struct {
	Data struct {
		Requests []struct {
			Request struct {
				RequestID   string `json:"requestId"`
				LoaderID    string `json:"loaderId"`
				DocumentURL string `json:"documentURL"`
				Request     struct {
					URL     string `json:"url"`
					Method  string `json:"method"`
					Headers struct {
						UpgradeInsecureRequests string `json:"Upgrade-Insecure-Requests"`
						UserAgent               string `json:"User-Agent"`
					} `json:"headers"`
					MixedContentType string `json:"mixedContentType"`
					InitialPriority  string `json:"initialPriority"`
					ReferrerPolicy   string `json:"referrerPolicy"`
				} `json:"request"`
				Timestamp float64 `json:"timestamp"`
				WallTime  float64 `json:"wallTime"`
				Initiator struct {
					Type string `json:"type"`
				} `json:"initiator"`
				Type           string `json:"type"`
				FrameID        string `json:"frameId"`
				HasUserGesture bool   `json:"hasUserGesture"`
			} `json:"request"`
			Response struct {
				EncodedDataLength int    `json:"encodedDataLength"`
				DataLength        int    `json:"dataLength"`
				RequestID         string `json:"requestId"`
				Type              string `json:"type"`
				Response          struct {
					URL        string `json:"url"`
					Status     int    `json:"status"`
					StatusText string `json:"statusText"`
					Headers    struct {
						Status                  string `json:"status"`
						Date                    string `json:"date"`
						Expires                 string `json:"expires"`
						CacheControl            string `json:"cache-control"`
						ContentType             string `json:"content-type"`
						StrictTransportSecurity string `json:"strict-transport-security"`
						P3P                     string `json:"p3p"`
						ContentEncoding         string `json:"content-encoding"`
						Server                  string `json:"server"`
						ContentLength           string `json:"content-length"`
						XXSSProtection          string `json:"x-xss-protection"`
						XFrameOptions           string `json:"x-frame-options"`
						SetCookie               string `json:"set-cookie"`
						AltSvc                  string `json:"alt-svc"`
					} `json:"headers"`
					MimeType       string `json:"mimeType"`
					RequestHeaders struct {
						Method                  string `json:":method"`
						Authority               string `json:":authority"`
						Scheme                  string `json:":scheme"`
						Path                    string `json:":path"`
						Pragma                  string `json:"pragma"`
						CacheControl            string `json:"cache-control"`
						UpgradeInsecureRequests string `json:"upgrade-insecure-requests"`
						UserAgent               string `json:"user-agent"`
						Accept                  string `json:"accept"`
						AcceptEncoding          string `json:"accept-encoding"`
					} `json:"requestHeaders"`
					RemoteIPAddress   string `json:"remoteIPAddress"`
					RemotePort        int    `json:"remotePort"`
					EncodedDataLength int    `json:"encodedDataLength"`
					Timing            struct {
						RequestTime       float64 `json:"requestTime"`
						ProxyStart        int     `json:"proxyStart"`
						ProxyEnd          int     `json:"proxyEnd"`
						DNSStart          int     `json:"dnsStart"`
						DNSEnd            int     `json:"dnsEnd"`
						ConnectStart      int     `json:"connectStart"`
						ConnectEnd        int     `json:"connectEnd"`
						SslStart          int     `json:"sslStart"`
						SslEnd            int     `json:"sslEnd"`
						WorkerStart       int     `json:"workerStart"`
						WorkerReady       int     `json:"workerReady"`
						SendStart         float64 `json:"sendStart"`
						SendEnd           float64 `json:"sendEnd"`
						PushStart         int     `json:"pushStart"`
						PushEnd           int     `json:"pushEnd"`
						ReceiveHeadersEnd float64 `json:"receiveHeadersEnd"`
					} `json:"timing"`
					Protocol        string `json:"protocol"`
					SecurityState   string `json:"securityState"`
					SecurityDetails struct {
						Protocol                          string        `json:"protocol"`
						KeyExchange                       string        `json:"keyExchange"`
						KeyExchangeGroup                  string        `json:"keyExchangeGroup"`
						Cipher                            string        `json:"cipher"`
						CertificateID                     int           `json:"certificateId"`
						SubjectName                       string        `json:"subjectName"`
						SanList                           []string      `json:"sanList"`
						Issuer                            string        `json:"issuer"`
						ValidFrom                         int           `json:"validFrom"`
						ValidTo                           int           `json:"validTo"`
						SignedCertificateTimestampList    []interface{} `json:"signedCertificateTimestampList"`
						CertificateTransparencyCompliance string        `json:"certificateTransparencyCompliance"`
					} `json:"securityDetails"`
					SecurityHeaders []struct {
						Name  string `json:"name"`
						Value string `json:"value"`
					} `json:"securityHeaders"`
				} `json:"response"`
				Hash string `json:"hash"`
				Size int    `json:"size"`
				Asn  struct {
					IP          string `json:"ip"`
					Asn         string `json:"asn"`
					Country     string `json:"country"`
					Registrar   string `json:"registrar"`
					Date        string `json:"date"`
					Description string `json:"description"`
					Route       string `json:"route"`
					Name        string `json:"name"`
				} `json:"asn"`
				Geoip struct {
					Range       [][]int `json:"range"`
					Country     string  `json:"country"`
					Region      string  `json:"region"`
					City        string  `json:"city"`
					Ll          []int   `json:"ll"`
					CountryName string  `json:"country_name"`
				} `json:"geoip"`
				Hashmatches []interface{} `json:"hashmatches"`
			} `json:"response"`
			InitiatorInfo struct {
				URL  string `json:"url"`
				Host string `json:"host"`
				Type string `json:"type"`
			} `json:"initiatorInfo,omitempty"`
		} `json:"requests"`
		Cookies []struct {
			Name     string  `json:"name"`
			Value    string  `json:"value"`
			Domain   string  `json:"domain"`
			Path     string  `json:"path"`
			Expires  float64 `json:"expires"`
			Size     int     `json:"size"`
			HTTPOnly bool    `json:"httpOnly"`
			Secure   bool    `json:"secure"`
			Session  bool    `json:"session"`
		} `json:"cookies"`
		Console []interface{} `json:"console"`
		Links   []struct {
			Href string `json:"href"`
			Text string `json:"text"`
		} `json:"links"`
		Timing struct {
			BeginNavigation      time.Time `json:"beginNavigation"`
			FrameStartedLoading  time.Time `json:"frameStartedLoading"`
			FrameNavigated       time.Time `json:"frameNavigated"`
			DomContentEventFired time.Time `json:"domContentEventFired"`
			LoadEventFired       time.Time `json:"loadEventFired"`
			FrameStoppedLoading  time.Time `json:"frameStoppedLoading"`
		} `json:"timing"`
		Globals []struct {
			Prop string `json:"prop"`
			Type string `json:"type"`
		} `json:"globals"`
	} `json:"data"`
	Stats struct {
		ResourceStats []struct {
			Count       int         `json:"count"`
			Size        int         `json:"size"`
			EncodedSize int         `json:"encodedSize"`
			Latency     int         `json:"latency"`
			Countries   []string    `json:"countries"`
			Ips         []string    `json:"ips"`
			Type        string      `json:"type"`
			Compression string      `json:"compression"`
			Percentage  interface{} `json:"percentage"`
		} `json:"resourceStats"`
		ProtocolStats []struct {
			Count         int      `json:"count"`
			Size          int      `json:"size"`
			EncodedSize   int      `json:"encodedSize"`
			Ips           []string `json:"ips"`
			Countries     []string `json:"countries"`
			SecurityState struct {
			} `json:"securityState"`
			Protocol string `json:"protocol"`
		} `json:"protocolStats"`
		TLSStats []struct {
			Count       int      `json:"count"`
			Size        int      `json:"size"`
			EncodedSize int      `json:"encodedSize"`
			Ips         []string `json:"ips"`
			Countries   []string `json:"countries"`
			Protocols   struct {
				TLS12ECDHEECDSAAES128GCM int `json:"TLS 1.2 / ECDHE_ECDSA / AES_128_GCM"`
				TLS12ECDHERSAAES128GCM   int `json:"TLS 1.2 / ECDHE_RSA / AES_128_GCM"`
			} `json:"protocols"`
			SecurityState string `json:"securityState"`
		} `json:"tlsStats"`
		ServerStats []struct {
			Count       int      `json:"count"`
			Size        int      `json:"size"`
			EncodedSize int      `json:"encodedSize"`
			Ips         []string `json:"ips"`
			Countries   []string `json:"countries"`
			Server      string   `json:"server"`
		} `json:"serverStats"`
		DomainStats []struct {
			Count       int      `json:"count"`
			Ips         []string `json:"ips"`
			Domain      string   `json:"domain"`
			Size        int      `json:"size"`
			EncodedSize int      `json:"encodedSize"`
			Countries   []string `json:"countries"`
			Index       int      `json:"index"`
			Initiators  []string `json:"initiators"`
			Redirects   int      `json:"redirects"`
		} `json:"domainStats"`
		RegDomainStats []struct {
			Count       int           `json:"count"`
			Ips         []string      `json:"ips"`
			RegDomain   string        `json:"regDomain"`
			Size        int           `json:"size"`
			EncodedSize int           `json:"encodedSize"`
			Countries   []interface{} `json:"countries"`
			Index       int           `json:"index"`
			SubDomains  []interface{} `json:"subDomains"`
			Redirects   int           `json:"redirects"`
		} `json:"regDomainStats"`
		SecureRequests   int `json:"secureRequests"`
		SecurePercentage int `json:"securePercentage"`
		IPv6Percentage   int `json:"IPv6Percentage"`
		UniqCountries    int `json:"uniqCountries"`
		TotalLinks       int `json:"totalLinks"`
		Malicious        int `json:"malicious"`
		AdBlocked        int `json:"adBlocked"`
		IPStats          []struct {
			Requests int      `json:"requests"`
			Domains  []string `json:"domains"`
			IP       string   `json:"ip"`
			Asn      struct {
				IP          string `json:"ip"`
				Asn         string `json:"asn"`
				Country     string `json:"country"`
				Registrar   string `json:"registrar"`
				Date        string `json:"date"`
				Description string `json:"description"`
				Route       string `json:"route"`
				Name        string `json:"name"`
			} `json:"asn"`
			DNS struct {
			} `json:"dns"`
			Geoip struct {
				Range       [][]int `json:"range"`
				Country     string  `json:"country"`
				Region      string  `json:"region"`
				City        string  `json:"city"`
				Ll          []int   `json:"ll"`
				CountryName string  `json:"country_name"`
			} `json:"geoip"`
			Size        int         `json:"size"`
			EncodedSize int         `json:"encodedSize"`
			Countries   []string    `json:"countries"`
			Index       int         `json:"index"`
			Ipv6        bool        `json:"ipv6"`
			Redirects   int         `json:"redirects"`
			Count       interface{} `json:"count"`
		} `json:"ipStats"`
	} `json:"stats"`
	Meta struct {
		Processors struct {
			Download struct {
				State string        `json:"state"`
				Data  []interface{} `json:"data"`
			} `json:"download"`
			Rdns struct {
				State string        `json:"state"`
				Data  []interface{} `json:"data"`
			} `json:"rdns"`
			Geoip struct {
				State string `json:"state"`
				Data  []struct {
					IP    string `json:"ip"`
					Geoip struct {
						Range       [][]int `json:"range"`
						Country     string  `json:"country"`
						Region      string  `json:"region"`
						City        string  `json:"city"`
						Ll          []int   `json:"ll"`
						CountryName string  `json:"country_name"`
					} `json:"geoip"`
				} `json:"data"`
			} `json:"geoip"`
			Cdnjs struct {
				State string        `json:"state"`
				Data  []interface{} `json:"data"`
			} `json:"cdnjs"`
			Gsb struct {
				State string `json:"state"`
				Data  struct {
				} `json:"data"`
			} `json:"gsb"`
			Abp struct {
				State string `json:"state"`
				Data  []struct {
					URL    string `json:"url"`
					Source string `json:"source"`
					Type   string `json:"type"`
				} `json:"data"`
			} `json:"abp"`
			Asn struct {
				State string `json:"state"`
				Data  []struct {
					IP          string `json:"ip"`
					Asn         string `json:"asn"`
					Country     string `json:"country"`
					Registrar   string `json:"registrar"`
					Date        string `json:"date"`
					Description string `json:"description"`
					Route       string `json:"route"`
					Name        string `json:"name"`
				} `json:"data"`
			} `json:"asn"`
			Wappa struct {
				State string `json:"state"`
				Data  []struct {
					App        string `json:"app"`
					Confidence []struct {
						Pattern    string `json:"pattern"`
						Confidence int    `json:"confidence"`
					} `json:"confidence"`
					ConfidenceTotal int    `json:"confidenceTotal"`
					Icon            string `json:"icon"`
					Website         string `json:"website"`
					Categories      []struct {
						Name     string `json:"name"`
						Priority string `json:"priority"`
					} `json:"categories"`
				} `json:"data"`
			} `json:"wappa"`
			Done struct {
				State string `json:"state"`
				Data  struct {
					State string `json:"state"`
				} `json:"data"`
			} `json:"done"`
		} `json:"processors"`
	} `json:"meta"`
	Task struct {
		UUID       string    `json:"uuid"`
		Time       time.Time `json:"time"`
		URL        string    `json:"url"`
		Visibility string    `json:"visibility"`
		Options    struct {
			Useragent string `json:"useragent"`
		} `json:"options"`
		Method        string `json:"method"`
		Source        string `json:"source"`
		UserAgent     string `json:"userAgent"`
		ReportURL     string `json:"reportURL"`
		ScreenshotURL string `json:"screenshotURL"`
		DomURL        string `json:"domURL"`
	} `json:"task"`
	Page struct {
		URL     string `json:"url"`
		Domain  string `json:"domain"`
		Country string `json:"country"`
		City    string `json:"city"`
		Server  string `json:"server"`
		IP      string `json:"ip"`
		Ptr     string `json:"ptr"`
		Asn     string `json:"asn"`
		Asnname string `json:"asnname"`
	} `json:"page"`
	Lists struct {
		Ips          []string `json:"ips"`
		Countries    []string `json:"countries"`
		Asns         []string `json:"asns"`
		Domains      []string `json:"domains"`
		Servers      []string `json:"servers"`
		Urls         []string `json:"urls"`
		LinkDomains  []string `json:"linkDomains"`
		Certificates []struct {
			SubjectName string `json:"subjectName"`
			Issuer      string `json:"issuer"`
			ValidFrom   int    `json:"validFrom"`
			ValidTo     int    `json:"validTo"`
		} `json:"certificates"`
		Hashes []string `json:"hashes"`
	} `json:"lists"`
}
