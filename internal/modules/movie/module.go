package movie

import (
	movieDomain "github.com/pawel2973/go-academy/internal/modules/movie/domain"
	"go.uber.org/fx"

	movieRepo "github.com/pawel2973/go-academy/internal/modules/movie/repository"
	movieSvc "github.com/pawel2973/go-academy/internal/modules/movie/service"
	movieHTTP "github.com/pawel2973/go-academy/internal/modules/movie/transport/http"
)

// Module bundles all dependencies for the movie bounded context.
var Module = fx.Options(
	fx.Provide(
		// inform FX: *MovieRepo implements domain.MovieRepository
		fx.Annotate(
			movieRepo.NewMovieRepo,
			fx.As(new(movieDomain.MovieRepository)),
		),
		movieSvc.NewMovieService,
		movieHTTP.NewMovieHandler,
	),
)
