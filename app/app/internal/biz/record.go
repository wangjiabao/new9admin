package biz

import (
	"context"
	v1 "dhb/app/app/api"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"strconv"
	"strings"
	"time"
)

type EthUserRecord struct {
	ID        int64
	UserId    int64
	Hash      string
	Status    string
	Type      string
	Amount    string
	RelAmount int64
	CoinType  string
	Last      int64
}

type Location struct {
	ID            int64
	UserId        int64
	Status        string
	CurrentLevel  int64
	Current       int64
	CurrentMax    int64
	CurrentMaxNew int64
	Row           int64
	Col           int64
	Count         int64
	Total         int64
	TotalTwo      int64
	TotalThree    int64
	LastLevel     int64
	StopDate      time.Time
	CreatedAt     time.Time
}

type LocationNew struct {
	ID                int64
	UserId            int64
	Num               int64
	Status            string
	Current           int64
	CurrentMax        int64
	CurrentMaxNew     int64
	StopLocationAgain int64
	OutRate           int64
	Count             int64
	StopCoin          int64
	Top               int64
	Usdt              int64
	Total             int64
	TotalTwo          int64
	TotalThree        int64
	Biw               int64
	TopNum            int64
	LastLevel         int64
	StopDate          time.Time
	CreatedAt         time.Time
}

type GlobalLock struct {
	ID     int64
	Status int64
}

type RecordUseCase struct {
	ethUserRecordRepo             EthUserRecordRepo
	userRecommendRepo             UserRecommendRepo
	configRepo                    ConfigRepo
	locationRepo                  LocationRepo
	userBalanceRepo               UserBalanceRepo
	userInfoRepo                  UserInfoRepo
	userCurrentMonthRecommendRepo UserCurrentMonthRecommendRepo
	tx                            Transaction
	log                           *log.Helper
}

type EthUserRecordRepo interface {
	GetEthUserRecordListByHash(ctx context.Context, hash ...string) (map[string]*EthUserRecord, error)
	GetEthUserRecordListByHash2(ctx context.Context, hash ...string) (map[string]*EthUserRecord, error)
	GetEthUserRecordLast(ctx context.Context) (int64, error)
	GetEthUserRecordLast2(ctx context.Context) (int64, error)
	CreateEthUserRecordListByHash(ctx context.Context, r *EthUserRecord) (*EthUserRecord, error)
	CreateEthUserRecordListByHash2(ctx context.Context, r *EthUserRecord) (*EthUserRecord, error)
}

type LocationRepo interface {
	CreateLocation(ctx context.Context, rel *Location) (*Location, error)
	GetLocationLast(ctx context.Context) (*LocationNew, error)
	GetMyLocationLast(ctx context.Context, userId int64) (*LocationNew, error)
	GetLocationById(ctx context.Context, id int64) (*LocationNew, error)
	GetMyLocationLastRunning(ctx context.Context, userId int64) (*LocationNew, error)
	GetLocationDailyYesterday(ctx context.Context, day int) ([]*LocationNew, error)
	GetBuyYesterday(ctx context.Context, day int) ([]*Reward, error)
	GetMyStopLocationLast(ctx context.Context, userId int64) (*Location, error)
	GetMyLocationRunningLast(ctx context.Context, userId int64) (*Location, error)
	GetLocationsByUserId(ctx context.Context, userId int64) ([]*Location, error)
	GetRewardLocationByRowOrCol(ctx context.Context, row int64, col int64, locationRowConfig int64) ([]*Location, error)
	GetRewardLocationByIds(ctx context.Context, ids ...int64) (map[int64]*Location, error)
	UpdateLocation(ctx context.Context, id int64, status string, current int64, stopDate time.Time) error
	UpdateLocationLastLevel(ctx context.Context, id int64, lastLevel int64) error
	GetLocations(ctx context.Context, b *Pagination, userId int64, status string) ([]*LocationNew, error, int64)
	GetLocations2(ctx context.Context, b *Pagination, userId int64) ([]*LocationNew, error, int64)
	GetUserBalanceRecords(ctx context.Context, b *Pagination, userId int64, coinType string) ([]*UserBalanceRecord, error, int64)
	GetLocationsAll(ctx context.Context, b *Pagination, userId int64) ([]*LocationNew, error, int64)
	UpdateLocationRowAndCol(ctx context.Context, id int64) error
	GetLocationsStopNotUpdate(ctx context.Context) ([]*Location, error)
	LockGlobalLocation(ctx context.Context) (bool, error)
	UnLockGlobalLocation(ctx context.Context) (bool, error)
	LockGlobalWithdraw(ctx context.Context) (bool, error)
	UnLockGlobalWithdraw(ctx context.Context) (bool, error)
	GetLockGlobalLocation(ctx context.Context) (*GlobalLock, error)
	GetLocationUserCount(ctx context.Context) int64
	GetLocationByIds(ctx context.Context, userIds ...int64) ([]*LocationNew, error)
	GetAllLocations(ctx context.Context) ([]*Location, error)
	GetAllLocationsNew(ctx context.Context, currentMax int64) ([]*LocationNew, error)
	GetAllLocationsNew2(ctx context.Context) ([]*LocationNew, error)
	GetLocationsByTop(ctx context.Context, top int64) ([]*LocationNew, error)
	GetLocationsByUserIds(ctx context.Context, userIds []int64) ([]*Location, error)

	CreateLocation2New(ctx context.Context, rel *LocationNew, amount int64) (*LocationNew, error)
	CreateLocationNew(ctx context.Context, rel *LocationNew, amount int64) (*LocationNew, error)
	GetMyStopLocationsLast(ctx context.Context, userId int64) ([]*LocationNew, error)
	GetMyStopLocations2Last(ctx context.Context, userId int64) ([]*LocationNew, error)
	GetLocationsNewByUserId(ctx context.Context, userId int64) ([]*LocationNew, error)
	GetLocationsNew2ByUserId(ctx context.Context, userId int64) ([]*LocationNew, error)
	UpdateLocationNew(ctx context.Context, id int64, status string, current int64, stopDate time.Time) error
	UpdateLocationNewNew(ctx context.Context, id int64, status string, current int64, amountB int64, biw int64, stopDate time.Time) error
	UpdateLocationNewNewNew(ctx context.Context, id int64, current int64) error
	UpdateLocationNew2(ctx context.Context, id int64, status string, current int64, stopDate time.Time) error
	UpdateLocationNewCurrent(ctx context.Context, id int64, current int64) error
	GetRunningLocations(ctx context.Context) ([]*LocationNew, error)
	GetLocationLastByNum(ctx context.Context) (*LocationNew, error)
	GetLocationsByNum(ctx context.Context, start int64, end int64) ([]*LocationNew, error)
	UpdateLocationNew7(ctx context.Context, id int64, status string, num int64, currentMax int64, stopDate time.Time) error
	UpdateLocationNewCount(ctx context.Context, id int64, count int64, total int64) error
	UpdateLocationNewTotal(ctx context.Context, id int64, count int64, total int64) error
	UpdateLocationNewTotalSub(ctx context.Context, id int64, count int64, total int64) error
	GetLocationFirst(ctx context.Context) (*LocationNew, error)
}

func NewRecordUseCase(
	ethUserRecordRepo EthUserRecordRepo,
	locationRepo LocationRepo,
	userBalanceRepo UserBalanceRepo,
	userRecommendRepo UserRecommendRepo,
	userInfoRepo UserInfoRepo,
	configRepo ConfigRepo,
	userCurrentMonthRecommendRepo UserCurrentMonthRecommendRepo,
	tx Transaction,
	logger log.Logger) *RecordUseCase {
	return &RecordUseCase{
		ethUserRecordRepo:             ethUserRecordRepo,
		locationRepo:                  locationRepo,
		configRepo:                    configRepo,
		userRecommendRepo:             userRecommendRepo,
		userBalanceRepo:               userBalanceRepo,
		userCurrentMonthRecommendRepo: userCurrentMonthRecommendRepo,
		userInfoRepo:                  userInfoRepo,
		tx:                            tx,
		log:                           log.NewHelper(logger),
	}
}

