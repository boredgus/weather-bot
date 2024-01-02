package domain

type ConversationState string

const (
	InitialState   ConversationState = "initial"
	SetLocation    ConversationState = "set_location"
	SetTime        ConversationState = "set_time"
	UpdateSub      ConversationState = "update_subscription"
	UpdateLocation ConversationState = "update_location"
	UpdateTime     ConversationState = "update_time"
	RemoveSub      ConversationState = "remove_subscription"
)

type Conversation struct {
	State               ConversationState `json:"state" bson:"state"`
	ChatId              int64             `json:"chatId" bson:"chatId"`
	CurrentSubscription Subscription      `json:"currentSubscription" bson:"currentSubscription,omitempty"`
}

func BaseConversation(chatId int64) Conversation {
	return Conversation{
		ChatId:              chatId,
		CurrentSubscription: Subscription{},
		State:               InitialState,
	}

}

func NewConversation(chatId int64, state ConversationState, sub Subscription) Conversation {
	return Conversation{ChatId: chatId, State: state, CurrentSubscription: sub}
}

func NewConversationWithState(chatId int64, state ConversationState) Conversation {
	return Conversation{ChatId: chatId, State: state}
}

type ConversationRepository interface {
	Get(chatId int64) Conversation
	Create(chatId int64) Conversation
	Update(conv Conversation) error
	RemoveAll(chatId int64) error
}

type ConversationSvc interface {
	Get(chatId int64) Conversation
	Create(chatId int64) Conversation
	Update(conv Conversation) error
	RemoveAll(chatId int64) error
}
