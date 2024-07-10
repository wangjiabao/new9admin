package biz

import (
	"context"
	"crypto/md5"
	v1 "dhb/app/app/api"
	"dhb/app/app/internal/pkg/middleware/auth"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwt2 "github.com/golang-jwt/jwt/v4"
	"sort"
	"strconv"
	"strings"
	"time"
)

type User struct {
	ID         int64
	Address    string
	AddressTwo string
	PrivateKey string
	Last       uint64
	Total      uint64
	CreatedAt  time.Time
}

type Admin struct {
	ID       int64
	Password string
	Account  string
	Type     string
}

type AdminAuth struct {
	ID      int64
	AdminId int64
	AuthId  int64
}

type Auth struct {
	ID   int64
	Name string
	Path string
	Url  string
}

type UserInfo struct {
	ID               int64
	UserId           int64
	Vip              int64
	HistoryRecommend int64
	LockVip          int64
	UseVip           int64
	TeamCsdBalance   int64
}

type UserRecommendArea struct {
	ID            int64
	RecommendCode string
	Num           int64
}

type UserRecommend struct {
	ID            int64
	UserId        int64
	RecommendCode string
	Total         int64
	CreatedAt     time.Time
}

type UserBalanceRecord struct {
	ID        int64
	UserId    int64
	Amount    int64
	CoinType  string
	CreatedAt time.Time
}

type BalanceReward struct {
	ID             int64
	UserId         int64
	Status         int64
	Amount         int64
	SetDate        time.Time
	LastRewardDate time.Time
	UpdatedAt      time.Time
	CreatedAt      time.Time
}

type UserCurrentMonthRecommend struct {
	ID              int64
	UserId          int64
	RecommendUserId int64
	Date            time.Time
}

type Config struct {
	ID      int64
	KeyName string
	Name    string
	Value   string
}

type UserBalance struct {
	ID          int64
	UserId      int64
	BalanceUsdt int64
	BalanceDhb  int64
}

type Withdraw struct {
	ID              int64
	UserId          int64
	Amount          int64
	RelAmount       int64
	BalanceRecordId int64
	Status          string
	Type            string
	AmountNew       float64
	AmountNewRel    float64
	CreatedAt       time.Time
}

type Trade struct {
	ID           int64
	UserId       int64
	AmountCsd    int64
	RelAmountCsd int64
	AmountHbs    int64
	CsdReward    int64
	RelAmountHbs int64
	Status       string
	CreatedAt    time.Time
}

type UserUseCase struct {
	repo                          UserRepo
	urRepo                        UserRecommendRepo
	configRepo                    ConfigRepo
	uiRepo                        UserInfoRepo
	ubRepo                        UserBalanceRepo
	locationRepo                  LocationRepo
	userCurrentMonthRecommendRepo UserCurrentMonthRecommendRepo
	tx                            Transaction
	log                           *log.Helper
}

type Reward struct {
	ID               int64
	UserId           int64
	Amount           int64
	AmountB          float64
	BalanceRecordId  int64
	Type             string
	TypeRecordId     int64
	Reason           string
	ReasonLocationId int64
	LocationType     string
	CreatedAt        time.Time
}

type Pagination struct {
	PageNum  int
	PageSize int
}

type UserArea struct {
	ID         int64
	UserId     int64
	Amount     int64
	SelfAmount int64
	Level      int64
}

type PriceChange struct {
	ID        int64
	Origin    int64
	Price     int64
	Status    int64
	CreatedAt time.Time
}

type UserSortRecommendReward struct {
	UserId int64
	Total  int64
}

type ConfigRepo interface {
	GetConfigByKeys(ctx context.Context, keys ...string) ([]*Config, error)
	GetConfigs(ctx context.Context) ([]*Config, error)
	UpdateConfig(ctx context.Context, id int64, value string) (bool, error)
	CreatePriceChangeConfig(ctx context.Context, origin int64, price int64) error
	UpdatePriceChangeStatus(ctx context.Context, id int64, status int64) (bool, error)
}

type UserBalanceRepo interface {
	GetPriceChangeConfig(ctx context.Context) (*PriceChange, error)
	CreateUserBalance(ctx context.Context, u *User) (*UserBalance, error)
	LocationReward(ctx context.Context, userId int64, amount int64, locationId int64, myLocationId int64, locationType string, status string) (int64, error)
	WithdrawReward(ctx context.Context, userId int64, amount int64, locationId int64, myLocationId int64, locationType string, status string) (int64, error)
	RecommendReward(ctx context.Context, userId int64, amount int64, locationId int64, status string) (int64, error)
	RecommendTeamReward(ctx context.Context, userId int64, rewardAmount int64, amount int64, amountDhb int64, locationId int64, recommendNum int64, status string) (int64, error)
	RecommendRewardBiw(ctx context.Context, userId int64, rewardAmount int64, recommendNum int64, stop string, tmpMaxNew int64, feeRate int64) (int64, error)
	LocationRewardBiw(ctx context.Context, userId int64, rewardAmount int64, stop string, currentMaxNew int64, feeRate int64) (int64, error)
	RecommendLocationRewardBiw(ctx context.Context, userId int64, rewardAmount int64, recommendNum int64, stop string, tmpMaxNew int64, feeRate int64) (int64, error)
	PriceChange(ctx context.Context, userId int64, rewardAmount int64, up string) error
	InRecordNew(ctx context.Context, userId int64, address string, amount int64) error
	FirstRewardBiw(ctx context.Context, userId int64, amount float64) error
	SecondRewardBiw(ctx context.Context, userId int64, amount float64, rewardType string) error
	AreaRewardBiw(ctx context.Context, userId int64, rewardAmount int64, tmpCurrentReward int64, areaType int64, stop string, tmpMaxNew int64, feeRate int64) (int64, error)
	FourRewardBiw(ctx context.Context, userId int64, rewardAmount int64, num int64) (int64, error)
	ExchangeBiw(ctx context.Context, userId int64, currentMaxNew int64, feeRate int64) (int64, error)
	SystemWithdrawReward(ctx context.Context, amount int64, locationId int64) error
	SystemReward(ctx context.Context, amount int64, locationId int64) error
	SystemDailyReward(ctx context.Context, amount int64, locationId int64) error
	GetSystemYesterdayDailyReward(ctx context.Context, day int) (*Reward, error)
	SystemFee(ctx context.Context, amount int64, locationId int64) error
	UserFee(ctx context.Context, userId int64, amount int64) (int64, error)
	UserDailyFee(ctx context.Context, userId int64, amount int64, status string) (int64, error)
	UserDailyRecommendArea(ctx context.Context, userId int64, rewardAmount int64, amount int64, amountDhb int64, status string) (int64, error)
	RecommendWithdrawReward(ctx context.Context, userId int64, amount int64, locationId int64, status string) (int64, error)
	RecommendWithdrawTopReward(ctx context.Context, userId int64, amount int64, locationId int64, vip int64, status string) (int64, error)
	NormalRecommendReward(ctx context.Context, userId int64, rewardAmount int64, rewardAmount2 int64, locationId int64, status string, status2 string, type1 string, reason string) (int64, error)
	NewNormalRecommendReward(ctx context.Context, userId int64, amount int64, locationId int64, tmpRecommendUserIdsInt []int64) (int64, error)
	NormalRecommendReward2(ctx context.Context, userId int64, rewardAmount int64, locationId int64, type1 string, reason string) (int64, error)
	NormalReward3(ctx context.Context, userId int64, rewardAmount int64, rewardAmount2 int64, locationId int64, status string, status2 string) (int64, error)
	NormalReward4(ctx context.Context, userId int64, rewardAmount int64, locationId int64) (int64, error)
	NormalRecommendTopReward(ctx context.Context, userId int64, amount int64, locationId int64, reasonId int64, status string) (int64, error)
	NormalWithdrawRecommendReward(ctx context.Context, userId int64, amount int64, locationId int64, status string) (int64, error)
	NormalWithdrawRecommendTopReward(ctx context.Context, userId int64, amount int64, locationId int64, reasonId int64, status string) (int64, error)
	Deposit(ctx context.Context, userId int64, amount int64, dhbAmount int64) (int64, error)
	DepositLast(ctx context.Context, userId int64, lastAmount int64, locationId int64) (int64, error)
	DepositDhb(ctx context.Context, userId int64, amount int64) (int64, error)
	GetUserBalance(ctx context.Context, userId int64) (*UserBalance, error)
	GetUserRewardByUserId(ctx context.Context, userId int64) ([]*Reward, error)
	GetUserRewards(ctx context.Context, b *Pagination, userId int64, reason string) ([]*Reward, error, int64)
	GetUserRewardsLastMonthFee(ctx context.Context) ([]*Reward, error)
	GetUserBalanceByUserIds(ctx context.Context, userIds ...int64) (map[int64]*UserBalance, error)
	GetUserBalanceLockByUserIds(ctx context.Context, userIds ...int64) (map[int64]*UserBalance, error)
	GetUserBalanceUsdtTotal(ctx context.Context) (float64, error)
	GetUserBalanceLockUsdtTotal(ctx context.Context) (int64, error)
	GetUserLocationNewCurrentMaxNew(ctx context.Context) (int64, error)
	GetUserLocationNewCurrentMax(ctx context.Context) (int64, error)
	GetUserLocationNewCurrent(ctx context.Context) (int64, error)
	GetUserBalanceDHBTotal(ctx context.Context) (int64, error)
	GreateWithdraw(ctx context.Context, userId int64, amount int64, coinType string) (*Withdraw, error)
	WithdrawUsdt(ctx context.Context, userId int64, amount int64) error
	WithdrawDhb(ctx context.Context, userId int64, amount int64) error
	GetWithdrawByUserId(ctx context.Context, userId int64) ([]*Withdraw, error)
	GetWithdraws(ctx context.Context, b *Pagination, userId int64, withdrawType string) ([]*Withdraw, error, int64)
	GetWithdrawPassOrRewarded(ctx context.Context) ([]*Withdraw, error)
	GetWithdrawPassOrRewardedFirst(ctx context.Context) (*Withdraw, error)
	GetTradeOk(ctx context.Context) (*Trade, error)
	GetTradeOkkCsd(ctx context.Context) (int64, error)
	GetTradeOkkHbs(ctx context.Context) (int64, error)
	UpdateWithdraw(ctx context.Context, id int64, status string) (*Withdraw, error)
	GetWithdrawById(ctx context.Context, id int64) (*Withdraw, error)
	GetWithdrawNotDeal(ctx context.Context) ([]*Withdraw, error)
	GetWithdrawByUserIds(ctx context.Context, userIds []int64) ([]*Withdraw, error)
	GetUserBalanceRecordUsdtTotal(ctx context.Context) (int64, error)
	GetUserBalanceRecordCsdTotal(ctx context.Context) (int64, error)
	GetUserBalanceRecordHbsTotal(ctx context.Context) (int64, error)
	GetUserBalanceRecordUsdtTotalToday(ctx context.Context) (int64, error)
	GetUserRewardLocationTotalToday(ctx context.Context, reason string) (int64, error)
	GetSystemWithdrawUsdtFeeTotalToday(ctx context.Context) (int64, error)
	GetUserWithdrawUsdtTotalToday(ctx context.Context) (float64, error)
	GetUserWithdrawDhbTotalToday(ctx context.Context) (int64, error)
	GetUserWithdrawUsdtTotal(ctx context.Context) (float64, error)
	GetUserWithdrawDhbTotal(ctx context.Context) (int64, error)
	GetUserWithdrawUsdtTotalByUserIds(ctx context.Context, userIds []int64) (int64, error)
	GetUserRewardUsdtTotal(ctx context.Context) (int64, error)
	GetUserRewardBalanceRewardTotal(ctx context.Context) (int64, error)
	GetBalanceRewardTotal(ctx context.Context) (int64, error)
	GetSystemRewardUsdtTotal(ctx context.Context) (int64, error)
	UpdateWithdrawAmount(ctx context.Context, id int64, status string, amount int64) (*Withdraw, error)
	GetUserRewardRecommendSort(ctx context.Context) ([]*UserSortRecommendReward, error)
	UpdateBalance(ctx context.Context, userId int64, amount int64) (bool, error)
	UpdateTrade(ctx context.Context, id int64, status string) (*Trade, error)
	GetTradeNotDeal(ctx context.Context) ([]*Trade, error)

	UpdateWithdrawPass(ctx context.Context, id int64) (*Withdraw, error)
	UserDailyBalanceReward(ctx context.Context, userId int64, rewardAmount int64, amount int64, amountDhb int64, status string) (int64, error)
	GetBalanceRewardCurrent(ctx context.Context, now time.Time) ([]*BalanceReward, error)
	GetUserTrades(ctx context.Context, b *Pagination, userId int64) ([]*Trade, error, int64)
	UserDailyLocationReward(ctx context.Context, userId int64, rewardAmount int64, amount int64, coinAmount int64, status string, locationId int64) (int64, error)
	DepositLastNew(ctx context.Context, userId int64, lastAmount int64) (int64, error)
	DepositLastNew2(ctx context.Context, userId int64, lastAmount int64) (int64, error)
	DepositLastNewDhb(ctx context.Context, userId int64, lastCoinAmount int64) error
	DepositLastNewCsd(ctx context.Context, userId int64, lastCoinAmount int64, tmpRecommendUserIdsInt []int64) error
	UpdateBalanceRewardLastRewardDate(ctx context.Context, id int64) error
	UpdateLocationAgain(ctx context.Context, locations []*LocationNew) error
	UpdateLocationAgain2(ctx context.Context, locations []*LocationNew) error
	LocationNewDailyReward(ctx context.Context, userId int64, amount int64, locationId int64) (int64, error)
	WithdrawNewRewardRecommend(ctx context.Context, userId int64, amount int64, amountB int64, locationId int64, tmpRecommendUserIdsInt []int64) (int64, error)
	WithdrawNewRewardTeamRecommend(ctx context.Context, userId int64, amount int64, amountB int64, locationId int64, tmpRecommendUserIdsInt []int64) (int64, error)
	WithdrawNewRewardSecondRecommend(ctx context.Context, userId int64, amount int64, amountB int64, locationId int64, tmpRecommendUserIdsInt []int64) (int64, error)
	WithdrawNewRewardLevelRecommend(ctx context.Context, userId int64, amount int64, amountB int64, locationId int64, tmpRecommendUserIdsInt []int64) (int64, error)
	UpdateLocationNewMax(ctx context.Context, userId int64, amount int64) (int64, error)
}

type UserRecommendRepo interface {
	GetUserRecommendByUserId(ctx context.Context, userId int64) (*UserRecommend, error)
	CreateUserRecommend(ctx context.Context, u *User, recommendUser *UserRecommend) (*UserRecommend, error)
	GetUserRecommendByCode(ctx context.Context, code string) ([]*UserRecommend, error)
	GetUserRecommendLikeCode(ctx context.Context, code string) ([]*UserRecommend, error)
	GetUserRecommends(ctx context.Context) ([]*UserRecommend, error)
	GetUserRecommendsFour(ctx context.Context) ([]*UserRecommend, error)
	CreateUserRecommendArea(ctx context.Context, recommendAreas []*UserRecommendArea) (bool, error)
	GetUserRecommendLowAreas(ctx context.Context) ([]*UserRecommendArea, error)
	UpdateUserAreaAmount(ctx context.Context, userId int64, amount int64) (bool, error)
	UpdateUserAreaSelfAmount(ctx context.Context, userId int64, amount int64) (bool, error)
	UpdateUserAreaLevel(ctx context.Context, userId int64, level int64) (bool, error)
	UpdateUserRecommendTotal(ctx context.Context, userId int64, total int64) error
	UpdateUserAreaLevelUp(ctx context.Context, userId int64, level int64) (bool, error)
	GetUserAreas(ctx context.Context, userIds []int64) ([]*UserArea, error)
	GetUserArea(ctx context.Context, userId int64) (*UserArea, error)
	CreateUserArea(ctx context.Context, u *User) (bool, error)
}

type UserCurrentMonthRecommendRepo interface {
	GetUserCurrentMonthRecommendByUserId(ctx context.Context, userId int64) ([]*UserCurrentMonthRecommend, error)
	GetUserCurrentMonthRecommendGroupByUserId(ctx context.Context, b *Pagination, userId int64) ([]*UserCurrentMonthRecommend, error, int64)
	CreateUserCurrentMonthRecommend(ctx context.Context, u *UserCurrentMonthRecommend) (*UserCurrentMonthRecommend, error)
	GetUserCurrentMonthRecommendCountByUserIds(ctx context.Context, userIds ...int64) (map[int64]int64, error)
	GetUserLastMonthRecommend(ctx context.Context) ([]int64, error)
}

type UserInfoRepo interface {
	CreateUserInfo(ctx context.Context, u *User) (*UserInfo, error)
	GetUserInfoByUserId(ctx context.Context, userId int64) (*UserInfo, error)
	UpdateUserPassword(ctx context.Context, userId int64, password string) (*User, error)
	UpdateUser(ctx context.Context, userId int64, amount uint64, originTotal uint64) error
	UpdateUserNew(ctx context.Context, userId int64, total uint64) error
	UpdateUserInfo(ctx context.Context, u *UserInfo) (*UserInfo, error)
	UpdateUserInfo2(ctx context.Context, u *UserInfo) (*UserInfo, error)
	UpdateUserInfoVip(ctx context.Context, userId, vip int64) (*UserInfo, error)
	GetUserInfoByUserIds(ctx context.Context, userIds ...int64) (map[int64]*UserInfo, error)
	GetUserInfosByVipAndLockVip(ctx context.Context) ([]*UserInfo, error)
}

type UserRepo interface {
	GetUserById(ctx context.Context, Id int64) (*User, error)
	UndoUser(ctx context.Context, userId int64, undo int64) (bool, error)
	GetAdminByAccount(ctx context.Context, account string, password string) (*Admin, error)
	GetAdminById(ctx context.Context, id int64) (*Admin, error)
	GetUserByAddresses(ctx context.Context, Addresses ...string) (map[string]*User, error)
	GetUsersNew(ctx context.Context) ([]*User, error)
	GetUserByAddress(ctx context.Context, address string) (*User, error)
	CreateUser(ctx context.Context, user *User) (*User, error)
	CreateAdmin(ctx context.Context, a *Admin) (*Admin, error)
	GetUserByUserIds(ctx context.Context, userIds ...int64) (map[int64]*User, error)
	GetAdmins(ctx context.Context) ([]*Admin, error)
	GetUsers(ctx context.Context, b *Pagination, address string, isLocation bool, vip int64) ([]*User, error, int64)
	GetAllUsers(ctx context.Context) ([]*User, error)
	GetAllUserInfos(ctx context.Context) ([]*UserInfo, error)
	GetUserCount(ctx context.Context) (int64, error)
	GetUserCountToday(ctx context.Context) (int64, error)
	CreateAdminAuth(ctx context.Context, adminId int64, authId int64) (bool, error)
	DeleteAdminAuth(ctx context.Context, adminId int64, authId int64) (bool, error)
	GetAuths(ctx context.Context) ([]*Auth, error)
	GetAuthByIds(ctx context.Context, ids ...int64) (map[int64]*Auth, error)
	GetAdminAuth(ctx context.Context, adminId int64) ([]*AdminAuth, error)
	UpdateAdminPassword(ctx context.Context, account string, password string) (*Admin, error)
}

func NewUserUseCase(repo UserRepo, tx Transaction, configRepo ConfigRepo, uiRepo UserInfoRepo, urRepo UserRecommendRepo, locationRepo LocationRepo, userCurrentMonthRecommendRepo UserCurrentMonthRecommendRepo, ubRepo UserBalanceRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{
		repo:                          repo,
		tx:                            tx,
		configRepo:                    configRepo,
		locationRepo:                  locationRepo,
		userCurrentMonthRecommendRepo: userCurrentMonthRecommendRepo,
		uiRepo:                        uiRepo,
		urRepo:                        urRepo,
		ubRepo:                        ubRepo,
		log:                           log.NewHelper(logger),
	}
}

func (uuc *UserUseCase) GetUserByAddress(ctx context.Context, Addresses ...string) (map[string]*User, error) {
	return uuc.repo.GetUserByAddresses(ctx, Addresses...)
}

func (uuc *UserUseCase) GetUsersNew(ctx context.Context) ([]*User, error) {
	return uuc.repo.GetUsersNew(ctx)
}

