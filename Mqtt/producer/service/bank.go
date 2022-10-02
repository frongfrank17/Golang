package service

import (
	"errors"
	"events"
	"log"
	commands "producer/command"

	"github.com/google/uuid"
)

type BankService interface {
	OpenAccount(command commands.OpenAccountCommand) (id string, err error)
	DepositFund(command commands.DepositFundCommand) error
	WithdrawFund(command commands.WithdrawFundCommand) error
	CloseAccount(command commands.CloseAccountCommand) error
}

type bankService struct {
	eventProducer EventProducer
}

func NewBankService(eventProducer EventProducer) bankService {
	return bankService{eventProducer}
}

func (obj bankService) OpenAccount(command commands.OpenAccountCommand) (id string, err error) {

	if command.AccountName == "" || command.AccountType == 0 || command.OpeningBalance == 0 {
		return "", errors.New("bad request")
	}

	event := events.OpenAccountEvent{
		ID:             uuid.NewString(),
		AccountName:    command.AccountName,
		AccountType:    command.AccountType,
		OpeningBalance: command.OpeningBalance,
	}

	log.Printf("%#v", event)
	return event.ID, obj.eventProducer.Produce(event)
}

func (obj bankService) DepositFund(command commands.DepositFundCommand) error {
	if command.ID == "" || command.Amount == 0 {
		return errors.New("bad request")
	}

	event := events.DepositFundEvent{
		ID:     command.ID,
		Amount: command.Amount,
	}

	log.Printf("%#v", event)
	return obj.eventProducer.Produce(event)
}

func (obj bankService) WithdrawFund(command commands.WithdrawFundCommand) error {
	if command.ID == "" || command.Amount == 0 {
		return errors.New("bad request")
	}

	event := events.WithdrawFundEvent{
		ID:     command.ID,
		Amount: command.Amount,
	}

	log.Printf("%#v", event)
	return obj.eventProducer.Produce(event)
}

func (obj bankService) CloseAccount(command commands.CloseAccountCommand) error {
	if command.ID == "" {
		return errors.New("bad request")
	}

	event := events.CloseAccountEvent{
		ID: command.ID,
	}

	log.Printf("%#v", event)
	return obj.eventProducer.Produce(event)
}
