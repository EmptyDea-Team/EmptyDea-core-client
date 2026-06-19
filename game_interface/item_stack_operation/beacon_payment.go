package item_stack_operation

import "github.com/EmptyDea-Team/EmptyDea-core-client/resources_control"

// BeaconPayment 指示信标支付操作。
type BeaconPayment struct {
	// PaymentPath 是支付物品所在位置。
	PaymentPath resources_control.SlotLocation
	// PrimaryEffect 是信标主效果 ID。
	PrimaryEffect int32
	// SecondaryEffect 是信标副效果 ID。
	SecondaryEffect int32
}

// ID 返回信标支付操作编号。
func (BeaconPayment) ID() uint8 {
	return IDItemStackOperationHighLevelBeaconPayment
}

// CanInline 返回该操作是否可以内联到普通物品堆栈请求中。
func (BeaconPayment) CanInline() bool {
	return false
}

// BeaconPaymentFromInventory 指示使用背包内物品执行信标支付的操作。
type BeaconPaymentFromInventory struct {
	// PaymentSlot 是背包内支付物品槽位。
	PaymentSlot resources_control.SlotID
	// PrimaryEffect 是信标主效果 ID。
	PrimaryEffect int32
	// SecondaryEffect 是信标副效果 ID。
	SecondaryEffect int32
}

// ID 返回信标支付操作编号。
func (BeaconPaymentFromInventory) ID() uint8 {
	return IDItemStackOperationHighLevelBeaconPayment
}

// CanInline 返回该操作是否可以内联到普通物品堆栈请求中。
func (BeaconPaymentFromInventory) CanInline() bool {
	return false
}
