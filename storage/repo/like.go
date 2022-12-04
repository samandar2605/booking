package repo

type Like struct {
	ID      int
	HotelId int
	UserId  int
	Status  bool
}

type LikesDislikesCountsResult struct {
	LikesCount    int
	DislikesCount int
}

type LikeStorageI interface {
	CreateOrUpdate(l *Like) error
	Get(userId, hotelId int) (*Like, error)
	GetLikesDislikesCount(HotelId int) (LikesDislikesCountsResult, error)
}