func (uuc *UserUseCase) GetDhbConfig(ctx context.Context) ([]*Config, error) {
	return uuc.configRepo.GetConfigByKeys(ctx, "level1Dhb", "level2Dhb", "level3Dhb")
}

func (uuc *UserUseCase) GetExistUserByAddressOrCreate(ctx context.Context, u *User, req *v1.EthAuthorizeRequest) (*User, error) {
	var (
		user *User
	)
	return user, nil
}

func (uuc *UserUseCase) UserInfo(ctx context.Context, user *User) (*v1.UserInfoReply, error) {
	return &v1.UserInfoReply{}, nil
}

func (uuc *UserUseCase) RewardList(ctx context.Context, req *v1.RewardListRequest, user *User) (*v1.RewardListReply, error) {
	res := &v1.RewardListReply{
		Rewards: make([]*v1.RewardListReply_List, 0),
	}

	return res, nil
}

func (uuc *UserUseCase) RecommendRewardList(ctx context.Context, user *User) (*v1.RecommendRewardListReply, error) {
	res := &v1.RecommendRewardListReply{
		Rewards: make([]*v1.RecommendRewardListReply_List, 0),
	}

	return res, nil
}

func (uuc *UserUseCase) FeeRewardList(ctx context.Context, user *User) (*v1.FeeRewardListReply, error) {
	res := &v1.FeeRewardListReply{
		Rewards: make([]*v1.FeeRewardListReply_List, 0),
	}

	return res, nil
}

func (uuc *UserUseCase) WithdrawList(ctx context.Context, user *User) (*v1.WithdrawListReply, error) {
	res := &v1.WithdrawListReply{
		Withdraw: make([]*v1.WithdrawListReply_List, 0),
	}

	return res, nil
}

func (uuc *UserUseCase) Withdraw(ctx context.Context, req *v1.WithdrawRequest, user *User) (*v1.WithdrawReply, error) {
	return &v1.WithdrawReply{
		Status: "ok",
	}, nil
}

func (uuc *UserUseCase) AdminRewardList(ctx context.Context, req *v1.AdminRewardListRequest) (*v1.AdminRewardListReply, error) {
	var (
		userSearch  *User
		userId      int64 = 0
		userRewards []*Reward
		users       map[int64]*User
		userIdsMap  map[int64]int64
		userIds     []int64
		err         error
		count       int64
	)
	res := &v1.AdminRewardListReply{
		Rewards: make([]*v1.AdminRewardListReply_List, 0),
	}

	// 地址查询
	if "" != req.Address {
		userSearch, err = uuc.repo.GetUserByAddress(ctx, req.Address)
		if nil != err {
			return res, nil
		}
		userId = userSearch.ID
	}

	userRewards, err, count = uuc.ubRepo.GetUserRewards(ctx, &Pagination{
		PageNum:  int(req.Page),
		PageSize: 10,
	}, userId, req.Reason)
	if nil != err {
		return res, nil
	}
	res.Count = count

	userIdsMap = make(map[int64]int64, 0)
	for _, vUserReward := range userRewards {
		userIdsMap[vUserReward.UserId] = vUserReward.UserId
	}
	for _, v := range userIdsMap {
		userIds = append(userIds, v)
	}

	users, err = uuc.repo.GetUserByUserIds(ctx, userIds...)
	for _, vUserReward := range userRewards {
		tmpUser := ""
		if nil != users {
			if _, ok := users[vUserReward.UserId]; ok {
				tmpUser = users[vUserReward.UserId].Address
			}
		}

		res.Rewards = append(res.Rewards, &v1.AdminRewardListReply_List{
			CreatedAt: vUserReward.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
			Amount:    vUserReward.Amount,
			AmountB:   vUserReward.AmountB,
			Address:   tmpUser,
			Reason:    vUserReward.Reason,
		})
	}

	return res, nil
}

func (uuc *UserUseCase) AdminTradeList(ctx context.Context, req *v1.AdminTradeListRequest) (*v1.AdminTradeListReply, error) {
	var (
		userSearch *User
		userId     int64 = 0
		userTrades []*Trade
		users      map[int64]*User
		userIdsMap map[int64]int64
		userIds    []int64
		err        error
		count      int64
	)
	res := &v1.AdminTradeListReply{
		Trades: make([]*v1.AdminTradeListReply_List, 0),
	}

	// 地址查询
	if "" != req.Address {
		userSearch, err = uuc.repo.GetUserByAddress(ctx, req.Address)
		if nil != err {
			return res, nil
		}
		userId = userSearch.ID
	}

	userTrades, err, count = uuc.ubRepo.GetUserTrades(ctx, &Pagination{
		PageNum:  int(req.Page),
		PageSize: 10,
	}, userId)
	if nil != err {
		return res, nil
	}
	res.Count = count

	userIdsMap = make(map[int64]int64, 0)
	for _, vUserReward := range userTrades {
		userIdsMap[vUserReward.UserId] = vUserReward.UserId
	}
	for _, v := range userIdsMap {
		userIds = append(userIds, v)
	}

	users, err = uuc.repo.GetUserByUserIds(ctx, userIds...)
	for _, vUserReward := range userTrades {
		tmpUser := ""
		if nil != users {
			if _, ok := users[vUserReward.UserId]; ok {
				tmpUser = users[vUserReward.UserId].Address
			}
		}

		res.Trades = append(res.Trades, &v1.AdminTradeListReply_List{
			CreatedAt: vUserReward.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
			AmountCsd: fmt.Sprintf("%.2f", float64(vUserReward.AmountCsd)/float64(100000)),
			AmountHbs: fmt.Sprintf("%.2f", float64(vUserReward.AmountHbs)/float64(100000)),
			Address:   tmpUser,
		})
	}

	return res, nil
}

func (uuc *UserUseCase) AdminUserList(ctx context.Context, req *v1.AdminUserListRequest) (*v1.AdminUserListReply, error) {
	var (
		users        []*User
		userIds      []int64
		userBalances map[int64]*UserBalance
		count        int64
		err          error
	)

	res := &v1.AdminUserListReply{
		Users: make([]*v1.AdminUserListReply_UserList, 0),
	}

	users, err, count = uuc.repo.GetUsers(ctx, &Pagination{
		PageNum:  int(req.Page),
		PageSize: 10,
	}, req.Address, true, req.Vip)
	if nil != err {
		return res, nil
	}
	res.Count = count

	for _, vUsers := range users {
		userIds = append(userIds, vUsers.ID)
	}

	userBalances, err = uuc.ubRepo.GetUserBalanceByUserIds(ctx, userIds...)
	if nil != err {
		return res, nil
	}

	for _, vUsers := range users {
		// 伞下业绩
		//locations, err = uuc.locationRepo.GetLocationsByUserId(ctx, vUsers.ID)
		//if nil != locations && 0 < len(locations) {
		//	for _, v := range locations {
		//
		//		if "running" == v.Status {
		//			areaAll = v.Total + v.TotalThree + v.TotalTwo
		//			if v.TotalTwo >= v.Total && v.TotalTwo >= v.TotalThree {
		//				areaMax = v.TotalTwo
		//				areaMin = v.Total + v.TotalThree
		//			}
		//			if v.Total >= v.TotalTwo && v.Total >= v.TotalThree {
		//				areaMax = v.Total
		//				areaMin = v.TotalTwo + v.TotalThree
		//			}
		//			if v.TotalThree >= v.Total && v.TotalThree >= v.TotalTwo {
		//				areaMax = v.TotalThree
		//				areaMin = v.TotalTwo + v.Total
		//			}
		//		}
		//
		//		if "stop" == v.Status {
		//			out++
		//		}
		//		var tmpLastLevel int64
		//		// 1大区
		//		if v.Total >= v.TotalTwo && v.Total >= v.TotalThree {
		//			if areaOne <= v.TotalTwo+v.TotalThree {
		//				tmpLastLevel = 1
		//			}
		//			if areaTwo <= v.TotalTwo+v.TotalThree {
		//				tmpLastLevel = 2
		//			}
		//			if areaThree <= v.TotalTwo+v.TotalThree {
		//				tmpLastLevel = 3
		//			}
		//			if areaFour <= v.TotalTwo+v.TotalThree {
		//				tmpLastLevel = 4
		//			}
		//			if areaFive <= v.TotalTwo+v.TotalThree {
		//				tmpLastLevel = 5
		//			}
		//		} else if v.TotalTwo >= v.Total && v.TotalTwo >= v.TotalThree {
		//			if areaOne <= v.Total+v.TotalThree {
		//				tmpLastLevel = 1
		//			}
		//			if areaTwo <= v.Total+v.TotalThree {
		//				tmpLastLevel = 2
		//			}
		//			if areaThree <= v.Total+v.TotalThree {
		//				tmpLastLevel = 3
		//			}
		//			if areaFour <= v.Total+v.TotalThree {
		//				tmpLastLevel = 4
		//			}
		//			if areaFive <= v.Total+v.TotalThree {
		//				tmpLastLevel = 5
		//			}
		//		} else if v.TotalThree >= v.Total && v.TotalThree >= v.TotalTwo {
		//			if areaOne <= v.TotalTwo+v.Total {
		//				tmpLastLevel = 1
		//			}
		//			if areaTwo <= v.TotalTwo+v.Total {
		//				tmpLastLevel = 2
		//			}
		//			if areaThree <= v.TotalTwo+v.Total {
		//				tmpLastLevel = 3
		//			}
		//			if areaFour <= v.TotalTwo+v.Total {
		//				tmpLastLevel = 4
		//			}
		//			if areaFive <= v.TotalTwo+v.Total {
		//				tmpLastLevel = 5
		//			}
		//		}
		//
		//		if tmpLastLevel > lastLevel {
		//			lastLevel = tmpLastLevel
		//		}
		//
		//		if v.LastLevel > lastLevel {
		//			lastLevel = v.LastLevel
		//		}
		//	}
		//}
		var (
			userRecommend      *UserRecommend
			myRecommendUsers   []*UserRecommend
			myRecommendUserIds []int64
		)
		userRecommend, err = uuc.urRepo.GetUserRecommendByUserId(ctx, vUsers.ID)
		if nil != err {
			return res, nil
		}
		myCode := userRecommend.RecommendCode + "D" + strconv.FormatInt(vUsers.ID, 10)
		myRecommendUsers, err = uuc.urRepo.GetUserRecommendByCode(ctx, myCode)
		if nil == err {
			// 找直推
			for _, vMyRecommendUsers := range myRecommendUsers {

				var (
					tmpMyRecommendLocations []*Location
				)
				tmpMyRecommendLocations, err = uuc.locationRepo.GetLocationsByUserId(ctx, vMyRecommendUsers.UserId)
				if nil != err {
					return nil, err
				}

				if 0 < len(tmpMyRecommendLocations) {
					myRecommendUserIds = append(myRecommendUserIds, vMyRecommendUsers.UserId)
				}

			}
		}

		if _, ok := userBalances[vUsers.ID]; !ok {
			continue
		}

		res.Users = append(res.Users, &v1.AdminUserListReply_UserList{
			UserId:           vUsers.ID,
			CreatedAt:        vUsers.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
			Address:          vUsers.Address,
			BalanceUsdt:      int64(vUsers.Total),
			BalanceDhb:       fmt.Sprintf("%.2f", float64(userBalances[vUsers.ID].BalanceDhb)/float64(100000)),
			Vip:              0,
			HistoryRecommend: int64(len(myRecommendUserIds)),
		})
	}

	return res, nil
}

func (uuc *UserUseCase) GetUserByUserIds(ctx context.Context, userIds ...int64) (map[int64]*User, error) {
	return uuc.repo.GetUserByUserIds(ctx, userIds...)
}

func (uuc *UserUseCase) AdminUndoUpdate(ctx context.Context, req *v1.AdminUndoUpdateRequest) (*v1.AdminUndoUpdateReply, error) {
	var (
		err  error
		undo int64
	)

	res := &v1.AdminUndoUpdateReply{}

	if 1 == req.SendBody.Undo {
		undo = 1
	} else {
		undo = 0
	}

	_, err = uuc.repo.UndoUser(ctx, req.SendBody.UserId, undo)
	if nil != err {
		return res, err
	}

	return res, nil
}

func (uuc *UserUseCase) AdminAreaLevelUpdate(ctx context.Context, req *v1.AdminAreaLevelUpdateRequest) (*v1.AdminAreaLevelUpdateReply, error) {
	var (
		err error
	)

	res := &v1.AdminAreaLevelUpdateReply{}

	_, err = uuc.urRepo.UpdateUserAreaLevel(ctx, req.SendBody.UserId, req.SendBody.Level)
	if nil != err {
		return res, err
	}

	return res, nil
}

func (uuc *UserUseCase) AdminRecordList(ctx context.Context, req *v1.RecordListRequest) (*v1.RecordListReply, error) {
	var (
		locations  []*EthUserRecord
		userSearch *User
		userId     int64
		userIds    []int64
		userIdsMap map[int64]int64
		users      map[int64]*User
		count      int64
		err        error
	)

	res := &v1.RecordListReply{
		Locations: make([]*v1.RecordListReply_LocationList, 0),
	}

	// 地址查询
	if "" != req.Address {
		userSearch, err = uuc.repo.GetUserByAddress(ctx, req.Address)
		if nil != err {
			return res, nil
		}
		userId = userSearch.ID
	}

	locations, err, count = uuc.locationRepo.GetUserBalanceRecords(ctx, &Pagination{
		PageNum:  int(req.Page),
		PageSize: 10,
	}, userId, "usdt")
	if nil != err {
		return res, nil
	}
	res.Count = count

	userIdsMap = make(map[int64]int64, 0)
	for _, vLocations := range locations {
		userIdsMap[vLocations.UserId] = vLocations.UserId
	}
	for _, v := range userIdsMap {
		userIds = append(userIds, v)
	}

	users, err = uuc.repo.GetUserByUserIds(ctx, userIds...)
	if nil != err {
		return res, nil
	}

	for _, v := range locations {
		if _, ok := users[v.UserId]; !ok {
			continue
		}

		res.Locations = append(res.Locations, &v1.RecordListReply_LocationList{
			CreatedAt: v.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
			Address:   users[v.UserId].Address,
			Amount:    v.AmountTwo,
		})
	}

	return res, nil

}

func (uuc *UserUseCase) AdminLocationList(ctx context.Context, req *v1.AdminLocationListRequest) (*v1.AdminLocationListReply, error) {
	var (
		locations  []*LocationNew
		userSearch *User
		userId     int64
		userIds    []int64
		userIdsMap map[int64]int64
		users      map[int64]*User
		count      int64
		err        error
	)

	res := &v1.AdminLocationListReply{
		Locations: make([]*v1.AdminLocationListReply_LocationList, 0),
	}

	// 地址查询
	if "" != req.Address {
		userSearch, err = uuc.repo.GetUserByAddress(ctx, req.Address)
		if nil != err {
			return res, nil
		}
		userId = userSearch.ID
	}

	status := "stop"
	if "running" == req.Status {
		status = "running"
	}
	locations, err, count = uuc.locationRepo.GetLocations(ctx, &Pagination{
		PageNum:  int(req.Page),
		PageSize: 10,
	}, userId, status)
	if nil != err {
		return res, nil
	}
	res.Count = count

	userIdsMap = make(map[int64]int64, 0)
	for _, vLocations := range locations {
		userIdsMap[vLocations.UserId] = vLocations.UserId
	}
	for _, v := range userIdsMap {
		userIds = append(userIds, v)
	}

	users, err = uuc.repo.GetUserByUserIds(ctx, userIds...)
	if nil != err {
		return res, nil
	}

	for _, v := range locations {
		if _, ok := users[v.UserId]; !ok {
			continue
		}

		var (
			myAllLocations []*Location
			out            int64
		)
		myAllLocations, err = uuc.locationRepo.GetLocationsByUserId(ctx, v.UserId)
		if nil == myAllLocations {
			return res, nil
		}

		for _, vMyAllLocations := range myAllLocations {
			if "stop" == vMyAllLocations.Status {
				out++
			}
		}

		res.Locations = append(res.Locations, &v1.AdminLocationListReply_LocationList{
			Address:       users[v.UserId].Address,
			Current:       fmt.Sprintf("%.2f", float64(v.Current)/float64(100000)),
			CurrentMax:    fmt.Sprintf("%.2f", float64(v.CurrentMax)/float64(100000)),
			Out:           out,
			CurrentMaxSub: fmt.Sprintf("%.2f", float64(v.CurrentMax-v.Current)/float64(100000)),
			Usdt:          fmt.Sprintf("%.2f", float64(v.Usdt)/float64(100000)),
		})
	}

	return res, nil

}

func (uuc *UserUseCase) AdminLocationListNew(ctx context.Context, req *v1.AdminLocationListRequest) (*v1.AdminLocationListReply, error) {
	var (
		locations  []*LocationNew
		userSearch *User
		userId     int64
		userIds    []int64
		userIdsMap map[int64]int64
		users      map[int64]*User
		count      int64
		err        error
	)

	res := &v1.AdminLocationListReply{
		Locations: make([]*v1.AdminLocationListReply_LocationList, 0),
	}

	// 地址查询
	if "" != req.Address {
		userSearch, err = uuc.repo.GetUserByAddress(ctx, req.Address)
		if nil != err {
			return res, nil
		}
		userId = userSearch.ID
	}

	locations, err, count = uuc.locationRepo.GetLocations2(ctx, &Pagination{
		PageNum:  int(req.Page),
		PageSize: 10,
	}, userId)
	if nil != err {
		return res, nil
	}
	res.Count = count

	userIdsMap = make(map[int64]int64, 0)
	for _, vLocations := range locations {
		userIdsMap[vLocations.UserId] = vLocations.UserId
	}
	for _, v := range userIdsMap {
		userIds = append(userIds, v)
	}

	users, err = uuc.repo.GetUserByUserIds(ctx, userIds...)
	if nil != err {
		return res, nil
	}

	for _, v := range locations {
		if _, ok := users[v.UserId]; !ok {
			continue
		}

		res.Locations = append(res.Locations, &v1.AdminLocationListReply_LocationList{
			Address:    users[v.UserId].Address,
			Current:    fmt.Sprintf("%.2f", float64(v.Current)/float64(100000)),
			CurrentMax: fmt.Sprintf("%.2f", float64(v.CurrentMax)/float64(100000)),
		})
	}

	return res, nil

}

func (uuc *UserUseCase) AdminLocationAllList(ctx context.Context, req *v1.AdminLocationAllListRequest) (*v1.AdminLocationAllListReply, error) {
	var (
		locations  []*LocationNew
		userSearch *User
		userId     int64
		userIds    []int64
		userIdsMap map[int64]int64
		users      map[int64]*User
		count      int64
		err        error
	)

	res := &v1.AdminLocationAllListReply{
		Locations: make([]*v1.AdminLocationAllListReply_LocationList, 0),
	}

	// 地址查询
	if "" != req.Address {
		userSearch, err = uuc.repo.GetUserByAddress(ctx, req.Address)
		if nil != err {
			return res, nil
		}
		userId = userSearch.ID
	}

	locations, err, count = uuc.locationRepo.GetLocationsAll(ctx, &Pagination{
		PageNum:  int(req.Page),
		PageSize: 10,
	}, userId)
	if nil != err {
		return res, nil
	}
	res.Count = count

	userIdsMap = make(map[int64]int64, 0)
	for _, vLocations := range locations {
		userIdsMap[vLocations.UserId] = vLocations.UserId
	}
	for _, v := range userIdsMap {
		userIds = append(userIds, v)
	}

	users, err = uuc.repo.GetUserByUserIds(ctx, userIds...)
	if nil != err {
		return res, nil
	}

	for _, v := range locations {
		if _, ok := users[v.UserId]; !ok {
			continue
		}

		res.Locations = append(res.Locations, &v1.AdminLocationAllListReply_LocationList{
			CreatedAt:  v.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
			Address:    users[v.UserId].Address,
			Status:     v.Status,
			Current:    fmt.Sprintf("%.2f", float64(v.Current)/float64(100000)),
			CurrentMax: fmt.Sprintf("%.2f", float64(v.CurrentMax)/float64(100000)),
		})
	}

	return res, nil

}

