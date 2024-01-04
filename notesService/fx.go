package notesService

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Invoke(RegisterNotesformEndPoint),
	fx.Provide(NewGetAuthService),
	fx.Provide(NewGetNotesService),
)
