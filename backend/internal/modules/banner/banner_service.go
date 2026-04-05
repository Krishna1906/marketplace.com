package banner

func UploadBannerService(imageURL, startTime, endTime string) error {
	return CreateBannerRepo(imageURL, startTime, endTime)
}

func GetAllBannersService() ([]Banner, error) {
	return GetAllBannersRepo()
}

func GetUserBannersService() ([]Banner, error) {
	return GetActiveBannersRepo()
}

func UpdateBannerServicePartial(id string, startTime, endTime *string) error {
	return UpdateBannerRepoPartial(id, startTime, endTime)
}