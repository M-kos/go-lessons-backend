package v1

type CreateUserRequest struct {
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Country  string `json:"country,omitempty"`
	City     string `json:"city,omitempty"`
	Zip      string `json:"zip,omitempty"`
	Street   string `json:"street,omitempty"`
}

type CreateUserResponse struct {
	ID int64 `json:"id"`
}

type GetUserByEmailRequest struct {
	Email string `json:"email"`
}

type ListUserByEmailsRequest struct {
	Emails []string `json:"emails"`
}

type UserResponse struct {
	ID         int64  `json:"id"`
	Email      string `json:"email"`
	FullName   string `json:"full_name"`
	CreateTime string `json:"create_time"`
	Country    string `json:"country"`
	City       string `json:"city"`
	Zip        string `json:"zip"`
	Street     string `json:"street"`
}

type ListUserResponse struct {
	Result []UserResponse `json:"result"`
}

type UpdateUserRequest struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Country  string `json:"country,omitempty"`
	City     string `json:"city,omitempty"`
	Zip      string `json:"zip,omitempty"`
	Street   string `json:"street,omitempty"`
}
