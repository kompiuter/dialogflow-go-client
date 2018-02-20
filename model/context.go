package model

type Context struct {
	Name       string                 `json:"name,omitempty"`
	Lifespan   int                    `json:"lifespan,omitempty"`
	Parameters map[string]interface{} `json:"parameters,omitempty"`
}

type ContextParameter struct {
	IntentAction string `json:"intent_action,omitempty"`
	Name         string `json:"name,omitempty"`
	Value        string `json:"value,omitempty"`
}
