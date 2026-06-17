package item_stack_operation

import "github.com/EmptyDea-Team/EmptyDea-core-client/resources_control"

// BeaconPayment 指示信标支付操作。
type BeaconPayment struct {
	PaymentPath     resources_control.SlotLocation
	PrimaryEffect   int32
	SecondaryEffect int32
}

func (BeaconPayment) ID() uint8 {
	return IDItemStackOperationHighLevelBeaconPayment
}

func (BeaconPayment) CanInline() bool {
	return false
}

// BeaconPaymentFromInventory 指示使用背包内物品执行信标支付的操作。
type BeaconPaymentFromInventory struct {
	PaymentSlot     resources_control.SlotID
	PrimaryEffect   int32
	SecondaryEffect int32
}

func (BeaconPaymentFromInventory) ID() uint8 {
	return IDItemStackOperationHighLevelBeaconPayment
}

func (BeaconPaymentFromInventory) CanInline() bool {
	return false
}