func (uuc *UserUseCase) AdminRecommendList(ctx context.Context, req *v1.AdminUserRecommendRequest) (*v1.AdminUserRecommendReply, error) {
	var (
		userRecommends []*UserRecommend
		userRecommend  *UserRecommend
		userWithdraws  []*Withdraw
		userIdsMap     map[int64]int64
		userIds        []int64
		users          map[int64]*User
		user           *User
		err            error
	)

	res := &v1.AdminUserRecommendReply{
		Users: make([]*v1.AdminUserRecommendReply_List, 0),
	}

	// 地址查询
	if 0 < req.UserId {
		userRecommend, err = uuc.urRepo.GetUserRecommendByUserId(ctx, req.UserId)
		if nil == userRecommend {
			return res, nil
		}

		userRecommends, err = uuc.urRepo.GetUserRecommendByCode(ctx, userRecommend.RecommendCode+"D"+strconv.FormatInt(userRecommend.UserId, 10))
		if nil != err {
			return res, nil
		}
	} else if "" != req.Address {
		user, err = uuc.repo.GetUserByAddress(ctx, req.Address)
		if nil != err {
			return res, nil
		}

		userRecommend, err = uuc.urRepo.GetUserRecommendByUserId(ctx, user.ID)
		if nil == userRecommend {
			return res, nil
		}

		userRecommends, err = uuc.urRepo.GetUserRecommendByCode(ctx, userRecommend.RecommendCode+"D"+strconv.FormatInt(userRecommend.UserId, 10))
		if nil != err {
			return res, nil
		}
	}

	userIdsMap = make(map[int64]int64, 0)
	for _, vLocations := range userRecommends {
		userIdsMap[vLocations.UserId] = vLocations.UserId
	}
	for _, v := range userIdsMap {
		userIds = append(userIds, v)
	}

	users, err = uuc.repo.GetUserByUserIds(ctx, userIds...)
	if nil != err {
		return res, nil
	}

	userWithdraws, err = uuc.ubRepo.GetWithdrawByUserIds(ctx, userIds)
	if nil != err {
		return res, nil
	}
	useWithdrawMapAmount := make(map[int64]int64, 0)
	useWithdrawMapRelAmount := make(map[int64]int64, 0)
	for _, vUserWithdraws := range userWithdraws {
		if _, ok := useWithdrawMapAmount[vUserWithdraws.UserId]; ok {
			useWithdrawMapAmount[vUserWithdraws.UserId] += vUserWithdraws.Amount
		} else {
			useWithdrawMapAmount[vUserWithdraws.UserId] = vUserWithdraws.Amount
		}

		if _, ok := useWithdrawMapRelAmount[vUserWithdraws.UserId]; ok {
			useWithdrawMapRelAmount[vUserWithdraws.UserId] += vUserWithdraws.RelAmount
		} else {
			useWithdrawMapRelAmount[vUserWithdraws.UserId] = vUserWithdraws.RelAmount
		}
	}

	for _, v := range userRecommends {
		if _, ok := users[v.UserId]; !ok {
			continue
		}

		var (
			myAllRecommends           []*UserRecommend
			tmpMyAllRecommendsUserIds []int64
			totalWithdraw             int64
		)

		myCode := v.RecommendCode + "D" + strconv.FormatInt(v.UserId, 10)
		myAllRecommends, err = uuc.urRepo.GetUserRecommendLikeCode(ctx, myCode)
		if 0 < len(myAllRecommends) {
			for _, vMyAllRecommends := range myAllRecommends {
				tmpMyAllRecommendsUserIds = append(tmpMyAllRecommendsUserIds, vMyAllRecommends.UserId)
			}

			if 0 < len(tmpMyAllRecommendsUserIds) {
				totalWithdraw, err = uuc.ubRepo.GetUserWithdrawUsdtTotalByUserIds(ctx, tmpMyAllRecommendsUserIds)
			}
		}

		var (
			tmpRelAmount int64
			tmpAmount    int64
		)
		if _, ok := useWithdrawMapRelAmount[v.UserId]; ok {
			tmpRelAmount = useWithdrawMapRelAmount[v.UserId]
		}

		if _, ok := useWithdrawMapAmount[v.UserId]; ok {
			tmpAmount = useWithdrawMapAmount[v.UserId]
		}

		res.Users = append(res.Users, &v1.AdminUserRecommendReply_List{
			Address:            users[v.UserId].Address,
			Id:                 v.ID,
			UserId:             v.UserId,
			CreatedAt:          v.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
			Amount:             fmt.Sprintf("%.2f", float64(tmpAmount)/float64(100000)),
			RelAmount:          fmt.Sprintf("%.2f", float64(tmpRelAmount)/float64(100000)),
			RecommendAllAmount: fmt.Sprintf("%.2f", float64(totalWithdraw)/float64(100000)),
		})
	}

	return res, nil
}

func (uuc *UserUseCase) AdminMonthRecommend(ctx context.Context, req *v1.AdminMonthRecommendRequest) (*v1.AdminMonthRecommendReply, error) {
	var (
		userCurrentMonthRecommends []*UserCurrentMonthRecommend
		searchUser                 *User
		userIdsMap                 map[int64]int64
		userIds                    []int64
		searchUserId               int64
		users                      map[int64]*User
		count                      int64
		err                        error
	)

	res := &v1.AdminMonthRecommendReply{
		Users: make([]*v1.AdminMonthRecommendReply_List, 0),
	}

	// 地址查询
	if "" != req.Address {
		searchUser, err = uuc.repo.GetUserByAddress(ctx, req.Address)
		if nil == searchUser {
			return res, nil
		}
		searchUserId = searchUser.ID
	}

	userCurrentMonthRecommends, err, count = uuc.userCurrentMonthRecommendRepo.GetUserCurrentMonthRecommendGroupByUserId(ctx, &Pagination{
		PageNum:  int(req.Page),
		PageSize: 10,
	}, searchUserId)
	if nil != err {
		return res, nil
	}
	res.Count = count

	userIdsMap = make(map[int64]int64, 0)
	for _, vRecommend := range userCurrentMonthRecommends {
		userIdsMap[vRecommend.UserId] = vRecommend.UserId
		userIdsMap[vRecommend.RecommendUserId] = vRecommend.RecommendUserId
	}
	for _, v := range userIdsMap {
		userIds = append(userIds, v)
	}

	users, err = uuc.repo.GetUserByUserIds(ctx, userIds...)
	if nil != err {
		return res, nil
	}

	for _, v := range userCurrentMonthRecommends {
		if _, ok := users[v.UserId]; !ok {
			continue
		}

		res.Users = append(res.Users, &v1.AdminMonthRecommendReply_List{
			Address:          users[v.UserId].Address,
			Id:               v.ID,
			RecommendAddress: users[v.RecommendUserId].Address,
			CreatedAt:        v.Date.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
		})
	}

	return res, nil
}

func (uuc *UserUseCase) AdminConfig(ctx context.Context, req *v1.AdminConfigRequest) (*v1.AdminConfigReply, error) {
	var (
		configs []*Config
	)

	res := &v1.AdminConfigReply{
		Config: make([]*v1.AdminConfigReply_List, 0),
	}

	configs, _ = uuc.configRepo.GetConfigs(ctx)
	if nil == configs {
		return res, nil
	}

	for _, v := range configs {
		res.Config = append(res.Config, &v1.AdminConfigReply_List{
			Id:    v.ID,
			Name:  v.Name,
			Value: v.Value,
		})
	}

	return res, nil
}

func (uuc *UserUseCase) AdminConfigUpdateListen(ctx context.Context, req *v1.AdminConfigUpdateListenRequest) (*v1.AdminConfigUpdateListenReply, error) {
	var (
		err         error
		priceChange *PriceChange
	)

	res := &v1.AdminConfigUpdateListenReply{}
	priceChange, err = uuc.ubRepo.GetPriceChangeConfig(ctx)
	if nil != err {
		return nil, err
	}

	if nil == priceChange {
		return res, nil
	}

	_, err = uuc.configRepo.UpdatePriceChangeStatus(ctx, priceChange.ID, 1)
	if nil != err {
		return nil, err
	}

	var (
		configs      []*Config
		bPrice       int64
		bPriceBase   int64
		originBprice int64
		feeRate      int64
		users        []*User
	)
	configs, _ = uuc.configRepo.GetConfigByKeys(ctx, "b_price_base", "exchange_rate")
	if nil != configs {
		for _, vConfig := range configs {
			if "b_price_base" == vConfig.KeyName {
				bPriceBase, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}

			if "exchange_rate" == vConfig.KeyName {
				feeRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
		}
	}

	bPrice = priceChange.Price
	originBprice = priceChange.Origin

	if 0 >= bPrice || 0 >= bPriceBase {
		return nil, err
	}

	users, err = uuc.repo.GetAllUsers(ctx)
	if nil != err {
		return nil, err
	}
	if nil == users {
		return nil, nil
	}
	for _, v := range users {
		var (
			runningLocation *LocationNew
			userBalance     *UserBalance
		)
		runningLocation, err = uuc.locationRepo.GetMyLocationLastRunning(ctx, v.ID)
		if nil != err {
			fmt.Println(err)
			continue
		}

		if nil == runningLocation {
			continue
		}

		userBalance, err = uuc.ubRepo.GetUserBalance(ctx, v.ID)
		if nil != err {
			fmt.Println(err)
			continue
		}

		if bPrice > originBprice {
			// 涨价
			tmp := userBalance.BalanceDhb*100/bPriceBase*bPrice - userBalance.BalanceDhb*100/bPriceBase*originBprice
			tmp = tmp / 100
			if tmp > 0 {

				if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
					runningLocation.Status = "running"
					if runningLocation.Current+tmp >= runningLocation.CurrentMax { // 占位分红人分满停止
						runningLocation.Status = "stop"
						runningLocation.StopDate = time.Now().UTC().Add(8 * time.Hour)

						tmp = runningLocation.CurrentMax - runningLocation.Current
					}

					if 0 < tmp {
						var tmpMaxNew int64
						if runningLocation.CurrentMax >= runningLocation.CurrentMaxNew {
							tmpMaxNew = runningLocation.CurrentMax - runningLocation.CurrentMaxNew
						}
						err = uuc.locationRepo.UpdateLocationNewNew(ctx, runningLocation.ID, runningLocation.Status, tmp, tmpMaxNew, 0, runningLocation.StopDate) // 分红占位数据修改
						if nil != err {
							return err
						}

						err = uuc.ubRepo.PriceChange(ctx, runningLocation.UserId, tmp, "up")
						if nil != err {
							return err
						}
					}

					// 业绩减掉
					if "stop" == runningLocation.Status {
						if runningLocation.CurrentMax >= runningLocation.CurrentMaxNew {
							_, err = uuc.ubRepo.ExchangeBiw(ctx, v.ID, runningLocation.CurrentMax-runningLocation.CurrentMaxNew, feeRate)
							if nil != err {
								return err
							}
						}

						tmpTop := runningLocation.Top
						tmpTopNum := runningLocation.TopNum
						for j := 0; j < 10000 && 0 < tmpTop && 0 < tmpTopNum; j++ {
							err = uuc.locationRepo.UpdateLocationNewTotalSub(ctx, tmpTop, tmpTopNum, runningLocation.Usdt/100000)
							if nil != err {
								return err
							}

							var (
								currentLocation *LocationNew
							)
							currentLocation, err = uuc.locationRepo.GetLocationById(ctx, tmpTop)
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
					fmt.Println("err price change", err, runningLocation)
					continue
				}
			}

		} else if bPrice < originBprice {
			// 降价
			tmp := userBalance.BalanceDhb*100/bPriceBase*originBprice - userBalance.BalanceDhb*100/bPriceBase*bPrice
			tmp = tmp / 100
			if tmp > 0 {
				if runningLocation.Current <= tmp { // 占位分红人分满停止
					tmp = runningLocation.Current
				}

				if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
					if 0 < tmp {
						err = uuc.locationRepo.UpdateLocationNewNewNew(ctx, runningLocation.ID, tmp) // 分红占位数据修改
						if nil != err {
							return err
						}

						err = uuc.ubRepo.PriceChange(ctx, runningLocation.UserId, tmp, "down")
						if nil != err {
							return err
						}
					}

					return nil
				}); nil != err {
					fmt.Println("err price change", err, runningLocation)
					continue
				}
			}
		}
	}

	return res, nil
}

func (uuc *UserUseCase) AdminConfigUpdate(ctx context.Context, req *v1.AdminConfigUpdateRequest) (*v1.AdminConfigUpdateReply, error) {
	var (
		err error
	)

	res := &v1.AdminConfigUpdateReply{}

	//configs, _ = uuc.configRepo.GetConfigByKeys(ctx, "b_price", "b_price_base", "exchange_rate")
	//if nil != configs {
	//	for _, vConfig := range configs {
	//		if "b_price" == vConfig.KeyName {
	//			originBprice, _ = strconv.ParseInt(vConfig.Value, 10, 64)
	//		}
	//		//else if "b_price_base" == vConfig.KeyName {
	//		//	bPriceBase, _ = strconv.ParseInt(vConfig.Value, 10, 64)
	//		//} else if "exchange_rate" == vConfig.KeyName {
	//		//	feeRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
	//		//}
	//	}
	//}

	if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		_, err = uuc.configRepo.UpdateConfig(ctx, req.SendBody.Id, req.SendBody.Value)
		if nil != err {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}
	//
	//_, err = uuc.configRepo.UpdateConfig(ctx, req.SendBody.Id, req.SendBody.Value)
	//if nil != err {
	//	return res, err
	//}

	return res, nil
}

func (uuc *UserUseCase) AdminWithdrawPass(ctx context.Context, req *v1.AdminWithdrawPassRequest) (*v1.AdminWithdrawPassReply, error) {
	//var (
	//	err error
	//)

	//res := &v1.AdminWithdrawPassReply{}
	//
	//_, err = uuc.ubRepo.UpdateWithdrawPass(ctx, req.SendBody.Id)
	//if nil != err {
	//	return res, err
	//}

	return nil, nil
}

func (uuc *UserUseCase) AdminPasswordUpdate(ctx context.Context, req *v1.AdminPasswordUpdateRequest) (*v1.AdminPasswordUpdateReply, error) {

	_, _ = uuc.uiRepo.UpdateUserPassword(ctx, req.SendBody.UserId, req.SendBody.Password)
	return &v1.AdminPasswordUpdateReply{}, nil
}

func (uuc *UserUseCase) AdminVipUpdate(ctx context.Context, req *v1.AdminVipUpdateRequest) (*v1.AdminVipUpdateReply, error) {
	var (
		user  *User
		err   error
		total uint64
	)

	if 1000 == req.SendBody.Vip {
		total = 1000
	} else if 3000 == req.SendBody.Vip {
		total = 3000
	} else if 5000 == req.SendBody.Vip {
		total = 5000
	} else if 15000 == req.SendBody.Vip {
		total = 15000
	} else if 30000 == req.SendBody.Vip {
		total = 30000
	} else {
		return nil, nil
	}

	user, err = uuc.repo.GetUserById(ctx, req.SendBody.UserId)
	if nil != err {
		return nil, err
	}

	// 推荐人
	if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		err = uuc.uiRepo.UpdateUserNew(ctx, req.SendBody.UserId, total)
		if nil != err {
			return err
		}

		// 充值记录
		err = uuc.ubRepo.InRecordNew(ctx, req.SendBody.UserId, user.Address, int64(total))
		if nil != err {
			return err
		}

		return nil
	}); nil != err {
		fmt.Println(err, "错误投资3", req.SendBody.UserId, total)
		return nil, err
	}

	return nil, nil
}

func (uuc *UserUseCase) AdminBalanceUpdate(ctx context.Context, req *v1.AdminBalanceUpdateRequest) (*v1.AdminBalanceUpdateReply, error) {
	var (
		err error
	)
	res := &v1.AdminBalanceUpdateReply{}

	amountFloat, _ := strconv.ParseFloat(req.SendBody.Amount, 10)
	amountFloat *= 100000
	amount, _ := strconv.ParseInt(strconv.FormatFloat(amountFloat, 'f', -1, 64), 10, 64)

	_, err = uuc.ubRepo.UpdateBalance(ctx, req.SendBody.UserId, amount) // 推荐人信息修改
	if nil != err {
		return res, err
	}

	return res, nil
}

func (uuc *UserUseCase) AdminLogin(ctx context.Context, req *v1.AdminLoginRequest, ca string) (*v1.AdminLoginReply, error) {
	var (
		admin *Admin
		err   error
	)

	res := &v1.AdminLoginReply{}
	password := fmt.Sprintf("%x", md5.Sum([]byte(req.SendBody.Password)))
	fmt.Println(password)
	admin, err = uuc.repo.GetAdminByAccount(ctx, req.SendBody.Account, password)
	if nil != err {
		return res, err
	}

	claims := auth.CustomClaims{
		UserId:   admin.ID,
		UserType: "admin",
		StandardClaims: jwt2.StandardClaims{
			NotBefore: time.Now().Unix(),              // 签名的生效时间
			ExpiresAt: time.Now().Unix() + 60*60*24*7, // 7天过期
			Issuer:    "DHB",
		},
	}
	token, err := auth.CreateToken(claims, ca)
	if err != nil {
		return nil, errors.New(500, "AUTHORIZE_ERROR", "生成token失败")
	}
	res.Token = token
	return res, nil
}

func (uuc *UserUseCase) AdminCreateAccount(ctx context.Context, req *v1.AdminCreateAccountRequest) (*v1.AdminCreateAccountReply, error) {
	var (
		admin    *Admin
		myAdmin  *Admin
		newAdmin *Admin
		err      error
	)

	res := &v1.AdminCreateAccountReply{}

	// 在上下文 context 中取出 claims 对象
	var adminId int64
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["UserId"] == nil {
			return nil, errors.New(500, "ERROR_TOKEN", "无效TOKEN")
		}
		adminId = int64(c["UserId"].(float64))
	}
	myAdmin, err = uuc.repo.GetAdminById(ctx, adminId)
	if nil == myAdmin {
		return res, err
	}
	if "super" != myAdmin.Type {
		return nil, errors.New(500, "ERROR_TOKEN", "非超管")
	}

	password := fmt.Sprintf("%x", md5.Sum([]byte(req.SendBody.Password)))
	admin, err = uuc.repo.GetAdminByAccount(ctx, req.SendBody.Account, password)
	if nil != admin {
		return nil, errors.New(500, "ERROR_TOKEN", "已存在账户")
	}

	newAdmin, err = uuc.repo.CreateAdmin(ctx, &Admin{
		Password: password,
		Account:  req.SendBody.Account,
	})

	if nil != newAdmin {
		return res, err
	}

	return res, nil
}

func (uuc *UserUseCase) AdminList(ctx context.Context, req *v1.AdminListRequest) (*v1.AdminListReply, error) {
	var (
		admins []*Admin
	)

	res := &v1.AdminListReply{Account: make([]*v1.AdminListReply_List, 0)}

	admins, _ = uuc.repo.GetAdmins(ctx)
	if nil == admins {
		return res, nil
	}

	for _, v := range admins {
		res.Account = append(res.Account, &v1.AdminListReply_List{
			Id:      v.ID,
			Account: v.Account,
		})
	}

	return res, nil
}

func (uuc *UserUseCase) AdminChangePassword(ctx context.Context, req *v1.AdminChangePasswordRequest) (*v1.AdminChangePasswordReply, error) {
	var (
		myAdmin *Admin
		admin   *Admin
		err     error
	)

	res := &v1.AdminChangePasswordReply{}

	// 在上下文 context 中取出 claims 对象
	var adminId int64
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["UserId"] == nil {
			return nil, errors.New(500, "ERROR_TOKEN", "无效TOKEN")
		}
		adminId = int64(c["UserId"].(float64))
	}
	myAdmin, err = uuc.repo.GetAdminById(ctx, adminId)
	if nil == myAdmin {
		return res, err
	}
	if "super" != myAdmin.Type {
		return nil, errors.New(500, "ERROR_TOKEN", "非超管")
	}

	password := fmt.Sprintf("%x", md5.Sum([]byte(req.SendBody.Password)))
	admin, err = uuc.repo.UpdateAdminPassword(ctx, req.SendBody.Account, password)
	if nil == admin {
		return res, err
	}

	return res, nil
}