func (ruc *RecordUseCase) GetEthUserRecordByTxHash(ctx context.Context, txHash ...string) (map[string]*EthUserRecord, error) {
	return ruc.ethUserRecordRepo.GetEthUserRecordListByHash(ctx, txHash...)
}

func (ruc *RecordUseCase) GetEthUserRecordByTxHash2(ctx context.Context, txHash ...string) (map[string]*EthUserRecord, error) {
	return ruc.ethUserRecordRepo.GetEthUserRecordListByHash2(ctx, txHash...)
}

func (ruc *RecordUseCase) GetEthUserRecordLast(ctx context.Context) (int64, error) {
	return ruc.ethUserRecordRepo.GetEthUserRecordLast(ctx)
}

func (ruc *RecordUseCase) GetEthUserRecordLast2(ctx context.Context) (int64, error) {
	return ruc.ethUserRecordRepo.GetEthUserRecordLast2(ctx)
}

func (ruc *RecordUseCase) GetGlobalLock(ctx context.Context) (*GlobalLock, error) {
	return ruc.locationRepo.GetLockGlobalLocation(ctx)
}

func (ruc *RecordUseCase) Deposit(ctx context.Context, userId int64, amount uint64, originTotal uint64) error {
	// 更新user last,

	// 推荐人
	var (
		err error
	)
	if err = ruc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		err = ruc.userInfoRepo.UpdateUser(ctx, userId, amount, originTotal)
		if nil != err {
			return err
		}

		// 充值记录
		err = ruc.userBalanceRepo.InRecordNew(ctx, userId, int64(amount))
		if nil != err {
			return err
		}

		return nil
	}); nil != err {
		fmt.Println(err, "错误投资3", userId, amount)
		return err
	}

	return nil
}

