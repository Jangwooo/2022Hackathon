package object

import "github.com/Jangwooo/2022Hackathon/interner/mysql/model"

type DTO interface {
	ConvertToDAO() model.DAO
}
