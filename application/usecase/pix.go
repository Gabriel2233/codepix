package usecase

import (
	"fmt"

	"github.com/codepix/domain/model"
)

type PixUseCase struct {
	PixRepository model.PixKeyRepositoryInterface
}

func (p *PixUseCase) RegisterKey(key, accountId, kind string) (*model.PixKey, error) {
	account, err := p.PixRepository.FindAccount(accountId)

	if err != nil {
		return nil, err
	}

	pixKey, err := model.NewPixKey(kind, account, key)

	if err != nil {
		return nil, err
	}

	p.PixRepository.Register(pixKey)

	if pixKey.ID == "" {
		return nil, fmt.Errorf("unable to create key at the moment")
	}

	return pixKey, nil
}

func (p *PixUseCase) FindKey(kind string, key string) (*model.PixKey, error) {
	pixKey, err := p.PixRepository.FindByKind(key, kind)

	if err != nil {
		return nil, err
	}

	return pixKey, nil
}
