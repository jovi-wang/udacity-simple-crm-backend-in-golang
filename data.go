package main

type Customer struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Contacted bool   `json:"contacted"`
}

var customerData = map[string]Customer{
	"iNov-6iQ9J": {
		Name:      "Urbanus",
		Role:      "Administrative Officer",
		Email:     "uabriani0@vistaprint.com",
		Phone:     "772-772-1211",
		Contacted: true,
	},
	"VVDvc6i99J": {
		Name:      "Bernetta",
		Role:      "Structural Engineer",
		Email:     "bchelnam0@google.co.uk",
		Phone:     "232-482-1454",
		Contacted: true,
	},
	"NVovc6-QQy": {
		Name:      "Tova",
		Role:      "Accountant",
		Email:     "tgreally0@youtube.com",
		Phone:     "288-930-0374",
		Contacted: false,
	},
	"VVoveui9QC": {
		Name:      "Whitby",
		Role:      "Quality Control Specialist",
		Email:     "wwoolcocks1@irs.gov",
		Phone:     "212-322-4456",
		Contacted: true,
	},
	"tFmGc6iQQs": {
		Name:      "Elsi",
		Role:      "Graphic Designer",
		Email:     "equinnelly4@java.com",
		Phone:     "905-575-7452",
		Contacted: false,
	},
	"KpTvcui99k": {
		Name:      "Marinna",
		Role:      "Senior Quality Engineer",
		Email:     "mstutt5@ameblo.jp",
		Phone:     "829-749-7736",
		Contacted: false,
	},
	"tFTvcu-QQt": {
		Name:      "Marlow",
		Role:      "Environmental Tech",
		Email:     "mrabbage6@mlb.com",
		Phone:     "870-415-3764",
		Contacted: false,
	},
	"KFmGeu-Q9O": {
		Name:      "Herman",
		Role:      "Community Outreach Specialist",
		Email:     "hsegges7@1688.com",
		Phone:     "420-667-6621",
		Contacted: true,
	},
}
