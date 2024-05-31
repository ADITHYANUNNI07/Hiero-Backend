package models

type EmployerLogin struct {
	Email    string `json:"email" binding:"required" validate:"required"`
	Password string `json:"password" binding:"required" validate:"min=6,max=20"`
}

type EmployerSignUp struct {
	ID                  uint   `json:"id"`
	CompanyName         string `json:"company_name"`
	Industry            string `json:"industry"`
	CompanySize         int    `json:"company_size"`
	Website             string `json:"website"`
	HeadquartersAddress string `json:"headquarters_address"`
	AboutCompany        string `json:"about_company"`
	ContactEmail        string `json:"contact_email"`
	ContactPhoneNumber  uint   `json:"contact_phone_number"`
	Password            string `json:"password"`
}

type EmployerDetailsResponse struct {
	ID                  uint   `json:"id"`
	CompanyName         string `json:"company_name"`
	Industry            string `json:"industry"`
	CompanySize         int    `json:"company_size"`
	Website             string `json:"website"`
	HeadquartersAddress string `json:"headquarters_address"`
	AboutCompany        string `json:"about_company"`
	ContactEmail        string `json:"contact_email"`
	ContactPhoneNumber  uint   `json:"contact_phone_number"`
}

type Employer struct {
	ID                   uint   `json:"id" gorm:"uniquekey; not null"`
	Company_name         string `json:"company_name" gorm:"validate:required"`
	Industry             string `json:"industry" gorm:"validate:required"`
	Company_size         int    `json:"company_size" gorm:"validate:required"`
	Website              string `json:"website"`
	Headquarters_address string `json:"headquarters_address"`
	About_company        string `json:"about_company" gorm:"type:text"`
	Contact_email        string `json:"contact_email" gorm:"validate:required"`
	Contact_phone_number uint   `json:"contact_phone_number" gorm:"type:numeric"`
	Password             string `json:"password" gorm:"validate:required"`
	Created_at           string `json:"created_at"`
	Updated_at           string `json:"updated_at"`
	Deleted_at           string `json:"deleted_at"`
}

type TokenEmployer struct {
	Employer EmployerDetailsResponse
	Token    string
}

type EmployerDetails struct {
	ID                  uint   `json:"id"`
	CompanyName         string `json:"company_name"`
	Industry            string `json:"industry"`
	CompanySize         int    `json:"company_size"`
	Website             string `json:"website"`
	HeadquartersAddress string `json:"headquarters_address"`
	AboutCompany        string `json:"about_company"`
	ContactEmail        string `json:"contact_email"`
	ContactPhoneNumber  uint   `json:"contact_phone_number"`
}
