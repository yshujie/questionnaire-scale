package port

import (
	"context"

	"github.com/yshujie/questionnaire-scale/internal/apiserver/domain/questionnaire"
)

// QuestionnaireRepositoryMySQL 问卷存储库接口（出站端口）
// 定义了与存储相关的所有操作契约
type QuestionnaireRepositoryMySQL interface {
	// 基础 CRUD 操作
	Create(ctx context.Context, questionnaire *questionnaire.Questionnaire) error
	FindByID(ctx context.Context, id uint64) (*questionnaire.Questionnaire, error)
	FindByCode(ctx context.Context, code string) (*questionnaire.Questionnaire, error)
	FindList(ctx context.Context, page, pageSize int, conditions map[string]string) ([]*questionnaire.Questionnaire, error)
	CountWithConditions(ctx context.Context, conditions map[string]string) (int64, error)
	Update(ctx context.Context, questionnaire *questionnaire.Questionnaire) error
	Remove(ctx context.Context, id uint64) error
}

// QuestionnaireRepository 问卷存储库接口（出站端口）
// 定义了与存储相关的所有操作契约
type QuestionnaireRepositoryMongo interface {
	Create(ctx context.Context, qDomain *questionnaire.Questionnaire) error
	FindByCode(ctx context.Context, code string) (*questionnaire.Questionnaire, error)
	FindByCodeVersion(ctx context.Context, code, version string) (*questionnaire.Questionnaire, error)
	Update(ctx context.Context, qDomain *questionnaire.Questionnaire) error
	Remove(ctx context.Context, code string) error
	HardDelete(ctx context.Context, code string) error
	ExistsByCode(ctx context.Context, code string) (bool, error)
	FindActiveQuestionnaires(ctx context.Context) ([]*questionnaire.Questionnaire, error)
}
