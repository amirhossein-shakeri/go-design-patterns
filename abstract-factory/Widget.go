package abstractfactory

import "fmt"

// In this implementation, we define two types of widgets:
// WindowsWidget and LinuxWidget. Both of these widgets implement the
// Widget interface, which has two methods: CreateButton and
// CreateCheckbox, each of which returns an Item.

// We also define four types of items: Button, Checkbox,
// WindowsButton, and WindowsCheckbox. Both WindowsButton and
// WindowsCheckbox embed Button and Checkbox, respectively, and
// override their GetName methods to add the Windows prefix.
// Similarly, LinuxButton and LinuxCheckbox embed Button and Checkbox,
// respectively, and override their GetName methods to add the Linux
// prefix.

// In createWidget, we accept a Widget parameter to create button and
// checkbox items. We then use the CreateButton and CreateCheckbox
// methods on the given widget to create and print the names of a
// button and a checkbox.

// In main, we create instances of WindowsWidget and LinuxWidget and
// then pass these to createWidget to create the corresponding items.

// This code demonstrates how the abstract factory pattern can be used
// to create families of related items without the client code needing
// to know about their concrete classes. Instead, the choice of which
// type of item to create is delegated to the widget, which can be
// swapped out as needed.

type Item interface {
	GetName() string
}

type Widget interface {
	CreateButton() Item
	CreateCheckbox() Item
}

type WindowsWidget struct{}

func (*WindowsWidget) CreateButton() Item {
	return &WindowsButton{}
}

func (*WindowsWidget) CreateCheckbox() Item {
	return &WindowsCheckbox{}
}

type LinuxWidget struct{}

func (*LinuxWidget) CreateButton() Item {
	return &LinuxButton{}
}

func (*LinuxWidget) CreateCheckbox() Item {
	return &LinuxCheckbox{}
}

type Button struct{}

func (*Button) GetName() string {
	return "Button"
}

type Checkbox struct{}

func (*Checkbox) GetName() string {
	return "Checkbox"
}

type WindowsButton struct {
	Button
}

type WindowsCheckbox struct {
	Checkbox
}

type LinuxButton struct {
	Button
}

type LinuxCheckbox struct {
	Checkbox
}

func CreateWidget(w Widget) {
	fmt.Printf("Button: %s\n", w.CreateButton().GetName())
	fmt.Printf("Checkbox: %s\n", w.CreateCheckbox().GetName())
}

func Run() {
	windowsWidget := &WindowsWidget{}
	linuxWidget := &LinuxWidget{}

	CreateWidget(windowsWidget)
	CreateWidget(linuxWidget)
}
