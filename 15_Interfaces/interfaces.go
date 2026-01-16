package main

import "fmt"

//Lets implement Payment GateWay without using any interface and see what happens if
//We had to implement multiple gateways and need to switch them as per need

//First Run Part 1 , then Comment it and run Part two to see the difference

//Part1

/*---------------------------------------------------------------------------------*/

// type Payment struct{
// 	gateway Stripe
// }

// func (p Payment) makePayment(amount float32){
// 	stripePaymentGW := Stripe{}
// 	stripePaymentGW.pay(amount)
// }

// type RazorPay struct {}
// func (r RazorPay) pay(amount float32){
// 	fmt.Println("Payment using Razor Payment Gateway" , amount)
// }

// type Stripe struct {}
// func (s Stripe) pay(amount float32){
// 	fmt.Println("Payment using Stripe Payment Gateway")
// }

// func main() {

// 	stripePaymentGW := Stripe{}

// 	newPayment := Payment{
// 		gateway: stripePaymentGW,
// 	}
// 	newPayment.makePayment(100)

// }

/*
!In above code , If one wants to switch from Stripe to Razor pay , he has to do
!changes in the Payment struct which Violates Open-Close Principle. Even for testing if
!one need to make "fakePaymentGW" , then it would require to change the Payment struct
!and the makePayment function.
*/

/*---------------------------------------------------------------------------------*/

//Part2
/*
Mental Model :- Interfaces are contracts written by the caller, not the callee

An interface is a collection of method signatures (name, args, returns) that defines a behavior contract
*/

type PaymentGateWay interface { //implemented a interface with same method signatures
	pay(amount float32)
}

type Payment struct {
	gateway PaymentGateWay //Changed the type of gateway from Stripe/RazorPay/PayPal
	// to an Interface (PaymentGateWay)
}

func (p Payment) makePayment(amount float32) {
	p.gateway.pay(amount)
}

//RazorPay
type RazorPay struct{}
func (r RazorPay) pay(amount float32) {
	fmt.Println("Payment using Razor Payment Gateway", amount)

}

//Stripe
type Stripe struct{}
func (s Stripe) pay(amount float32) {
	fmt.Println("Payment using Stripe Payment Gateway",amount)
}

//Fake Payment Gateway
type Fake struct{}
func (f Fake) pay(amount float32){
	fmt.Println("Payment using Fake Payment Gateway",amount)
}

func main() {

	stripePaymentGW := Stripe{}
	razorPaymentGW := RazorPay{}

	//Payment using Stripe
	payment1 := Payment{
		gateway: stripePaymentGW,
	}
	payment1.makePayment(100)

	//Payment using RazorPay
	payment2 := Payment{
		gateway: razorPaymentGW,
	}
	payment2.makePayment(200)

	
	
	//Fake Payment for testing 

	fakePaymentGW := Fake{}

	payment3 := Payment{
		gateway: fakePaymentGW,
	}
	payment3.makePayment(300)

}


/*
Important Go rule:

You never write implements PaymentGateWay
Go automatically checks:
“Does this type have all the methods?”
If yes → it implements the interface


Points to be noted:
-Payment struct depends ONLY on the interface.
			This means:
				Payment does not know
				-Stripe
				-RazorPay
				-Fake
			It only knows:
				“I have something that can pay()”

-The magic happens in makePayment
			At runtime:
				-p.gateway holds a real object
				-Go looks at the actual concrete type
				-Calls the correct pay()
			This is called dynamic dispatch

-What happens in main()



-Why Part 1 Violates OCP but Part 2 Doesn’t
		Part 1
		type Payment struct{
			gateway Stripe
		}
		Tightly coupled
	Every new gateway → change Payment
*/