func (uuc *UserUseCase) AuthList(ctx context.Context, req *v1.AuthListRequest) (*v1.AuthListReply, error) {
	var (
		myAdmin *Admin
		Auths   []*Auth
		err     error
	)

	res := &v1.AuthListReply{}

	// 在上下文 context 中取出 claims 对象
	var adminId int64
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["UserId"] == nil {
			return nil, errors.New(500, "ERROR_TOKEN", "无效TOKEN")
		}
		adminId = int64(c["UserId"].(float64))
	}
	myAdmin, err = uuc.repo.GetAdminById(ctx, adminId)
	if nil == myAdmin {
		return res, err
	}
	if "super" != myAdmin.Type {
		return nil, errors.New(500, "ERROR_TOKEN", "非超管")
	}

	Auths, err = uuc.repo.GetAuths(ctx)
	if nil == Auths {
		return res, err
	}

	for _, v := range Auths {
		res.Auth = append(res.Auth, &v1.AuthListReply_List{
			Id:   v.ID,
			Name: v.Name,
			Path: v.Path,
		})
	}

	return res, nil
}

func (uuc *UserUseCase) MyAuthList(ctx context.Context, req *v1.MyAuthListRequest) (*v1.MyAuthListReply, error) {
	var (
		myAdmin   *Admin
		adminAuth []*AdminAuth
		auths     map[int64]*Auth
		authIds   []int64
		err       error
	)

	res := &v1.MyAuthListReply{}

	// 在上下文 context 中取出 claims 对象
	var adminId int64
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["UserId"] == nil {
			return nil, errors.New(500, "ERROR_TOKEN", "无效TOKEN")
		}
		adminId = int64(c["UserId"].(float64))
	}
	myAdmin, err = uuc.repo.GetAdminById(ctx, adminId)
	if nil == myAdmin {
		return res, err
	}
	if "super" == myAdmin.Type {
		res.Super = int64(1)
		return res, nil
	}

	adminAuth, err = uuc.repo.GetAdminAuth(ctx, adminId)
	if nil == adminAuth {
		return res, err
	}

	for _, v := range adminAuth {
		authIds = append(authIds, v.AuthId)
	}

	if 0 >= len(authIds) {
		return res, nil
	}

	auths, err = uuc.repo.GetAuthByIds(ctx, authIds...)
	for _, v := range adminAuth {
		if _, ok := auths[v.AuthId]; !ok {
			continue
		}
		res.Auth = append(res.Auth, &v1.MyAuthListReply_List{
			Id:   v.ID,
			Name: auths[v.AuthId].Name,
			Path: auths[v.AuthId].Path,
		})
	}

	return res, nil
}

func (uuc *UserUseCase) UserAuthList(ctx context.Context, req *v1.UserAuthListRequest) (*v1.UserAuthListReply, error) {
	var (
		myAdmin   *Admin
		adminAuth []*AdminAuth
		auths     map[int64]*Auth
		authIds   []int64
		err       error
	)

	res := &v1.UserAuthListReply{}

	// 在上下文 context 中取出 claims 对象
	var adminId int64
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["UserId"] == nil {
			return nil, errors.New(500, "ERROR_TOKEN", "无效TOKEN")
		}
		adminId = int64(c["UserId"].(float64))
	}
	myAdmin, err = uuc.repo.GetAdminById(ctx, adminId)
	if nil == myAdmin {
		return res, err
	}
	if "super" != myAdmin.Type {
		return nil, errors.New(500, "ERROR_TOKEN", "非超管")
	}

	adminAuth, err = uuc.repo.GetAdminAuth(ctx, req.AdminId)
	if nil == adminAuth {
		return res, err
	}

	for _, v := range adminAuth {
		authIds = append(authIds, v.AuthId)
	}

	if 0 >= len(authIds) {
		return res, nil
	}

	auths, err = uuc.repo.GetAuthByIds(ctx, authIds...)
	for _, v := range adminAuth {
		if _, ok := auths[v.AuthId]; !ok {
			continue
		}
		res.Auth = append(res.Auth, &v1.UserAuthListReply_List{
			Id:   v.ID,
			Name: auths[v.AuthId].Name,
			Path: auths[v.AuthId].Path,
		})
	}

	return res, nil
}

func (uuc *UserUseCase) AuthAdminCreate(ctx context.Context, req *v1.AuthAdminCreateRequest) (*v1.AuthAdminCreateReply, error) {
	var (
		myAdmin *Admin
		err     error
	)

	res := &v1.AuthAdminCreateReply{}

	// 在上下文 context 中取出 claims 对象
	var adminId int64
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["UserId"] == nil {
			return nil, errors.New(500, "ERROR_TOKEN", "无效TOKEN")
		}
		adminId = int64(c["UserId"].(float64))
	}
	myAdmin, err = uuc.repo.GetAdminById(ctx, adminId)
	if nil == myAdmin {
		return res, err
	}
	if "super" != myAdmin.Type {
		return nil, errors.New(500, "ERROR_TOKEN", "非超管")
	}

	_, err = uuc.repo.CreateAdminAuth(ctx, req.SendBody.AdminId, req.SendBody.AuthId)
	if nil != err {
		return nil, errors.New(500, "ERROR_TOKEN", "创建失败")
	}

	return res, err
}

func (uuc *UserUseCase) AuthAdminDelete(ctx context.Context, req *v1.AuthAdminDeleteRequest) (*v1.AuthAdminDeleteReply, error) {
	var (
		myAdmin *Admin
		err     error
	)

	res := &v1.AuthAdminDeleteReply{}

	// 在上下文 context 中取出 claims 对象
	var adminId int64
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["UserId"] == nil {
			return nil, errors.New(500, "ERROR_TOKEN", "无效TOKEN")
		}
		adminId = int64(c["UserId"].(float64))
	}
	myAdmin, err = uuc.repo.GetAdminById(ctx, adminId)
	if nil == myAdmin {
		return res, err
	}
	if "super" != myAdmin.Type {
		return nil, errors.New(500, "ERROR_TOKEN", "非超管")
	}

	_, err = uuc.repo.DeleteAdminAuth(ctx, req.SendBody.AdminId, req.SendBody.AuthId)
	if nil != err {
		return nil, errors.New(500, "ERROR_TOKEN", "删除失败")
	}

	return res, err
}

func (uuc *UserUseCase) GetWithdrawPassOrRewardedFirst(ctx context.Context) (*Withdraw, error) {
	return uuc.ubRepo.GetWithdrawPassOrRewardedFirst(ctx)
}

func (uuc *UserUseCase) GetTradeOk(ctx context.Context) (*Trade, error) {
	return uuc.ubRepo.GetTradeOk(ctx)
}

func (uuc *UserUseCase) UpdateWithdrawDoing(ctx context.Context, id int64) (*Withdraw, error) {
	return uuc.ubRepo.UpdateWithdraw(ctx, id, "doing")
}

func (uuc *UserUseCase) UpdateWithdrawSuccess(ctx context.Context, id int64) (*Withdraw, error) {
	return uuc.ubRepo.UpdateWithdraw(ctx, id, "success")
}

func (uuc *UserUseCase) UpdateTrade(ctx context.Context, id int64) (*Trade, error) {
	return uuc.ubRepo.UpdateTrade(ctx, id, "okk")
}

func (uuc *UserUseCase) UpdateTradeDoing(ctx context.Context, id int64) (*Trade, error) {
	return uuc.ubRepo.UpdateTrade(ctx, id, "doing")
}

func (uuc *UserUseCase) AdminWithdrawList(ctx context.Context, req *v1.AdminWithdrawListRequest) (*v1.AdminWithdrawListReply, error) {
	var (
		withdraws  []*Withdraw
		userIds    []int64
		userSearch *User
		userId     int64
		userIdsMap map[int64]int64
		users      map[int64]*User
		count      int64
		err        error
	)

	res := &v1.AdminWithdrawListReply{
		Withdraw: make([]*v1.AdminWithdrawListReply_List, 0),
	}

	// 地址查询
	if "" != req.Address {
		userSearch, err = uuc.repo.GetUserByAddress(ctx, req.Address)
		if nil != err {
			return res, nil
		}
		userId = userSearch.ID
	}

	withdraws, err, count = uuc.ubRepo.GetWithdraws(ctx, &Pagination{
		PageNum:  int(req.Page),
		PageSize: 10,
	}, userId, req.WithDrawType)
	if nil != err {
		return res, err
	}
	res.Count = count

	userIdsMap = make(map[int64]int64, 0)
	for _, vWithdraws := range withdraws {
		userIdsMap[vWithdraws.UserId] = vWithdraws.UserId
	}
	for _, v := range userIdsMap {
		userIds = append(userIds, v)
	}

	users, err = uuc.repo.GetUserByUserIds(ctx, userIds...)
	if nil != err {
		return res, nil
	}

	for _, v := range withdraws {
		if _, ok := users[v.UserId]; !ok {
			continue
		}
		res.Withdraw = append(res.Withdraw, &v1.AdminWithdrawListReply_List{
			Id:        v.ID,
			CreatedAt: v.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
			Amount:    v.AmountNew,
			Status:    v.Status,
			Type:      v.Type,
			Address:   users[v.UserId].Address,
			RelAmount: v.AmountNewRel,
		})
	}

	return res, nil

}

func (uuc *UserUseCase) AdminFee(ctx context.Context, req *v1.AdminFeeRequest) (*v1.AdminFeeReply, error) {
	return &v1.AdminFeeReply{}, nil
}

func (uuc *UserUseCase) AdminFeeDaily(ctx context.Context, req *v1.AdminDailyFeeRequest) (*v1.AdminDailyFeeReply, error) {
	return &v1.AdminDailyFeeReply{}, nil
}

func (uuc *UserUseCase) AdminAll(ctx context.Context, req *v1.AdminAllRequest) (*v1.AdminAllReply, error) {

	var (
		userBalanceUsdtTotal            float64
		userBalanceRecordUsdtTotal      int64
		userBalanceRecordUsdtTotalToday int64
		userWithdrawUsdtTotalToday      float64
		userWithdrawUsdtTotal           float64
		//userBalanceDhbTotal              int64
		userLocationCount int64
		//userRewardLocationTotal          int64
		//userRewardAreaTotal              int64
		//userRewardRecommendTotal         int64
		//userRewardFourTotal              int64
		//userRewardRecommendLocationTotal int64
	)
	userLocationCount, _ = uuc.repo.GetUserCount(ctx)
	////userTodayCount, _ = uuc.repo.GetUserCountToday(ctx)
	userBalanceUsdtTotal, _ = uuc.ubRepo.GetUserBalanceUsdtTotal(ctx)
	//userBalanceDhbTotal, _ = uuc.ubRepo.GetUserBalanceDHBTotal(ctx)
	////userLocationNewCurrentMaxNew, _ = uuc.ubRepo.GetUserLocationNewCurrentMaxNew(ctx)
	////userLocationNewCurrentMax, _ = uuc.ubRepo.GetUserLocationNewCurrentMax(ctx)
	////userLocationNewCurrent, _ = uuc.ubRepo.GetUserLocationNewCurrent(ctx)
	////tmpUserLocationNewCurrent = userLocationNewCurrentMaxNew/100000000 + userLocationNewCurrentMax/100000000 - userLocationNewCurrent/100000000
	////userBalanceLockUsdtTotal, _ = uuc.ubRepo.GetUserBalanceLockUsdtTotal(ctx)
	userBalanceRecordUsdtTotal, _ = uuc.ubRepo.GetUserBalanceRecordUsdtTotal(ctx)
	////userBalanceRecordHbsTotal, _ = uuc.ubRepo.GetUserBalanceRecordHbsTotal(ctx)
	////userBalanceRecordCsdTotal, _ = uuc.ubRepo.GetUserBalanceRecordCsdTotal(ctx)
	userBalanceRecordUsdtTotalToday, _ = uuc.ubRepo.GetUserBalanceRecordUsdtTotalToday(ctx)
	userWithdrawUsdtTotalToday, _ = uuc.ubRepo.GetUserWithdrawUsdtTotalToday(ctx)
	////userWithdrawDhbTotalToday, _ = uuc.ubRepo.GetUserWithdrawDhbTotalToday(ctx)
	userWithdrawUsdtTotal, _ = uuc.ubRepo.GetUserWithdrawUsdtTotal(ctx)
	////userWithdrawDhbTotal, _ = uuc.ubRepo.GetUserWithdrawDhbTotal(ctx)
	////systemRewardUsdtTotal, _ = uuc.ubRepo.GetSystemRewardUsdtTotal(ctx)
	//userLocationCount = uuc.locationRepo.GetLocationUserCount(ctx)
	//
	//var (
	//	err error
	//)
	//userRewardLocationTotal, err = uuc.ubRepo.GetUserRewardLocationTotalToday(ctx, "location")
	//if nil != err {
	//
	//}
	//userRewardRecommendTotal, err = uuc.ubRepo.GetUserRewardLocationTotalToday(ctx, "recommend")
	//if nil != err {
	//
	//}
	//userRewardAreaTotal, err = uuc.ubRepo.GetUserRewardLocationTotalToday(ctx, "area")
	//if nil != err {
	//
	//}
	//userRewardFourTotal, err = uuc.ubRepo.GetUserRewardLocationTotalToday(ctx, "four")
	//if nil != err {
	//
	//}
	//userRewardRecommendLocationTotal, err = uuc.ubRepo.GetUserRewardLocationTotalToday(ctx, "recommend_location")
	//if nil != err {
	//
	//}
	//balanceRewardRewarded, _ = uuc.ubRepo.GetUserRewardBalanceRewardTotal(ctx)
	//balanceReward, _ = uuc.ubRepo.GetBalanceRewardTotal(ctx)
	//amountCsd, _ = uuc.ubRepo.GetTradeOkkCsd(ctx)
	//amountHbs, _ = uuc.ubRepo.GetTradeOkkHbs(ctx)

	return &v1.AdminAllReply{
		TotalUser:     userLocationCount,
		TodayLocation: fmt.Sprintf("%.2f", float64(userBalanceRecordUsdtTotalToday)/float64(100000)),
		AllLocation:   fmt.Sprintf("%.2f", float64(userBalanceRecordUsdtTotal)/float64(100000)),
		//TodayLocationReward:          fmt.Sprintf("%.2f", float64(userRewardLocationTotal)/float64(100000)),
		//TodayRecommendReward:         fmt.Sprintf("%.2f", float64(userRewardRecommendTotal)/float64(100000)),
		//TodayRecommendLocationReward: fmt.Sprintf("%.2f", float64(userRewardRecommendLocationTotal)/float64(100000)),
		//TodayAreaReward:              fmt.Sprintf("%.2f", float64(userRewardAreaTotal)/float64(100000)),
		//TodayFourReward:              fmt.Sprintf("%.2f", float64(userRewardFourTotal)/float64(100000)),
		//TotalIsps:                    fmt.Sprintf("%.2f", float64(userBalanceDhbTotal)/float64(100000)),
		TotalUsdt:         fmt.Sprintf("%.2f", float64(userBalanceUsdtTotal)/float64(100000)), // 可提biw
		TodayWithdrawUsdt: fmt.Sprintf("%.2f", float64(userWithdrawUsdtTotalToday)/float64(100000)),
		TotalWithdrawUsdt: fmt.Sprintf("%.2f", float64(userWithdrawUsdtTotal)/float64(100000)),
	}, nil
	return nil, nil
}

func (uuc *UserUseCase) GetConfigWithdrawDestroyRate(ctx context.Context) ([]*Config, error) {
	return uuc.configRepo.GetConfigByKeys(ctx, "withdraw_destroy_rate")
}

func (uuc *UserUseCase) AdminTrade(ctx context.Context, req *v1.AdminTradeRequest) (*v1.AdminTradeReply, error) {
	//time.Sleep(30 * time.Second) // 错开时间和充值
	var (
		tradeNotDeal                []*Trade
		configs                     []*Config
		withdrawRate                int64
		withdrawRecommendRate       int64
		withdrawRecommendSecondRate int64
		withdrawTeamVipRate         int64
		withdrawTeamVipSecondRate   int64
		withdrawTeamVipThirdRate    int64
		withdrawTeamVipFourthRate   int64
		withdrawTeamVipFifthRate    int64
		withdrawTeamVipLevelRate    int64
		vip0Balance                 int64
		err                         error
	)
	// 配置
	configs, _ = uuc.configRepo.GetConfigByKeys(ctx, "withdraw_rate",
		"withdraw_recommend_rate", "withdraw_recommend_second_rate",
		"withdraw_team_vip_rate", "withdraw_team_vip_second_rate",
		"withdraw_team_vip_third_rate", "withdraw_team_vip_fourth_rate",
		"withdraw_team_vip_fifth_rate", "withdraw_team_vip_level_rate",
		"withdraw_destroy_rate", "vip_0_balance",
	)
	if nil != configs {
		for _, vConfig := range configs {
			if "withdraw_rate" == vConfig.KeyName {
				withdrawRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "withdraw_recommend_rate" == vConfig.KeyName {
				withdrawRecommendRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "withdraw_recommend_second_rate" == vConfig.KeyName {
				withdrawRecommendSecondRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "withdraw_team_vip_rate" == vConfig.KeyName {
				withdrawTeamVipRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "withdraw_team_vip_second_rate" == vConfig.KeyName {
				withdrawTeamVipSecondRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "withdraw_team_vip_third_rate" == vConfig.KeyName {
				withdrawTeamVipThirdRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "withdraw_team_vip_fourth_rate" == vConfig.KeyName {
				withdrawTeamVipFourthRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "withdraw_team_vip_fifth_rate" == vConfig.KeyName {
				withdrawTeamVipFifthRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "withdraw_team_vip_level_rate" == vConfig.KeyName {
				withdrawTeamVipLevelRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "vip_0_balance" == vConfig.KeyName {
				vip0Balance, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
		}
	}

	tradeNotDeal, err = uuc.ubRepo.GetTradeNotDeal(ctx)
	if nil == tradeNotDeal {
		return &v1.AdminTradeReply{}, nil
	}

	for _, withdraw := range tradeNotDeal {
		if "default" != withdraw.Status {
			continue
		}

		//if "dhb" == withdraw.Type { // 提现dhb
		//	if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		//		_, err = uuc.ubRepo.UpdateWithdrawAmount(ctx, withdraw.ID, "rewarded", currentValue)
		//		if nil != err {
		//			return err
		//		}
		//
		//		return nil
		//	}); nil != err {
		//		return nil, err
		//	}
		//
		//	continue
		//}

		//withdraw.Amount*withdrawRate/100
		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			//currentValue -= withdraw.Amount * withdrawRate / 100 // 手续费
			//currentValue -= withdraw.Amount * withdrawDestroyRate / 100
			//fmt.Println(withdraw.Amount, currentValue)
			// 手续费记录
			//err = uuc.ubRepo.SystemFee(ctx, withdraw.Amount*withdrawRate/100, withdraw.ID)
			//if nil != err {
			//	return err
			//}

			_, err = uuc.ubRepo.UpdateTrade(ctx, withdraw.ID, "ok")
			if nil != err {
				return err
			}

			return nil
		}); nil != err {
			continue
		}

		var (
			userRecommend       *UserRecommend
			tmpRecommendUserIds []string
		)

		// 推荐人
		userRecommend, err = uuc.urRepo.GetUserRecommendByUserId(ctx, withdraw.UserId)
		if nil == userRecommend {
			continue
		}
		if "" != userRecommend.RecommendCode {
			tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")
		}

		lastKey := len(tmpRecommendUserIds) - 1
		if 1 > lastKey {
			continue
		}

		lastVip := int64(1)
		level1RewardCount := 1
		level2RewardCount := 1
		level3RewardCount := 1
		level4RewardCount := 1
		level5RewardCount := 1

		withdrawTeamVip := int64(0)
		levelOk := 0
		for i := 0; i <= lastKey; i++ {
			// 有占位信息，推荐人推荐人的上一代
			if lastKey-i <= 0 {
				break
			}

			tmpMyTopUserRecommendUserId, _ := strconv.ParseInt(tmpRecommendUserIds[lastKey-i], 10, 64) // 最后一位是直推人
			myUserTopRecommendUserInfo, _ := uuc.uiRepo.GetUserInfoByUserId(ctx, tmpMyTopUserRecommendUserId)
			if nil == myUserTopRecommendUserInfo {
				continue
			}
			//
			rewardAmount := withdraw.AmountCsd * withdrawRate / 100
			rewardAmountDhb := withdraw.AmountHbs * withdrawRate / 100
			tmpRecommendUserIdsInt := make([]int64, 0)
			if 1 < lastKey-i {
				for _, va := range tmpRecommendUserIds[1 : lastKey-i] {
					tmpRecommendUserIdsInt1, _ := strconv.ParseInt(va, 10, 64)
					tmpRecommendUserIdsInt = append(tmpRecommendUserIdsInt, tmpRecommendUserIdsInt1)
				}
			}

			if lastVip <= myUserTopRecommendUserInfo.Vip { // 上一个级别比我高
				// 会员团队
				if lastVip < myUserTopRecommendUserInfo.Vip && withdrawTeamVipFifthRate >= withdrawTeamVip {
					var tmp int64
					if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务

						if 2 == myUserTopRecommendUserInfo.Vip {
							tmp = withdrawTeamVipRate - withdrawTeamVip
							withdrawTeamVip = withdrawTeamVipRate

						} else if 3 == myUserTopRecommendUserInfo.Vip {
							tmp = withdrawTeamVipSecondRate - withdrawTeamVip
							withdrawTeamVip = withdrawTeamVipSecondRate

						} else if 4 == myUserTopRecommendUserInfo.Vip {
							tmp = withdrawTeamVipThirdRate - withdrawTeamVip
							withdrawTeamVip = withdrawTeamVipThirdRate

						} else if 5 == myUserTopRecommendUserInfo.Vip {
							tmp = withdrawTeamVipFourthRate - withdrawTeamVip
							withdrawTeamVip = withdrawTeamVipFourthRate

						} else if 6 == myUserTopRecommendUserInfo.Vip {
							tmp = withdrawTeamVipFifthRate - withdrawTeamVip
							withdrawTeamVip = withdrawTeamVipFifthRate
						}

						_, err = uuc.ubRepo.WithdrawNewRewardTeamRecommend(ctx, myUserTopRecommendUserInfo.UserId, rewardAmount*tmp/100, rewardAmountDhb*tmp/100, withdraw.ID, tmpRecommendUserIdsInt)
						if nil != err {
							return err
						}

						return nil
					}); nil != err {
						continue
					}

					lastVip = myUserTopRecommendUserInfo.Vip
					levelOk = 1
					continue
				}

				// 平级奖
				if 0 < levelOk && lastVip == myUserTopRecommendUserInfo.Vip { // 上一个是vip1和以上且和我平级
					tmpCurrent := 0
					if 2 == myUserTopRecommendUserInfo.Vip {
						if 0 < level1RewardCount {
							tmpCurrent = level1RewardCount
							level1RewardCount--
						}
					} else if 3 == myUserTopRecommendUserInfo.Vip {
						if 0 < level2RewardCount {
							tmpCurrent = level2RewardCount
							level2RewardCount--
						}
					} else if 4 == myUserTopRecommendUserInfo.Vip {
						if 0 < level3RewardCount {
							tmpCurrent = level3RewardCount
							level3RewardCount--
						}
					} else if 5 == myUserTopRecommendUserInfo.Vip {
						if 0 < level4RewardCount {
							tmpCurrent = level4RewardCount
							level4RewardCount--
						}
					} else if 6 == myUserTopRecommendUserInfo.Vip {
						if 0 < level5RewardCount {
							tmpCurrent = level5RewardCount
							level5RewardCount--
						}
					}

					if 0 < tmpCurrent {
						if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
							_, err = uuc.ubRepo.WithdrawNewRewardLevelRecommend(ctx, myUserTopRecommendUserInfo.UserId, rewardAmount*withdrawTeamVipLevelRate/100, rewardAmountDhb*withdrawTeamVipLevelRate/100, withdraw.ID, tmpRecommendUserIdsInt)
							if nil != err {
								return err
							}

							return nil
						}); nil != err {
							continue
						}

						lastVip = myUserTopRecommendUserInfo.Vip
						continue
					}
				}
			}

			if 0 == i { // 当前用户被此人直推

				var userBalance *UserBalance
				userBalance, err = uuc.ubRepo.GetUserBalance(ctx, myUserTopRecommendUserInfo.UserId)
				if nil != err {
					continue
				}

				if userBalance.BalanceUsdt/100000 < vip0Balance {
					continue
				}

				if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
					_, err = uuc.ubRepo.WithdrawNewRewardRecommend(ctx, myUserTopRecommendUserInfo.UserId, rewardAmount*withdrawRecommendRate/100, rewardAmountDhb*withdrawRecommendRate/100, withdraw.ID, tmpRecommendUserIdsInt)
					if nil != err {
						return err
					}

					return nil
				}); nil != err {
					continue
				}

				continue
			} else if 1 == i { // 间接推
				var userBalance *UserBalance
				userBalance, err = uuc.ubRepo.GetUserBalance(ctx, myUserTopRecommendUserInfo.UserId)
				if nil != err {
					continue
				}

				if userBalance.BalanceUsdt/100000 < vip0Balance {
					continue
				}

				if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
					_, err = uuc.ubRepo.WithdrawNewRewardSecondRecommend(ctx, myUserTopRecommendUserInfo.UserId, rewardAmount*withdrawRecommendSecondRate/100, rewardAmountDhb*withdrawRecommendSecondRate/100, withdraw.ID, tmpRecommendUserIdsInt)
					if nil != err {
						return err
					}

					return nil
				}); nil != err {
					continue
				}

				continue
			}
		}
	}

	return &v1.AdminTradeReply{}, nil
}

