package server

import (
	"context"

	"github.com/micro/go-micro/v3/router"
	"github.com/micro/micro/v3/service/errors"
	pb "github.com/micro/micro/v3/service/router/proto"
)

type Table struct {
	Router router.Router
}

func (t *Table) Create(ctx context.Context, route *pb.Route, resp *pb.CreateResponse) error {
	err := t.Router.Table().Create(router.Route{
		Service:  route.Service,
		Address:  route.Address,
		Gateway:  route.Gateway,
		Network:  route.Network,
		Router:   route.Router,
		Link:     route.Link,
		Metric:   route.Metric,
		Metadata: route.Metadata,
	})
	if err != nil {
		return errors.InternalServerError("router.Table.Create", "failed to create route: %s", err)
	}

	return nil
}

func (t *Table) Update(ctx context.Context, route *pb.Route, resp *pb.UpdateResponse) error {
	err := t.Router.Table().Update(router.Route{
		Service:  route.Service,
		Address:  route.Address,
		Gateway:  route.Gateway,
		Network:  route.Network,
		Router:   route.Router,
		Link:     route.Link,
		Metric:   route.Metric,
		Metadata: route.Metadata,
	})
	if err != nil {
		return errors.InternalServerError("router.Table.Update", "failed to update route: %s", err)
	}

	return nil
}

func (t *Table) Delete(ctx context.Context, route *pb.Route, resp *pb.DeleteResponse) error {
	err := t.Router.Table().Delete(router.Route{
		Service:  route.Service,
		Address:  route.Address,
		Gateway:  route.Gateway,
		Network:  route.Network,
		Router:   route.Router,
		Link:     route.Link,
		Metric:   route.Metric,
		Metadata: route.Metadata,
	})
	if err != nil {
		return errors.InternalServerError("route.Table.Delete", "failed to delete route: %s", err)
	}

	return nil
}

// List returns all routes in the routing table
func (t *Table) List(ctx context.Context, req *pb.Request, resp *pb.ListResponse) error {
	routes, err := t.Router.Table().List()
	if err != nil {
		return errors.InternalServerError("router.Table.List", "failed to list routes: %s", err)
	}

	respRoutes := make([]*pb.Route, 0, len(routes))
	for _, route := range routes {
		respRoute := &pb.Route{
			Service:  route.Service,
			Address:  route.Address,
			Gateway:  route.Gateway,
			Network:  route.Network,
			Router:   route.Router,
			Link:     route.Link,
			Metric:   route.Metric,
			Metadata: route.Metadata,
		}
		respRoutes = append(respRoutes, respRoute)
	}

	resp.Routes = respRoutes

	return nil
}

func (t *Table) Query(ctx context.Context, req *pb.QueryRequest, resp *pb.QueryResponse) error {
	routes, err := t.Router.Table().Query(req.Service)
	if err != nil {
		return errors.InternalServerError("router.Table.Query", "failed to lookup routes: %s", err)
	}

	respRoutes := make([]*pb.Route, 0, len(routes))
	for _, route := range routes {
		respRoute := &pb.Route{
			Service:  route.Service,
			Address:  route.Address,
			Gateway:  route.Gateway,
			Network:  route.Network,
			Router:   route.Router,
			Link:     route.Link,
			Metric:   route.Metric,
			Metadata: route.Metadata,
		}
		respRoutes = append(respRoutes, respRoute)
	}

	resp.Routes = respRoutes

	return nil
}
