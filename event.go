/*
Package event provides GopherJS bindings around generic JavaScript Event objects.
*/

package event

import (
	"time"
	"github.com/gopherjs/gopherjs/js"
)

type Phase uint8

const (
	NONE Phase = 0
	CAPTURING_PHASE Phase = 1
	AT_TARGET Phase = 2
	BUBBLING_PHASE Phase = 3
)

type Event interface {
	Bubbles() bool
	Cancelable() bool
	CurrentTarget() *js.Object
	DefaultPrevented() bool
	EventPhase() Phase
	IsTrusted() bool
	Target() *js.Object
	TimeStamp() uint64
	Timestamp() time.Time
	Type() string
	PreventDefault()
	StopImmediatePropagation()
	StopPropagation()
}

func (e *BasicEvent) Bubbles() bool {
	return e.Get("bubbles").Bool()
}

func (e *BasicEvent) Cancelable() bool {
	return e.Get("cancelable").Bool()
}

func (e *BasicEvent) CurrentTarget() *js.Object {
	return e.Get("currentTarget")
}

func (e *BasicEvent) DefaultPrevented() bool {
	return e.Get("defaultPrevented").Bool()
}

func (e *BasicEvent) EventPhase() Phase {
	return Phase( e.Get("eventPhase").Int() )
}

func (e *BasicEvent) IsTrusted() bool {
	return e.Get("isTrusted").Bool()
}

func (e *BasicEvent) Target() *js.Object {
	return e.Get("target")
}

// TimeStamp returns the number of milliseconds since the epoch when the event was created
func (e *BasicEvent) TimeStamp() uint64 {
	return e.Get("timeStamp").Uint64()
}

// Timestamp returns the Unix time, in seconds, since the epoch when the even twas created
func (e *BasicEvent) Timestamp() time.Time {
	ms := e.TimeStamp()
	sec := ms % 1000;
	ms = ms-sec*1000;
	return time.Unix(int64(sec), int64(ms))
}

func (e *BasicEvent) Type() string {
	return e.Get("type").String()
}

func (e *BasicEvent) PreventDefault() {
	e.Call("preventDefault")
}

func (e *BasicEvent) StopImmediatePropagation() {
	e.Call("stopImmediatePropagation")
}

func (e *BasicEvent) StopPropagation() {
	e.Call("stopPropagation")
}

type BasicEvent struct {
	js.Object
}

func New(t string) Event {
	e := js.Global.Get("Event").New(t)
	return &BasicEvent{*e}
}

func Custom(t string, data interface{}) Event {
	e := js.Global.Get("CustomEvent").New(t, data)
	return &BasicEvent{*e}
}