func (uuc *UserUseCase) AdminDailyBalanceReward(ctx context.Context, req *v1.AdminDailyBalanceRewardRequest) (*v1.AdminDailyBalanceRewardReply, error) {
	var (
		balanceRewards    []*BalanceReward
		configs           []*Config
		balanceRewardRate int64
		coinPrice         int64
		coinRewardRate    int64
		rewardRate        int64
		err               error
	)
	configs, _ = uuc.configRepo.GetConfigByKeys(ctx, "balance_reward_rate", "coin_price", "coin_reward_rate", "reward_rate")
	if nil != configs {
		for _, vConfig := range configs {
			if "balance_reward_rate" == vConfig.KeyName {
				balanceRewardRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "coin_price" == vConfig.KeyName {
				coinPrice, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "coin_reward_rate" == vConfig.KeyName {
				coinRewardRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "reward_rate" == vConfig.KeyName {
				rewardRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
		}
	}

	now := time.Now()
	if "" != req.Date { // 测试条件
		now, err = time.Parse("2006-01-02 15:04:05", req.Date) // 时间进行格式校验
		if nil != err {
			return nil, err
		}
	}

	now = now.UTC()
	balanceRewards, err = uuc.ubRepo.GetBalanceRewardCurrent(ctx, now)

	timeLimit := time.Now().UTC().Add(-23 * time.Hour)

	for _, vBalanceRewards := range balanceRewards {
		if "" == req.Date { // 测试条件
			if vBalanceRewards.LastRewardDate.After(timeLimit) {
				continue
			}
		}

		// 今天发
		tmpCurrentReward := vBalanceRewards.Amount * balanceRewardRate / 1000
		var myLocationLast *LocationNew
		// 获取当前用户的占位信息，已经有运行中的跳过
		myLocationLast, err = uuc.locationRepo.GetMyLocationLast(ctx, vBalanceRewards.UserId)
		if nil == myLocationLast { // 无占位信息
			continue
		}

		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			tmpCurrentStatus := myLocationLast.Status // 现在还在运行中

			tmpBalanceUsdtAmount := tmpCurrentReward * rewardRate / 100 // 记录下一次
			tmpBalanceCoinAmount := tmpCurrentReward * coinRewardRate / 100 * 1000 / coinPrice

			myLocationLast.Status = "running"
			myLocationLast.Current += tmpCurrentReward
			if myLocationLast.Current >= myLocationLast.CurrentMax { // 占位分红人分满停止
				if "running" == tmpCurrentStatus {
					myLocationLast.StopDate = time.Now().UTC().Add(8 * time.Hour)

					lastRewardAmount := tmpCurrentReward - (myLocationLast.Current - myLocationLast.CurrentMax)
					tmpBalanceUsdtAmount = lastRewardAmount * rewardRate / 100 // 记录下一次
					tmpBalanceCoinAmount = lastRewardAmount * coinRewardRate / 100 * 1000 / coinPrice
				}
				myLocationLast.Status = "stop"
			}

			if 0 < tmpCurrentReward {
				err = uuc.locationRepo.UpdateLocationNew(ctx, myLocationLast.ID, myLocationLast.Status, tmpCurrentReward, myLocationLast.StopDate) // 分红占位数据修改
				if nil != err {
					return err
				}

				_, err = uuc.ubRepo.UserDailyBalanceReward(ctx, vBalanceRewards.UserId, tmpCurrentReward, tmpBalanceUsdtAmount, tmpBalanceCoinAmount, tmpCurrentStatus)
				if nil != err {
					return err
				}

				err = uuc.ubRepo.UpdateBalanceRewardLastRewardDate(ctx, vBalanceRewards.ID)
				if nil != err {
					return err
				}
			}

			return nil
		}); nil != err {
			continue
		}

	}

	return &v1.AdminDailyBalanceRewardReply{}, nil
}

func (uuc *UserUseCase) AdminDailyBuyReward(ctx context.Context, req *v1.AdminDailyLocationRewardRequest) (*v1.AdminDailyLocationRewardReply, error) {
	var (
		configs   []*Config
		one       float64
		two       float64
		three     float64
		four      float64
		five      float64
		oneTwo    float64
		twoTwo    float64
		threeTwo  float64
		fiveThree float64
		fourThree float64
		today     float64
		err       error
	)

	configs, err = uuc.configRepo.GetConfigByKeys(ctx,
		"one", "two", "three", "four", "five",
		"one_two", "two_two", "three_two", "four_two", "five_two", "four_three", "five_three", "today",
	)
	if nil != err {
		fmt.Println("分红,配置", err)
		return nil, err
	}

	if nil != configs {
		for _, vConfig := range configs {
			if "one" == vConfig.KeyName {
				one, err = strconv.ParseFloat(vConfig.Value, 10)
				if nil != err {
					return nil, err
				}
			} else if "two" == vConfig.KeyName {
				two, err = strconv.ParseFloat(vConfig.Value, 10)
				if nil != err {
					return nil, err
				}
			} else if "three" == vConfig.KeyName {
				three, err = strconv.ParseFloat(vConfig.Value, 10)
				if nil != err {
					return nil, err
				}
			} else if "four" == vConfig.KeyName {
				four, err = strconv.ParseFloat(vConfig.Value, 10)
				if nil != err {
					return nil, err
				}
			} else if "five" == vConfig.KeyName {
				five, err = strconv.ParseFloat(vConfig.Value, 10)
				if nil != err {
					return nil, err
				}
			} else if "one_two" == vConfig.KeyName {
				oneTwo, err = strconv.ParseFloat(vConfig.Value, 10)
				if nil != err {
					return nil, err
				}
			} else if "two_two" == vConfig.KeyName {
				twoTwo, err = strconv.ParseFloat(vConfig.Value, 10)
				if nil != err {
					return nil, err
				}
			} else if "three_two" == vConfig.KeyName {
				threeTwo, err = strconv.ParseFloat(vConfig.Value, 10)
				if nil != err {
					return nil, err
				}
			} else if "four_three" == vConfig.KeyName {
				fourThree, err = strconv.ParseFloat(vConfig.Value, 10)
				if nil != err {
					return nil, err
				}
			} else if "five_three" == vConfig.KeyName {
				fiveThree, err = strconv.ParseFloat(vConfig.Value, 10)
				if nil != err {
					return nil, err
				}
			} else if "today" == vConfig.KeyName {
				today, err = strconv.ParseFloat(vConfig.Value, 10)
				if nil != err {
					return nil, err
				}
			}
		}
	}

	var (
		users []*User
	)
	users, err = uuc.repo.GetUsersNew(ctx)
	if nil != err {
		fmt.Println("分红", err)
		return nil, nil
	}

	fourUsers := make([]*User, 0)
	fiveUsers := make([]*User, 0)

	var (
		amountSecond float64
	)
	for _, vUsers := range users {
		// 定义阈值数组
		var (
			amount float64
		)

		// 按顺序遍历阈值数组
		number := vUsers.Total
		if 30000 <= number {
			fiveUsers = append(fiveUsers, vUsers)
		} else if 15000 <= number {
			fourUsers = append(fourUsers, vUsers)
		}

		for number > 0 {
			if 30000 <= number {
				number -= 30000
				amount += five
			} else if 15000 <= number {
				number -= 15000
				amount += four
			} else if 5000 <= number {
				number -= 5000
				amount += three
				amountSecond += threeTwo
			} else if 3000 <= number {
				number -= 3000
				amount += two
				amountSecond += twoTwo
			} else if 1000 <= number {
				number -= 1000
				amount += one
				amountSecond += oneTwo
			} else {
				break
			}
		}

		// 发奖励
		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务

			// 充值记录
			err = uuc.ubRepo.FirstRewardBiw(ctx, vUsers.ID, amount)
			if nil != err {
				return err
			}

			return nil
		}); nil != err {
			fmt.Println(err, "错误发奖励1", vUsers, amount)
			continue
		}

	}

	// 查询昨日入金
	// 全网
	var (
		perFour float64
		perFive float64
	)

	if 0 < today {
		amountSecond = today
	}

	if 0 >= amountSecond {
		return nil, nil
	}

	if 0 < len(fourUsers) {
		perFour = amountSecond * fourThree / 100 / float64(len(fourUsers))
		fmt.Println(amountSecond, fourThree, perFour)
		if 0 < perFour {
			for _, vFourUser := range fourUsers {
				// 发奖励
				if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务

					err = uuc.ubRepo.SecondRewardBiw(ctx, vFourUser.ID, perFour, "four")
					if nil != err {
						return err
					}

					return nil
				}); nil != err {
					fmt.Println(err, "错误发奖励2,4", vFourUser, perFour, "four")
					continue
				}
			}
		}
	}

	if 0 < len(fiveUsers) {
		perFive = amountSecond * fiveThree / 100 / float64(len(fiveUsers))
		fmt.Println(amountSecond, fiveThree, perFive)

		if 0 < perFive {
			for _, vFiveUser := range fiveUsers {
				// 发奖励
				if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务

					err = uuc.ubRepo.SecondRewardBiw(ctx, vFiveUser.ID, perFive, "five")
					if nil != err {
						return err
					}

					return nil
				}); nil != err {
					fmt.Println(err, "错误发奖励2,5", vFiveUser, perFive)
					continue
				}
			}
		}
	}

	return nil, nil
}

