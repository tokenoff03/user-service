package interceptor

import (
	"context"
	"user-service/internal/logger"

	"github.com/pkg/errors"
	"github.com/tokenoff03/lib_ad1lek/pkg/sys"
	"github.com/tokenoff03/lib_ad1lek/pkg/sys/codes"
	"github.com/tokenoff03/lib_ad1lek/pkg/sys/validate"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	grpcCodes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GRPCStatusInterface interface {
	GRPCStatus() *status.Status
}

func ErrorCodesInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (res interface{}, err error) {
	res, err = handler(ctx, req)
	if nil == err {
		return res, nil
	}

	logger.Error("Error from interceptor error", zap.Error(err))

	switch {
	case sys.IsCommonError(err):
		commEr := sys.GetCommonError(err)
		code := toGRPCCode(commEr.Code()) //сommEr.Code() по сути возвращает grpc коды точь в точ. Свой кастомный туда не добавлял

		err = status.Error(code, commEr.Error())

	case validate.IsValidationError(err):
		err = status.Error(grpcCodes.InvalidArgument, err.Error())

	default:
		var se GRPCStatusInterface //ошибки с свервисов grpc
		if errors.As(err, &se) {   //Проверка на эти ошибки что они с сервиса
			return nil, se.GRPCStatus().Err()
		} else {
			if errors.Is(err, context.DeadlineExceeded) {
				err = status.Error(grpcCodes.DeadlineExceeded, err.Error())
			} else if errors.Is(err, context.Canceled) {
				err = status.Error(grpcCodes.Canceled, err.Error())
			} else {
				err = status.Error(grpcCodes.Internal, "internal error") //Если мы там забыли обернуть ошибку с бд или т.д то просто возрващается internal
			}
		}
	}

	return res, err
}

// Это конвертер для grpc кода, если мы используем свой кастомный код то можем переписать!
func toGRPCCode(code codes.Code) grpcCodes.Code {
	var res grpcCodes.Code

	switch code {
	case codes.OK:
		res = grpcCodes.OK
	case codes.Canceled:
		res = grpcCodes.Canceled
	case codes.InvalidArgument:
		res = grpcCodes.InvalidArgument
	case codes.DeadlineExceeded:
		res = grpcCodes.DeadlineExceeded
	case codes.NotFound:
		res = grpcCodes.NotFound
	case codes.AlreadyExists:
		res = grpcCodes.AlreadyExists
	case codes.PermissionDenied:
		res = grpcCodes.PermissionDenied
	case codes.ResourceExhausted:
		res = grpcCodes.ResourceExhausted
	case codes.FailedPrecondition:
		res = grpcCodes.FailedPrecondition
	case codes.Aborted:
		res = grpcCodes.Aborted
	case codes.OutOfRange:
		res = grpcCodes.OutOfRange
	case codes.Unimplemented:
		res = grpcCodes.Unimplemented
	case codes.Internal:
		res = grpcCodes.Internal
	case codes.Unavailable:
		res = grpcCodes.Unavailable
	case codes.DataLoss:
		res = grpcCodes.DataLoss
	case codes.Unauthenticated:
		res = grpcCodes.Unauthenticated
	default:
		res = grpcCodes.Unknown
	}

	return res
}
