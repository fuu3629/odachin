package presentation

import (
	"context"

	"connectrpc.com/connect"
	"github.com/fuu3629/odachin/apps/service/gen/v1/odachin"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ServerStruct) ApplicateUsage(
	ctx context.Context,
	req *connect.Request[odachin.ApplicateUsageRequest],
) (*connect.Response[emptypb.Empty], error) {
	err := s.usageUsecase.ApplicateUsage(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (s *ServerStruct) GetUsageCategories(
	ctx context.Context,
	req *connect.Request[emptypb.Empty],
) (*connect.Response[odachin.GetUsageCategoriesResponse], error) {
	categories, err := s.usageUsecase.GetUsageCategories(ctx)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&odachin.GetUsageCategoriesResponse{
		Categories: categories,
	}), nil
}

func (s *ServerStruct) ApproveUsage(
	ctx context.Context,
	req *connect.Request[odachin.ApproveUsageRequest],
) (*connect.Response[emptypb.Empty], error) {
	err := s.usageUsecase.ApproveUsage(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (s *ServerStruct) GetUsageApplication(
	ctx context.Context,
	req *connect.Request[odachin.GetUsageApplicationRequest],
) (*connect.Response[odachin.GetUsageApplicationResponse], error) {
	usages, err := s.usageUsecase.GetUsageApplication(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&odachin.GetUsageApplicationResponse{
		UsageApplications: usages,
	}), nil
}

func (s *ServerStruct) GetUsageSummary(
	ctx context.Context,
	req *connect.Request[emptypb.Empty],
) (*connect.Response[odachin.GetUsageSummaryResponse], error) {
	summary, summary_monthly, err := s.usageUsecase.GetUsageSummary(ctx)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&odachin.GetUsageSummaryResponse{
		UsageSummaries:        summary,
		UsageSummariesMonthly: summary_monthly,
	}), nil
}

func (s *ServerStruct) RejectUsage(
	ctx context.Context,
	req *connect.Request[odachin.RejectUsageRequest],
) (*connect.Response[emptypb.Empty], error) {
	err := s.usageUsecase.RejectUsage(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}
