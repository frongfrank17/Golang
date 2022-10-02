package repository

import "gorm.io/gorm"

type BankAccount struct {
	gorm.Model
	ID          string  `json:"id"`
	AccountName string  `json:"name"`
	AccountType int     `json:"type"`
	Balance     float64 `json:"balance"`
}

type AccountRepository interface {
	Save(bankAccount BankAccount) error
	Delete(id string) error
	FindAll() ([]BankAccount, error)
	FindByID(id string) (*BankAccount, error)
}

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {

	return accountRepository{db}
}

func (obj accountRepository) Save(bankAccount BankAccount) error {
	err := obj.db.Create(&bankAccount).Error
	if err != nil {

		return err
	}
	return nil
}

func (obj accountRepository) Delete(id string) error {
	err := obj.db.Model(BankAccount{}).Where("id=?", id).Delete(&BankAccount{}).Error
	if err != nil {

		return err
	}
	return nil
}

func (obj accountRepository) FindAll() ([]BankAccount, error) {
	var bankAcc []BankAccount
	err := obj.db.Find(&bankAcc).Error
	if err != nil {
		return nil, err
	}
	return bankAcc, nil
}

func (obj accountRepository) FindByID(id string) (*BankAccount, error) {
	var bankAcc BankAccount
	err := obj.db.Where("id=?", id).First(&bankAcc).Error
	if err != nil {
		return nil, err
	}
	return &bankAcc, nil
}
