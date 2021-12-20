package main

type WashingMachine interface {
	wash()
	dry()
}

type dryer struct {
}

func (d dryer) wash() {
	print("洗一洗")
}

type haier struct {
	dryer
}

func (h haier) dry() {
	print("甩一甩")
}
