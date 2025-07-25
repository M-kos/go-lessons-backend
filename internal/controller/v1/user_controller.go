package v1

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/meetmorrowsolonmars/education-pet-project/internal/domain"
)

type UserStore interface {
	Create(ctx context.Context, user domain.User) (domain.User, error)
	GetByEmail(ctx context.Context, email string) (domain.User, error)
	ListUsersByEmail(ctx context.Context, emails []string) ([]domain.User, error)
	UpdateUser(ctx context.Context, user domain.User) error
}

type UserController struct {
	userStore UserStore
}

func NewUserController(userStore UserStore) *UserController {
	return &UserController{
		userStore: userStore,
	}
}

func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	req, err := DecodeRequest[CreateUserRequest](r)
	if err != nil {
		slog.Error("Decode create user request", slog.String("error", err.Error()))

		err = WriteErrorResponse(w, http.StatusBadRequest, "Invalid request body: %s", err)
		if err != nil {
			slog.Error("Write error response", slog.String("error", err.Error()))
		}

		return
	}

	user, err := c.userStore.Create(r.Context(), domain.User{
		Email:    req.Email,
		FullName: req.FullName,
	})
	if err != nil {
		slog.Error("Create user", slog.String("error", err.Error()))

		err = WriteErrorResponse(w, http.StatusInternalServerError, "Create user error")
		if err != nil {
			slog.Error("Write error response", slog.String("error", err.Error()))
		}

		return
	}

	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(CreateUserResponse{
		ID: user.ID,
	})
	if err != nil {
		slog.Error("Write create user response", slog.String("error", err.Error()))
	}
}

func (c *UserController) UserByEmail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	req, err := DecodeRequest[GetUserByEmailRequest](r)
	if err != nil {
		slog.Error("Decode get user request", slog.String("error", err.Error()))

		err = WriteErrorResponse(w, http.StatusBadRequest, "Invalid request body: %s", err)
		if err != nil {
			slog.Error("Write error response", slog.String("error", err.Error()))
		}

		return
	}

	if req.Email == "" {
		err := WriteErrorResponse(w, http.StatusBadRequest, "Invalid email")
		if err != nil {
			slog.Error("Write error response", slog.String("error", err.Error()))
		}

		return
	}

	user, err := c.userStore.GetByEmail(r.Context(), req.Email)
	if err != nil {
		slog.Error("Get user by email", slog.String("error", err.Error()))

		err = WriteErrorResponse(w, http.StatusInternalServerError, "Get user error")
		if err != nil {
			slog.Error("Write error response", slog.String("error", err.Error()))
		}

		return
	}

	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(UserResponse{
		ID:         user.ID,
		Email:      user.Email,
		FullName:   user.FullName,
		CreateTime: user.CreateTime.String(),
	})
	if err != nil {
		slog.Error("Write get user response", slog.String("error", err.Error()))
	}
}

func (c *UserController) ListUserByEmail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	req, err := DecodeRequest[ListUserByEmailsRequest](r)
	if err != nil {
		slog.Error("Decode list user request", slog.String("error", err.Error()))

		err = WriteErrorResponse(w, http.StatusBadRequest, "Invalid request body: %s", err)
		if err != nil {
			slog.Error("Write error response", slog.String("error", err.Error()))
		}

		return
	}

	users, err := c.userStore.ListUsersByEmail(r.Context(), req.Emails)
	if err != nil {
		slog.Error("List user by email", slog.String("error", err.Error()))

		err = WriteErrorResponse(w, http.StatusInternalServerError, "Get user error")
		if err != nil {
			slog.Error("Write error response", slog.String("error", err.Error()))
		}

		return
	}

	w.WriteHeader(http.StatusOK)

	usersRespons := make([]UserResponse, len(users))

	for _, u := range users {
		usersRespons = append(usersRespons, UserResponse{
			ID:         u.ID,
			Email:      u.Email,
			FullName:   u.FullName,
			CreateTime: u.CreateTime.String(),
		})
	}

	err = json.NewEncoder(w).Encode(ListUserResponse{
		Result: usersRespons,
	})
	if err != nil {
		slog.Error("Write list user response", slog.String("error", err.Error()))
	}
}

func (c *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	req, err := DecodeRequest[UpdateUserRequest](r)
	if err != nil {
		slog.Error("Decode update user request", slog.String("error", err.Error()))

		err = WriteErrorResponse(w, http.StatusBadRequest, "Invalid request body: %s", err)
		if err != nil {
			slog.Error("Write error response", slog.String("error", err.Error()))
		}

		return
	}

	err = c.userStore.UpdateUser(r.Context(), domain.User{
		ID:       req.ID,
		Email:    req.Email,
		FullName: req.FullName,
	})
	if err != nil {
		slog.Error("update user", slog.String("error", err.Error()))

		err = WriteErrorResponse(w, http.StatusInternalServerError, "update user error")
		if err != nil {
			slog.Error("Write error response", slog.String("error", err.Error()))
		}

		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *UserController) Register(mux *http.ServeMux) {
	mux.HandleFunc("POST /v1/users", c.CreateUser)
	mux.HandleFunc("POST /v1/users/getOne", c.UserByEmail)
	mux.HandleFunc("POST /v1/users/list", c.ListUserByEmail)
	mux.HandleFunc("PUT /v1/users", c.UpdateUser)
}
