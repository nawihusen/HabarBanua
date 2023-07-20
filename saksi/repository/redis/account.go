package redis

import (
	"be-service-saksi-management/domain"

	"github.com/go-redis/redis/v8"
)

type redisAccountRepository struct {
	Conn *redis.Client
}

// NewRedisAccountRepository is constructor of Redis repository
func NewRedisAccountRepository(Conn *redis.Client) domain.AccountRedisRepository {
	return &redisAccountRepository{Conn}
}

// // SetSession is save to Redis if any user login
// func (r *redisAccountRepository) SetSession(ctx context.Context, token string, add domain.SaksiAdditionalData) (err error) {
// 	var params []interface{}
// 	params = append(params, "recruiter_id", add.RecruiterID, "saksi_id", add.SaksiID, "coordinator_id", add.CoordinatorID, "tps_no", add.TpsNo, "attendance", add.Attendance, "recruiter_role", add.RoleRecruiter, "saksi_role", add.RoleSaksi, "coordinator_role", add.RoleCoordinator)
// 	_, err = r.Conn.HSet(ctx, token, params...).Result()

// 	// _, err = r.Conn.Set(ctx, token, role, 0).Result()
// 	return
// }

// // GetSession is get role from Redis if still have session
// func (r *redisAccountRepository) GetSession(ctx context.Context, token string) (add domain.SaksiAdditionalData, err error) {
// 	data := r.Conn.HGetAll(ctx, token)
// 	res, err := data.Result()
// 	if err != nil {
// 		return
// 	}

// 	if len(res) == 0 {
// 		err = errors.New("not found")
// 		return
// 	}

// 	err = data.Scan(&add)
// 	if err != nil {
// 		return
// 	}

// 	// role, err = r.Conn.Get(ctx, token).Result()
// 	return
// }

// // DeleteSession id delete session (key) if have session no more from auth
// func (r *redisAccountRepository) DeleteSession(ctx context.Context, token string) (err error) {
// 	_, err = r.Conn.Del(ctx, token).Result()

// 	return
// }

// // DeleteSession id delete session (key) if have session no more from auth
// func (r *redisAccountRepository) AttendanceSession(ctx context.Context, token string) (err error) {
// 	_, err = r.Conn.HSet(ctx, token, "attendance", true).Result()

// 	return
// }
