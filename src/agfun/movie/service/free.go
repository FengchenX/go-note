package service

import (
	"agfun/movie/dto"
	"agfun/movie/entity"
	"github.com/kataras/iris"
	"util"
)

func (s *MovieSvc) AddFree(c iris.Context) {

}
func (s *MovieSvc) GetFrees(c iris.Context) {
	res := dto.FreeMovies{
		Total: 5,
		Rows:  []dto.FreeMovie{
			{
				Movie: dto.Movie{
					Movie: entity.Movie{
						ID:       "1",
						VideoID:  "video1",
						Name:     "video1",
						Creator:  "",
						CreateAt: "",
						Thumb:    "",
						Describe: "",
					},
				},
			},
			{
				Movie: dto.Movie{
					Movie: entity.Movie{
						ID:       "2",
						VideoID:  "video2",
						Name:     "video2",
						Creator:  "",
						CreateAt: "",
						Thumb:    "",
						Describe: "",
					},
				},
			},
			{
				Movie: dto.Movie{
					Movie: entity.Movie{
						ID:       "3",
						VideoID:  "video3",
						Name:     "video3",
						Creator:  "",
						CreateAt: "",
						Thumb:    "",
						Describe: "",
					},
				},
			},
			{
				Movie: dto.Movie{
					Movie: entity.Movie{
						ID:       "4",
						VideoID:  "video4",
						Name:     "video4",
						Creator:  "",
						CreateAt: "",
						Thumb:    "",
						Describe: "",
					},
				},
			},
			{
				Movie: dto.Movie{
					Movie: entity.Movie{
						ID:       "5",
						VideoID:  "video5",
						Name:     "video5",
						Creator:  "",
						CreateAt: "",
						Thumb:    "",
						Describe: "",
					},
				},
			},
		},
	}
	util.Success(c, &res)
}
func (s *MovieSvc) UpdateFree(c iris.Context) {

}
func (s *MovieSvc) DeleteFree(c iris.Context) {

}
