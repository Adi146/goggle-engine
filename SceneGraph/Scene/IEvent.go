package Scene

type IEvent interface {
	RegisterListener(listener chan IEventArgs)
	DeregisterListener(listener chan IEventArgs)
	NotifyListeners(eventArgs IEventArgs)
}

type IEventArgs interface {
}

type Event struct {
	Listeners []chan IEventArgs
}

func (event *Event) RegisterListener(listener chan IEventArgs) {
	for _, existingListener := range event.Listeners {
		if listener == existingListener {
			return
		}
	}

	event.Listeners = append(event.Listeners, listener)
}

func (event *Event) DeregisterListener(listener chan IEventArgs) {
	for i, existingListener := range event.Listeners {
		if listener == existingListener {
			event.Listeners[i] = event.Listeners[len(event.Listeners)-1]
			event.Listeners = event.Listeners[:len(event.Listeners)-1]
			return
		}
	}
}

func (event *Event) NotifyListeners(eventArgs IEventArgs) {
	for _, listener := range event.Listeners {
		select {
		case listener <- eventArgs:
		default:
		}
	}
}
