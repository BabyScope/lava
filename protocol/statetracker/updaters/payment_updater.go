package updaters

import (
	"sync"

	"github.com/lavanet/lava/protocol/rpcprovider/rewardserver"
	"golang.org/x/net/context"
)

const (
	CallbackKeyForPaymentUpdate = "payment-update"
)

type PaymentUpdatable interface {
	PaymentHandler(*rewardserver.PaymentRequest)
	Description() string
}

type PaymentUpdater struct {
	lock             sync.RWMutex
	paymentUpdatable map[string]*PaymentUpdatable
	eventTracker     *EventTracker
}

func NewPaymentUpdater(eventTracker *EventTracker) *PaymentUpdater {
	return &PaymentUpdater{paymentUpdatable: map[string]*PaymentUpdatable{}, eventTracker: eventTracker}
}

func (pu *PaymentUpdater) RegisterPaymentUpdatable(ctx context.Context, paymentUpdatable *PaymentUpdatable) {
	pu.lock.Lock()
	defer pu.lock.Unlock()
	pu.paymentUpdatable[(*paymentUpdatable).Description()] = paymentUpdatable
}

func (pu *PaymentUpdater) UpdaterKey() string {
	return CallbackKeyForPaymentUpdate
}

func (pu *PaymentUpdater) updateInner() {
	pu.lock.RLock()
	defer pu.lock.RUnlock()
	payments, err := pu.eventTracker.getLatestPaymentEvents()
	if err != nil {
		return
	}
	for _, payment := range payments {
		updatable, foundUpdatable := pu.paymentUpdatable[payment.Description]
		if foundUpdatable {
			(*updatable).PaymentHandler(payment)
		}
	}
}

func (pu *PaymentUpdater) Reset(latestBlock int64) {
	// in case we need a reset we don't have much to do as we might have lost some data due to pruning. we can just continue parsing our transactions
	// TODO: we can fetch the latest transactions for our account and see if we have some missing data but for that we need to keep track all transactions per provider account
	pu.updateInner()
}

func (pu *PaymentUpdater) Update(latestBlock int64) {
	pu.updateInner()
}
