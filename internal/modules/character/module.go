package character

import (
	charDomain "github.com/pawel2973/go-academy/internal/modules/character/domain"
	"go.uber.org/fx"

	charRepo "github.com/pawel2973/go-academy/internal/modules/character/repository"
	charSvc "github.com/pawel2973/go-academy/internal/modules/character/service"
	charHTTP "github.com/pawel2973/go-academy/internal/modules/character/transport/http"
)

// Module bundles all dependencies for the character bounded context.
var Module = fx.Options(
	fx.Provide(
		// tell FX that *CharacterRepo implements domain.CharacterRepository
		fx.Annotate(
			charRepo.NewCharacterRepo,
			fx.As(new(charDomain.CharacterRepository)),
		),
		charSvc.NewCharacterService,
		charHTTP.NewCharacterHandler,
	),
)
