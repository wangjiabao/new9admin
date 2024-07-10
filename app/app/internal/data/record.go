package data

import (
	"context"
	"dhb/app/app/internal/biz"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"time"
)

type EthUserRecord struct {
	ID        int64     `gorm:"primarykey;type:int"`
	Hash      string    `gorm:"type:varchar(100);not null"`
	UserId    int64     `gorm:"type:int;not null"`
	Status    string    `gorm:"type:varchar(45);not null"`
	Type      string    `gorm:"type:varchar(45);not null"`
	Amount    string    `gorm:"type:varchar(45);not null"`
	AmountTwo uint64    `gorm:"type:bigint;not null"`
	CoinType  string    `gorm:"type:varchar(45);not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
	Last      int64     `gorm:"type:int;not null"`
}

type EthUserRecordRepo struct {
	data *Data
	log  *log.Helper
}

func NewEthUserRecordRepo(data *Data, logger log.Logger) biz.EthUserRecordRepo {
	return &EthUserRecordRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (e *EthUserRecordRepo) GetEthUserRecordListByHash(ctx context.Context, hash ...string) (map[string]*biz.EthUserRecord, error) {
	var ethUserRecord []*EthUserRecord
	if err := e.data.DB(ctx).Table("eth_user_record").Where("hash IN (?)", hash).Find(&ethUserRecord).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("USER_RECOMMEND_NOT_FOUND", "user recommend not found")
		}

		return nil, errors.New(500, "USER RECOMMEND ERROR", err.Error())
	}

	res := make(map[string]*biz.EthUserRecord, 0)
	for _, item := range ethUserRecord {
		res[item.Hash] = &biz.EthUserRecord{
			ID:       item.ID,
			UserId:   item.UserId,
			Hash:     item.Hash,
			Status:   item.Status,
			Type:     item.Type,
			Amount:   item.Amount,
			CoinType: item.CoinType,
		}
	}

	return res, nil
}

func (e *EthUserRecordRepo) GetEthUserRecordListByHash2(ctx context.Context, hash ...string) (map[string]*biz.EthUserRecord, error) {
	var ethUserRecord []*EthUserRecord
	if err := e.data.DB(ctx).Table("eth_user_record_2").Where("hash IN (?)", hash).Find(&ethUserRecord).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("USER_RECOMMEND_NOT_FOUND", "user recommend not found")
		}

		return nil, errors.New(500, "USER RECOMMEND ERROR", err.Error())
	}

	res := make(map[string]*biz.EthUserRecord, 0)
	for _, item := range ethUserRecord {
		res[item.Hash] = &biz.EthUserRecord{
			ID:       item.ID,
			UserId:   item.UserId,
			Hash:     item.Hash,
			Status:   item.Status,
			Type:     item.Type,
			Amount:   item.Amount,
			CoinType: item.CoinType,
		}
	}

	return res, nil
}

func (e *EthUserRecordRepo) GetEthUserRecordLast(ctx context.Context) (int64, error) {
	var ethUserRecord *EthUserRecord
	if err := e.data.DB(ctx).Table("eth_user_record").Order("last desc").First(&ethUserRecord).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, nil
		}

		return -1, errors.New(500, "USER RECOMMEND ERROR", err.Error())
	}

	return ethUserRecord.Last, nil
}

func (e *EthUserRecordRepo) GetEthUserRecordLast2(ctx context.Context) (int64, error) {
	var ethUserRecord *EthUserRecord
	if err := e.data.DB(ctx).Table("eth_user_record_2").Order("last desc").First(&ethUserRecord).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, nil
		}

		return -1, errors.New(500, "USER RECOMMEND ERROR", err.Error())
	}

	return ethUserRecord.Last, nil
}

func (e *EthUserRecordRepo) CreateEthUserRecordListByHash(ctx context.Context, r *biz.EthUserRecord) (*biz.EthUserRecord, error) {
	var ethUserRecord EthUserRecord
	ethUserRecord.UserId = r.UserId
	ethUserRecord.Hash = r.Hash
	ethUserRecord.Type = r.Type
	ethUserRecord.Status = r.Status
	ethUserRecord.Amount = r.Amount
	ethUserRecord.AmountTwo = r.AmountTwo
	ethUserRecord.CoinType = r.CoinType
	ethUserRecord.Last = r.Last

	res := e.data.DB(ctx).Table("eth_user_record").Create(&ethUserRecord)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_ETH_USER_RECORD_ERROR", "以太坊交易信息创建失败")
	}

	return &biz.EthUserRecord{
		ID:       ethUserRecord.ID,
		UserId:   ethUserRecord.UserId,
		Hash:     ethUserRecord.Hash,
		Status:   ethUserRecord.Status,
		Type:     ethUserRecord.Type,
		Amount:   ethUserRecord.Amount,
		CoinType: ethUserRecord.CoinType,
		Last:     ethUserRecord.Last,
	}, nil
}

func (e *EthUserRecordRepo) CreateEthUserRecordListByHash2(ctx context.Context, r *biz.EthUserRecord) (*biz.EthUserRecord, error) {
	var ethUserRecord EthUserRecord
	ethUserRecord.UserId = r.UserId
	ethUserRecord.Hash = r.Hash
	ethUserRecord.Type = r.Type
	ethUserRecord.Status = r.Status
	ethUserRecord.Amount = r.Amount
	ethUserRecord.CoinType = r.CoinType
	ethUserRecord.Last = r.Last

	res := e.data.DB(ctx).Table("eth_user_record_2").Create(&ethUserRecord)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_ETH_USER_RECORD_ERROR", "以太坊交易信息创建失败")
	}

	return &biz.EthUserRecord{
		ID:       ethUserRecord.ID,
		UserId:   ethUserRecord.UserId,
		Hash:     ethUserRecord.Hash,
		Status:   ethUserRecord.Status,
		Type:     ethUserRecord.Type,
		Amount:   ethUserRecord.Amount,
		CoinType: ethUserRecord.CoinType,
		Last:     ethUserRecord.Last,
	}, nil
}
