package question_types

import (
	"github.com/yshujie/questionnaire-scale/internal/apiserver/domain/question"
	"github.com/yshujie/questionnaire-scale/internal/apiserver/domain/question/calculation"
	"github.com/yshujie/questionnaire-scale/internal/apiserver/domain/question/option"
	"github.com/yshujie/questionnaire-scale/internal/apiserver/domain/question/validation"
)

// RadioQuestion 单选问题
type RadioQuestion struct {
	BaseQuestion
	validation.ValidationAbility
	calculation.CalculationAbility

	options []option.Option
}

// 注册单选问题
func init() {
	RegisterQuestionFactory(question.QuestionTypeRadio, func(builder *QuestionBuilder) question.Question {
		// 创建单选问题
		q := newRadioQuestion(builder.GetCode(), builder.GetTitle())

		// 设置选项
		q.setOptions(builder.GetOptions())

		// 设置校验规则
		for _, rule := range builder.GetValidationRules() {
			q.addValidationRule(rule)
		}

		// 设置计算规则
		if builder.GetCalculationRule() != nil {
			q.setCalculationRule(builder.GetCalculationRule())
		}
		return q
	})
}

// NewRadioQuestion 创建单选问题
func newRadioQuestion(code question.QuestionCode, title string) *RadioQuestion {
	return &RadioQuestion{
		BaseQuestion: NewBaseQuestion(code, title, question.QuestionTypeRadio),
	}
}

// setOptions 设置选项
func (q *RadioQuestion) setOptions(options []option.Option) {
	q.options = options
}

// AddValidationRule 添加校验规则
func (q *RadioQuestion) addValidationRule(rule validation.ValidationRule) {
	q.ValidationAbility.AddValidationRule(rule)
}

// setCalculationRule 设置计算规则
func (q *RadioQuestion) setCalculationRule(rule *calculation.CalculationRule) {
	q.CalculationAbility.SetCalculationRule(rule)
}

// GetOptions 获取选项
func (q *RadioQuestion) GetOptions() []option.Option {
	return q.options
}

// GetValidationRules 获取校验规则 - 重写BaseQuestion的默认实现
func (q *RadioQuestion) GetValidationRules() []validation.ValidationRule {
	return q.ValidationAbility.GetValidationRules()
}

// GetCalculationRule 获取计算规则 - 重写BaseQuestion的默认实现
func (q *RadioQuestion) GetCalculationRule() *calculation.CalculationRule {
	return q.CalculationAbility.GetCalculationRule()
}