func (ruc *RecordUseCase) EthUserRecordHandle(ctx context.Context, ethUserRecord ...*EthUserRecord) (bool, error) {

	var (
		configs      []*Config
		buyOne       int64
		buyTwo       int64
		buyThree     int64
		buyFour      int64
		buyFive      int64
		buySix       int64
		areaOne      int64
		areaTwo      int64
		areaThree    int64
		areaFour     int64
		areaFive     int64
		recommendOne int64
		recommendTwo int64
		bPrice       int64
		bPriceBase   int64
		feeRate      int64
		//recommendRate1 int64
		//recommendRate2 int64
		//recommendRate3 int64
		//recommendRate4 int64
		//recommendRate5 int64
		//recommendRate6 int64
		//recommendRate7 int64
		//recommendRate8 int64
		//recommendBase  = int64(100)
	)
	// 配置
	configs, _ = ruc.configRepo.GetConfigByKeys(ctx,
		"area_one", "area_two", "area_three", "area_four", "area_five", "recommend_new_one", "recommend_new_two", "exchange_rate",
		"buy_one", "buy_two", "buy_six", "buy_three", "buy_four", "buy_five", "b_price", "b_price_base", "recommend_rate_1", "recommend_rate_2", "recommend_rate_3", "recommend_rate_4", "recommend_rate_5", "recommend_rate_6", "recommend_rate_7", "recommend_rate_8")
	if nil != configs {
		for _, vConfig := range configs {
			if "buy_one" == vConfig.KeyName {
				buyOne, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
			if "buy_two" == vConfig.KeyName {
				buyTwo, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
			if "buy_three" == vConfig.KeyName {
				buyThree, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
			if "buy_four" == vConfig.KeyName {
				buyFour, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
			if "buy_five" == vConfig.KeyName {
				buyFive, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}

			if "buy_six" == vConfig.KeyName {
				buySix, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}

			if "area_one" == vConfig.KeyName {
				areaOne, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
			if "area_two" == vConfig.KeyName {
				areaTwo, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
			if "area_three" == vConfig.KeyName {
				areaThree, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
			if "area_four" == vConfig.KeyName {
				areaFour, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
			if "area_five" == vConfig.KeyName {
				areaFive, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
			if "recommend_new_one" == vConfig.KeyName {
				recommendOne, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
			if "recommend_new_two" == vConfig.KeyName {
				recommendTwo, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
			if "b_price" == vConfig.KeyName {
				bPrice, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
			if "b_price_base" == vConfig.KeyName {
				bPriceBase, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
			if "exchange_rate" == vConfig.KeyName {
				feeRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}

			//		if "b_price_base" == vConfig.KeyName {
			//			bPriceBase, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			//		}
			//		if "recommend_rate_1" == vConfig.KeyName {
			//			recommendRate1, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			//		}
			//		if "recommend_rate_2" == vConfig.KeyName {
			//			recommendRate2, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			//		}
			//		if "recommend_rate_3" == vConfig.KeyName {
			//			recommendRate3, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			//		}
			//		if "recommend_rate_4" == vConfig.KeyName {
			//			recommendRate4, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			//		}
			//		if "recommend_rate_5" == vConfig.KeyName {
			//			recommendRate5, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			//		}
			//		if "recommend_rate_6" == vConfig.KeyName {
			//			recommendRate6, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			//		}
			//		if "recommend_rate_7" == vConfig.KeyName {
			//			recommendRate7, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			//		}
			//		if "recommend_rate_8" == vConfig.KeyName {
			//			recommendRate8, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			//		}
		}
	}

	for _, v := range ethUserRecord {
		var (
			allLocations   []*LocationNew
			myLocations    []*LocationNew
			myLastLocation *LocationNew
			lastLevel      int64
			err            error
		)

		// 获取当前用户的占位信息，已经有运行中的跳过
		allLocations, err = ruc.locationRepo.GetAllLocationsNew(ctx, v.RelAmount*25/10) // 同倍率的
		if nil == allLocations {                                                        // 查询异常跳过本次循环
			fmt.Println(err, "错误投资", v)
			continue
		}
		if nil != err {
			fmt.Println(err, "123")
			continue
		}

		if 10000000 == v.RelAmount && int64(len(allLocations)) < buySix {

		} else if 30000000 == v.RelAmount && int64(len(allLocations)) < buyOne {

		} else if 100000000 == v.RelAmount && int64(len(allLocations)) < buyTwo {

		} else if 300000000 == v.RelAmount && int64(len(allLocations)) < buyThree {

		} else if 500000000 == v.RelAmount && int64(len(allLocations)) < buyFour {

		} else if 1000000000 == v.RelAmount && int64(len(allLocations)) < buyFive {

		} else {
			fmt.Println(v, "1234")
			continue
		}

		// 获取当前用户的占位信息，已经有运行中的跳过
		myLocations, err = ruc.locationRepo.GetLocationsNewByUserId(ctx, v.UserId)
		if nil == myLocations { // 查询异常跳过本次循环
			fmt.Println(err, "错误投资2", v)
			continue
		}
		if nil != err {
			fmt.Println(err, "12")
			continue
		}

		if 0 < len(myLocations) {
			var (
				stop bool
			)

			for _, vMyLocations := range myLocations {
				if "stop" != vMyLocations.Status {
					stop = true
					fmt.Println(err, "已投资", v)
					break
				}

				myLastLocation = vMyLocations // 遍历到最后一个
				var tmpLastLevel int64
				// 1大区
				if vMyLocations.Total >= vMyLocations.TotalTwo && vMyLocations.Total >= vMyLocations.TotalThree {
					if areaOne <= vMyLocations.TotalTwo+vMyLocations.TotalThree {
						tmpLastLevel = 1
					}
					if areaTwo <= vMyLocations.TotalTwo+vMyLocations.TotalThree {
						tmpLastLevel = 2
					}
					if areaThree <= vMyLocations.TotalTwo+vMyLocations.TotalThree {
						tmpLastLevel = 3
					}
					if areaFour <= vMyLocations.TotalTwo+vMyLocations.TotalThree {
						tmpLastLevel = 4
					}
					if areaFive <= vMyLocations.TotalTwo+vMyLocations.TotalThree {
						tmpLastLevel = 5
					}
				} else if vMyLocations.TotalTwo >= vMyLocations.Total && vMyLocations.TotalTwo >= vMyLocations.TotalThree {
					if areaOne <= vMyLocations.Total+vMyLocations.TotalThree {
						tmpLastLevel = 1
					}
					if areaTwo <= vMyLocations.Total+vMyLocations.TotalThree {
						tmpLastLevel = 2
					}
					if areaThree <= vMyLocations.Total+vMyLocations.TotalThree {
						tmpLastLevel = 3
					}
					if areaFour <= vMyLocations.Total+vMyLocations.TotalThree {
						tmpLastLevel = 4
					}
					if areaFive <= vMyLocations.Total+vMyLocations.TotalThree {
						tmpLastLevel = 5
					}
				} else if vMyLocations.TotalThree >= vMyLocations.Total && vMyLocations.TotalThree >= vMyLocations.TotalTwo {
					if areaOne <= vMyLocations.TotalTwo+vMyLocations.Total {
						tmpLastLevel = 1
					}
					if areaTwo <= vMyLocations.TotalTwo+vMyLocations.Total {
						tmpLastLevel = 2
					}
					if areaThree <= vMyLocations.TotalTwo+vMyLocations.Total {
						tmpLastLevel = 3
					}
					if areaFour <= vMyLocations.TotalTwo+vMyLocations.Total {
						tmpLastLevel = 4
					}
					if areaFive <= vMyLocations.TotalTwo+vMyLocations.Total {
						tmpLastLevel = 5
					}
				}

				if tmpLastLevel > lastLevel {
					lastLevel = tmpLastLevel
				}

				if vMyLocations.LastLevel > lastLevel {
					lastLevel = vMyLocations.LastLevel
				}
			}

			if stop {
				continue // 跳过已经投资
			}
		}

		// 推荐人
		var (
			userRecommend         *UserRecommend
			myUserRecommendUserId int64
			tmpRecommendUserIds   []string
		)
		userRecommend, err = ruc.userRecommendRepo.GetUserRecommendByUserId(ctx, v.UserId)
		if nil != err {
			continue
		}
		if "" != userRecommend.RecommendCode {
			tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")
			if 2 <= len(tmpRecommendUserIds) {
				myUserRecommendUserId, _ = strconv.ParseInt(tmpRecommendUserIds[len(tmpRecommendUserIds)-1], 10, 64) // 最后一位是直推人
			}
		}

		// 直推人投资
		var (
			myRecommmendLocation *LocationNew
		)
		if 0 < myUserRecommendUserId {
			myRecommmendLocation, err = ruc.locationRepo.GetMyLocationLastRunning(ctx, myUserRecommendUserId)
			if nil != err {
				continue
			}
		}

		// 顺位
		var (
			lastLocation *LocationNew
		)

		// 有直推人占位且第一次入金，挂在直推人名下，按位查找
		if nil != myRecommmendLocation && nil == myLastLocation {
			var (
				selectLocation *LocationNew // 选中的
			)
			if 3 <= myRecommmendLocation.Count {
				var tmpSopFor bool

				tmpIds := make([]int64, 0)
				tmpIds = append(tmpIds, myRecommmendLocation.ID)
				for _, vTmpId := range tmpIds { // 小于3个人
					// 查找
					var (
						topLocations []*LocationNew
					)
					topLocations, err = ruc.locationRepo.GetLocationsByTop(ctx, vTmpId)
					if nil != err {
						break
					}

					// 没数据没数据, 正常最少三个
					if 0 >= len(topLocations) {
						tmpSopFor = true
						break
					}

					for _, vTopLocations := range topLocations {
						if 3 > vTopLocations.Count {
							selectLocation = vTopLocations
							break
						}
						tmpIds = append(tmpIds, vTopLocations.ID)
					}

					if nil != selectLocation {
						break
					}
					//
				}

				if tmpSopFor {
					continue
				}

			} else {
				selectLocation = myRecommmendLocation
			}

			lastLocation = selectLocation
		} else if nil != myLastLocation || nil == myRecommmendLocation { // 2复投，直接顺位 2直推无位置或一号用户无直推人，顺位补齐

			var (
				firstLocation  *LocationNew
				tmpSopFor      bool
				selectLocation *LocationNew // 选中的
			)

			firstLocation, err = ruc.locationRepo.GetLocationFirst(ctx)
			if nil != err {
				continue
			}
			if nil != firstLocation {
				if 3 <= firstLocation.Count {
					tmpIds := make([]int64, 0)
					tmpIds = append(tmpIds, firstLocation.ID)
					for _, vTmpId := range tmpIds { // 小于3个人
						// 查找
						var (
							topLocations []*LocationNew
						)
						topLocations, err = ruc.locationRepo.GetLocationsByTop(ctx, vTmpId)
						if nil != err {
							break
						}

						// 没数据, 正常最少三个
						if 0 >= len(topLocations) {
							tmpSopFor = true
							break
						}

						for _, vTopLocations := range topLocations {
							if 3 > vTopLocations.Count {
								selectLocation = vTopLocations
								break
							}
							tmpIds = append(tmpIds, vTopLocations.ID)
						}

						if nil != selectLocation {
							break
						}
						//
					}

					if tmpSopFor {
						continue
					}
				} else {
					selectLocation = firstLocation
				}
			}

			lastLocation = selectLocation
		} else {
			continue
		}

		if err = ruc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务

			// 推荐人
			if 0 < len(tmpRecommendUserIds) {
				lastKey := len(tmpRecommendUserIds) - 1 // 有直推len比>=2 ,key是0则是空格，1是直推，键位最后一个人
				if 1 <= lastKey {
					for i := 0; i <= 1; i++ { // 两代
						// 有占位信息，推荐人推荐人的上一代
						if lastKey-i <= 0 {
							break
						}

						tmpMyTopUserRecommendUserId, _ := strconv.ParseInt(tmpRecommendUserIds[lastKey-i], 10, 64) // 最后一位是直推人
						if 0 >= tmpMyTopUserRecommendUserId {
							break
						}

						var myUserRecommendUserLocationsLast []*LocationNew
						myUserRecommendUserLocationsLast, err = ruc.locationRepo.GetLocationsNewByUserId(ctx, tmpMyTopUserRecommendUserId)
						if nil != myUserRecommendUserLocationsLast {

							var tmpMyTopUserRecommendUserLocationLast *LocationNew
							if 1 <= len(myUserRecommendUserLocationsLast) {
								for _, vMyUserRecommendUserLocationLast := range myUserRecommendUserLocationsLast {
									if "running" == vMyUserRecommendUserLocationLast.Status {
										tmpMyTopUserRecommendUserLocationLast = vMyUserRecommendUserLocationLast
										break
									}
								}

								if nil == tmpMyTopUserRecommendUserLocationLast { // 无位
									continue
								}

								tmpMinUsdt := tmpMyTopUserRecommendUserLocationLast.Usdt
								if v.RelAmount < tmpMinUsdt {
									tmpMinUsdt = v.RelAmount
								}

								var tmpMyRecommendAmount int64
								if 0 == i { // 当前用户被此人直推
									tmpMyRecommendAmount = tmpMinUsdt / 1000 * recommendOne
								} else if 1 == i {
									tmpMyRecommendAmount = tmpMinUsdt / 1000 * recommendTwo
								} else {
									continue
								}

								if 0 < tmpMyRecommendAmount { // 扣除推荐人分红
									bAmount := tmpMyRecommendAmount * bPriceBase / bPrice
									tmpStatus := tmpMyTopUserRecommendUserLocationLast.Status
									tmpStopDate := time.Now().UTC().Add(8 * time.Hour)
									// 过了
									if tmpMyTopUserRecommendUserLocationLast.Current+tmpMyRecommendAmount >= tmpMyTopUserRecommendUserLocationLast.CurrentMax { // 占位分红人分满停止
										tmpStatus = "stop"
										tmpStopDate = time.Now().UTC().Add(8 * time.Hour)

										tmpMyRecommendAmount = tmpMyTopUserRecommendUserLocationLast.CurrentMax - tmpMyTopUserRecommendUserLocationLast.Current
										bAmount = tmpMyRecommendAmount * bPriceBase / bPrice
									}

									if 0 < tmpMyRecommendAmount && 0 < bAmount {
										var tmpMaxNew int64
										if tmpMyTopUserRecommendUserLocationLast.CurrentMaxNew < tmpMyTopUserRecommendUserLocationLast.CurrentMax {
											tmpMaxNew = tmpMyTopUserRecommendUserLocationLast.CurrentMax - tmpMyTopUserRecommendUserLocationLast.CurrentMaxNew
										}

										if err = ruc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
											err = ruc.locationRepo.UpdateLocationNewNew(ctx, tmpMyTopUserRecommendUserLocationLast.ID, tmpStatus, tmpMyRecommendAmount, tmpMaxNew, bAmount, tmpStopDate) // 分红占位数据修改
											if nil != err {
												return err
											}

											_, err = ruc.userBalanceRepo.RecommendLocationRewardBiw(ctx, tmpMyTopUserRecommendUserId, bAmount, int64(i+1), tmpStatus, tmpMaxNew, feeRate) // 推荐人奖励
											if nil != err {
												return err
											}

											// 业绩减掉
											if "stop" == tmpStatus {
												tmpTop := tmpMyTopUserRecommendUserLocationLast.Top
												tmpTopNum := tmpMyTopUserRecommendUserLocationLast.TopNum
												for j := 0; j < 10000 && 0 < tmpTop && 0 < tmpTopNum; j++ {
													err = ruc.locationRepo.UpdateLocationNewTotalSub(ctx, tmpTop, tmpTopNum, tmpMyTopUserRecommendUserLocationLast.Usdt/100000)
													if nil != err {
														return err
													}

													var (
														currentLocation *LocationNew
													)
													currentLocation, err = ruc.locationRepo.GetLocationById(ctx, tmpTop)
													if nil != err {
														return err
													}

													if nil != currentLocation && 0 < currentLocation.Top {
														tmpTop = currentLocation.Top
														tmpTopNum = currentLocation.TopNum
														continue
													}

													break
												}
											}

											return nil
										}); nil != err {
											fmt.Println("err reward daily recommend", err, myUserRecommendUserLocationsLast)
											continue
										}
									}
								}
							}
						}

					}

				}

			}

			var (
				tmpTop int64
				tmpNum int64
			)

			// 顺位
			if nil != lastLocation {
				err = ruc.locationRepo.UpdateLocationNewCount(ctx, lastLocation.ID, lastLocation.Count+1, v.RelAmount/100000)
				if nil != err {
					return err
				}
				tmpTop = lastLocation.ID
				tmpNum = lastLocation.Count + 1

				var (
					currentTop    = lastLocation.Top
					currentTopNum = lastLocation.TopNum
				)
				// 大小区业绩
				for j := 0; j < 10000 && 0 < currentTop && 0 < currentTopNum; j++ {
					err = ruc.locationRepo.UpdateLocationNewTotal(ctx, currentTop, currentTopNum, v.RelAmount/100000)
					if nil != err {
						return err
					}

					var (
						currentLocation *LocationNew
					)
					currentLocation, err = ruc.locationRepo.GetLocationById(ctx, currentTop)
					if nil != err {
						return err
					}

					if nil != currentLocation && 0 < currentLocation.Top && 0 < currentLocation.TopNum {
						currentTop = currentLocation.Top
						currentTopNum = currentLocation.TopNum
						continue
					}

					break
				}
			}

			_, err = ruc.locationRepo.CreateLocationNew(ctx, &LocationNew{ // 占位
				UserId:     v.UserId,
				Status:     "running",
				Current:    0,
				CurrentMax: v.RelAmount * 25 / 10, // 2.5倍率
				Num:        1,
				Top:        tmpTop,
				TopNum:     tmpNum,
				LastLevel:  lastLevel,
			}, v.RelAmount)

			if nil != err {
				return err
			}

			// 充值记录
			_, err = ruc.ethUserRecordRepo.CreateEthUserRecordListByHash(ctx, &EthUserRecord{
				Hash:     v.Hash,
				UserId:   v.UserId,
				Status:   v.Status,
				Type:     v.Type,
				Amount:   v.Amount,
				CoinType: v.CoinType,
				Last:     v.Last,
			})

			if nil != err {
				return err
			}

			if 0 < myUserRecommendUserId {
				err = ruc.userRecommendRepo.UpdateUserRecommendTotal(ctx, myUserRecommendUserId, v.RelAmount/100000)
				if nil != err {
					return err
				}
			}

			//if 0 < len(tmpRecommendUserIdsInt) {
			//	var (
			//		currentK int64
			//	)
			//	for k, vTmpRecommendUserIdsInt := range tmpRecommendUserIdsInt {
			//		var (
			//			myUserRecommendUserLocationsLast []*LocationNew
			//		)
			//
			//		// 有占位信息，推荐人第一代
			//		myUserRecommendUserLocationsLast, err = ruc.locationRepo.GetLocationsNewByUserId(ctx, myUserRecommendUserInfo.UserId)
			//		if nil != myUserRecommendUserLocationsLast {
			//			var myUserRecommendUserLocationLast *LocationNew
			//			if 1 <= len(myUserRecommendUserLocationsLast) {
			//				myUserRecommendUserLocationLast = myUserRecommendUserLocationsLast[0]
			//				for _, vMyUserRecommendUserLocationLast := range myUserRecommendUserLocationsLast {
			//					if "running" == vMyUserRecommendUserLocationLast.Status {
			//						myUserRecommendUserLocationLast = vMyUserRecommendUserLocationLast
			//						break
			//					}
			//				}
			//
			//				tmpStatus := myUserRecommendUserLocationLast.Status // 现在还在运行中
			//				var tmpRecommendAmount int64
			//				if 0 == currentK { // 直推
			//					if 1 <= len(myUserRecommendUserLocationsLast) {
			//						tmpRecommendAmount = v.RelAmount * recommendRate1 / recommendBase
			//					}
			//				} else if 1 == currentK {
			//					if 1 <= len(myUserRecommendUserLocationsLast) {
			//						tmpRecommendAmount = v.RelAmount * recommendRate2 / recommendBase
			//					}
			//				} else if 2 == currentK {
			//					if 2 <= len(myUserRecommendUserLocationsLast) {
			//						tmpRecommendAmount = v.RelAmount * recommendRate3 / recommendBase
			//					}
			//				} else if 3 == currentK {
			//					if 3 <= len(myUserRecommendUserLocationsLast) { // 复投2次拿第4代
			//						tmpRecommendAmount = v.RelAmount * recommendRate4 / recommendBase
			//					}
			//				} else if 4 == currentK {
			//					if 4 <= len(myUserRecommendUserLocationsLast) { // 复投3次拿第5代
			//						tmpRecommendAmount = v.RelAmount * recommendRate5 / recommendBase
			//					}
			//				} else if 5 == currentK {
			//					if 5 <= len(myUserRecommendUserLocationsLast) { // 复投4次拿第6代
			//						tmpRecommendAmount = v.RelAmount * recommendRate6 / recommendBase
			//					}
			//				} else if 6 == currentK {
			//					if 6 <= len(myUserRecommendUserLocationsLast) { // 复投5次拿第7代
			//						tmpRecommendAmount = v.RelAmount * recommendRate7 / recommendBase
			//					}
			//				} else if 7 == currentK {
			//					if 7 <= len(myUserRecommendUserLocationsLast) { // 复投6次拿第8代
			//						tmpRecommendAmount = v.RelAmount * recommendRate8 / recommendBase
			//					}
			//				}
			//
			//				// 奖励usdt
			//				tmpRewardAmount := tmpRecommendAmount * bPrice / bPriceBase
			//
			//				myUserRecommendUserLocationLast.Status = "running"
			//				myUserRecommendUserLocationLast.Current += tmpRewardAmount
			//
			//				tmpRewardAmount2 := tmpRewardAmount
			//				if myUserRecommendUserLocationLast.Current >= myUserRecommendUserLocationLast.CurrentMax { // 占位分红人分满停止
			//					myUserRecommendUserLocationLast.Status = "stop"
			//					if "running" == tmpStatus {
			//						myUserRecommendUserLocationLast.StopDate = time.Now().UTC().Add(8 * time.Hour)
			//						tmpRewardAmount2 = tmpRewardAmount - (myUserRecommendUserLocationLast.Current - myUserRecommendUserLocationLast.CurrentMax)
			//					} else {
			//						tmpRewardAmount2 = 0
			//					}
			//				}
			//
			//				if 0 < tmpRewardAmount {
			//					err = ruc.locationRepo.UpdateLocationNew(ctx, myUserRecommendUserLocationLast.ID, myUserRecommendUserLocationLast.Status, tmpRewardAmount, myUserRecommendUserLocationLast.StopDate) // 分红占位数据修改
			//					if nil != err {
			//						return err
			//					}
			//					_, err = ruc.userBalanceRepo.NormalRecommendReward(ctx, myUserRecommendUserId, tmpRewardAmount, tmpRewardAmount2, currentLocationNew.ID, tmpStatus, myUserRecommendUserLocationLast.Status, "recommend", "recommend") // 直推人奖励
			//					if nil != err {
			//						return err
			//					}
			//				}
			//			}
			//
			//		}
			//	}
			//}

			return nil
		}); nil != err {
			fmt.Println(err, "错误投资3", v)
			continue
		}
	}

	return true, nil
}

func (ruc *RecordUseCase) EthUserRecordHandle5(ctx context.Context, ethUserRecord ...*EthUserRecord) (bool, error) {

	var (
		configs       []*Config
		timeAgain     int64
		recommendRate int64
	)
	// 配置
	configs, _ = ruc.configRepo.GetConfigByKeys(ctx, "time_again", "num", "recommend_rate_2")

	if nil != configs {
		for _, vConfig := range configs {
			if "recommend_rate_2" == vConfig.KeyName {
				recommendRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "time_again" == vConfig.KeyName {
				timeAgain, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
			//else if "num" == vConfig.KeyName {
			//	num, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			//}
		}
	}

	for _, v := range ethUserRecord {
		var (
			locationCurrent         int64
			locationCurrentMax      int64
			currentLocationNew      *LocationNew
			userRecommend           *UserRecommend
			myUserRecommendUserId   int64
			myUserRecommendUserInfo *UserInfo
			myLocations             []*LocationNew
			locationNum             int64
			tmpRecommendUserIds     []string
			err                     error
		)

		// 获取当前用户的占位信息，已经有运行中的跳过
		myLocations, err = ruc.locationRepo.GetLocationsNew2ByUserId(ctx, v.UserId)
		if nil == myLocations { // 查询异常跳过本次循环
			continue
		}
		if 0 < len(myLocations) {
			tmpStatusRunning := false
			for _, vMyLocations := range myLocations {
				locationNum = vMyLocations.Num
				if "running" == vMyLocations.Status {
					tmpStatusRunning = true
					break
				}
			}

			if tmpStatusRunning { // 有运行中直接跳过本次循环
				continue
			}
		}

		locationCurrentMax = v.RelAmount * 25 / 10

		// 推荐人
		userRecommend, err = ruc.userRecommendRepo.GetUserRecommendByUserId(ctx, v.UserId)
		if nil != err {
			continue
		}

		if "" != userRecommend.RecommendCode {
			tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")
			if 2 <= len(tmpRecommendUserIds) {
				myUserRecommendUserId, _ = strconv.ParseInt(tmpRecommendUserIds[len(tmpRecommendUserIds)-1], 10, 64) // 最后一位是直推人
			}
		}

		if 0 < myUserRecommendUserId {
			myUserRecommendUserInfo, err = ruc.userInfoRepo.GetUserInfoByUserId(ctx, myUserRecommendUserId)
		}

		// 冻结
		var (
			myLastStopLocations []*LocationNew
		)
		myLastStopLocations, err = ruc.locationRepo.GetMyStopLocations2Last(ctx, v.UserId)
		now := time.Now().UTC().Add(8 * time.Hour)
		if nil != myLastStopLocations {
			for _, vMyLastStopLocations := range myLastStopLocations {
				if now.Before(vMyLastStopLocations.StopDate.Add(time.Duration(timeAgain) * time.Minute)) {
					locationCurrent += vMyLastStopLocations.Current - vMyLastStopLocations.CurrentMax // 补上
				}
			}
		}

		// 修改用户推荐人区数据，修改自身区数据
		myVip := int64(1)
		if 30000000 == v.RelAmount {
			myVip = 2
		} else if 50000000 == v.RelAmount {
			myVip = 3
		}

		if err = ruc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			tmpLocationStatus := "running"
			var tmpStopDate time.Time
			if locationCurrent >= locationCurrentMax {
				tmpLocationStatus = "stop"
				tmpStopDate = time.Now().UTC().Add(8 * time.Hour)
			}
			currentLocationNew, err = ruc.locationRepo.CreateLocation2New(ctx, &LocationNew{ // 占位
				UserId:     v.UserId,
				Status:     tmpLocationStatus,
				Current:    locationCurrent,
				CurrentMax: locationCurrentMax,
				StopDate:   tmpStopDate,
				Num:        locationNum,
			}, v.RelAmount)
			if nil != err {
				return err
			}

			_, err = ruc.userInfoRepo.UpdateUserInfoVip(ctx, v.UserId, myVip)
			if nil != err {
				return err
			}

			// 推荐人
			if nil != myUserRecommendUserInfo {
				var (
					myUserRecommendUserLocationsLast []*LocationNew
				)
				// 有占位信息，推荐人第一代
				myUserRecommendUserLocationsLast, err = ruc.locationRepo.GetLocationsNew2ByUserId(ctx, myUserRecommendUserInfo.UserId)
				if nil != myUserRecommendUserLocationsLast {
					var myUserRecommendUserLocationLast *LocationNew
					if 1 <= len(myUserRecommendUserLocationsLast) {
						for _, vMyUserRecommendUserLocationLast := range myUserRecommendUserLocationsLast {
							if "running" == vMyUserRecommendUserLocationLast.Status {
								myUserRecommendUserLocationLast = vMyUserRecommendUserLocationLast
								break
							}
						}

						// 奖励usdt
						tmpRewardAmount := v.RelAmount * recommendRate / 100
						if 0 < tmpRewardAmount && nil != myUserRecommendUserLocationLast {
							_, err = ruc.userBalanceRepo.NormalRecommendReward2(ctx, myUserRecommendUserId, tmpRewardAmount, currentLocationNew.ID, "recommend_token", "recommend") // 直推人奖励
							if nil != err {
								return err
							}
						}
					}

				}
			}

			// 清算冻结
			if nil != myLastStopLocations {
				err = ruc.userBalanceRepo.UpdateLocationAgain2(ctx, myLastStopLocations) // 充值
				if nil != err {
					return err
				}

				if 0 < locationCurrent {
					var tmpCurrentAmount int64
					if locationCurrent > locationCurrentMax {
						tmpCurrentAmount = locationCurrentMax
					} else {
						tmpCurrentAmount = locationCurrent
					}

					_, err = ruc.userBalanceRepo.DepositLastNew2(ctx, v.UserId, tmpCurrentAmount) // 充值
					if nil != err {
						return err
					}
				}
			}

			_, err = ruc.ethUserRecordRepo.CreateEthUserRecordListByHash2(ctx, &EthUserRecord{
				Hash:     v.Hash,
				UserId:   v.UserId,
				Status:   v.Status,
				Type:     v.Type,
				Amount:   v.Amount,
				CoinType: v.CoinType,
				Last:     v.Last,
			})
			if nil != err {
				return err
			}

			return nil
		}); nil != err {
			continue
		}
	}

	return true, nil
}

func (ruc *RecordUseCase) EthUserRecordHandle2(ctx context.Context, ethUserRecord ...*EthUserRecord) (bool, error) {

	var (
	//configs       []*Confi
	//level1Price   int64
	//level2Price   int64
	//level3Price   int64
	//level4Price   int64
	//csdPrice      int64
	//vip1          bool
	//recommendRate int64
	)
	// 配置
	//configs, _ = ruc.configRepo.GetConfigByKeys(ctx,
	//	"term", "level_1_price", "level_2_price", "level_3_price", "level_4_price",
	//	"csd_price", "recommend_rate",
	//)
	//if nil != configs {
	//	for _, vConfig := range configs {
	//		if "term" == vConfig.KeyName {
	//			term, _ = strconv.ParseInt(vConfig.Value, 10, 64)
	//		} else if "level_1_price" == vConfig.KeyName {
	//			level1Price, _ = strconv.ParseInt(vConfig.Value, 10, 64)
	//		} else if "level_2_price" == vConfig.KeyName {
	//			level2Price, _ = strconv.ParseInt(vConfig.Value, 10, 64)
	//		} else if "level_3_price" == vConfig.KeyName {
	//			level3Price, _ = strconv.ParseInt(vConfig.Value, 10, 64)
	//		} else if "level_4_price" == vConfig.KeyName {
	//			level4Price, _ = strconv.ParseInt(vConfig.Value, 10, 64)
	//		} else if "csd_price" == vConfig.KeyName {
	//			csdPrice, _ = strconv.ParseInt(vConfig.Value, 10, 64)
	//		} else if "recommend_rate" == vConfig.KeyName {
	//			recommendRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
	//		}
	//	}
	//}

	for _, v := range ethUserRecord {
		var (
			err error
		)

		// 获取当前用户的占位信息，已经有运行中的跳过
		//myLocations, err = ruc.locationRepo.GetLocationsNewByUserId(ctx, v.UserId)
		//if nil == myLocations { // 查询异常跳过本次循环
		//	continue
		//}
		//if 0 < len(myLocations) {
		//	tmpStatusRunning := false
		//	for _, vMyLocations := range myLocations {
		//		if term == vMyLocations.Term {
		//			tmpStatusRunning = true
		//			break
		//		}
		//	}
		//
		//	if tmpStatusRunning { // 有运行中直接跳过本次循环
		//		continue
		//	}
		//}

		//if v.RelAmount >= level1Price*100000 && v.RelAmount < level2Price*100000 {
		//	locationCurrentMax = level1Price * 100000 * csdPrice / 1000
		//} else if v.RelAmount >= level2Price*100000 && v.RelAmount < level3Price*100000 {
		//	locationCurrentMax = level2Price * 100000 * csdPrice / 1000
		//} else if v.RelAmount >= level3Price*100000 && v.RelAmount < level4Price*100000 {
		//	locationCurrentMax = level3Price * 100000 * csdPrice / 1000
		//} else if v.RelAmount >= level4Price*100000 {
		//	locationCurrentMax = level4Price * 100000 * csdPrice / 1000
		//	vip1 = true
		//}

		// 推荐人
		//userRecommend, err = ruc.userRecommendRepo.GetUserRecommendByUserId(ctx, v.UserId)
		//if nil != err {
		//	continue
		//}
		//if "" != userRecommend.RecommendCode {
		//	tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")
		//	if 2 <= len(tmpRecommendUserIds) {
		//		myUserRecommendUserId, _ = strconv.ParseInt(tmpRecommendUserIds[len(tmpRecommendUserIds)-1], 10, 64) // 最后一位是直推人
		//	}
		//}
		//if 0 < myUserRecommendUserId {
		//	myUserRecommendUserInfo, err = ruc.userInfoRepo.GetUserInfoByUserId(ctx, myUserRecommendUserId)
		//}

		// 冻结
		//myLastStopLocations, err = ruc.locationRepo.GetMyStopLocationsLast(ctx, v.UserId)
		//now := time.Now().UTC().Add(8 * time.Hour)
		//if nil != myLastStopLocations {
		//	for _, vMyLastStopLocations := range myLastStopLocations {
		//		if now.Before(vMyLastStopLocations.StopDate.Add(time.Duration(timeAgain) * time.Minute)) {
		//			locationCurrent += vMyLastStopLocations.Current - vMyLastStopLocations.CurrentMax // 补上
		//		}
		//	}
		//}
		//myUserInfo, err = ruc.userInfoRepo.GetUserInfoByUserId(ctx, v.UserId)
		//if nil != err {
		//	continue
		//}
		//
		if err = ruc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			//	tmpLocationStatus := "running"
			//	var tmpStopDate time.Time
			//	//if locationCurrent >= locationCurrentMax {
			//	//	tmpLocationStatus = "stop"
			//	//	tmpStopDate = time.Now().UTC().Add(8 * time.Hour)
			//	//}
			//	currentLocationNew, err = ruc.locationRepo.CreateLocationNew(ctx, &LocationNew{ // 占位
			//		UserId:     v.UserId,
			//		Term:       term,
			//		Status:     tmpLocationStatus,
			//		Current:    locationCurrent,
			//		CurrentMax: locationCurrentMax,
			//		StopDate:   tmpStopDate,
			//	})
			//	if nil != err {
			//		return err
			//	}
			//
			//	if vip1 && 0 == myUserInfo.Vip {
			//		_, err = ruc.userInfoRepo.UpdateUserInfoVip(ctx, v.UserId, 1) // 推荐人信息修改
			//		if nil != err {
			//			return err
			//		}
			//	}
			//
			//	// 推荐人
			//	if nil != myUserRecommendUserInfo {
			//		if 0 == len(myLocations) { // vip 等级调整，被推荐人首次入单
			//			myUserRecommendUserInfo.HistoryRecommend += 1
			//			_, err = ruc.userInfoRepo.UpdateUserInfo(ctx, myUserRecommendUserInfo) // 推荐人信息修改
			//			if nil != err {
			//				return err
			//			}
			//		}
			//
			//		// 有占位信息，推荐人第一代
			//		//	myUserRecommendUserLocationsLast, err = ruc.locationRepo.GetLocationsNewByUserId(ctx, myUserRecommendUserInfo.UserId)
			//		//	if nil != myUserRecommendUserLocationsLast {
			//		//		var myUserRecommendUserLocationLast *LocationNew
			//		//		if 1 <= len(myUserRecommendUserLocationsLast) {
			//		//			myUserRecommendUserLocationLast = myUserRecommendUserLocationsLast[0]
			//		//			for _, vMyUserRecommendUserLocationLast := range myUserRecommendUserLocationsLast {
			//		//				if "running" == vMyUserRecommendUserLocationLast.Status {
			//		//					myUserRecommendUserLocationLast = vMyUserRecommendUserLocationLast
			//		//					break
			//		//				}
			//		//			}
			//		//
			//		//			tmpStatus := myUserRecommendUserLocationLast.Status // 现在还在运行中
			//		//
			//		//			// 奖励usdt
			//		//			tmpRewardAmount := currentValue * recommendNeed / 100
			//		//
			//		//			tmpBalanceAmount := tmpRewardAmount * rewardRate / 100 // 记录下一次
			//		//			tmpBalanceCoinAmount := tmpRewardAmount * coinRewardRate / 100 * 1000 / coinPrice
			//		//
			//		//			myUserRecommendUserLocationLast.Status = "running"
			//		//			myUserRecommendUserLocationLast.Current += tmpRewardAmount
			//		//
			//		//			if myUserRecommendUserLocationLast.Current >= myUserRecommendUserLocationLast.CurrentMax { // 占位分红人分满停止
			//		//				myUserRecommendUserLocationLast.Status = "stop"
			//		//				if "running" == tmpStatus {
			//		//					myUserRecommendUserLocationLast.StopDate = time.Now().UTC().Add(8 * time.Hour)
			//		//					// 这里刚刚停止
			//		//					tmpLastAmount := tmpRewardAmount - (myUserRecommendUserLocationLast.Current - myUserRecommendUserLocationLast.CurrentMax)
			//		//					tmpBalanceAmount = tmpLastAmount * rewardRate / 100 // 记录下一次
			//		//					tmpBalanceCoinAmount = tmpLastAmount * coinRewardRate / 100 * 1000 / coinPrice
			//		//				}
			//		//			}
			//		//
			//		//			if 0 < tmpRewardAmount {
			//		//				err = ruc.locationRepo.UpdateLocationNew(ctx, myUserRecommendUserLocationLast.ID, myUserRecommendUserLocationLast.Status, tmpRewardAmount, myUserRecommendUserLocationLast.StopDate) // 分红占位数据修改
			//		//				if nil != err {
			//		//					return err
			//		//				}
			//		//
			//		//				_, err = ruc.userBalanceRepo.NormalRecommendReward(ctx, myUserRecommendUserId, tmpRewardAmount, tmpBalanceAmount, tmpBalanceCoinAmount, currentLocationNew.ID, tmpStatus) // 直推人奖励
			//		//				if nil != err {
			//		//					return err
			//		//				}
			//		//
			//		//			}
			//		//		}
			//		//
			//		//	}
			//
			//		_, err = ruc.userBalanceRepo.NewNormalRecommendReward(ctx, myUserRecommendUserId, locationCurrentMax*recommendRate/1000, currentLocationNew.ID) // 直推人奖励
			//		if nil != err {
			//			return err
			//		}
			//	}

			// 修改用户推荐人区数据，修改自身区数据
			//_, err = ruc.userRecommendRepo.UpdateUserAreaSelfAmount(ctx, v.UserId, locationCurrentMax/100000)
			//if nil != err {
			//	return err
			//}
			//for _, vTmpRecommendUserIds := range tmpRecommendUserIds {
			//	vTmpRecommendUserId, _ := strconv.ParseInt(vTmpRecommendUserIds, 10, 64)
			//	if vTmpRecommendUserId > 0 {
			//		_, err = ruc.userRecommendRepo.UpdateUserAreaAmount(ctx, vTmpRecommendUserId, locationCurrentMax/100000)
			//		if nil != err {
			//			return err
			//		}
			//	}
			//}

			//_, err = ruc.userBalanceRepo.Deposit(ctx, v.UserId, locationCurrentMax, dhbAmount) // 充值
			//if nil != err {
			//	return err
			//}

			// 清算冻结
			//if nil != myLastStopLocations {
			//	err = ruc.userBalanceRepo.UpdateLocationAgain(ctx, myLastStopLocations) // 充值
			//	if nil != err {
			//		return err
			//	}
			//
			//	if 0 < locationCurrent {
			//		var tmpCurrentAmount int64
			//		if locationCurrent > locationCurrentMax {
			//			tmpCurrentAmount = locationCurrentMax
			//		} else {
			//			tmpCurrentAmount = locationCurrent
			//		}
			//
			//		stopUsdt += tmpCurrentAmount * rewardRate / 100 // 记录下一次
			//		stopCoin += tmpCurrentAmount * coinRewardRate / 100 * 1000 / coinPrice
			//
			if "CSD" == v.CoinType {
				// 推荐人
				var (
					tmpRecommendUserIds    []string
					tmpRecommendUserIdsInt []int64
				)
				var userRecommend *UserRecommend
				userRecommend, err = ruc.userRecommendRepo.GetUserRecommendByUserId(ctx, v.UserId)
				if nil == err {
					if "" != userRecommend.RecommendCode {
						tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")
						lastKey := len(tmpRecommendUserIds) - 1
						if 1 <= lastKey {
							for i := 0; i <= lastKey; i++ {
								// 有占位信息，推荐人推荐人的上一代
								if lastKey-i <= 0 {
									break
								}

								tmpMyTopUserRecommendUserId, _ := strconv.ParseInt(tmpRecommendUserIds[lastKey-i], 10, 64) // 最后一位是直推人
								tmpRecommendUserIdsInt = append(tmpRecommendUserIdsInt, tmpMyTopUserRecommendUserId)
							}
						}
					}
				}

				err = ruc.userBalanceRepo.DepositLastNewCsd(ctx, v.UserId, v.RelAmount, tmpRecommendUserIdsInt) // 充值
				if nil != err {
					return err
				}
			} else {
				err = ruc.userBalanceRepo.DepositLastNewDhb(ctx, v.UserId, v.RelAmount) // 充值
				if nil != err {
					return err
				}
			}
			//	}
			//}

			//err = ruc.userBalanceRepo.SystemReward(ctx, amount, currentLocationNew.ID)
			//if nil != err {
			//	return err
			//}

			_, err = ruc.ethUserRecordRepo.CreateEthUserRecordListByHash(ctx, &EthUserRecord{
				Hash:     v.Hash,
				UserId:   v.UserId,
				Status:   v.Status,
				Type:     v.Type,
				Amount:   v.Amount,
				CoinType: v.CoinType,
			})
			if nil != err {
				return err
			}

			return nil
		}); nil != err {
			continue
		}
	}

	return true, nil
}

func (ruc *RecordUseCase) AdminLocationInsert(ctx context.Context, userId int64, amount int64) (bool, error) {

	//var (
	//	currentLocation         *LocationNew
	//	myLastStopLocations     []*LocationNew
	//	stopCoin                int64
	//	stopUsdt                int64
	//	err                     error
	//	configs                 []*Config
	//	myLocations             []*LocationNew
	//	userRecommend           *UserRecommend
	//	tmpRecommendUserIds     []string
	//	myUserRecommendUserInfo *UserInfo
	//	myUserRecommendUserId   int64
	//	locationCurrent         int64
	//	coinPrice               int64
	//	coinRewardRate          int64
	//	rewardRate              int64
	//	outRate                 int64
	//	timeAgain               int64
	//)
	//// 配置
	//configs, _ = ruc.configRepo.GetConfigByKeys(ctx, "recommend_need", "time_again", "out_rate", "coin_price", "reward_rate", "coin_reward_rate")
	//if nil != configs {
	//	for _, vConfig := range configs {
	//		if "time_again" == vConfig.KeyName {
	//			timeAgain, _ = strconv.ParseInt(vConfig.Value, 10, 64)
	//		} else if "out_rate" == vConfig.KeyName {
	//			outRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
	//		} else if "coin_price" == vConfig.KeyName {
	//			coinPrice, _ = strconv.ParseInt(vConfig.Value, 10, 64)
	//		} else if "coin_reward_rate" == vConfig.KeyName {
	//			coinRewardRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
	//		} else if "reward_rate" == vConfig.KeyName {
	//			rewardRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
	//		}
	//	}
	//}
	//
	//// 获取当前用户的占位信息，已经有运行中的跳过
	//myLocations, err = ruc.locationRepo.GetLocationsNewByUserId(ctx, userId)
	//if nil == myLocations { // 查询异常跳过本次循环
	//	return false, errors.New(500, "ERROR", "输入金额错误，重试")
	//}
	//if 0 < len(myLocations) { // 也代表复投
	//	tmpStatusRunning := false
	//	for _, vMyLocations := range myLocations {
	//		if "running" == vMyLocations.Status {
	//			tmpStatusRunning = true
	//			break
	//		}
	//	}
	//
	//	if tmpStatusRunning { // 有运行中直接跳过本次循环
	//		return false, errors.New(500, "ERROR", "已存在运行中位置信息")
	//	}
	//}
	//
	//// 冻结
	//myLastStopLocations, err = ruc.locationRepo.GetMyStopLocationsLast(ctx, userId)
	//now := time.Now().UTC().Add(8 * time.Hour)
	//if nil != myLastStopLocations {
	//	for _, vMyLastStopLocations := range myLastStopLocations {
	//		if now.Before(vMyLastStopLocations.StopDate.Add(time.Duration(timeAgain) * time.Minute)) {
	//			locationCurrent += vMyLastStopLocations.Current - vMyLastStopLocations.CurrentMax // 补上
	//		}
	//	}
	//}
	//
	//// 推荐人
	//userRecommend, err = ruc.userRecommendRepo.GetUserRecommendByUserId(ctx, userId)
	//if nil != err {
	//	return false, errors.New(500, "ERROR", "输入金额错误，重试")
	//}
	//if "" != userRecommend.RecommendCode {
	//	tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")
	//	if 2 <= len(tmpRecommendUserIds) {
	//		myUserRecommendUserId, _ = strconv.ParseInt(tmpRecommendUserIds[len(tmpRecommendUserIds)-1], 10, 64) // 最后一位是直推人
	//	}
	//}
	//
	//if 0 < myUserRecommendUserId {
	//	myUserRecommendUserInfo, err = ruc.userInfoRepo.GetUserInfoByUserId(ctx, myUserRecommendUserId)
	//}
	//// 推荐人
	//if nil != myUserRecommendUserInfo {
	//	if 0 == len(myLocations) { // vip 等级调整，被推荐人首次入单
	//		myUserRecommendUserInfo.HistoryRecommend += 1
	//		if myUserRecommendUserInfo.HistoryRecommend >= 10 {
	//			myUserRecommendUserInfo.Vip = 5
	//		} else if myUserRecommendUserInfo.HistoryRecommend >= 8 {
	//			myUserRecommendUserInfo.Vip = 4
	//		} else if myUserRecommendUserInfo.HistoryRecommend >= 6 {
	//			myUserRecommendUserInfo.Vip = 3
	//		} else if myUserRecommendUserInfo.HistoryRecommend >= 4 {
	//			myUserRecommendUserInfo.Vip = 2
	//		} else if myUserRecommendUserInfo.HistoryRecommend >= 2 {
	//			myUserRecommendUserInfo.Vip = 1
	//		}
	//	}
	//}
	//
	//if err = ruc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
	//	tmpLocationStatus := "running"
	//	var tmpStopDate time.Time
	//	if locationCurrent >= amount*100000*outRate {
	//		tmpLocationStatus = "stop"
	//		tmpStopDate = time.Now().UTC().Add(8 * time.Hour)
	//	}
	//
	//	currentLocation, err = ruc.locationRepo.CreateLocationNew(ctx, &LocationNew{ // 占位
	//		UserId:     userId,
	//		Status:     tmpLocationStatus,
	//		Current:    locationCurrent,
	//		OutRate:    outRate,
	//		StopDate:   tmpStopDate,
	//		CurrentMax: amount * 100000 * outRate,
	//	})
	//	if nil != err {
	//		return err
	//	}
	//
	//	_, err = ruc.userInfoRepo.UpdateUserInfo(ctx, myUserRecommendUserInfo) // 推荐人信息修改
	//	if nil != err {
	//		return err
	//	}
	//
	//	_, err = ruc.userCurrentMonthRecommendRepo.CreateUserCurrentMonthRecommend(ctx, &UserCurrentMonthRecommend{ // 直推人本月推荐人数
	//		UserId:          myUserRecommendUserId,
	//		RecommendUserId: userId,
	//		Date:            time.Now().UTC().Add(8 * time.Hour),
	//	})
	//	if nil != err {
	//		return err
	//	}
	//
	//	// 清算冻结
	//	if nil != myLastStopLocations {
	//		err = ruc.userBalanceRepo.UpdateLocationAgain(ctx, myLastStopLocations) // 充值
	//		if nil != err {
	//			return err
	//		}
	//		if 0 < locationCurrent {
	//			var tmpCurrentAmount int64
	//			if locationCurrent > amount*100000*outRate {
	//				tmpCurrentAmount = amount * 100000 * outRate
	//			} else {
	//				tmpCurrentAmount = locationCurrent
	//			}
	//
	//			stopUsdt += tmpCurrentAmount * rewardRate / 100 // 记录下一次
	//			stopCoin += tmpCurrentAmount * coinRewardRate / 100 * 1000 / coinPrice
	//
	//			_, err = ruc.userBalanceRepo.DepositLastNew(ctx, userId, tmpCurrentAmount, stopUsdt, stopCoin) // 充值
	//			if nil != err {
	//				return err
	//			}
	//		}
	//	}
	//
	//	// 修改用户推荐人区数据，修改自身区数据
	//	_, err = ruc.userRecommendRepo.UpdateUserAreaSelfAmount(ctx, userId, amount*100000)
	//	if nil != err {
	//		return err
	//	}
	//	for _, vTmpRecommendUserIds := range tmpRecommendUserIds {
	//		vTmpRecommendUserId, _ := strconv.ParseInt(vTmpRecommendUserIds, 10, 64)
	//		if vTmpRecommendUserId > 0 {
	//			_, err = ruc.userRecommendRepo.UpdateUserAreaAmount(ctx, vTmpRecommendUserId, amount*100000)
	//			if nil != err {
	//				return err
	//			}
	//		}
	//	}
	//
	//	return nil
	//}); nil != err {
	//	return false, errors.New(500, "ERROR", "错误，重试")
	//
	//}

	return true, nil
}

func (ruc *RecordUseCase) LockSystem(ctx context.Context, req *v1.LockSystemRequest) (*v1.LockSystemReply, error) {
	_, _ = ruc.locationRepo.LockGlobalLocation(ctx)
	return nil, nil
}

func (ruc *RecordUseCase) UnLockEthUserRecordHandle(ctx context.Context, ethUserRecord ...*EthUserRecord) (bool, error) {
	return ruc.locationRepo.UnLockGlobalLocation(ctx)
}
