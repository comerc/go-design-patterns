package main

import "fmt"

// Цепочка обязанностей - передаёт запрос по цепочке потенциальных обработчиков,
// пока один из них не обработает запрос.

// Избавляет от жёсткой привязки отправителя запроса к его получателю,
// позволяя выстраивать цепь из различных обработчиков динамически.

type Executer interface {
	execute(*Patient)
}

type Department struct {
	next Executer
}

func (d *Department) setNext(next Executer) {
	d.next = next
}

type Reception struct {
	Department
}

func (r *Reception) execute(p *Patient) {
	if p.registrationDone {
		fmt.Println("Patient registration already done")
		r.next.execute(p)
		return
	}
	fmt.Println("Reception registering patient")
	p.registrationDone = true
	r.next.execute(p)
}

type Doctor struct {
	Department
}

func (d *Doctor) execute(p *Patient) {
	if p.doctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
		d.next.execute(p)
		return
	}
	fmt.Println("Doctor checking patient")
	p.doctorCheckUpDone = true
	d.next.execute(p)
}

type Medical struct {
	Department
}

func (m *Medical) execute(p *Patient) {
	if p.medicineDone {
		fmt.Println("Medicine already given to patient")
		m.next.execute(p)
		return
	}
	fmt.Println("Medical giving medicine to patient")
	p.medicineDone = true
	m.next.execute(p)
}

type Cashier struct {
	Department
}

func (c *Cashier) execute(p *Patient) {
	if p.paymentDone {
		fmt.Println("Payment Done")
	}
	fmt.Println("Cashier getting money from patient patient")
}

type Patient struct {
	name              string
	registrationDone  bool
	doctorCheckUpDone bool
	medicineDone      bool
	paymentDone       bool
}

func main() {

	cashier := &Cashier{}

	//Set next for medical department
	medical := &Medical{}
	medical.setNext(cashier)

	//Set next for doctor department
	doctor := &Doctor{}
	doctor.setNext(medical)

	//Set next for reception department
	reception := &Reception{}
	reception.setNext(doctor)

	patient := &Patient{name: "abc"}
	//Patient visiting
	reception.execute(patient)
}
