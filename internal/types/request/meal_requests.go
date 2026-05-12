package request

// ==================== MEALS ====================

// CreateMealRequest represents creating a new meal
type CreateMealRequest struct {
	Name        string `form:"name" json:"name" name:"Meal Name"`
	Description string `form:"description" json:"description" name:"Description"`
	Price       int64  `form:"price" json:"price" name:"Price"` // in kobo
	Category    string `form:"category" json:"category" name:"Category"`
}

// UpdateMealRequest represents updating an existing meal
type UpdateMealRequest struct {
	Name        string `form:"name" json:"name" name:"Meal Name"`
	Description string `form:"description" json:"description" name:"Description"`
	Price       int64  `form:"price" json:"price" name:"Price"`
	Category    string `form:"category" json:"category" name:"Category"`
	Status      string `form:"status" json:"status" name:"Status"`
}

// UploadMealPicturesRequest represents uploading pictures for a meal
type UploadMealPicturesRequest struct {
	MealID     string `form:"meal_id" json:"meal_id" name:"Meal ID"`
	IsPrimary  bool   `form:"is_primary" json:"is_primary" name:"Is Primary"`
	// Pictures handled via multipart/form-data
}

// ==================== CUSTOMERS ====================

// RegisterCustomerRequest represents customer registration
type RegisterCustomerRequest struct {
	FullName string `form:"full_name" json:"full_name" name:"Full Name"`
	Email    string `form:"email" json:"email" name:"Email Address"`
	Phone    string `form:"phone" json:"phone" name:"Phone Number"`
	Password string `form:"password" json:"password" name:"Password"`
}

// CustomerLoginRequest represents customer login
type CustomerLoginRequest struct {
	Email    string `form:"email" json:"email" name:"Email Address"`
	Password string `form:"password" json:"password" name:"Password"`
}

// UpdateCustomerProfileRequest represents customer profile update
type UpdateCustomerProfileRequest struct {
	FullName string `form:"full_name" json:"full_name" name:"Full Name"`
	Phone    string `form:"phone" json:"phone" name:"Phone Number"`
}

// ==================== ORDERS ====================

// PlaceOrderRequest represents placing a new order
type PlaceOrderRequest struct {
	VendorID        string                    `json:"vendor_id" name:"Vendor ID"`
	Items           []OrderItemRequest        `json:"items" name:"Order Items"`
	DeliveryAddress string                    `json:"delivery_address" name:"Delivery Address"`
}

// OrderItemRequest represents an item in the order
type OrderItemRequest struct {
	MealID   string `json:"meal_id" name:"Meal ID"`
	Quantity int    `json:"quantity" name:"Quantity"`
}

// UpdateOrderStatusRequest represents updating order status (vendor/admin)
type UpdateOrderStatusRequest struct {
	Status string `form:"status" json:"status" name:"Status"`
}

// ==================== REVIEWS ====================

// AddReviewRequest represents adding a review for an order/meal
type AddReviewRequest struct {
	OrderID string `json:"order_id" name:"Order ID"`
	MealID  string `json:"meal_id" name:"Meal ID"`
	Rating  int    `json:"rating" name:"Rating"` // 1-5
	Comment string `json:"comment" name:"Comment"`
}

// ==================== SEARCH & DISCOVERY ====================

// SearchMealsRequest represents searching for meals
type SearchMealsRequest struct {
	Query    string `form:"query" json:"query" name:"Search Query"`
	Category string `form:"category" json:"category" name:"Category"`
	VendorID string `form:"vendor_id" json:"vendor_id" name:"Vendor ID"`
	Page     int    `form:"page" json:"page" name:"Page"`
	Limit    int    `form:"limit" json:"limit" name:"Limit"`
}

// SearchVendorsRequest represents searching for vendors/restaurants
type SearchVendorsRequest struct {
	Query string `form:"query" json:"query" name:"Search Query"`
	State string `form:"state" json:"state" name:"State"`
	Page  int    `form:"page" json:"page" name:"Page"`
	Limit int    `form:"limit" json:"limit" name:"Limit"`
}