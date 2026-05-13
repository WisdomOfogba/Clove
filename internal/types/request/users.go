package request

// RegisterCustomerRequest represents customer registration
type RegisterCustomerRequest struct {
	FullName string `form:"full_name" json:"full_name" name:"Full Name" validate:"required,lte=40"`
	Email    string `form:"email" json:"email" name:"Email Address" validate:"required,email"`
	Phone    string `form:"phone" json:"phone" name:"Phone Number" validate:"required,gte=10"`
	Password string `form:"password" json:"password" name:"Password" validate:"required,lte=30"`
}

// CustomerLoginRequest represents customer login
type CustomerLoginRequest struct {
	Email    string `form:"email" json:"email" name:"Email Address" validate:"required,email"`
	Password string `form:"password" json:"password" name:"Password" validate:"required,lte=30"`
}

// UpdateCustomerProfileRequest represents customer profile update
type UpdateCustomerProfileRequest struct {
	FullName string `form:"full_name" json:"full_name" name:"Full Name"`
	Phone    string `form:"phone" json:"phone" name:"Phone Number"`
}
