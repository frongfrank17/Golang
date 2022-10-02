package service

import (
	"consumer/repository"
	"encoding/json"
	"events"
	"fmt"
	"log"
	"reflect"
)

type EventService interface {
	Service(topic string, eventBytes []byte)
}
type accountEventService struct {
	accountRepo repository.AccountRepository
}

func NewAccountService(accountRepo repository.AccountRepository) EventService {
	return accountEventService{accountRepo: accountRepo}
}

func (account accountEventService) Service(topic string, eventBytes []byte) {
	switch topic {
	case reflect.TypeOf(events.OpenAccountEvent{}).Name():
		event := &events.OpenAccountEvent{}
		err := json.Unmarshal(eventBytes, event)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Print(event)
		bankAcc := repository.BankAccount{
			ID:          event.ID,
			AccountName: event.AccountName,
			AccountType: event.AccountType,
			Balance:     event.OpeningBalance,
		}
		err = account.accountRepo.Save(bankAcc)
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("[%v] %#v", topic, event)
	case reflect.TypeOf(events.DepositFundEvent{}).Name():
		event := &events.DepositFundEvent{}
		err := json.Unmarshal(eventBytes, event)
		if err != nil {
			log.Println(err)
			return
		}
		FindOne, err := account.accountRepo.FindByID(event.ID)
		if err != nil {
			log.Println(err)
			return
		}
		FindOne.Balance += event.Amount
		bankAcc := repository.BankAccount{
			ID:          FindOne.ID,
			AccountName: FindOne.AccountName,
			AccountType: FindOne.AccountType,
			Balance:     FindOne.Balance,
		}
		err = account.accountRepo.Save(bankAcc)
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("[%v] %#v", topic, event)

	case reflect.TypeOf(events.WithdrawFundEvent{}).Name():
		event := &events.WithdrawFundEvent{}
		err := json.Unmarshal(eventBytes, event)
		if err != nil {
			log.Println(err)
			return
		}
		bankAccount, err := account.accountRepo.FindByID(event.ID)
		if err != nil {
			log.Println(err)
			return
		}
		bankAccount.Balance -= event.Amount
		bankAcc := repository.BankAccount{
			ID:          bankAccount.ID,
			AccountName: bankAccount.AccountName,
			AccountType: bankAccount.AccountType,
			Balance:     bankAccount.Balance,
		}
		err = account.accountRepo.Save(bankAcc)
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("[%v] %#v", topic, event)
	case reflect.TypeOf(events.CloseAccountEvent{}).Name():
		event := &events.CloseAccountEvent{}
		err := json.Unmarshal(eventBytes, event)
		if err != nil {
			log.Println(err)
			return
		}
		err = account.accountRepo.Delete(event.ID)
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("[%v] %#v", topic, event)

	default:
		log.Println("no event handler")
	}
}
