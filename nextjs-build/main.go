package main

import (
	"context"
)

type NextjsBuild struct {}

func (m *NextjsBuild) MyFunction(ctx context.Context, stringArg string) (*Container, error) {
	return dag.Container().From("alpine:latest").WithExec([]string{"echo", stringArg}).Sync(ctx)
}
