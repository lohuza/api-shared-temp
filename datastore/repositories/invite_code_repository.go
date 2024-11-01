package repositories

import (
	"github.com/lohuza/api-shared-temp/models/invitationmodel"
	"gorm.io/gorm"
)

type InviteCodeRepository interface {
	GetInviteCode(code string) (*invitationmodel.InviteCode, error)
}

type inviteCodeRepository struct {
	db *gorm.DB
}

func NewInviteCodeRepository(db *gorm.DB) InviteCodeRepository {
	return &inviteCodeRepository{
		db: db,
	}
}

func (repo *inviteCodeRepository) GetInviteCode(inviteCode string) (*invitationmodel.InviteCode, error) {
	model := new(invitationmodel.InviteCode)
	err := repo.db.Table((*invitationmodel.InviteCode)(nil).TableName()).First(model, "code = ?", inviteCode).Error
	return model, err
}
