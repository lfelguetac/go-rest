package repos

import (
	"github.com/PQMISSCMP/go-rest/models"
	"gorm.io/gorm"
)

type UsuarioRepository struct {
	conn *gorm.DB
}

func NewUsuarioRepo(c *gorm.DB) *UsuarioRepository {
	return &UsuarioRepository{conn: c}
}

func (r *UsuarioRepository) Loging(rs *models.RequestSession) (*models.USuario, error) {
	// return r.conn.Create(u).Error
	return nil, nil
}
