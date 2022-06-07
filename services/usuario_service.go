package services

import (
	"errors"
	"strings"

	"github.com/PQMISSCMP/go-rest/models"
	"github.com/PQMISSCMP/go-rest/repos"
)

type UsuarioService struct {
	repository *repos.UsuarioRepository
}

func NewUsuarioService(r *repos.UsuarioRepository) *UsuarioService {
	return &UsuarioService{repository: r}
}

func (s *UsuarioService) Logging(rs *models.RequestSession) (*models.USuario, error) {
	lenUsr := len(strings.TrimSpace(rs.Username))
	lenPwd := len(strings.TrimSpace(rs.Password))

	if lenUsr < 5 || lenPwd < 5 {
		return nil, errors.New("El nombre de usuario y la contraseña deben tener más de 5 caracteres")
	}

	return s.repository.Loging(rs)
}
