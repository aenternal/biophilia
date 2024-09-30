package services

import (
	"biophilia/internal/domain/entities"
	"biophilia/internal/domain/interfaces/data"
	"biophilia/internal/domain/interfaces/domain"
	"github.com/redis/go-redis/v9"
	"log/slog"
	"strings"
)

type BiomoleculeService struct {
	log                   *slog.Logger
	biomoleculeRepository data.BiomoleculeRepository
	storageRepository     data.StorageRepository
	imageService          domain.ImageService
	blastRepository       data.BlastClient
	redisClient           *redis.Client
}

func NewBiomoleculeService(
	log *slog.Logger,
	biomoleculeRepository data.BiomoleculeRepository,
	storageRepository data.StorageRepository,
	imageService domain.ImageService,
	blastRepository data.BlastClient,
	redisClient *redis.Client,
) *BiomoleculeService {
	return &BiomoleculeService{
		log:                   log,
		biomoleculeRepository: biomoleculeRepository,
		storageRepository:     storageRepository,
		imageService:          imageService,
		blastRepository:       blastRepository,
		redisClient:           redisClient,
	}
}

func (service *BiomoleculeService) AddBiomolecule(biomolecule entities.AddBiomoleculeRequest) error {
	return service.biomoleculeRepository.Add(biomolecule)
}

func (service *BiomoleculeService) GetBiomolecules() ([]entities.Biomolecule, error) {
	return service.biomoleculeRepository.GetAll()
}

func (service *BiomoleculeService) GetBiomoleculeByID(id int) (*entities.Biomolecule, error) {
	return service.biomoleculeRepository.GetByID(id)
}

func (service *BiomoleculeService) UpdateBiomolecule(id int, biomolecule entities.UpdateBiomoleculeRequest) error {
	return service.biomoleculeRepository.Update(id, biomolecule)
}

func (service *BiomoleculeService) DeleteBiomolecule(id int) error {
	return service.biomoleculeRepository.Delete(id)
}

func (_ *BiomoleculeService) transcribe(dna string) string {
	return strings.ReplaceAll(dna, "T", "U")
}

func (_ *BiomoleculeService) reverseTranscribe(rna string) string {
	return strings.ReplaceAll(rna, "U", "T")
}

func (_ *BiomoleculeService) translate(mrna string) string {
	var peptide strings.Builder
	for i := 0; i < len(mrna)-2; i += 3 {
		codon := mrna[i : i+3]
		if aminoAcid, ok := entities.CodonTable()[codon]; ok {
			if aminoAcid == "*" {
				break
			}
			peptide.WriteString(aminoAcid)
		}
	}
	return peptide.String()
}

//func countSequenceUnits(sequence string) map[string]int {
//	counts := make(map[string]int)
//	for _, aa := range sequence {
//		counts[string(aa)]++
//	}
//	return counts
//}
