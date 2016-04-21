package gam

func ActorOf(props Properties) ActorRef {
	return spawnChild(props, nil)
}

func spawnChild(props Properties, parent ActorRef) ActorRef {
	cell := NewActorCell(props, parent)
	mailbox := props.ProduceMailbox()
	mailbox.RegisterHandlers(cell.invokeUserMessage, cell.invokeSystemMessage)
	ref := NewLocalActorRef(mailbox)
	cell.self = ref
	cell.invokeUserMessage(Started{})
	return ref
}