func (uuc *UserUseCase) AdminDailyLocationReward(ctx context.Context, req *v1.AdminDailyLocationRewardRequest) (*v1.AdminDailyLocationRewardReply, error) {
	var (
		userLocations      []*LocationNew
		configs            []*Config
		locationRewardRate int64
		bPrice             int64
		bPriceBase         int64
		recommendOneRate   int64
		recommendTwoRate   int64
		recommendThreeRate int64
		recommendFourRate  int64
		recommendFiveRate  int64
		recommendSixRate   int64
		recommendSevenRate int64
		recommendEightRate int64
		feeRate            int64
		err                error
	)

	configs, _ = uuc.configRepo.GetConfigByKeys(ctx,
		"location_reward_rate", "b_price", "b_price_base", "exchange_rate",
		"recommend_one_rate", "recommend_two_rate",
		"recommend_three_rate", "recommend_four_rate",
		"recommend_five_rate", "recommend_six_rate",
		"recommend_seven_rate", "recommend_eight_rate",
		"area_one", "area_two", "area_three", "area_four", "area_five",
		"area_num_one", "area_num_two", "area_num_three", "area_num_four", "area_num_five", "one", "two", "three", "four", "total",
	)
	if nil != configs {
		for _, vConfig := range configs {
			if "location_reward_rate" == vConfig.KeyName {
				locationRewardRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "b_price" == vConfig.KeyName {
				bPrice, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "b_price_base" == vConfig.KeyName {
				bPriceBase, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "recommend_one_rate" == vConfig.KeyName {
				recommendOneRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "recommend_two_rate" == vConfig.KeyName {
				recommendTwoRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "recommend_three_rate" == vConfig.KeyName {
				recommendThreeRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "recommend_four_rate" == vConfig.KeyName {
				recommendFourRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "recommend_five_rate" == vConfig.KeyName {
				recommendFiveRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "recommend_six_rate" == vConfig.KeyName {
				recommendSixRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "recommend_seven_rate" == vConfig.KeyName {
				recommendSevenRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "recommend_eight_rate" == vConfig.KeyName {
				recommendEightRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "exchange_rate" == vConfig.KeyName {
				feeRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
		}
	}

	var (
		stopLocationIds map[int64]int64
	)

	stopLocationIds = make(map[int64]int64, 0)
	// 每日
	userLocations, err = uuc.locationRepo.GetRunningLocations(ctx)
	if nil != err {
		return &v1.AdminDailyLocationRewardReply{}, nil
	}

	for _, vUserLocations := range userLocations {
		// 奖励
		tmpCurrentReward := vUserLocations.Usdt / 1000 * locationRewardRate
		tmpCurrentReward = tmpCurrentReward / 6
		bLocationRewardAmount := tmpCurrentReward * bPriceBase / bPrice

		if 0 < tmpCurrentReward && 0 < bLocationRewardAmount {
			if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
				tmpStatus := vUserLocations.Status
				tmpStopDate := time.Now().UTC().Add(8 * time.Hour)
				if vUserLocations.Current+tmpCurrentReward >= vUserLocations.CurrentMax { // 占位分红人分满停止
					tmpStatus = "stop"
					tmpStopDate = time.Now().UTC().Add(8 * time.Hour)

					tmpCurrentReward = vUserLocations.CurrentMax - vUserLocations.Current
					bLocationRewardAmount = tmpCurrentReward * bPriceBase / bPrice
					stopLocationIds[vUserLocations.ID] = vUserLocations.ID
				}

				var tmpMaxNew int64
				if vUserLocations.CurrentMaxNew < vUserLocations.CurrentMax {
					tmpMaxNew = vUserLocations.CurrentMax - vUserLocations.CurrentMaxNew
				}
				if 0 < tmpCurrentReward && 0 < bLocationRewardAmount {
					err = uuc.locationRepo.UpdateLocationNewNew(ctx, vUserLocations.ID, tmpStatus, tmpCurrentReward, tmpMaxNew, bLocationRewardAmount, tmpStopDate) // 分红占位数据修改
					if nil != err {
						return err
					}

					_, err = uuc.ubRepo.LocationRewardBiw(ctx, vUserLocations.UserId, bLocationRewardAmount, tmpStatus, tmpMaxNew, feeRate)
					if nil != err {
						return err
					}
				}

				// 业绩减掉
				if "stop" == tmpStatus {
					tmpTop := vUserLocations.Top
					tmpTopNum := vUserLocations.TopNum
					for j := 0; j < 10000 && 0 < tmpTop && 0 < tmpTopNum; j++ {
						err = uuc.locationRepo.UpdateLocationNewTotalSub(ctx, tmpTop, tmpTopNum, vUserLocations.Usdt/100000)
						if nil != err {
							return err
						}

						var (
							currentLocation *LocationNew
						)
						currentLocation, err = uuc.locationRepo.GetLocationById(ctx, tmpTop)
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
				fmt.Println("err reward daily", err, vUserLocations)
				continue
			}
		}
	}

	for _, vUserLocations := range userLocations {
		var (
			userRecommend       *UserRecommend
			tmpRecommendUserIds []string
		)

		// 推荐人
		userRecommend, err = uuc.urRepo.GetUserRecommendByUserId(ctx, vUserLocations.UserId)
		if nil != userRecommend {
			if "" != userRecommend.RecommendCode {
				tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")
			}

			lastKey := len(tmpRecommendUserIds) - 1
			if 1 <= lastKey {
				for i := 0; i <= 7; i++ {
					// 有占位信息，推荐人推荐人的上一代
					if lastKey-i <= 0 {
						break
					}

					tmpMyTopUserRecommendUserId, _ := strconv.ParseInt(tmpRecommendUserIds[lastKey-i], 10, 64) // 最后一位是直推人
					if 0 >= tmpMyTopUserRecommendUserId {
						break
					}

					var myUserRecommendUserLocationsLast []*LocationNew
					myUserRecommendUserLocationsLast, err = uuc.locationRepo.GetLocationsNewByUserId(ctx, tmpMyTopUserRecommendUserId)
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

							if _, ok := stopLocationIds[tmpMyTopUserRecommendUserLocationLast.ID]; ok { // 上一轮已经停止了
								continue
							}

							tmpMinUsdt := tmpMyTopUserRecommendUserLocationLast.Usdt
							if vUserLocations.Usdt < tmpMinUsdt {
								tmpMinUsdt = vUserLocations.Usdt
							}

							var tmpMyRecommendAmount int64
							if 0 == i { // 当前用户被此人直推
								tmpMyRecommendAmount = tmpMinUsdt / 1000 * locationRewardRate / 100 * recommendOneRate
							} else if 1 == i {
								tmpMyRecommendAmount = tmpMinUsdt / 1000 * locationRewardRate / 100 * recommendTwoRate
							} else if 2 == i && 2 <= len(myUserRecommendUserLocationsLast) { // 3代需要复投1次
								tmpMyRecommendAmount = tmpMinUsdt / 1000 * locationRewardRate / 100 * recommendThreeRate
							} else if 3 == i && 3 <= len(myUserRecommendUserLocationsLast) { // 4代需要复投2次
								tmpMyRecommendAmount = tmpMinUsdt / 1000 * locationRewardRate / 100 * recommendFourRate
							} else if 4 == i && 4 <= len(myUserRecommendUserLocationsLast) { // 5代需要复投3次
								tmpMyRecommendAmount = tmpMinUsdt / 1000 * locationRewardRate / 100 * recommendFiveRate
							} else if 5 == i && 5 <= len(myUserRecommendUserLocationsLast) { // 6代需要复投4次
								tmpMyRecommendAmount = tmpMinUsdt / 1000 * locationRewardRate / 100 * recommendSixRate
							} else if 6 == i && 6 <= len(myUserRecommendUserLocationsLast) { // 7代需要复投5次
								tmpMyRecommendAmount = tmpMinUsdt / 1000 * locationRewardRate / 100 * recommendSevenRate
							} else if 7 == i && 7 <= len(myUserRecommendUserLocationsLast) { // 8代需要复投6次
								tmpMyRecommendAmount = tmpMinUsdt / 1000 * locationRewardRate / 100 * recommendEightRate
							} else {
								continue
							}

							tmpMyRecommendAmount = tmpMyRecommendAmount / 6
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
									stopLocationIds[vUserLocations.ID] = vUserLocations.ID
								}

								if 0 < tmpMyRecommendAmount && 0 < bAmount {
									var tmpMaxNew int64
									if tmpMyTopUserRecommendUserLocationLast.CurrentMaxNew < tmpMyTopUserRecommendUserLocationLast.CurrentMax {
										tmpMaxNew = tmpMyTopUserRecommendUserLocationLast.CurrentMax - tmpMyTopUserRecommendUserLocationLast.CurrentMaxNew
									}

									if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
										err = uuc.locationRepo.UpdateLocationNewNew(ctx, tmpMyTopUserRecommendUserLocationLast.ID, tmpStatus, tmpMyRecommendAmount, tmpMaxNew, bAmount, tmpStopDate) // 分红占位数据修改
										if nil != err {
											return err
										}

										_, err = uuc.ubRepo.RecommendRewardBiw(ctx, tmpMyTopUserRecommendUserId, bAmount, int64(i+1), tmpStatus, tmpMaxNew, feeRate) // 推荐人奖励
										if nil != err {
											return err
										}

										// 业绩减掉
										if "stop" == tmpStatus {
											tmpTop := tmpMyTopUserRecommendUserLocationLast.Top
											tmpTopNum := tmpMyTopUserRecommendUserLocationLast.TopNum
											for j := 0; j < 10000 && 0 < tmpTop && 0 < tmpTopNum; j++ {
												err = uuc.locationRepo.UpdateLocationNewTotalSub(ctx, tmpTop, tmpTopNum, tmpMyTopUserRecommendUserLocationLast.Usdt/100000)
												if nil != err {
													return err
												}

												var (
													currentLocation *LocationNew
												)
												currentLocation, err = uuc.locationRepo.GetLocationById(ctx, tmpTop)
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
	}

	return nil, err
}

func (uuc *UserUseCase) AdminDailyAreaReward(ctx context.Context, req *v1.AdminDailyLocationRewardRequest) (*v1.AdminDailyLocationRewardReply, error) {
	var (
		userLocations []*LocationNew
		configs       []*Config
		bPrice        int64
		bPriceBase    int64
		areaOne       int64
		areaTwo       int64
		areaThree     int64
		areaFour      int64
		areaFive      int64
		areaNumOne    int64
		areaNumTwo    int64
		areaNumThree  int64
		areaNumFour   int64
		areaNumFive   int64
		one           int64
		two           int64
		three         int64
		four          int64
		total         int64
		feeRate       int64
		err           error
	)

	configs, _ = uuc.configRepo.GetConfigByKeys(ctx,
		"location_reward_rate", "b_price", "b_price_base", "exchange_rate",
		"recommend_one_rate", "recommend_two_rate",
		"recommend_three_rate", "recommend_four_rate",
		"recommend_five_rate", "recommend_six_rate",
		"recommend_seven_rate", "recommend_eight_rate",
		"area_one", "area_two", "area_three", "area_four", "area_five",
		"area_num_one", "area_num_two", "area_num_three", "area_num_four", "area_num_five", "one", "two", "three", "four", "total",
	)
	if nil != configs {
		for _, vConfig := range configs {
			if "b_price" == vConfig.KeyName {
				bPrice, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "b_price_base" == vConfig.KeyName {
				bPriceBase, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "area_one" == vConfig.KeyName {
				areaOne, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "area_two" == vConfig.KeyName {
				areaTwo, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "area_three" == vConfig.KeyName {
				areaThree, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "area_four" == vConfig.KeyName {
				areaFour, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "area_five" == vConfig.KeyName {
				areaFive, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "area_num_one" == vConfig.KeyName {
				areaNumOne, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "area_num_two" == vConfig.KeyName {
				areaNumTwo, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "area_num_three" == vConfig.KeyName {
				areaNumThree, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "area_num_four" == vConfig.KeyName {
				areaNumFour, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "area_num_five" == vConfig.KeyName {
				areaNumFive, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "one" == vConfig.KeyName {
				one, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "two" == vConfig.KeyName {
				two, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "three" == vConfig.KeyName {
				three, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "four" == vConfig.KeyName {
				four, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "total" == vConfig.KeyName {
				total, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "exchange_rate" == vConfig.KeyName {
				feeRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
		}
	}

	// 获取今日收益
	var (
		day               = -1
		userLocationsYes  []*LocationNew
		userLocationsBef  []*LocationNew
		rewardLocationYes int64
		rewardLocationBef int64
	)
	// 全网
	userLocationsYes, err = uuc.locationRepo.GetLocationDailyYesterday(ctx, day)
	for _, userLocationYes := range userLocationsYes {
		rewardLocationYes += userLocationYes.Usdt
	}

	if 0 >= rewardLocationYes {
		return &v1.AdminDailyLocationRewardReply{}, nil
	}

	// 团队奖励
	userLocationsOne := make([]*LocationNew, 0)
	userLocations, err = uuc.locationRepo.GetRunningLocations(ctx)
	if nil != err {
		return &v1.AdminDailyLocationRewardReply{}, nil
	}
	for _, vUserLocations := range userLocations {
		// 1大区
		if vUserLocations.Total >= vUserLocations.TotalTwo && vUserLocations.Total >= vUserLocations.TotalThree {
			if areaOne <= vUserLocations.TotalTwo+vUserLocations.TotalThree || 1 <= vUserLocations.LastLevel {
				userLocationsOne = append(userLocationsOne, vUserLocations)
			}

		} else if vUserLocations.TotalTwo >= vUserLocations.Total && vUserLocations.TotalTwo >= vUserLocations.TotalThree {
			if areaOne <= vUserLocations.Total+vUserLocations.TotalThree || 1 <= vUserLocations.LastLevel {
				userLocationsOne = append(userLocationsOne, vUserLocations)
			}

		} else if vUserLocations.TotalThree >= vUserLocations.Total && vUserLocations.TotalThree >= vUserLocations.TotalTwo {
			if areaOne <= vUserLocations.TotalTwo+vUserLocations.Total || 1 <= vUserLocations.LastLevel {
				userLocationsOne = append(userLocationsOne, vUserLocations)
			}
		}
	}

	if 0 < len(userLocationsOne) {
		rewardLocationYesOne := rewardLocationYes / 1000 * areaNumOne / int64(len(userLocationsOne))
		if 0 < rewardLocationYesOne {
			for _, vUserLocationsItem := range userLocationsOne {
				// 奖励
				tmpCurrentReward := rewardLocationYesOne
				bLocationRewardAmount := tmpCurrentReward * bPriceBase / bPrice

				if 0 < tmpCurrentReward && 0 < bLocationRewardAmount {
					if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
						if vUserLocationsItem.Current+tmpCurrentReward >= vUserLocationsItem.CurrentMax { // 占位分红人分满停止
							vUserLocationsItem.Status = "stop"
							vUserLocationsItem.StopDate = time.Now().UTC().Add(8 * time.Hour)

							tmpCurrentReward = vUserLocationsItem.CurrentMax - vUserLocationsItem.Current
							bLocationRewardAmount = tmpCurrentReward * bPriceBase / bPrice
						}

						var tmpMaxNew int64
						if vUserLocationsItem.CurrentMaxNew < vUserLocationsItem.CurrentMax {
							tmpMaxNew = vUserLocationsItem.CurrentMax - vUserLocationsItem.CurrentMaxNew
						}

						if 0 < tmpCurrentReward && 0 < bLocationRewardAmount {
							err = uuc.locationRepo.UpdateLocationNewNew(ctx, vUserLocationsItem.ID, vUserLocationsItem.Status, tmpCurrentReward, tmpMaxNew, bLocationRewardAmount, vUserLocationsItem.StopDate) // 分红占位数据修改
							if nil != err {
								return err
							}

							_, err = uuc.ubRepo.AreaRewardBiw(ctx, vUserLocationsItem.UserId, bLocationRewardAmount, tmpCurrentReward, 1, vUserLocationsItem.Status, tmpMaxNew, feeRate)
							if nil != err {
								return err
							}
						}

						// 业绩减掉
						if "stop" == vUserLocationsItem.Status {
							tmpTop := vUserLocationsItem.Top
							tmpTopNum := vUserLocationsItem.TopNum
							for j := 0; j < 10000 && 0 < tmpTop && 0 < tmpTopNum; j++ {
								err = uuc.locationRepo.UpdateLocationNewTotalSub(ctx, tmpTop, tmpTopNum, vUserLocationsItem.Usdt/100000)
								if nil != err {
									return err
								}

								var (
									currentLocation *LocationNew
								)
								currentLocation, err = uuc.locationRepo.GetLocationById(ctx, tmpTop)
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
						fmt.Println("err reward daily three", err, vUserLocationsItem)
						continue
					}
				}
			}
		}
	}

	// 团队奖励
	userLocationsTwo := make([]*LocationNew, 0)
	userLocations, err = uuc.locationRepo.GetRunningLocations(ctx)
	if nil != err {
		return &v1.AdminDailyLocationRewardReply{}, nil
	}
	for _, vUserLocations := range userLocations {
		// 1大区
		if vUserLocations.Total >= vUserLocations.TotalTwo && vUserLocations.Total >= vUserLocations.TotalThree {
			if areaTwo <= vUserLocations.TotalTwo+vUserLocations.TotalThree || 2 <= vUserLocations.LastLevel {
				userLocationsTwo = append(userLocationsTwo, vUserLocations)
			}
		} else if vUserLocations.TotalTwo >= vUserLocations.Total && vUserLocations.TotalTwo >= vUserLocations.TotalThree {
			if areaTwo <= vUserLocations.Total+vUserLocations.TotalThree || 2 <= vUserLocations.LastLevel {
				userLocationsTwo = append(userLocationsTwo, vUserLocations)
			}
		} else if vUserLocations.TotalThree >= vUserLocations.Total && vUserLocations.TotalThree >= vUserLocations.TotalTwo {
			if areaTwo <= vUserLocations.TotalTwo+vUserLocations.Total || 2 <= vUserLocations.LastLevel {
				userLocationsTwo = append(userLocationsTwo, vUserLocations)
			}
		}
	}

	if 0 < len(userLocationsTwo) {
		rewardLocationYesTwo := rewardLocationYes / 1000 * areaNumTwo / int64(len(userLocationsTwo))
		if 0 < rewardLocationYesTwo {
			for _, vUserLocationsItem := range userLocationsTwo {
				// 奖励
				tmpCurrentReward := rewardLocationYesTwo
				bLocationRewardAmount := tmpCurrentReward * bPriceBase / bPrice

				if 0 < tmpCurrentReward && 0 < bLocationRewardAmount {
					if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
						if vUserLocationsItem.Current+tmpCurrentReward >= vUserLocationsItem.CurrentMax { // 占位分红人分满停止
							vUserLocationsItem.Status = "stop"
							vUserLocationsItem.StopDate = time.Now().UTC().Add(8 * time.Hour)

							tmpCurrentReward = vUserLocationsItem.CurrentMax - vUserLocationsItem.Current
							bLocationRewardAmount = tmpCurrentReward * bPriceBase / bPrice
						}

						var tmpMaxNew int64
						if vUserLocationsItem.CurrentMaxNew < vUserLocationsItem.CurrentMax {
							tmpMaxNew = vUserLocationsItem.CurrentMax - vUserLocationsItem.CurrentMaxNew
						}

						if 0 < tmpCurrentReward && 0 < bLocationRewardAmount {
							err = uuc.locationRepo.UpdateLocationNewNew(ctx, vUserLocationsItem.ID, vUserLocationsItem.Status, tmpCurrentReward, tmpMaxNew, bLocationRewardAmount, vUserLocationsItem.StopDate) // 分红占位数据修改
							if nil != err {
								return err
							}

							_, err = uuc.ubRepo.AreaRewardBiw(ctx, vUserLocationsItem.UserId, bLocationRewardAmount, tmpCurrentReward, 2, vUserLocationsItem.Status, tmpMaxNew, feeRate)
							if nil != err {
								return err
							}
						}

						// 业绩减掉
						if "stop" == vUserLocationsItem.Status {
							tmpTop := vUserLocationsItem.Top
							tmpTopNum := vUserLocationsItem.TopNum
							for j := 0; j < 10000 && 0 < tmpTop && 0 < tmpTopNum; j++ {
								err = uuc.locationRepo.UpdateLocationNewTotalSub(ctx, tmpTop, tmpTopNum, vUserLocationsItem.Usdt/100000)
								if nil != err {
									return err
								}

								var (
									currentLocation *LocationNew
								)
								currentLocation, err = uuc.locationRepo.GetLocationById(ctx, tmpTop)
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
						fmt.Println("err reward daily three", err, vUserLocationsItem)
						continue
					}
				}
			}
		}
	}

	userLocationsThree := make([]*LocationNew, 0)
	userLocations, err = uuc.locationRepo.GetRunningLocations(ctx)
	if nil != err {
		return &v1.AdminDailyLocationRewardReply{}, nil
	}
	for _, vUserLocations := range userLocations {
		// 1大区
		if vUserLocations.Total >= vUserLocations.TotalTwo && vUserLocations.Total >= vUserLocations.TotalThree {
			if areaThree <= vUserLocations.TotalTwo+vUserLocations.TotalThree || 3 <= vUserLocations.LastLevel {
				userLocationsThree = append(userLocationsThree, vUserLocations)
			}
		} else if vUserLocations.TotalTwo >= vUserLocations.Total && vUserLocations.TotalTwo >= vUserLocations.TotalThree {
			if areaThree <= vUserLocations.Total+vUserLocations.TotalThree || 3 <= vUserLocations.LastLevel {
				userLocationsThree = append(userLocationsThree, vUserLocations)
			}
		} else if vUserLocations.TotalThree >= vUserLocations.Total && vUserLocations.TotalThree >= vUserLocations.TotalTwo {
			if areaThree <= vUserLocations.TotalTwo+vUserLocations.Total || 3 <= vUserLocations.LastLevel {
				userLocationsThree = append(userLocationsThree, vUserLocations)
			}
		}
	}

	if 0 < len(userLocationsThree) {
		rewardLocationYesThree := rewardLocationYes / 1000 * areaNumThree / int64(len(userLocationsThree))
		if 0 < rewardLocationYesThree {
			for _, vUserLocationsItem := range userLocationsThree {
				// 奖励
				tmpCurrentReward := rewardLocationYesThree
				bLocationRewardAmount := tmpCurrentReward * bPriceBase / bPrice

				if 0 < tmpCurrentReward && 0 < bLocationRewardAmount {
					if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
						if vUserLocationsItem.Current+tmpCurrentReward >= vUserLocationsItem.CurrentMax { // 占位分红人分满停止
							vUserLocationsItem.Status = "stop"
							vUserLocationsItem.StopDate = time.Now().UTC().Add(8 * time.Hour)

							tmpCurrentReward = vUserLocationsItem.CurrentMax - vUserLocationsItem.Current
							bLocationRewardAmount = tmpCurrentReward * bPriceBase / bPrice
						}

						var tmpMaxNew int64
						if vUserLocationsItem.CurrentMaxNew < vUserLocationsItem.CurrentMax {
							tmpMaxNew = vUserLocationsItem.CurrentMax - vUserLocationsItem.CurrentMaxNew
						}

						if 0 < tmpCurrentReward && 0 < bLocationRewardAmount {
							err = uuc.locationRepo.UpdateLocationNewNew(ctx, vUserLocationsItem.ID, vUserLocationsItem.Status, tmpCurrentReward, tmpMaxNew, bLocationRewardAmount, vUserLocationsItem.StopDate) // 分红占位数据修改
							if nil != err {
								return err
							}

							_, err = uuc.ubRepo.AreaRewardBiw(ctx, vUserLocationsItem.UserId, bLocationRewardAmount, tmpCurrentReward, 3, vUserLocationsItem.Status, tmpMaxNew, feeRate)
							if nil != err {
								return err
							}
						}

						// 业绩减掉
						if "stop" == vUserLocationsItem.Status {
							tmpTop := vUserLocationsItem.Top
							tmpTopNum := vUserLocationsItem.TopNum
							for j := 0; j < 10000 && 0 < tmpTop && 0 < tmpTopNum; j++ {
								err = uuc.locationRepo.UpdateLocationNewTotalSub(ctx, tmpTop, tmpTopNum, vUserLocationsItem.Usdt/100000)
								if nil != err {
									return err
								}

								var (
									currentLocation *LocationNew
								)
								currentLocation, err = uuc.locationRepo.GetLocationById(ctx, tmpTop)
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
						fmt.Println("err reward daily three", err, vUserLocationsItem)
						continue
					}
				}
			}
		}
	}

	userLocationsFour := make([]*LocationNew, 0)
	userLocations, err = uuc.locationRepo.GetRunningLocations(ctx)
	if nil != err {
		return &v1.AdminDailyLocationRewardReply{}, nil
	}
	for _, vUserLocations := range userLocations {
		// 1大区
		if vUserLocations.Total >= vUserLocations.TotalTwo && vUserLocations.Total >= vUserLocations.TotalThree {
			if areaFour <= vUserLocations.TotalTwo+vUserLocations.TotalThree || 4 <= vUserLocations.LastLevel {
				userLocationsFour = append(userLocationsFour, vUserLocations)
			}
		} else if vUserLocations.TotalTwo >= vUserLocations.Total && vUserLocations.TotalTwo >= vUserLocations.TotalThree {
			if areaFour <= vUserLocations.Total+vUserLocations.TotalThree || 4 <= vUserLocations.LastLevel {
				userLocationsFour = append(userLocationsFour, vUserLocations)
			}
		} else if vUserLocations.TotalThree >= vUserLocations.Total && vUserLocations.TotalThree >= vUserLocations.TotalTwo {
			if areaFour <= vUserLocations.TotalTwo+vUserLocations.Total || 4 <= vUserLocations.LastLevel {
				userLocationsFour = append(userLocationsFour, vUserLocations)
			}
		}
	}

	if 0 < len(userLocationsFour) {
		rewardLocationYesFour := rewardLocationYes / 1000 * areaNumFour / int64(len(userLocationsFour))
		if 0 < rewardLocationYesFour {
			for _, vUserLocationsItem := range userLocationsFour {
				// 奖励
				tmpCurrentReward := rewardLocationYesFour
				bLocationRewardAmount := tmpCurrentReward * bPriceBase / bPrice

				if 0 < tmpCurrentReward && 0 < bLocationRewardAmount {
					if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
						if vUserLocationsItem.Current+tmpCurrentReward >= vUserLocationsItem.CurrentMax { // 占位分红人分满停止
							vUserLocationsItem.Status = "stop"
							vUserLocationsItem.StopDate = time.Now().UTC().Add(8 * time.Hour)

							tmpCurrentReward = vUserLocationsItem.CurrentMax - vUserLocationsItem.Current
							bLocationRewardAmount = tmpCurrentReward * bPriceBase / bPrice
						}

						var tmpMaxNew int64
						if vUserLocationsItem.CurrentMaxNew < vUserLocationsItem.CurrentMax {
							tmpMaxNew = vUserLocationsItem.CurrentMax - vUserLocationsItem.CurrentMaxNew
						}

						if 0 < tmpCurrentReward && 0 < bLocationRewardAmount {
							err = uuc.locationRepo.UpdateLocationNewNew(ctx, vUserLocationsItem.ID, vUserLocationsItem.Status, tmpCurrentReward, tmpMaxNew, bLocationRewardAmount, vUserLocationsItem.StopDate) // 分红占位数据修改
							if nil != err {
								return err
							}

							_, err = uuc.ubRepo.AreaRewardBiw(ctx, vUserLocationsItem.UserId, bLocationRewardAmount, tmpCurrentReward, 4, vUserLocationsItem.Status, tmpMaxNew, feeRate)
							if nil != err {
								return err
							}
						}

						// 业绩减掉
						if "stop" == vUserLocationsItem.Status {
							tmpTop := vUserLocationsItem.Top
							tmpTopNum := vUserLocationsItem.TopNum
							for j := 0; j < 10000 && 0 < tmpTop && 0 < tmpTopNum; j++ {
								err = uuc.locationRepo.UpdateLocationNewTotalSub(ctx, tmpTop, tmpTopNum, vUserLocationsItem.Usdt/100000)
								if nil != err {
									return err
								}

								var (
									currentLocation *LocationNew
								)
								currentLocation, err = uuc.locationRepo.GetLocationById(ctx, tmpTop)
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
						fmt.Println("err reward daily three", err, vUserLocationsItem)
						continue
					}
				}
			}
		}
	}

	userLocationsFive := make([]*LocationNew, 0)
	userLocations, err = uuc.locationRepo.GetRunningLocations(ctx)
	if nil != err {
		return &v1.AdminDailyLocationRewardReply{}, nil
	}
	for _, vUserLocations := range userLocations {
		// 1大区
		if vUserLocations.Total >= vUserLocations.TotalTwo && vUserLocations.Total >= vUserLocations.TotalThree {
			if areaFive <= vUserLocations.TotalTwo+vUserLocations.TotalThree || 5 <= vUserLocations.LastLevel {
				userLocationsFive = append(userLocationsFive, vUserLocations)
			}
		} else if vUserLocations.TotalTwo >= vUserLocations.Total && vUserLocations.TotalTwo >= vUserLocations.TotalThree {
			if areaFive <= vUserLocations.Total+vUserLocations.TotalThree || 5 <= vUserLocations.LastLevel {
				userLocationsFive = append(userLocationsFive, vUserLocations)
			}
		} else if vUserLocations.TotalThree >= vUserLocations.Total && vUserLocations.TotalThree >= vUserLocations.TotalTwo {
			if areaFive <= vUserLocations.TotalTwo+vUserLocations.Total || 5 <= vUserLocations.LastLevel {
				userLocationsFive = append(userLocationsFive, vUserLocations)
			}
		}
	}

	if 0 < len(userLocationsFive) {
		rewardLocationYesFive := rewardLocationYes / 1000 * areaNumFive / int64(len(userLocationsFive))
		if 0 < rewardLocationYesFive {
			for _, vUserLocationsItem := range userLocationsFive {
				// 奖励
				tmpCurrentReward := rewardLocationYesFive
				bLocationRewardAmount := tmpCurrentReward * bPriceBase / bPrice

				if 0 < tmpCurrentReward && 0 < bLocationRewardAmount {
					if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
						if vUserLocationsItem.Current+tmpCurrentReward >= vUserLocationsItem.CurrentMax { // 占位分红人分满停止
							vUserLocationsItem.Status = "stop"
							vUserLocationsItem.StopDate = time.Now().UTC().Add(8 * time.Hour)

							tmpCurrentReward = vUserLocationsItem.CurrentMax - vUserLocationsItem.Current
							bLocationRewardAmount = tmpCurrentReward * bPriceBase / bPrice
						}

						var tmpMaxNew int64
						if vUserLocationsItem.CurrentMaxNew < vUserLocationsItem.CurrentMax {
							tmpMaxNew = vUserLocationsItem.CurrentMax - vUserLocationsItem.CurrentMaxNew
						}

						if 0 < tmpCurrentReward && 0 < bLocationRewardAmount {
							err = uuc.locationRepo.UpdateLocationNewNew(ctx, vUserLocationsItem.ID, vUserLocationsItem.Status, tmpCurrentReward, tmpMaxNew, bLocationRewardAmount, vUserLocationsItem.StopDate) // 分红占位数据修改
							if nil != err {
								return err
							}

							_, err = uuc.ubRepo.AreaRewardBiw(ctx, vUserLocationsItem.UserId, bLocationRewardAmount, tmpCurrentReward, 5, vUserLocationsItem.Status, tmpMaxNew, feeRate)
							if nil != err {
								return err
							}
						}

						// 业绩减掉
						if "stop" == vUserLocationsItem.Status {
							tmpTop := vUserLocationsItem.Top
							tmpTopNum := vUserLocationsItem.TopNum
							for j := 0; j < 10000 && 0 < tmpTop && 0 < tmpTopNum; j++ {
								err = uuc.locationRepo.UpdateLocationNewTotalSub(ctx, tmpTop, tmpTopNum, vUserLocationsItem.Usdt/100000)
								if nil != err {
									return err
								}

								var (
									currentLocation *LocationNew
								)
								currentLocation, err = uuc.locationRepo.GetLocationById(ctx, tmpTop)
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
						fmt.Println("err reward daily three", err, vUserLocationsItem)
						continue
					}
				}
			}
		}
	}

	// 全网前天
	userLocationsBef, err = uuc.locationRepo.GetLocationDailyYesterday(ctx, day-1)
	for _, userLocationBef := range userLocationsBef {
		rewardLocationBef += userLocationBef.Usdt
	}
	// 全球
	totalReward := rewardLocationYes/100/100*70*total + rewardLocationBef/100/100*30*total

	var (
		fourUserRecommendTotal map[int64]int64
	)

	fourUserRecommendTotal = make(map[int64]int64, 0)
	for _, userLocationYes := range userLocationsYes {
		// 获取直推

		var (
			fourUserRecommend         *UserRecommend
			myFourUserRecommendUserId int64
			//myFourRecommendUser *User
		)
		fourUserRecommend, err = uuc.urRepo.GetUserRecommendByUserId(ctx, userLocationYes.UserId)
		if nil == fourUserRecommend {
			continue
		}

		if "" != fourUserRecommend.RecommendCode {
			tmpFourRecommendUserIds := strings.Split(fourUserRecommend.RecommendCode, "D")
			if 2 <= len(tmpFourRecommendUserIds) {
				myFourUserRecommendUserId, _ = strconv.ParseInt(tmpFourRecommendUserIds[len(tmpFourRecommendUserIds)-1], 10, 64) // 最后一位是直推人
			}
			//myFourRecommendUser, err = uuc.repo.GetUserById(ctx, myFourUserRecommendUserId)
			//if nil != err {
			//	return nil, err
			//}

			if _, ok := fourUserRecommendTotal[myFourUserRecommendUserId]; ok {
				fourUserRecommendTotal[myFourUserRecommendUserId] += userLocationYes.Usdt
			} else {
				fourUserRecommendTotal[myFourUserRecommendUserId] = userLocationYes.Usdt
			}
		}
	}

	if 0 >= len(fourUserRecommendTotal) {
		return &v1.AdminDailyLocationRewardReply{}, nil
	}

	// 前四名
	type KeyValuePair struct {
		Key   int64
		Value int64
	}
	var keyValuePairs []KeyValuePair
	for key, value := range fourUserRecommendTotal {
		keyValuePairs = append(keyValuePairs, KeyValuePair{key, value})
	}

	// 按值排序切片
	sort.Slice(keyValuePairs, func(i, j int) bool {
		return keyValuePairs[i].Value > keyValuePairs[j].Value
	})

	// 获取前四项
	var topFour []KeyValuePair
	if 4 <= len(keyValuePairs) {
		topFour = keyValuePairs[:4]
	} else {
		topFour = keyValuePairs[:len(keyValuePairs)]
	}

	for k, vTopFour := range topFour {
		var (
			tmpMyRecommendAmount int64
		)
		if 0 == k {
			tmpMyRecommendAmount = totalReward / 100 * one
		} else if 1 == k {
			tmpMyRecommendAmount = totalReward / 100 * two
		} else if 2 == k {
			tmpMyRecommendAmount = totalReward / 100 * three
		} else if 3 == k {
			tmpMyRecommendAmount = totalReward / 100 * four
		}

		if 0 >= tmpMyRecommendAmount {
			continue
		}

		if 0 < tmpMyRecommendAmount {
			if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
				_, err = uuc.ubRepo.FourRewardBiw(ctx, vTopFour.Key, tmpMyRecommendAmount, int64(k+1)) // 推荐人奖励
				if nil != err {
					return err
				}

				return nil
			}); nil != err {
				fmt.Println("err reward daily four", err, vTopFour)
				continue
			}
		}
	}

	return &v1.AdminDailyLocationRewardReply{}, nil
}

func (uuc *UserUseCase) AdminUpdateLocationNewMax(ctx context.Context, req *v1.AdminUpdateLocationNewMaxRequest) (*v1.AdminUpdateLocationNewMaxReply, error) {
	var (
		err error
	)
	res := &v1.AdminUpdateLocationNewMaxReply{}
	amountFloat, _ := strconv.ParseFloat(req.SendBody.Amount, 10)
	amountFloat *= 100000
	amount, _ := strconv.ParseInt(strconv.FormatFloat(amountFloat, 'f', -1, 64), 10, 64)

	_, err = uuc.ubRepo.UpdateLocationNewMax(ctx, req.SendBody.UserId, amount)

	if nil != err {
		return res, err
	}

	return nil, err
}

func (uuc *UserUseCase) AdminDailyLocationRewardNew(ctx context.Context, req *v1.AdminDailyLocationRewardNewRequest) (*v1.AdminDailyLocationRewardNewReply, error) {
	var (
		userLocations    []*LocationNew
		userLocationsMap map[int64]*LocationNew
		userLocations1   []*LocationNew
		userLocations2   []*LocationNew
		userLocations3   []*LocationNew
		v1r              int64
		v2r              int64
		v3r              int64
		configs          []*Config
		amount           int64
		amountV1         int64
		amountV2         int64
		amountV3         int64
		v1Count          int64
		v2Count          int64
		v3Count          int64
		err              error
		userInfos        []*UserInfo
	)

	configs, _ = uuc.configRepo.GetConfigByKeys(ctx,
		"v1", "v2", "v3",
	)

	if nil != configs {
		for _, vConfig := range configs {
			if "v1" == vConfig.KeyName {
				v1r, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "v2" == vConfig.KeyName {
				v2r, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "v3" == vConfig.KeyName {
				v3r, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
		}
	}

	amount, err = uuc.ubRepo.GetSystemWithdrawUsdtFeeTotalToday(ctx)
	if nil != err {
		return &v1.AdminDailyLocationRewardNewReply{}, nil
	}

	// 获取手动设置的
	userInfos, err = uuc.uiRepo.GetUserInfosByVipAndLockVip(ctx)
	if nil != err {
		return &v1.AdminDailyLocationRewardNewReply{}, nil
	}

	if nil != userInfos {
		for _, v := range userInfos {
			if 3 == v.Vip {
				v3Count += 1
			}

			if 2 == v.Vip {
				v2Count += 1
			}

			if 1 == v.Vip {
				v1Count += 1
			}
		}
	}

	userLocations, err = uuc.locationRepo.GetAllLocationsNew2(ctx)
	if nil != err {
		return &v1.AdminDailyLocationRewardNewReply{}, nil
	}

	userLocationsMap = make(map[int64]*LocationNew, 0)
	userLocations3 = make([]*LocationNew, 0)
	userLocations2 = make([]*LocationNew, 0)
	userLocations1 = make([]*LocationNew, 0)
	for _, vUserLocations := range userLocations {

		if _, ok := userLocationsMap[vUserLocations.UserId]; ok {
			continue
		}

		var (
			userInfo *UserInfo
		)

		userLocationsMap[vUserLocations.UserId] = vUserLocations

		userInfo, err = uuc.uiRepo.GetUserInfoByUserId(ctx, vUserLocations.UserId)
		if nil != err {
			continue
		}

		if 3 == userInfo.Vip {
			v3Count += 1
			userLocations3 = append(userLocations3, vUserLocations)
		}

		if 2 == userInfo.Vip {
			v2Count += 1
			userLocations2 = append(userLocations2, vUserLocations)
		}

		if 1 == userInfo.Vip {
			v1Count += 1
			userLocations1 = append(userLocations1, vUserLocations)
		}
	}

	if v1Count > 0 {
		amountV1 = amount * v1r / 100 / v1Count
	}

	if v2Count > 0 {
		amountV2 = amount * v2r / 100 / v2Count
	}

	if v3Count > 0 {
		amountV3 = amount * v3r / 100 / v3Count
	}

	if nil != userInfos {
		for _, v := range userInfos {
			var tmpAmount int64
			if 3 == v.Vip {
				tmpAmount = amountV3
			}

			if 2 == v.Vip {
				tmpAmount = amountV2
			}

			if 1 == v.Vip {
				tmpAmount = amountV1
			}
			if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error {
				_, err = uuc.ubRepo.NormalReward4(ctx, v.UserId, tmpAmount, 0)
				if nil != err {
					return err
				}
				return nil
			}); nil != err {
				continue
			}
		}
	}

	for _, vUserLocations1 := range userLocations1 {

		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			tmpStatus := vUserLocations1.Status // 现在还在运行中

			// 奖励usdt
			tmpRewardAmount := amountV1

			vUserLocations1.Status = "running"
			vUserLocations1.Current += tmpRewardAmount

			tmpRewardAmount2 := tmpRewardAmount
			if vUserLocations1.Current >= vUserLocations1.CurrentMax { // 占位分红人分满停止
				vUserLocations1.Status = "stop"
				if "running" == tmpStatus {
					vUserLocations1.StopDate = time.Now().UTC().Add(8 * time.Hour)
					tmpRewardAmount2 = tmpRewardAmount - (vUserLocations1.Current - vUserLocations1.CurrentMax)
				} else {
					tmpRewardAmount2 = 0
				}
			}

			if 0 < tmpRewardAmount {
				err = uuc.locationRepo.UpdateLocationNew2(ctx, vUserLocations1.ID, vUserLocations1.Status, tmpRewardAmount, vUserLocations1.StopDate) // 分红占位数据修改
				if nil != err {
					return err
				}
				_, err = uuc.ubRepo.NormalReward3(ctx, vUserLocations1.UserId, tmpRewardAmount, tmpRewardAmount2, vUserLocations1.ID, tmpStatus, vUserLocations1.Status) // 直推人奖励
				if nil != err {
					return err
				}
			}
			return nil
		}); nil != err {
			continue
		}
	}

	for _, vUserLocations3 := range userLocations3 {

		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			tmpStatus := vUserLocations3.Status // 现在还在运行中

			// 奖励usdt
			tmpRewardAmount := amountV3

			vUserLocations3.Status = "running"
			vUserLocations3.Current += tmpRewardAmount

			tmpRewardAmount2 := tmpRewardAmount
			if vUserLocations3.Current >= vUserLocations3.CurrentMax { // 占位分红人分满停止
				vUserLocations3.Status = "stop"
				if "running" == tmpStatus {
					vUserLocations3.StopDate = time.Now().UTC().Add(8 * time.Hour)
					tmpRewardAmount2 = tmpRewardAmount - (vUserLocations3.Current - vUserLocations3.CurrentMax)
				} else {
					tmpRewardAmount2 = 0
				}
			}

			if 0 < tmpRewardAmount {
				err = uuc.locationRepo.UpdateLocationNew2(ctx, vUserLocations3.ID, vUserLocations3.Status, tmpRewardAmount, vUserLocations3.StopDate) // 分红占位数据修改
				if nil != err {
					return err
				}
				_, err = uuc.ubRepo.NormalReward3(ctx, vUserLocations3.UserId, tmpRewardAmount, tmpRewardAmount2, vUserLocations3.ID, tmpStatus, vUserLocations3.Status) // 直推人奖励
				if nil != err {
					return err
				}
			}
			return nil
		}); nil != err {
			continue
		}
	}

	for _, vUserLocations2 := range userLocations2 {

		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			tmpStatus := vUserLocations2.Status // 现在还在运行中

			// 奖励usdt
			tmpRewardAmount := amountV2

			vUserLocations2.Status = "running"
			vUserLocations2.Current += tmpRewardAmount

			tmpRewardAmount2 := tmpRewardAmount
			if vUserLocations2.Current >= vUserLocations2.CurrentMax { // 占位分红人分满停止
				vUserLocations2.Status = "stop"
				if "running" == tmpStatus {
					vUserLocations2.StopDate = time.Now().UTC().Add(8 * time.Hour)
					tmpRewardAmount2 = tmpRewardAmount - (vUserLocations2.Current - vUserLocations2.CurrentMax)
				} else {
					tmpRewardAmount2 = 0
				}
			}

			if 0 < tmpRewardAmount {
				err = uuc.locationRepo.UpdateLocationNew2(ctx, vUserLocations2.ID, vUserLocations2.Status, tmpRewardAmount, vUserLocations2.StopDate) // 分红占位数据修改
				if nil != err {
					return err
				}
				_, err = uuc.ubRepo.NormalReward3(ctx, vUserLocations2.UserId, tmpRewardAmount, tmpRewardAmount2, vUserLocations2.ID, tmpStatus, vUserLocations2.Status) // 直推人奖励
				if nil != err {
					return err
				}
			}
			return nil
		}); nil != err {
			continue
		}
	}

	return &v1.AdminDailyLocationRewardNewReply{}, nil
}

func (uuc *UserUseCase) AdminDailyRecommendReward(ctx context.Context, req *v1.AdminDailyRecommendRewardRequest) (*v1.AdminDailyRecommendRewardReply, error) {

	var (
		users                  []*User
		userLocations          []*LocationNew
		configs                []*Config
		recommendAreaOne       int64
		recommendAreaOneRate   int64
		recommendAreaTwo       int64
		recommendAreaTwoRate   int64
		recommendAreaThree     int64
		recommendAreaThreeRate int64
		recommendAreaFour      int64
		recommendAreaFourRate  int64
		fee                    int64
		rewardRate             int64
		coinPrice              int64
		coinRewardRate         int64
		day                    = -1
		err                    error
	)

	if 1 == req.Day {
		day = 0
	}

	// 全网手续费
	userLocations, err = uuc.locationRepo.GetLocationDailyYesterday(ctx, day)
	if nil != err {
		return nil, err
	}
	for _, userLocation := range userLocations {
		fee += userLocation.CurrentMax * 100 / userLocation.OutRate
	}
	if 0 >= fee {
		return &v1.AdminDailyRecommendRewardReply{}, nil
	}

	configs, _ = uuc.configRepo.GetConfigByKeys(ctx, "recommend_area_one",
		"recommend_area_one_rate", "recommend_area_two_rate", "recommend_area_three_rate", "recommend_area_four_rate",
		"recommend_area_two", "recommend_area_three", "recommend_area_four", "coin_price", "coin_reward_rate", "reward_rate")
	if nil != configs {
		for _, vConfig := range configs {
			if "recommend_area_one" == vConfig.KeyName {
				recommendAreaOne, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "recommend_area_one_rate" == vConfig.KeyName {
				recommendAreaOneRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "recommend_area_two" == vConfig.KeyName {
				recommendAreaTwo, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "recommend_area_two_rate" == vConfig.KeyName {
				recommendAreaTwoRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "recommend_area_three" == vConfig.KeyName {
				recommendAreaThree, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "recommend_area_three_rate" == vConfig.KeyName {
				recommendAreaThreeRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "recommend_area_four" == vConfig.KeyName {
				recommendAreaFour, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "recommend_area_four_rate" == vConfig.KeyName {
				recommendAreaFourRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "coin_price" == vConfig.KeyName {
				coinPrice, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "coin_reward_rate" == vConfig.KeyName {
				coinRewardRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "reward_rate" == vConfig.KeyName {
				rewardRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
		}
	}

	users, err = uuc.repo.GetAllUsers(ctx)
	if nil != err {
		return nil, err
	}

	level1 := make(map[int64]int64, 0)
	level2 := make(map[int64]int64, 0)
	level3 := make(map[int64]int64, 0)
	level4 := make(map[int64]int64, 0)

	for _, user := range users {
		var userArea *UserArea
		userArea, err = uuc.urRepo.GetUserArea(ctx, user.ID)
		if nil != err {
			continue
		}

		if userArea.Level > 0 {
			if userArea.Level >= 1 {
				level1[user.ID] = user.ID
			}
			if userArea.Level >= 2 {
				level2[user.ID] = user.ID
			}
			if userArea.Level >= 3 {
				level3[user.ID] = user.ID
			}
			if userArea.Level >= 4 {
				level4[user.ID] = user.ID
			}
			continue
		}

		var userRecommend *UserRecommend
		userRecommend, err = uuc.urRepo.GetUserRecommendByUserId(ctx, user.ID)
		if nil != err {
			continue
		}

		// 伞下业绩
		var (
			myRecommendUsers   []*UserRecommend
			userAreas          []*UserArea
			maxAreaAmount      int64
			areaAmount         int64
			myRecommendUserIds []int64
		)
		myCode := userRecommend.RecommendCode + "D" + strconv.FormatInt(user.ID, 10)
		myRecommendUsers, err = uuc.urRepo.GetUserRecommendByCode(ctx, myCode)
		if nil == err {
			// 找直推
			for _, vMyRecommendUsers := range myRecommendUsers {
				myRecommendUserIds = append(myRecommendUserIds, vMyRecommendUsers.UserId)
			}
		}
		if 0 < len(myRecommendUserIds) {
			userAreas, err = uuc.urRepo.GetUserAreas(ctx, myRecommendUserIds)
			if nil == err {
				var (
					tmpTotalAreaAmount int64
				)
				for _, vUserAreas := range userAreas {
					tmpAreaAmount := vUserAreas.Amount + vUserAreas.SelfAmount
					tmpTotalAreaAmount += tmpAreaAmount
					if tmpAreaAmount > maxAreaAmount {
						maxAreaAmount = tmpAreaAmount
					}
				}

				areaAmount = tmpTotalAreaAmount - maxAreaAmount
			}
		}

		// 比较级别
		if areaAmount >= recommendAreaOne*100000 {
			level1[user.ID] = user.ID
		}

		if areaAmount >= recommendAreaTwo*100000 {
			level2[user.ID] = user.ID
		}

		if areaAmount >= recommendAreaThree*100000 {
			level3[user.ID] = user.ID
		}

		if areaAmount >= recommendAreaFour*100000 {
			level4[user.ID] = user.ID
		}
	}
	fmt.Println(level4, level3, level2, level1)
	// 分红
	fee /= 100000 // 这里多除五个0
	fmt.Println(fee)
	if 0 < len(level1) {
		feeLevel1 := fee * recommendAreaOneRate / 100 / int64(len(level1))
		feeLevel1 *= 100000

		for _, vLevel1 := range level1 {
			if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
				var myLocationLast *LocationNew
				// 获取当前用户的占位信息，已经有运行中的跳过
				myLocationLast, err = uuc.locationRepo.GetMyLocationLast(ctx, vLevel1)
				if nil == myLocationLast { // 无占位信息
					return err
				}

				feeLevel1Usdt := feeLevel1 * rewardRate / 100
				feeLevel1Coin := feeLevel1 * coinRewardRate / 100 * 1000 / coinPrice

				tmpCurrentStatus := myLocationLast.Status // 现在还在运行中
				myLocationLast.Status = "running"
				myLocationLast.Current += feeLevel1
				if myLocationLast.Current >= myLocationLast.CurrentMax { // 占位分红人分满停止
					if "running" == tmpCurrentStatus {
						myLocationLast.StopDate = time.Now().UTC().Add(8 * time.Hour)

						tmpLastAmount := feeLevel1 - (myLocationLast.Current - myLocationLast.CurrentMax)
						feeLevel1Usdt = tmpLastAmount * rewardRate / 100
						feeLevel1Coin = tmpLastAmount * coinRewardRate / 100 * 1000 / coinPrice
					}
					myLocationLast.Status = "stop"
				}

				if 0 < feeLevel1 {
					err = uuc.locationRepo.UpdateLocationNew(ctx, myLocationLast.ID, myLocationLast.Status, feeLevel1, myLocationLast.StopDate) // 分红占位数据修改
					if nil != err {
						return err
					}

					_, err = uuc.ubRepo.UserDailyRecommendArea(ctx, vLevel1, feeLevel1, feeLevel1Usdt, feeLevel1Coin, tmpCurrentStatus)
					if nil != err {
						return err
					}

				}

				return nil
			}); nil != err {
				continue
			}
		}
	}

	// 分红
	if 0 < len(level2) {
		feeLevel2 := fee * recommendAreaTwoRate / 100 / int64(len(level2))
		feeLevel2 *= 100000
		for _, vLevel2 := range level2 {
			if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
				var myLocationLast *LocationNew
				// 获取当前用户的占位信息，已经有运行中的跳过
				myLocationLast, err = uuc.locationRepo.GetMyLocationLast(ctx, vLevel2)
				if nil == myLocationLast { // 无占位信息
					return err
				}

				feeLevel2Usdt := feeLevel2 * rewardRate / 100
				feeLevel2Coin := feeLevel2 * coinRewardRate / 100 * 1000 / coinPrice

				tmpCurrentStatus := myLocationLast.Status // 现在还在运行中
				myLocationLast.Status = "running"
				myLocationLast.Current += feeLevel2
				if myLocationLast.Current >= myLocationLast.CurrentMax { // 占位分红人分满停止
					if "running" == tmpCurrentStatus {
						myLocationLast.StopDate = time.Now().UTC().Add(8 * time.Hour)

						tmpLastAmount := feeLevel2 - (myLocationLast.Current - myLocationLast.CurrentMax)
						feeLevel2Usdt = tmpLastAmount * rewardRate / 100
						feeLevel2Coin = tmpLastAmount * coinRewardRate / 100 * 1000 / coinPrice
					}
					myLocationLast.Status = "stop"
				}

				if 0 < feeLevel2 {
					err = uuc.locationRepo.UpdateLocationNew(ctx, myLocationLast.ID, myLocationLast.Status, feeLevel2, myLocationLast.StopDate) // 分红占位数据修改
					if nil != err {
						return err
					}

					_, err = uuc.ubRepo.UserDailyRecommendArea(ctx, vLevel2, feeLevel2, feeLevel2Usdt, feeLevel2Coin, tmpCurrentStatus)
					if nil != err {
						return err
					}

				}

				return nil
			}); nil != err {
				continue
			}
		}
	}

	// 分红
	if 0 < len(level3) {
		feeLevel3 := fee * recommendAreaThreeRate / 100 / int64(len(level3))
		feeLevel3 *= 100000
		for _, vLevel3 := range level3 {
			if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
				var myLocationLast *LocationNew
				// 获取当前用户的占位信息，已经有运行中的跳过
				myLocationLast, err = uuc.locationRepo.GetMyLocationLast(ctx, vLevel3)
				if nil == myLocationLast { // 无占位信息
					return err
				}

				feeLevel3Usdt := feeLevel3 * rewardRate / 100
				feeLevel3Coin := feeLevel3 * coinRewardRate / 100 * 1000 / coinPrice

				tmpCurrentStatus := myLocationLast.Status // 现在还在运行中
				myLocationLast.Status = "running"
				myLocationLast.Current += feeLevel3
				if myLocationLast.Current >= myLocationLast.CurrentMax { // 占位分红人分满停止
					if "running" == tmpCurrentStatus {
						myLocationLast.StopDate = time.Now().UTC().Add(8 * time.Hour)

						tmpLastAmount := feeLevel3 - (myLocationLast.Current - myLocationLast.CurrentMax)
						feeLevel3Usdt = tmpLastAmount * rewardRate / 100
						feeLevel3Coin = tmpLastAmount * coinRewardRate / 100 * 1000 / coinPrice
					}
					myLocationLast.Status = "stop"
				}

				if 0 < feeLevel3 {
					err = uuc.locationRepo.UpdateLocationNew(ctx, myLocationLast.ID, myLocationLast.Status, feeLevel3, myLocationLast.StopDate) // 分红占位数据修改
					if nil != err {
						return err
					}

					_, err = uuc.ubRepo.UserDailyRecommendArea(ctx, vLevel3, feeLevel3, feeLevel3Usdt, feeLevel3Coin, tmpCurrentStatus)
					if nil != err {
						return err
					}

				}

				return nil
			}); nil != err {
				continue
			}
		}
	}

	// 分红
	if 0 < len(level4) {
		feeLevel4 := fee * recommendAreaFourRate / 100 / int64(len(level4))
		feeLevel4 *= 100000
		for _, vLevel4 := range level4 {
			if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
				var myLocationLast *LocationNew
				// 获取当前用户的占位信息，已经有运行中的跳过
				myLocationLast, err = uuc.locationRepo.GetMyLocationLast(ctx, vLevel4)
				if nil == myLocationLast { // 无占位信息
					return err
				}

				feeLevel4Usdt := feeLevel4 * rewardRate / 100
				feeLevel4Coin := feeLevel4 * coinRewardRate / 100 * 1000 / coinPrice

				tmpCurrentStatus := myLocationLast.Status // 现在还在运行中
				myLocationLast.Status = "running"
				myLocationLast.Current += feeLevel4
				if myLocationLast.Current >= myLocationLast.CurrentMax { // 占位分红人分满停止
					if "running" == tmpCurrentStatus {
						myLocationLast.StopDate = time.Now().UTC().Add(8 * time.Hour)

						tmpLastAmount := feeLevel4 - (myLocationLast.Current - myLocationLast.CurrentMax)
						feeLevel4Usdt = tmpLastAmount * rewardRate / 100
						feeLevel4Coin = tmpLastAmount * coinRewardRate / 100 * 1000 / coinPrice

					}
					myLocationLast.Status = "stop"
				}

				if 0 < feeLevel4 {
					err = uuc.locationRepo.UpdateLocationNew(ctx, myLocationLast.ID, myLocationLast.Status, feeLevel4, myLocationLast.StopDate) // 分红占位数据修改
					if nil != err {
						return err
					}

					_, err = uuc.ubRepo.UserDailyRecommendArea(ctx, vLevel4, feeLevel4, feeLevel4Usdt, feeLevel4Coin, tmpCurrentStatus)
					if nil != err {
						return err
					}

				}

				return nil
			}); nil != err {
				continue
			}
		}
	}

	return &v1.AdminDailyRecommendRewardReply{}, nil
}

func (uuc *UserUseCase) CheckAndInsertRecommendArea(ctx context.Context, req *v1.CheckAndInsertRecommendAreaRequest) (*v1.CheckAndInsertRecommendAreaReply, error) {

	var (
		userRecommends         []*UserRecommend
		userRecommendAreaCodes []string
		userRecommendAreas     []*UserRecommendArea
		err                    error
	)
	userRecommends, err = uuc.urRepo.GetUserRecommends(ctx)
	if nil != err {
		return &v1.CheckAndInsertRecommendAreaReply{}, nil
	}

	for _, vUserRecommends := range userRecommends {
		tmp := vUserRecommends.RecommendCode + "D" + strconv.FormatInt(vUserRecommends.UserId, 10)
		tmpNoHas := true
		for k, vUserRecommendAreaCodes := range userRecommendAreaCodes {
			if strings.HasPrefix(vUserRecommendAreaCodes, tmp) {
				tmpNoHas = false
			} else if strings.HasPrefix(tmp, vUserRecommendAreaCodes) {
				userRecommendAreaCodes[k] = tmp
				tmpNoHas = false
			}
		}

		if tmpNoHas {
			userRecommendAreaCodes = append(userRecommendAreaCodes, tmp)
		}
	}

	userRecommendAreas = make([]*UserRecommendArea, 0)
	for _, vUserRecommendAreaCodes := range userRecommendAreaCodes {
		userRecommendAreas = append(userRecommendAreas, &UserRecommendArea{
			RecommendCode: vUserRecommendAreaCodes,
			Num:           int64(len(strings.Split(vUserRecommendAreaCodes, "D")) - 1),
		})
	}

	if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		_, err = uuc.urRepo.CreateUserRecommendArea(ctx, userRecommendAreas)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &v1.CheckAndInsertRecommendAreaReply{}, nil
}

func (uuc *UserUseCase) VipCheck(ctx context.Context, req *v1.VipCheckRequest) (*v1.VipCheckReply, error) {

	var (
		users           []*UserInfo
		configs         []*Config
		vip5Balance     int64
		vip4Balance     int64
		vip3Balance     int64
		vip2Balance     int64
		vip1Balance     int64
		vip0Balance     int64
		vip5BalanceTeam int64
		vip4BalanceTeam int64
		vip3BalanceTeam int64
		vip2BalanceTeam int64
		vip1BalanceTeam int64
		err             error
	)
	users, err = uuc.repo.GetAllUserInfos(ctx)
	if nil != err {
		return nil, err
	}

	configs, _ = uuc.configRepo.GetConfigByKeys(ctx, "vip_5_balance",
		"vip_4_balance", "vip_3_balance", "vip_2_balance", "vip_1_balance", "vip_0_balance",
		"vip_5_balance_team", "vip_4_balance_team", "vip_3_balance_team", "vip_2_balance_team", "vip_1_balance_team")
	if nil != configs {
		for _, vConfig := range configs {
			if "vip_5_balance" == vConfig.KeyName {
				vip5Balance, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "vip_4_balance" == vConfig.KeyName {
				vip4Balance, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "vip_3_balance" == vConfig.KeyName {
				vip3Balance, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "vip_2_balance" == vConfig.KeyName {
				vip2Balance, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "vip_0_balance" == vConfig.KeyName {
				vip0Balance, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "vip_1_balance" == vConfig.KeyName {
				vip1Balance, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "vip_4_balance_team" == vConfig.KeyName {
				vip4BalanceTeam, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "vip_3_balance_team" == vConfig.KeyName {
				vip3BalanceTeam, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "vip_2_balance_team" == vConfig.KeyName {
				vip2BalanceTeam, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "vip_1_balance_team" == vConfig.KeyName {
				vip1BalanceTeam, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "vip_5_balance_team" == vConfig.KeyName {
				vip5BalanceTeam, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
		}
	}

	for _, user := range users {
		if 0 < user.LockVip {
			continue
		}

		var (
			userRecommend  *UserRecommend
			userBalance    *UserBalance
			myCode         string
			teamCsdBalance int64
			myUserBalance  int64
			myVip          int64 = 0
		)

		vip1Count1 := make(map[int64]int64, 0)
		vip2Count1 := make(map[int64]int64, 0)
		vip3Count1 := make(map[int64]int64, 0)
		vip4Count1 := make(map[int64]int64, 0)

		userRecommend, err = uuc.urRepo.GetUserRecommendByUserId(ctx, user.ID)
		if nil != err {
			continue
		}

		userBalance, err = uuc.ubRepo.GetUserBalance(ctx, user.ID)
		if nil != err {
			continue
		}

		// 我的伞下所有用户
		myCode = userRecommend.RecommendCode + "D" + strconv.FormatInt(user.ID, 10)

		var (
			UserInfos             map[int64]*UserInfo
			userRecommends        []*UserRecommend
			userRecommendsUserIds []int64
		)

		userRecommends, err = uuc.urRepo.GetUserRecommendByCode(ctx, myCode)
		if nil == err {
			for _, vUserRecommends := range userRecommends {
				userRecommendsUserIds = append(userRecommendsUserIds, vUserRecommends.UserId)
			}
		}
		if 0 < len(userRecommendsUserIds) {
			UserInfos, err = uuc.uiRepo.GetUserInfoByUserIds(ctx, userRecommendsUserIds...)
		}
		for _, vUserInfos := range UserInfos {
			if 2 == vUserInfos.Vip {
				vip1Count1[vUserInfos.UserId] += 1
			} else if 3 == vUserInfos.Vip {
				vip2Count1[vUserInfos.UserId] += 1
			} else if 4 == vUserInfos.Vip {
				vip3Count1[vUserInfos.UserId] += 1
			} else if 5 == vUserInfos.Vip {
				vip4Count1[vUserInfos.UserId] += 1
			}
		}

		if 0 < len(userRecommends) {
			for _, vUserRecommendsQ := range userRecommends {

				var (
					userRecommends1        []*UserRecommend
					userRecommendsUserIds1 []int64
				)
				myCode1 := vUserRecommendsQ.RecommendCode + "D" + strconv.FormatInt(vUserRecommendsQ.UserId, 10)
				userRecommends1, err = uuc.urRepo.GetUserRecommendLikeCode(ctx, myCode1)
				if nil == err {
					for _, vUserRecommends1 := range userRecommends1 {
						userRecommendsUserIds1 = append(userRecommendsUserIds1, vUserRecommends1.UserId)
					}
				}

				var UserInfos1 map[int64]*UserInfo
				if 0 < len(userRecommendsUserIds1) {
					UserInfos1, err = uuc.uiRepo.GetUserInfoByUserIds(ctx, userRecommendsUserIds1...)
				}

				for _, vUserInfos1 := range UserInfos1 {
					if 2 == vUserInfos1.Vip {
						vip1Count1[vUserRecommendsQ.UserId] += 1
					} else if 3 == vUserInfos1.Vip {
						vip2Count1[vUserRecommendsQ.UserId] += 1
					} else if 4 == vUserInfos1.Vip {
						vip3Count1[vUserRecommendsQ.UserId] += 1
					} else if 5 == vUserInfos1.Vip {
						vip4Count1[vUserRecommendsQ.UserId] += 1
					}
				}
			}
		}

		var (
			vip1Count int64
			vip2Count int64
			vip3Count int64
			vip4Count int64
		)
		for _, vv1 := range vip1Count1 {
			if vv1 > 0 {
				vip1Count++
			}
		}
		for _, vv2 := range vip2Count1 {
			if vv2 > 0 {
				vip2Count++
			}
		}
		for _, vv3 := range vip3Count1 {
			if vv3 > 0 {
				vip3Count++
			}
		}
		for _, vv4 := range vip4Count1 {
			if vv4 > 0 {
				vip4Count++
			}
		}

		teamCsdBalance = user.TeamCsdBalance / 100000
		myUserBalance = userBalance.BalanceUsdt / 100000
		if teamCsdBalance >= vip5BalanceTeam && 2 <= vip4Count && 5 <= user.HistoryRecommend && myUserBalance >= vip5Balance {
			myVip = 6
		} else if teamCsdBalance >= vip4BalanceTeam && 2 <= vip3Count && 5 <= user.HistoryRecommend && myUserBalance >= vip4Balance {
			myVip = 5
		} else if teamCsdBalance >= vip3BalanceTeam && 2 <= vip2Count && 5 <= user.HistoryRecommend && myUserBalance >= vip3Balance {
			myVip = 4
		} else if teamCsdBalance >= vip2BalanceTeam && 2 <= vip1Count && 5 <= user.HistoryRecommend && myUserBalance >= vip2Balance {
			myVip = 3
		} else if teamCsdBalance >= vip1BalanceTeam && 5 <= user.HistoryRecommend && myUserBalance >= vip1Balance {
			myVip = 2
		} else if myUserBalance >= vip0Balance {
			myVip = 1
		}

		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务

			// 修改用户推荐人区数据，修改自身区数据
			_, err = uuc.uiRepo.UpdateUserInfoVip(ctx, user.ID, myVip)
			if nil != err {
				return err
			}

			return nil
		}); err != nil {
			return nil, err
		}
	}

	return &v1.VipCheckReply{}, nil
}

func (uuc *UserUseCase) CheckAdminUserArea(ctx context.Context, req *v1.CheckAdminUserAreaRequest) (*v1.CheckAdminUserAreaReply, error) {
	return &v1.CheckAdminUserAreaReply{}, nil
}

func (uuc *UserUseCase) CheckAndInsertLocationsRecommendUser(ctx context.Context, req *v1.CheckAndInsertLocationsRecommendUserRequest) (*v1.CheckAndInsertLocationsRecommendUserReply, error) {
	return &v1.CheckAndInsertLocationsRecommendUserReply{}, nil
}
