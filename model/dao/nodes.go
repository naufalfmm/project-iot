package dao

import nodeDTO "github.com/naufalfmm/project-iot/model/dto/node"

type Nodes []Node

func (ns Nodes) Len() int {
	return len(ns)
}

func (ns Nodes) ToResponsesDTO() []nodeDTO.ResponseDTO {
	nsLen := ns.Len()
	nsDTO := make([]nodeDTO.ResponseDTO, nsLen)

	for i := 0; i < nsLen; i++ {
		nsDTO[i] = ns[i].ToResponseDTO()
	}

	return nsDTO
}
