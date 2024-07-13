package http

// // ProvidePalmClient creates a new palm.Client.
// func ProvidePalmClient(ctx context.Context) *genai.GenerativeModel {
// 	apiKey := os.Getenv("PALM_API_KEY")
// 	return palm.NewClient(ctx, apiKey)
// }
//
// // ProvideChatHandler creates a new chat handler using a palm.Client.
// func ProvideChatHandler(client *genai.GenerativeModel) *chat.Handler {
// 	return chat.NewHandler(client)
// }
//
// var chatSet = wire.NewSet(
// 	ProvidePalmClient,
// 	ProvideChatHandler,
// )
//
// // InitializeChatHandler creates a chat handler with its dependencies.
// func InitializeChatHandler(ctx context.Context) (*chat.Handler, error) {
// 	wire.Build(chatSet)
// 	return nil, nil
// }
