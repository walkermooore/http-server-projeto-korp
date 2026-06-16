package services

import (
	"http-server-projeto-korp/models"
	"time"
)

func GerarRespostaJson() models.Resposta {

	agora := time.Now().UTC()

	return models.Resposta{
		Nome:    "Projeto Korp",
		Horario: agora.String(),
	}
}
