package model

import "github.com/Jangwooo/2022Hackathon/interner/domain/object"

type DAO interface {
	ConvertToDTO() object.DTO
}
